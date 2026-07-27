package unstable

import "testing"

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
