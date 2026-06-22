#include "textflag.h"

// func maskBlock(b []byte) (quote, bslash, lbrace, rbrace, lbrack, rbrack uint32)
//
// Returns, for the 32 bytes at b[:32], the bitmap of each of '"' '\\' '{' '}'
// '[' ']' (bit j set when b[j] is that byte). NEON; the caller guarantees 32
// readable bytes.
//
// NEON has no PMOVMSKB, so each 16-byte VCMEQ result (lanes 0x00/0xFF) is turned
// into a 16-bit mask by the byte-LSB gather: AND each 8-lane half (moved to a GP
// register) with 0x0101010101010101, multiply by 0x0102040810204080, and shift
// right 56 — bit j of the byte lands from lane j, matching amd64 PMOVMSKB. The
// two 16-byte halves of the block give the low and high 16 bits of each uint32.
TEXT ·maskBlock(SB), NOSPLIT, $0-48
	MOVD b_base+0(FP), R0
	MOVD $0x22, R1
	VDUP R1, V0.B16            // '"'
	MOVD $0x5c, R1
	VDUP R1, V1.B16            // '\\'
	MOVD $0x7b, R1
	VDUP R1, V2.B16            // '{'
	MOVD $0x7d, R1
	VDUP R1, V3.B16            // '}'
	MOVD $0x5b, R1
	VDUP R1, V4.B16            // '['
	MOVD $0x5d, R1
	VDUP R1, V5.B16            // ']'
	MOVD $0x0101010101010101, R2 // LSB-per-byte mask
	MOVD $0x0102040810204080, R3 // gather magic

	// First 16 bytes -> low 16 bits of each result (R10..R15).
	VLD1 (R0), [V6.B16]

	VCMEQ V0.B16, V6.B16, V7.B16
	VMOV  V7.D[0], R5
	VMOV  V7.D[1], R6
	AND   R2, R5, R5
	MUL   R3, R5, R5
	LSR   $56, R5, R5
	AND   R2, R6, R6
	MUL   R3, R6, R6
	LSR   $56, R6, R6
	LSL   $8, R6, R6
	ORR   R6, R5, R10

	VCMEQ V1.B16, V6.B16, V7.B16
	VMOV  V7.D[0], R5
	VMOV  V7.D[1], R6
	AND   R2, R5, R5
	MUL   R3, R5, R5
	LSR   $56, R5, R5
	AND   R2, R6, R6
	MUL   R3, R6, R6
	LSR   $56, R6, R6
	LSL   $8, R6, R6
	ORR   R6, R5, R11

	VCMEQ V2.B16, V6.B16, V7.B16
	VMOV  V7.D[0], R5
	VMOV  V7.D[1], R6
	AND   R2, R5, R5
	MUL   R3, R5, R5
	LSR   $56, R5, R5
	AND   R2, R6, R6
	MUL   R3, R6, R6
	LSR   $56, R6, R6
	LSL   $8, R6, R6
	ORR   R6, R5, R12

	VCMEQ V3.B16, V6.B16, V7.B16
	VMOV  V7.D[0], R5
	VMOV  V7.D[1], R6
	AND   R2, R5, R5
	MUL   R3, R5, R5
	LSR   $56, R5, R5
	AND   R2, R6, R6
	MUL   R3, R6, R6
	LSR   $56, R6, R6
	LSL   $8, R6, R6
	ORR   R6, R5, R13

	VCMEQ V4.B16, V6.B16, V7.B16
	VMOV  V7.D[0], R5
	VMOV  V7.D[1], R6
	AND   R2, R5, R5
	MUL   R3, R5, R5
	LSR   $56, R5, R5
	AND   R2, R6, R6
	MUL   R3, R6, R6
	LSR   $56, R6, R6
	LSL   $8, R6, R6
	ORR   R6, R5, R14

	VCMEQ V5.B16, V6.B16, V7.B16
	VMOV  V7.D[0], R5
	VMOV  V7.D[1], R6
	AND   R2, R5, R5
	MUL   R3, R5, R5
	LSR   $56, R5, R5
	AND   R2, R6, R6
	MUL   R3, R6, R6
	LSR   $56, R6, R6
	LSL   $8, R6, R6
	ORR   R6, R5, R15

	// Second 16 bytes -> high 16 bits (OR in shifted by 16).
	ADD  $16, R0, R4
	VLD1 (R4), [V6.B16]

	VCMEQ V0.B16, V6.B16, V7.B16
	VMOV  V7.D[0], R5
	VMOV  V7.D[1], R6
	AND   R2, R5, R5
	MUL   R3, R5, R5
	LSR   $56, R5, R5
	AND   R2, R6, R6
	MUL   R3, R6, R6
	LSR   $56, R6, R6
	LSL   $8, R6, R6
	ORR   R6, R5, R5
	LSL   $16, R5, R5
	ORR   R5, R10, R10

	VCMEQ V1.B16, V6.B16, V7.B16
	VMOV  V7.D[0], R5
	VMOV  V7.D[1], R6
	AND   R2, R5, R5
	MUL   R3, R5, R5
	LSR   $56, R5, R5
	AND   R2, R6, R6
	MUL   R3, R6, R6
	LSR   $56, R6, R6
	LSL   $8, R6, R6
	ORR   R6, R5, R5
	LSL   $16, R5, R5
	ORR   R5, R11, R11

	VCMEQ V2.B16, V6.B16, V7.B16
	VMOV  V7.D[0], R5
	VMOV  V7.D[1], R6
	AND   R2, R5, R5
	MUL   R3, R5, R5
	LSR   $56, R5, R5
	AND   R2, R6, R6
	MUL   R3, R6, R6
	LSR   $56, R6, R6
	LSL   $8, R6, R6
	ORR   R6, R5, R5
	LSL   $16, R5, R5
	ORR   R5, R12, R12

	VCMEQ V3.B16, V6.B16, V7.B16
	VMOV  V7.D[0], R5
	VMOV  V7.D[1], R6
	AND   R2, R5, R5
	MUL   R3, R5, R5
	LSR   $56, R5, R5
	AND   R2, R6, R6
	MUL   R3, R6, R6
	LSR   $56, R6, R6
	LSL   $8, R6, R6
	ORR   R6, R5, R5
	LSL   $16, R5, R5
	ORR   R5, R13, R13

	VCMEQ V4.B16, V6.B16, V7.B16
	VMOV  V7.D[0], R5
	VMOV  V7.D[1], R6
	AND   R2, R5, R5
	MUL   R3, R5, R5
	LSR   $56, R5, R5
	AND   R2, R6, R6
	MUL   R3, R6, R6
	LSR   $56, R6, R6
	LSL   $8, R6, R6
	ORR   R6, R5, R5
	LSL   $16, R5, R5
	ORR   R5, R14, R14

	VCMEQ V5.B16, V6.B16, V7.B16
	VMOV  V7.D[0], R5
	VMOV  V7.D[1], R6
	AND   R2, R5, R5
	MUL   R3, R5, R5
	LSR   $56, R5, R5
	AND   R2, R6, R6
	MUL   R3, R6, R6
	LSR   $56, R6, R6
	LSL   $8, R6, R6
	ORR   R6, R5, R5
	LSL   $16, R5, R5
	ORR   R5, R15, R15

	MOVW R10, quote+24(FP)
	MOVW R11, bslash+28(FP)
	MOVW R12, lbrace+32(FP)
	MOVW R13, rbrace+36(FP)
	MOVW R14, lbrack+40(FP)
	MOVW R15, rbrack+44(FP)
	RET
