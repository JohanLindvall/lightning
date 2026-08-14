#include "textflag.h"

// No 16-byte index constant is needed for position recovery: scalar RBIT/CLZ on
// the result mask moved into general registers finds the first nonzero byte.
//
// indexQuoteOrBackslashNEON builds its two comparison splats ('"', '\\') per call
// with VMOVI rather than loading a RODATA table: on amd64 the equivalent constants
// live in RODATA because building them inline needs VPBROADCASTB (a GP→XMM domain
// crossing that dominates a single-block call), but VMOVI materializes the splat
// inside the vector unit from an 8-bit immediate, so the inline form costs one
// instruction with no load-use latency *and* no domain crossing at all.
//
// This replaced a MOVD-immediate + VDUP pair per splat, which was already cheaper
// than a RODATA VLD1 (~2% on the get benchmark) but still spent two instructions
// and a GP→SIMD transfer per constant. VMOVI is a strict improvement on both: the
// splats no longer depend on any general register, so they are ready before the
// argument loads retire. Verified byte-identical to the VDUP form (see
// TestIndexFunctionsMatchScalar and the exhaustive oracle tests in simd_test.go).
//
// Each scanner also peels its first 16-byte block out of the loop. The scanners
// are called on `string-body + rest-of-document`, and the overwhelmingly common
// outcome is a match inside that first block (median JSON string length is 8-16
// bytes across the corpus), so the peeled copy loads straight from the base
// pointer with no offset arithmetic, no loop-bound recomputation and no counter
// update — 20 executed instructions instead of 26 for the dominant case. Only a
// string that survives 16 clean bytes pays for the loop counters, and the loop
// itself hoists its bound (len-16) out rather than recomputing `remaining` per
// iteration.
//
// indexStructuralNEON instead classifies its five target bytes with simdjson's
// shuffle trick — two TBL (NEON's VPSHUFB) lookups into the nibble tables below,
// structLo[lowNibble] & structHi[highNibble] != 0 — which a single-byte VDUP
// cannot build, so those tables are kept in RODATA and loaded once per call (the
// scalar prescan means the NEON loop, hence this load, runs only for long skips).
// Two bits encode the groups: '"' (lo 0x2 / hi 0x2) and the brackets/braces
// (lo 0xB|0xD / hi 0x5|0x7). The 0x0f mask isolates the low nibble.
DATA structTablesArm<>+0(SB)/8, $0x0000000000010000  // structLo nibbles 0–7
DATA structTablesArm<>+8(SB)/8, $0x0000020002000000  // structLo nibbles 8–15
DATA structTablesArm<>+16(SB)/8, $0x0200020000010000 // structHi nibbles 0–7
DATA structTablesArm<>+24(SB)/8, $0x0000000000000000 // structHi nibbles 8–15
DATA structTablesArm<>+32(SB)/8, $0x0f0f0f0f0f0f0f0f // low-nibble mask
DATA structTablesArm<>+40(SB)/8, $0x0f0f0f0f0f0f0f0f
GLOBL structTablesArm<>(SB), RODATA|NOPTR, $48

// func indexQuoteOrBackslashNEON(b []byte) int
//
// Returns the index of the first '"' or '\\' byte in b, or len(b) if neither
// is present. Scans 16 bytes per iteration with NEON, then a scalar tail.
TEXT ·indexQuoteOrBackslashNEON(SB), NOSPLIT, $0-32
	MOVD  b_base+0(FP), R0
	MOVD  b_len+8(FP), R1
	VMOVI $0x22, V0.B16          // V0 = '"' x16
	VMOVI $0x5c, V1.B16          // V1 = '\\' x16
	CMP   $16, R1
	BLT   shortInput

	// Peeled first block (see the file header): loads at the base pointer, so
	// no offset register, no bound, no address ADD on the critical path.
	VLD1  (R0), [V2.B16]
	VCMEQ V0.B16, V2.B16, V3.B16 // V3 = (chunk == '"')
	VCMEQ V1.B16, V2.B16, V4.B16 // V4 = (chunk == '\\')
	VORR  V3.B16, V4.B16, V3.B16 // V3 = either match (0xFF per lane)
	VMOV  V3.D[0], R9            // low 8 lanes
	VMOV  V3.D[1], R10           // high 8 lanes
	CBNZ  R9, firstLow8
	CBNZ  R10, firstHigh8

	MOVD  $16, R2                // R2 = current offset
	SUB   $16, R1, R12           // R12 = last offset holding a full block

