#include "textflag.h"

// Constants for parseFloatRunAVX2 (see floatrun_amd64.go).
DATA frZero<>+0(SB)/8, $0x3030303030303030 // '0'
DATA frZero<>+8(SB)/8, $0x3030303030303030
DATA frZero<>+16(SB)/8, $0x3030303030303030
DATA frZero<>+24(SB)/8, $0x3030303030303030
GLOBL frZero<>(SB), RODATA|NOPTR, $32

DATA frNine<>+0(SB)/8, $0x0909090909090909 // 9: (c-'0') <= 9 via VPMINUB
DATA frNine<>+8(SB)/8, $0x0909090909090909
DATA frNine<>+16(SB)/8, $0x0909090909090909
DATA frNine<>+24(SB)/8, $0x0909090909090909
GLOBL frNine<>(SB), RODATA|NOPTR, $32

DATA frComma<>+0(SB)/8, $0x2c2c2c2c2c2c2c2c // ','
DATA frComma<>+8(SB)/8, $0x2c2c2c2c2c2c2c2c
DATA frComma<>+16(SB)/8, $0x2c2c2c2c2c2c2c2c
DATA frComma<>+24(SB)/8, $0x2c2c2c2c2c2c2c2c
GLOBL frComma<>(SB), RODATA|NOPTR, $32

DATA frClose<>+0(SB)/8, $0x5d5d5d5d5d5d5d5d // ']'
DATA frClose<>+8(SB)/8, $0x5d5d5d5d5d5d5d5d
DATA frClose<>+16(SB)/8, $0x5d5d5d5d5d5d5d5d
DATA frClose<>+24(SB)/8, $0x5d5d5d5d5d5d5d5d
GLOBL frClose<>(SB), RODATA|NOPTR, $32

DATA frSpace<>+0(SB)/8, $0x2020202020202020 // 0x20: c <= 0x20 via VPMINUB (SkipWS's rule)
DATA frSpace<>+8(SB)/8, $0x2020202020202020
DATA frSpace<>+16(SB)/8, $0x2020202020202020
DATA frSpace<>+24(SB)/8, $0x2020202020202020
GLOBL frSpace<>(SB), RODATA|NOPTR, $32

// The digit folds, as in parseIntRunAVX2, and the weights that join the two
// eight-digit halves in double precision: [1e8, 1].
DATA frW10<>+0(SB)/8, $0x010a010a010a010a
DATA frW10<>+8(SB)/8, $0x010a010a010a010a
GLOBL frW10<>(SB), RODATA|NOPTR, $16

DATA frW100<>+0(SB)/8, $0x0001006400010064
DATA frW100<>+8(SB)/8, $0x0001006400010064
GLOBL frW100<>(SB), RODATA|NOPTR, $16

DATA frW1e4<>+0(SB)/8, $0x0001271000012710
DATA frW1e4<>+8(SB)/8, $0x0001271000012710
GLOBL frW1e4<>(SB), RODATA|NOPTR, $16

DATA frJoin<>+0(SB)/8, $0x4197d78400000000 // 1e8
DATA frJoin<>+8(SB)/8, $0x3ff0000000000000 // 1.0
GLOBL frJoin<>(SB), RODATA|NOPTR, $16

// CLASSIFY builds the window's three 64-bit lane masks from Y0 (bytes 0-31)
// and Y1 (32-63): R11 = NOT a digit, R12 = a comma, R13 = NOT whitespace
// (<= 0x20).
#define CLASSIFY \
	VPSUBB    Y4, Y0, Y2   \
	VPMINUB   Y5, Y2, Y3   \
	VPCMPEQB  Y2, Y3, Y3   \
	VPMOVMSKB Y3, R11      \
	VPCMPEQB  Y6, Y0, Y3   \
	VPMOVMSKB Y3, R12      \
	VPMINUB   Y7, Y0, Y3   \
	VPCMPEQB  Y0, Y3, Y3   \
	VPMOVMSKB Y3, R13      \
	VPSUBB    Y4, Y1, Y2   \
	VPMINUB   Y5, Y2, Y3   \
	VPCMPEQB  Y2, Y3, Y3   \
	VPMOVMSKB Y3, AX       \
	SHLQ      $32, AX      \
	ORQ       AX, R11      \
	VPCMPEQB  Y6, Y1, Y3   \
	VPMOVMSKB Y3, AX       \
	SHLQ      $32, AX      \
	ORQ       AX, R12      \
	VPMINUB   Y7, Y1, Y3   \
	VPCMPEQB  Y1, Y3, Y3   \
	VPMOVMSKB Y3, AX       \
	SHLQ      $32, AX      \
	ORQ       AX, R13      \
	NOTQ      R11          \
	NOTQ      R13

// CLOSES builds the window's ']' lane mask in BX, for the last element when
// no comma is left after the cursor. (When one is, and it is the enclosing
// container's, the ']' is the first non-whitespace byte after the element and
// wsBeforeComma finds it there; computing this mask for every window instead
// cost the integer kernel up to 0.7 cycles an element.)
#define CLOSES \
	VPCMPEQB  frClose<>(SB), Y0, Y2 \
	VPMOVMSKB Y2, BX       \
	VPCMPEQB  frClose<>(SB), Y1, Y2 \
	VPMOVMSKB Y2, AX       \
	SHLQ      $32, AX      \
	ORQ       AX, BX

// PARSE measures the number whose region starts at window lane CX and ends
// at a delimiter lane (a comma or the ']', in BX): CX = s, its first digit
// (whitespace and a '-' before it skipped, DI = 1 for the '-'), AX = L1
// integer digits, R14 = L2 fraction digits (0 without a '.'), DX = e, the lane
// after the number. The delimiter is neither whitespace nor a digit, so every
// lane this measures is at or below it — below 64 — which is what keeps the
// shift counts and the byte loads in the window. Anything this kernel does not
// convert jumps to bad: no integer digit, more than 15 digits in all, a '.'
// with no digit after it.
#define PARSE(bad, nofrac, parsed) \
	SHRXQ   CX, R13, AX           \
	TZCNTQ  AX, AX                \
	ADDQ    AX, CX                \
	XORL    DI, DI                \
	CMPB    (SI)(CX*1), $0x2d     \
	SETEQ   DI                    \
	ADDQ    DI, CX                \
	SHRXQ   CX, R11, AX           \
	TZCNTQ  AX, AX                \
	LEAQ    -1(AX), DX            \
	CMPQ    DX, $14               \
	JA      bad                   \
	LEAQ    (CX)(AX*1), DX        \
	XORL    R14, R14              \
	CMPB    (SI)(DX*1), $0x2e     \
	JNE     nofrac                \
	LEAQ    1(DX), R14            \
	SHRXQ   R14, R11, R14         \
	TZCNTQ  R14, R14              \
	TESTQ   R14, R14              \
	JZ      bad                   \
	LEAQ    1(DX)(R14*1), DX      \
nofrac:                           \
	SUBQ    CX, DX                \
	CMPQ    DX, $16               \
	JA      bad                   \
	ADDQ    CX, DX                \
parsed:

// CONVERT folds the digits PARSE measured and stores ±mantissa / 10^L2 at
// out[R10]. The sixteen bytes at s hold the integer digits, the '.', and the
// fraction digits; the (L1, L2) control right-aligns the L1+L2 digits with
// zeros in front and drops the '.', the folds give the top and bottom eight
// digits as two dwords, and with at most 15 digits the top half is below 1e7,
// so hi*1e8 + lo is exact in double precision. The divide by the exact power
// 10^L2 (L2 <= 14) is then one correctly rounded IEEE operation on two exact
// operands — Clinger's fast path, which is what strconv returns — and the
// divisor carries the sign, so -0 comes out as -0. AX, DI, R14 are clobbered.
#define CONVERT \
	VMOVDQU    (SI)(CX*1), X2       \
	VPSUBB     X4, X2, X2           \
	SHLQ       $5, DI               \
	ADDQ       R14, DI              \
	SHLQ       $8, AX               \
	SHLQ       $4, R14              \
	ADDQ       R14, AX              \
	VPSHUFB    (R15)(AX*1), X2, X2  \
	VPMADDUBSW X8, X2, X2           \
	VPMADDWD   X9, X2, X2           \
	VPACKUSDW  X2, X2, X2           \
	VPMADDWD   X10, X2, X2          \
	VCVTDQ2PD  X2, X2               \
	VMULPD     X11, X2, X2          \
	VPERMILPD  $1, X2, X3           \
	VADDSD     X3, X2, X2           \
	VDIVSD     4096(R15)(DI*8), X2, X2 \
	VMOVSD     X2, (R8)(R10*8)      \
	INCQ       R10

// func parseFloatRunAVX2(data []byte, i int, out []float64) (n, p, closed int)
//
// parseIntRunAVX2's walk for arrays of decimal numbers: parses as many
// "ws* -? digits (. digits)? ws* ','" groups as it can from data[i:] — at most
// 15 digits in all, no exponent — one float64 per group into out, and also the
// array's last element when it is terminated by ']' (closed = 1, p at the
// ']'). Stops, with p at the start of the unconsumed element's region (or
// part-way into its whitespace) or right after the last consumed comma, when
// fewer than 80 bytes remain from the
// window (the fold loads 16 bytes at any lane of a 64-byte window), out is
// full, or an element is anything else: an exponent, 16+ digits, a '+', null,
// a '.' without digits after it. Every stop position is a state the scalar
// loop resumes from, and every value the kernel does write is Clinger's exact
// fast path, so it is the value strconv returns.
//
// The window walk is parseIntRunAVX2's (64-byte windows at a fixed 48-byte
// stride, commas walked with TZCNT/BLSR); see intrun_amd64.s. The array's last
// element, which has no comma, is finished in place when its ']' is in the
// window and otherwise handed back.
//
// Registers: SI the window, CX the cursor, R8/R9 out base/len, R10 values
// written, R11/R12/R13 the window masks, R15 floatRunTab (the shuffle controls,
// then ±10^k at +4096; the VBMI body's tables follow), BX the comma, and
// AX/DX/DI/R14 the element's measures.
// The data base and the window limit are reloaded from the frame when needed.
// Y4-Y7 and X8-X11 hold the constants.
TEXT ·parseFloatRunAVX2(SB), NOSPLIT, $0-80
	MOVBLZX ·useFloatRunLong(SB), AX
	TESTL   AX, AX
	JNZ     toVBMI
	MOVQ    data_base+0(FP), SI
	MOVQ    i+24(FP), CX
	MOVQ    out_base+32(FP), R8
	MOVQ    out_len+40(FP), R9
	XORQ    R10, R10
	ADDQ    CX, SI
	XORL    CX, CX
	VMOVDQU frZero<>(SB), Y4
	VMOVDQU frNine<>(SB), Y5
	VMOVDQU frComma<>(SB), Y6
	VMOVDQU frSpace<>(SB), Y7
	VMOVDQU frW10<>(SB), X8
	VMOVDQU frW100<>(SB), X9
	VMOVDQU frW1e4<>(SB), X10
	VMOVDQU frJoin<>(SB), X11
	LEAQ    ·floatRunTab(SB), R15

window:
	MOVQ    data_base+0(FP), AX
	ADDQ    data_len+8(FP), AX
	SUBQ    $80, AX                 // the last window start with 80 bytes
	CMPQ    SI, AX
	JHI     stop
	VMOVDQU (SI), Y0
	VMOVDQU 32(SI), Y1
	CLASSIFY
	MOVQ    $-1, AX
	SHLXQ   CX, AX, AX
	ANDQ    AX, R12                 // the commas at or after b
	JZ      noComma

elem:
	TZCNTQ  R12, BX                 // c: this element's comma
	PARSE(stopElem, elemNoFrac, elemParsed)
	CMPQ    DX, BX
	JNE     wsBeforeComma
fold:
	CMPQ    R10, R9
	JAE     stopElem                // out is full: leave this element to the caller
	CONVERT
	LEAQ    1(BX), CX               // b: just past the comma
	BLSRQ   R12, R12                // the comma is consumed; ZF: none left
	JNZ     elem

noComma:
	// No comma left in the window: step by 48 when the cursor is in the last
	// 16 bytes. Otherwise this is the array's last element if the window
	// holds a ']' after the cursor, and that ']' bounds it exactly as a comma
	// bounds the others — every lane PARSE measures is then below it, so no
	// shift count reaches 64 and no load leaves the window. Without one the
	// region runs past the window and the window restarts at it.
	CMPQ     CX, $48
	JAE      step
	CLOSES
	MOVQ     $-1, AX
	SHLXQ    CX, AX, AX
	ANDQ     AX, BX                 // the ']'s at or after b
	JZ       restart
	TZCNTQ   BX, BX                 // r: the ']'
	PARSE(stopElem, lastNoFrac, lastParsed)
	CMPQ     DX, BX
	JNE      lastWs
lastFold:
	CMPQ     R10, R9
	JAE      stopElem
	CONVERT
	ADDQ     BX, SI
	SUBQ     data_base+0(FP), SI
	MOVQ     R10, n+56(FP)
	MOVQ     SI, p+64(FP)
	MOVQ     $1, closed+72(FP)
	VZEROUPPER
	RET

lastWs:
	SHRXQ   DX, R13, AX
	TZCNTQ  AX, AX
	ADDQ    DX, AX
	CMPQ    AX, BX
	JNE     stopElem
	SHRXQ   CX, R11, AX
	TZCNTQ  AX, AX
	JMP     lastFold

wsBeforeComma:
	// e < c: the bytes between must all be whitespace — or the first one that
	// is not must be the array's ']', the comma being the enclosing
	// container's (an array is followed by more of the document, often in the
	// same window), and then this element is the last. L1 is measured again,
	// AX having been needed for the test.
	SHRXQ   DX, R13, AX
	TZCNTQ  AX, AX
	ADDQ    DX, AX
	CMPQ    AX, BX
	JNE     closedBefore
	SHRXQ   CX, R11, AX
	TZCNTQ  AX, AX
	JMP     fold
closedBefore:
	CMPB    (SI)(AX*1), $0x5d
	JNE     stopElem
	MOVQ    AX, BX                  // r: the ']'
	SHRXQ   CX, R11, AX
	TZCNTQ  AX, AX
	JMP     lastFold

