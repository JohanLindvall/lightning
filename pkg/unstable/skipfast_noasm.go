//go:build !amd64 && !arm64

package unstable

// fastSkipAvail is false on arches without a SIMD maskBlock (the scalar one is
// slower than the indexStructural skip), so SkipValue stays on the current path.
const fastSkipAvail = false

// maskBlock returns the character-class bitmaps for the 64 bytes at b[:64] (the
// caller guarantees len(b) >= 64): quote, backslash, and the container's own open
// and close brackets — `[`/`]` when isArray, else `{`/`}`. Bit j is set when b[j]
// is that byte. Scalar fallback for arches without a SIMD routine.
func maskBlock(b []byte, isArray bool) (quote, bslash, open, close uint64) {
	openByte, closeByte := byte('{'), byte('}')
	if isArray {
		openByte, closeByte = '[', ']'
	}
	_ = b[63]
	for j := 0; j < 64; j++ {
		switch b[j] {
		case '"':
			quote |= uint64(1) << uint(j)
		case '\\':
			bslash |= uint64(1) << uint(j)
		case openByte:
			open |= uint64(1) << uint(j)
		case closeByte:
			close |= uint64(1) << uint(j)
		}
	}
	return
}