loop16:
	CMP  R12, R2
	BGT  tail                    // fewer than 16 bytes left
	ADD  R0, R2, R8              // R8 = &b[offset]
	VLD1 (R8), [V2.B16]
	VCMEQ V0.B16, V2.B16, V3.B16
	VCMEQ V1.B16, V2.B16, V4.B16
	VORR V3.B16, V4.B16, V3.B16
	VMOV V3.D[0], R9
	VMOV V3.D[1], R10
	CBNZ R9, low8
	CBNZ R10, high8
	ADD  $16, R2
	B    loop16

	// The peeled block's recoveries are separate from the loop's so that the
	// common case adds no offset (its block starts at 0).
firstLow8:
	RBIT R9, R11
	CLZ  R11, R11                // trailing zeros of R9 = first set bit
	LSR  $3, R11, R11            // /8 -> first matching byte (lane 0..7)
	MOVD R11, ret+24(FP)
	RET

firstHigh8:
	RBIT R10, R11
	CLZ  R11, R11
	LSR  $3, R11, R11            // lane within high half (0..7)
	ADD  $8, R11, R11            // lanes 8..15
	MOVD R11, ret+24(FP)
	RET

low8:
	RBIT R9, R11
	CLZ  R11, R11                // trailing zeros of R9 = first set bit
	LSR  $3, R11, R11            // /8 -> first matching byte (lane 0..7)
	ADD  R2, R11, R11
	MOVD R11, ret+24(FP)
	RET

high8:
	RBIT R10, R11
	CLZ  R11, R11
	LSR  $3, R11, R11            // lane within high half (0..7)
	ADD  $8, R11, R11            // lanes 8..15
	ADD  R2, R11, R11
	MOVD R11, ret+24(FP)
	RET

shortInput:
	MOVD ZR, R2                  // buffer shorter than one block: scan from 0

tail:
	CMP  R1, R2
	BGE  notfound

tailloop:
	ADD   R0, R2, R8
	MOVBU (R8), R9
	CMP   $0x22, R9
	BEQ   tfound
	CMP   $0x5c, R9
	BEQ   tfound
	ADD   $1, R2
	CMP   R1, R2
	BLT   tailloop

notfound:
	MOVD R1, ret+24(FP)
	RET

tfound:
	MOVD R2, ret+24(FP)
	RET

// func indexEscapeNEON(b []byte) int
//
// Returns the index of the first byte JSON string encoding must escape — '"'
// (0x22), '\\' (0x5c) or a control byte < 0x20 — or len(b) if none. Mirrors
// indexQuoteOrBackslashNEON (16 bytes/iter, scalar tail) with one extra per-block
// test for control bytes: VUMIN(chunk, 0x1f) == chunk, true exactly when the lane
// is <= 0x1f (the NEON form of amd64's PMINUB(v, 0x1f) == v). The three compare
// splats are built with VMOVI, no RODATA load and no GP→SIMD transfer; the first
// block is peeled as in indexQuoteOrBackslashNEON (see the file header).
TEXT ·indexEscapeNEON(SB), NOSPLIT, $0-32
	MOVD  b_base+0(FP), R0
	MOVD  b_len+8(FP), R1
	VMOVI $0x22, V0.B16          // V0 = '"' x16
	VMOVI $0x5c, V1.B16          // V1 = '\\' x16
	VMOVI $0x1f, V5.B16          // V5 = 0x1f x16 (largest control byte)
	CMP   $16, R1
	BLT   eshortInput

	VLD1  (R0), [V2.B16]
	VCMEQ V0.B16, V2.B16, V3.B16 // V3 = (chunk == '"')
	VCMEQ V1.B16, V2.B16, V4.B16 // V4 = (chunk == '\\')
	VORR  V4.B16, V3.B16, V3.B16
	VUMIN V5.B16, V2.B16, V4.B16 // V4 = min(chunk, 0x1f)
	VCMEQ V4.B16, V2.B16, V4.B16 // V4 = (chunk == min) -> control byte (<= 0x1f)
	VORR  V4.B16, V3.B16, V3.B16 // V3 = any of the three (0xFF per matching lane)
	VMOV  V3.D[0], R9            // low 8 lanes
	VMOV  V3.D[1], R10           // high 8 lanes
	CBNZ  R9, efirstLow8
	CBNZ  R10, efirstHigh8

	MOVD  $16, R2                // R2 = current offset
	SUB   $16, R1, R12           // R12 = last offset holding a full block

