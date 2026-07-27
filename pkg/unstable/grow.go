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

// GrowSliceEst is GrowSlice with a progress-based capacity estimate, for a
// decode loop that is about to append past cap(s) and knows where in the
// document it is. The caller has decoded len(s) elements while the scan
// advanced from start (the index of the array's '[') to i, in a document
// ending at end (len(data)). Assuming the bytes not yet consumed hold elements
// at the same density as the bytes already consumed, the final element count
// is about
//
//	len(s) * (end - start) / (i - start)
//
// This is a fine hint for an array that spans the rest of the document (a
// root array of large records), and it costs no extra scanning — all three
// indexes are already live in the decode loop, which is what distinguishes it
// from the rejected counting presizes.
//
// The raw ratio is then padded by 1/8 (+1) to make it genuinely upper-ish. A
// hint one element SHORT is the worst outcome: the loop grows once more for
// the tail and the 2x floor below doubles a nearly-final capacity — measured
// on github_events (30 large records), the unpadded estimate landed at 29,
// the last element forced 29 -> 58, and the decode allocated MORE than flat
// doubling (234 KB vs 165 KB B/op). Elements are never uniform, so a
// near-exact estimate is common; the 12.5% pad turns it into a small
// over-allocation (bounded, unlike the miss it prevents, which costs a full
// extra backing + memmove). With the pad the same decode allocates 103 KB —
// under both the unpadded estimate and the flat-2x baseline.
//
// The estimate is only ever a capacity hint, so it is clamped to
// [2*cap(s), 8*cap(s)]:
//   - never below 2*cap(s): each grow is at least the flat doubling GrowSlice
//     does, so no caller regresses against that behavior (an array whose
//     remaining bytes are NOT elements — the document continues after it —
//     under-estimates, and the floor restores plain doubling);
//   - never above 8*cap(s): a nested slice early in a large document sees an
//     estimate inflated by all the trailing non-element bytes; the ceiling
//     bounds the over-allocation to two extra doublings per grow (8x = 2x^3),
//     and the next grow re-estimates from better progress.
//
// GrowSlice's floor of 4 applies to the lower bound, so for cap(s) < 2 the
// clamp window is [4, 16] rather than [2*cap, 8*cap] — harmless, since the
// generated decoders give a fresh slice a ~256-byte first-append hint and only
// ever grow from there.
//
// Overflow and degenerate inputs: the product len(s)*(end-start) is computed
// in uint64 because both factors can be large at once (a huge element count
// late in a huge document overflows int64; in uint64 it is exact for any
// document under ~2^32 bytes even in the worst one-byte-element case, since
// len(s) is bounded by the bytes consumed — and were it ever to wrap, the
// result is still only a capacity hint inside the clamp window: a mis-sized
// slice, never a misdecode). Both factors are non-negative under the guards —
// i <= start (no progress; includes a caller passing a bad start) or
// end <= start (start past the document) would divide by zero or yield
// nonsense, so those fall back to plain doubling, as does len(s) == 0. The
// clamped result is at most 4*lo, which fits uint64 exactly (lo is an int); a
// slice so large that 8*cap(s) exceeds int is already beyond what make/append
// could have built, matching GrowSlice's own 2*cap(s) exposure.
func GrowSliceEst[T any](s []T, start, i, end int) []T {
	lo := 2 * cap(s)
	if lo < 4 {
		lo = 4
	}
	n := lo
	if len(s) > 0 && i > start && end > start {
		est := uint64(len(s)) * uint64(end-start) / uint64(i-start)
		est += est/8 + 1 // upper-ish: don't land one short (see above)
		if hi := uint64(lo) * 4; est > hi {
			est = hi
		}
		if est > uint64(n) {
			n = int(est)
		}
	}
	t := make([]T, len(s), n)
	copy(t, s)
	return t
}
