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

// 32-byte splat of 0x1f, the largest control byte. indexEscapeSSE2 flags a control
// byte (< 0x20) per lane with PMINUB(v, 0x1f) == v: min(c, 0x1f) equals c exactly
// when c <= 0x1f.
DATA ctrlMask<>+0(SB)/8, $0x1f1f1f1f1f1f1f1f
DATA ctrlMask<>+8(SB)/8, $0x1f1f1f1f1f1f1f1f
DATA ctrlMask<>+16(SB)/8, $0x1f1f1f1f1f1f1f1f
DATA ctrlMask<>+24(SB)/8, $0x1f1f1f1f1f1f1f1f
GLOBL ctrlMask<>(SB), RODATA|NOPTR, $32

// Splats for indexStructuralAVX2's AVX2 body. A byte is structural iff
// (c|0x20) == '{' (which holds for exactly '{' and '['), (c|0x20) == '}'
// (exactly '}' and ']'), or c == '"': the two bracket pairs differ only in bit
// 0x20, so ORing it in folds each pair onto one compare, and the quote is
// compared unfolded because 0x22|0x20 is also 0x02|0x20. Three compares and one
// OR per vector, with no shuffle — the nibble-table classification this
// replaced spent two VPSHUFB, a shift and three ANDs, and needed the mask
// inverted afterwards.
DATA stBit5<>+0(SB)/8, $0x2020202020202020
DATA stBit5<>+8(SB)/8, $0x2020202020202020
DATA stBit5<>+16(SB)/8, $0x2020202020202020
DATA stBit5<>+24(SB)/8, $0x2020202020202020
GLOBL stBit5<>(SB), RODATA|NOPTR, $32

DATA stOpen<>+0(SB)/8, $0x7b7b7b7b7b7b7b7b
DATA stOpen<>+8(SB)/8, $0x7b7b7b7b7b7b7b7b
DATA stOpen<>+16(SB)/8, $0x7b7b7b7b7b7b7b7b
DATA stOpen<>+24(SB)/8, $0x7b7b7b7b7b7b7b7b
GLOBL stOpen<>(SB), RODATA|NOPTR, $32

DATA stClose<>+0(SB)/8, $0x7d7d7d7d7d7d7d7d
DATA stClose<>+8(SB)/8, $0x7d7d7d7d7d7d7d7d
DATA stClose<>+16(SB)/8, $0x7d7d7d7d7d7d7d7d
DATA stClose<>+24(SB)/8, $0x7d7d7d7d7d7d7d7d
GLOBL stClose<>(SB), RODATA|NOPTR, $32

// The AVX-512 VBMI body's classification table, one VPERMB and one compare per
// 64 bytes. VPERMB indexes a 64-byte table with the low six bits of each input
// byte, and the five structural bytes have distinct low six bits (0x22, 0x1b,
// 0x1d, 0x3b, 0x3d), so entry k holds the structural byte whose low six bits are
// k and every other entry holds k^1 — a byte whose low six bits are NOT k, which
// no input c with c&63 == k can equal. So table[c&63] == c holds for exactly the
// five structural bytes. Entry 0 is 0x01, so a zeroed lane of a masked load
// never matches either.
DATA structPerm<>+0(SB)/8, $0x0607040502030001
DATA structPerm<>+8(SB)/8, $0x0e0f0c0d0a0b0809
DATA structPerm<>+16(SB)/8, $0x1617141512131011
DATA structPerm<>+24(SB)/8, $0x1e1f5d1d5b1b1819
DATA structPerm<>+32(SB)/8, $0x2627242522222021
DATA structPerm<>+40(SB)/8, $0x2e2f2c2d2a2b2829
DATA structPerm<>+48(SB)/8, $0x3637343532333031
DATA structPerm<>+56(SB)/8, $0x3e3f7d3d7b3b3839
GLOBL structPerm<>(SB), RODATA|NOPTR, $64

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
//
// The two splats are not loaded into registers up front: every compare takes
// its splat as a memory operand instead. Legacy SSE requires a 16-byte aligned
// memory operand, which these 32-byte RODATA symbols have (the linker aligns a
// symbol to the power of two covering its size, up to 32). Almost every call
// is one block, so the two up-front loads were two instructions of eight in
// front of the first compare, on the call that runs once per object key and
// once per string value; folded, the loads issue with the compares they feed
// and the call is two instructions shorter.
TEXT ·indexQuoteOrBackslashSSE2(SB), NOSPLIT, $0-40
	MOVQ b_base+0(FP), SI
	MOVQ b_len+8(FP), CX
	MOVQ i+24(FP), DI            // scan offset; DI stays the absolute index
	SUBQ DI, CX                  // CX = bytes left from i
	JS   sse_notfound            // i past the end: no bytes, answer len(b)

