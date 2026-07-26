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
