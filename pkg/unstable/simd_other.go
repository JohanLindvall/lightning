package unstable

import (
	"bytes"
	"encoding/binary"
)

// SWAR (SIMD-within-a-register) helpers operating on eight packed bytes at once,
// used by indexEscapeScalar. escSwarHasLess reports (nonzero, high bit per lane)
// whether any byte lane of v is < n (1 <= n <= 128); escSwarHasByte reports
// whether any lane equals b.
const (
	escSwarLo = 0x0101010101010101 // 1 in every byte lane
	escSwarHi = 0x8080808080808080 // high bit of every byte lane
)

func escSwarHasLess(v, n uint64) uint64 { return (v - escSwarLo*n) & ^v & escSwarHi }
func escSwarHasByte(v uint64, b byte) uint64 {
	x := v ^ (escSwarLo * uint64(b))
	return (x - escSwarLo) & ^x & escSwarHi
}

// indexEscapeScalar is the portable fallback for indexEscape: it returns the index
// of the first byte that JSON string encoding must escape — a control byte < 0x20,
// '"' or '\\' — or len(b) if none, scanning eight bytes at a time via SWAR.
func indexEscapeScalar(b []byte) int {
	i := 0
	for ; i+8 <= len(b); i += 8 {
		v := binary.LittleEndian.Uint64(b[i:])
		if escSwarHasLess(v, 0x20)|escSwarHasByte(v, '"')|escSwarHasByte(v, '\\') != 0 {
			break
		}
	}
	for ; i < len(b); i++ {
		if c := b[i]; c < 0x20 || c == '"' || c == '\\' {
			return i
		}
	}
	return i
}

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
