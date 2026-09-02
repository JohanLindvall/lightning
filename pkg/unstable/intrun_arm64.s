#include "textflag.h"

// Constants for parseIntRunNEON, loaded as one three-vector VLD1: the byte
// weights of the UDOT that folds the shuffled digits into three groups
// (100,10,1 for the first two, 10,1 for the last), the word weights 10^5,
// 10^2, 1 that turn the groups into the eight-digit value, and the bit
// weights the byte-class compares are ANDed with so an ADDP cascade sums
// them into a bitmask (see skipfast_arm64.s) — in the order {128,...,1},
// which with a byte reversal of the packed result gives the bit-reversed
// mask the walk wants (see PREP).
DATA irConst<>+0(SB)/8, $0x00010a6400010a64
DATA irConst<>+8(SB)/8, $0x000000000000010a
DATA irConst<>+16(SB)/8, $0x00000064000186a0
DATA irConst<>+24(SB)/8, $0x0000000000000001
DATA irConst<>+32(SB)/8, $0x0102040810204080
DATA irConst<>+40(SB)/8, $0x0102040810204080
GLOBL irConst<>(SB), RODATA|NOPTR, $48

// The assembler floor is the module's go directive — CI assembles with Go
// 1.25 — and its arm64 assembler has no CMHI or MUL (vector) mnemonic, nor
// UDOT; those are WORDs, each on its own line with the mnemonic sveasm
// derives the encoding from (a macro line cannot carry that comment), which
// is why the compares below sit between the macros rather than inside them.

// CASCADE folds the four class compares in t0..t3 (0x00/0xFF per lane) to
// the 16-byte form two ADDP steps short of the block's 64-bit mask, in t0 —
// the bit-weight/ADDP cascade of skipfast_arm64.s. Clobbers t1..t3.
#define CASCADE(t0, t1, t2, t3) \
	VAND  V10.B16, t0, t0  \
	VAND  V10.B16, t1, t1  \
	VAND  V10.B16, t2, t2  \
	VAND  V10.B16, t3, t3  \
	VADDP t1, t0, t0       \
	VADDP t3, t2, t2       \
	VADDP t2, t0, t0

// CLASS3 is CASCADE with the compare of the four chunks c0..c3 against the
// splat s in front, for a compare the assembler can spell (VCMEQ).
#define CLASS3(cmp, s, c0, c1, c2, c3, t0, t1, t2, t3) \
	cmp   s, c0, t0        \
	cmp   s, c1, t1        \
	cmp   s, c2, t2        \
	cmp   s, c3, t3        \
	CASCADE(t0, t1, t2, t3)

// PREP1 and PREP2, with the four not-digit compares between them, classify
// the 64-byte block at address a: its bytes minus '0' land in d0..d3 (the
// fold's TBL source) and two masks, bit-reversed so lane j is bit 63-j, in
// rD (not a digit: c-'0' > 9, the CMHI WORDs against the 9 splat V5, into
// V12..V15) and rC (a comma). The digit class is compared negated because
// the walk wants its complement. Each class is the cascade movemask, and
// the reversal costs nothing on the way: the weights run {128,...,1}
// within each eight-lane half, so the cascade packs each half reversed,
// and a VREV64 of the packed bytes then reverses the halves and chunks
// (lane 16k+8h+l lands at bit 63-(16k+8h+l)). The two classes share their
// final ADDP, so one VREV64 and one VMOV pair serve both. Clobbers V0-V3,
// V12-V19.
#define PREP1(a, d0, d1, d2, d3) \
	VLD1  (a), [V0.B16, V1.B16, V2.B16, V3.B16] \
	VSUB  V4.B16, V0.B16, d0                    \
	VSUB  V4.B16, V1.B16, d1                    \
	VSUB  V4.B16, V2.B16, d2                    \
	VSUB  V4.B16, V3.B16, d3

#define PREP2(rD, rC) \
	CASCADE(V12.B16, V13.B16, V14.B16, V15.B16)                                              \
	CLASS3(VCMEQ, V6.B16, V0.B16, V1.B16, V2.B16, V3.B16, V16.B16, V17.B16, V18.B16, V19.B16) \
	VADDP  V12.B16, V16.B16, V16.B16                                                          \
	VREV64 V16.B16, V16.B16                                                                   \
	VMOV   V16.D[1], rD                                                                       \
	VMOV   V16.D[0], rC

