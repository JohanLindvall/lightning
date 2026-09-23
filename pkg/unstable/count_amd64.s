#include "textflag.h"

DATA cbcClose<>+0(SB)/8, $0x5d5d5d5d5d5d5d5d // ']' splat
DATA cbcClose<>+8(SB)/8, $0x5d5d5d5d5d5d5d5d
DATA cbcClose<>+16(SB)/8, $0x5d5d5d5d5d5d5d5d
DATA cbcClose<>+24(SB)/8, $0x5d5d5d5d5d5d5d5d
GLOBL cbcClose<>(SB), RODATA|NOPTR, $32

// func countKernel(data []byte, i int, c byte, hint bool) (rb, n int)
//
// Scans data[i:] for the first ']' and returns its offset from i — what
// bytes.IndexByte(data[i:], ']') returns — and the number of c bytes before it;
// rb is -1 and n is 0 when there is no ']'. One pass and one call where the
// presize counters it serves made two library calls over the same bytes
// (IndexByte, then Count over the span it found), each with its own dispatch
// and setup; for the 3- and 4-element coordinate arrays that dominate
// marine_ik that setup was most of the count.
//
// With hint set n is instead the element count
// CountArrayScalars wants: the commas plus one clamped to (rb+1)/2, and for a
// span with no comma 1 if it holds a byte above 0x20 — this library's whitespace rule, the one
// SkipWS applies between tokens — and 0 if it is blank. That decision was a Go
// loop, and it was what kept CountArrayScalars from inlining into the batch
// readers; here it is a walk that stops at the first non-blank byte, which in
// a one-element array is the element's first byte.
//
// AVX2 (·useCountAVX2) takes 32 bytes a step: one compare for ']' and one for
// c, and the step that holds the ']' counts only the c lanes below it (BZHI).
// The last < 32 bytes are the buffer's LAST 32 with the lanes before the
// cursor shifted out of both masks, so a buffer of 32 bytes or more has no
// byte loop at all; a shorter one, or a host without AVX2, is walked a byte at
// a time.
//
// Registers: DX start pointer, R8 cursor, R9 end, R10 count, R11 c, R12 hint.
TEXT ·countKernel(SB), NOSPLIT, $0-56
	MOVQ    data_base+0(FP), SI
	MOVQ    data_len+8(FP), CX
	MOVQ    i+24(FP), DI
	MOVBLZX c+32(FP), R11
	MOVBLZX hint+33(FP), R12
	XORQ    R10, R10
	CMPQ    DI, CX
	JGE     cbc_notfound
	LEAQ    (SI)(CX*1), R9
	LEAQ    (SI)(DI*1), R8
	MOVQ    R8, DX
	MOVBLZX ·useCountAVX2(SB), AX
	TESTL   AX, AX
	JZ      cbc_bytes
	CMPQ    CX, $32
	JLT     cbc_bytes
	VMOVDQU cbcClose<>(SB), Y0
	MOVQ    R11, X1
	VPBROADCASTB X1, Y1
	LEAQ    -32(R9), BX             // the last start of a full block
	CMPQ    R8, BX
	JHI     cbc_tail
cbc_loop:
	VMOVDQU  (R8), Y2
	VPCMPEQB Y0, Y2, Y3
	VPCMPEQB Y1, Y2, Y4
	VPMOVMSKB Y3, AX
	VPMOVMSKB Y4, CX
	TESTL    AX, AX
	JNZ      cbc_found
	POPCNTL  CX, CX
	ADDQ     CX, R10
	ADDQ     $32, R8
	CMPQ     R8, BX
	JLS      cbc_loop
cbc_tail:
	CMPQ     R8, R9
	JEQ      cbc_vznotfound
	VMOVDQU  -32(R9), Y2
	VPCMPEQB Y0, Y2, Y3
	VPCMPEQB Y1, Y2, Y4
	VPMOVMSKB Y3, AX
	VPMOVMSKB Y4, CX
	MOVQ     R8, DI
	SUBQ     BX, DI                 // lanes already counted: R8 - (end-32)
	SHRXL    DI, AX, AX
	SHRXL    DI, CX, CX
	TESTL    AX, AX
	JZ       cbc_vznotfound
cbc_found:
	VZEROUPPER
	TZCNTL   AX, AX                 // the ']' lane
	BZHIL    AX, CX, CX             // the c lanes below it
	POPCNTL  CX, CX
	ADDQ     CX, R10
	SUBQ     DX, R8
	ADDQ     R8, AX
cbc_result:
	// AX = rb, R10 = count.
	TESTL    R12, R12
	JNZ      cbc_hint
	MOVQ     AX, rb+40(FP)
	MOVQ     R10, n+48(FP)
	RET
cbc_hint:
	MOVQ     AX, rb+40(FP)
	TESTQ    R10, R10
	JZ       cbc_blank
	INCQ     R10                    // elements = commas + 1
	// Clamp to the element count the span can hold: n elements need n-1
	// commas and a byte each, so rb >= 2n-1 and n <= (rb+1)/2 (see
	// CountArrayScalars for why this bound matters and never clips an
	// honest count).
	LEAQ     1(AX), BX
	SHRQ     $1, BX
	CMPQ     R10, BX
	CMOVQHI  BX, R10
	MOVQ     R10, n+48(FP)
	RET
cbc_blank:
	// No comma before the ']' at DX+AX: one element unless every byte is
	// whitespace. The walk stops at the first byte above 0x20, which in any
	// one-element array is the element's first byte.
	LEAQ     (DX)(AX*1), BX
cbc_blankloop:
	CMPQ     DX, BX
	JEQ      cbc_empty
	MOVBLZX  (DX), CX
	INCQ     DX
	CMPL     CX, $0x20
	JLS      cbc_blankloop
	MOVQ     $1, n+48(FP)
	RET
cbc_empty:
	MOVQ     $0, n+48(FP)
	RET

cbc_bytes:
	MOVBLZX (R8), AX
	CMPL    AX, $0x5d
	JEQ     cbc_bytefound
	CMPL    AX, R11
	JNE     cbc_bytenext
	INCQ    R10
cbc_bytenext:
	INCQ    R8
	CMPQ    R8, R9
	JNE     cbc_bytes
	JMP     cbc_notfound
cbc_bytefound:
	SUBQ    DX, R8
	MOVQ    R8, AX
	JMP     cbc_result

cbc_vznotfound:
	VZEROUPPER
cbc_notfound:
	MOVQ    $-1, rb+40(FP)
	MOVQ    $0, n+48(FP)
	RET