sse_loop32:
	CMPQ CX, $32
	JL   sse_loop16
	MOVOU (SI)(DI*1), X2         // first 16 bytes
	MOVOU X2, X3
	PCMPEQB quoteMask<>(SB), X2  // the splats are memory operands: see above
	PCMPEQB bslashMask<>(SB), X3
	POR     X3, X2
	PMOVMSKB X2, AX
	TESTL    AX, AX
	JNZ      sse_found
	MOVOU 16(SI)(DI*1), X4       // second 16 bytes (only if first had no match)
	MOVOU X4, X5
	PCMPEQB quoteMask<>(SB), X4
	PCMPEQB bslashMask<>(SB), X5
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
	MOVQ AX, ret+32(FP)
	VZEROUPPER
	RET

qb_avx_done:
	VZEROUPPER
	// <32 bytes remain; finish in SSE.

sse_loop16:
	CMPQ CX, $16
	JL   sse_tail
	MOVOU (SI)(DI*1), X2
	MOVOU X2, X3
	PCMPEQB quoteMask<>(SB), X2
	PCMPEQB bslashMask<>(SB), X3
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
	MOVQ AX, ret+32(FP)
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
	MOVQ AX, ret+32(FP)
	RET

sse_tfound:
	MOVQ DI, ret+32(FP)
	RET

// func indexStructuralAVX2(b []byte, i int) int
//
// Returns the index of the first '{', '}', '[', ']' or '"' byte in b at or
// after i, or len(b) if none is present. Two bodies, selected here rather than
// in Go for the reason the string scanner gives: ·useStructural512 (AVX-512
// VBMI) takes one VPERMB and one VPCMPEQB into a mask register per 64 bytes and
// ends in a masked load, so it has no scalar tail at all; otherwise AVX2 scans
// 64 bytes per iteration with the compare classification above and finishes on
// an OVERLAPPING 32-byte block (the buffer's last 32 bytes, the lanes before
// the cursor shifted out of the mask) instead of a byte loop.
//
// Both loops advance a pointer rather than an index so that every load has a
// base-only address (an indexed VEX operand splits into two uops on Intel),
// and both test the loop condition at the bottom, which leaves one taken branch
// per iteration. The skip-heavy case is this loop and nothing else, and on Zen 4
// the loop it replaced ran at 5.8 ops a cycle — the dispatch limit — so the
// currency was ops per byte: 15 per 32 bytes before, 17 per 64 on AVX2 and 5 per
// 64 on VBMI.
//
// The VBMI body uses only Z16-Z31, which no SSE or VEX instruction can reach,
// so it needs no VZEROUPPER (the glibc EVEX string functions' reason for the
// same choice).
TEXT ·indexStructuralAVX2(SB), NOSPLIT, $0-40
	MOVQ b_base+0(FP), DX           // DX = base
	MOVQ b_len+8(FP), CX
	MOVQ i+24(FP), SI
	LEAQ (DX)(CX*1), R9             // R9 = end
	CMPQ SI, CX
	JGE  st_notfound
	LEAQ (DX)(SI*1), SI             // SI = p
	MOVBLZX ·useStructural512(SB), AX
	TESTL   AX, AX
	JNZ     st_vbmi

	// AVX2: 64 bytes a step while a full step remains. A buffer under 32 bytes
	// holds no block at all, so it is walked before any vector state is touched.
	CMPQ CX, $32
	JLT  st_bytes
	VMOVDQU stBit5<>(SB), Y10
	VMOVDQU stOpen<>(SB), Y11
	VMOVDQU stClose<>(SB), Y12
	VMOVDQU quoteMask<>(SB), Y13
	LEAQ -64(R9), R8                // last start of a full 64-byte step
	CMPQ SI, R8
	JHI  st_avx_32