// WSPACK finishes the third mask, not whitespace (c > 0x20), from its four
// compares in t0..t3 (CMHI WORDs against the 0x20 splat V7), into rW; v is
// t0's bare register name, for the lane move.
#define WSPACK(t0, t1, t2, t3, v, rW) \
	CASCADE(t0, t1, t2, t3) \
	VADDP  t0, t0, t0       \
	VREV64 t0, t0           \
	VMOV   v.D[0], rW

// FOLD1 and FOLD2, with the UDOT and MUL between them, turn the L digits at
// block offset s (R6 = s, R12 = L; the block's bytes minus '0' in V20..V23)
// into an int64 stored at the out cursor R3, which they advance. The TBL
// control is intRunShuffleNEON[s*8+L]: the digits of the eight-digit
// zero-padded number laid out three, three and two to a word (0x80
// elsewhere, which a four-register TBL reads as zero, any index >= 64 being
// out of range), so one UDOT with byte weights 100,10,1 / 100,10,1 / 10,1
// makes the three groups, and a word multiply by 10^5, 10^2, 1 with an
// across-vector add makes the value — a byte dot product cannot weight a
// fourth digit (1000 does not fit a byte), which is why the groups are
// three and not four, and why the control is sixteen bytes rather than
// eight. UDOT accumulates, so the value register is zeroed first. Clobbers
// R16, V29-V31.
#define FOLD1 \
	ADD    R6<<3, R12, R16                                        \
	ADD    R16<<4, R10, R16                                       \
	VLD1   (R16), [V30.B16]                                       \
	VTBL   V30.B16, [V20.B16, V21.B16, V22.B16, V23.B16], V29.B16 \
	VMOVI  $0, V31.B16

#define FOLD2 \
	VUADDLV V31.S4, V31 \
	FMOVD.P F31, 8(R3)

