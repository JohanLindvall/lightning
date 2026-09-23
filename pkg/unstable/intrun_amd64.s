#include "textflag.h"

// Constants for parseIntRunAVX2 (see intrun_amd64.go).
DATA irZero<>+0(SB)/8, $0x3030303030303030 // '0'
DATA irZero<>+8(SB)/8, $0x3030303030303030
DATA irZero<>+16(SB)/8, $0x3030303030303030
DATA irZero<>+24(SB)/8, $0x3030303030303030
GLOBL irZero<>(SB), RODATA|NOPTR, $32

DATA irNine<>+0(SB)/8, $0x0909090909090909 // 9: (c-'0') <= 9 via VPMINUB
DATA irNine<>+8(SB)/8, $0x0909090909090909
DATA irNine<>+16(SB)/8, $0x0909090909090909
DATA irNine<>+24(SB)/8, $0x0909090909090909
GLOBL irNine<>(SB), RODATA|NOPTR, $32

DATA irComma<>+0(SB)/8, $0x2c2c2c2c2c2c2c2c // ','
DATA irComma<>+8(SB)/8, $0x2c2c2c2c2c2c2c2c
DATA irComma<>+16(SB)/8, $0x2c2c2c2c2c2c2c2c
DATA irComma<>+24(SB)/8, $0x2c2c2c2c2c2c2c2c
GLOBL irComma<>(SB), RODATA|NOPTR, $32

DATA irSpace<>+0(SB)/8, $0x2020202020202020 // 0x20: c <= 0x20 via VPMINUB (SkipWS's rule)
DATA irSpace<>+8(SB)/8, $0x2020202020202020
DATA irSpace<>+16(SB)/8, $0x2020202020202020
DATA irSpace<>+24(SB)/8, $0x2020202020202020
GLOBL irSpace<>(SB), RODATA|NOPTR, $32

// The three fold weights. Bytes (10,1) pair digits into two-digit words,
// words (100,1) pair those into four-digit dwords, and after PACKUSDW words
// (10000,1) pair those into the eight-digit value.
DATA irW10<>+0(SB)/8, $0x010a010a010a010a
DATA irW10<>+8(SB)/8, $0x010a010a010a010a
GLOBL irW10<>(SB), RODATA|NOPTR, $16

DATA irW100<>+0(SB)/8, $0x0001006400010064
DATA irW100<>+8(SB)/8, $0x0001006400010064
GLOBL irW100<>(SB), RODATA|NOPTR, $16

DATA irW1e4<>+0(SB)/8, $0x0001271000012710
DATA irW1e4<>+8(SB)/8, $0x0001271000012710
GLOBL irW1e4<>(SB), RODATA|NOPTR, $16

// irCtrl holds the fold's PSHUFB control for each digit count L (entry L at
// byte 16*L): the eight bytes loaded from the digit start hold the L digits in
// their first L lanes, and the control moves them to the top of the low eight
// lanes with zeros in front (0x80 zeroes a lane) and zeroes the upper eight, so
// the fixed-weight folds read an eight-digit number with leading zeros and the
// last fold's upper dword is zero — the low quadword IS the int64, stored with
// one VMOVQ.
DATA irCtrl<>+0(SB)/8, $0x8080808080808080
DATA irCtrl<>+8(SB)/8, $0x8080808080808080
DATA irCtrl<>+16(SB)/8, $0x0080808080808080
DATA irCtrl<>+24(SB)/8, $0x8080808080808080
DATA irCtrl<>+32(SB)/8, $0x0100808080808080
DATA irCtrl<>+40(SB)/8, $0x8080808080808080
DATA irCtrl<>+48(SB)/8, $0x0201008080808080
DATA irCtrl<>+56(SB)/8, $0x8080808080808080
DATA irCtrl<>+64(SB)/8, $0x0302010080808080
DATA irCtrl<>+72(SB)/8, $0x8080808080808080
DATA irCtrl<>+80(SB)/8, $0x0403020100808080
DATA irCtrl<>+88(SB)/8, $0x8080808080808080
DATA irCtrl<>+96(SB)/8, $0x0504030201008080
DATA irCtrl<>+104(SB)/8, $0x8080808080808080
DATA irCtrl<>+112(SB)/8, $0x0605040302010080
DATA irCtrl<>+120(SB)/8, $0x8080808080808080
DATA irCtrl<>+128(SB)/8, $0x0706050403020100
DATA irCtrl<>+136(SB)/8, $0x8080808080808080
GLOBL irCtrl<>(SB), RODATA|NOPTR, $144