restart:
	// Neither a comma nor a ']' after the cursor: the element region runs past
	// the window (longer than the stride's 16 bytes of slack). Start the next
	// window at its first byte that is not whitespace — or past the window if
	// all of it is — and stop only if that would not move.
	SHRXQ   CX, R13, AX
	TESTQ   AX, AX
	JZ      skipBlank
	TZCNTQ  AX, AX
	ADDQ    AX, CX
	TESTQ   CX, CX
	JZ      stop
	ADDQ    CX, SI
	XORL    CX, CX
	JMP     window

skipBlank:
	ADDQ    $64, SI
	XORL    CX, CX
	JMP     window

step:
	ADDQ    $48, SI
	SUBQ    $48, CX
	JMP     window

stopElem:
	// A stop after PARSE: CX is past the element's '-', if it has one, and
	// the scalar loop must see the sign — resuming at the digits would read
	// a negative number as positive. DI is the sign flag PARSE set.
	SUBQ    DI, CX

stop:
	VZEROUPPER
	ADDQ    CX, SI
	SUBQ    data_base+0(FP), SI
	MOVQ    R10, n+56(FP)
	MOVQ    SI, p+64(FP)
	MOVQ    $0, closed+72(FP)
	RET

toVBMI:
	JMP     ·parseFloatRunVBMI(SB)

// ---- The same walk as a check, for SkipValueStrict ------------------------------

// CHECK is PARSE for validNumberRun: the same measurement, with no limit on
// the digit count and no span test — nothing is converted, and a number
// bounded by its delimiter in a 64-byte window has fewer than 64 digits, so no
// magnitude it can spell overflows a float64 — and the same rejections: no
// integer digit, or a '.' with no digit after it.
#define CHECK(bad, nofrac, checked) \
	SHRXQ   CX, R13, AX           \
	TZCNTQ  AX, AX                \
	ADDQ    AX, CX                \
	XORL    DI, DI                \
	CMPB    (SI)(CX*1), $0x2d     \
	SETEQ   DI                    \
	ADDQ    DI, CX                \
	SHRXQ   CX, R11, AX           \
	TZCNTQ  AX, AX                \
	TESTQ   AX, AX                \
	JZ      bad                   \
	LEAQ    (CX)(AX*1), DX        \
	CMPB    (SI)(DX*1), $0x2e     \
	JNE     nofrac                \
	LEAQ    1(DX), AX             \
	SHRXQ   AX, R11, AX           \
	TZCNTQ  AX, AX                \
	TESTQ   AX, AX                \
	JZ      bad                   \
	LEAQ    1(DX)(AX*1), DX       \
nofrac:                           \
checked:

// func validNumberRun(data []byte, i int) (p, closed int)
//
// parseFloatRunAVX2's window walk as a check, for SkipValueStrict: it passes
// over "ws* -? digits (. digits)? ws* ','" elements from data[i:], and over the
// array's last element when a ']' ends it (closed = 1, p at the ']'),
// converting nothing. Otherwise p is where the scalar walk resumes: the start
// of the first element region it did not take (its '-' included), a place in
// that region's leading whitespace, or just past the last comma it passed. It
// takes only numbers the decoder's own reader accepts — plain decimals, which
// cannot overflow at under 64 digits — and leaves everything else (an
// exponent, a '+', a bare '.', a leading '.', a string, a container, a stray
// byte) to the scalar walk, which decides exactly as before. The walk reads
// whole 64-byte windows only, so it stops when fewer than 64 bytes remain.
//
// Registers as in parseFloatRunAVX2, without the output and the conversion
// tables: SI the window, CX the cursor, R11/R12/R13 the window masks, BX the
// delimiter, AX/DX/DI temporaries; Y4-Y7 the constants.
TEXT ·validNumberRun(SB), NOSPLIT, $0-48
	MOVBLZX ·useValidPoints(SB), AX
	TESTL   AX, AX
	JNZ     to512
	MOVQ    data_base+0(FP), SI
	MOVQ    i+24(FP), CX
	ADDQ    CX, SI
	XORL    CX, CX
	VMOVDQU frZero<>(SB), Y4
	VMOVDQU frNine<>(SB), Y5
	VMOVDQU frComma<>(SB), Y6
	VMOVDQU frSpace<>(SB), Y7

cwindow:
	MOVQ    data_base+0(FP), AX
	ADDQ    data_len+8(FP), AX
	SUBQ    $64, AX                 // the last window start with 64 bytes
	CMPQ    SI, AX
	JHI     cstop
	VMOVDQU (SI), Y0
	VMOVDQU 32(SI), Y1
	CLASSIFY
	MOVQ    $-1, AX
	SHLXQ   CX, AX, AX
	ANDQ    AX, R12                 // the commas at or after b
	JZ      cnoComma

celem:
	TZCNTQ  R12, BX                 // c: this element's comma
	CHECK(cstopElem, celemNoFrac, celemChecked)
	CMPQ    DX, BX
	JNE     cwsBeforeComma
cnext:
	LEAQ    1(BX), CX               // b: just past the comma
	BLSRQ   R12, R12                // the comma is consumed; ZF: none left
	JNZ     celem

cnoComma:
	// As in parseFloatRunAVX2: step by 48 when the cursor is in the last 16
	// bytes, take the array's last element when a ']' after the cursor bounds
	// it, and otherwise restart the window at the element region.
	CMPQ     CX, $48
	JAE      cstep
	CLOSES
	MOVQ     $-1, AX
	SHLXQ    CX, AX, AX
	ANDQ     AX, BX                 // the ']'s at or after b
	JZ       crestart
	TZCNTQ   BX, BX                 // r: the ']'
	CHECK(cstopElem, clastNoFrac, clastChecked)
	CMPQ     DX, BX
	JEQ      cclosed
	SHRXQ    DX, R13, AX            // whitespace, then the ']'
	TZCNTQ   AX, AX
	ADDQ     DX, AX
	CMPQ     AX, BX
	JNE      cstopElem

cclosed:
	ADDQ    BX, SI
	SUBQ    data_base+0(FP), SI
	MOVQ    SI, p+32(FP)
	MOVQ    $1, closed+40(FP)
	VZEROUPPER
	RET

cwsBeforeComma:
	// Whitespace to the comma, or whitespace to the array's ']' — the comma
	// being the enclosing container's — and then this element was the last.
	SHRXQ   DX, R13, AX
	TZCNTQ  AX, AX
	ADDQ    DX, AX
	CMPQ    AX, BX
	JEQ     cnext
	CMPB    (SI)(AX*1), $0x5d
	JNE     cstopElem
	MOVQ    AX, BX                  // r: the ']'
	JMP     cclosed

crestart:
	SHRXQ   CX, R13, AX
	TESTQ   AX, AX
	JZ      cskipBlank
	TZCNTQ  AX, AX
	ADDQ    AX, CX
	TESTQ   CX, CX
	JZ      cstop
	ADDQ    CX, SI
	XORL    CX, CX
	JMP     cwindow

cskipBlank:
	ADDQ    $64, SI
	XORL    CX, CX
	JMP     cwindow

