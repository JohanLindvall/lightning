#include "textflag.h"

// The bit weights {1, 2, ..., 128} of the ADDP-cascade movemask
// (skipfast_arm64.s), in NORMAL order: lane j lands at bit j, so the first
// ']' is a count of trailing zeros.
DATA cntWeights<>+0(SB)/8, $0x8040201008040201
DATA cntWeights<>+8(SB)/8, $0x8040201008040201
GLOBL cntWeights<>(SB), RODATA|NOPTR, $16

// CNTMASK folds four compares (0x00/0xFF lanes) of a 64-byte block into the
// block's 64-bit mask in rd, lane j at bit j: the bit weights, a four-step
// ADDP cascade, one lane move. Clobbers t0..t3.
#define CNTMASK(t0, t1, t2, t3, v, rd) \
	VAND  V16.B16, t0, t0 \
	VAND  V16.B16, t1, t1 \
	VAND  V16.B16, t2, t2 \
	VAND  V16.B16, t3, t3 \
	VADDP t1, t0, t0      \
	VADDP t3, t2, t2      \
	VADDP t2, t0, t0      \
	VADDP t0, t0, t0      \
	VMOV  v.D[0], rd

// POPCNT counts the set bits of rs into rd through the vector unit (arm64 has
// no general-register popcount). Clobbers V31.
#define POPCNT(rs, rd) \
	FMOVD   rs, F31        \
	VCNT    V31.B8, V31.B8 \
	VUADDLV V31.B8, V31    \
	FMOVD   F31, rd

// func countKernel(data []byte, i int, c byte, hint bool) (rb, n int)
//
// The presize scan, one pass: the first ']' in data[i:] (rb, its offset
// from i, or -1 with n 0) and the number of c bytes before it — or, with
// hint, the element count CountArrayScalars wants (the commas plus one
// clamped to (rb+1)/2, or for a span with no comma 1 or 0 by whether it holds
// a byte above 0x20). The contract is count_amd64.s's; see there and count.go.
// It replaces two runtime calls over the same bytes (bytes.IndexByte, then
// bytes.Count), which on the short coordinate arrays of marine_ik and mesh
// were mostly call and setup: 13-17% of those decodes once their numbers got
// cheap.
//
// 64 bytes a step: four ']' compares ORed and reduced with one UMAXP decide
// whether the ']' is in the block, and the four c compares are subtracted
// into per-lane byte counters (flushed with UADDLV every 63 steps, before a
// lane can wrap). Only the block that holds the ']' pays for exact positions:
// the ADDP cascade makes both of its 64-bit masks, and the count is the c
// bits below the ']'. The final < 64 bytes are the buffer's LAST 64 with the
// lanes already walked shifted out of both masks, so a buffer of 64 bytes or
// more from i has no byte loop at all; a shorter one is walked a byte at a
// time.
//
// Registers: R0 the start, R1 the end, R2 the cursor, R3 c, R4 hint, R5 the
// count, R6 the steps left before a flush, R7-R12 temporaries. V0 the ']'
// splat, V1 the c splat, V16 the bit weights, V17 the per-lane counters,
// V2-V15 the block and its compares.
TEXT ·countKernel(SB), NOSPLIT, $0-56
	MOVD  data_base+0(FP), R0
	MOVD  data_len+8(FP), R1
	MOVD  i+24(FP), R2
	MOVBU c+32(FP), R3
	MOVBU hint+33(FP), R4
	MOVD  ZR, R5
	CMP   R1, R2
	BGE   notfound
	ADD   R0, R1, R1        // end
	ADD   R2, R0, R0        // start = &data[i]
	MOVD  R0, R2            // cursor
	SUB   R0, R1, R7
	CMP   $64, R7
	BLO   bytes
	VMOVI $0x5d, V0.B16     // ']'
	VDUP  R3, V1.B16
	MOVD  $cntWeights<>(SB), R7
	VLD1  (R7), [V16.B16]
	VMOVI $0, V17.B16
	MOVD  $63, R6
	SUB   $64, R1, R8       // the last start of a full block

