#include "textflag.h"

// 32-byte splats of the six bytes maskBlock classifies.
DATA mbQuote<>+0(SB)/8, $0x2222222222222222
DATA mbQuote<>+8(SB)/8, $0x2222222222222222
DATA mbQuote<>+16(SB)/8, $0x2222222222222222
DATA mbQuote<>+24(SB)/8, $0x2222222222222222
GLOBL mbQuote<>(SB), RODATA|NOPTR, $32

DATA mbBslash<>+0(SB)/8, $0x5c5c5c5c5c5c5c5c
DATA mbBslash<>+8(SB)/8, $0x5c5c5c5c5c5c5c5c
DATA mbBslash<>+16(SB)/8, $0x5c5c5c5c5c5c5c5c
DATA mbBslash<>+24(SB)/8, $0x5c5c5c5c5c5c5c5c
GLOBL mbBslash<>(SB), RODATA|NOPTR, $32

DATA mbLbrace<>+0(SB)/8, $0x7b7b7b7b7b7b7b7b
DATA mbLbrace<>+8(SB)/8, $0x7b7b7b7b7b7b7b7b
DATA mbLbrace<>+16(SB)/8, $0x7b7b7b7b7b7b7b7b
DATA mbLbrace<>+24(SB)/8, $0x7b7b7b7b7b7b7b7b
GLOBL mbLbrace<>(SB), RODATA|NOPTR, $32

DATA mbRbrace<>+0(SB)/8, $0x7d7d7d7d7d7d7d7d
DATA mbRbrace<>+8(SB)/8, $0x7d7d7d7d7d7d7d7d
DATA mbRbrace<>+16(SB)/8, $0x7d7d7d7d7d7d7d7d
DATA mbRbrace<>+24(SB)/8, $0x7d7d7d7d7d7d7d7d
GLOBL mbRbrace<>(SB), RODATA|NOPTR, $32

DATA mbLbrack<>+0(SB)/8, $0x5b5b5b5b5b5b5b5b
DATA mbLbrack<>+8(SB)/8, $0x5b5b5b5b5b5b5b5b
DATA mbLbrack<>+16(SB)/8, $0x5b5b5b5b5b5b5b5b
DATA mbLbrack<>+24(SB)/8, $0x5b5b5b5b5b5b5b5b
GLOBL mbLbrack<>(SB), RODATA|NOPTR, $32

DATA mbRbrack<>+0(SB)/8, $0x5d5d5d5d5d5d5d5d
DATA mbRbrack<>+8(SB)/8, $0x5d5d5d5d5d5d5d5d
DATA mbRbrack<>+16(SB)/8, $0x5d5d5d5d5d5d5d5d
DATA mbRbrack<>+24(SB)/8, $0x5d5d5d5d5d5d5d5d
GLOBL mbRbrack<>(SB), RODATA|NOPTR, $32

// func maskBlock(b []byte) (quote, bslash, lbrace, rbrace, lbrack, rbrack uint32)
//
// Returns, for the 32 bytes at b[:32], the bitmap of each of '"' '\\' '{' '}'
// '[' ']' (bit j set when b[j] is that byte). AVX2; the caller guarantees 32
// readable bytes.
TEXT ·maskBlock(SB), NOSPLIT, $0-48
	MOVQ    b_base+0(FP), SI
	VMOVDQU (SI), Y0
	VMOVDQU mbQuote<>(SB), Y2
	VMOVDQU mbBslash<>(SB), Y3
	VMOVDQU mbLbrace<>(SB), Y4
	VMOVDQU mbRbrace<>(SB), Y5
	VMOVDQU mbLbrack<>(SB), Y6
	VMOVDQU mbRbrack<>(SB), Y7

	VPCMPEQB  Y2, Y0, Y8
	VPMOVMSKB Y8, AX
	MOVL      AX, quote+24(FP)
	VPCMPEQB  Y3, Y0, Y8
	VPMOVMSKB Y8, AX
	MOVL      AX, bslash+28(FP)
	VPCMPEQB  Y4, Y0, Y8
	VPMOVMSKB Y8, AX
	MOVL      AX, lbrace+32(FP)
	VPCMPEQB  Y5, Y0, Y8
	VPMOVMSKB Y8, AX
	MOVL      AX, rbrace+36(FP)
	VPCMPEQB  Y6, Y0, Y8
	VPMOVMSKB Y8, AX
	MOVL      AX, lbrack+40(FP)
	VPCMPEQB  Y7, Y0, Y8
	VPMOVMSKB Y8, AX
	MOVL      AX, rbrack+44(FP)
	VZEROUPPER
	RET
