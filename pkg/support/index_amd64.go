//go:build amd64

package support

import "golang.org/x/sys/cpu"

//go:noescape
func indexQuoteOrBackslashAVX2(b []byte) int

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
