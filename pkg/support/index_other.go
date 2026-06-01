//go:build !amd64 && !arm64

package support

// indexCloseOrEscape returns the index of the first '"' or '\\' byte in b, or
// len(b) if neither is present. Platforms without a SIMD implementation use
// the scalar fallback.
func indexCloseOrEscape(b []byte) int {
	return indexCloseOrEscapeScalar(b)
}
