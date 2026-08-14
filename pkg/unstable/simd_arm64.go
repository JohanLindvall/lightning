//go:build arm64

package unstable

import "golang.org/x/sys/cpu"

// Three of the four NEON routines below have no Go caller: the SVE2 entry points
// in simd_arm64.s reach them with a tail branch (`B ·indexEscapeNEON(SB)`) when
// useSVE2 is clear, which is what keeps the Go dispatch a single unconditional
// call. The `unused` linter reads Go only, so it sees a dead declaration and says
// so — a blind spot of the same family as asmdecl's, not a finding. The
// declarations themselves are load-bearing regardless of who calls them: asmdecl
// validates the assembly's frame offsets against exactly these signatures, and
// `go vet` reports assembly with no Go prototype.
//
// (CI never sees this, because the Lint step runs on amd64 only, where all four
// are behind a build tag. It is a real gap — `unused` is architecture-dependent
// wherever per-arch files are — and was checked by hand here.)

//nolint:unused // called from assembly; see above
//go:noescape
func indexQuoteOrBackslashNEON(b []byte) int

//go:noescape
func indexStructuralNEON(b []byte) int

//nolint:unused // called from assembly; see above
//go:noescape
func indexEscapeNEON(b []byte) int

//nolint:unused // called from assembly; see above
//go:noescape
func indexEscapeNonASCIINEON(b []byte) int

// The three entry points below each read useSVE2 in assembly and run an SVE2
// body or tail-call the NEON routine above it; see the SVE2 section of
// simd_arm64.s for why the gate lives there rather than in Go.

//go:noescape
func indexQuoteOrBackslashArm64(b []byte) int

//go:noescape
func indexEscapeArm64(b []byte) int

//go:noescape
func indexEscapeNonASCIIArm64(b []byte) int

//go:noescape
func indexStructuralSVE2(b []byte) int

// useSVE2 selects the SVE2 body of each scanner (see the SVE2 section of
// simd_arm64.s) on cores that implement SVE2 — Neoverse N2/V2, Graviton4 and
// later — leaving Apple M-series, Neoverse N1/V1 and every other ARMv8-A core on
// the NEON baseline. MATCH is the only SVE2 (rather than SVE) instruction those
// bodies use, so SVE2 is the right gate.
//
// It is a var rather than a const so TestIndexVariantsFlip can drive both bodies
// on one machine, the way useAVX2 and useSkipBlocks are used elsewhere here.
var useSVE2 = cpu.ARM64.HasSVE2

// indexCloseOrEscape returns the index of the first '"' or '\\' byte in b, or
// len(b) if neither is present. Both bodies behind indexQuoteOrBackslashArm64
// handle every length themselves — the NEON one falls through to a scalar tail
// for the final <16 bytes, the SVE2 one covers the ragged end with its WHILELO
// predicate — and the SVE2-vs-NEON choice is made inside that assembly, so this
// dispatch stays a single unconditional call with no branch of its own.
//
// That is what keeps indexCloseOrEscape inlinable into its callers (ReadKey, the
// Read*String funcs, SkipString, decodeEscaped) and into every generated
// decoder, removing a call layer from the hottest function in object decoding,
// where the scanner is ~40% of the work. Spelling the feature gate here instead
// would cost two calls (124 against the inliner's budget of 80) and lose that —
// re-check `-gcflags=-m` after editing here, the inline is load-bearing.
func indexCloseOrEscape(b []byte) int {
	return indexQuoteOrBackslashArm64(b)
}

// indexEscape returns the index of the first byte JSON string encoding must escape
// (control byte < 0x20, '"' or '\\'), or len(b) if none. Both bodies handle every
// length themselves and pick themselves in assembly, so this is a single
// unconditional call — inlinable into EscapeStringInto's clean-run scan, which is
// what the per-run vector/SWAR gate there depends on. The < 0x20 lane test is
// VUMIN(chunk, 0x1f) == chunk under NEON (its form of amd64's PMINUB(v, 0x1f) ==
// v) and one unsigned CMPLO against #32 under SVE2.
func indexEscape(b []byte) int {
	return indexEscapeArm64(b)
}

// indexEscapeNonASCII is indexEscape with the predicate widened by non-ASCII
// bytes (>= 0x80) — the scan behind EscapeStringInto's UTF-8 handling. Same
// structure and inlinability as indexEscape; the widening is VUSHR $7 + VORR of
// the raw chunk per block under NEON, and under SVE2 is folded into the
// control-byte test as one unsigned "outside [0x20, 0x7f]" compare (see the asm,
// where that fold is what keeps the loop off the predicate pipe's limit).
func indexEscapeNonASCII(b []byte) int {
	return indexEscapeNonASCIIArm64(b)
}

// structuralPrescan is how many leading bytes indexStructural scans with the
// scalar loop before falling back to the vector routine. Unlike indexCloseOrEscape
// (which the decode hot path calls for every key and string value), indexStructural
// is reached only while skipping a value (skipObject/skipArray), where the input
// is typically token-dense JSON whose next structural byte is only a few bytes
// away. A short scalar scan finds it without the per-call cost of the vector
// routine (the Go→assembly boundary and, under NEON, the 48-byte nibble-table
// load), which is not amortized when the routine early-exits after a single
// 16-byte block. Only long runs survive the prescan and reach the SIMD pass. This
// stays out of the decode path entirely, so it cannot regress the generated
// unmarshalers.
const structuralPrescan = 16

// indexStructural returns the index of the first '{', '}', '[', ']' or '"' byte
// in b, or len(b) if none is present.
//
// This is the one arm64 scanner whose SVE2 gate is an ordinary Go branch rather
// than a load inside the assembly: the function already carries a length test
// and the scalar prescan loop, so it is far past the inline budget either way
// and the branch costs nothing. The three scanners the decode path calls per
// key and per string value cannot afford that — see indexCloseOrEscape.
func indexStructural(b []byte) int {
	if len(b) < structuralPrescan+16 {
		return indexStructuralScalar(b)
	}
	for i, c := range b[:structuralPrescan] {
		switch c {
		case '{', '}', '[', ']', '"':
			return i
		}
	}
	if useSVE2 {
		return structuralPrescan + indexStructuralSVE2(b[structuralPrescan:])
	}
	return structuralPrescan + indexStructuralNEON(b[structuralPrescan:])
}