eloop16:
	CMP  R12, R2
	BGT  etail                   // fewer than 16 bytes left
	ADD  R0, R2, R8             // R8 = &b[offset]
	VLD1 (R8), [V2.B16]
	VCMEQ V0.B16, V2.B16, V3.B16
	VCMEQ V1.B16, V2.B16, V4.B16
	VORR  V4.B16, V3.B16, V3.B16
	VUMIN V5.B16, V2.B16, V4.B16
	VCMEQ V4.B16, V2.B16, V4.B16
	VORR  V4.B16, V3.B16, V3.B16
	VMOV V3.D[0], R9
	VMOV V3.D[1], R10
	CBNZ R9, elow8
	CBNZ R10, ehigh8
	ADD  $16, R2
	B    eloop16

efirstLow8:
	RBIT R9, R11
	CLZ  R11, R11               // trailing zeros of R9 = first set bit
	LSR  $3, R11, R11           // /8 -> first matching byte (lane 0..7)
	MOVD R11, ret+24(FP)
	RET

efirstHigh8:
	RBIT R10, R11
	CLZ  R11, R11
	LSR  $3, R11, R11           // lane within high half (0..7)
	ADD  $8, R11, R11           // lanes 8..15
	MOVD R11, ret+24(FP)
	RET

elow8:
	RBIT R9, R11
	CLZ  R11, R11               // trailing zeros of R9 = first set bit
	LSR  $3, R11, R11           // /8 -> first matching byte (lane 0..7)
	ADD  R2, R11, R11
	MOVD R11, ret+24(FP)
	RET

ehigh8:
	RBIT R10, R11
	CLZ  R11, R11
	LSR  $3, R11, R11           // lane within high half (0..7)
	ADD  $8, R11, R11           // lanes 8..15
	ADD  R2, R11, R11
	MOVD R11, ret+24(FP)
	RET

eshortInput:
	MOVD ZR, R2                 // buffer shorter than one block: scan from 0

etail:
	CMP  R1, R2
	BGE  enotfound

etailloop:
	ADD   R0, R2, R8
	MOVBU (R8), R9
	CMP   $0x20, R9
	BLO   etfound              // control byte < 0x20 (unsigned)
	CMP   $0x22, R9
	BEQ   etfound
	CMP   $0x5c, R9
	BEQ   etfound
	ADD   $1, R2
	CMP   R1, R2
	BLT   etailloop

enotfound:
	MOVD R1, ret+24(FP)
	RET

etfound:
	MOVD R2, ret+24(FP)
	RET

// func indexStructuralNEON(b []byte) int
//
// Returns the index of the first '{', '}', '[', ']' or '"' byte, or len(b).
TEXT ·indexStructuralNEON(SB), NOSPLIT, $0-32
	MOVD b_base+0(FP), R0
	MOVD b_len+8(FP), R1
	MOVD $structTablesArm<>(SB), R3
	VLD1 (R3), [V0.B16, V1.B16, V2.B16] // structLo, structHi, 0x0f mask
	CMP  $16, R1
	BLT  sshortInput

	VLD1  (R0), [V5.B16]              // chunk
	VAND  V2.B16, V5.B16, V6.B16     // low nibbles
	VTBL  V6.B16, [V0.B16], V6.B16   // structLo[lowNibble]
	VUSHR $4, V5.B16, V7.B16         // high nibbles (per-byte shift)
	VTBL  V7.B16, [V1.B16], V7.B16   // structHi[highNibble]
	VAND  V7.B16, V6.B16, V6.B16     // nonzero byte where structural
	VMOV  V6.D[0], R9
	VMOV  V6.D[1], R10
	CBNZ  R9, sfirstLow8
	CBNZ  R10, sfirstHigh8

	MOVD  $16, R2
	SUB   $16, R1, R12                // last offset holding a full block

