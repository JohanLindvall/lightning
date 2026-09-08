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
	n := max(2*cap(s), 4)
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
	return growTo(s, estimateCap(len(s), cap(s), start, i, end, 4))
}

// estimateCap is the shared body of the two extrapolating grows: the density
// estimate, its pad and its clamp, given cur = cap(s), length = len(s) and a
// ceiling expressed as a multiple of the 2x floor (so ceil=4 means 8*cap).
// Separated only so GrowSliceSpan, whose end is measured rather than assumed,
// can carry a looser ceiling than GrowSliceEst's.
func estimateCap(length, cur, start, i, end, ceil int) int {
	lo := max(2*cur, 4)
	n := lo
	if length > 0 && i > start && end > start {
		est := uint64(length) * uint64(end-start) / uint64(i-start)
		est += est/8 + 1 // upper-ish: don't land one short (see above)
		est = min(est, uint64(lo)*uint64(ceil))
		if est > uint64(n) {
			n = int(est)
		}
	}
	return n
}

func growTo[T any](s []T, n int) []T {
	t := make([]T, len(s), n)
	copy(t, s)
	return t
}

// The three thresholds that gate the scan GrowSliceSpan pays for. Each was
// measured, and each keeps out a different way for the scan to lose.
//
// The scan is O(the array's remaining JSON bytes) — about 0.9 instructions per
// byte measured in the shape this gate sees, one scan plus the call around it
// (the "quarter of an instruction per byte" quoted for skipContainerFast
// elsewhere holds only where most blocks take its popcount bulk path). The
// saving is the Go bytes the remaining doublings would otherwise allocate,
// zero, copy and hand to the collector, which on this corpus is worth 1.5 to 7
// instructions a byte. So the scan pays when the array still has several
// doublings left AND its JSON is not much fatter than the Go value it decodes
// into.
//
//   - arrayScanMinElems: the backing must already hold this many elements. The
//     saving is bounded by the doublings that remain, and an array of 30 fat
//     records has none worth having — payload_large's 344-byte Topics entries
//     cross the byte floor at 16 elements, then end: the scan cost 11 KB of
//     reading to size the last backing 34 instead of 32 (+2.9% instructions for
//     no B/op at all). Requiring 32 elements — five doublings past the capacity
//     hint — is the cheapest evidence that an array is long.
//   - arrayScanMin: and it must be a real backing. A scan carries a call and a
//     block-loop prologue, which a few hundred bytes of elements cannot repay.
//   - arrayScanRatio: the JSON consumed (i - start) must be at most this many
//     times the backing filled (cap * sizeof(T)) — i.e. the document must not
//     be much fatter than the Go value. citm_catalog's performances array is
//     136 bytes of Go per ~7 KB of JSON: scanning it read the whole 1.7 MB
//     document to save 18 KB of allocation, 9.8% of that decode. The test is
//     scale-invariant, so an array either passes it at every grow or at none.
//
// Scanning EVERY array is worse than any of these (skipBlocks 17.9% of
// citm_catalog, +23%), because nested arrays — performances -> seatCategories
// -> areas -> blockIds — each re-read the same bytes once per level.
const (
	arrayScanMinElems = 32
	arrayScanMin      = 4 << 10
	arrayScanRatio    = 5
)

// ScanWorth answers the gate for a slice about to grow: capacity elements of
// size bytes each are already filled and consumed bytes of JSON produced them.
// It is separate from GrowSliceSpan, and called from the generated loop rather
// than from inside it, because GrowSliceSpan is far past the inline budget
// (cost 231) while this is a few compares: an array that never scans then keeps
// exactly the inlined GrowSlice it had before, instead of paying a call per
// grow for a decision that is nearly always no. Measured on citm_catalog, whose
// 591 grows per decode never scan: +1.2% instructions with the gate inside,
// nothing with it here.
func ScanWorth(capacity, size, consumed int) bool {
	got := uintptr(capacity) * uintptr(size)
	return capacity >= arrayScanMinElems && got >= arrayScanMin &&
		uintptr(consumed) <= arrayScanRatio*got
}

// GrowSliceSpan is the growth of a NESTED slice: the same progress
// extrapolation GrowSliceEst does, over the array's OWN byte span rather than
// the document's.
//
// GrowSliceEst's premise — that the array runs to the end of the document — is
// structurally true only at the root. For a slice reached as a field the
// document tail says nothing about the array's length, which is why nested
// slices grew by blind doubling: the estimate would always saturate its clamp.
// What was missing was the array's own end, and arrayEndAt supplies it, forward
// from the cursor.
//
// end is the array's end index, 0 while it is still unknown; the caller keeps
// the returned value and hands it back, so the scan happens at most once per
// array however many times it grows. ScanWorth decides whether to call this at
// all — see there, and see the thresholds above for why an ungated scan loses.
//
// The ceiling is looser than GrowSliceEst's (32*lo, i.e. 64x the current
// capacity, against 8x): that clamp exists because a nested array's estimate
// used to be inflated by every trailing byte of the document, and with a
// measured end the only error left is variance in element size across the
// array — which the next grow re-estimates away from a longer sample.
func GrowSliceSpan[T any](s []T, data []byte, start, i, end int) ([]T, int) {
	if end == 0 {
		end = arrayEndAt(data, i)
	}
	return growTo(s, estimateCap(len(s), cap(s), start, i, end, 32)), end
}

// arrayEndAt returns the index just past the ']' closing the array whose
// element boundary i sits on — i is at depth 1 inside that array and outside
// every string, which is exactly the state a slice decoder's loop is in when
// it is about to append.
//
// Scanning FORWARD FROM i, rather than from the array's '[', is what makes the
// cost proportional to what is left rather than to the whole array: an array
// about to end is settled in a block or two, which is the case a growth
// estimate must not overpay for. It is the container skip's own block loop
// entered at depth 1 (skipContainerFast does the same thing one byte past an
// open bracket), over bytes the decode is about to read anyway.
//
// A malformed or truncated array has no end to find; len(data) is then the
// answer, which is the honest bound on what is left and only ever sizes a
// backing.
func arrayEndAt(data []byte, i int) int {
	if useSkipBlocks && i+64 <= len(data) {
		end, d, pe, pis := skipBlocks(data, i, 1, true)
		if end >= 0 {
			return end
		}
		tail := i + ((len(data) - i) &^ 63)
		if end, err := skipContainerBlocks(data, tail, '[', d, pe, pis); err == nil {
			return end
		}
		return len(data)
	}
	if fastSkipAvail {
		if end, err := skipContainerBlocks(data, i, '[', 1, 0, 0); err == nil {
			return end
		}
		return len(data)
	}
	// No block machinery: walk the remaining elements. This is the estimate's
	// slow path on an architecture whose skip is scalar anyway.
	for {
		i = SkipWS(data, i)
		if uint(i) >= uint(len(data)) {
			return len(data)
		}
		switch data[i] {
		case ']':
			return i + 1
		case ',':
			i++
			continue
		}
		end, err := SkipValue(data, i)
		if err != nil {
			return len(data)
		}
		i = end
	}
}
