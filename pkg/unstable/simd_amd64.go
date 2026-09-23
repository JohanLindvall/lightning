//go:build amd64

package unstable

import (
	"math/bits"

	"golang.org/x/sys/cpu"
)

//go:noescape
func indexQuoteOrBackslashSSE2(b []byte, i int) int

//go:noescape
func indexStructuralAVX2(b []byte, i int) int

//go:noescape
func indexEscapeSSE2(b []byte) int

//go:noescape
func indexEscapeNonASCIISSE2(b []byte) int

var useAVX2 = cpu.X86.HasAVX2

// useStructural512 selects indexStructuralAVX2's AVX-512 VBMI body: VPERMB is
// what lets one table lookup and one compare classify all five structural
// bytes exactly (see structPerm in simd_amd64.s), and BW supplies the byte
// compare into a mask register and the masked tail load. The gate is read in
// the assembly, so the Go dispatch stays a single call.
var useStructural512 = useAVX2 && cpu.X86.HasAVX512F && cpu.X86.HasAVX512BW && cpu.X86.HasAVX512VBMI && cpu.X86.HasBMI2

// indexEscape returns the index of the first byte that JSON string encoding must
// escape — a control byte < 0x20, '"' or '\\' — or len(b) if none. It is a single
// (inlinable) call into indexEscapeSSE2, which begins in SSE2 (baseline, no
// VZEROUPPER), switches to AVX2 once 32 bytes are clean (a long unescaped run, e.g.
// a log line), and drops a sub-block buffer to a scalar tail. The < 0x20 SIMD test
// is a per-block PMINUB(v, 0x1f) == v.
func indexEscape(b []byte) int {
	return indexEscapeSSE2(b)
}

// indexEscapeNonASCII is indexEscape with the predicate widened by non-ASCII
// bytes (>= 0x80) — the scan behind EscapeStringInto's UTF-8 handling. Same
// structure and inlinability as indexEscape; the widening costs one POR of the
// raw chunk per block (PMOVMSKB reads sign bits, which for the raw bytes are
// exactly the non-ASCII lanes).
func indexEscapeNonASCII(b []byte) int {
	return indexEscapeNonASCIISSE2(b)
}

// indexCloseOrEscape returns the index of the first '"' or '\\' byte in b, or
// len(b) if neither is present. String scanning begins in SSE2, not AVX2: SSE2 is
// baseline on every amd64 CPU (no feature gate) and avoids the VZEROUPPER every
// AVX2 call must execute, which is pure overhead for the short keys and string
// values that dominate JSON — those are found within the first 32-byte SSE block
// and return before any AVX2 state is touched. Only when the first 32 bytes hold
// no '"' or '\\' (a long string, e.g. a unicode text field) does the routine
// switch to AVX2 inside indexQuoteOrBackslashSSE2: one 32-byte compare per
// iteration instead of two 16-byte ones, paying the single VZEROUPPER over many
// bytes. So short strings keep the SSE2 cost and long strings get AVX2 throughput
// (string_unicode −9%, no regression on cloudflare/pretty/time-array).
//
// The routine handles every length itself (the 32- and 16-byte loops fall
// through to a scalar tail for the final <16 bytes), so the dispatch is a single
// call with no length branch. That keeps indexCloseOrEscape inlinable into its
// callers (ReadKey, the Read*String funcs, SkipString), removing a call layer
// from the hottest path in object decoding.
func indexCloseOrEscape(b []byte) int {
	return indexQuoteOrBackslashSSE2(b, 0)
}

// indexCloseOrEscapeAt is indexCloseOrEscape starting at i and returning an
// absolute index. The offset is an argument rather than a b[i:] at the call site
// because Go lowers that reslice to seven instructions (len and cap
// subtractions, the negative-length clamp on the base) at every one of them, and
// this scanner runs once per object key and once per string value; DI is the
// index register the assembly already carries, so it simply starts at i.
func indexCloseOrEscapeAt(b []byte, i int) int {
	return indexQuoteOrBackslashSSE2(b, i)
}

// structuralPrescan is how many leading bytes indexStructural scans with the
// scalar loop before falling back to the AVX2 routine. Unlike indexCloseOrEscape
// (which the decode hot path calls for every key and string value), indexStructural
// is reached only while skipping a value (skipObject/skipArray), where the input
// is typically token-dense JSON whose next structural byte is only a few bytes
// away. A short scalar scan finds it without the per-call cost of the AVX2 routine
// (the Go→assembly boundary and VZEROUPPER), which is not amortized when the
// routine early-exits after a single block. Only long runs survive the prescan and
// reach the SIMD pass. This stays out of the decode path entirely, so it cannot
// regress the generated unmarshalers.
const structuralPrescan = 16

// indexStructural returns the index of the first '{', '}', '[', ']' or '"' byte
// in b, or len(b) if none is present.
func indexStructural(b []byte) int { return indexStructuralAt(b, 0) }

// indexStructuralAt is indexStructural starting at i, returning an absolute
// index — the shape every caller wants, since all three of them (skipObject,
// skipArray and the scanner's container balance) wrote i += indexStructural(
// data[i:]) and put i back. That reslice is seven instructions of len and cap
// subtraction and a negative-length clamp, per structural jump, which on an
// array of small values is once or twice an element; taking the offset costs
// nothing but an argument word, and the prescan simply starts there. It is the
// same trade IndexCloseOrEscapeAt already makes for the string scanner, and the
// one reslice left is on the path where the first structuralPrescan bytes hold
// no structural byte at all, where an assembly call is about to be paid anyway.
func indexStructuralAt(b []byte, i int) int {
	if !useAVX2 {
		return i + indexStructuralScalar(b[i:])
	}
	// The prescan is two SWAR words rather than sixteen byte compares; see
	// structuralMask. Both loads are unchecked because the guard has already
	// established structuralPrescan bytes past i. The assembly takes any
	// remainder after it — its tails are an overlapping block (AVX2) or a
	// masked load (VBMI) — where this used to fall back to the byte loop below
	// 32 bytes, at five compares a byte: a 40-byte scan cost 21 ns and now
	// costs 6. Under structuralPrescan bytes only the masked load beats the
	// byte loop; the AVX2 body would load four splats and pay a VZEROUPPER to
	// classify at most fifteen bytes.
	if uint(i)+structuralPrescan <= uint(len(b)) {
		if m := structuralMask(load64(b, i)); m != 0 {
			return i + bits.TrailingZeros64(m)>>3
		}
		if m := structuralMask(load64(b, i+8)); m != 0 {
			return i + 8 + bits.TrailingZeros64(m)>>3
		}
		return indexStructuralAVX2(b, i+structuralPrescan)
	}
	if useStructural512 {
		return indexStructuralAVX2(b, i)
	}
	return i + indexStructuralScalar(b[i:])
}