cstep:
	ADDQ    $48, SI
	SUBQ    $48, CX
	JMP     cwindow

cstopElem:
	// A stop after CHECK: back over the element's '-', so that the scalar
	// walk reads the element whole ("--5" must not become "-5").
	SUBQ    DI, CX

cstop:
	VZEROUPPER
	ADDQ    CX, SI
	SUBQ    data_base+0(FP), SI
	MOVQ    SI, p+32(FP)
	MOVQ    $0, closed+40(FP)
	RET

to512:
	JMP     ·validNumberRun512(SB)

// func validNumberRun512(data []byte, i int) (p, closed int)
//
// validNumberRun with the window classified by AVX-512 compares into mask
// registers: one 64-byte load and three compares and three KMOVQs make the
// three masks CLASSIFY builds from two ymm halves in 26 instructions, and the
// ']' mask is one more compare. It matters because the walk is often called
// for one short array — a coordinate point — so the window's classification
// is most of a call. Reached only by validNumberRun's tail jump (the flag is
// read in the assembly so that the Go side stays a single call); the walk is
// validNumberRun's, label for label.
TEXT ·validNumberRun512(SB), NOSPLIT, $0-48
	MOVQ    data_base+0(FP), SI
	MOVQ    i+24(FP), CX
	ADDQ    CX, SI
	XORL    CX, CX
	VPBROADCASTB fvZero<>(SB), Z4
	VPBROADCASTB fvNine<>(SB), Z5
	VPBROADCASTB fvComma<>(SB), Z6
	VPBROADCASTB fvSpace<>(SB), Z7
	VPBROADCASTB fvClose<>(SB), Z3

dwindow:
	MOVQ    data_base+0(FP), AX
	ADDQ    data_len+8(FP), AX
	SUBQ    $64, AX                 // the last window start with 64 bytes
	CMPQ    SI, AX
	JHI     dstop
	VMOVDQU64 (SI), Z0
	VPSUBB    Z4, Z0, Z1
	VPCMPUB   $6, Z5, Z1, K1        // not a digit
	VPCMPEQB  Z6, Z0, K2            // a comma
	VPCMPUB   $6, Z7, Z0, K3        // not whitespace
	KMOVQ     K1, R11
	KMOVQ     K2, R12
	KMOVQ     K3, R13
	MOVQ    $-1, AX
	SHLXQ   CX, AX, AX
	ANDQ    AX, R12                 // the commas at or after b
	JZ      dnoComma

delem:
	TZCNTQ  R12, BX                 // c: this element's comma
	CHECK(dstopElem, delemNoFrac, delemChecked)
	CMPQ    DX, BX
	JNE     dwsBeforeComma
dnext:
	LEAQ    1(BX), CX               // b: just past the comma
	BLSRQ   R12, R12                // the comma is consumed; ZF: none left
	JNZ     delem

dnoComma:
	CMPQ     CX, $48
	JAE      dstep
	VPCMPEQB Z3, Z0, K4             // a ']'
	KMOVQ    K4, BX
	MOVQ     $-1, AX
	SHLXQ    CX, AX, AX
	ANDQ     AX, BX                 // the ']'s at or after b
	JZ       drestart
	TZCNTQ   BX, BX                 // r: the ']'
	CHECK(dstopElem, dlastNoFrac, dlastChecked)
	CMPQ     DX, BX
	JEQ      dclosed
	SHRXQ    DX, R13, AX            // whitespace, then the ']'
	TZCNTQ   AX, AX
	ADDQ     DX, AX
	CMPQ     AX, BX
	JNE      dstopElem

dclosed:
	ADDQ    BX, SI
	SUBQ    data_base+0(FP), SI
	MOVQ    SI, p+32(FP)
	MOVQ    $1, closed+40(FP)
	VZEROUPPER
	RET

dwsBeforeComma:
	SHRXQ   DX, R13, AX
	TZCNTQ  AX, AX
	ADDQ    DX, AX
	CMPQ    AX, BX
	JEQ     dnext
	CMPB    (SI)(AX*1), $0x5d
	JNE     dstopElem
	MOVQ    AX, BX                  // r: the ']'
	JMP     dclosed

drestart:
	SHRXQ   CX, R13, AX
	TESTQ   AX, AX
	JZ      dskipBlank
	TZCNTQ  AX, AX
	ADDQ    AX, CX
	TESTQ   CX, CX
	JZ      dstop
	ADDQ    CX, SI
	XORL    CX, CX
	JMP     dwindow

dskipBlank:
	ADDQ    $64, SI
	XORL    CX, CX
	JMP     dwindow

dstep:
	ADDQ    $48, SI
	SUBQ    $48, CX
	JMP     dwindow

dstopElem:
	SUBQ    DI, CX

dstop:
	VZEROUPPER
	ADDQ    CX, SI
	SUBQ    data_base+0(FP), SI
	MOVQ    SI, p+32(FP)
	MOVQ    $0, closed+40(FP)
	RET

// func validPointsRun(data []byte, i int) (p, closed int)
//
// The points walk (parseFloatPointsVBMI) as a check, for SkipValueStrict: i is
// at a point's '[' in an array of flat numeric arrays — a coordinate ring —
// and it passes over as many points as it can, each "[" number ("," number)*
// "]" with any whitespace, the numbers as validNumberRun's CHECK takes them,
// returning closed = 1 with p at the ring's ']' when it reaches it. Otherwise p
// is the start of the first point it did not take (the region after the
// previous point's ','), which the scalar walk then checks whole — an empty
// point, an exponent, a nested array, a string: anything but plain decimals.
// Unlike the conversion walk it takes points of any length that fit a window,
// since any count of numbers is valid. One window per point, at the point, and
// the separator after it looked for across windows of whitespace, as in the
// conversion walk; windows are read whole, so it stops when fewer than 64
// bytes remain. The caller runs it only below MaxDepth, the points being one
// level deeper than the ring.
//
// Registers as in validNumberRun512, plus R8 the point's ']'; the frame holds
// the point's start and the window limit.
TEXT ·validPointsRun(SB), NOSPLIT, $16-48
	MOVQ    data_base+0(FP), SI
	MOVQ    data_len+8(FP), AX
	LEAQ    -64(SI)(AX*1), AX
	MOVQ    AX, qlimit-16(SP)       // the last window start with 64 bytes
	ADDQ    i+24(FP), SI
	VPBROADCASTB fvZero<>(SB), Z4
	VPBROADCASTB fvNine<>(SB), Z5
	VPBROADCASTB fvComma<>(SB), Z6
	VPBROADCASTB fvSpace<>(SB), Z7
	VPBROADCASTB fvClose<>(SB), Z3

