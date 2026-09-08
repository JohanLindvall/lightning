package unstable

import (
	"strings"
	"testing"
)

// TestGrowSlice checks the growth helper preserves length and contents, at least
// doubles capacity, and never damps the way runtime.nextslicecap does past 256
// elements — the whole reason it exists.
func TestGrowSlice(t *testing.T) {
	// From nil / empty: a floor of 4, not 0.
	if got := GrowSlice[int](nil); cap(got) < 4 || len(got) != 0 {
		t.Errorf("GrowSlice(nil) len=%d cap=%d, want len 0 cap >= 4", len(got), cap(got))
	}
	if got := GrowSlice(make([]int, 0)); cap(got) < 4 {
		t.Errorf("GrowSlice(empty) cap=%d, want >= 4", cap(got))
	}

	// Contents and length preserved, capacity at least doubled, at sizes either
	// side of nextslicecap's 256-element damping threshold.
	for _, n := range []int{1, 2, 3, 4, 255, 256, 257, 1024, 5000} {
		s := make([]int, n)
		for i := range s {
			s[i] = i * 7
		}
		got := GrowSlice(s)
		if len(got) != n {
			t.Fatalf("n=%d: len=%d, want %d", n, len(got), n)
		}
		if cap(got) < 2*n {
			t.Errorf("n=%d: cap=%d, want >= %d (growth must not damp)", n, cap(got), 2*n)
		}
		for i := range s {
			if got[i] != i*7 {
				t.Fatalf("n=%d: element %d = %d, want %d", n, i, got[i], i*7)
			}
		}
		// The result must be a fresh backing array, not an alias.
		if n > 0 {
			got[0] = -1
			if s[0] == -1 {
				t.Errorf("n=%d: GrowSlice aliased the original backing array", n)
			}
		}
	}

	// A partially filled slice keeps its length, not its capacity, as the length.
	s := make([]int, 3, 10)
	if got := GrowSlice(s); len(got) != 3 || cap(got) < 20 {
		t.Errorf("partial: len=%d cap=%d, want len 3 cap >= 20", len(got), cap(got))
	}
}

// TestGrowSliceEst checks the progress-extrapolating variant: length and
// contents preserved with a fresh backing, capacity at least doubled always
// (the lower clamp — never worse than GrowSlice), the estimate honored between
// the clamp bounds, the 8x ceiling for arrays that do not span the rest of the
// document, and plain doubling on degenerate progress inputs.
func TestGrowSliceEst(t *testing.T) {
	fill := func(n int) []int {
		s := make([]int, n)
		for i := range s {
			s[i] = i * 7
		}
		return s
	}
	check := func(name string, s []int, got []int, wantCap int) {
		t.Helper()
		if len(got) != len(s) {
			t.Fatalf("%s: len=%d, want %d", name, len(got), len(s))
		}
		if cap(got) != wantCap {
			t.Errorf("%s: cap=%d, want %d", name, cap(got), wantCap)
		}
		for i := range s {
			if got[i] != i*7 {
				t.Fatalf("%s: element %d = %d, want %d", name, i, got[i], i*7)
			}
		}
		if len(s) > 0 {
			got[0] = -1
			if s[0] == -1 {
				t.Errorf("%s: aliased the original backing array", name)
			}
			got[0] = s[0]
		}
	}

	// Middle of the clamp window: the (padded) estimate itself is used.
	// 100 elements over bytes 0..250 of a 1000-byte document extrapolate to
	// 100*1000/250 = 400, padded upper-ish to 400+400/8+1 = 451, inside
	// [200, 800].
	s := fill(100)
	check("estimate", s, GrowSliceEst(s, 0, 250, 1000), 451)

	// A non-zero start indexes the array's '[', not the document head: the
	// same shape shifted by 50 bytes must give the same 451.
	check("estimate/offset", s, GrowSliceEst(s, 50, 300, 1050), 451)

	// Below the lower clamp: 100 elements with only 1000-900 bytes left
	// estimate 100*1000/900 = 111 (125 padded) < 2*cap; the floor keeps flat
	// doubling.
	check("clamp-low", s, GrowSliceEst(s, 0, 900, 1000), 200)

	// Above the upper clamp: 10 elements in the first 10 bytes of a 10000-byte
	// document estimate 10000 elements; the ceiling holds at 8*cap. This is the
	// nested-slice case — the trailing bytes are not this array's elements.
	s10 := fill(10)
	check("clamp-high", s10, GrowSliceEst(s10, 0, 10, 10000), 80)

	// The github_events shape that motivated the upper-ish pad: 4 of 30 large
	// records decoded, raw estimate 4*65132/8855 = 29 — one SHORT of the real
	// 30, which would force a 29 -> 58 doubling at the last element. The pad
	// lifts it to 33, the 8x ceiling trims to 32, and the array finishes in
	// this single grow.
	s4 := fill(4)
	check("pad-covers-near-miss", s4, GrowSliceEst(s4, 0, 8855, 65132), 32)

	// Degenerate progress falls back to plain doubling: no bytes consumed
	// (i == start, would divide by zero) and start at/past the document end.
	check("degenerate/i==start", s10, GrowSliceEst(s10, 5, 5, 10000), 20)
	check("degenerate/i<start", s10, GrowSliceEst(s10, 50, 40, 10000), 20)
	check("degenerate/start==end", s10, GrowSliceEst(s10, 100, 200, 100), 20)
	check("degenerate/start>end", s10, GrowSliceEst(s10, 200, 300, 100), 20)

	// Empty/nil input keeps GrowSlice's floor of 4 (no elements to
	// extrapolate from).
	if got := GrowSliceEst[int](nil, 0, 10, 100); cap(got) != 4 || len(got) != 0 {
		t.Errorf("nil: len=%d cap=%d, want len 0 cap 4", len(got), cap(got))
	}
	if got := GrowSliceEst(make([]int, 0, 1), 0, 10, 100); cap(got) != 4 || len(got) != 0 {
		t.Errorf("empty: len=%d cap=%d, want len 0 cap 4", len(got), cap(got))
	}

	// A partially filled slice keeps its length and clamps against its cap:
	// len 3, cap 10 with a huge estimate lands on 8*cap = 80.
	p := make([]int, 3, 10)
	for i := range p {
		p[i] = i * 7
	}
	check("partial", p, GrowSliceEst(p, 0, 1, 1000), 80)
}

