#include "textflag.h"

// 16-byte constant of byte indices is not needed; we recover the position with
// scalar RBIT/CLZ on the comparison mask moved into general registers.

// Precomputed 16-byte splats of the search bytes, laid out contiguously so a
// single multi-register VLD1 fills all the comparison vectors at once. This
// hoists the per-call constant setup (a MOVD + VDUP for each byte) out of the
// hot functions: instead of rebuilding the splats on every call — wasteful when
// the routine returns after a single 16-byte block, the common case for
// token-dense JSON — they are loaded once from RODATA.
//
// qbMask holds '"' then '\\' (indexQuoteOrBackslashNEON, 32 bytes).
DATA qbMask<>+0(SB)/8, $0x2222222222222222
DATA qbMask<>+8(SB)/8, $0x2222222222222222
DATA qbMask<>+16(SB)/8, $0x5c5c5c5c5c5c5c5c
DATA qbMask<>+24(SB)/8, $0x5c5c5c5c5c5c5c5c
GLOBL qbMask<>(SB), RODATA|NOPTR, $32

// structMask holds '{', '}', '[', ']' then '"' (indexStructuralNEON, 80 bytes).
DATA structMask<>+0(SB)/8, $0x7b7b7b7b7b7b7b7b
DATA structMask<>+8(SB)/8, $0x7b7b7b7b7b7b7b7b
DATA structMask<>+16(SB)/8, $0x7d7d7d7d7d7d7d7d
DATA structMask<>+24(SB)/8, $0x7d7d7d7d7d7d7d7d
DATA structMask<>+32(SB)/8, $0x5b5b5b5b5b5b5b5b
DATA structMask<>+40(SB)/8, $0x5b5b5b5b5b5b5b5b
DATA structMask<>+48(SB)/8, $0x5d5d5d5d5d5d5d5d
DATA structMask<>+56(SB)/8, $0x5d5d5d5d5d5d5d5d
DATA structMask<>+64(SB)/8, $0x2222222222222222
DATA structMask<>+72(SB)/8, $0x2222222222222222
GLOBL structMask<>(SB), RODATA|NOPTR, $80

// func indexQuoteOrBackslashNEON(b []byte) int
//
// Returns the index of the first '"' or '\\' byte in b, or len(b) if neither
// is present. Scans 16 bytes per iteration with NEON, then a scalar tail.
TEXT ·indexQuoteOrBackslashNEON(SB), NOSPLIT, $0-32
	MOVD b_base+0(FP), R0
	MOVD b_len+8(FP), R1
	MOVD $0, R2                  // R2 = current offset
	MOVD $qbMask<>(SB), R3
	VLD1 (R3), [V0.B16, V1.B16]  // V0 = '"' x16, V1 = '\\' x16

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
	MOVD $structMask<>(SB), R3
	VLD1 (R3), [V0.B16, V1.B16, V2.B16, V3.B16] // '{', '}', '[', ']'
	ADD  $64, R3, R3
	VLD1 (R3), [V4.B16]                          // '"'

sloop:
	SUB  R2, R1, R7
	CMP  $16, R7
	BLT  stail
	ADD  R0, R2, R8
	VLD1 (R8), [V5.B16]          // chunk
	VCMEQ V0.B16, V5.B16, V6.B16
	VCMEQ V1.B16, V5.B16, V7.B16
	VORR V7.B16, V6.B16, V6.B16
	VCMEQ V2.B16, V5.B16, V7.B16
	VORR V7.B16, V6.B16, V6.B16
	VCMEQ V3.B16, V5.B16, V7.B16
	VORR V7.B16, V6.B16, V6.B16
	VCMEQ V4.B16, V5.B16, V7.B16
	VORR V7.B16, V6.B16, V6.B16
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