qpoint:
	MOVQ    SI, qstart-8(SP)        // where the point's region starts
	CMPQ    SI, qlimit-16(SP)
	JHI     qstop
	VMOVDQU64 (SI), Z0
	VPSUBB    Z4, Z0, Z1
	VPCMPUB   $6, Z5, Z1, K1        // not a digit
	VPCMPEQB  Z6, Z0, K2            // a comma
	VPCMPUB   $6, Z7, Z0, K3        // not whitespace
	VPCMPEQB  Z3, Z0, K4            // a ']'
	KMOVQ     K1, R11
	KMOVQ     K2, R12
	KMOVQ     K3, R13
	TZCNTQ    R13, CX               // the first byte that is not whitespace
	CMPQ      CX, $16
	JAE       qrewindow             // far into the window: start one at it
	CMPB      (SI)(CX*1), $0x5b
	JNE       qstop                 // not a point ('[')
	INCQ      CX                    // b: past the '['
	// Everything before b is whitespace and the '[', so the window's first
	// ']' is the point's (or a nested array's, which CHECK then refuses), and
	// its commas are the ones below that.
	KMOVQ     K4, R8
	TZCNTQ    R8, R8
	JCS       qlong                 // the point's ']' is past the window
	BZHIQ     R8, R12, R12

qelem:
	TZCNTQ  R12, BX                 // an inner number's delimiter: its comma
	JCC     qcheck
	MOVQ    R8, BX                  // the last number's: the point's ']'
qcheck:
	CHECK(qstop, qNoFrac, qChecked)
	CMPQ    DX, BX
	JEQ     qnext
	SHRXQ   DX, R13, AX             // whitespace to the delimiter
	TZCNTQ  AX, AX
	ADDQ    DX, AX
	CMPQ    AX, BX
	JNE     qstop
qnext:
	LEAQ    1(BX), CX               // past the delimiter
	BLSRQ   R12, R12                // the comma is consumed; CF: there was none
	JCC     qelem                   // left, so that was the point's ']'

	// The first byte after the point that is not whitespace must be the ','
	// before the next point or the ring's ']'.
	CMPQ    CX, $64
	JEQ     qsep
	SHRXQ   CX, R13, AX
	TZCNTQ  AX, AX
	JCS     qsep                    // the separator is past the window
	ADDQ    CX, AX
qsepAt:
	CMPB    (SI)(AX*1), $0x2c
	JEQ     qcomma
	CMPB    (SI)(AX*1), $0x5d
	JNE     qstop
	ADDQ    AX, SI                  // the ring's ']': done, and closed
	SUBQ    data_base+0(FP), SI
	MOVQ    SI, p+32(FP)
	MOVQ    $1, closed+40(FP)
	VZEROUPPER
	RET

qcomma:
	LEAQ    1(SI)(AX*1), SI         // the next point's region starts after the ','
	JMP     qpoint

qsep:
	// The rest of the window is whitespace: look for the separator in the
	// windows after it while they hold nothing else.
	ADDQ    $64, SI
	CMPQ    SI, qlimit-16(SP)
	JHI     qstop
	VMOVDQU64 (SI), Z0
	VPCMPUB   $6, Z7, Z0, K3        // not whitespace
	KMOVQ     K3, AX
	TZCNTQ    AX, AX
	JCS       qsep                  // all of it whitespace
	JMP       qsepAt

qlong:
	// As in the conversion walk: a point that did not fit behind whitespace
	// gets one more window, at its '['.
	CMPQ    CX, $1
	JEQ     qstop
	DECQ    CX

qrewindow:
	// Sixteen or more bytes of whitespace before the point, or a point that
	// did not fit behind less: start the next window at the first byte that
	// is not whitespace.
	ADDQ    CX, SI
	JMP     qpoint

qstop:
	// Hand back the point being checked, whole, from its region's start.
	VZEROUPPER
	MOVQ    qstart-8(SP), SI
	SUBQ    data_base+0(FP), SI
	MOVQ    SI, p+32(FP)
	MOVQ    $0, closed+40(FP)
	RET

// ---- The AVX-512 VBMI body: up to 19 digits, and Eisel-Lemire ----------------

DATA fvZero<>+0(SB)/1, $0x30
GLOBL fvZero<>(SB), RODATA|NOPTR, $1
DATA fvNine<>+0(SB)/1, $0x09
GLOBL fvNine<>(SB), RODATA|NOPTR, $1
DATA fvComma<>+0(SB)/1, $0x2c
GLOBL fvComma<>(SB), RODATA|NOPTR, $1
DATA fvSpace<>+0(SB)/1, $0x20
GLOBL fvSpace<>(SB), RODATA|NOPTR, $1
DATA fvClose<>+0(SB)/1, $0x5d
GLOBL fvClose<>(SB), RODATA|NOPTR, $1
DATA fvMinus<>+0(SB)/1, $0x2d
GLOBL fvMinus<>(SB), RODATA|NOPTR, $1

// VPARSE is PARSE for this body: the sign is read from the window's '-' mask
// in R9 (a bit test is two cycles on the cursor's chain where PARSE's byte
// load and compare are seven, and every number waits on that chain for its
// gather), and neither the digit count nor the span is bounded here. The
// number lies between the region start and its delimiter — a non-digit below
// lane 64 — so its end is at most the delimiter's lane whatever it measures,
// and the one limit left, 19 digits in all for the gather's template, is
// tested where a number takes that path (vlong, plong); a number of fifteen
// digits or fewer is within every limit.
#define VPARSE(bad, nofrac, parsed) \
	SHRXQ   CX, R13, AX           \
	TZCNTQ  AX, AX                \
	ADDQ    AX, CX                \
	XORL    DI, DI                \
	BTQ     CX, R9                \
	SETCS   DI                    \
	ADCQ    $0, CX                \
	SHRXQ   CX, R11, AX           \
	TZCNTQ  AX, AX                \
	TESTQ   AX, AX                \
	JZ      bad                   \
	LEAQ    (CX)(AX*1), DX        \
	XORL    R14, R14              \
	CMPB    (SI)(DX*1), $0x2e     \
	JNE     nofrac                \
	LEAQ    1(DX), R14            \
	SHRXQ   R14, R11, R14         \
	TZCNTQ  R14, R14              \
	TESTQ   R14, R14              \
	JZ      bad                   \
	LEAQ    1(DX)(R14*1), DX      \
nofrac:                           \
parsed:

// LONGCONV converts a number VPARSE measured at 16 to 19 digits and stores it
// at out[R10], R10 incremented, falling through when done: AX = L1, R14 = L2,
// CX = s, DI = the sign, and out[R10] in bounds. It clobbers the cursor and
// the sign, so a caller that hands a declined element back from its first
// byte keeps that elsewhere first (the flat walk parks it in the output slot;
// the points walk hands back the whole point from its start). One VPERMI2B
// gathers the digits from the 32 bytes at the first digit and a register of
// '0' under the (L1, L2) template, the folds give three eight-digit groups,
// the mantissa is joined in a general register (below 10^19), and then
// Clinger when it is below 2^53, or Eisel-Lemire (eiselLemire64 transcribed;
// see parseFloatRunVBMI).
//
// The transcription leaves out eiselLemire64's two range tests, which cannot
// fire on this body's numbers: a mantissa of 2^53 or more over at most 10^18
// is at least 2^-7, and below 10^19 is below 2^64, so the biased exponent is
// between about 1016 and 1087 — far from subnormal (0) and infinity (2047).
// The one carry left to handle, rounding up to 2^53, needs no test either:
// the result is assembled as the mantissa plus (retExp2 - 1) << 52, the
// mantissa's own bit 52 supplying the exponent's last 1, so a mantissa of
// exactly 2^53 carries into the exponent field and gives 2^52 at retExp2 + 1,
// the value eiselLemire64's shift and increment produce.
#define LONGCONV(el, norefine, round, decline, done) \
	SHLQ    $5, AX \
	ADDQ    R14, AX \
	SHLQ    $5, AX \ // the (L1, L2) template: 32 bytes at 4608+32*(32*L1+L2)
	VMOVDQU (SI)(CX*1), Y13 \ // the 32 bytes from the first digit
	VMOVDQU 4608(R15)(AX*1), Y12 \
	VPERMI2B Y4, Y13, Y12 \ // the digits, right-aligned, from those or '0'
	VPSUBB  Y4, Y12, Y12 \
	VPMADDUBSW Y8, Y12, Y12 \
	VPMADDWD Y9, Y12, Y12 \
	VPACKUSDW Y12, Y12, Y12 \
	VPMADDWD Y10, Y12, Y12 \ // dwords: hi8, mid8 | lo8
	VMOVQ   X12, AX \
	VEXTRACTI128 $1, Y12, X13 \
	VMOVD   X13, DX \ // lo8
	MOVL    AX, CX \ // hi8
	SHRQ    $32, AX \ // mid8
	IMUL3Q  $100000000, CX, CX \
	ADDQ    AX, CX \
	IMUL3Q  $100000000, CX, CX \
	ADDQ    DX, CX \ // the mantissa, below 10^19
	MOVQ    CX, AX \
	SHRQ    $53, AX \
	JNZ     el \
	VCVTSI2SDQ CX, X13, X13 \
	SHLQ    $5, DI \
	ADDQ    R14, DI \
	VDIVSD  4096(R15)(DI*8), X13, X13 \
	VMOVSD  X13, (R8)(R10*8) \
	INCQ    R10 \
	JMP     done \
el: \
	BSRQ    CX, AX \
	XORQ    $63, AX \ // leading zeros
	SHLXQ   AX, CX, CX \ // normalized: the top bit set
	MOVQ    25088+160(R15)(R14*8), DX \ // the biased binary exponent estimate, less 2
	SUBQ    AX, DX \
	MOVQ    25088(R15)(R14*8), AX \ // the high word of 10^-L2
	MOVQ    DX, R14 \
	MULQ    CX \ // DX:AX = xHi:xLo
	CMPB    DX, $0xFF \ // the low nine bits of xHi all ones: one product in 512,
	JNE     norefine \ // so this is tested first — whether xLo + man wraps is a
	BTL     $8, DX \ // coin toss, and as the first branch it mispredicted on
	JCC     norefine \ // every other long number
	ADDQ    AX, CX \ // CF: xLo + man wraps (man is not needed again)
	JCS     decline \ // eiselLemire64 would refine with the low word
norefine: \
	MOVQ    DX, CX \
	SHRQ    $63, CX \ // msb
	ADDQ    CX, R14 \ // estimate - lz - 2 + msb: eiselLemire64's retExp2, after its -= 1^msb, less 1
	ADDQ    $9, CX \
	SHRXQ   CX, DX, CX \ // retMantissa = xHi >> (msb + 9): 54 bits
	TESTQ   AX, AX \
	JNZ     round \
	TESTL   $0x1FF, DX \
	JNZ     round \
	MOVL    CX, AX \
	ANDL    $3, AX \
	CMPL    AX, $1 \
	JEQ     decline \ // exactly halfway
round: \
	SHRQ    $1, CX \
	ADCQ    $0, CX \ // round to 53 bits: (m >> 1) + (m & 1) is (m + (m & 1)) >> 1
	SHLQ    $52, R14 \
	ADDQ    R14, CX \ // the mantissa's bit 52 (or a rounding carry) completes the exponent
	SHLQ    $63, DI \
	ORQ     DI, CX \
	MOVQ    CX, (R8)(R10*8) \
	INCQ    R10

// func parseFloatRunVBMI(data []byte, i int, out []float64) (n, p, closed int)
//
// parseFloatRunAVX2's contract and window walk, for CPUs with AVX-512 VBMI,
// taking numbers of up to 19 digits: the 16- and 17-digit coordinates of
// canada, large-json and mesh_pretty, which the AVX2 body hands back. It is
// reached only by parseFloatRunAVX2's tail jump.
//
// The digits are gathered by one VPERMI2B from the 32 bytes at the number's
// first digit and a register of '0' bytes (an index with bit 5 set):
// right-aligned into 24 lanes with the '.' dropped, under a per-(L1, L2)
// template that depends on nothing else, so the gather waits only on the load.
// (Gathering from the window register instead needs the element's lane added
// to the template — a broadcast from a general register on the chain — and a
// 64-lane permute.) The
// three eight-digit groups fold as in the AVX2 body and combine in a general
// register — at most 19 digits is below 2^64 — and then:
//
//   - a mantissa below 2^53 is Clinger's fast path, exactly as in the AVX2
//     body: one convert and one correctly rounded divide by an exact ±10^L2;
//   - anything larger is Eisel-Lemire, eiselLemire64 transcribed: normalize,
//     one 64x64 multiply by the high word of the 128-bit power, round to 53
//     bits. Where eiselLemire64 would refine with the low word (the product's
//     low nine bits all ones and xLo + man wrapping: one canada coordinate in
//     about 750, but far more often for a value exact in binary, whose
//     truncated product lands just under it — see TestFloatRunEiselLemire) or
//     decline (an exact halfway value, an exponent out of range),
//     this declines instead, and the element goes to the scalar loop — which
//     runs eiselLemire64 and strconv — so every value this writes is the one
//     they return.
//
// A number of fifteen digits or fewer takes the AVX2 body's conversion
// instead. The gather loads 32 bytes at the first digit, so the window needs
// 96 bytes. Registers as in the AVX2 body; R15 is floatRunTab (the
// AVX2 controls, the divisors at +4096, the templates at +4608, the
// Eisel-Lemire entries at +25088); Z4-Z7 hold the splats, Y8-Y10 the fold
// weights and X11 the join weights.
TEXT ·parseFloatRunVBMI(SB), NOSPLIT, $0-80
	MOVQ    data_base+0(FP), SI
	MOVQ    i+24(FP), CX
	MOVQ    out_base+32(FP), R8
	XORQ    R10, R10
	ADDQ    CX, SI
	XORL    CX, CX
	VPBROADCASTB fvZero<>(SB), Z4
	VPBROADCASTB fvNine<>(SB), Z5
	VPBROADCASTB fvComma<>(SB), Z6
	VPBROADCASTB fvSpace<>(SB), Z7
	VPBROADCASTB fvMinus<>(SB), Z15
	VBROADCASTI128 frW10<>(SB), Y8
	VBROADCASTI128 frW100<>(SB), Y9
	VBROADCASTI128 frW1e4<>(SB), Y10
	VMOVDQU frJoin<>(SB), X11
	LEAQ    ·floatRunTab(SB), R15