sloop:
	CMP  R12, R2
	BGT  stail
	ADD  R0, R2, R8
	VLD1  (R8), [V5.B16]
	VAND  V2.B16, V5.B16, V6.B16
	VTBL  V6.B16, [V0.B16], V6.B16
	VUSHR $4, V5.B16, V7.B16
	VTBL  V7.B16, [V1.B16], V7.B16
	VAND  V7.B16, V6.B16, V6.B16
	VMOV V6.D[0], R9
	VMOV V6.D[1], R10
	CBNZ R9, slow8
	CBNZ R10, shigh8
	ADD  $16, R2
	B    sloop

sfirstLow8:
	RBIT R9, R11
	CLZ  R11, R11
	LSR  $3, R11, R11
	MOVD R11, ret+24(FP)
	RET

sfirstHigh8:
	RBIT R10, R11
	CLZ  R11, R11
	LSR  $3, R11, R11
	ADD  $8, R11, R11
	MOVD R11, ret+24(FP)
	RET

slow8:
	RBIT R9, R11
	CLZ  R11, R11
	LSR  $3, R11, R11
	ADD  R2, R11, R11
	MOVD R11, ret+24(FP)
	RET

shigh8:
	RBIT R10, R11
	CLZ  R11, R11
	LSR  $3, R11, R11
	ADD  $8, R11, R11
	ADD  R2, R11, R11
	MOVD R11, ret+24(FP)
	RET

sshortInput:
	MOVD ZR, R2                       // buffer shorter than one block: scan from 0

stail:
	CMP  R1, R2
	BGE  snone
stloop:
	ADD   R0, R2, R8
	MOVBU (R8), R9
	CMP   $0x7b, R9
	BEQ   stf
	CMP   $0x7d, R9
	BEQ   stf
	CMP   $0x5b, R9
	BEQ   stf
	CMP   $0x5d, R9
	BEQ   stf
	CMP   $0x22, R9
	BEQ   stf
	ADD  $1, R2
	CMP  R1, R2
	BLT  stloop
snone:
	MOVD R1, ret+24(FP)
	RET
stf:
	MOVD R2, ret+24(FP)
	RET

// func indexEscapeNonASCIINEON(b []byte) int
//
// indexEscapeNEON's scan with the predicate widened by non-ASCII bytes: returns
// the index of the first '"' (0x22), '\\' (0x5c), control byte < 0x20 OR byte
// >= 0x80, or len(b) if none. The widening is VUSHR $7 of the raw chunk (0x01
// per non-ASCII lane, no splat needed) VORR-ed into the match vector — the
// RBIT/CLZ recovery finds the first nonzero LANE, and a 0x01 lane's set bit
// stays inside its lane, so any nonzero marker value recovers the same index as
// the 0xFF compare lanes. Everything else is byte-for-byte indexEscapeNEON.
TEXT ·indexEscapeNonASCIINEON(SB), NOSPLIT, $0-32
	MOVD  b_base+0(FP), R0
	MOVD  b_len+8(FP), R1
	VMOVI $0x22, V0.B16          // V0 = '"' x16
	VMOVI $0x5c, V1.B16          // V1 = '\\' x16
	VMOVI $0x1f, V5.B16          // V5 = 0x1f x16 (largest control byte)
	CMP   $16, R1
	BLT   eushortInput

	VLD1  (R0), [V2.B16]
	VCMEQ V0.B16, V2.B16, V3.B16 // V3 = (chunk == '"')
	VCMEQ V1.B16, V2.B16, V4.B16 // V4 = (chunk == '\\')
	VORR  V4.B16, V3.B16, V3.B16
	VUMIN V5.B16, V2.B16, V4.B16 // V4 = min(chunk, 0x1f)
	VCMEQ V4.B16, V2.B16, V4.B16 // V4 = (chunk == min) -> control byte (<= 0x1f)
	VORR  V4.B16, V3.B16, V3.B16
	VUSHR $7, V2.B16, V4.B16     // V4 = 0x01 per non-ASCII lane (high bit set)
	VORR  V4.B16, V3.B16, V3.B16 // V3 = any of the four (nonzero per matching lane)
	VMOV  V3.D[0], R9            // low 8 lanes
	VMOV  V3.D[1], R10           // high 8 lanes
	CBNZ  R9, eufirstLow8
	CBNZ  R10, eufirstHigh8

	MOVD  $16, R2                // R2 = current offset
	SUB   $16, R1, R12           // R12 = last offset holding a full block

