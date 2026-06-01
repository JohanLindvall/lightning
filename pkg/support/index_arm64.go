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

// indexStructural returns the index of the first '{', '}', '[', ']' or '"' byte
// in b, or len(b) if none is present.
func indexStructural(b []byte) int {
	if useNEON && len(b) >= 16 {
		return indexStructuralNEON(b)
	}
	return indexStructuralScalar(b)
}
