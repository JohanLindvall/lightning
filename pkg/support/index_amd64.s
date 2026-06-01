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
