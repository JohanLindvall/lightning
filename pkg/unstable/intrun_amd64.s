#include "textflag.h"

// Constants for parseIntRunSSE (see intrun_amd64.go).
DATA irZero<>+0(SB)/8, $0x3030303030303030 // '0' splat
DATA irZero<>+8(SB)/8, $0x3030303030303030
GLOBL irZero<>(SB), RODATA|NOPTR, $16

DATA irNine<>+0(SB)/8, $0x0909090909090909 // 9 splat: (c-'0') <= 9 via PMINUB
DATA irNine<>+8(SB)/8, $0x0909090909090909
GLOBL irNine<>(SB), RODATA|NOPTR, $16

DATA irComma<>+0(SB)/8, $0x2c2c2c2c2c2c2c2c // ',' splat
DATA irComma<>+8(SB)/8, $0x2c2c2c2c2c2c2c2c
GLOBL irComma<>(SB), RODATA|NOPTR, $16

DATA irClose<>+0(SB)/8, $0x5d5d5d5d5d5d5d5d // ']' splat
DATA irClose<>+8(SB)/8, $0x5d5d5d5d5d5d5d5d
GLOBL irClose<>(SB), RODATA|NOPTR, $16

DATA irSpace<>+0(SB)/8, $0x2020202020202020 // 0x20 splat: c <= 0x20 via PMINUB (SkipWS's rule)
DATA irSpace<>+8(SB)/8, $0x2020202020202020
GLOBL irSpace<>(SB), RODATA|NOPTR, $16

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

// func parseIntRunSSE(data []byte, i int, out []int64) (n, p, closed int)
//
// Parses as many "ws* digits{1..8} ws* ','" groups as it can from data[i:],
// one int64 per group into out, and also the array's last element when it
// is terminated by ']' (closed = 1, p at the ']'). Stops, with p at the start
// of the unconsumed element (its leading whitespace skipped) or right after
// the last consumed comma, when: fewer than 16 bytes remain, out is full, an
// element is not shaped like that (a sign, null, a fraction or exponent, 9+
// digits, a missing comma), or an element and its whitespace straddle the
// block beginning at p (the block cannot move forward without consuming
// something). Every stop position is a state the scalar loop resumes from,
// so the values and every error come from the same code as before.
//
// Registers: SI data base, DI s (absolute block start), R8 out base, R9 out
// len, R10 j (values written), CX b (offset in the block), R11 digit mask,
// R12 the commas not yet consumed, R13 whitespace mask, R14 shuffle table,
// AX/BX/DX/R15 temporaries. X1 holds the block's bytes minus '0' for the
// whole block; X4-X11 the constants.
//
// Two things this loop had to learn on Zen 4, both from the counters (see
// CLAUDE.md): the shifts are SHRX followed by an explicit TEST, never a
// shift by CL whose flags are then branched on — a variable-count shift
// writes its flags only conditionally, and reading them costs a decoder-fed
// sequence (de_src_op_disp.decoder counted seven ops per element, and the
// first version ran 38 cycles an element against the scalar loop's 21 with
// a third of the instructions); and each element's comma comes from the
// block's comma mask (BLSR/TZCNT), not from the digit count, so the
// element-to-element chain is two cycles rather than shift, count, add,
// shift, count. The encodings are VEX.128 for the three-operand forms;
// legacy SSE measured the same once the shift was fixed.
TEXT ·parseIntRunSSE(SB), NOSPLIT, $0-80
	MOVQ data_base+0(FP), SI
	MOVQ i+24(FP), DI
	MOVQ out_base+32(FP), R8
	MOVQ out_len+40(FP), R9
	XORQ R10, R10
	VMOVDQU irZero<>(SB), X4
	VMOVDQU irNine<>(SB), X5
	VMOVDQU irComma<>(SB), X6
	VMOVDQU irSpace<>(SB), X7
	VMOVDQU irW10<>(SB), X8
	VMOVDQU irW100<>(SB), X9
	VMOVDQU irW1e4<>(SB), X10
	VMOVDQU irClose<>(SB), X11
	LEAQ ·intRunShuffle(SB), R14

block:
	XORL CX, CX
	MOVQ data_len+8(FP), AX
	SUBQ $16, AX
	CMPQ DI, AX
	JGT  done // fewer than 16 bytes from s: the scalar loop takes it from here
	VMOVDQU (SI)(DI*1), X0
	VPSUBB X4, X0, X1 // X1 = c - '0'
	VPMINUB X5, X1, X2
	VPCMPEQB X1, X2, X2
	VPMOVMSKB X2, R11 // digits: min(c-'0', 9) == c-'0'
	VPCMPEQB X6, X0, X2
	VPMOVMSKB X2, R12 // commas; consumed with BLSR as elements complete
	VPMINUB X7, X0, X2
	VPCMPEQB X0, X2, X2
	VPMOVMSKB X2, R13 // whitespace: min(c, 0x20) == c

