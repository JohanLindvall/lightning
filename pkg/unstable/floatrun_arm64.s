#include "textflag.h"

// Constants for parseFloatRunNEON, loaded as two VLD1s: the digit-fold
// weights (10,1 per byte pair for the MUL; 100,100,1,1 per word for the UDOT;
// 10^4,1 per word pair for the second MUL), the bit weights of the ADDP
// cascade in the reversed order the walk wants (see intrun_arm64.s), the
// top-digit UDOT weights (0,100,10,1: one to three digits in lanes 1-3 of a
// word), and the doubles 1e8 and 1 that join the two eight-digit halves.
DATA frConst<>+0(SB)/8, $0x010a010a010a010a
DATA frConst<>+8(SB)/8, $0x010a010a010a010a
DATA frConst<>+16(SB)/8, $0x0101646401016464
DATA frConst<>+24(SB)/8, $0x0101646401016464
DATA frConst<>+32(SB)/8, $0x0102040810204080
DATA frConst<>+40(SB)/8, $0x0102040810204080
DATA frConst<>+48(SB)/8, $0x0000000100002710
DATA frConst<>+56(SB)/8, $0x0000000100002710
DATA frConst<>+64(SB)/8, $0x010a6400010a6400
DATA frConst<>+72(SB)/8, $0x010a6400010a6400
DATA frConst<>+80(SB)/8, $0x4197d78400000000
DATA frConst<>+88(SB)/8, $0x3ff0000000000000
GLOBL frConst<>(SB), RODATA|NOPTR, $96

// The assembler floor is Go 1.25 (see intrun_arm64.s), which has no CMHI,
// vector MUL, UDOT, UADDLP, UCVTF, vector FMUL or FADDP: those are WORDs, on
// lines of their own so sveasm can derive them from the mnemonics.

// CASCADE folds the four class compares in t0..t3 (0x00/0xFF per lane) to
// the 16-byte form two ADDP steps short of the block's 64-bit mask, in t0 —
// the bit-weight/ADDP cascade of skipfast_arm64.s, with the reversed weights
// in V10. Clobbers t1..t3.
#define CASCADE(t0, t1, t2, t3) \
	VAND  V10.B16, t0, t0  \
	VAND  V10.B16, t1, t1  \
	VAND  V10.B16, t2, t2  \
	VAND  V10.B16, t3, t3  \
	VADDP t1, t0, t0       \
	VADDP t3, t2, t2       \
	VADDP t2, t0, t0

// FRPREP1 and FRPREP2, with the four not-digit compares between them (CMHI
// of the bytes minus '0' against 9, into V12-V15), classify the 64-byte
// block at address a into two bit-reversed masks, lane j at bit 63-j: rD, not
// a digit, and rC, a comma. The block's bytes stay in V0-V3 for the
// whitespace class. Clobbers V12-V19.
#define FRPREP1(a) \
	VLD1  (a), [V0.B16, V1.B16, V2.B16, V3.B16] \
	VSUB  V4.B16, V0.B16, V12.B16               \
	VSUB  V4.B16, V1.B16, V13.B16               \
	VSUB  V4.B16, V2.B16, V14.B16               \
	VSUB  V4.B16, V3.B16, V15.B16

#define FRPREP2(rD, rC) \
	CASCADE(V12.B16, V13.B16, V14.B16, V15.B16) \
	VCMEQ  V6.B16, V0.B16, V16.B16              \
	VCMEQ  V6.B16, V1.B16, V17.B16              \
	VCMEQ  V6.B16, V2.B16, V18.B16              \
	VCMEQ  V6.B16, V3.B16, V19.B16              \
	CASCADE(V16.B16, V17.B16, V18.B16, V19.B16) \
	VADDP  V12.B16, V16.B16, V16.B16            \
	VREV64 V16.B16, V16.B16                     \
	VMOV   V16.D[1], rD                         \
	VMOV   V16.D[0], rC

// FRPREP2V is FRPREP2 for a window classified ahead: the two masks stay in
// the vector unit, parked in V28 (not-digit in D[1], comma in D[0]), until
// the walk reaches the window. Clobbers V12-V19.
#define FRPREP2V \
	CASCADE(V12.B16, V13.B16, V14.B16, V15.B16) \
	VCMEQ  V6.B16, V0.B16, V16.B16              \
	VCMEQ  V6.B16, V1.B16, V17.B16              \
	VCMEQ  V6.B16, V2.B16, V18.B16              \
	VCMEQ  V6.B16, V3.B16, V19.B16              \
	CASCADE(V16.B16, V17.B16, V18.B16, V19.B16) \
	VADDP  V12.B16, V16.B16, V16.B16            \
	VREV64 V16.B16, V28.B16

// WSPACK finishes the third mask, not whitespace (c > 0x20), from its four
// compares in t0..t3 (CMHI against the 0x20 splat V7), into rW; v is t0's
// bare register name, for the lane move.
#define WSPACK(t0, t1, t2, t3, v, rW) \
	CASCADE(t0, t1, t2, t3) \
	VADDP  t0, t0, t0       \
	VREV64 t0, t0           \
	VMOV   v.D[0], rW

// SHORTCONV converts the number at lane R16 of the block at R11 — L1 (R17)
// integer and L2 (R14) fraction digits, 15 at most, the sign in R15 — and
// stores it at the out cursor R3, which it advances. The sixteen bytes at the
// first digit, minus '0', are right-aligned under the (L1, L2) TBL control
// with the '.' dropped; a byte MUL by 10,1 and a UDOT by 100,100,1,1 make four
// 4-digit groups; a word MUL by 10^4,1 and a UADDLP two 8-digit halves; and
// those are converted, joined (exactly: the top half is below 10^7) and
// divided once by ±10^L2 — Clinger's fast path, correctly rounded, the divisor
// carrying the sign. Clobbers R20, V22-V25.
#define SHORTCONV \
	WORD  $0x3cf06976 /* ldr q22, [x11, x16] */ \
	ADD   R17<<4, R14, R20 \
	WORD  $0x3cf47957 /* ldr q23, [x10, x20, lsl #4] */ \
	VSUB  V4.B16, V22.B16, V22.B16 \
	VTBL  V23.B16, [V22.B16], V22.B16 \
	WORD  $0x4e289ed6 /* mul v22.16b, v22.16b, v8.16b */ \
	VMOVI $0, V24.B16 \
	WORD  $0x6e8996d8 /* udot v24.4s, v22.16b, v9.16b */ \
	WORD  $0x4eab9f18 /* mul v24.4s, v24.4s, v11.4s */ \
	WORD  $0x6ea02b18 /* uaddlp v24.2d, v24.4s */ \
	WORD  $0x6e61db18 /* ucvtf v24.2d, v24.2d */ \
	WORD  $0x6e75df18 /* fmul v24.2d, v24.2d, v21.2d */ \
	WORD  $0x7e70db18 /* faddp d24, v24.2d */ \
	ADD   R15<<4, R14, R20 \
	FMOVD (R10)(R20<<3), F25 \
	FDIVD F25, F24, F24 \
	FMOVD.P F24, 8(R3)

