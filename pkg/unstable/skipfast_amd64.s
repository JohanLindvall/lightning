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

// func maskBlock(b []byte, isArray bool) (quote, bslash, open, close uint64)
//
// Returns, for the 64 bytes at b[:64], the bitmap of `"`, `\\`, and the
// container's open/close brackets — `[`/`]` when isArray, else `{`/`}`; bit j set
// when b[j] is that byte. AVX2; the caller guarantees 64 readable bytes. Each
// uint64 is the low 32-byte half's 32-bit movemask in bits 0-31 and the high
// half's in bits 32-63.
//
// CLASS computes one class: compare both 32-byte halves against splat \s, combine
// the two movemasks into a uint64, and store at frame offset \off.
#define CLASS(s, off) \
	VPCMPEQB  s, Y0, Y2          \
	VPMOVMSKB Y2, AX             \
	VPCMPEQB  s, Y1, Y3          \
	VPMOVMSKB Y3, BX             \
	SHLQ      $32, BX            \
	ORQ       BX, AX             \
	MOVQ      AX, off(FP)

TEXT ·maskBlock(SB), NOSPLIT, $0-64
	MOVQ    b_base+0(FP), SI
	VMOVDQU (SI), Y0            // bytes 0-31
	VMOVDQU 32(SI), Y1          // bytes 32-63
	VMOVDQU mbQuote<>(SB), Y10
	VMOVDQU mbBslash<>(SB), Y11
	// Select the container's bracket splats: '['/']' for an array, else '{'/'}'.
	MOVBLZX isArray+24(FP), AX
	TESTL   AX, AX
	JNZ     arrayBrackets
	VMOVDQU mbLbrace<>(SB), Y12
	VMOVDQU mbRbrace<>(SB), Y13
	JMP     haveBrackets
arrayBrackets:
	VMOVDQU mbLbrack<>(SB), Y12
	VMOVDQU mbRbrack<>(SB), Y13
haveBrackets:
	CLASS(Y10, quote+32)
	CLASS(Y11, bslash+40)
	CLASS(Y12, open+48)
	CLASS(Y13, close+56)

	VZEROUPPER
	RET
