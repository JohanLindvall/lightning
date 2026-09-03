package unstable

import (
	"unsafe"

	"github.com/JohanLindvall/arena"
)

// Arena is the chunk-backed store the Decode*SliceArena readers carve slice
// backings from, and which the //lightning:arena decoders declare one of per
// element type per UnmarshalJSON call. It is github.com/JohanLindvall/arena's
// Arena, aliased here so generated code names only this package.
//
// On documents shaped like a skeletal animation or a mesh — many thousands of
// 3-4 element []float64 fields — the per-slice make() is the decode's dominant
// allocation source (marine_ik: 95% of allocated objects, ~20% of CPU in
// mallocgc); carving those backings out of shared chunks turns tens of
// thousands of tiny allocations into a few hundred chunk allocations. The
// total bytes allocated and zeroed are unchanged (a fresh chunk is zeroed just
// as each make() was), so the win is precisely the removed per-object mallocgc
// work, plus the GC tracking that many fewer objects.
//
// An arena is per-decode: the generated method declares it, which bounds its
// lifetime — the Arena itself is garbage as soon as the decode returns, and
// each chunk lives exactly as long as some result slice still references it.
// Nothing here ever calls Reset, which is what keeps two properties the readers
// rely on: a carved region is always freshly make-zeroed (so the capacity past
// a decoded length reads as zeros, exactly as a make([]T, 0, n) backing does),
// and no region is ever handed out twice.
//
// That chunk lifetime is the deliberate trade-off of //lightning:arena: a
// single surviving 3-element slice keeps its whole arenaChunkBytes chunk
// reachable. Callers that decode, process, and discard together (the common
// shape for these documents) lose nothing; callers that retain a few small
// slices out of a large decode should not opt in.
type Arena[T any] = arena.Arena[T]

const (
	// arenaChunkBytes is the size of each chunk. Large enough that a
	// small-slice-heavy decode amortizes one chunk allocation over ~a hundred
	// carves; small enough that the pinning granularity (and the waste when an
	// arena decode has only a handful of small slices) stays modest. It is well
	// under the arena package's own 64 KiB default, which is why the generated
	// code builds its arenas with NewArena rather than taking zero values.
	arenaChunkBytes = 4096
	// arenaMaxCarve is the largest backing carved from a chunk. A bigger
	// backing gets its own make() — one allocation for many elements is
	// already cheap, carving it would burn most of a chunk, and a huge
	// surviving slice pinning its chunk is exactly the trade-off large
	// backings don't need. It also bounds the waste when a carve doesn't fit
	// the current chunk's tail (< arenaMaxCarve bytes are abandoned).
	arenaMaxCarve = 512
)

// NewArena returns the Arena the generated decoders declare, sized to
// arenaChunkBytes. It is a value rather than a pointer so the generated arena
// struct holds its arenas inline, and exported because generated code calls it.
func NewArena[T any]() Arena[T] { return arena.Make[T](arenaChunkBytes) }

// arenaCarve returns a length-0, capacity-n []T for the decode loop to append
// into, carved from a's current chunk (or freshly allocated when the backing is
// over arenaMaxCarve). Reserve advances the arena past the full capacity up
// front, so the region is exclusively this slice's: no later carve overlaps it,
// and because the capacity is exactly n, a caller appending past it reallocates
// to the heap rather than clobbering a neighbour.
//
// Unlike the []byte-chunk allocator this used to carve from, the arena is typed
// — its chunks are []T — so there is no pointer reinterpretation to get wrong,
// no alignment to maintain, and no need to restrict T to pointer-free kinds:
// the GC scans a chunk correctly for whatever T is.
func arenaCarve[T any](a *Arena[T], n int) []T {
	var zero T
	if n*int(unsafe.Sizeof(zero)) > arenaMaxCarve {
		return make([]T, 0, n)
	}
	return a.Reserve(n)
}
