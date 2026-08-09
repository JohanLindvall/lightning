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

// TestSkipBackslashAlignmentCliff pins the one property of the backslash
// divergence that makes it more than a curiosity: it depends on where the 64-byte
// block grid falls, so padding a document without changing its meaning flips the
// fast path's verdict. At 63 spaces the backslash is the block's last byte, so
// findEscaped64's prevEscaped carry reaches the tail and masks the container's own
// closing brace out of the close bitmap.
func TestSkipBackslashAlignmentCliff(t *testing.T) {
	if !fastSkipAvail {
		t.Skip("no SIMD skip path on this build/CPU")
	}
	for _, tc := range []struct {
		pad     int
		wantErr error
	}{
		{62, nil},
		{63, ErrTruncated},
		{64, nil},
	} {
		data := []byte(`{` + strings.Repeat(" ", tc.pad) + `\}`)
		_, err := skipContainerFast(data, 0, '{')
		if err != tc.wantErr {
			t.Errorf("pad=%d: skipContainerFast err = %v, want %v", tc.pad, err, tc.wantErr)
		}
		// The scalar path has no block grid and no escape state, so it accepts
		// every one of these regardless of padding.
		if _, err := skipObject(data, 0); err != nil {
			t.Errorf("pad=%d: scalar skipObject err = %v, want nil", tc.pad, err)
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
