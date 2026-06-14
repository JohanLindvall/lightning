//go:build amd64

package support

import "golang.org/x/sys/cpu"

//go:noescape
func indexQuoteOrBackslashAVX2(b []byte) int

//go:noescape
func indexStructuralAVX2(b []byte) int

var useAVX2 = cpu.X86.HasAVX2

// indexCloseOrEscape returns the index of the first '"' or '\\' byte in b, or
// len(b) if neither is present. On AVX2 hardware, strings of at least 32 bytes
// are scanned in a single SIMD pass; shorter inputs and non-AVX2 CPUs use the
// scalar (bytes.IndexByte-based) fallback.
func indexCloseOrEscape(b []byte) int {
	if useAVX2 && len(b) >= 32 {
		return indexQuoteOrBackslashAVX2(b)
	}
	return indexCloseOrEscapeScalar(b)
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
