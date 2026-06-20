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

// Shuffle-classification tables for indexStructuralAVX2 (the simdjson
// find_structurals trick). A byte is one of '{' '}' '[' ']' '"' iff
// structLoTable[lowNibble] & structHiTable[highNibble] != 0. The two bits encode
// the two nibble groups — '"' (lo 0x2, hi 0x2) and the brackets/braces (lo 0xB/0xD,
// hi 0x5/0x7) — so cross combinations like 0x52 'R' classify as non-structural.
// Each 16-byte table is duplicated across both AVX2 lanes (VPSHUFB is per-lane).
DATA structLoTable<>+0(SB)/8, $0x0000000000010000
DATA structLoTable<>+8(SB)/8, $0x0000020002000000
DATA structLoTable<>+16(SB)/8, $0x0000000000010000
DATA structLoTable<>+24(SB)/8, $0x0000020002000000
GLOBL structLoTable<>(SB), RODATA|NOPTR, $32

DATA structHiTable<>+0(SB)/8, $0x0200020000010000
DATA structHiTable<>+8(SB)/8, $0x0000000000000000
DATA structHiTable<>+16(SB)/8, $0x0200020000010000
DATA structHiTable<>+24(SB)/8, $0x0000000000000000
GLOBL structHiTable<>(SB), RODATA|NOPTR, $32

DATA nibbleMask<>+0(SB)/8, $0x0f0f0f0f0f0f0f0f
DATA nibbleMask<>+8(SB)/8, $0x0f0f0f0f0f0f0f0f
DATA nibbleMask<>+16(SB)/8, $0x0f0f0f0f0f0f0f0f
DATA nibbleMask<>+24(SB)/8, $0x0f0f0f0f0f0f0f0f
GLOBL nibbleMask<>(SB), RODATA|NOPTR, $32

// func indexQuoteOrBackslashSSE2(b []byte) int
//
// Returns the index of the first '"' or '\\' byte in b, or len(b) if neither
// is present, then a scalar tail. It uses no 256-bit registers, so it needs no
// VZEROUPPER — the AVX2 variant pays that on every call, which dominates for the
// short keys and string values that make up most JSON. The dispatch routes
// string scanning here. The main loop covers 32 bytes per iteration with two
// 16-byte SSE2 compares so it matches AVX2's stride (a 17-32 byte string still
// finishes in one iteration) without the VZEROUPPER, and the second 16-byte load
// is skipped entirely when the match lands in the first half. A 16-byte loop
// then handles a final 16-31 byte span before the scalar tail.
TEXT ·indexQuoteOrBackslashSSE2(SB), NOSPLIT, $0-32
	MOVQ b_base+0(FP), SI
	MOVQ b_len+8(FP), CX
	XORQ DI, DI
	MOVOU quoteMask<>(SB), X0    // low 16 bytes of the splat: '"' x16
	MOVOU bslashMask<>(SB), X1   // '\\' x16

sse_loop32:
	CMPQ CX, $32
	JL   sse_loop16
	MOVOU (SI)(DI*1), X2         // first 16 bytes
	MOVOU X2, X3
	PCMPEQB X0, X2
	PCMPEQB X1, X3
	POR     X3, X2
	PMOVMSKB X2, AX
	TESTL    AX, AX
	JNZ      sse_found
	MOVOU 16(SI)(DI*1), X4       // second 16 bytes (only if first had no match)
	MOVOU X4, X5
	PCMPEQB X0, X4
	PCMPEQB X1, X5
	POR     X5, X4
	PMOVMSKB X4, AX
	TESTL    AX, AX
	JNZ      sse_found16
	ADDQ     $32, DI
	SUBQ     $32, CX
	// The first 32 bytes held no '"' or '\\': a long string (the short keys and
	// values that dominate JSON already returned above, never reaching here). Such
	// strings amortize AVX2's one VZEROUPPER over many bytes, and its single 32-byte
	// compare per iteration halves the two 16-byte SSE compares the loop above does.
	MOVBLZX ·useAVX2(SB), AX
	TESTL   AX, AX
	JNZ     qb_avx_setup
	JMP     sse_loop32

qb_avx_setup:
	VMOVDQU quoteMask<>(SB), Y0
	VMOVDQU bslashMask<>(SB), Y1

qb_avx_loop:
	CMPQ CX, $32
	JL   qb_avx_done
	VMOVDQU (SI)(DI*1), Y2
	VPCMPEQB Y0, Y2, Y3
	VPCMPEQB Y1, Y2, Y4
	VPOR      Y4, Y3, Y3
	VPMOVMSKB Y3, AX
	TESTL     AX, AX
	JNZ       qb_avx_found
	ADDQ      $32, DI
	SUBQ      $32, CX
	JMP       qb_avx_loop

qb_avx_found:
	BSFL AX, AX
	ADDQ DI, AX
	MOVQ AX, ret+24(FP)
	VZEROUPPER
	RET

qb_avx_done:
	VZEROUPPER
	// <32 bytes remain; the X-register splats survive VZEROUPPER, so finish in SSE.

sse_loop16:
	CMPQ CX, $16
	JL   sse_tail
	MOVOU (SI)(DI*1), X2
	MOVOU X2, X3
	PCMPEQB X0, X2
	PCMPEQB X1, X3
	POR     X3, X2
	PMOVMSKB X2, AX
	TESTL    AX, AX
	JNZ      sse_found
	ADDQ     $16, DI
	SUBQ     $16, CX
	JMP      sse_loop16

sse_found16:
	ADDQ $16, DI                 // match was in the second 16-byte half

sse_found:
	BSFL AX, AX
	ADDQ DI, AX
	MOVQ AX, ret+24(FP)
	RET

sse_tail:
	TESTQ CX, CX
	JZ    sse_notfound

sse_tailloop:
	MOVBLZX (SI)(DI*1), AX
	CMPL    AX, $0x22
	JE      sse_tfound
	CMPL    AX, $0x5c
	JE      sse_tfound
	INCQ    DI
	DECQ    CX
	JNZ     sse_tailloop

sse_notfound:
	MOVQ b_len+8(FP), AX
	MOVQ AX, ret+24(FP)
	RET

sse_tfound:
	MOVQ DI, ret+24(FP)
	RET

// func indexStructuralAVX2(b []byte) int
//
// Returns the index of the first '{', '}', '[', ']' or '"' byte in b, or
// len(b) if none is present.
TEXT ·indexStructuralAVX2(SB), NOSPLIT, $0-32
	MOVQ b_base+0(FP), SI
	MOVQ b_len+8(FP), CX
	XORQ DI, DI
	VMOVDQU structLoTable<>(SB), Y0
	VMOVDQU structHiTable<>(SB), Y1
	VMOVDQU nibbleMask<>(SB), Y2
	VPXOR   Y3, Y3, Y3

loops:
	CMPQ CX, $32
	JL   tails
	VMOVDQU  (SI)(DI*1), Y5
	VPAND    Y2, Y5, Y6  // low nibbles
	VPSHUFB  Y6, Y0, Y6  // structLoTable[lowNibble]
	VPSRLW   $4, Y5, Y7
	VPAND    Y2, Y7, Y7  // high nibbles
	VPSHUFB  Y7, Y1, Y7  // structHiTable[highNibble]
	VPAND    Y7, Y6, Y6  // nonzero where structural
	VPCMPEQB Y3, Y6, Y6  // 0xFF where NOT structural
	VPMOVMSKB Y6, AX     // 1-bits mark non-structural bytes
	NOTL     AX          // ... so invert to mark structural bytes
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