// LONGCONV converts a number of sixteen to nineteen digits — SHORTCONV's
// inputs, and its end e in R12 — and stores it at the out cursor R3, which it
// advances, falling through (to done) when it has. Its last sixteen digits
// come from the loads at its first digit and ending at its last byte under a
// two-register TBL control, the one to three before them from the first load
// into lanes 1-3 of a word of their own (the top-digit UDOT); the mantissa is
// joined in a general register, below 10^19, and it is Clinger below 2^53 or
// Eisel-Lemire above: eiselLemire64 transcribed, its refinement with the
// power's low word included, branching to decline exactly where it declines —
// a product still ambiguous after the refinement, or an exact halfway value.
// (The amd64 LONGCONV declines wherever the refinement would be needed. That
// looks like 0.13% of canada's numbers, but the cost is not the numbers: a
// refused point goes to the reader's per-point path, and a refusal at the
// first point of a walk switches the walk off for the rest of its ring — on
// canada, whose six-decimal coordinates printed to fifteen fraction digits
// need the refinement one number in ~260 (429 of 111,080, e.g.
// "46.851662000000033"), that was 8,573 of 55,563 points.) The range tests
// go for the amd64 macro's reason: a mantissa of 2^53 or more over at most
// 10^19 is far from both subnormals and infinity, and a rounding carry into
// bit 53 completes the exponent field by itself. k1e8 holds 10^8 (for the
// join); a caller that hands a declined element back from its first byte
// keeps that somewhere first, as R16 is clobbered. Clobbers R12, R14, R16,
// R17, R20, V22-V27, F30, F31; el, norefine, round, refine and done are
// labels.
#define LONGCONV(decline, el, norefine, round, refine, done, k1e8) \
	WORD  $0x3cf06976 /* ldr q22, [x11, x16] */ \
	ADD   R11, R12, R20 \
	WORD  $0x3cdf0297 /* ldur q23, [x20, #-16] */ \
	ADD   R17<<5, R14, R20 \
	ADD   $160, R20, R20 \
	ADD   R20<<5, R10, R20 \
	VLD1  (R20), [V24.B16, V25.B16] \
	VSUB  V4.B16, V22.B16, V22.B16 \
	VSUB  V4.B16, V23.B16, V23.B16 \
	VTBL  V24.B16, [V22.B16, V23.B16], V24.B16 \
	VTBL  V25.B16, [V22.B16], V25.B16 \
	WORD  $0x4e289f18 /* mul v24.16b, v24.16b, v8.16b */ \
	VMOVI $0, V26.B16 \
	VMOVI $0, V27.B16 \
	WORD  $0x6e89971a /* udot v26.4s, v24.16b, v9.16b */ \
	WORD  $0x6e94973b /* udot v27.4s, v25.16b, v20.16b */ \
	WORD  $0x4eab9f5a /* mul v26.4s, v26.4s, v11.4s */ \
	WORD  $0x6ea02b5a /* uaddlp v26.2d, v26.4s */ \
	VMOV  V27.S[0], R17 \
	VMOV  V26.D[0], R20 \
	VMOV  V26.D[1], R12 \
	MADD  k1e8, R20, R17, R20 \
	MADD  k1e8, R12, R20, R20 \
	LSR   $53, R20, R12 \
	CBNZ  R12, el \
	SCVTFD R20, F24 \
	ADD   R15<<5, R14, R20 \
	ADD   $512, R20, R20 \
	FMOVD (R10)(R20<<3), F25 \
	FDIVD F25, F24, F24 \
	FMOVD.P F24, 8(R3) \
	B     done \
el: \
	CLZ   R20, R16 \
	LSL   R16, R20, R20 \
	ADD   $576, R14, R14 \
	FMOVD R14, F30 \
	MOVD  (R10)(R14<<3), R17 \
	ADD   $20, R14, R12 \
	MOVD  (R10)(R12<<3), R12 \
	SUB   R16, R12, R16 \
	UMULH R17, R20, R12 \
	MUL   R17, R20, R17 \
	AND   $0x1ff, R12, R14 \
	CMP   $0x1ff, R14 \
	BEQ   refine \
norefine: \
	LSR   $63, R12, R14 \
	ADD   R14, R16, R16 \
	ADD   $9, R14, R14 \
	LSR   R14, R12, R14 \
	CBNZ  R17, round \
	TST   $0x1ff, R12 \
	BNE   round \
	AND   $3, R14, R20 \
	CMP   $1, R20 \
	BEQ   decline \
round: \
	AND   $1, R14, R17 \
	ADD   R14>>1, R17, R14 \
	ADD   R16<<52, R14, R14 \
	ORR   R15<<63, R14, R14 \
	MOVD.P R14, 8(R3) \
	B     done \
refine: \
	ADDS  R20, R17, R14 \
	BCC   norefine \
	FMOVD F30, R14 \
	ADD   $40, R14, R14 \
	MOVD  (R10)(R14<<3), R14 \
	FMOVD R15, F31 \
	UMULH R14, R20, R15 \
	MUL   R14, R20, R14 \
	ADDS  R20, R14, R14 \
	CSET  HS, R20 \
	ADDS  R15, R17, R17 \
	ADC   ZR, R12, R12 \
	FMOVD F31, R15 \
	AND   $0x1ff, R12, R14 \
	CMP   $0x1ff, R14 \
	BNE   norefine \
	CMN   $1, R17 \
	BNE   norefine \
	CBNZ  R20, decline \
	B     norefine \
done:

// func parseFloatRunNEON(data []byte, i int, out []float64) (n, p, closed int)
//
// The arm64 decimal-array kernel: the contract of the amd64 parseFloatRunAVX2
// and parseFloatRunVBMI (floatrun_amd64.s), in one body that takes numbers of
// up to 19 digits. It parses as many "ws* -? digits (. digits)? ws* ','"
// groups as it can from data[i:], one float64 per group into out, and the
// array's last element when a ']' ends it (closed = 1, p at the ']'). It stops
// — with p at the element it did not take, after any whitespace before it —
// when fewer than 80 bytes remain from the block, out is full, or an element is
// anything else: an exponent, 20+ digits, a '+', null, a '.' without digits
// after it. Every stop position is a state the scalar loop resumes from, and
// every value it writes is the one strconv returns.
//
// The block walk is parseIntRunNEON's (intrun_arm64.s): 64-byte blocks at a
// fixed 48-byte stride, the next block classified while this one is walked,
// bit-reversed class masks so every run length is one LSL and one CLZ, the
// capacity bounded once per block rather than per element, and the whitespace
// class computed only for blocks that need it — with, here, an element loop of
// its own for those blocks (welem), which skips the whitespace in front of
// every element unconditionally instead of first trying the compact shape.
// What differs from the integer walk is the cursor: a decimal's end costs two
// runs and a '.' test to find, so the walk does not derive the next element
// from it — it walks COMMAS (LSL, CLZ, ADD from each element's start), and each
// element's measure only has to agree with the comma it already has. That
// keeps the loop-carried chain at a few cycles whatever the number looks
// like; the measure, the '.' and sign loads and the conversion hang off it and
// overlap from element to element.
//
// The conversion gathers the digits from memory at the first digit — a TBL
// under a per-(L1, L2) control right-aligns them and drops the '.' — and folds
// them in the vector unit: a byte MUL by 10,1 and a UDOT by 100,100,1,1 make
// four 4-digit groups, a word MUL by 10^4,1 and a UADDLP two 8-digit halves.
// Up to 15 digits (short) the halves are joined in double precision, exactly
// (the top half is below 10^7), and divided once by ±10^L2: Clinger's fast
// path, correctly rounded, the divisor carrying the sign. From 16 to 19 digits
// (long) a second TBL and UDOT fold the one to three top digits, the mantissa
// is joined in a general register, and it is Clinger again below 2^53 or
// Eisel-Lemire above — eiselLemire64 transcribed with its low-word refinement
// (which the amd64 LONGCONV leaves out, declining there instead), declining
// (handing the element to the scalar loop) exactly where eiselLemire64
// declines. The long loads are sixteen bytes at the first digit and sixteen
// ending at the number's last byte, so no load leaves the 80 bytes the block
// guarantees.
//
// Vector MUL, UCVTF and FDIV all issue on the one V0 pipe of a Neoverse N2
// (measured), while UDOT, UADDLP, TBL and FMUL issue on either; the fold keeps
// the V0-only ops to four a number.
//
// Registers: R0 data base, R1 len-80 (the last block start), R2 s (the
// block's absolute start), R3 the out cursor, R4 out end, R5 the lane limit
// (48, or less when out has fewer than 25 slots left), R6 b (the element
// region's lane), R7/R8/R9 the reversed not-digit / comma / not-whitespace
// masks (R19 whether the last is computed, R22 whether the block before
// needed it), R10 floatRunTab, R11 the block's address, R13 the element's
// comma (or ']') lane, R14 L2, R15 the sign, R16 the first digit's lane, R17
// L1, R21 set when the element ends the array, R23-R25 the next block's masks
// and R26 its address (0: none), R12/R20 temporaries (R20 also names
// wsCompute's return site). V4-V7 the '0', 9, ',' and 0x20 splats, V8-V11 and
// V20-V21 the constants above, V0-V3/V12-V19 the classification's
// temporaries, V22-V27 the conversion's.
TEXT ·parseFloatRunNEON(SB), NOSPLIT, $0-80
	MOVD  data_base+0(FP), R0
	MOVD  data_len+8(FP), R1
	SUB   $80, R1
	MOVD  i+24(FP), R2
	MOVD  out_base+32(FP), R3
	MOVD  out_len+40(FP), R4
	ADD   R4<<3, R3, R4
	MOVD  $·floatRunTab(SB), R10
	MOVD  $frConst<>(SB), R12
	VLD1  (R12), [V8.B16, V9.B16, V10.B16, V11.B16]
	ADD   $64, R12
	VLD1  (R12), [V20.B16, V21.B16]
	VMOVI $0x30, V4.B16 // '0'
	VMOVI $9, V5.B16
	VMOVI $0x2c, V6.B16 // ','
	VMOVI $0x20, V7.B16 // whitespace is c <= 0x20, SkipWS's rule
	MOVD  ZR, R21
	MOVD  ZR, R19
	MOVD  ZR, R6

