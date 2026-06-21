#include "textflag.h"

// No 16-byte index constant is needed for position recovery: scalar RBIT/CLZ on
// the result mask moved into general registers finds the first nonzero byte.
//
// indexQuoteOrBackslashNEON builds its two comparison splats ('"', '\\') per call
// with MOVD-immediate + VDUP rather than loading a RODATA table: on amd64 the
// equivalent constants live in RODATA because building them inline needs
// VPBROADCASTB (a GP→XMM domain crossing that dominates a single-block call), but
// arm64's VDUP-from-GPR has no such penalty, so the inline form is cheaper than a
// VLD1 (no load-use latency) and measures ~2% faster on the get benchmark.
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
	MOVD b_base+0(FP), R0
	MOVD b_len+8(FP), R1
	MOVD $0, R2                  // R2 = current offset
	MOVD $0x22, R3
	VDUP R3, V0.B16              // V0 = '"' x16
	MOVD $0x5c, R3
	VDUP R3, V1.B16              // V1 = '\\' x16

loop16:
	SUB  R2, R1, R7              // R7 = remaining = len - offset
	CMP  $16, R7
	BLT  tail
	ADD  R0, R2, R8              // R8 = &b[offset]
	VLD1 (R8), [V2.B16]
	VCMEQ V0.B16, V2.B16, V3.B16 // V3 = (chunk == '"')
	VCMEQ V1.B16, V2.B16, V4.B16 // V4 = (chunk == '\\')
	VORR V3.B16, V4.B16, V3.B16  // V3 = either match (0xFF per lane)
	VMOV V3.D[0], R9             // low 8 lanes
	VMOV V3.D[1], R10            // high 8 lanes
	CBNZ R9, low8
	CBNZ R10, high8
	ADD  $16, R2
	B    loop16

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

// func indexStructuralNEON(b []byte) int
//
// Returns the index of the first '{', '}', '[', ']' or '"' byte, or len(b).
TEXT ·indexStructuralNEON(SB), NOSPLIT, $0-32
	MOVD b_base+0(FP), R0
	MOVD b_len+8(FP), R1
	MOVD $0, R2
	MOVD $structTablesArm<>(SB), R3
	VLD1 (R3), [V0.B16, V1.B16, V2.B16] // structLo, structHi, 0x0f mask

sloop:
	SUB  R2, R1, R7
	CMP  $16, R7
	BLT  stail
	ADD  R0, R2, R8
	VLD1  (R8), [V5.B16]              // chunk
	VAND  V2.B16, V5.B16, V6.B16     // low nibbles
	VTBL  V6.B16, [V0.B16], V6.B16   // structLo[lowNibble]
	VUSHR $4, V5.B16, V7.B16         // high nibbles (per-byte shift)
	VTBL  V7.B16, [V1.B16], V7.B16   // structHi[highNibble]
	VAND  V7.B16, V6.B16, V6.B16     // nonzero byte where structural
	VMOV V6.D[0], R9
	VMOV V6.D[1], R10
	CBNZ R9, slow8
	CBNZ R10, shigh8
	ADD  $16, R2
	B    sloop

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
