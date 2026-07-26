package unstable

// GrowSlice returns s with its capacity at least doubled, its length and contents
// preserved, for a decode loop that is about to append past cap(s).
//
// It exists to bypass runtime.nextslicecap's damping: bare append doubles only
// while cap is under 256 elements and then grows by cap += (cap+768)>>2, about
// 1.25x. Since the bytes a growing slice allocates in total come to
// final_cap * f/(f-1), the 1.25x regime allocates roughly 5x the final size and
// memmoves about 4x it, where a flat 2x allocates 2x and memmoves 1x. Arrays that
// stay under 256 elements are unaffected either way, so this only changes the
// large-array regime.
func GrowSlice[T any](s []T) []T {
	n := 2 * cap(s)
	if n < 4 {
		n = 4
	}
	t := make([]T, len(s), n)
	copy(t, s)
	return t
}