euloop16:
	CMP  R12, R2
	BGT  eutail                  // fewer than 16 bytes left
	ADD  R0, R2, R8             // R8 = &b[offset]
	VLD1 (R8), [V2.B16]
	VCMEQ V0.B16, V2.B16, V3.B16
	VCMEQ V1.B16, V2.B16, V4.B16
	VORR  V4.B16, V3.B16, V3.B16
	VUMIN V5.B16, V2.B16, V4.B16
	VCMEQ V4.B16, V2.B16, V4.B16
	VORR  V4.B16, V3.B16, V3.B16
	VUSHR $7, V2.B16, V4.B16
	VORR  V4.B16, V3.B16, V3.B16
	VMOV V3.D[0], R9
	VMOV V3.D[1], R10
	CBNZ R9, eulow8
	CBNZ R10, euhigh8
	ADD  $16, R2
	B    euloop16

eufirstLow8:
	RBIT R9, R11
	CLZ  R11, R11               // trailing zeros of R9 = first set bit
	LSR  $3, R11, R11           // /8 -> first matching byte (lane 0..7)
	MOVD R11, ret+24(FP)
	RET

eufirstHigh8:
	RBIT R10, R11
	CLZ  R11, R11
	LSR  $3, R11, R11           // lane within high half (0..7)
	ADD  $8, R11, R11           // lanes 8..15
	MOVD R11, ret+24(FP)
	RET

eulow8:
	RBIT R9, R11
	CLZ  R11, R11               // trailing zeros of R9 = first set bit
	LSR  $3, R11, R11           // /8 -> first matching byte (lane 0..7)
	ADD  R2, R11, R11
	MOVD R11, ret+24(FP)
	RET

euhigh8:
	RBIT R10, R11
	CLZ  R11, R11
	LSR  $3, R11, R11           // lane within high half (0..7)
	ADD  $8, R11, R11           // lanes 8..15
	ADD  R2, R11, R11
	MOVD R11, ret+24(FP)
	RET

eushortInput:
	MOVD ZR, R2                 // buffer shorter than one block: scan from 0

eutail:
	CMP  R1, R2
	BGE  eunotfound

eutailloop:
	// Sign-extending load (MOVB) + one signed compare covers control AND
	// non-ASCII bytes: as int8, 0x80..0xFF are negative and 0x00..0x1F are
	// below 0x20, while clean ASCII 0x20..0x7F is not — the same three
	// compares per byte as indexEscapeNEON's tail.
	ADD   R0, R2, R8
	MOVB  (R8), R9
	CMP   $0x20, R9
	BLT   eutfound             // control byte or non-ASCII byte (signed)
	CMP   $0x22, R9
	BEQ   eutfound
	CMP   $0x5c, R9
	BEQ   eutfound
	ADD   $1, R2
	CMP   R1, R2
	BLT   eutailloop

eunotfound:
	MOVD R1, ret+24(FP)
	RET

eutfound:
	MOVD R2, ret+24(FP)
	RET