block:
	// The block at s, classified here: the first one, and the one a
	// straddle restarts at an element's own start. Its whitespace class is
	// computed on demand.
	CMP   R1, R2
	BGT   done // fewer than 80 bytes from s: the scalar loop takes it from here
	ADD   R0, R2, R11
	FRPREP1(R11)
	WORD  $0x6e25358c // cmhi v12.16b, v12.16b, v5.16b
	WORD  $0x6e2535ad // cmhi v13.16b, v13.16b, v5.16b
	WORD  $0x6e2535ce // cmhi v14.16b, v14.16b, v5.16b
	WORD  $0x6e2535ef // cmhi v15.16b, v15.16b, v5.16b
	FRPREP2(R7, R8)
	MOVD  R19, R22
	MOVD  ZR, R9
	MOVD  ZR, R19

blockAt:
	// The block at s is classified and its walk starts at b. First the
	// capacity: the block stores at most 25 values (elements start at least
	// two lanes apart, so at most 24 start in its first 48 lanes, plus the one
	// a ']' closes), so with 25 slots free the walk runs unchecked to lane 48.
	MOVD  $48, R5
	SUB   R3, R4, R12
	CMP   $200, R12
	BLO   capacity

pipeline:
	// Then classify the block at s+48 if there is one, so that work overlaps
	// the walk — its whitespace class too when this block's predecessor
	// needed it (R22), since an array's separators are alike. Only a walk that
	// can reach lane 48 needs it: when the capacity limit ends the walk inside
	// this block — a short array whose target was presized to it, the common
	// case on marine_ik — that classification was forty instructions for
	// nothing.
	ADD   $48, R2, R12
	CMP   R1, R12
	BGT   noNext
	ADD   $48, R11, R26
	FRPREP1(R26)
	WORD  $0x6e25358c // cmhi v12.16b, v12.16b, v5.16b
	WORD  $0x6e2535ad // cmhi v13.16b, v13.16b, v5.16b
	WORD  $0x6e2535ce // cmhi v14.16b, v14.16b, v5.16b
	WORD  $0x6e2535ef // cmhi v15.16b, v15.16b, v5.16b
	FRPREP2(R23, R24)
	MOVD  ZR, R25
	CBZ   R22, dispatch
	WORD  $0x6e27340c // cmhi v12.16b, v0.16b, v7.16b
	WORD  $0x6e27342d // cmhi v13.16b, v1.16b, v7.16b
	WORD  $0x6e27344e // cmhi v14.16b, v2.16b, v7.16b
	WORD  $0x6e27346f // cmhi v15.16b, v3.16b, v7.16b
	WSPACK(V12.B16, V13.B16, V14.B16, V15.B16, V12, R25)
	B     dispatch

noNext:
	MOVD  ZR, R26

dispatch:
	CBNZ  R19, welem

elem:
	// The element region starting at lane b, in a block with no whitespace
	// class: the compact shape first.
	CMP   R5, R6
	BHS   limit
	LSL   R6, R8, R12       // commas << b
	CBZ   R12, noComma
	CLZ   R12, R13
	ADD   R6, R13, R13      // c: this element's comma
	MOVBU (R11)(R6), R12    // a '-' at b: s = b+1
	CMP   $0x2d, R12
	CSET  EQ, R15
	ADD   R15, R6, R16      // s
	LSL   R16, R7, R12      // not-digit << s
	CLZ   R12, R17          // L1: 0 when the byte at s is not a digit
	CBZ   R17, elemWS

elemDigits:
	ADD   R16, R17, R12     // e1: the byte after the integer digits
	MOVBU (R11)(R12), R14
	CMP   $0x2e, R14
	BNE   elemNoFrac
	ADD   $1, R12, R12
	LSL   R12, R7, R14      // not-digit << e1+1
	CLZ   R14, R14          // L2
	CBZ   R14, stopElem     // a '.' without digits after it
	ADD   R14, R12, R12     // e: the byte after the number
	CMP   R13, R12
	BNE   wsBeforeComma

convert:
	// The element is the number at s, L1 integer and L2 fraction digits,
	// ending at e (R12), delimited at R13 by a comma or, with R21 set, by the
	// array's ']'.
	ADD   R17, R14, R20     // L
	CMP   $15, R20
	BHI   long
	SHORTCONV
	CBNZ  R21, closed
	ADD   $1, R13, R6       // b: just past the comma
	CBZ   R19, elem
	B     welem

elemNoFrac:
	MOVD  ZR, R14
	CMP   R13, R12
	BEQ   convert

wsBeforeComma:
	// e < c: the bytes between must be whitespace — or the first one that
	// is not must be the array's ']', the comma being the enclosing
	// container's, and then this element is the last.
	MOVBU (R11)(R12), R20
	CMP   $0x5d, R20
	BEQ   closeAtE
	CBNZ  R19, wsBeforeCommaAt
	MOVD  $2, R20
	B     wsCompute

wsBeforeCommaAt:
	LSL   R12, R9, R20
	CLZ   R20, R20
	ADD   R12, R20, R20     // the first byte after e that is not whitespace
	CMP   R13, R20
	BEQ   convert
	MOVBU (R11)(R20), R13
	CMP   $0x5d, R13
	BNE   stopElem
	MOVD  R20, R13          // the ']' delimits this, the array's last element
	MOVD  $1, R21
	B     convert

closeAtE:
	MOVD  R12, R13
	MOVD  $1, R21
	B     convert

welem:
	// The element region starting at lane b, in a block whose whitespace
	// class is computed: skip the whitespace in front unconditionally.
	CMP   R5, R6
	BHS   limit
	LSL   R6, R8, R12       // commas << b
	CBZ   R12, noComma
	CLZ   R12, R13
	ADD   R6, R13, R13      // c

welemAt:
	LSL   R6, R9, R12       // not-whitespace << b
	CLZ   R12, R12
	ADD   R6, R12, R16      // the first byte that is not whitespace (c at most)
	MOVBU (R11)(R16), R12
	CMP   $0x2d, R12
	CSET  EQ, R15
	ADD   R15, R16, R16     // s
	LSL   R16, R7, R12
	CLZ   R12, R17          // L1
	CBNZ  R17, elemDigits
	B     stopElem          // not a number (an empty element among them)

elemWS:
	// The compact loop met a byte at s that is not a digit: after a '-'
	// this is not a number; else it may be whitespace before one, and the
	// block's whitespace class is computed and its walk continues in welem.
	CBNZ  R15, stopElem
	MOVD  $1, R20
	B     wsCompute

long:
	// Sixteen to nineteen digits.
	CMP   $19, R20
	BHI   stopElem          // 20 digits or more: the scalar loop's
	SUB   R15, R16, R20
	MOVD  R20, (R3)         // parked for elDecline: the element's first byte
	MOVD  $100000000, R6    // b is free until the comma sets it again
	LONGCONV(elDecline, el, elNoRefine, elRound, elRefine, converted, R6)
	CBNZ  R21, closed
	ADD   $1, R13, R6
	CBZ   R19, elem
	B     welem

elDecline:
	// Hand the element back from its first byte, parked in its output slot.
	MOVD  (R3), R16
	MOVD  ZR, R15
	B     stopElem

noComma:
	// No comma left in the block: the element can only be the array's last,
	// its ']' in this block; one that runs past the block restarts the next
	// block at it.
	MOVBU (R11)(R6), R12
	CMP   $0x2d, R12
	CSET  EQ, R15
	ADD   R15, R6, R16
	LSL   R16, R7, R12
	CLZ   R12, R17
	CBZ   R17, lastWS

lastDigits:
	ADD   R16, R17, R12     // e1
	CMP   $64, R12
	BHS   straddle          // digits to the block's end
	MOVBU (R11)(R12), R14
	CMP   $0x2e, R14
	BNE   lastNoFrac
	ADD   $1, R12, R12
	// A '.' in lane 63 has its digits in the next block — and an LSL by 64
	// is an LSL by 0, so the test is not optional.
	CMP   $64, R12
	BHS   straddle
	LSL   R12, R7, R14
	CLZ   R14, R14          // L2
	CBZ   R14, stopElem
	ADD   R14, R12, R12     // e
	CMP   $64, R12
	BHS   straddle
	B     lastTerm

lastNoFrac:
	MOVD  ZR, R14

