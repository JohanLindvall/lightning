//go:build !amd64 && !arm64

package unstable

// These platforms have no SIMD implementation, so indexCloseOrEscape and
// indexStructural dispatch to the generic scalar implementations in simd_other.go.

func indexCloseOrEscape(b []byte) int { return indexCloseOrEscapeScalar(b) }

func indexStructural(b []byte) int { return indexStructuralScalar(b) }

func indexEscape(b []byte) int { return indexEscapeScalar(b) }
