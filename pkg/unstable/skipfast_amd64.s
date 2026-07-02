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

// skipBlocksAVX2 / skipBlocksAVX512 run skipContainerFast's whole 64-byte block
// loop in assembly (see the Go declarations in skipfast_amd64.go): character
// classification, the escape / in-string bit math, and the bracket balancing,
// with the splats loaded once and the carried state (depth, prevEscaped,
// prevInString) held in registers across blocks — where the maskBlock form paid
// a call, four result stores/loads, five splat loads and an isArray branch per
// block. The prefix XOR that turns quote positions into the inside-string mask
// is a single carryless multiply by all-ones (VPCLMULQDQ) instead of the
// six-step shift/XOR chain, shortening the loop-carried dependency the Go loop
// is latency-bound on.
//
// Register map (both variants): SI=data base, DX=pos, CX=len-64 (last valid
// block start), R8=depth, R9=prevEscaped (0/1), R10=prevInString (0/all-ones),
// DI=evenBits 0x5555..., X7=all-ones CLMUL operand; R11-R14 hold the block's
// quote/bslash/open/close bitmaps, AX/BX/R15 scratch.
//
// BLOCKTAIL is the shared per-block bit math: the Go loop's escaped -> inStr ->
// bracket-balance steps, including its fast paths (no backslash pending: skip
// findEscaped; no unescaped quote: in-string state unchanged; popcount bulk
// depth update whenever the block cannot cross depth 0). Expanded once per
// TEXT symbol, so its labels stay function-scoped. Falls through to nextBlock
// (advance and loop) or exits at foundEnd.
#define BLOCKTAIL \
	MOVQ  R12, AX               \ // escaped: skip the add-carry chain when no
	ORQ   R9, AX                \ // backslash in the block and none pending
	JNE   computeEscaped        \
	XORQ  R15, R15              \
	JMP   haveEscaped           \
computeEscaped:                 \
	ANDNQ R12, R9, R12          \ // backslash &^= prevEscaped
	LEAQ  (R9)(R12*2), BX       \ // followsEscape = backslash<<1 | prevEscaped
	ANDNQ R12, DI, AX           \ // backslash &^ evenBits
	ANDNQ AX, BX, AX            \ //   ... &^ followsEscape = oddSequenceStarts
	ADDQ  R12, AX               \ // seq = odd + backslash, CF = carry out
	MOVQ  $0, R9                \ // (MOV keeps flags)
	ADCQ  $0, R9                \ // prevEscaped = carry
	SHLQ  $1, AX                \ // invertMask = seq << 1
	XORQ  DI, AX                \
	ANDQ  BX, AX                \ // escaped = (evenBits ^ invertMask) & follows
	MOVQ  AX, R15               \
haveEscaped:                    \
	ANDNQ R11, R15, AX          \ // q = quote &^ escaped
	TESTQ AX, AX                \ // no unescaped quote: in-string mask carries
	JNE   computeInStr          \
	MOVQ  R10, BX               \
	JMP   haveInStr             \
computeInStr:                   \
	VMOVQ AX, X4                \
	VPCLMULQDQ $0x00, X7, X4, X4 \ // prefix XOR: clmul by all-ones, low 64
	VMOVQ X4, BX                \
	XORQ  R10, BX               \ // inStr = prefixXor(q) ^ prevInString
	MOVQ  BX, R10               \
	SARQ  $63, R10              \ // carry: inside-string at byte 63
haveInStr:                      \
	ANDNQ R13, BX, R13          \ // op &^= inStr
	ANDNQ R14, BX, R14          \ // cl &^= inStr
	TESTQ R14, R14              \
	JNE   hasClose              \
	POPCNTQ R13, AX             \ // no close: depth += popcount(op), no crossing
	ADDQ  AX, R8                \
	JMP   nextBlock             \
hasClose:                       \
	POPCNTQ R14, AX             \
	CMPQ  R8, AX                \
	JLE   bitLoopSetup          \ // depth <= closes: a crossing is possible
	POPCNTQ R13, BX             \ // depth > closes: bulk update, min depth >= 1
	ADDQ  BX, R8                \
	SUBQ  AX, R8                \
	JMP   nextBlock             \
bitLoopSetup:                   \
	MOVQ  R13, AX               \
	ORQ   R14, AX               \ // brackets, walked in byte order
bitLoop:                        \
	TESTQ AX, AX                \
	JE    nextBlock             \
	TZCNTQ AX, BX               \ // j = next bracket
	BTQ   BX, R13               \ // CF = open at j
	JC    isOpen                \
	DECQ  R8                    \
	JE    foundEnd              \ // depth == 0: close of the container
	JMP   clearBit              \
isOpen:                         \
	INCQ  R8                    \
clearBit:                       \
	BLSRQ AX, AX                \ // brackets &= brackets-1
	JMP   bitLoop               \
foundEnd:                       \
	LEAQ  1(DX)(BX*1), AX       \ // end = pos + j + 1
	MOVQ  AX, end+48(FP)        \
	MOVQ  R8, ndepth+56(FP)     \
	MOVQ  R9, prevEscaped+64(FP) \
	MOVQ  R10, prevInString+72(FP) \
	VZEROUPPER                  \
	RET                         \
nextBlock:                      \
	ADDQ  $64, DX               \
	JMP   blockLoop

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


