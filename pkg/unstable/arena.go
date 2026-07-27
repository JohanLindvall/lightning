package unstable

import "unsafe"

// Arena is a chunked bump allocator for the backings of small scalar slices,
// used by the Decode*SliceArena readers that //lightning:arena decoders call.
// On documents shaped like a skeletal animation or a mesh — many thousands of
// 3–4 element []float64 fields — the per-slice make() is the decode's dominant
// allocation source (marine_ik: 95% of allocated objects, ~20% of CPU in
// mallocgc); carving those backings out of shared chunks turns tens of
// thousands of tiny allocations into a few hundred chunk allocations. The
// total bytes allocated and zeroed are unchanged (a fresh chunk is zeroed just
// as each make() was), so the win is precisely the removed per-object mallocgc
// work, plus the GC tracking that many fewer objects.
//
// The zero value is ready to use. An Arena is not safe for concurrent use; the
// generated decoders declare one per UnmarshalJSON call, which also bounds its
// lifetime — the Arena itself is garbage as soon as the decode returns, and
// each chunk lives exactly as long as some result slice still references it.
//
// That last property is the deliberate trade-off of //lightning:arena: a
// single surviving 3-element slice keeps its whole arenaChunkBytes chunk
// reachable. Callers that decode, process, and discard together (the common
// shape for these documents) lose nothing; callers that retain a few small
// slices out of a large decode should not opt in.
type Arena struct {
	free []byte // the unused tail of the current chunk
}

const (
	// arenaChunkBytes is the size of each chunk. Large enough that a
	// small-slice-heavy decode amortizes one chunk allocation over ~a hundred
	// carves; small enough that the pinning granularity (and the waste when an
	// arena decode has only a handful of small slices) stays modest.
	arenaChunkBytes = 4096
	// arenaMaxCarve is the largest backing carved from a chunk. A bigger
	// backing gets its own make() — one allocation for many elements is
	// already cheap, carving it would burn most of a chunk, and a huge
	// surviving slice pinning its chunk is exactly the trade-off large
	// backings don't need. It also bounds the waste when a carve doesn't fit
	// the current chunk's tail (< arenaMaxCarve bytes are abandoned).
	arenaMaxCarve = 512
)

// arenaScalar constrains carving to pointer-free element kinds. The backing is
// carved from a []byte chunk, which the GC does not scan, so a type holding
// pointers must never be carved from one — its pointees would be collectable
// while still referenced. Every kind here is 8 bytes or smaller, so the
// 8-aligned carve offsets below align any of them.
type arenaScalar interface {
	intKind | uintKind | ~float32 | ~float64
}

// arenaCarve returns a length-0, capacity-n []T for the decode loop to append
// into, carved from the arena's current chunk (or freshly allocated when the
// backing is over arenaMaxCarve). The cursor advances past the full capacity
// up front, so the region is exclusively this slice's: no later carve overlaps
// it, and a caller appending into leftover capacity afterwards can never
// clobber a neighbour. Chunks are make-zeroed and a carved region is never
// reused, so the capacity beyond the decoded length reads as zeros — exactly
// what a make([]T, 0, n) backing shows.
//
// Offsets stay 8-aligned: chunks are 8-aligned by the allocator and each carve
// rounds its byte size up to a multiple of 8.
func arenaCarve[T arenaScalar](a *Arena, n int) []T {
	var zero T
	size := (n*int(unsafe.Sizeof(zero)) + 7) &^ 7
	if size > arenaMaxCarve {
		return make([]T, 0, n)
	}
	if size > len(a.free) {
		a.free = make([]byte, arenaChunkBytes)
	}
	p := (*T)(unsafe.Pointer(unsafe.SliceData(a.free)))
	a.free = a.free[size:]
	return unsafe.Slice(p, n)[:0]
}
