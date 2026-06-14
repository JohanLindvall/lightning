#include "textflag.h"

// 16-byte constant of byte indices is not needed; we recover the position with
// scalar RBIT/CLZ on the comparison mask moved into general registers.
//
// The comparison splats are built per call with MOVD-immediate + VDUP rather
// than loaded from a RODATA table. On amd64 the equivalent constants are kept in
// RODATA because building them inline needs VPBROADCASTB, whose GP→XMM domain
// crossing dominates when the routine early-exits after a single block; arm64's
// VDUP-from-GPR has no such penalty, so the inline form is cheaper than a VLD1
// load (no load-use latency, no memory dependency) and measures ~2% faster on
// the get benchmark. See index_amd64.s for the contrasting choice.

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
	MOVD $0x7b, R3
	VDUP R3, V0.B16   // '{'
	MOVD $0x7d, R3
	VDUP R3, V1.B16   // '}'
	MOVD $0x5b, R3
	VDUP R3, V2.B16   // '['
	MOVD $0x5d, R3
	VDUP R3, V3.B16   // ']'
	MOVD $0x22, R3
	VDUP R3, V4.B16   // '"'

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
