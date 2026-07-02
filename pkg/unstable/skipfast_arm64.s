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

// POPCNT64 counts the set bits of GP register rs into rd through the vector
// unit (arm64 has no GP popcount) — the same FMOV/CNT/UADDLV sequence the
// compiler emits for bits.OnesCount64. Clobbers V24.
#define POPCNT64(rs, rd) \
	FMOVD   rs, F24          \
	VCNT    V24.B8, V24.B8   \
	VUADDLV V24.B8, V24      \
	FMOVD   F24, rd

// func skipBlocksNEON(data []byte, pos, depth int, isArray bool) (end, ndepth int, prevEscaped, prevInString uint64)
//
// The arm64 whole-loop form of skipContainerFast's block scan (see the amd64
// twin in skipfast_amd64.s): splats and the bit-weight vector loaded once,
// depth/prevEscaped/prevInString carried in registers, and the same fast paths
// as the Go loop — skip the escape math when no backslash is in play, skip the
// prefix XOR when the block has no unescaped quote, and update depth with
// popcounts in bulk whenever the block cannot cross depth 0. Per block it does
// exactly the work maskBlock + the Go bit math did, minus the per-block call,
// four results through memory, five splat VDUPs and the isArray branch.
//
// Unlike the amd64 twin there is no carryless-multiply prefix XOR: the mask is
// in the GP domain (the escape math needs GP add-with-carry) and a
// GP->SIMD->GP round trip for PMULL costs more on M-class cores than the
// shift-XOR chain — which arm64's shifted-register EOR does in six single
// instructions anyway.
//
// Register map: R0=data base, R1=len-64 (last valid block start), R2=pos,
// R16=data cursor (R0+R2), R3=depth, R4=prevEscaped (0/1), R5=prevInString
// (0/all-ones), R6=evenBits 0x5555...; R10-R13 the block's quote/bslash/open/
// close masks (CLASS outputs), R7-R9/R14/R15 scratch. V0-V3 splats, V16 bit
// weights, V6-V9 chunks, V20-V24 scratch.
TEXT ·skipBlocksNEON(SB), NOSPLIT, $0-80
	MOVD data_base+0(FP), R0
	MOVD data_len+8(FP), R1
	SUB  $64, R1
	MOVD pos+24(FP), R2
	MOVD depth+32(FP), R3
	MOVD ZR, R4
	MOVD ZR, R5
	MOVD $0x5555555555555555, R6
	MOVD $0x22, R7
	VDUP R7, V0.B16            // '"'
	MOVD $0x5c, R7
	VDUP R7, V1.B16            // '\\'
	MOVBU isArray+40(FP), R7
	CBNZ R7, arrayBrackets
	MOVD $0x7b, R7
	VDUP R7, V2.B16            // '{'
	MOVD $0x7d, R7
	VDUP R7, V3.B16            // '}'
	JMP  haveBrackets
arrayBrackets:
	MOVD $0x5b, R7
	VDUP R7, V2.B16            // '['
	MOVD $0x5d, R7
	VDUP R7, V3.B16            // ']'
haveBrackets:
	MOVD $bitWeights<>(SB), R7
	VLD1 (R7), [V16.B16]
	ADD  R0, R2, R16           // cursor = &data[pos]

blockLoop:
	CMP  R1, R2
	BGT  exhausted
	VLD1 (R16), [V6.B16, V7.B16, V8.B16, V9.B16]
	CLASS(V0.B16, R10)         // quote
	CLASS(V1.B16, R11)         // bslash
	CLASS(V2.B16, R12)         // open
	CLASS(V3.B16, R13)         // close

	// escaped (R15): skip the add-carry chain when no backslash in the block
	// and none pending from the previous one.
	ORR  R4, R11, R7
	CBNZ R7, computeEscaped
	MOVD ZR, R15
	JMP  haveEscaped
computeEscaped:
	BIC  R4, R11, R11          // backslash &^= prevEscaped
	LSL  $1, R11, R7
	ORR  R4, R7, R7            // followsEscape = backslash<<1 | prevEscaped
	BIC  R6, R11, R8           // backslash &^ evenBits
	BIC  R7, R8, R8            //   ... &^ followsEscape = oddSequenceStarts
	ADDS R11, R8, R8           // seq = odd + backslash, C = carry out
	ADC  ZR, ZR, R4            // prevEscaped = carry
	LSL  $1, R8, R8            // invertMask = seq << 1
	EOR  R6, R8, R8
	AND  R7, R8, R15           // escaped = (evenBits ^ invertMask) & follows
haveEscaped:
	// inStr (R14): a block with no unescaped quote cannot change the
	// in-string state; otherwise the prefix XOR runs as six shifted-register
	// EORs (x ^= x<<1; x<<2; ... x<<32).
	BIC  R15, R10, R7          // q = quote &^ escaped
	CBNZ R7, computeInStr
	MOVD R5, R14
	JMP  haveInStr
computeInStr:
	EOR  R7<<1, R7, R7
	EOR  R7<<2, R7, R7
	EOR  R7<<4, R7, R7
	EOR  R7<<8, R7, R7
	EOR  R7<<16, R7, R7
	EOR  R7<<32, R7, R7
	EOR  R5, R7, R14           // inStr = prefixXor(q) ^ prevInString
	ASR  $63, R14, R5          // carry: inside-string at byte 63
haveInStr:
	BIC  R14, R12, R12         // op &^= inStr
	BIC  R14, R13, R13         // cl &^= inStr
	ORR  R13, R12, R7          // brackets
	CBZ  R7, nextBlock
	CBZ  R13, onlyOpens
	POPCNT64(R13, R8)          // nc = popcount(cl)
	CMP  R8, R3
	BGT  bulk                  // depth > closes: no crossing possible
	// Per-bit walk in byte order: only blocks that might close the container.
bitLoop:
	CBZ  R7, nextBlock
	RBIT R7, R9
	CLZ  R9, R9                // j = index of the lowest set bit
	LSR  R9, R12, R8
	TBNZ $0, R8, isOpen        // open at j?
	SUB  $1, R3
	CBZ  R3, foundEnd          // depth == 0: the container's close
	JMP  clearBit
isOpen:
	ADD  $1, R3
clearBit:
	SUB  $1, R7, R8
	AND  R8, R7, R7            // brackets &= brackets-1
	JMP  bitLoop
onlyOpens:
	POPCNT64(R12, R8)          // no close: depth += popcount(op)
	ADD  R8, R3
	JMP  nextBlock
bulk:
	POPCNT64(R12, R9)          // depth += popcount(op) - popcount(cl)
	ADD  R9, R3
	SUB  R8, R3
nextBlock:
	ADD  $64, R2
	ADD  $64, R16
	JMP  blockLoop

foundEnd:
	ADD  R2, R9, R7            // end = pos + j + 1
	ADD  $1, R7
	MOVD R7, end+48(FP)
	MOVD R3, ndepth+56(FP)
	MOVD R4, prevEscaped+64(FP)
	MOVD R5, prevInString+72(FP)
	RET

exhausted:
	MOVD $-1, R7
	MOVD R7, end+48(FP)
	MOVD R3, ndepth+56(FP)
	MOVD R4, prevEscaped+64(FP)
	MOVD R5, prevInString+72(FP)
	RET