st_avx_loop:
	VMOVDQU  (SI), Y0
	VMOVDQU  32(SI), Y3
	VPOR     Y10, Y0, Y1            // t = c | 0x20
	VPCMPEQB Y13, Y0, Y0            // c == '"'
	VPCMPEQB Y11, Y1, Y2            // t == '{'  ('{' or '[')
	VPCMPEQB Y12, Y1, Y1            // t == '}'  ('}' or ']')
	VPOR     Y2, Y0, Y0
	VPOR     Y1, Y0, Y0
	VPOR     Y10, Y3, Y4
	VPCMPEQB Y13, Y3, Y3
	VPCMPEQB Y11, Y4, Y5
	VPCMPEQB Y12, Y4, Y4
	VPOR     Y5, Y3, Y3
	VPOR     Y4, Y3, Y3
	VPOR     Y3, Y0, Y5
	VPMOVMSKB Y5, AX
	TESTL    AX, AX
	JNZ      st_avx_found64
	ADDQ     $64, SI
	CMPQ     SI, R8
	JLS      st_avx_loop

st_avx_32:
	// < 64 bytes left: one 32-byte block if it fits, then the overlapping tail.
	LEAQ -32(R9), R8                // last start of a full 32-byte block
	CMPQ SI, R8
	JHI  st_avx_tail
	VMOVDQU  (SI), Y0
	VPOR     Y10, Y0, Y1
	VPCMPEQB Y13, Y0, Y0
	VPCMPEQB Y11, Y1, Y2
	VPCMPEQB Y12, Y1, Y1
	VPOR     Y2, Y0, Y0
	VPOR     Y1, Y0, Y0
	VPMOVMSKB Y0, AX
	TESTL    AX, AX
	JNZ      st_avx_found
	ADDQ     $32, SI

st_avx_tail:
	// < 32 bytes left from SI. The buffer holds at least 32 (checked on entry),
	// so classify its LAST 32 — they end at R9 and include everything from SI
	// on — and shift the lanes before SI out of the mask.
	CMPQ SI, R9
	JEQ  st_avx_notfound
	VMOVDQU  -32(R9), Y0
	VPOR     Y10, Y0, Y1
	VPCMPEQB Y13, Y0, Y0
	VPCMPEQB Y11, Y1, Y2
	VPCMPEQB Y12, Y1, Y1
	VPOR     Y2, Y0, Y0
	VPOR     Y1, Y0, Y0
	VPMOVMSKB Y0, AX
	VZEROUPPER
	MOVQ     SI, BX
	SUBQ     R8, BX                 // lanes already covered: SI - (end-32)
	SHRXQ    BX, AX, AX
	TESTL    AX, AX
	JZ       st_notfound
	TZCNTL   AX, AX
	SUBQ     DX, SI
	ADDQ     SI, AX
	MOVQ     AX, ret+32(FP)
	RET

st_avx_found64:
	VPMOVMSKB Y0, BX                // the first half's lanes
	TESTL    BX, BX
	JNZ      st_avx_found64lo
	ADDQ     $32, SI                // the match is in the second half
	JMP      st_avx_found
st_avx_found64lo:
	MOVL     BX, AX
st_avx_found:
	VZEROUPPER
	TZCNTL   AX, AX
	SUBQ     DX, SI
	ADDQ     SI, AX
	MOVQ     AX, ret+32(FP)
	RET

st_bytes:
	MOVBLZX (SI), AX
	CMPL    AX, $0x7b
	JE      st_bytefound
	CMPL    AX, $0x7d
	JE      st_bytefound
	CMPL    AX, $0x5b
	JE      st_bytefound
	CMPL    AX, $0x5d
	JE      st_bytefound
	CMPL    AX, $0x22
	JE      st_bytefound
	INCQ    SI
	CMPQ    SI, R9
	JNE     st_bytes
	JMP     st_notfound
st_bytefound:
	SUBQ    DX, SI
	MOVQ    SI, ret+32(FP)
	RET

st_avx_notfound:
	VZEROUPPER
st_notfound:
	MOVQ CX, ret+32(FP)
	RET

st_vbmi:
	// AVX-512 VBMI: VPERMB looks every byte up in structPerm and the compare
	// keeps the lanes whose byte came back unchanged (see the table). VPERMB on
	// a zmm register issues once every two cycles on Zen 4, and a zmm load that
	// is not 64-byte aligned always spans two cache lines, which measured 45%
	// slower on a load-only loop; so the first block is read where the scan
	// starts, and the bulk from the next 64-byte boundary on, two blocks a step.
	VMOVDQU64 structPerm<>(SB), Z16
	LEAQ -64(R9), R8
	CMPQ SI, R8
	JHI  st_vbmi_tail
	VMOVDQU64 (SI), Z17
	VPERMB    Z16, Z17, Z18
	VPCMPEQB  Z17, Z18, K1
	KORTESTQ  K1, K1
	JNZ       st_vbmi_found
	ADDQ      $64, SI
	ANDQ      $-64, SI              // re-scanning up to 63 clean bytes is free
	LEAQ      -128(R9), R8
	CMPQ      SI, R8
	JHI       st_vbmi_one