// CLASSIFY builds the window's three 64-bit lane masks from Y0 (bytes 0-31)
// and Y1 (32-63): R11 = NOT a digit, R12 = a comma, R13 = NOT whitespace
// (<= 0x20). The two NOT masks are what the walk counts trailing zeros of.
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

// FOLD parses the L digits (AX = L, 1..8) starting at window lane CX and
// stores the value at out[R10]; AX is clobbered.
#define FOLD \
	VMOVQ      (SI)(CX*1), X2   \
	VPSUBB     X4, X2, X2       \
	SHLL       $4, AX           \
	VPSHUFB    (R15)(AX*1), X2, X2 \
	VPMADDUBSW X8, X2, X2       \
	VPMADDWD   X9, X2, X2       \
	VPACKUSDW  X2, X2, X2       \
	VPMADDWD   X10, X2, X2      \
	VMOVQ      X2, (R8)(R10*8)  \
	INCQ       R10

// func parseIntRunAVX2(data []byte, i int, out []int64) (n, p, closed int)
//
// Parses as many "ws* digits{1..8} ws* ','" groups as it can from data[i:],
// one int64 per group into out, and also the array's last element when it is
// terminated by ']' (closed = 1, p at the ']'). Stops, with p at the start of
// the unconsumed element (its leading whitespace possibly skipped) or right
// after the last consumed comma, when: fewer than 72 bytes remain from the
// window, out is full, or an element is not shaped like that (a sign, null, a
// fraction or exponent, 9+ digits, a missing comma, an empty element). Every
// stop position is a state the scalar loop resumes from, so the values and
// every error come from the same code as before.
//
// The walk is over 64-byte WINDOWS stepped at a fixed 48 bytes, the shape the
// arm64 kernel took for the same reason: the previous form reclassified a
// 16-byte block at whatever element straddled its end, which put every block's
// load on the cursor's dependency chain and cost mesh's ", "-separated
// four-digit elements 14.5 cycles each (mesh_pretty's newline-and-indent ones
// straddled on every element). A fixed stride makes the next window's address
// independent of the walk. Elements are consumed while their comma is in the
// window; if the last one consumed ends in the window's final 16 bytes the
// window steps by 48, and otherwise (a region longer than 16 bytes, or the
// array's end) the element is finished in place or the window restarts at it.
//
// Within a window the commas are walked with TZCNT/BLSR, so the loop-carried
// chain is one BLSR — whose zero flag also closes the loop — and each
// element's digit count and whitespace checks hang off the comma position,
// not off the previous element's digits. The fold loads
// the element's eight bytes from memory at its first digit, so it needs no
// window register either.
//
// Registers: R14 data base, DI the last window start with 72 bytes, SI the
// window, CX the cursor b (then s) in it, R8/R9 out base/len, R10 values
// written, R11/R12/R13 the window masks, R15 irCtrl, BX the comma, DX e, AX
// temporaries. Y4-Y7 and X8-X10 hold the constants.
TEXT ·parseIntRunAVX2(SB), NOSPLIT, $0-80
	MOVQ    data_base+0(FP), R14
	MOVQ    data_len+8(FP), DI
	MOVQ    i+24(FP), CX
	MOVQ    out_base+32(FP), R8
	MOVQ    out_len+40(FP), R9
	XORQ    R10, R10
	LEAQ    (R14)(CX*1), SI
	XORL    CX, CX
	LEAQ    -72(R14)(DI*1), DI
	CMPQ    SI, DI
	JHI     done_novz               // under 72 bytes: nothing to do here
	VMOVDQU irZero<>(SB), Y4
	VMOVDQU irNine<>(SB), Y5
	VMOVDQU irComma<>(SB), Y6
	VMOVDQU irSpace<>(SB), Y7
	VMOVDQU irW10<>(SB), X8
	VMOVDQU irW100<>(SB), X9
	VMOVDQU irW1e4<>(SB), X10
	LEAQ    irCtrl<>(SB), R15