// TestArrayEndAt holds the forward container scan to SkipValue over the array
// it is inside: the walk starts at an element boundary at depth 1, which is
// where a slice decoder's loop is when it grows, so the answer must be the
// same index SkipValue reports for the whole array.
func TestArrayEndAt(t *testing.T) { checkArrayEndAt(t) }

// checkArrayEndAt is the body, so the amd64 test that flips useSkipBlocks and
// fastSkipAvail can drive all three arms of the dispatch on one machine — the
// live flags only ever exercise whichever the host selects, and the scalar
// fallback would otherwise never run anywhere the corpus is measured.
func checkArrayEndAt(t *testing.T) {
	t.Helper()
	// Each case is an array; the scan starts at each element boundary in turn.
	arrays := []string{
		`[]`,
		`[1,2,3]`,
		`["a","b","c"]`,
		`[{"a":1},{"b":[2,3]},{"c":{"d":"]"}}]`,
		`[[1,2],[3,4],[5,6]]`,
		`["]","[","\"]","\\"]`,
		`[` + strings.Repeat(`{"k":"vvvvvvvvvvvvvvvvvvvvvvvvvvvvvv"},`, 40) + `{"k":"x"}]`,
		`[` + strings.Repeat(`123456789,`, 200) + `0]`,
	}
	for _, arr := range arrays {
		for _, pad := range []string{"", ",\"tail\"", ":" + strings.Repeat("z", 200)} {
			doc := []byte(`{"a":` + arr + pad + `}`)
			start := 5
			want, err := SkipValue(doc, start)
			if err != nil {
				t.Fatalf("SkipValue(%q): %v", arr, err)
			}
			// Walk every element boundary: just past '[' and just past each
			// separating comma at depth 1.
			i := start + 1
			for {
				i = SkipWS(doc, i)
				if got := arrayEndAt(doc, i); got != want {
					t.Fatalf("arrayEndAt(%q, %d) = %d, want %d", doc, i, got, want)
				}
				if doc[i] == ']' {
					break
				}
				e, err := SkipValue(doc, i)
				if err != nil {
					t.Fatalf("SkipValue element: %v", err)
				}
				i = SkipWS(doc, e)
				if doc[i] != ',' {
					break
				}
				i++
			}
		}
	}
}

// TestGrowSliceSpan pins the estimate, and TestScanWorth the gate in front of
// it: an array below any of the three thresholds must not be scanned at all,
// and one above them must be sized from the span the scan measured.
func TestScanWorth(t *testing.T) {
	const sz = 8
	cases := []struct {
		name                     string
		capacity, size, consumed int
		want                     bool
	}{
		{"too few elements", 16, 344, 16 * 344, false},
		{"too few bytes", 64, 8, 64 * 8, false},
		{"document too fat", 600, sz, arrayScanRatio*600*sz + 8, false},
		{"worth it", 600, sz, 600 * sz, true},
		{"exactly at the element floor", arrayScanMinElems, 512, 4096, true},
	}
	for _, c := range cases {
		if got := ScanWorth(c.capacity, c.size, c.consumed); got != c.want {
			t.Errorf("%s: ScanWorth(%d, %d, %d) = %v, want %v",
				c.name, c.capacity, c.size, c.consumed, got, c.want)
		}
	}
}

func TestGrowSliceSpan(t *testing.T) {
	// 4000 elements of 5 bytes each, and the caller is 600 elements in.
	doc := []byte(`[` + strings.Repeat(`1234,`, 4000) + `0]`)
	big := make([]int, 600)
	for i := range big {
		big[i] = i
	}
	got, end := GrowSliceSpan(big, doc, 0, 1+600*5, 0)
	if end != len(doc) {
		t.Fatalf("end=%d, want %d", end, len(doc))
	}
	if cap(got) < 4000 || cap(got) > 4001+4001/8 {
		t.Errorf("cap=%d, want the measured ~4001 plus the 1/8 pad", cap(got))
	}
	if len(got) != 600 {
		t.Fatalf("len=%d, want 600", len(got))
	}
	for i := range big {
		if got[i] != i {
			t.Fatalf("element %d = %d, want %d", i, got[i], i)
		}
	}

	// A known end is reused rather than scanned for again.
	got2, end2 := GrowSliceSpan(big, doc, 0, 1+600*5, len(doc))
	if end2 != len(doc) || cap(got2) != cap(got) {
		t.Errorf("cached end: cap=%d end=%d, want cap %d end %d", cap(got2), end2, cap(got), len(doc))
	}

	// The estimate never shrinks below the 2x floor GrowSlice would have given,
	// however little of the array is left.
	nearlyDone := make([]int, 600)
	if got, _ := GrowSliceSpan(nearlyDone, doc, 0, len(doc)-6, 0); cap(got) < 1200 {
		t.Errorf("nearly done: cap=%d, want at least the 1200 doubling gives", cap(got))
	}
}