lastTerm:
	MOVBU (R11)(R12), R13
	CMP   $0x5d, R13
	BNE   lastTermWS
	MOVD  R12, R13          // the ']' right after the number
	MOVD  $1, R21
	B     convert

lastTermWS:
	CBNZ  R19, lastTermAt
	MOVD  $4, R20
	B     wsCompute

lastTermAt:
	LSL   R12, R9, R13
	CLZ   R13, R13
	ADD   R12, R13, R13     // the first byte after e that is not whitespace
	CMP   $64, R13
	BHS   straddle
	MOVBU (R11)(R13), R20
	CMP   $0x5d, R20
	BNE   stopElem
	MOVD  $1, R21
	B     convert

lastWS:
	CBNZ  R15, stopElem
	CBNZ  R19, lastWSAt
	MOVD  $3, R20
	B     wsCompute

lastWSAt:
	LSL   R6, R9, R12
	CLZ   R12, R12
	ADD   R6, R12, R16      // s
	CMP   $64, R16
	BHS   straddle          // whitespace to the block's end
	MOVBU (R11)(R16), R12
	CMP   $0x2d, R12
	CSET  EQ, R15
	ADD   R15, R16, R16
	CMP   $64, R16
	BHS   straddle          // a '-' in lane 63
	LSL   R16, R7, R12
	CLZ   R12, R17
	CBZ   R17, stopElem
	B     lastDigits

wsCompute:
	// The whitespace class, not whitespace (c > 0x20), for the block at R11,
	// into R9 with R19 set: computed only when the walk first needs it, from
	// a reload of the block. R20 says which site asked.
	VLD1  (R11), [V12.B16, V13.B16, V14.B16, V15.B16]
	WORD  $0x6e273590 // cmhi v16.16b, v12.16b, v7.16b
	WORD  $0x6e2735b1 // cmhi v17.16b, v13.16b, v7.16b
	WORD  $0x6e2735d2 // cmhi v18.16b, v14.16b, v7.16b
	WORD  $0x6e2735f3 // cmhi v19.16b, v15.16b, v7.16b
	WSPACK(V16.B16, V17.B16, V18.B16, V19.B16, V16, R9)
	MOVD  $1, R19
	CMP   $2, R20
	BLO   welemAt
	BEQ   wsBeforeCommaAt
	CMP   $3, R20
	BEQ   lastWSAt
	B     lastTermAt

capacity:
	// Fewer than 25 slots: the element after the free-th comma at or after b
	// has no slot, so the lane just past that comma is the limit — exactly,
	// which a count of the block's commas is not: an array is followed by
	// more of the document, whose commas would make a target sized to the
	// array stop short of its end, and on a document of short arrays that is
	// every array. A walk of at most 24 comma bits, once per array.
	LSR   $3, R12, R12      // free slots
	CBZ   R12, done
	LSL   R6, R8, R20       // commas << b
	MOVD  R6, R5

capLoop:
	CBZ   R20, capAll       // fewer commas than slots: no limit in this block
	CLZ   R20, R16
	ADD   R16, R5, R5
	ADD   $1, R5, R5        // the lane after this comma
	LSL   R16, R20, R20
	LSL   $1, R20, R20
	SUBS  $1, R12, R12
	BNE   capLoop
	CMP   $48, R5
	BHS   capAll
	MOVD  ZR, R26           // the walk ends in this block: no next block to classify
	B     dispatch

capAll:
	MOVD  $48, R5
	B     pipeline

limit:
	// b at the lane limit: 48 means the element starts in the next block,
	// anything less that out is full.
	CMP   $48, R6
	BLO   done

step:
	// The element starts in the next block, at lane b-48, which is already
	// classified: make it the current one.
	CBZ   R26, done         // no next block: the scalar loop takes it from here
	ADD   $48, R2
	SUB   $48, R6
	MOVD  R26, R11
	MOVD  R23, R7
	MOVD  R24, R8
	MOVD  R25, R9           // its whitespace class, if R22 had it computed ...
	MOVD  R22, R12
	MOVD  R19, R22          // ... and whether this block needed it decides for the next
	MOVD  R12, R19
	B     blockAt

straddle:
	// The element runs past the block: start the next block at its first
	// byte that is not whitespace (its '-' included), or past the block when
	// the whitespace reaches the end — every straddle is found after s is
	// measured. At lane 0 that would not move, so stop.
	SUB   R15, R16, R12
	CMP   $64, R12
	BLS   straddleAt
	MOVD  $64, R12

straddleAt:
	CBZ   R12, stopElem
	ADD   R12, R2, R2
	MOVD  ZR, R6
	B     block

stopElem:
	// A stop at an element measured from s: resume at its first byte that
	// is not whitespace, its '-' included.
	SUB   R15, R16, R6

done:
	ADD   R6, R2, R2
	MOVD  out_base+32(FP), R12
	SUB   R12, R3, R12
	LSR   $3, R12, R12
	MOVD  R12, n+56(FP)
	MOVD  R2, p+64(FP)
	MOVD  ZR, closed+72(FP)
	RET

closed:
	ADD   R13, R2, R2       // p: the ']'
	MOVD  out_base+32(FP), R12
	SUB   R12, R3, R12
	LSR   $3, R12, R12
	MOVD  R12, n+56(FP)
	MOVD  R2, p+64(FP)
	MOVD  $1, R12
	MOVD  R12, closed+72(FP)
	RET

// func parseFloatPointsNEON(data []byte, i int, out []float64, n int) (np, p, closed int)
//
// The points walk over an array of fixed-size numeric arrays — a ring of
// coordinate points, "[[x,y],[x,y],…]" — for DecodeFloat64Points: the
// contract of the amd64 parseFloatPointsVBMI. i is at a point's '[' (or
// whitespace before it); out is the flat backing of the points' spare slots,
// n floats a point. It converts as many points as it can, each exactly n
// numbers of the flat walk's shape, writing the k-th point's numbers at
// out[n*k:], and returns how many (np), and either closed = 1 with p at the
// ring's ']', or p at the start of the first point it did not take — which the
// reader then decodes element by element. A point is counted only once the
// ',' after it has been seen (or the ring's ']'), so p is always a place the
// reader resumes by reading an element; the separator is looked for past the
// point's window only while the bytes between are whitespace.
//
// One window per point, at the point: a coordinate is at most 21 bytes, so a
// point of up to three and its brackets fits in 64. The window is classified
// as a block of the flat walk is — not-digit and comma masks, the whitespace
// class only for a window that needs it (a pretty-printed ring) — and each
// number is measured and converted exactly as there, SHORTCONV and LONGCONV
// being the flat walk's own macros: the inner numbers are delimited by the
// commas at or after their start, the last by the point's ']' right after it
// (or after whitespace). A point of fewer numbers meets its ']' where an inner
// comma must be, one of more meets a comma where its ']' must be, and either
// is handed back whole — as is anything the flat walk would not take.
//
// The walk's chain is point to point: a window's classification, its commas,
// the next window's address. So each point classifies its successor's window
// as soon as its own commas are known — at the byte after the n-th comma,
// which in a ring of n-number points is the separator — and keeps the two
// masks in V28 until the walk gets there; a point that ends anywhere else
// only means that window is classified when it is reached.
//
// Registers as in parseFloatRunNEON, plus R5 n, R21 the numbers left in the
// point, R22 the out cursor at the point's start and R23 its region's start
// (what a hand-back restores and returns), R24 the lane of its '[', R25 10^8
// for LONGCONV, R26 the offset of the window classified ahead (-1: none) and
// V28 its masks.
TEXT ·parseFloatPointsNEON(SB), NOSPLIT, $0-88
	MOVD  data_base+0(FP), R0
	MOVD  data_len+8(FP), R1
	SUB   $80, R1
	MOVD  i+24(FP), R2
	MOVD  out_base+32(FP), R3
	MOVD  out_len+40(FP), R4
	ADD   R4<<3, R3, R4
	MOVD  n+56(FP), R5
	MOVD  $·floatRunTab(SB), R10
	MOVD  $frConst<>(SB), R12
	VLD1  (R12), [V8.B16, V9.B16, V10.B16, V11.B16]
	ADD   $64, R12
	VLD1  (R12), [V20.B16, V21.B16]
	VMOVI $0x30, V4.B16 // '0'
	VMOVI $9, V5.B16
	VMOVI $0x2c, V6.B16 // ','
	VMOVI $0x20, V7.B16 // whitespace is c <= 0x20, SkipWS's rule
	MOVD  $100000000, R25   // LONGCONV's join
	MOVD  $-1, R26          // no window classified ahead

