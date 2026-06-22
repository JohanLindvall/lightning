//go:build !amd64 && !arm64

package unstable

// fastSkipAvail is false on arches without a SIMD maskBlock (the scalar one is
// slower than the indexStructural skip), so SkipValue stays on the current path.
const fastSkipAvail = false

// maskBlock returns the character-class bitmaps for the 32 bytes at b[:32] (the
// caller guarantees len(b) >= 32). Bit j of each result is set when b[j] is the
// corresponding byte. Scalar fallback for arches without the AVX2 routine.
func maskBlock(b []byte) (quote, bslash, lbrace, rbrace, lbrack, rbrack uint32) {
	_ = b[31]
	for j := 0; j < 32; j++ {
		switch b[j] {
		case '"':
			quote |= uint32(1) << uint(j)
		case '\\':
			bslash |= uint32(1) << uint(j)
		case '{':
			lbrace |= uint32(1) << uint(j)
		case '}':
			rbrace |= uint32(1) << uint(j)
		case '[':
			lbrack |= uint32(1) << uint(j)
		case ']':
			rbrack |= uint32(1) << uint(j)
		}
	}
	return
}
