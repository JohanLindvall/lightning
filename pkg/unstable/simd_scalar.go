//go:build !amd64 && !arm64

package unstable

// These platforms have no SIMD implementation, so indexCloseOrEscape and
// indexStructural dispatch to the generic scalar implementations in simd_other.go.

func indexCloseOrEscape(b []byte) int { return indexCloseOrEscapeScalar(b) }

// indexCloseOrEscapeAt is indexCloseOrEscape starting at i, returning an
// absolute index. The SIMD architectures take the offset all the way into the
// assembly (it is the index register the scan already carries, and it saves the
// caller a seven-instruction b[i:] reslice); here it is just the reslice.
func indexCloseOrEscapeAt(b []byte, i int) int {
	if i >= len(b) {
		return len(b) // match the SIMD bodies, which answer len(b) for an empty span
	}
	return i + indexCloseOrEscapeScalar(b[i:])
}

func indexStructural(b []byte) int { return indexStructuralScalar(b) }

func indexEscape(b []byte) int { return indexEscapeScalar(b) }

func indexEscapeNonASCII(b []byte) int { return indexEscapeNonASCIIScalar(b) }