// func parseIntRunNEON(data []byte, i int, out []int64) (n, p, closed int)
//
// The arm64 twin of parseIntRunSSE (intrun_amd64.s): the same element walk,
// stop positions and contract, over a different block shape. What the
// counters said about the direct port of the 16-byte block (CLAUDE.md,
// 2026-09-02): the next block's address depended on where the walk left the
// previous one, so the whole load, compare, narrow, VMOV prologue sat on the
// element chain (two elements a block ran 16.7 cycles each straddling and
// 10.9 not); at 40 instructions an element the walk was issue-bound; and
// even with the address dependence gone, a ~25-cycle classification chain
// was exposed at every block start, because the walk executes at dispatch
// rate and so nothing older is pending for the chain to overlap with.
//
// So the block is 64 bytes, walked at a fixed 48-byte stride: elements
// starting in the first 48 lanes are consumed here, one starting at lane 48
// or beyond is deferred to the next block, which is entered at lane b-48
// (the consumed element's tail can reach past lane 48; the comma test's
// shift to each element's end is what skips what lies before it), and the
// next block's address is s+48 whatever the walk found. And the
// classification is software-pipelined: block k+1 is
// loaded and classified into a second register set at the top of block k's
// walk, so its chain runs under the walk and the transition is a handful
// of moves. Only an element that starts below lane 48 and still runs past
// lane 63 — a whitespace run or digit run of 16+ bytes — restarts the next
// block at its own start (the straddle case), which is the one data-
// dependent block address left and the one place the pipelined block is
// discarded. Classes are the bit-weight/ADDP-cascade movemask of
// skipfast_arm64.s, two of them packed by one ADDP into a single VMOV; the
// masks are kept BIT-REVERSED (free, see PREP), so every run length the
// walk needs is one LSL to the run's start plus one CLZ, and the comma
// that ends an element is one bit of the comma mask shifted to the digits'
// end — the mask is never consumed, so nothing but b itself carries from
// one element to the next. The digit and whitespace classes are compared
// negated (CMHI against 9 and 0x20), so the complement the counts want
// costs nothing, and the whitespace class is computed only for a block
// whose walk meets a non-digit — on demand the first time (see wsCompute),
// pipelined with the rest once a block has needed it. The out-capacity
// test is per block, not per element. A block stores at most 25 values
// (elements start at least two lanes apart, so at most 24 start in its
// first 48 lanes, plus the one ']' can close), so with 25 slots free the
// walk runs unchecked: one compare of the free bytes (computed from the
// cursor, since out may be empty with any base pointer, nil included).
// Below that, the block's exact bound is one value per comma in its
// first 48 lanes at or after b plus one (the element whose comma lies
// beyond them, or the one ']' closes), counted with CNT; a target presized
// from the array's comma count always meets it, and the walk again runs
// unchecked. Only a target too small for the block is walked under a lane
// limit — at most ceil(m/2) elements start in the first m lanes, so lanes
// below 2*free-1 store at most free values — and hands the rest back at
// the first element past it, filling what fits as the amd64 kernel does.
//
// Registers: R0 data base, R1 len-64 (the last block start), R2 s (the
// block's absolute start), R3 the out cursor, R4 out end, R20 out base,
// R21 48, R6 b (offset in the block), R7 the reversed not-digit mask, R8
// the reversed comma mask, R9 the reversed not-whitespace mask (R19
// whether it has been computed, R22 whether the previous block needed it),
// R10 the shuffle table, R11 the block's address, R13 the lane limit,
// R23-R25 the next block's masks, R26 its address, R27 whether there is
// one, R12/R14-R17 temporaries (R17 also names wsCompute's return site).
// V20-V23 the block minus '0' and V24-V27 the next block's, V4-V7 splats,
// V8-V10 constants, V0-V3/V12-V19 the temporaries of the classification,
// V29 the capacity test's, V29-V31 the fold's.
TEXT ·parseIntRunNEON(SB), NOSPLIT, $0-80
	MOVD  data_base+0(FP), R0
	MOVD  data_len+8(FP), R1
	SUB   $64, R1
	MOVD  i+24(FP), R2
	MOVD  out_base+32(FP), R3
	MOVD  out_len+40(FP), R4
	MOVD  R3, R20
	ADD   R4<<3, R3, R4
	MOVD  $·intRunShuffleNEON(SB), R10
	MOVD  $irConst<>(SB), R12
	VLD1  (R12), [V8.B16, V9.B16, V10.B16]
	VMOVI $0x30, V4.B16 // '0'
	VMOVI $9, V5.B16
	VMOVI $0x2c, V6.B16 // ','
	VMOVI $0x20, V7.B16 // whitespace is c <= 0x20, SkipWS's rule
	MOVD  $48, R21
	MOVD  ZR, R19

block:
	// The block at s, classified here: the first one, and the one a
	// straddle restarts at an element's own start. Its whitespace class
	// is computed on demand: R9 zero, so the walk's first probe of it
	// lands in wsZero, and R19 clear to say it is not there yet.
	MOVD  ZR, R6
	CMP   R1, R2
	BGT   done // fewer than 64 bytes from s: the scalar loop takes it from here
	ADD   R0, R2, R11
	PREP1(R11, V20.B16, V21.B16, V22.B16, V23.B16)
	WORD  $0x6e25368c // cmhi v12.16b, v20.16b, v5.16b
	WORD  $0x6e2536ad // cmhi v13.16b, v21.16b, v5.16b
	WORD  $0x6e2536ce // cmhi v14.16b, v22.16b, v5.16b
	WORD  $0x6e2536ef // cmhi v15.16b, v23.16b, v5.16b
	PREP2(R7, R8)
	MOVD  R19, R22
	MOVD  ZR, R9
	MOVD  ZR, R19