// ---------------------------------------------------------------------------
// SVE2 scanners.
//
// The routines above are the NEON baseline, which every arm64 CPU can run. On a
// core that also implements SVE2 (Neoverse N2/V2, Graviton4, ...) the three hot
// string scanners take an SVE2 body instead, selected by the useSVE2 flag read
// at the top of each entry point below. Three properties of SVE make that body
// strictly shorter than the NEON one — this is instruction removal, which is
// what the N2 measurements say these scanners are bound by (IPC 3.49 with a
// 0.057% branch-miss rate: issue-bound, not mispredict-bound):
//
//  1. PREDICATION REMOVES THE TAIL. WHILELO builds a predicate covering only the
//     bytes actually left, and an SVE load never touches memory for an inactive
//     lane. One loop body therefore handles the short buffer, the full blocks and
//     the ragged final block alike: the NEON form's `CMP $16` head test, its
//     scalar byte tail and its separate short-input entry all disappear, along
//     with the branches that pick between them.
//
//  2. MATCH DOES A WHOLE CHARACTER CLASS IN ONE OP. SVE2's MATCH compares every
//     byte of the data against ALL SIXTEEN bytes of the corresponding 128-bit
//     segment of its second operand, so a set of up to 16 target bytes costs one
//     instruction: two VCMEQ plus a VORR collapse to one MATCH, and for the
//     structural set simdjson's whole shuffle trick (VAND, VTBL, VUSHR, VTBL,
//     VAND — five ops and a 48-byte RODATA load) collapses to the same one.
//
//  3. THE RESULT NEVER LEAVES THE VECTOR DOMAIN. MATCH — and CMPLO/CMPLT/ORRS —
//     set NZCV directly, so "did this block match" is a plain branch rather than
//     NEON's two VMOV cross-domain moves feeding two CBNZs. Position recovery is
//     BRKB (the lanes before the first true) plus INCP (add their count to the
//     offset): one cross-domain transfer instead of two, no RBIT/CLZ/LSR bit
//     arithmetic, and the two half-block recovery paths fold into one because
//     BRKB counts across the whole vector.
//
// The bodies are VECTOR-LENGTH AGNOSTIC: WHILELO sizes the predicate, INCB
// advances by however many bytes a vector holds, INCP counts lanes. The same
// code therefore scans 16 bytes per block on a 128-bit implementation (N2) and
// 32 or 64 on a wider one, with no change and no re-verification. MATCH's
// per-segment semantics are why each match set must hold the target bytes in
// EVERY 128-bit segment; TRN1-of-two-DUPs and LD1RQB (load-and-replicate-quad)
// both do that by construction, at any vector length.
//
// WHY THE FEATURE GATE IS HERE AND NOT IN GO. The natural spelling is an `if
// useSVE2` in the Go dispatch wrapper, but that gives the wrapper two calls,
// which costs 124 against the inliner's budget of 80 — indexCloseOrEscape stops
// inlining into ReadKey/ReadStringOrNull/SkipString/decodeEscaped and every
// generated decoder, the regression CLAUDE.md records as worth ~5% on cloudflare
// when it was fixed in the other direction. Reading the flag in assembly keeps
// the Go side a single unconditional call (cost 61, still inlined) and costs the
// SVE2 path three instructions: an ADRP+LDRB pair independent of the argument
// setup, and a perfectly-predicted not-taken CBZ. Non-SVE2 cores pay those three
// plus one taken branch into the unchanged NEON routine. Same shape as the
// amd64 side, which reads ·useAVX2(SB) inside simd_amd64.s for the same reason.
//
// The Go assembler has no SVE mnemonics (checked on Go 1.26), so the SVE
// instructions are emitted as WORD constants. Every encoding below was produced
// by GNU as 2.42 from the mnemonic in the trailing comment, and `objdump -d` of
// the built archive is what re-checks that they landed as intended; the
// differential tests (TestIndexVariantsFlip, TestIndexFunctionsMatchScalar) pin
// the behaviour. Control flow stays in ordinary Go asm: SVE's condition aliases
// are just NZCV, so ANY is NE and NONE is EQ, and BNE/BEQ read them directly.
//
// Async preemption cannot fire inside hand-written assembly — isAsyncSafePoint
// rejects any frame with abi.FuncFlagAsm set — so no goroutine switch can happen
// between these instructions. That is what makes predicate registers safe here:
// runtime.asyncPreempt saves V0-V31 but knows nothing of P0-P15, so a preemption
// mid-body would be the one way to lose them. The only interruption a leaf asm
// body can take is a signal, across which the kernel saves and restores the full
// SVE state, and no predicate value outlives such a context.

// The structural match set: one 128-bit quad holding '{', '}', '[' and ']' with
// the remaining lanes filled with '"'. LD1RQB replicates it into every segment.
// Bytes: 7b 7d 5b 5d 22 22 22 22 | 22 22 22 22 22 22 22 22
DATA sveStructSet<>+0(SB)/8, $0x222222225d5b7d7b
DATA sveStructSet<>+8(SB)/8, $0x2222222222222222
GLOBL sveStructSet<>(SB), RODATA|NOPTR, $16