st_vbmi_loop:
	VMOVDQU64 (SI), Z17
	VMOVDQU64 64(SI), Z19
	VPERMB    Z16, Z17, Z18
	VPERMB    Z16, Z19, Z20
	VPCMPEQB  Z17, Z18, K1
	VPCMPEQB  Z19, Z20, K2
	KORTESTQ  K1, K2
	JNZ       st_vbmi_found2
	SUBQ      $-128, SI
	CMPQ      SI, R8
	JLS       st_vbmi_loop
st_vbmi_one:
	// < 128 bytes left: one more full block if 64 of them remain.
	LEAQ      -64(R9), R8
	CMPQ      SI, R8
	JHI       st_vbmi_tail
	VMOVDQU64 (SI), Z17
	VPERMB    Z16, Z17, Z18
	VPCMPEQB  Z17, Z18, K1
	KORTESTQ  K1, K1
	JNZ       st_vbmi_found
	ADDQ      $64, SI
st_vbmi_tail:
	// < 64 bytes left: a masked load reads exactly them (masked-off lanes are
	// zeroed and cannot fault), so there is no byte loop and no overlap.
	MOVQ      R9, BX
	SUBQ      SI, BX                // n = end - p, 0..63
	JZ        st_notfound
	MOVQ      $-1, AX
	BZHIQ     BX, AX, AX
	KMOVQ     AX, K2
	VMOVDQU8.Z (SI), K2, Z17
	VPERMB    Z16, Z17, Z18
	VPCMPEQB  Z17, Z18, K1
	KORTESTQ  K1, K1
	JZ        st_notfound
st_vbmi_found:
	KMOVQ     K1, AX
st_vbmi_foundAX:
	TZCNTQ    AX, AX
	SUBQ      DX, SI
	ADDQ      SI, AX
	MOVQ      AX, ret+32(FP)
	RET
st_vbmi_found2:
	KMOVQ     K1, AX
	TESTQ     AX, AX
	JNZ       st_vbmi_foundAX
	KMOVQ     K2, AX
	ADDQ      $64, SI
	JMP       st_vbmi_foundAX

// func indexEscapeSSE2(b []byte) int
//
// Returns the index of the first byte JSON string encoding must escape — '"'
// (0x22), '\\' (0x5c) or a control byte < 0x20 — or len(b) if none. It mirrors
// indexQuoteOrBackslashSSE2's structure (SSE2 first 32 bytes with no VZEROUPPER,
// switch to AVX2 only for a long clean run, then a 16-byte SSE loop and a scalar
// tail) and adds, per block, a PMINUB(v, 0x1f) == v test that marks control bytes:
// min(c, 0x1f) equals c exactly when c <= 0x1f. The SSE blocks take the three
// splats as aligned memory operands, as indexQuoteOrBackslashSSE2 does; the AVX2
// loop holds them in Y0 ('"'), Y1 ('\\') and Y6 (0x1f). X7/Y4 are per-block
// scratch.
TEXT ·indexEscapeSSE2(SB), NOSPLIT, $0-32
	MOVQ b_base+0(FP), SI
	MOVQ b_len+8(FP), CX
	XORQ DI, DI
	CMPQ CX, $16
	JL   esc_tail

esc_loop32:
	CMPQ CX, $32
	JL   esc_loop16
	MOVOU (SI)(DI*1), X2         // first 16 bytes
	MOVOU X2, X3
	PCMPEQB quoteMask<>(SB), X3              // == '"'
	MOVOU X2, X7
	PCMPEQB bslashMask<>(SB), X7             // == '\\'
	POR     X7, X3
	MOVOU X2, X7
	PMINUB ctrlMask<>(SB), X7              // min(v, 0x1f)
	PCMPEQB X2, X7            // == v  -> control byte
	POR     X7, X3
	PMOVMSKB X3, AX
	TESTL    AX, AX
	JNZ      esc_found
	MOVOU 16(SI)(DI*1), X4      // second 16 bytes
	MOVOU X4, X5
	PCMPEQB quoteMask<>(SB), X5
	MOVOU X4, X7
	PCMPEQB bslashMask<>(SB), X7
	POR     X7, X5
	MOVOU X4, X7
	PMINUB ctrlMask<>(SB), X7
	PCMPEQB X4, X7
	POR     X7, X5
	PMOVMSKB X5, AX
	TESTL    AX, AX
	JNZ      esc_found16
	ADDQ     $32, DI
	SUBQ     $32, CX
	MOVBLZX ·useAVX2(SB), AX
	TESTL   AX, AX
	JNZ     esc_avx_setup
	JMP     esc_loop32

