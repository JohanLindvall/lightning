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

// SwarNeedsEscape reports (nonzero, high bit per matching lane) which of the
// eight packed bytes in v JSON string encoding must escape — a control byte
// < 0x20, '"' or '\\'. It is the one shared spelling of that predicate:
// indexEscapeScalar's clean-run scan and pkg/json's escape-walk per-run probe
// (escapeValidInto, and via SwarNeedsEscapeOrNonASCII below, EscapeStringInto)
// all build on it, so the escape byte set lives in one place. Pure bit math
// with no calls, so it inlines into all of them (the escape walks' gates depend
// on that — re-check -gcflags=-m if this grows).
func SwarNeedsEscape(v uint64) uint64 {
	return escSwarHasLess(v, 0x20) | escSwarHasByte(v, '"') | escSwarHasByte(v, '\\')
}

// indexEscapeScalar is the portable fallback for indexEscape: it returns the index
// of the first byte that JSON string encoding must escape — a control byte < 0x20,
// '"' or '\\' — or len(b) if none, scanning eight bytes at a time via SWAR.
func indexEscapeScalar(b []byte) int {
	i := 0
	for ; i+8 <= len(b); i += 8 {
		v := binary.LittleEndian.Uint64(b[i : i+8])
		if SwarNeedsEscape(v) != 0 {
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

// SwarNeedsEscapeOrNonASCII is SwarNeedsEscape widened by the non-ASCII lanes
// (high bit set): a nonzero result means some lane JSON string encoding must
// escape OR is >= 0x80, and the LOWEST set bit marks the first such lane. It is
// the predicate behind EscapeStringInto's UTF-8 handling: the escape walk runs
// on it until the first non-ASCII byte decides (once, via utf8.Valid) whether
// the rest of the input is clean UTF-8 or needs U+FFFD substitution.
//
// Contract: only the LOWEST set bit is meaningful — higher bits may be false
// positives. The has-less trick's `&^ v` term exists to keep a borrow out of a
// lane whose own value is >= 0x20 from flagging it; dropping it and OR-ing v
// instead makes the widened predicate one op CHEAPER than SwarNeedsEscape
// itself ((v-0x20·lo)|v vs (v-0x20·lo)&^v, then &hi either way). The dropped
// term readmits exactly two lane kinds: high-bit lanes (wanted) and a lane
// whose 0x20-subtraction borrowed from a LOWER underflowing lane — and such a
// lower lane is < 0x20, i.e. a true match below the false positive, so
// TrailingZeros64 never lands on the false one. Both callers (EscapeStringInto's
// per-run probe, indexEscapeNonASCIIScalar's word loop) take only the first
// match. Same inlining constraint as SwarNeedsEscape: pure bit math, no calls.
func SwarNeedsEscapeOrNonASCII(v uint64) uint64 {
	return ((v-escSwarLo*0x20)|v)&escSwarHi | escSwarHasByte(v, '"') | escSwarHasByte(v, '\\')
}

// indexEscapeNonASCIIScalar is the portable fallback for indexEscapeNonASCII:
// indexEscapeScalar's scan with the widened SwarNeedsEscapeOrNonASCII predicate —
// the first escape byte or non-ASCII byte (>= 0x80), or len(b) if neither.
func indexEscapeNonASCIIScalar(b []byte) int {
	i := 0
	for ; i+8 <= len(b); i += 8 {
		v := binary.LittleEndian.Uint64(b[i : i+8])
		if SwarNeedsEscapeOrNonASCII(v) != 0 {
			break
		}
	}
	for ; i < len(b); i++ {
		if c := b[i]; c < 0x20 || c >= 0x80 || c == '"' || c == '\\' {
			return i
		}
	}
	return i
}

// The generic, architecture-independent scalar implementations of the SIMD scan
// primitives. The amd64 (SSE2/AVX2) and arm64 (NEON, or SVE2 where the core has
// it) dispatchers fall back to these for short buffers, and on architectures with
// no SIMD version they are the sole implementation (simd_scalar.go is the dispatch
// file that routes to them). TestIndexFunctionsMatchScalar checks the SIMD
// routines against them, and the per-arch TestIndexVariantsFlip tests each
// dispatch arm separately — a scanner with more than one vector body would
// otherwise have only the host's preferred one covered.

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