ppoint:
	// The point whose region starts at s: its window, classified — here, or
	// ahead of time by the point before it (R26).
	MOVD  R2, R23           // the hand-back position
	MOVD  R3, R22           // and what was written before the point
	CMP   R1, R2
	BGT   pstop             // fewer than 80 bytes from s
	SUB   R3, R4, R12
	CMP   R5<<3, R12
	BLO   pstop             // no room for a whole point
	ADD   R0, R2, R11
	CMP   R26, R2
	BEQ   pahead
	FRPREP1(R11)
	WORD  $0x6e25358c // cmhi v12.16b, v12.16b, v5.16b
	WORD  $0x6e2535ad // cmhi v13.16b, v13.16b, v5.16b
	WORD  $0x6e2535ce // cmhi v14.16b, v14.16b, v5.16b
	WORD  $0x6e2535ef // cmhi v15.16b, v15.16b, v5.16b
	FRPREP2(R7, R8)
	B     pwindow

pahead:
	VMOV  V28.D[1], R7
	VMOV  V28.D[0], R8

pwindow:
	MOVD  ZR, R19
	MOVBU (R11), R12
	CMP   $0x5b, R12
	BNE   plead             // whitespace before the '[', or not a point
	MOVD  ZR, R24           // the '[' in lane 0
	MOVD  $1, R6            // b: past it

pfirst:
	// Classify the next point's window now, at the byte after the n-th comma
	// from b: in a ring of points of n numbers that comma is the separator
	// (whitespace holds none), and the window's address is then known long
	// before this point's numbers are measured, which is the walk's chain —
	// classification is ~23 cycles of latency, and a point whose successor
	// waited for its last number and its ']' paid it on top of theirs. A
	// point that turns out not to end there, or no n-th comma in the window,
	// only means the walk classifies the next window when it reaches it.
	LSL   R6, R8, R12       // commas << b
	MOVD  R5, R21
	MOVD  R6, R13

paheadComma:
	CBZ   R12, paheadNone
	CLZ   R12, R14
	ADD   R14, R13, R13
	ADD   $1, R13, R13      // the lane after this comma
	LSL   R14, R12, R12
	LSL   $1, R12, R12
	SUBS  $1, R21, R21
	BNE   paheadComma
	ADD   R2, R13, R26
	CMP   R1, R26
	BGT   paheadNone
	ADD   R0, R26, R12
	FRPREP1(R12)
	WORD  $0x6e25358c // cmhi v12.16b, v12.16b, v5.16b
	WORD  $0x6e2535ad // cmhi v13.16b, v13.16b, v5.16b
	WORD  $0x6e2535ce // cmhi v14.16b, v14.16b, v5.16b
	WORD  $0x6e2535ef // cmhi v15.16b, v15.16b, v5.16b
	FRPREP2V
	B     paheadDone

paheadNone:
	MOVD  $-1, R26

paheadDone:
	MOVD  R5, R21

pnum:
	// A number whose region starts at lane b.
	CMP   $64, R6
	BHS   pfar
	SUBS  $1, R21, R21
	BEQ   plast
	LSL   R6, R8, R12       // commas << b
	CBZ   R12, pfar
	CLZ   R12, R13
	ADD   R6, R13, R13      // c: an inner number's comma
	MOVBU (R11)(R6), R12
	CMP   $0x2d, R12
	CSET  EQ, R15
	ADD   R15, R6, R16      // s
	LSL   R16, R7, R12
	CLZ   R12, R17          // L1
	CBZ   R17, pinWS

pinDigits:
	ADD   R16, R17, R12     // e1
	MOVBU (R11)(R12), R14
	CMP   $0x2e, R14
	BNE   pinNoFrac
	ADD   $1, R12, R12
	LSL   R12, R7, R14
	CLZ   R14, R14          // L2
	CBZ   R14, pback
	ADD   R14, R12, R12     // e
	CMP   R13, R12
	BNE   pinWSComma

pinConv:
	ADD   R17, R14, R20     // L
	CMP   $15, R20
	BHI   pinLong
	SHORTCONV
	ADD   $1, R13, R6
	B     pnum

pinLong:
	CMP   $19, R20
	BHI   pback
	LONGCONV(pback, pinEl, pinNoRefine, pinRound, pinRefine, pinDone, R25)
	ADD   $1, R13, R6
	B     pnum

pinNoFrac:
	MOVD  ZR, R14
	CMP   R13, R12
	BEQ   pinConv

pinWSComma:
	// e < c: whitespace up to the comma, or the number is not this point's
	// (a ']' there: the point has fewer numbers than n).
	CBNZ  R19, pinWSCommaAt
	MOVD  $1, R20
	B     pwsCompute

pinWSCommaAt:
	LSL   R12, R9, R20
	CLZ   R20, R20
	ADD   R12, R20, R20
	CMP   R13, R20
	BEQ   pinConv
	B     pback

pinWS:
	// The byte at s is not a digit: whitespace before the number (after a
	// '-', not a number at all).
	CBNZ  R15, pback
	CBNZ  R19, pinWSAt
	MOVD  $2, R20
	B     pwsCompute

pinWSAt:
	LSL   R6, R9, R12
	CLZ   R12, R12
	ADD   R6, R12, R16      // c at most: the comma is not whitespace
	MOVBU (R11)(R16), R12
	CMP   $0x2d, R12
	CSET  EQ, R15
	ADD   R15, R16, R16
	LSL   R16, R7, R12
	CLZ   R12, R17
	CBZ   R17, pback
	B     pinDigits

plast:
	// The point's last number, delimited by the point's ']'.
	MOVBU (R11)(R6), R12
	CMP   $0x2d, R12
	CSET  EQ, R15
	ADD   R15, R6, R16
	CMP   $64, R16
	BHS   pfar
	LSL   R16, R7, R12
	CLZ   R12, R17
	CBZ   R17, plastWS

plastDigits:
	ADD   R16, R17, R12     // e1
	CMP   $64, R12
	BHS   pfar
	MOVBU (R11)(R12), R14
	CMP   $0x2e, R14
	BNE   plastNoFrac
	ADD   $1, R12, R12
	CMP   $64, R12
	BHS   pfar
	LSL   R12, R7, R14
	CLZ   R14, R14          // L2
	CBZ   R14, pback
	ADD   R14, R12, R12     // e
	CMP   $64, R12
	BHS   pfar
	B     plastTerm

plastNoFrac:
	MOVD  ZR, R14

plastTerm:
	MOVBU (R11)(R12), R13
	CMP   $0x5d, R13
	BNE   plastTermWS
	MOVD  R12, R13          // r: the point's ']', right after the number

plastConv:
	ADD   R17, R14, R20
	CMP   $15, R20
	BHI   plastLong
	SHORTCONV
	B     psep

plastLong:
	CMP   $19, R20
	BHI   pback
	LONGCONV(pback, plastEl, plastNoRefine, plastRound, plastRefine, psep, R25)

	// The point is written; the first byte after its ']' that is not
	// whitespace must be the ',' before the next point or the ring's ']'.
	ADD   $1, R13, R12
	CMP   $64, R12
	BHS   psepNext
	MOVBU (R11)(R12), R20
	CMP   $0x2c, R20
	BEQ   pcomma
	CMP   $0x5d, R20
	BEQ   pclosed
	CMP   $0x20, R20
	BHI   pback             // neither: not a ring of points
	CBNZ  R19, psepAt
	MOVD  $3, R20
	B     pwsCompute

psepAt:
	LSL   R12, R9, R20
	CLZ   R20, R20
	ADD   R20, R12, R12     // the first byte after the ']' that is not whitespace
	CMP   $64, R12
	BHS   psepNext
	MOVBU (R11)(R12), R20
	CMP   $0x2c, R20
	BEQ   pcomma
	CMP   $0x5d, R20
	BEQ   pclosed
	B     pback

pcomma:
	ADD   R12, R2, R2
	ADD   $1, R2, R2        // the next point's region starts after the ','
	B     ppoint

pclosed:
	ADD   R12, R2, R2       // p: the ring's ']'
	MOVD  out_base+32(FP), R12
	SUB   R12, R3, R12
	LSR   $3, R12, R12
	UDIV  R5, R12, R12
	MOVD  R12, np+64(FP)
	MOVD  R2, p+72(FP)
	MOVD  $1, R12
	MOVD  R12, closed+80(FP)
	RET

psepNext:
	// The rest of the window is whitespace — in a pretty-printed ring the
	// last point's ']' is followed by a line break and the ring's own
	// indentation — so the separator is looked for in the windows after it,
	// while they hold nothing else. The point stays uncounted until then.
	ADD   $64, R2, R2

