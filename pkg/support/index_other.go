//go:build !amd64 && !arm64

package support

// These platforms have no SIMD implementation, so every primitive uses its
// scalar fallback.

func indexCloseOrEscape(b []byte) int { return indexCloseOrEscapeScalar(b) }

func indexStructural(b []byte) int { return indexStructuralScalar(b) }
