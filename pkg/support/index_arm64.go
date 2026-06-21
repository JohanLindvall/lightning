//go:build arm64

package support

//go:noescape
func indexQuoteOrBackslashNEON(b []byte) int

//go:noescape
func indexStructuralNEON(b []byte) int

// indexCloseOrEscape returns the index of the first '"' or '\\' byte in b, or
// len(b) if neither is present. The NEON routine handles every length itself —
// its 16-byte loop falls through to a scalar tail for the final <16 bytes — so
// the dispatch is a single unconditional call with no length branch and no
// feature gate (Advanced SIMD is mandatory in the ARMv8-A baseline Go targets).
// That keeps indexCloseOrEscape inlinable into its callers (ReadKey, the
// Read*String funcs, skipString, decodeEscaped), removing a call layer from the
// hottest function in object decoding, where the scanner is ~40% of the work.
func indexCloseOrEscape(b []byte) int {
	return indexQuoteOrBackslashNEON(b)
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
	if len(b) < structuralPrescan+16 {
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