esc_avx_setup:
	VMOVDQU quoteMask<>(SB), Y0
	VMOVDQU bslashMask<>(SB), Y1
	VMOVDQU ctrlMask<>(SB), Y6

esc_avx_loop:
	CMPQ CX, $32
	JL   esc_avx_done
	VMOVDQU (SI)(DI*1), Y2
	VPCMPEQB Y0, Y2, Y3
	VPCMPEQB Y1, Y2, Y4
	VPOR      Y4, Y3, Y3
	VPMINUB   Y6, Y2, Y4        // min(v, 0x1f)
	VPCMPEQB  Y2, Y4, Y4        // == v  -> control byte
	VPOR      Y4, Y3, Y3
	VPMOVMSKB Y3, AX
	TESTL     AX, AX
	JNZ       esc_avx_found
	ADDQ      $32, DI
	SUBQ      $32, CX
	JMP       esc_avx_loop

esc_avx_found:
	BSFL AX, AX
	ADDQ DI, AX
	MOVQ AX, ret+24(FP)
	VZEROUPPER
	RET

esc_avx_done:
	VZEROUPPER

esc_loop16:
	CMPQ CX, $16
	JL   esc_tail
	MOVOU (SI)(DI*1), X2
	MOVOU X2, X3
	PCMPEQB quoteMask<>(SB), X3
	MOVOU X2, X7
	PCMPEQB bslashMask<>(SB), X7
	POR     X7, X3
	MOVOU X2, X7
	PMINUB ctrlMask<>(SB), X7
	PCMPEQB X2, X7
	POR     X7, X3
	PMOVMSKB X3, AX
	TESTL    AX, AX
	JNZ      esc_found
	ADDQ     $16, DI
	SUBQ     $16, CX
	JMP      esc_loop16

esc_found16:
	ADDQ $16, DI                 // match was in the second 16-byte half

esc_found:
	BSFL AX, AX
	ADDQ DI, AX
	MOVQ AX, ret+24(FP)
	RET

esc_tail:
	TESTQ CX, CX
	JZ    esc_notfound

esc_tailloop:
	MOVBLZX (SI)(DI*1), AX
	CMPL    AX, $0x20
	JL      esc_tfound                 // control byte < 0x20
	CMPL    AX, $0x22
	JE      esc_tfound
	CMPL    AX, $0x5c
	JE      esc_tfound
	INCQ    DI
	DECQ    CX
	JNZ     esc_tailloop

esc_notfound:
	MOVQ b_len+8(FP), AX
	MOVQ AX, ret+24(FP)
	RET

esc_tfound:
	MOVQ DI, ret+24(FP)
	RET

// func indexEscapeNonASCIISSE2(b []byte) int
//
// indexEscapeSSE2's scan with the predicate widened by non-ASCII bytes: returns
// the index of the first '"' (0x22), '\\' (0x5c), control byte < 0x20 OR byte
// >= 0x80, or len(b) if none. The widening is one POR of the raw chunk into the
// match vector before each PMOVMSKB: PMOVMSKB reads only the per-lane sign bit,
// the match lanes are already 0x00/0xFF, so OR-ing the raw bytes sets exactly the
// non-ASCII lanes' sign bits — no extra compare, no extra splat. Everything else
// (SSE2 first 32 bytes, AVX2 switch for long clean runs, 16-byte loop, scalar
// tail) is byte-for-byte indexEscapeSSE2.
TEXT ·indexEscapeNonASCIISSE2(SB), NOSPLIT, $0-32
	MOVQ b_base+0(FP), SI
	MOVQ b_len+8(FP), CX
	XORQ DI, DI
	CMPQ CX, $16
	JL   escu_tail