psepLoop:
	ADD   $16, R1, R12      // a window read whole: 64 bytes, not 80
	CMP   R12, R2
	BGT   pback
	ADD   R0, R2, R11
	VLD1  (R11), [V0.B16, V1.B16, V2.B16, V3.B16]
	WORD  $0x6e27340c // cmhi v12.16b, v0.16b, v7.16b
	WORD  $0x6e27342d // cmhi v13.16b, v1.16b, v7.16b
	WORD  $0x6e27344e // cmhi v14.16b, v2.16b, v7.16b
	WORD  $0x6e27346f // cmhi v15.16b, v3.16b, v7.16b
	WSPACK(V12.B16, V13.B16, V14.B16, V15.B16, V12, R20)
	CBNZ  R20, psepFound
	ADD   $64, R2, R2
	B     psepLoop

psepFound:
	CLZ   R20, R12
	MOVBU (R11)(R12), R20
	CMP   $0x2c, R20
	BEQ   pcomma
	CMP   $0x5d, R20
	BEQ   pclosed
	B     pback

plastTermWS:
	// The byte after the last number is not the ']': whitespace before it,
	// or the point has more numbers than n.
	CMP   $0x20, R13
	BHI   pback
	CBNZ  R19, plastTermAt
	MOVD  $4, R20
	B     pwsCompute

plastTermAt:
	LSL   R12, R9, R13
	CLZ   R13, R13
	ADD   R12, R13, R13     // the first byte after e that is not whitespace
	CMP   $64, R13
	BHS   pfar
	MOVBU (R11)(R13), R20
	CMP   $0x5d, R20
	BNE   pback
	B     plastConv

plastWS:
	CBNZ  R15, pback
	CBNZ  R19, plastWSAt
	MOVD  $5, R20
	B     pwsCompute

plastWSAt:
	LSL   R6, R9, R12
	CLZ   R12, R12
	ADD   R6, R12, R16
	CMP   $64, R16
	BHS   pfar
	MOVBU (R11)(R16), R12
	CMP   $0x2d, R12
	CSET  EQ, R15
	ADD   R15, R16, R16
	CMP   $64, R16
	BHS   pfar
	LSL   R16, R7, R12
	CLZ   R12, R17
	CBZ   R17, pback
	B     plastDigits

plead:
	// The window does not start at a '[': whitespace before it (a
	// pretty-printed ring), or this is not a point.
	CMP   $0x20, R12
	BHI   pstop
	MOVD  $6, R20
	B     pwsCompute

pleadAt:
	CLZ   R9, R12           // the first byte that is not whitespace
	CMP   $16, R12
	BHS   prewindow         // far into the window: start one at it
	MOVBU (R11)(R12), R20
	CMP   $0x5b, R20
	BNE   pstop
	MOVD  R12, R24
	ADD   $1, R12, R6
	B     pfirst

prewindow:
	// Sixteen or more bytes of whitespace before the point (64: all of the
	// window): start the next window at the first byte that is not.
	ADD   R12, R2, R2
	B     ppoint

pfar:
	// The point runs past its window: when whitespace before its '[' took
	// some of the window, start one at the '[' and try once more (at lane 0
	// there is nothing left to gain, and the point is handed back).
	CBZ   R24, pback
	MOVD  R22, R3
	ADD   R24, R2, R2
	B     ppoint

pwsCompute:
	// The whitespace class, not whitespace (c > 0x20), for the window at
	// R11, into R9 with R19 set: computed only when the window first needs
	// it, from a reload. R20 says which site asked.
	VLD1  (R11), [V12.B16, V13.B16, V14.B16, V15.B16]
	WORD  $0x6e273590 // cmhi v16.16b, v12.16b, v7.16b
	WORD  $0x6e2735b1 // cmhi v17.16b, v13.16b, v7.16b
	WORD  $0x6e2735d2 // cmhi v18.16b, v14.16b, v7.16b
	WORD  $0x6e2735f3 // cmhi v19.16b, v15.16b, v7.16b
	WSPACK(V16.B16, V17.B16, V18.B16, V19.B16, V16, R9)
	MOVD  $1, R19
	CMP   $2, R20
	BLO   pinWSCommaAt
	BEQ   pinWSAt
	CMP   $4, R20
	BLO   psepAt
	BEQ   plastTermAt
	CMP   $5, R20
	BEQ   plastWSAt
	B     pleadAt

pback:
	// The point cannot be taken here: forget its numbers already written and
	// hand it back from its region's start.
	MOVD  R22, R3

pstop:
	MOVD  out_base+32(FP), R12
	SUB   R12, R3, R12
	LSR   $3, R12, R12
	UDIV  R5, R12, R12
	MOVD  R12, np+64(FP)
	MOVD  R23, p+72(FP)
	MOVD  ZR, closed+80(FP)
	RET

// ---- The same walks as checks, for SkipValueStrict -----------------------------

// func validNumberRunNEON(data []byte, i int) (p, closed int)
//
// parseFloatRunNEON's walk as a check, for SkipValueStrict — the contract of
// the amd64 validNumberRun: it passes over "ws* -? digits (. digits)? ws* ','"
// elements from data[i:], and over the array's last element when a ']' ends
// it (closed = 1, p at the ']'), converting nothing. Otherwise p is where the
// scalar walk resumes: the first byte that is not whitespace of the first
// element it did not take (its '-' included), or the start of that element's
// region, or just past the last comma it passed. It takes only numbers the
// decoder's own reader accepts — plain decimals, which cannot overflow at
// under 64 digits, and a number delimited inside a 64-byte block has fewer —
// and leaves everything else (an exponent, a '+', a bare '.', a leading '.',
// a string, a container, a stray byte) to the scalar walk, which decides
// exactly as before. It reads whole 64-byte blocks only, so it stops when
// fewer than 64 bytes remain.
//
// The block walk is parseFloatRunNEON's, label for label, without the output:
// no capacity bound, no digit limit, and where that walk converts, this one
// takes the next element.
TEXT ·validNumberRunNEON(SB), NOSPLIT, $0-48
	MOVD  data_base+0(FP), R0
	MOVD  data_len+8(FP), R1
	SUB   $64, R1           // the last block start with 64 bytes
	MOVD  i+24(FP), R2
	MOVD  $frConst<>(SB), R12
	ADD   $32, R12
	VLD1  (R12), [V10.B16]  // the cascade's bit weights
	VMOVI $0x30, V4.B16 // '0'
	VMOVI $9, V5.B16
	VMOVI $0x2c, V6.B16 // ','
	VMOVI $0x20, V7.B16
	MOVD  ZR, R19
	MOVD  ZR, R6

vblock:
	CMP   R1, R2
	BGT   vdone
	ADD   R0, R2, R11
	FRPREP1(R11)
	WORD  $0x6e25358c // cmhi v12.16b, v12.16b, v5.16b
	WORD  $0x6e2535ad // cmhi v13.16b, v13.16b, v5.16b
	WORD  $0x6e2535ce // cmhi v14.16b, v14.16b, v5.16b
	WORD  $0x6e2535ef // cmhi v15.16b, v15.16b, v5.16b
	FRPREP2(R7, R8)
	MOVD  R19, R22
	MOVD  ZR, R9
	MOVD  ZR, R19

vblockAt:
	ADD   $48, R2, R12
	CMP   R1, R12
	BGT   vnoNext
	ADD   $48, R11, R26
	FRPREP1(R26)
	WORD  $0x6e25358c // cmhi v12.16b, v12.16b, v5.16b
	WORD  $0x6e2535ad // cmhi v13.16b, v13.16b, v5.16b
	WORD  $0x6e2535ce // cmhi v14.16b, v14.16b, v5.16b
	WORD  $0x6e2535ef // cmhi v15.16b, v15.16b, v5.16b
	FRPREP2(R23, R24)
	MOVD  ZR, R25
	CBZ   R22, vdispatch
	WORD  $0x6e27340c // cmhi v12.16b, v0.16b, v7.16b
	WORD  $0x6e27342d // cmhi v13.16b, v1.16b, v7.16b
	WORD  $0x6e27344e // cmhi v14.16b, v2.16b, v7.16b
	WORD  $0x6e27346f // cmhi v15.16b, v3.16b, v7.16b
	WSPACK(V12.B16, V13.B16, V14.B16, V15.B16, V12, R25)
	B     vdispatch

vnoNext:
	MOVD  ZR, R26

vdispatch:
	CBNZ  R19, vwelem