vwindow:
	MOVQ    data_base+0(FP), AX
	ADDQ    data_len+8(FP), AX
	SUBQ    $96, AX                 // the last window start with 96 bytes
	CMPQ    SI, AX
	JHI     vstop
	VMOVDQU64 (SI), Z0
	VPSUBB    Z4, Z0, Z1
	VPCMPUB   $6, Z5, Z1, K1        // (c-'0') > 9: not a digit
	VPCMPEQB  Z6, Z0, K2            // a comma
	VPCMPUB   $6, Z7, Z0, K3        // c > 0x20: not whitespace
	VPCMPEQB  Z15, Z0, K5           // a '-'
	KMOVQ     K1, R11
	KMOVQ     K2, R12
	KMOVQ     K3, R13
	KMOVQ     K5, R9
	MOVQ    $-1, AX
	SHLXQ   CX, AX, AX
	ANDQ    AX, R12                 // the commas at or after b
	JZ      vnoComma

velem:
	TZCNTQ  R12, BX                 // c: this element's comma
	VPARSE(vstopElem, velemNoFrac, velemParsed)
	CMPQ    DX, BX
	JNE     vwsBeforeComma

vfold:
	// AX = L1, R14 = L2, CX = s, DI = the sign, BX = the element's comma.
	CMPQ    R10, out_len+40(FP)
	JAE     vstopElem               // out is full: leave this element to the caller
	LEAQ    (AX)(R14*1), DX         // L1 + L2
	CMPQ    DX, $15
	JA      vlong
	// Fifteen digits or fewer: exactly the AVX2 body's conversion and step,
	// which on these is cheaper than the gather below — a 16-byte load and one
	// PSHUFB against a broadcast, a variable permute and two trips through a
	// general register — so an array of short numbers costs this body only
	// the length test above.
	CONVERT
	LEAQ    1(BX), CX
	BLSRQ   R12, R12
	JNZ     velem
	JMP     vnoComma

vlastFold:
	// The array's last element: BX = its ']'.
	CMPQ    R10, out_len+40(FP)
	JAE     vstopElem
	LEAQ    (AX)(R14*1), DX
	CMPQ    DX, $15
	JA      vlong
	CONVERT
	JMP     vclosed

vlong:
	CMPQ    DX, $19
	JA      vstopElem               // 20 digits or more: the scalar loop's
	MOVQ    CX, DX
	SUBQ    DI, DX
	MOVQ    DX, (R8)(R10*8)         // parked for vdecline: the element's first byte
	LONGCONV(vel, velNoRefine, velRound, vdecline, vconverted)

vconverted:
	// After a long element the delimiter says whether the walk goes on (a
	// comma) or the array has ended here (its ']').
	CMPB    (SI)(BX*1), $0x2c
	JNE     vclosed
	LEAQ    1(BX), CX               // b: just past the comma
	BLSRQ   R12, R12                // the comma is consumed; ZF: none left
	JNZ     velem

vnoComma:
	CMPQ     CX, $48
	JAE      vstep
	VPBROADCASTB fvClose<>(SB), Z2
	VPCMPEQB Z2, Z0, K4
	KMOVQ    K4, BX
	MOVQ     $-1, AX
	SHLXQ    CX, AX, AX
	ANDQ     AX, BX                 // the ']'s at or after b
	JZ       vrestart
	TZCNTQ   BX, BX                 // r: the ']'
	VPARSE(vstopElem, vlastNoFrac, vlastParsed)
	CMPQ     DX, BX
	JEQ      vlastFold
	SHRXQ    DX, R13, AX            // whitespace, then the ']'
	TZCNTQ   AX, AX
	ADDQ     DX, AX
	CMPQ     AX, BX
	JNE      vstopElem
	SHRXQ    CX, R11, AX            // L1 again
	TZCNTQ   AX, AX
	JMP      vlastFold

vwsBeforeComma:
	// e < c: whitespace to the comma, or whitespace to the array's ']' (the
	// comma being the enclosing container's).
	SHRXQ   DX, R13, AX
	TZCNTQ  AX, AX
	ADDQ    DX, AX
	CMPQ    AX, BX
	JNE     vwsClose
	SHRXQ   CX, R11, AX             // L1 again
	TZCNTQ  AX, AX
	JMP     vfold
vwsClose:
	CMPB    (SI)(AX*1), $0x5d
	JNE     vstopElem
	MOVQ    AX, BX                  // the ']' delimits this, the last element
	SHRXQ   CX, R11, AX             // L1 again
	TZCNTQ  AX, AX
	JMP     vlastFold

vclosed:
	ADDQ    BX, SI
	SUBQ    data_base+0(FP), SI
	MOVQ    R10, n+56(FP)
	MOVQ    SI, p+64(FP)
	MOVQ    $1, closed+72(FP)
	VZEROUPPER
	RET

vrestart:
	SHRXQ   CX, R13, AX
	TESTQ   AX, AX
	JZ      vskipBlank
	TZCNTQ  AX, AX
	ADDQ    AX, CX
	TESTQ   CX, CX
	JZ      vstop
	ADDQ    CX, SI
	XORL    CX, CX
	JMP     vwindow

vskipBlank:
	ADDQ    $64, SI
	XORL    CX, CX
	JMP     vwindow

vstep:
	ADDQ    $48, SI
	SUBQ    $48, CX
	JMP     vwindow

vdecline:
	// Eisel-Lemire declined the element: hand it back from its first byte,
	// parked in its output slot.
	MOVQ    (R8)(R10*8), CX
	JMP     vstop

vstopElem:
	// A stop at an element VPARSE measured: CX is past its '-', if any.
	SUBQ    DI, CX

vstop:
	VZEROUPPER
	ADDQ    CX, SI
	SUBQ    data_base+0(FP), SI
	MOVQ    R10, n+56(FP)
	MOVQ    SI, p+64(FP)
	MOVQ    $0, closed+72(FP)
	RET

