//go:build arm64

package support

import "golang.org/x/sys/cpu"

//go:noescape
func indexQuoteOrBackslashNEON(b []byte) int

//go:noescape
func indexStructuralNEON(b []byte) int

var useNEON = cpu.ARM64.HasASIMD

// indexCloseOrEscape returns the index of the first '"' or '\\' byte in b, or
// len(b) if neither is present. On arm64 with Advanced SIMD, strings of at
// least 16 bytes are scanned in a single NEON pass; shorter inputs use the
// scalar (bytes.IndexByte-based) fallback.
func indexCloseOrEscape(b []byte) int {
	if useNEON && len(b) >= 16 {
		return indexQuoteOrBackslashNEON(b)
	}
	return indexCloseOrEscapeScalar(b)
}

// structuralPrescan is how many leading bytes indexStructural scans with the
// scalar loop before falling back to the NEON routine. Unlike indexCloseOrEscape
// (which the decode hot path calls for every key and string value), indexStructural
// is reached only while skipping a value (skipObject/skipArray), where the input
// is typically token-dense JSON whose next structural byte is only a few bytes
// away. A short scalar scan finds it without the per-call cost of the NEON routine
// (the Go→assembly boundary and the 80-byte structMask load), which is not
// amortized when the routine early-exits after a single 16-byte block. Only long
// runs survive the prescan and reach the SIMD pass. This stays out of the decode
// path entirely, so it cannot regress the generated unmarshalers.
const structuralPrescan = 16

// indexStructural returns the index of the first '{', '}', '[', ']' or '"' byte
// in b, or len(b) if none is present.
func indexStructural(b []byte) int {
	if !useNEON || len(b) < structuralPrescan+16 {
		return indexStructuralScalar(b)
	}
	for i, c := range b[:structuralPrescan] {
		switch c {
		case '{', '}', '[', ']', '"':
			return i
		}
	}
	return structuralPrescan + indexStructuralNEON(b[structuralPrescan:])
}