velem:
	CMP   $48, R6
	BHS   vstep
	LSL   R6, R8, R12       // commas << b
	CBZ   R12, vnoComma
	CLZ   R12, R13
	ADD   R6, R13, R13      // c
	MOVBU (R11)(R6), R12
	CMP   $0x2d, R12
	CSET  EQ, R15
	ADD   R15, R6, R16      // s
	LSL   R16, R7, R12
	CLZ   R12, R17          // L1
	CBZ   R17, velemWS

vdigits:
	ADD   R16, R17, R12     // e1
	MOVBU (R11)(R12), R14
	CMP   $0x2e, R14
	BNE   vnoFrac
	ADD   $1, R12, R12
	LSL   R12, R7, R14
	CLZ   R14, R14          // L2
	CBZ   R14, vstopElem
	ADD   R14, R12, R12     // e
	CMP   R13, R12
	BNE   vwsBeforeComma

vnext:
	ADD   $1, R13, R6       // b: just past the comma
	CBZ   R19, velem
	B     vwelem

vnoFrac:
	CMP   R13, R12
	BEQ   vnext

vwsBeforeComma:
	MOVBU (R11)(R12), R20
	CMP   $0x5d, R20
	BEQ   vcloseAtE
	CBNZ  R19, vwsBeforeCommaAt
	MOVD  $2, R20
	B     vwsCompute

vwsBeforeCommaAt:
	LSL   R12, R9, R20
	CLZ   R20, R20
	ADD   R12, R20, R20     // the first byte after e that is not whitespace
	CMP   R13, R20
	BEQ   vnext
	MOVBU (R11)(R20), R13
	CMP   $0x5d, R13
	BNE   vstopElem
	MOVD  R20, R13          // the ']' ends this, the array's last element
	B     vclosed

vcloseAtE:
	MOVD  R12, R13
	B     vclosed

vwelem:
	CMP   $48, R6
	BHS   vstep
	LSL   R6, R8, R12
	CBZ   R12, vnoComma
	CLZ   R12, R13
	ADD   R6, R13, R13

vwelemAt:
	LSL   R6, R9, R12
	CLZ   R12, R12
	ADD   R6, R12, R16
	MOVBU (R11)(R16), R12
	CMP   $0x2d, R12
	CSET  EQ, R15
	ADD   R15, R16, R16
	LSL   R16, R7, R12
	CLZ   R12, R17
	CBNZ  R17, vdigits
	B     vstopElem

velemWS:
	CBNZ  R15, vstopElem
	MOVD  $1, R20
	B     vwsCompute

vnoComma:
	MOVBU (R11)(R6), R12
	CMP   $0x2d, R12
	CSET  EQ, R15
	ADD   R15, R6, R16
	LSL   R16, R7, R12
	CLZ   R12, R17
	CBZ   R17, vlastWS

vlastDigits:
	ADD   R16, R17, R12     // e1
	CMP   $64, R12
	BHS   vstraddle
	MOVBU (R11)(R12), R14
	CMP   $0x2e, R14
	BNE   vlastTerm
	ADD   $1, R12, R12
	CMP   $64, R12
	BHS   vstraddle
	LSL   R12, R7, R14
	CLZ   R14, R14
	CBZ   R14, vstopElem
	ADD   R14, R12, R12     // e
	CMP   $64, R12
	BHS   vstraddle

vlastTerm:
	MOVBU (R11)(R12), R13
	CMP   $0x5d, R13
	BNE   vlastTermWS
	MOVD  R12, R13
	B     vclosed

vlastTermWS:
	CBNZ  R19, vlastTermAt
	MOVD  $4, R20
	B     vwsCompute

vlastTermAt:
	LSL   R12, R9, R13
	CLZ   R13, R13
	ADD   R12, R13, R13
	CMP   $64, R13
	BHS   vstraddle
	MOVBU (R11)(R13), R20
	CMP   $0x5d, R20
	BNE   vstopElem
	B     vclosed

vlastWS:
	CBNZ  R15, vstopElem
	CBNZ  R19, vlastWSAt
	MOVD  $3, R20
	B     vwsCompute

vlastWSAt:
	LSL   R6, R9, R12
	CLZ   R12, R12
	ADD   R6, R12, R16
	CMP   $64, R16
	BHS   vstraddle
	MOVBU (R11)(R16), R12
	CMP   $0x2d, R12
	CSET  EQ, R15
	ADD   R15, R16, R16
	CMP   $64, R16
	BHS   vstraddle
	LSL   R16, R7, R12
	CLZ   R12, R17
	CBZ   R17, vstopElem
	B     vlastDigits

vwsCompute:
	VLD1  (R11), [V12.B16, V13.B16, V14.B16, V15.B16]
	WORD  $0x6e273590 // cmhi v16.16b, v12.16b, v7.16b
	WORD  $0x6e2735b1 // cmhi v17.16b, v13.16b, v7.16b
	WORD  $0x6e2735d2 // cmhi v18.16b, v14.16b, v7.16b
	WORD  $0x6e2735f3 // cmhi v19.16b, v15.16b, v7.16b
	WSPACK(V16.B16, V17.B16, V18.B16, V19.B16, V16, R9)
	MOVD  $1, R19
	CMP   $2, R20
	BLO   vwelemAt
	BEQ   vwsBeforeCommaAt
	CMP   $3, R20
	BEQ   vlastWSAt
	B     vlastTermAt

vstep:
	CBZ   R26, vdone
	ADD   $48, R2
	SUB   $48, R6
	MOVD  R26, R11
	MOVD  R23, R7
	MOVD  R24, R8
	MOVD  R25, R9
	MOVD  R22, R12
	MOVD  R19, R22
	MOVD  R12, R19
	B     vblockAt

vstraddle:
	SUB   R15, R16, R12
	CMP   $64, R12
	BLS   vstraddleAt
	MOVD  $64, R12

vstraddleAt:
	CBZ   R12, vstopElem
	ADD   R12, R2, R2
	MOVD  ZR, R6
	B     vblock

vstopElem:
	SUB   R15, R16, R6

vdone:
	ADD   R6, R2, R2
	MOVD  R2, p+32(FP)
	MOVD  ZR, closed+40(FP)
	RET

vclosed:
	ADD   R13, R2, R2
	MOVD  R2, p+32(FP)
	MOVD  $1, R12
	MOVD  R12, closed+40(FP)
	RET

// func validPointsRunNEON(data []byte, i int) (p, closed int)
//
// The points walk (parseFloatPointsNEON) as a check, for SkipValueStrict — the
// contract of the amd64 validPointsRun: i is at a point's '[' in an array of
// flat numeric arrays — a coordinate ring — and it passes over as many points
// as it can, each "[" number ("," number)* "]" with any whitespace, the
// numbers as validNumberRunNEON takes them, returning closed = 1 with p at the
// ring's ']' when it reaches it. Otherwise p is the start of the first point
// it did not take (the region after the previous point's ','), which the
// scalar walk then checks whole — an empty point, an exponent, a nested array,
// a string: anything but plain decimals. It takes points of any length that
// fit a window, since any count of numbers is valid, and so it does not walk
// commas: each number's delimiter is the byte after it (or after whitespace),
// a ',' going on to the next number and a ']' ending the point. It does
// classify commas, for the look-ahead: the next window is classified as soon
// as this one's commas are known, where the previous point's count of numbers
// says it starts (see qfirst). One window per point, at the point, and the
// separator after it looked for across windows of whitespace, as in the
// conversion walk; windows are read whole, so it stops when fewer than 64
// bytes remain. The caller runs it only below MaxDepth, the points being one
// level deeper than the ring.
TEXT ·validPointsRunNEON(SB), NOSPLIT, $0-48
	MOVD  data_base+0(FP), R0
	MOVD  data_len+8(FP), R1
	SUB   $64, R1
	MOVD  i+24(FP), R2
	MOVD  $frConst<>(SB), R12
	ADD   $32, R12
	VLD1  (R12), [V10.B16]
	VMOVI $0x30, V4.B16
	VMOVI $9, V5.B16
	VMOVI $0x2c, V6.B16     // ',': the comma class is the look-ahead's
	VMOVI $0x20, V7.B16
	MOVD  $-1, R26          // no window classified ahead
	MOVD  ZR, R22           // the previous point's count: none yet

