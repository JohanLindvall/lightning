#include "textflag.h"

// bitWeights<> is the per-byte bit-position weight {1,2,...,128} repeated over
// both 8-lane halves, used to turn a 0x00/0xFF compare mask into a packed bitmask.
DATA bitWeights<>+0(SB)/8, $0x8040201008040201
DATA bitWeights<>+8(SB)/8, $0x8040201008040201
GLOBL bitWeights<>(SB), RODATA|NOPTR, $16

// CLASS folds splat s against all four 16-byte chunks (V6..V9) into the 64-bit
// class mask in dst — chunk k's 16 bits at bit offset 16k — with a single
// vector→GP move.
//
// Each chunk's compare (lanes 0x00/0xFF) is ANDed with the {1,2,...,128} bit
// weights, so a matching lane becomes its bit value. A four-step ADDP cascade then
// collapses each chunk's two 8-lane halves to one byte (the half's 8-bit mask) and
// packs all four chunks into the low eight bytes in order
// [lo0,hi0,lo1,hi1,lo2,hi2,lo3,hi3] — which is exactly the 64-bit mask, chunk k at
// bit 16k — so one VMOV of the D lane lifts the whole result. (ADDP(Vn,Vm) puts the
// pairwise sums of Vn in the low half, Vm in the high half; the cascade P=ADDP(A0,
// A1), Q=ADDP(A2,A3), R=ADDP(P,Q), S=ADDP(R,R) keeps chunk order and never
// overflows a byte — a full 8-lane half sums to 255.) This replaces a per-chunk
// gather that did four cross-domain VMOVs and twelve ADDPs plus shift/OR stitching.
#define CLASS(s, dst) \
	VCMEQ s, V6.B16, V20.B16        \
	VAND  V16.B16, V20.B16, V20.B16 \
	VCMEQ s, V7.B16, V21.B16        \
	VAND  V16.B16, V21.B16, V21.B16 \
	VCMEQ s, V8.B16, V22.B16        \
	VAND  V16.B16, V22.B16, V22.B16 \
	VCMEQ s, V9.B16, V23.B16        \
	VAND  V16.B16, V23.B16, V23.B16 \
	VADDP V21.B16, V20.B16, V20.B16 \
	VADDP V23.B16, V22.B16, V22.B16 \
	VADDP V22.B16, V20.B16, V20.B16 \
	VADDP V20.B16, V20.B16, V20.B16 \
	VMOV  V20.D[0], dst

// func maskBlock(b []byte, isArray bool) (quote, bslash, open, close uint64)
//
// Returns, for the 64 bytes at b[:64], the bitmap of `"`, `\\`, and the
// container's open/close brackets — `[`/`]` when isArray, else `{`/`}`; bit j set
// when b[j] is that byte. NEON; the caller guarantees 64 readable bytes. The
// 64-byte block is four 16-byte chunks (V6..V9); each class folds to its 64-bit
// mask via the CLASS cascade above.
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