window:
	CMPQ    SI, DI
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
	SHRXQ   CX, R13, AX
	TZCNTQ  AX, AX                  // whitespace run from b (bounded by the comma)
	ADDQ    AX, CX                  // s
	SHRXQ   CX, R11, AX
	TZCNTQ  AX, AX                  // L: digits from s
	LEAQ    -1(AX), DX
	CMPQ    DX, $7
	JA      stop                    // no digit at s, or 9+ of them
	LEAQ    (CX)(AX*1), DX          // e
	CMPQ    DX, BX
	JNE     wsBeforeComma
fold:
	CMPQ    R10, R9
	JAE     stop                    // out is full: leave this element to the caller
	FOLD
	LEAQ    1(BX), CX               // b: just past the comma
	BLSRQ   R12, R12                // the comma is consumed; ZF: none left
	JNZ     elem

noComma:
	// No comma left in the window. If the cursor is in its last 16 bytes the
	// window steps by 48 (the element at b is then wholly in the next one);
	// otherwise this is the array's last element, finished in place when its
	// ']' is in the window, or an element region too long for the stride,
	// which restarts the window at it.
	CMPQ    CX, $48
	JAE     step
	BTQ     CX, R13
	JC      lastDigits
	SHRXQ   CX, R13, AX
	TESTQ   AX, AX                  // SHRX sets no flags
	JZ      restart                 // whitespace to the window's end
	TZCNTQ  AX, AX
	ADDQ    AX, CX                  // s
lastDigits:
	SHRXQ   CX, R11, AX
	TESTQ   AX, AX
	JZ      restart                 // digits to the window's end
	TZCNTQ  AX, AX
	LEAQ    -1(AX), DX
	CMPQ    DX, $7
	JA      stop
	LEAQ    (CX)(AX*1), DX          // e
	SHRXQ   DX, R13, BX
	TESTQ   BX, BX
	JZ      restart                 // whitespace after the digits to the window's end
	TZCNTQ  BX, BX
	ADDQ    DX, BX                  // the byte after the element
	MOVBLZX (SI)(BX*1), DX
	CMPL    DX, $0x5d
	JNE     stop                    // not ']': not this kernel's element
	CMPQ    R10, R9
	JAE     stop
	FOLD
	ADDQ    BX, SI
	SUBQ    R14, SI
	MOVQ    R10, n+56(FP)
	MOVQ    SI, p+64(FP)
	MOVQ    $1, closed+72(FP)
	VZEROUPPER
	RET

wsBeforeComma:
	// e < c: the bytes between must all be whitespace — or the first one that
	// is not must be the array's ']', the comma being the enclosing
	// container's (an array is followed by more of the document, often in the
	// same window), and then this element is the last; handing it back cost
	// every short array its final element.
	SHRXQ   DX, R13, AX
	TZCNTQ  AX, AX
	ADDQ    DX, AX
	CMPQ    AX, BX
	JNE     closedBefore
	MOVQ    DX, AX
	SUBQ    CX, AX                  // L again
	JMP     fold
closedBefore:
	CMPB    (SI)(AX*1), $0x5d
	JNE     stop
	MOVQ    AX, BX                  // the ']'
	CMPQ    R10, R9
	JAE     stop
	MOVQ    DX, AX
	SUBQ    CX, AX                  // L again
	FOLD
	ADDQ    BX, SI
	SUBQ    R14, SI
	MOVQ    R10, n+56(FP)
	MOVQ    SI, p+64(FP)
	MOVQ    $1, closed+72(FP)
	VZEROUPPER
	RET

restart:
	// The element at CX runs past the window: start the next window at it.
	// At CX == 0 that would not move, so stop.
	TESTQ   CX, CX
	JZ      stop
	ADDQ    CX, SI
	XORL    CX, CX
	JMP     window

step:
	ADDQ    $48, SI
	SUBQ    $48, CX
	JMP     window

stop:
	VZEROUPPER
done_novz:
	ADDQ    CX, SI
	SUBQ    R14, SI
	MOVQ    R10, n+56(FP)
	MOVQ    SI, p+64(FP)
	MOVQ    $0, closed+72(FP)
	RET
