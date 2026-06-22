#include "textflag.h"

// bitWeights<> is the per-byte bit-position weight {1,2,4,...,128} repeated over
// both 8-lane halves, used to turn a 0x00/0xFF compare mask into a packed bitmask.
DATA bitWeights<>+0(SB)/8, $0x8040201008040201
DATA bitWeights<>+8(SB)/8, $0x8040201008040201
GLOBL bitWeights<>(SB), RODATA|NOPTR, $16

// func maskBlock(b []byte) (quote, bslash, lbrace, rbrace, lbrack, rbrack uint32)
//
// Returns, for the 32 bytes at b[:32], the bitmap of each of '"' '\\' '{' '}'
// '[' ']' (bit j set when b[j] is that byte). NEON; the caller guarantees 32
// readable bytes.
//
// NEON has no PMOVMSKB, so each 16-byte VCMEQ result (lanes 0x00/0xFF) is turned
// into a 16-bit mask by weight-and-fold: AND with the {1,2,4,…,128} repeated
// bit-weight vector (each matching lane becomes its bit value), then three
// pairwise adds (VADDP) collapse each 8-lane half to a single byte whose bits are
// exactly that half's mask — byte 0 is lanes 0–7, byte 1 is lanes 8–15 — so one
// halfword VMOV lifts the full 16-bit mask to a GP register. This replaces the
// older per-half "AND/MUL/LSR" gather: it halves the cross-domain VMOVs (one per
// 16-byte half instead of two) and drops the integer multiplies entirely, doing
// the fold with cheap NEON pairwise adds instead. The block's two 16-byte halves
// give the low and high 16 bits of each uint32.
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
	MOVD $bitWeights<>(SB), R2
	VLD1 (R2), [V16.B16]      // bit-position weights

	// First 16 bytes -> low 16 bits of each result (R10..R15).
	VLD1 (R0), [V6.B16]

	VCMEQ V0.B16, V6.B16, V7.B16
	VAND  V16.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VMOV  V7.H[0], R10

	VCMEQ V1.B16, V6.B16, V7.B16
	VAND  V16.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VMOV  V7.H[0], R11

	VCMEQ V2.B16, V6.B16, V7.B16
	VAND  V16.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VMOV  V7.H[0], R12

	VCMEQ V3.B16, V6.B16, V7.B16
	VAND  V16.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VMOV  V7.H[0], R13

	VCMEQ V4.B16, V6.B16, V7.B16
	VAND  V16.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VMOV  V7.H[0], R14

	VCMEQ V5.B16, V6.B16, V7.B16
	VAND  V16.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VMOV  V7.H[0], R15

	// Second 16 bytes -> high 16 bits (OR in shifted by 16).
	ADD  $16, R0, R4
	VLD1 (R4), [V6.B16]

	VCMEQ V0.B16, V6.B16, V7.B16
	VAND  V16.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VMOV  V7.H[0], R5
	ORR   R5<<16, R10, R10

	VCMEQ V1.B16, V6.B16, V7.B16
	VAND  V16.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VMOV  V7.H[0], R5
	ORR   R5<<16, R11, R11

	VCMEQ V2.B16, V6.B16, V7.B16
	VAND  V16.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VMOV  V7.H[0], R5
	ORR   R5<<16, R12, R12

	VCMEQ V3.B16, V6.B16, V7.B16
	VAND  V16.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VMOV  V7.H[0], R5
	ORR   R5<<16, R13, R13

	VCMEQ V4.B16, V6.B16, V7.B16
	VAND  V16.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VMOV  V7.H[0], R5
	ORR   R5<<16, R14, R14

	VCMEQ V5.B16, V6.B16, V7.B16
	VAND  V16.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VADDP V7.B16, V7.B16, V7.B16
	VMOV  V7.H[0], R5
	ORR   R5<<16, R15, R15

	MOVW R10, quote+24(FP)
	MOVW R11, bslash+28(FP)
	MOVW R12, lbrace+32(FP)
	MOVW R13, rbrace+36(FP)
	MOVW R14, lbrack+40(FP)
	MOVW R15, rbrack+44(FP)
	RET