// func skipBlocksAVX2(data []byte, pos, depth int, isArray bool) (end, ndepth int, prevEscaped, prevInString uint64)
TEXT ·skipBlocksAVX2(SB), NOSPLIT, $0-80
	MOVQ    data_base+0(FP), SI
	MOVQ    data_len+8(FP), CX
	SUBQ    $64, CX
	MOVQ    pos+24(FP), DX
	MOVQ    depth+32(FP), R8
	XORQ    R9, R9
	XORQ    R10, R10
	MOVQ    $0x5555555555555555, DI
	VPCMPEQB X7, X7, X7         // all-ones CLMUL operand
	VMOVDQU mbQuote<>(SB), Y10
	VMOVDQU mbBslash<>(SB), Y11
	MOVBLZX isArray+40(FP), AX
	TESTL   AX, AX
	JNZ     arrayBrackets
	VMOVDQU mbLbrace<>(SB), Y12
	VMOVDQU mbRbrace<>(SB), Y13
	JMP     blockLoop
arrayBrackets:
	VMOVDQU mbLbrack<>(SB), Y12
	VMOVDQU mbRbrack<>(SB), Y13
blockLoop:
	CMPQ    DX, CX
	JGT     exhausted
	VMOVDQU (SI)(DX*1), Y0
	VMOVDQU 32(SI)(DX*1), Y1
	// quote bitmap -> R11: compare both halves, fold the two 32-bit movemasks.
	VPCMPEQB Y10, Y0, Y2
	VPMOVMSKB Y2, R11
	VPCMPEQB Y10, Y1, Y3
	VPMOVMSKB Y3, AX
	SHLQ    $32, AX
	ORQ     AX, R11
	// backslash -> R12
	VPCMPEQB Y11, Y0, Y2
	VPMOVMSKB Y2, R12
	VPCMPEQB Y11, Y1, Y3
	VPMOVMSKB Y3, AX
	SHLQ    $32, AX
	ORQ     AX, R12
	// open -> R13
	VPCMPEQB Y12, Y0, Y2
	VPMOVMSKB Y2, R13
	VPCMPEQB Y12, Y1, Y3
	VPMOVMSKB Y3, AX
	SHLQ    $32, AX
	ORQ     AX, R13
	// close -> R14
	VPCMPEQB Y13, Y0, Y2
	VPMOVMSKB Y2, R14
	VPCMPEQB Y13, Y1, Y3
	VPMOVMSKB Y3, AX
	SHLQ    $32, AX
	ORQ     AX, R14
	BLOCKTAIL
exhausted:
	MOVQ    $-1, end+48(FP)
	MOVQ    R8, ndepth+56(FP)
	MOVQ    R9, prevEscaped+64(FP)
	MOVQ    R10, prevInString+72(FP)
	VZEROUPPER
	RET

// func skipBlocksAVX512(data []byte, pos, depth int, isArray bool) (end, ndepth int, prevEscaped, prevInString uint64)
//
// The AVX-512 form of the same loop: one 64-byte load, and each class is a
// VPCMPEQB straight into a k-mask plus a KMOVQ — two instructions per class
// where AVX2 needs seven. Splats are broadcast from immediates once (AVX512BW's
// GP-source VPBROADCASTB). Only ZMM0-15 are used, so the closing VZEROUPPER
// covers all dirtied upper state.
TEXT ·skipBlocksAVX512(SB), NOSPLIT, $0-80
	MOVQ    data_base+0(FP), SI
	MOVQ    data_len+8(FP), CX
	SUBQ    $64, CX
	MOVQ    pos+24(FP), DX
	MOVQ    depth+32(FP), R8
	XORQ    R9, R9
	XORQ    R10, R10
	MOVQ    $0x5555555555555555, DI
	VPCMPEQB X7, X7, X7         // all-ones CLMUL operand
	MOVL    $0x22, AX           // '"'
	VPBROADCASTB AX, Z10
	MOVL    $0x5c, AX           // '\\'
	VPBROADCASTB AX, Z11
	MOVBLZX isArray+40(FP), AX
	TESTL   AX, AX
	JNZ     arrayBrackets
	MOVL    $0x7b, AX           // '{'
	VPBROADCASTB AX, Z12
	MOVL    $0x7d, AX           // '}'
	VPBROADCASTB AX, Z13
	JMP     blockLoop
arrayBrackets:
	MOVL    $0x5b, AX           // '['
	VPBROADCASTB AX, Z12
	MOVL    $0x5d, AX           // ']'
	VPBROADCASTB AX, Z13
blockLoop:
	CMPQ    DX, CX
	JGT     exhausted
	VMOVDQU64 (SI)(DX*1), Z0
	VPCMPEQB Z10, Z0, K1
	KMOVQ   K1, R11
	VPCMPEQB Z11, Z0, K2
	KMOVQ   K2, R12
	VPCMPEQB Z12, Z0, K3
	KMOVQ   K3, R13
	VPCMPEQB Z13, Z0, K4
	KMOVQ   K4, R14
	BLOCKTAIL
exhausted:
	MOVQ    $-1, end+48(FP)
	MOVQ    R8, ndepth+56(FP)
	MOVQ    R9, prevEscaped+64(FP)
	MOVQ    R10, prevInString+72(FP)
	VZEROUPPER
	RET
