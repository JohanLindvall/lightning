#include "textflag.h"

// 32-byte constant vectors of '"' (0x22) and '\\' (0x5c) for VPCMPEQB.
DATA quoteMask<>+0(SB)/8, $0x2222222222222222
DATA quoteMask<>+8(SB)/8, $0x2222222222222222
DATA quoteMask<>+16(SB)/8, $0x2222222222222222
DATA quoteMask<>+24(SB)/8, $0x2222222222222222
GLOBL quoteMask<>(SB), RODATA|NOPTR, $32

DATA bslashMask<>+0(SB)/8, $0x5c5c5c5c5c5c5c5c
DATA bslashMask<>+8(SB)/8, $0x5c5c5c5c5c5c5c5c
DATA bslashMask<>+16(SB)/8, $0x5c5c5c5c5c5c5c5c
DATA bslashMask<>+24(SB)/8, $0x5c5c5c5c5c5c5c5c
GLOBL bslashMask<>(SB), RODATA|NOPTR, $32

// 32-byte constant vectors of the JSON structural bytes '{' (0x7b), '}' (0x7d),
// '[' (0x5b), ']' (0x5d) and '"' (0x22). Precomputed in RODATA and loaded with
// VMOVDQU so indexStructuralAVX2 does not rebuild them with VPBROADCASTB (and
// the GP→XMM domain crossings that entails) on every call — that per-call setup
// dominated when the routine early-exits after a single block, which is the
// common case for token-dense JSON.
DATA lbraceMask<>+0(SB)/8, $0x7b7b7b7b7b7b7b7b
DATA lbraceMask<>+8(SB)/8, $0x7b7b7b7b7b7b7b7b
DATA lbraceMask<>+16(SB)/8, $0x7b7b7b7b7b7b7b7b
DATA lbraceMask<>+24(SB)/8, $0x7b7b7b7b7b7b7b7b
GLOBL lbraceMask<>(SB), RODATA|NOPTR, $32

DATA rbraceMask<>+0(SB)/8, $0x7d7d7d7d7d7d7d7d
DATA rbraceMask<>+8(SB)/8, $0x7d7d7d7d7d7d7d7d
DATA rbraceMask<>+16(SB)/8, $0x7d7d7d7d7d7d7d7d
DATA rbraceMask<>+24(SB)/8, $0x7d7d7d7d7d7d7d7d
GLOBL rbraceMask<>(SB), RODATA|NOPTR, $32

DATA lbrackMask<>+0(SB)/8, $0x5b5b5b5b5b5b5b5b
DATA lbrackMask<>+8(SB)/8, $0x5b5b5b5b5b5b5b5b
DATA lbrackMask<>+16(SB)/8, $0x5b5b5b5b5b5b5b5b
DATA lbrackMask<>+24(SB)/8, $0x5b5b5b5b5b5b5b5b
GLOBL lbrackMask<>(SB), RODATA|NOPTR, $32

DATA rbrackMask<>+0(SB)/8, $0x5d5d5d5d5d5d5d5d
DATA rbrackMask<>+8(SB)/8, $0x5d5d5d5d5d5d5d5d
DATA rbrackMask<>+16(SB)/8, $0x5d5d5d5d5d5d5d5d
DATA rbrackMask<>+24(SB)/8, $0x5d5d5d5d5d5d5d5d
GLOBL rbrackMask<>(SB), RODATA|NOPTR, $32

// func indexQuoteOrBackslashAVX2(b []byte) int
//
// Returns the index of the first '"' or '\\' byte in b, or len(b) if neither
// is present. Scans 32 bytes per iteration with AVX2, then a scalar tail.
TEXT ·indexQuoteOrBackslashAVX2(SB), NOSPLIT, $0-32
	MOVQ b_base+0(FP), SI
	MOVQ b_len+8(FP), CX
	XORQ DI, DI                  // DI = current offset
	VMOVDQU quoteMask<>(SB), Y0
	VMOVDQU bslashMask<>(SB), Y1

loop:
	CMPQ CX, $32
	JL   tail
	VMOVDQU (SI)(DI*1), Y2
	VPCMPEQB Y0, Y2, Y3          // Y3 = (chunk == '"')
	VPCMPEQB Y1, Y2, Y4          // Y4 = (chunk == '\\')
	VPOR     Y4, Y3, Y3          // Y3 = either
	VPMOVMSKB Y3, AX             // 32-bit mask of matching lanes
	TESTL    AX, AX
	JNZ      found
	ADDQ     $32, DI
	SUBQ     $32, CX
	JMP      loop

found:
	BSFL AX, AX                  // first set bit = offset within chunk
	ADDQ DI, AX
	MOVQ AX, ret+24(FP)
	VZEROUPPER
	RET

tail:
	TESTQ CX, CX
	JZ    notfound

tailloop:
	MOVBLZX (SI)(DI*1), AX
	CMPL    AX, $0x22
	JE      tfound
	CMPL    AX, $0x5c
	JE      tfound
	INCQ    DI
	DECQ    CX
	JNZ     tailloop

notfound:
	MOVQ b_len+8(FP), AX
	MOVQ AX, ret+24(FP)
	VZEROUPPER
	RET

tfound:
	MOVQ DI, ret+24(FP)
	VZEROUPPER
	RET

// func indexStructuralAVX2(b []byte) int
//
// Returns the index of the first '{', '}', '[', ']' or '"' byte in b, or
// len(b) if none is present.
TEXT ·indexStructuralAVX2(SB), NOSPLIT, $0-32
	MOVQ b_base+0(FP), SI
	MOVQ b_len+8(FP), CX
	XORQ DI, DI
	VMOVDQU lbraceMask<>(SB), Y0 // '{'
	VMOVDQU rbraceMask<>(SB), Y1 // '}'
	VMOVDQU lbrackMask<>(SB), Y2 // '['
	VMOVDQU rbrackMask<>(SB), Y3 // ']'
	VMOVDQU quoteMask<>(SB), Y4  // '"'

loops:
	CMPQ CX, $32
	JL   tails
	VMOVDQU (SI)(DI*1), Y5
	VPCMPEQB Y0, Y5, Y6
	VPCMPEQB Y1, Y5, Y7
	VPOR     Y7, Y6, Y6
	VPCMPEQB Y2, Y5, Y7
	VPOR     Y7, Y6, Y6
	VPCMPEQB Y3, Y5, Y7
	VPOR     Y7, Y6, Y6
	VPCMPEQB Y4, Y5, Y7
	VPOR     Y7, Y6, Y6
	VPMOVMSKB Y6, AX
	TESTL    AX, AX
	JNZ      founds
	ADDQ     $32, DI
	SUBQ     $32, CX
	JMP      loops

founds:
	BSFL AX, AX
	ADDQ DI, AX
	MOVQ AX, ret+24(FP)
	VZEROUPPER
	RET

tails:
	TESTQ CX, CX
	JZ    notfounds

tailsloop:
	MOVBLZX (SI)(DI*1), AX
	CMPL    AX, $0x7b
	JE      tfs
	CMPL    AX, $0x7d
	JE      tfs
	CMPL    AX, $0x5b
	JE      tfs
	CMPL    AX, $0x5d
	JE      tfs
	CMPL    AX, $0x22
	JE      tfs
	INCQ    DI
	DECQ    CX
	JNZ     tailsloop

notfounds:
	MOVQ b_len+8(FP), AX
	MOVQ AX, ret+24(FP)
	VZEROUPPER
	RET

tfs:
	MOVQ DI, ret+24(FP)
	VZEROUPPER
	RET
