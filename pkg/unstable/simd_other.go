package unstable

import "bytes"

// The generic, architecture-independent scalar implementations of the SIMD scan
// primitives. The amd64 (SSE2/AVX2) and arm64 (NEON) dispatchers fall back to
// these for short buffers, and on architectures with no SIMD version they are the
// sole implementation (see simd_noasm.go). TestIndexFunctionsMatchScalar checks
// the SIMD routines against them.

// indexCloseOrEscapeScalar is the portable fallback for indexCloseOrEscape: it
// returns the index of the first '"' or '\\' in b, or len(b) if neither is
// present, using the runtime's (already SIMD-optimized) bytes.IndexByte.
func indexCloseOrEscapeScalar(b []byte) int {
	q := bytes.IndexByte(b, '"')
	if q < 0 {
		if bs := bytes.IndexByte(b, '\\'); bs >= 0 {
			return bs
		}
		return len(b)
	}
	if bs := bytes.IndexByte(b[:q], '\\'); bs >= 0 {
		return bs
	}
	return q
}

// indexStructuralScalar is the portable fallback for indexStructural: it returns
// the index of the first '{', '}', '[', ']' or '"' in b, or len(b) if none.
func indexStructuralScalar(b []byte) int {
	for i := 0; i < len(b); i++ {
		switch b[i] {
		case '{', '}', '[', ']', '"':
			return i
		}
	}
	return len(b)
}