// func indexQuoteOrBackslashArm64(b []byte) int
//
// The arm64 entry point for the '"'-or-'\\' scan: index of the first such byte,
// or len(b). Runs the SVE2 body below on an SVE2 core and tail-calls the NEON
// routine otherwise. The match set is built in-register as alternating '"'/'\\'
// bytes, so there is no table and no memory operand on the critical path.
TEXT ·indexQuoteOrBackslashArm64(SB), NOSPLIT, $0-32
	MOVBU ·useSVE2(SB), R9
	CBZ   R9, qbNEON
	MOVD  b_base+0(FP), R0
	MOVD  b_len+8(FP), R1
	WORD $0x2538c441 // mov z1.b, #34 // '"'
	WORD $0x2538cb82 // mov z2.b, #92 // '\\'
	WORD $0x05227021 // trn1 z1.b, z1.b, z2.b // -> 22 5c 22 5c ... per segment
	MOVD  ZR, R2

qbLoop:
	WORD $0x25211c41 // whilelo p1.b, x2, x1 // active lanes = min(VL, len-x2)
	BEQ  qbNone      // b.none: the offset reached the end
	WORD $0xa4024400 // ld1b {z0.b}, p1/z, [x0, x2]
	WORD $0x45218402 // match p2.b, p1/z, z0.b, z1.b
	BNE  qbFound     // b.any
	WORD $0x0430e3e2 // incb x2
	B    qbLoop

qbFound:
	WORD $0x25904443 // brkb p3.b, p1/z, p2.b // lanes before the first match
	WORD $0x252c8862 // incp x2, p3.b // x2 += how many there were
	MOVD R2, ret+24(FP)
	RET

qbNone:
	MOVD R1, ret+24(FP)
	RET

qbNEON:
	B ·indexQuoteOrBackslashNEON(SB) // args are already in place at FP

// func indexEscapeArm64(b []byte) int
//
// The arm64 entry point for the escape scan: first '"', '\\' or control byte
// < 0x20, or len(b). The control-byte half is one unsigned CMPLO against #32,
// where NEON needs a VUMIN and a VCMEQ.
TEXT ·indexEscapeArm64(SB), NOSPLIT, $0-32
	MOVBU ·useSVE2(SB), R9
	CBZ   R9, esNEON
	MOVD  b_base+0(FP), R0
	MOVD  b_len+8(FP), R1
	WORD $0x2538c441 // mov z1.b, #34
	WORD $0x2538cb82 // mov z2.b, #92
	WORD $0x05227021 // trn1 z1.b, z1.b, z2.b
	MOVD  ZR, R2

esLoop:
	WORD $0x25211c41 // whilelo p1.b, x2, x1
	BEQ  esNone
	WORD $0xa4024400 // ld1b {z0.b}, p1/z, [x0, x2]
	WORD $0x45218402 // match p2.b, p1/z, z0.b, z1.b
	WORD $0x24282403 // cmplo p3.b, p1/z, z0.b, #32 // unsigned <, i.e. <= 0x1f
	WORD $0x25c34442 // orrs p2.b, p1/z, p2.b, p3.b
	BNE  esFound
	WORD $0x0430e3e2 // incb x2
	B    esLoop

esFound:
	WORD $0x25904443 // brkb p3.b, p1/z, p2.b
	WORD $0x252c8862 // incp x2, p3.b
	MOVD R2, ret+24(FP)
	RET

esNone:
	MOVD R1, ret+24(FP)
	RET

esNEON:
	B ·indexEscapeNEON(SB)

