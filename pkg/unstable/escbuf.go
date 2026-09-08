package unstable

import "sync"

// An escaped string decodes into a buffer that the result then ALIASES
// (decodeEscaped hands it out with unsafeStr), so that buffer is not scratch at
// all — it is the string's backing, and it used to cost one `make` per escaped
// string. On an escape-heavy document that is most of the decode: makeslice
// under decodeStringEscaped measured 22% of gsoc_2018 on a Neoverse N2, about
// half of it the allocation itself and half the collector work that thousands
// of extra objects per decode pace (measured by rerunning under GOGC=off, where
// the same change is worth 13.6% instead of 25.9%).
//
// escapeScratch carves those buffers from a chunk instead. Every carve is
// exclusive — cap is set to exactly n, so a caller's append reallocates onto the
// heap rather than reaching a neighbour — and no chunk is ever reused, so a
// carved buffer behaves exactly like a make'd one. The only observable
// difference is which object the garbage collector sees it in, and therefore
// how much memory ONE surviving string keeps alive.
//
// That retention is the whole design constraint, because the copying string
// readers exist precisely so a caller need not hold on to the document. Three
// rules bound it:
//
//   - A chunk is never larger than escapeChunkMax, so a string cannot pin more
//     than that however long it lives.
//   - A chunk is never larger than the document that needed it, so decoding a
//     small document cannot create a large chunk. The bound is really
//     min(escapeChunkMax, the largest document decoded into this chunk) — no
//     more than a nocopy decode of that document already keeps alive.
//   - A body over escapeMaxCarve gets its own make, which both keeps the tail a
//     chunk can waste to that bound and keeps one large string off a chunk.
//
// The chunk lives in a sync.Pool rather than a package variable: Get removes it,
// so the bump is exclusive for the one carve it is checked out for, and Put
// returns it for the next string. A pool emptied at a GC costs only the tail of
// whatever chunk was in it. The pair measures 25.0ns against a make's 42.5ns
// (BenchmarkEscapeScratch); a package variable would be 11.7ns and is not an
// option.
const (
	escapeChunkMin = 1 << 10
	escapeChunkMax = 16 << 10
	escapeMaxCarve = 4096
)

type escapeBump struct{ buf []byte }

var escapeBumpPool = sync.Pool{New: func() any { return new(escapeBump) }}

// escapeScratch returns an empty buffer with capacity exactly n for a body of
// the document data, together with the chunk it came from (nil when the buffer
// is its own allocation). The caller must hand the chunk back to escapeRelease
// with the number of bytes the decode actually wrote, which is at most n —
// unescaping only ever shrinks — so the rest of the carve is returned rather
// than left dead in the chunk. A dense \uXXXX body decodes to half its escaped
// length, and that half is worth having back.
func escapeScratch(data []byte, n int) (buf []byte, chunk *escapeBump, off int) {
	if n > escapeMaxCarve {
		return make([]byte, 0, n), nil, 0
	}
	a := escapeBumpPool.Get().(*escapeBump)
	if cap(a.buf)-len(a.buf) < n {
		sz := len(data)
		if sz > escapeChunkMax {
			sz = escapeChunkMax
		} else if sz < escapeChunkMin {
			sz = escapeChunkMin
		}
		if sz < n {
			sz = n // n <= len(data) always, so this is belt and braces
		}
		a.buf = make([]byte, 0, sz)
	}
	off = len(a.buf)
	a.buf = a.buf[:off+n]
	return a.buf[off : off : off+n], a, off
}

// escapeRelease returns a's chunk to the pool, keeping only the used bytes of
// the carve escapeScratch made. used is clamped to the carve: a decode whose
// result outgrew the estimate (only reachable on truncated input, where the
// estimate is the escape-free prefix) reallocated onto the heap, and then the
// whole carve is dead and must simply not be handed out again.
func escapeRelease(a *escapeBump, off, n, used int) {
	if a == nil {
		return
	}
	if used > n {
		used = n
	}
	a.buf = a.buf[:off+used]
	escapeBumpPool.Put(a)
}