// func parseFloatPointsVBMI(data []byte, i int, out []float64, n int) (np, p, closed int)
//
// The VBMI body's number conversion over an array of fixed-size numeric
// arrays — a ring of coordinate points, "[[x,y],[x,y],…]" — for
// DecodeFloat64Points. i is at a point's '[' (or whitespace before it); out is
// the flat backing of the points' spare slots, n floats a point. It converts as
// many points as it can, each exactly n numbers of this body's shape, writing
// the i-th point's numbers at out[n*i:], and returns how many (np), and either
// closed = 1 with p at the ring's ']', or p at the start of the first point it
// did not take — which the reader then decodes element by element. A point is
// counted only once the ',' after it has been seen (or the ring's ']'), so p is
// always a place the reader resumes by reading an element; the separator is
// looked for past the point's window only while the bytes between are
// whitespace, and a point whose separator is still not found is handed back
// whole.
//
// One window per point, at the point: a coordinate is at most 20 bytes, so a
// point of up to three of them and its brackets fits in 64 bytes, and the
// window's address — the byte after the previous point's separator — is known
// as soon as that separator is found, well before the previous point's
// conversions finish. Each number is delimited exactly as in the flat walk: the
// inner ones by their commas, the last by the point's ']' (found with a ']'
// mask kept in K4), and its region must be whitespace, the number, whitespace.
// A point of fewer or more numbers fails that on its own (a ']' or a ',' where
// whitespace must be) and is handed back.
//
// Registers as in parseFloatRunVBMI; the frame holds the numbers left in the
// point and the point's start.
TEXT ·parseFloatPointsVBMI(SB), NOSPLIT, $48-88
	MOVQ    data_base+0(FP), SI
	MOVQ    data_len+8(FP), AX
	LEAQ    -96(SI)(AX*1), AX       // the last window start with 96 bytes: the gather
	MOVQ    AX, plimit-24(SP)       // reads 32 bytes at any lane
	ADDQ    i+24(FP), SI
	MOVQ    out_base+32(FP), R8
	MOVQ    out_len+40(FP), DX
	MOVQ    n+56(FP), AX
	SUBQ    AX, DX
	MOVQ    DX, proom-48(SP)        // the last slot a whole point can start at
	DECQ    AX
	MOVQ    AX, pcommas-40(SP)      // the commas inside a point
	XORQ    R10, R10
	VPBROADCASTB fvZero<>(SB), Z4
	VPBROADCASTB fvNine<>(SB), Z5
	VPBROADCASTB fvComma<>(SB), Z6
	VPBROADCASTB fvSpace<>(SB), Z7
	VPBROADCASTB fvClose<>(SB), Z14 // not Z3: CONVERT writes X3, which clears it
	VPBROADCASTB fvMinus<>(SB), Z15
	VBROADCASTI128 frW10<>(SB), Y8
	VBROADCASTI128 frW100<>(SB), Y9
	VBROADCASTI128 frW1e4<>(SB), Y10
	VMOVDQU frJoin<>(SB), X11
	LEAQ    ·floatRunTab(SB), R15

ppoint:
	MOVQ    SI, pstart-8(SP)        // where the point's region starts
	MOVQ    R10, pr10-32(SP)        // and what was written before it
	CMPQ    SI, plimit-24(SP)
	JHI     pstop
	CMPQ    R10, proom-48(SP)
	JGT     pstop                   // no room for a whole point
	VMOVDQU64 (SI), Z0
	VPSUBB    Z4, Z0, Z1
	VPCMPUB   $6, Z5, Z1, K1        // not a digit
	VPCMPEQB  Z6, Z0, K2            // a comma
	VPCMPUB   $6, Z7, Z0, K3        // not whitespace
	VPCMPEQB  Z14, Z0, K4           // a ']'
	VPCMPEQB  Z15, Z0, K5           // a '-'
	KMOVQ     K1, R11
	KMOVQ     K2, R12
	KMOVQ     K3, R13
	KMOVQ     K5, R9
	TZCNTQ    R13, CX               // the first byte that is not whitespace
	CMPQ      CX, $16
	JAE       prewindow             // far into the window: start one at it
	CMPB      (SI)(CX*1), $0x5b
	JNE       pstop                 // not a point ('[')
	INCQ      CX                    // b: past the '['
	// Everything before b is whitespace and the '[', so the window's first
	// ']' is the point's, and its commas are the ones below that.
	KMOVQ     K4, BX
	TZCNTQ    BX, BX
	JCS       plong2                // the point's ']' is past the window
	MOVQ      BX, pclose-16(SP)
	BZHIQ     BX, R12, R12          // the point's commas
	POPCNTQ   R12, AX
	CMPQ      AX, pcommas-40(SP)
	JNE       pback                 // not n numbers (or not a flat point)

pelem:
	TZCNTQ  R12, BX                 // an inner number's delimiter: its comma
	JCC     pparse
	MOVQ    pclose-16(SP), BX       // the last number's: the point's ']'
pparse:
	VPARSE(pback, pNoFrac, pParsed)
	CMPQ    DX, BX
	JEQ     pfold
	SHRXQ   DX, R13, AX             // whitespace to the delimiter
	TZCNTQ  AX, AX
	ADDQ    DX, AX
	CMPQ    AX, BX
	JNE     pback
	SHRXQ   CX, R11, AX             // L1 again
	TZCNTQ  AX, AX
pfold:
	LEAQ    (AX)(R14*1), DX
	CMPQ    DX, $15
	JA      plong
	CONVERT
	JMP     pconverted
plong:
	CMPQ    DX, $19
	JA      pback                   // 20 digits or more: the reader's
	LONGCONV(pel, pNoRefine, pRound, pback, pconverted)

pconverted:
	LEAQ    1(BX), CX               // past the delimiter
	BLSRQ   R12, R12                // the comma is consumed; CF: there was none
	JCC     pelem                   // left, so that was the point's ']'

pdone:
	// The point's ']' is at BX. The first byte after it that is not
	// whitespace must be the ',' before the next point or the ring's ']'.
	CMPQ    CX, $64
	JEQ     psep
	SHRXQ   CX, R13, AX
	TZCNTQ  AX, AX
	JCS     psep                    // the separator is past the window
	ADDQ    CX, AX
psepAt:
	CMPB    (SI)(AX*1), $0x2c
	JEQ     pnext
	CMPB    (SI)(AX*1), $0x5d
	JNE     pback
	// The ring's ']': done, and closed.
	ADDQ    AX, SI
	SUBQ    data_base+0(FP), SI
	MOVQ    R10, AX
	XORQ    DX, DX
	DIVQ    n+56(FP)
	MOVQ    AX, np+64(FP)
	MOVQ    SI, p+72(FP)
	MOVQ    $1, closed+80(FP)
	VZEROUPPER
	RET

pnext:
	LEAQ    1(SI)(AX*1), SI         // the next point's region starts after the ','
	JMP     ppoint

psep:
	// The rest of the window is whitespace — in a pretty-printed ring, the
	// last point's ']' is on the next line — so the separator is looked for
	// in the next window, and the one after that while they hold nothing
	// else. The point stays uncounted until it is found (pback forgets it).
	ADDQ    $64, SI
	MOVQ    plimit-24(SP), AX
	ADDQ    $32, AX                 // a window read whole: 64 bytes, not 96
	CMPQ    SI, AX
	JHI     pback
	VMOVDQU64 (SI), Z0
	VPCMPUB   $6, Z7, Z0, K3        // not whitespace
	KMOVQ     K3, AX
	TZCNTQ    AX, AX
	JCS       psep                  // all of it whitespace
	JMP       psepAt

plong2:
	// The point's ']' is past the window: when whitespace before its '['
	// took some of the window, start one at the '[' and try once more (at
	// lane 0 there is nothing left to gain, and the point is handed back).
	CMPQ    CX, $1
	JEQ     pback
	DECQ    CX

prewindow:
	// Sixteen or more bytes of whitespace before the point (64: all of the
	// window), or a point that did not fit behind less: start the next window
	// at the first byte that is not whitespace.
	ADDQ    CX, SI
	JMP     ppoint

pback:
	// The point cannot be taken here: forget its numbers already written —
	// all n of them when it is its separator that failed — and hand it back
	// from its start.
	MOVQ    pr10-32(SP), R10

pstop:
	VZEROUPPER
	MOVQ    pstart-8(SP), SI
	SUBQ    data_base+0(FP), SI
	MOVQ    R10, AX
	XORQ    DX, DX
	DIVQ    n+56(FP)
	MOVQ    AX, np+64(FP)
	MOVQ    SI, p+72(FP)
	MOVQ    $0, closed+80(FP)
	RET