blockAt:
	// The block at s is classified (V20-V23, R7-R9) and its walk starts at
	// b. Before it, classify the block at s+48 if there is one, so that
	// work overlaps the walk — including its whitespace class when the
	// block before this one needed it (R22), since an array's separators
	// are alike: a compact array never pays for the class, and a
	// whitespace-separated one pays the on-demand form only on its first
	// block.
	ADD   $48, R2, R12
	CMP   R1, R12
	BGT   noNext
	ADD   $48, R11, R26
	PREP1(R26, V24.B16, V25.B16, V26.B16, V27.B16)
	WORD  $0x6e25370c // cmhi v12.16b, v24.16b, v5.16b
	WORD  $0x6e25372d // cmhi v13.16b, v25.16b, v5.16b
	WORD  $0x6e25374e // cmhi v14.16b, v26.16b, v5.16b
	WORD  $0x6e25376f // cmhi v15.16b, v27.16b, v5.16b
	PREP2(R23, R24)
	MOVD  ZR, R25
	CBZ   R22, haveNextNoWS
	WORD  $0x6e27340c // cmhi v12.16b, v0.16b, v7.16b
	WORD  $0x6e27342d // cmhi v13.16b, v1.16b, v7.16b
	WORD  $0x6e27344e // cmhi v14.16b, v2.16b, v7.16b
	WORD  $0x6e27346f // cmhi v15.16b, v3.16b, v7.16b
	WSPACK(V12.B16, V13.B16, V14.B16, V15.B16, V12, R25)

haveNextNoWS:
	MOVD  $1, R27
	B     haveNext

noNext:
	MOVD  ZR, R27

haveNext:
	MOVD  R21, R13                  // the lane limit: the whole 48-lane stride ...
	SUB   R3, R4, R16
	CMP   $200, R16
	BLO   capacity                  // ... unless fewer than 25 slots remain

elem:
	// Leading whitespace, only when the byte at b is not a digit (compact
	// input never enters the arithmetic): s = b + the whitespace run there,
	// and a non-digit after it is not this kernel's element.
	LSL   R6, R7, R12       // not-digit << b
	TBZ   $63, R12, digits

elemWS:
	LSL   R6, R9, R17       // not-whitespace << b
	CBZ   R17, wsZero       // to the block's end, or the class not computed yet
	CLZ   R17, R17
	ADD   R17, R6, R6       // s = b + run
	LSL   R6, R7, R12       // not-digit << s
	TBNZ  $63, R12, done

digits:
	// The digit run at s: L = leading zeros of not-digit << s, must be
	// 1..8. Digits to the block's end read as 64 and take the 9+ exit,
	// which is a stop the scalar loop resumes from like any other; the
	// kernel is then called again at s, so the block restarts there.
	CLZ   R12, R12          // L
	CMP   $8, R12
	BHI   done              // 9+ digits
	ADD   R12, R6, R14      // e: the byte after the digits
	// The common case is the comma right after the digits: one bit of the
	// comma mask shifted to e. The mask is never consumed — the shift to
	// each element's end excludes the commas before it — so no register
	// carries from one element to the next except b itself.
	LSL   R14, R8, R16      // commas << e
	TBZ   $63, R16, wsComma

fold:
	FOLD1
	WORD  $0x6e8897bf // udot v31.4s, v29.16b, v8.16b
	WORD  $0x4ea99fff // mul v31.4s, v31.4s, v9.4s
	FOLD2
	ADD   $1, R14, R6       // b = just past the comma
	CMP   R13, R6
	BLO   elem
	CMP   $48, R6
	BLO   done              // the slot limit, not the block's: leave the rest to the caller
	CBZ   R27, done         // no next block: the scalar loop takes it from here
	// The next element starts in the next block, at lane b-48, which is
	// already classified: make it the current one.
	ADD   $48, R2
	SUB   $48, R6, R6
	MOVD  R26, R11
	MOVD  R23, R7
	MOVD  R24, R8
	MOVD  R25, R9          // its whitespace class, if R22 had it computed ...
	MOVD  R22, R16
	MOVD  R19, R22         // ... and whether this block needed the class decides for the next
	MOVD  R16, R19
	VORR  V24.B16, V24.B16, V20.B16
	VORR  V25.B16, V25.B16, V21.B16
	VORR  V26.B16, V26.B16, V22.B16
	VORR  V27.B16, V27.B16, V23.B16
	B     blockAt

wsZero:
	CBNZ  R19, straddle
	MOVD  $1, R17
	B     wsCompute

wsComma:
	// No comma right after the digits: whitespace before it, or none in
	// the block. c > e (a comma cannot fall inside the digit run, and any
	// comma before s would have ended the whitespace skip with L == 0), so
	// the bytes e..c-1 must all be whitespace, i.e. the whitespace run
	// from e reaches c; then the fold proceeds with e moved to c.
	CBZ   R16, noComma
	CBNZ  R19, wsCommaAt
	MOVD  $2, R17
	B     wsCompute