escu_loop32:
	CMPQ CX, $32
	JL   escu_loop16
	MOVOU (SI)(DI*1), X2         // first 16 bytes
	MOVOU X2, X3
	PCMPEQB quoteMask<>(SB), X3              // == '"'
	MOVOU X2, X7
	PCMPEQB bslashMask<>(SB), X7             // == '\\'
	POR     X7, X3
	MOVOU X2, X7
	PMINUB ctrlMask<>(SB), X7              // min(v, 0x1f)
	PCMPEQB X2, X7            // == v  -> control byte
	POR     X7, X3
	POR     X2, X3             // raw sign bits -> non-ASCII lanes
	PMOVMSKB X3, AX
	TESTL    AX, AX
	JNZ      escu_found
	MOVOU 16(SI)(DI*1), X4      // second 16 bytes
	MOVOU X4, X5
	PCMPEQB quoteMask<>(SB), X5
	MOVOU X4, X7
	PCMPEQB bslashMask<>(SB), X7
	POR     X7, X5
	MOVOU X4, X7
	PMINUB ctrlMask<>(SB), X7
	PCMPEQB X4, X7
	POR     X7, X5
	POR     X4, X5             // raw sign bits -> non-ASCII lanes
	PMOVMSKB X5, AX
	TESTL    AX, AX
	JNZ      escu_found16
	ADDQ     $32, DI
	SUBQ     $32, CX
	MOVBLZX ·useAVX2(SB), AX
	TESTL   AX, AX
	JNZ     escu_avx_setup
	JMP     escu_loop32

escu_avx_setup:
	VMOVDQU quoteMask<>(SB), Y0
	VMOVDQU bslashMask<>(SB), Y1
	VMOVDQU ctrlMask<>(SB), Y6

escu_avx_loop:
	CMPQ CX, $32
	JL   escu_avx_done
	VMOVDQU (SI)(DI*1), Y2
	VPCMPEQB Y0, Y2, Y3
	VPCMPEQB Y1, Y2, Y4
	VPOR      Y4, Y3, Y3
	VPMINUB   Y6, Y2, Y4        // min(v, 0x1f)
	VPCMPEQB  Y2, Y4, Y4        // == v  -> control byte
	VPOR      Y4, Y3, Y3
	VPOR      Y2, Y3, Y3        // raw sign bits -> non-ASCII lanes
	VPMOVMSKB Y3, AX
	TESTL     AX, AX
	JNZ       escu_avx_found
	ADDQ      $32, DI
	SUBQ      $32, CX
	JMP       escu_avx_loop

escu_avx_found:
	BSFL AX, AX
	ADDQ DI, AX
	MOVQ AX, ret+24(FP)
	VZEROUPPER
	RET

escu_avx_done:
	VZEROUPPER

escu_loop16:
	CMPQ CX, $16
	JL   escu_tail
	MOVOU (SI)(DI*1), X2
	MOVOU X2, X3
	PCMPEQB quoteMask<>(SB), X3
	MOVOU X2, X7
	PCMPEQB bslashMask<>(SB), X7
	POR     X7, X3
	MOVOU X2, X7
	PMINUB ctrlMask<>(SB), X7
	PCMPEQB X2, X7
	POR     X7, X3
	POR     X2, X3             // raw sign bits -> non-ASCII lanes
	PMOVMSKB X3, AX
	TESTL    AX, AX
	JNZ      escu_found
	ADDQ     $16, DI
	SUBQ     $16, CX
	JMP      escu_loop16

escu_found16:
	ADDQ $16, DI                 // match was in the second 16-byte half

escu_found:
	BSFL AX, AX
	ADDQ DI, AX
	MOVQ AX, ret+24(FP)
	RET

escu_tail:
	TESTQ CX, CX
	JZ    escu_notfound

escu_tailloop:
	// Sign-extended load + one signed compare covers control AND non-ASCII
	// bytes: as int8, 0x80..0xFF are negative and 0x00..0x1F are below 0x20,
	// while clean ASCII 0x20..0x7F is not — so the tail costs exactly the same
	// three compares per byte as indexEscapeSSE2's.
	MOVBLSX (SI)(DI*1), AX
	CMPL    AX, $0x20
	JL      escu_tfound                // control byte or non-ASCII byte
	CMPL    AX, $0x22
	JE      escu_tfound
	CMPL    AX, $0x5c
	JE      escu_tfound
	INCQ    DI
	DECQ    CX
	JNZ     escu_tailloop

escu_notfound:
	MOVQ b_len+8(FP), AX
	MOVQ AX, ret+24(FP)
	RET

escu_tfound:
	MOVQ DI, ret+24(FP)
	RET
