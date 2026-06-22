#include "textflag.h"

// bitWeights<> is the per-byte bit-position weight {1,2,4,...,128} repeated over
// both 8-lane halves, used to turn a 0x00/0xFF compare mask into a packed bitmask.
DATA bitWeights<>+0(SB)/8, $0x8040201008040201
DATA bitWeights<>+8(SB)/8, $0x8040201008040201
GLOBL bitWeights<>(SB), RODATA|NOPTR, $16

// FOLD16 computes the 16-bit mask of (chunk c == splat s) into GP register dst,
// via weight-and-fold: AND the 0x00/0xFF compare with the {1,2,4,…,128} bit
// weights, then three pairwise adds collapse each 8-lane half to one byte whose
// bits are that half's mask (byte 0 = lanes 0–7, byte 1 = lanes 8–15); one
// halfword VMOV lifts both bytes to dst.
#define FOLD16(s, c, dst) \
	VCMEQ s, c, V20.B16          \
	VAND  V16.B16, V20.B16, V20.B16 \
	VADDP V20.B16, V20.B16, V20.B16 \
	VADDP V20.B16, V20.B16, V20.B16 \
	VADDP V20.B16, V20.B16, V20.B16 \
	VMOV  V20.H[0], dst

// CLASS folds splat s against all four 16-byte chunks (V6..V9) and assembles the
// 64-bit class mask in dst: chunk k's 16 bits at bit offset 16k.
#define CLASS(s, dst) \
	FOLD16(s, V6.B16, dst) \
	FOLD16(s, V7.B16, R2)  \
	ORR   R2<<16, dst, dst \
	FOLD16(s, V8.B16, R2)  \
	ORR   R2<<32, dst, dst \
	FOLD16(s, V9.B16, R2)  \
	ORR   R2<<48, dst, dst

// func maskBlock(b []byte, isArray bool) (quote, bslash, open, close uint64)
//
// Returns, for the 64 bytes at b[:64], the bitmap of `"`, `\\`, and the
// container's open/close brackets — `[`/`]` when isArray, else `{`/`}`; bit j set
// when b[j] is that byte. NEON; the caller guarantees 64 readable bytes. The
// 64-byte block is four 16-byte chunks (V6..V9); each class folds to a 16-bit
// mask per chunk that is shifted to bits 16k of the uint64.
TEXT ·maskBlock(SB), NOSPLIT, $0-64
	MOVD b_base+0(FP), R0
	MOVD $0x22, R1
	VDUP R1, V0.B16            // '"'
	MOVD $0x5c, R1
	VDUP R1, V1.B16            // '\\'
	// Select the container's bracket splats: '['/']' for an array, else '{'/'}'.
	MOVBU isArray+24(FP), R1
	CBNZ  R1, arrayBrackets
	MOVD  $0x7b, R1
	VDUP  R1, V2.B16           // '{'
	MOVD  $0x7d, R1
	VDUP  R1, V3.B16           // '}'
	JMP   haveBrackets
arrayBrackets:
	MOVD $0x5b, R1
	VDUP R1, V2.B16            // '['
	MOVD $0x5d, R1
	VDUP R1, V3.B16            // ']'
haveBrackets:
	MOVD $bitWeights<>(SB), R1
	VLD1 (R1), [V16.B16]

	VLD1 (R0), [V6.B16, V7.B16, V8.B16, V9.B16] // four 16-byte chunks

	CLASS(V0.B16, R10)
	CLASS(V1.B16, R11)
	CLASS(V2.B16, R12)
	CLASS(V3.B16, R13)

	MOVD R10, quote+32(FP)
	MOVD R11, bslash+40(FP)
	MOVD R12, open+48(FP)
	MOVD R13, close+56(FP)
	RET