wsCommaAt:
	CLZ   R16, R15          // c - e
	LSL   R14, R9, R17      // not-whitespace << e
	CLZ   R17, R17          // the whitespace run from e, 64 if to the block's end
	CMP   R15, R17
	BLO   done              // something other than whitespace before the comma
	ADD   R15, R14, R14
	B     fold

wsCompute:
	// The whitespace class, not whitespace (c > 0x20), for the block at
	// R11, in R9 with R19 set to say it is there. It is computed only when
	// the walk first needs it — a compact array never does, and its 13
	// vector ops a block were the largest removable item once the walk was
	// issue-bound — from a reload of the block, which is cheaper than
	// keeping its bytes in registers across the next block's PREP; R17
	// says which of the three sites asked.
	VLD1  (R11), [V12.B16, V13.B16, V14.B16, V15.B16]
	WORD  $0x6e273590 // cmhi v16.16b, v12.16b, v7.16b
	WORD  $0x6e2735b1 // cmhi v17.16b, v13.16b, v7.16b
	WORD  $0x6e2735d2 // cmhi v18.16b, v14.16b, v7.16b
	WORD  $0x6e2735f3 // cmhi v19.16b, v15.16b, v7.16b
	WSPACK(V16.B16, V17.B16, V18.B16, V19.B16, V16, R9)
	MOVD  $1, R19
	CMP   $1, R17
	BEQ   elemWS
	CMP   $2, R17
	BEQ   wsCommaAt
	B     noCommaAt

capacity:
	// Fewer than 25 slots: the block's exact bound is its commas in lanes
	// b..47 plus one. Met, the walk is unchecked as above; not met, the
	// lane limit min(48, 2*free-1) bounds the stores by the free slots.
	SUB   R3, R4, R13
	LSR   $3, R13, R13      // free slots
	CBZ   R13, done
	LSL   R6, R8, R16       // the commas at or after b ...
	LSR   R6, R16, R16
	AND   $0xffffffffffff0000, R16, R16 // ... and below lane 48
	FMOVD R16, F29
	VCNT  V29.B8, V29.B8
	VUADDLV V29.B8, V29
	FMOVD F29, R16
	ADD   $1, R16, R16      // the bound
	CMP   R16, R13
	BLO   partial
	MOVD  R21, R13
	B     elem

partial:
	LSL   $1, R13, R13
	SUB   $1, R13, R13      // 2*free - 1
	CMP   R21, R13
	CSEL  LO, R13, R21, R13
	B     elem

noComma:
	// No comma left in the block: the element can only be the array's
	// last, terminated by ']' (after optional whitespace) in this block.
	CBNZ  R19, noCommaAt
	MOVD  $3, R17
	B     wsCompute

noCommaAt:
	LSL   R14, R9, R16      // not-whitespace << e
	CBZ   R16, straddle
	CLZ   R16, R16
	ADD   R14, R16, R16     // the terminator's offset
	MOVBU (R11)(R16), R17
	CMP   $0x5d, R17
	BNE   done              // not ']': not this kernel's element
	MOVD  R16, R15

foldClose:
	// The array's last element, terminated by ']' at R15: fold it, then
	// return with p at the ']' and closed set.
	FOLD1
	WORD  $0x6e8897bf // udot v31.4s, v29.16b, v8.16b
	WORD  $0x4ea99fff // mul v31.4s, v31.4s, v9.4s
	FOLD2
	ADD   R15, R2, R2
	SUB   R20, R3, R16
	LSR   $3, R16, R16
	MOVD  R16, n+56(FP)
	MOVD  R2, p+64(FP)
	MOVD  $1, R16
	MOVD  R16, closed+72(FP)
	RET

straddle:
	// The element (or the whitespace before it) runs to the block's end:
	// start the next block at it. At b == 0 that would not move, so stop.
	CBZ   R6, done
	ADD   R6, R2, R2
	B     block

done:
	ADD   R6, R2, R2
	SUB   R20, R3, R16
	LSR   $3, R16, R16
	MOVD  R16, n+56(FP)
	MOVD  R2, p+64(FP)
	MOVD  ZR, closed+72(FP)
	RET