qpoint:
	MOVD  R2, R23           // the hand-back position
	CMP   R1, R2
	BGT   qstop
	ADD   R0, R2, R11
	CMP   R26, R2
	BEQ   qahead
	FRPREP1(R11)
	WORD  $0x6e25358c // cmhi v12.16b, v12.16b, v5.16b
	WORD  $0x6e2535ad // cmhi v13.16b, v13.16b, v5.16b
	WORD  $0x6e2535ce // cmhi v14.16b, v14.16b, v5.16b
	WORD  $0x6e2535ef // cmhi v15.16b, v15.16b, v5.16b
	FRPREP2(R7, R8)
	B     qwindow

qahead:
	VMOV  V28.D[1], R7
	VMOV  V28.D[0], R8

qwindow:
	MOVD  ZR, R19
	MOVBU (R11), R12
	CMP   $0x5b, R12
	BNE   qlead
	MOVD  ZR, R24
	MOVD  $1, R6

qfirst:
	// Classify the next point's window now, where the previous point's
	// count of numbers says it starts — the byte after that many commas from
	// b: a ring's points are alike, and the walk's chain is window to window
	// (classification, then every number's delimiter, then the separator),
	// which left alone ran it at 52 cycles a canada point with the core
	// mostly idle. A guess that is wrong only means the next window is
	// classified when it is reached.
	MOVD  $-1, R26
	CBZ   R22, qfirstDone
	LSL   R6, R8, R12       // commas << b
	MOVD  R22, R21
	MOVD  R6, R13

qaheadComma:
	CBZ   R12, qfirstDone
	CLZ   R12, R14
	ADD   R14, R13, R13
	ADD   $1, R13, R13      // the lane after this comma
	LSL   R14, R12, R12
	LSL   $1, R12, R12
	SUBS  $1, R21, R21
	BNE   qaheadComma
	ADD   R2, R13, R26
	CMP   R1, R26
	BGT   qaheadNone
	ADD   R0, R26, R12
	FRPREP1(R12)
	WORD  $0x6e25358c // cmhi v12.16b, v12.16b, v5.16b
	WORD  $0x6e2535ad // cmhi v13.16b, v13.16b, v5.16b
	WORD  $0x6e2535ce // cmhi v14.16b, v14.16b, v5.16b
	WORD  $0x6e2535ef // cmhi v15.16b, v15.16b, v5.16b
	FRPREP2V
	B     qfirstDone

qaheadNone:
	MOVD  $-1, R26

qfirstDone:
	MOVD  ZR, R21           // the numbers in this point so far

qnum:
	CMP   $64, R6
	BHS   qfar
	MOVBU (R11)(R6), R12
	CMP   $0x2d, R12
	CSET  EQ, R15
	ADD   R15, R6, R16
	CMP   $64, R16
	BHS   qfar
	LSL   R16, R7, R12
	CLZ   R12, R17
	CBZ   R17, qnumWS

qdigits:
	ADD   R16, R17, R12     // e1
	CMP   $64, R12
	BHS   qfar
	MOVBU (R11)(R12), R14
	CMP   $0x2e, R14
	BNE   qdelim
	ADD   $1, R12, R12
	CMP   $64, R12
	BHS   qfar
	LSL   R12, R7, R14
	CLZ   R14, R14
	CBZ   R14, qback
	ADD   R14, R12, R12     // e
	CMP   $64, R12
	BHS   qfar
	MOVBU (R11)(R12), R14

qdelim:
	// R14 is the byte at e (R12): a ',' goes on to the next number, a ']'
	// ends the point, whitespace comes before either.
	CMP   $0x2c, R14
	BEQ   qcomma
	CMP   $0x5d, R14
	BEQ   qclose
	CMP   $0x20, R14
	BHI   qback
	CBNZ  R19, qdelimAt
	MOVD  $1, R20
	B     qwsCompute

qdelimAt:
	LSL   R12, R9, R14
	CLZ   R14, R14
	ADD   R14, R12, R12     // the first byte after e that is not whitespace
	CMP   $64, R12
	BHS   qfar
	MOVBU (R11)(R12), R14
	CMP   $0x2c, R14
	BEQ   qcomma
	CMP   $0x5d, R14
	BEQ   qclose
	B     qback

qcomma:
	ADD   $1, R21, R21
	ADD   $1, R12, R6
	B     qnum

qnumWS:
	CBNZ  R15, qback
	CBNZ  R19, qnumWSAt
	MOVD  $2, R20
	B     qwsCompute

qnumWSAt:
	LSL   R6, R9, R12
	CLZ   R12, R12
	ADD   R6, R12, R16
	CMP   $64, R16
	BHS   qfar
	MOVBU (R11)(R16), R12
	CMP   $0x2d, R12
	CSET  EQ, R15
	ADD   R15, R16, R16
	CMP   $64, R16
	BHS   qfar
	LSL   R16, R7, R12
	CLZ   R12, R17
	CBZ   R17, qback
	B     qdigits

qclose:
	// The point's ']' at R12: the first byte after it that is not whitespace
	// must be the ',' before the next point or the ring's ']'.
	ADD   $1, R21, R22      // its count of numbers, the next point's guess
	ADD   $1, R12, R12
	CMP   $64, R12
	BHS   qsepNext
	MOVBU (R11)(R12), R20
	CMP   $0x2c, R20
	BEQ   qsepComma
	CMP   $0x5d, R20
	BEQ   qclosed
	CMP   $0x20, R20
	BHI   qback
	CBNZ  R19, qsepAt
	MOVD  $3, R20
	B     qwsCompute

qsepAt:
	LSL   R12, R9, R20
	CLZ   R20, R20
	ADD   R20, R12, R12
	CMP   $64, R12
	BHS   qsepNext
	MOVBU (R11)(R12), R20
	CMP   $0x2c, R20
	BEQ   qsepComma
	CMP   $0x5d, R20
	BEQ   qclosed
	B     qback

qsepComma:
	ADD   R12, R2, R2
	ADD   $1, R2, R2
	B     qpoint

qclosed:
	ADD   R12, R2, R2
	MOVD  R2, p+32(FP)
	MOVD  $1, R12
	MOVD  R12, closed+40(FP)
	RET

qsepNext:
	// The rest of the window is whitespace: look for the separator in the
	// windows after it while they hold nothing else.
	ADD   $64, R2, R2

qsepLoop:
	CMP   R1, R2
	BGT   qback
	ADD   R0, R2, R11
	VLD1  (R11), [V0.B16, V1.B16, V2.B16, V3.B16]
	WORD  $0x6e27340c // cmhi v12.16b, v0.16b, v7.16b
	WORD  $0x6e27342d // cmhi v13.16b, v1.16b, v7.16b
	WORD  $0x6e27344e // cmhi v14.16b, v2.16b, v7.16b
	WORD  $0x6e27346f // cmhi v15.16b, v3.16b, v7.16b
	WSPACK(V12.B16, V13.B16, V14.B16, V15.B16, V12, R20)
	CBNZ  R20, qsepFound
	ADD   $64, R2, R2
	B     qsepLoop

qsepFound:
	CLZ   R20, R12
	MOVBU (R11)(R12), R20
	CMP   $0x2c, R20
	BEQ   qsepComma
	CMP   $0x5d, R20
	BEQ   qclosed
	B     qback

qlead:
	CMP   $0x20, R12
	BHI   qstop
	MOVD  $4, R20
	B     qwsCompute

qleadAt:
	CLZ   R9, R12
	CMP   $16, R12
	BHS   qrewindow
	MOVBU (R11)(R12), R20
	CMP   $0x5b, R20
	BNE   qstop
	MOVD  R12, R24
	ADD   $1, R12, R6
	B     qfirst

qrewindow:
	ADD   R12, R2, R2
	B     qpoint

qfar:
	// The point runs past its window: when whitespace before its '[' took
	// some of it, start one at the '[' and try once more.
	CBZ   R24, qback
	ADD   R24, R2, R2
	B     qpoint

qwsCompute:
	VLD1  (R11), [V12.B16, V13.B16, V14.B16, V15.B16]
	WORD  $0x6e273590 // cmhi v16.16b, v12.16b, v7.16b
	WORD  $0x6e2735b1 // cmhi v17.16b, v13.16b, v7.16b
	WORD  $0x6e2735d2 // cmhi v18.16b, v14.16b, v7.16b
	WORD  $0x6e2735f3 // cmhi v19.16b, v15.16b, v7.16b
	WSPACK(V16.B16, V17.B16, V18.B16, V19.B16, V16, R9)
	MOVD  $1, R19
	CMP   $2, R20
	BLO   qdelimAt
	BEQ   qnumWSAt
	CMP   $3, R20
	BEQ   qsepAt
	B     qleadAt

qback:
qstop:
	MOVD  R23, p+32(FP)
	MOVD  ZR, closed+40(FP)
	RET
