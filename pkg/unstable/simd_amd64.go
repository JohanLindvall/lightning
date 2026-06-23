//go:build amd64

package unstable

import "golang.org/x/sys/cpu"

//go:noescape
func indexQuoteOrBackslashSSE2(b []byte) int

//go:noescape
func indexStructuralAVX2(b []byte) int

//go:noescape
func indexEscapeSSE2(b []byte) int

var useAVX2 = cpu.X86.HasAVX2

// indexEscape returns the index of the first byte that JSON string encoding must
// escape — a control byte < 0x20, '"' or '\\' — or len(b) if none. It is a single
// (inlinable) call into indexEscapeSSE2, which begins in SSE2 (baseline, no
// VZEROUPPER), switches to AVX2 once 32 bytes are clean (a long unescaped run, e.g.
// a log line), and drops a sub-block buffer to a scalar tail. The < 0x20 SIMD test
// is a per-block PMINUB(v, 0x1f) == v.
func indexEscape(b []byte) int {
	return indexEscapeSSE2(b)
}

// indexCloseOrEscape returns the index of the first '"' or '\\' byte in b, or
// len(b) if neither is present. String scanning begins in SSE2, not AVX2: SSE2 is
// baseline on every amd64 CPU (no feature gate) and avoids the VZEROUPPER every
// AVX2 call must execute, which is pure overhead for the short keys and string
// values that dominate JSON — those are found within the first 32-byte SSE block
// and return before any AVX2 state is touched. Only when the first 32 bytes hold
// no '"' or '\\' (a long string, e.g. a unicode text field) does the routine
// switch to AVX2 inside indexQuoteOrBackslashSSE2: one 32-byte compare per
// iteration instead of two 16-byte ones, paying the single VZEROUPPER over many
// bytes. So short strings keep the SSE2 cost and long strings get AVX2 throughput
// (string_unicode −9%, no regression on cloudflare/pretty/time-array).
//
// The routine handles every length itself (the 32- and 16-byte loops fall
// through to a scalar tail for the final <16 bytes), so the dispatch is a single
// call with no length branch. That keeps indexCloseOrEscape inlinable into its
// callers (ReadKey, the Read*String funcs, SkipString), removing a call layer
// from the hottest path in object decoding.
func indexCloseOrEscape(b []byte) int {
	return indexQuoteOrBackslashSSE2(b)
}

// structuralPrescan is how many leading bytes indexStructural scans with the
// scalar loop before falling back to the AVX2 routine. Unlike indexCloseOrEscape
// (which the decode hot path calls for every key and string value), indexStructural
// is reached only while skipping a value (skipObject/skipArray), where the input
// is typically token-dense JSON whose next structural byte is only a few bytes
// away. A short scalar scan finds it without the per-call cost of the AVX2 routine
// (the Go→assembly boundary and VZEROUPPER), which is not amortized when the
// routine early-exits after a single block. Only long runs survive the prescan and
// reach the SIMD pass. This stays out of the decode path entirely, so it cannot
// regress the generated unmarshalers.
const structuralPrescan = 16

// indexStructural returns the index of the first '{', '}', '[', ']' or '"' byte
// in b, or len(b) if none is present.
func indexStructural(b []byte) int {
	if !useAVX2 || len(b) < 32 {
		return indexStructuralScalar(b)
	}
	for i, c := range b[:structuralPrescan] {
		switch c {
		case '{', '}', '[', ']', '"':
			return i
		}
	}
	return structuralPrescan + indexStructuralAVX2(b[structuralPrescan:])
}