// func indexEscapeNonASCIIArm64(b []byte) int
//
// The arm64 entry point for the widened escape scan: indexEscapeArm64's
// predicate plus non-ASCII bytes.
//
// The two range halves are folded into ONE compare. What the widened predicate
// wants is "the byte is not in [0x20, 0x7f]" — a control byte below that range
// or a non-ASCII byte above it — and subtracting the low end turns the interval
// into a single unsigned test: (c - 0x20) >= 0x60 holds for exactly 0x00-0x1f
// (which wrap to 0xe0-0xff) and 0x80-0xff (which land on 0x60-0xdf), and fails
// for the printable ASCII between them. So a SUB and a CMPHS replace the
// separate CMPLO/CMPLT pair and the ORR that joined them.
//
// That is worth more than the one instruction it saves, because the two it
// removes are PREDICATE ops while the one it adds is an ordinary vector op, and
// predicate issue is this loop's limit on N2. The first version of this routine
// — MATCH, CMPLO, ORR, CMPLT, ORRS, five predicate ops per block — was the one
// scanner of the four that LOST to NEON on long clean runs (EscapeString
// log_line_clean +4.4%, mostly_clean_one_quote +5.3%) despite executing fewer
// instructions in total: NEON spreads its eight vector ops over two vector
// pipes, where predicate ops serialize on one. At three predicate ops the long
// runs come back. SUB is unpredicated, so it also computes on the lanes the
// zeroing load left at zero; harmless, because CMPHS is predicated and those
// lanes cannot reach the result.
TEXT ·indexEscapeNonASCIIArm64(SB), NOSPLIT, $0-32
	MOVBU ·useSVE2(SB), R9
	CBZ   R9, euNEON
	MOVD  b_base+0(FP), R0
	MOVD  b_len+8(FP), R1
	WORD $0x2538c441 // mov z1.b, #34
	WORD $0x2538cb82 // mov z2.b, #92
	WORD $0x05227021 // trn1 z1.b, z1.b, z2.b
	WORD $0x2538c403 // mov z3.b, #32 // the interval's low end
	MOVD  ZR, R2

euLoop:
	WORD $0x25211c41 // whilelo p1.b, x2, x1
	BEQ  euNone
	WORD $0xa4024400 // ld1b {z0.b}, p1/z, [x0, x2]
	WORD $0x45218402 // match p2.b, p1/z, z0.b, z1.b
	WORD $0x04230404 // sub z4.b, z0.b, z3.b
	WORD $0x24380483 // cmphs p3.b, p1/z, z4.b, #96 // outside [0x20, 0x7f]
	WORD $0x25c34442 // orrs p2.b, p1/z, p2.b, p3.b
	BNE  euFound
	WORD $0x0430e3e2 // incb x2
	B    euLoop

euFound:
	WORD $0x25904443 // brkb p3.b, p1/z, p2.b
	WORD $0x252c8862 // incp x2, p3.b
	MOVD R2, ret+24(FP)
	RET

euNone:
	MOVD R1, ret+24(FP)
	RET

euNEON:
	B ·indexEscapeNonASCIINEON(SB)

// func indexStructuralSVE2(b []byte) int
//
// SVE2 twin of indexStructuralNEON: first '{', '}', '[', ']' or '"', or len(b).
// All five bytes live in one MATCH. Unlike the three scanners above this one is
// reached through a Go-level `if useSVE2` rather than an in-assembly gate: its
// dispatcher (indexStructural) already carries a length test and a 16-byte
// scalar prescan, so it is far past the inline budget either way and the branch
// costs nothing there.
TEXT ·indexStructuralSVE2(SB), NOSPLIT, $0-32
	MOVD b_base+0(FP), R0
	MOVD b_len+8(FP), R1
	MOVD $sveStructSet<>(SB), R3
	WORD $0x2518e3e0 // ptrue p0.b
	WORD $0xa4002061 // ld1rqb {z1.b}, p0/z, [x3] // 16 bytes -> every segment
	MOVD ZR, R2

stLoop:
	WORD $0x25211c41 // whilelo p1.b, x2, x1
	BEQ  stNone
	WORD $0xa4024400 // ld1b {z0.b}, p1/z, [x0, x2]
	WORD $0x45218402 // match p2.b, p1/z, z0.b, z1.b
	BNE  stFound
	WORD $0x0430e3e2 // incb x2
	B    stLoop

stFound:
	WORD $0x25904443 // brkb p3.b, p1/z, p2.b
	WORD $0x252c8862 // incp x2, p3.b
	MOVD R2, ret+24(FP)
	RET

stNone:
	MOVD R1, ret+24(FP)
	RET