elem:
	// Leading whitespace, only when the byte at b is whitespace (compact
	// input never enters the arithmetic): s = b + trailing zeros of ~W >> b.
	BTL  CX, R13
	JNC  digits
	MOVL R13, AX
	XORL $0xffff, AX
	SHRXL CX, AX, AX
	TESTL AX, AX
	JZ   straddle
	TZCNTL AX, AX
	ADDL AX, CX

digits:
	// The digit run at s: L = trailing zeros of ~D >> s, must be 1..8.
	MOVL R11, AX
	XORL $0xffff, AX
	SHRXL CX, AX, AX
	TESTL AX, AX
	JZ   straddle
	TZCNTL AX, AX
	LEAL -1(AX), DX
	CMPL DX, $7
	JA   done // no digit here, or 9+ of them
	LEAL (CX)(AX*1), DX // e: the byte after the digits (<= 15)
	// The element's comma is the lowest one left in the mask, found
	// without the cursor: this is what keeps the element-to-element chain
	// at a BLSR and a TZCNT rather than shift, count, add, shift, count.
	TESTL R12, R12
	JZ   noComma
	TZCNTL R12, BX // c
	CMPL BX, DX
	JE   fold // ',' right after the digits: the common case
	// c > e (a comma cannot fall inside the digit run, and any comma before
	// s would have ended the whitespace skip with L == 0): the bytes e..c-1
	// must all be whitespace, i.e. the whitespace run from e reaches c.
	MOVL R13, R15
	XORL $0xffff, R15
	SHRXL DX, R15, R15
	TZCNTL R15, R15 // 32 when ~W >> e is zero, which also reaches c
	MOVL BX, DX
	SUBL CX, DX
	SUBL AX, DX // c - e
	CMPL R15, DX
	JB   done // something other than whitespace before the comma

fold:
	CMPQ R10, R9
	JAE  done // out is full: leave this element to the caller
	// control = intRunShuffle[s*8+L] (unique for L in 1..8), eight bytes
	// that right-align digits s..s+L-1 into the low lane with zeros in
	// front; the upper lane's control is whatever VMOVQ zeroes to, and the
	// garbage it produces stays in the words the final fold never reads.
	LEAL (AX)(CX*8), R15
	VMOVQ (R14)(R15*8), X2
	VPSHUFB X2, X1, X3
	VPMADDUBSW X8, X3, X3
	VPMADDWD X9, X3, X3
	VPACKUSDW X3, X3, X3
	VPMADDWD X10, X3, X3
	VMOVD X3, R15
	MOVQ R15, (R8)(R10*8)
	INCQ R10
	BLSRL R12, R12 // the comma is consumed
	LEAL 1(BX), CX // b = just past it
	CMPL CX, $16
	JB   elem
	ADDQ $16, DI // the comma was the block's last byte
	JMP  block

noComma:
	// No comma left in the block: the element can only be the array's
	// last, terminated by ']' (after optional whitespace) in this block.
	VPCMPEQB X11, X0, X2
	VPMOVMSKB X2, R15 // close brackets
	MOVL R13, BX
	XORL $0xffff, BX
	SHRXL DX, BX, BX
	TESTL BX, BX
	JZ   straddle
	TZCNTL BX, BX
	ADDL DX, BX // the terminator's offset
	BTL  BX, R15
	JNC  done // not ']': not this kernel's element
	MOVL BX, DX

foldClose:
	// The array's last element, terminated by ']' at DX: fold it, then
	// return with p at the ']' and closed set.
	CMPQ R10, R9
	JAE  done
	LEAL (AX)(CX*8), R15
	VMOVQ (R14)(R15*8), X2
	VPSHUFB X2, X1, X3
	VPMADDUBSW X8, X3, X3
	VPMADDWD X9, X3, X3
	VPACKUSDW X3, X3, X3
	VPMADDWD X10, X3, X3
	VMOVD X3, R15
	MOVQ R15, (R8)(R10*8)
	INCQ R10
	ADDQ DX, DI
	MOVQ R10, n+56(FP)
	MOVQ DI, p+64(FP)
	MOVQ $1, closed+72(FP)
	RET

straddle:
	// The element (or the whitespace before it) runs to the block's end:
	// start the next block at it. At b == 0 that would not move, so stop.
	TESTL CX, CX
	JZ   done
	ADDQ CX, DI
	JMP  block

done:
	ADDQ CX, DI
	MOVQ R10, n+56(FP)
	MOVQ DI, p+64(FP)
	MOVQ $0, closed+72(FP)
	RET