loop:
	CMP   R8, R2
	BHI   tail
	VLD1  (R2), [V2.B16, V3.B16, V4.B16, V5.B16]
	VCMEQ V0.B16, V2.B16, V6.B16
	VCMEQ V0.B16, V3.B16, V7.B16
	VCMEQ V0.B16, V4.B16, V8.B16
	VCMEQ V0.B16, V5.B16, V9.B16
	VORR  V7.B16, V6.B16, V10.B16
	VORR  V9.B16, V8.B16, V11.B16
	VORR  V11.B16, V10.B16, V10.B16
	WORD  $0x6eaaa54a // umaxp v10.4s, v10.4s, v10.4s
	FMOVD F10, R7
	VCMEQ V1.B16, V2.B16, V12.B16
	VCMEQ V1.B16, V3.B16, V13.B16
	VCMEQ V1.B16, V4.B16, V14.B16
	VCMEQ V1.B16, V5.B16, V15.B16
	CBNZ  R7, found
	VSUB  V12.B16, V17.B16, V17.B16
	VSUB  V13.B16, V17.B16, V17.B16
	VSUB  V14.B16, V17.B16, V17.B16
	VSUB  V15.B16, V17.B16, V17.B16
	ADD   $64, R2
	SUBS  $1, R6
	BNE   loop
	VUADDLV V17.B16, V18    // flush the per-lane counters before a lane can wrap
	FMOVD F18, R7
	ADD   R7, R5
	VMOVI $0, V17.B16
	MOVD  $63, R6
	B     loop

tail:
	// Fewer than 64 bytes from the cursor: the buffer's last 64, the lanes
	// before the cursor shifted out of both masks.
	CMP   R1, R2
	BEQ   notfoundFlush
	SUB   $64, R1, R9
	VLD1  (R9), [V2.B16, V3.B16, V4.B16, V5.B16]
	VCMEQ V0.B16, V2.B16, V6.B16
	VCMEQ V0.B16, V3.B16, V7.B16
	VCMEQ V0.B16, V4.B16, V8.B16
	VCMEQ V0.B16, V5.B16, V9.B16
	VCMEQ V1.B16, V2.B16, V12.B16
	VCMEQ V1.B16, V3.B16, V13.B16
	VCMEQ V1.B16, V4.B16, V14.B16
	VCMEQ V1.B16, V5.B16, V15.B16
	CNTMASK(V6.B16, V7.B16, V8.B16, V9.B16, V6, R10)
	CNTMASK(V12.B16, V13.B16, V14.B16, V15.B16, V12, R11)
	SUB   R9, R2, R9        // lanes already walked
	LSR   R9, R10, R10
	LSR   R9, R11, R11
	CBZ   R10, notfoundFlush
	B     haveMasks

found:
	// The ']' is in the block at the cursor: both of its masks exactly.
	CNTMASK(V6.B16, V7.B16, V8.B16, V9.B16, V6, R10)
	CNTMASK(V12.B16, V13.B16, V14.B16, V15.B16, V12, R11)

haveMasks:
	// R10: the ']' lanes from the cursor, R11: the c lanes.
	RBIT  R10, R10
	CLZ   R10, R10          // the ']' lane
	MOVD  $1, R12
	LSL   R10, R12, R12
	SUB   $1, R12, R12
	AND   R12, R11, R11     // the c lanes below it
	POPCNT(R11, R11)
	ADD   R11, R5
	VUADDLV V17.B16, V18
	FMOVD F18, R7
	ADD   R7, R5
	SUB   R0, R2, R2
	ADD   R10, R2, R10      // rb

result:
	// R10 = rb, R5 = the count.
	MOVD  R10, rb+40(FP)
	CBNZ  R4, hint
	MOVD  R5, n+48(FP)
	RET

hint:
	CBZ   R5, blank
	ADD   $1, R5            // elements = commas + 1
	// Clamp to the element count the span can hold: n elements need n-1
	// commas and a byte each, so rb >= 2n-1 and n <= (rb+1)/2 (see
	// CountArrayScalars for why the bound matters and never clips an honest
	// count).
	ADD   $1, R10, R7
	LSR   $1, R7, R7
	CMP   R7, R5
	CSEL  HI, R7, R5, R5
	MOVD  R5, n+48(FP)
	RET

blank:
	// No comma before the ']' at start+rb: one element unless every byte is
	// whitespace. The walk stops at the first byte above 0x20, which in any
	// one-element array is the element's first byte.
	ADD   R0, R10, R7

blankLoop:
	CMP   R7, R0
	BEQ   empty
	MOVBU.P 1(R0), R8
	CMP   $0x20, R8
	BLS   blankLoop
	MOVD  $1, R8
	MOVD  R8, n+48(FP)
	RET

empty:
	MOVD  ZR, n+48(FP)
	RET

bytes:
	// Fewer than 64 bytes from i: a byte at a time.
	CMP   R1, R2
	BEQ   notfound
	MOVBU (R2), R7
	CMP   $0x5d, R7
	BEQ   bytesFound
	CMP   R3, R7
	CINC  EQ, R5, R5
	ADD   $1, R2
	B     bytes

bytesFound:
	SUB   R0, R2, R10
	B     result

notfoundFlush:
notfound:
	MOVD  $-1, R7
	MOVD  R7, rb+40(FP)
	MOVD  ZR, n+48(FP)
	RET
