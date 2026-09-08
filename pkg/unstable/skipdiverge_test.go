package unstable

import (
	"strings"
	"testing"
)

// TestSkipPathsDivergeOnMalformed pins the three ways skipContainerFast and
// skipObject/skipArray disagree, so neither side drifts silently — the same job
// TestValidDivergesFromStdlib does for Valid's deliberate disagreements with
// encoding/json.
//
// The divergences are documented in skipfast.go's header and are all confined to
// malformed input: on a bracket-balanced value the two paths return the identical
// end index, which the differential fuzz and the boundary corpus already lock.
// This test exists because a comment claiming "the SIMD and scalar paths never
// diverge (including on malformed input)" survived in the tree for a long time
// while being false, and because the depth case below is true only by virtue of a
// bound added elsewhere (skip.go) — exactly the kind of cross-file claim that rots.
//
// A failure here is not automatically a bug. It means the relationship changed:
// either a path was fixed (delete the case and say so) or one drifted (fix it).
func TestSkipPathsDivergeOnMalformed(t *testing.T) {
	if !fastSkipAvail {
		t.Skip("no SIMD skip path on this build/CPU; there is nothing to diverge from")
	}

	scalar := func(data []byte) (int, error) {
		if data[0] == '{' {
			return skipObject(data, 0)
		}
		return skipArray(data, 0)
	}

	type want struct {
		end int
		err error
	}
	cases := []struct {
		name       string
		in         string
		fast, slow want
		why        string
	}{{
		// Class 1: an unbalanced bracket of the *other* type. The block loop is
		// told which pair to count and never sees the '['; skipObject descends
		// into it via skipArray and keeps looking for the inner close.
		name: "other-type bracket, truncated",
		in:   `{"a":[}` + strings.Repeat(" ", 80),
		fast: want{7, nil},
		slow: want{87, ErrTruncated},
		why:  "fast accepts, scalar runs off the end",
	}, {
		name: "other-type bracket, balanced outer",
		in:   `{"a":[}]}`,
		fast: want{7, nil},
		slow: want{9, nil},
		why:  "both accept, different end index",
	}, {
		// Class 3: a stray backslash outside a string. findEscaped64 is pure bit
		// math on the backslash bitmap and cannot know the backslash is not in a
		// string, so it masks the next byte out of the quote bitmap. Note the
		// direction is the opposite of class 1 here: the fast path rejects.
		name: "stray backslash, fast rejects",
		in:   `{"a":1,\"b":2,"pad":"` + strings.Repeat("x", 70) + `"}`,
		fast: want{93, ErrTruncated},
		slow: want{93, nil},
		why:  "fast rejects what scalar accepts",
	}, {
		name: "stray backslash, fast accepts",
		in:   `{\"}` + strings.Repeat(" ", 80),
		fast: want{4, nil},
		slow: want{84, ErrTruncated},
		why:  "fast accepts what scalar rejects",
	}}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			data := []byte(c.in)
			open := data[0]
			gotEnd, gotErr := skipContainerFast(data, 0, open)
			if gotEnd != c.fast.end || gotErr != c.fast.err {
				t.Errorf("skipContainerFast = (%d, %v), want (%d, %v) [%s]", gotEnd, gotErr, c.fast.end, c.fast.err, c.why)
			}
			gotEnd, gotErr = scalar(data)
			if gotEnd != c.slow.end || gotErr != c.slow.err {
				t.Errorf("scalar skip = (%d, %v), want (%d, %v) [%s]", gotEnd, gotErr, c.slow.end, c.slow.err, c.why)
			}
		})
	}
}

// TestSkipBackslashLengthCliff pins where the backslash divergence still turns
// on something other than the document's meaning: the fast path answers with the
// block math whenever it has 64 bytes to read and with a byte walk when it does
// not, and the two disagree about a stray backslash OUTSIDE a string. The block
// math masks the byte after it out of the quote bitmap (findEscaped64 is pure bit
// math and cannot know the backslash is not in a string); the byte walk only ever
// sets its escape flag inside one. So `{\"a}` is a truncated container under a
// block and an accepted one over it, and the flip is at 64 bytes of input.
//
// This USED to be an alignment cliff — the same document flipping on where the
// 64-byte grid fell, at 62/63/64 spaces of padding — because the block loop
// handed its final < 64 bytes to that byte walk together with a carried
// prevEscaped the walk applied to the very next byte, brace or not. The tail is
// now one more (overlapping) block, so the block math decides every byte of a
// document that reaches one block, and the grid no longer shows: `{`+pad+`\}` is
// accepted at every padding. The remaining threshold is a property of the input's
// LENGTH, which is at least visible in the input.
func TestSkipBackslashLengthCliff(t *testing.T) {
	if !fastSkipAvail {
		t.Skip("no SIMD skip path on this build/CPU")
	}
	// The old cliff, gone: padding does not change the verdict.
	for pad := 0; pad <= 130; pad++ {
		data := []byte(`{` + strings.Repeat(" ", pad) + `\}`)
		if _, err := skipContainerFast(data, 0, '{'); err != nil {
			t.Fatalf("pad=%d: skipContainerFast err = %v, want nil (the grid must not show)", pad, err)
		}
	}
	// The one that is left: under a block the byte walk decides, over it the
	// block math does, and they read a stray backslash differently.
	for _, tc := range []struct {
		pad     int
		wantErr error
	}{
		{0, ErrTruncated},  // 5 bytes: byte walk, the backslash escapes nothing
		{58, ErrTruncated}, // 63 bytes: still the byte walk
		{59, nil},          // 64 bytes: the block math, which escapes the quote
		{60, nil},
	} {
		data := []byte(`{` + strings.Repeat(" ", tc.pad) + `\"a}`)
		if _, err := skipContainerFast(data, 0, '{'); err != tc.wantErr {
			t.Errorf("pad=%d (len %d): skipContainerFast err = %v, want %v", tc.pad, len(data), err, tc.wantErr)
		}
		// The scalar path has neither a block grid nor escape state, so it reads
		// the quote as opening a string at every length.
		if _, err := skipObject(data, 0); err != ErrTruncated {
			t.Errorf("pad=%d: scalar skipObject err = %v, want %v", tc.pad, err, ErrTruncated)
		}
	}
}

// TestSkipDepthDivergence pins the divergence the MaxDepth bound in skip.go
// created: the iterative SIMD path has no stack to overflow and accepts any
// nesting, while the recursive scalar path must bound itself. If the assembly
// block loops ever learn to test depth, this test is the one that should fail.
func TestSkipDepthDivergence(t *testing.T) {
	if !fastSkipAvail {
		t.Skip("no SIMD skip path on this build/CPU")
	}
	n := MaxDepth + 1
	data := []byte(strings.Repeat("[", n) + strings.Repeat("]", n))

	if end, err := skipContainerFast(data, 0, '['); err != nil || end != len(data) {
		t.Errorf("skipContainerFast past MaxDepth = (%d, %v), want (%d, nil) — the SIMD path is deliberately unbounded", end, err, len(data))
	}
	if _, err := skipArray(data, 0); err != ErrMaxDepth {
		t.Errorf("scalar skipArray past MaxDepth = %v, want ErrMaxDepth", err)
	}
}
