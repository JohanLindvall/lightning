//go:build amd64

package unstable

import "golang.org/x/sys/cpu"

// fastSkipAvail reports whether skipContainerFast's maskBlock has a SIMD
// implementation on this machine. The AVX2 routine requires AVX2; without it the
// scalar maskBlock is slower than the current indexStructural skip, so SkipValue
// stays on the latter.
var fastSkipAvail = useAVX2

// maskBlock returns the character-class bitmaps for b[:64] (see skipfast.go):
// quote, backslash, and the container's own open/close brackets (`[`/`]` when
// isArray, else `{`/`}`). AVX2 implementation in skipfast_amd64.s; used only when
// fastSkipAvail (AVX2 present) and only by the Go block loop, which is the
// fallback when the whole-loop skipBlocks assembly below is unavailable.
//
//go:noescape
func maskBlock(b []byte, isArray bool) (quote, bslash, open, close uint64)

// useSkipBlocks gates the whole-loop assembly form of the block scan
// (skipBlocks and the AVX-512 body it selects): the per-block character classes, the
// escape/in-string bit math, and the bracket balancing all run in one assembly
// loop with the splats loaded once and the carried state in registers,
// removing the per-block Go<->asm call and result marshaling the maskBlock
// form pays. PCLMULQDQ does the prefix XOR in one carryless multiply; BMI1
// (ANDN/TZCNT/BLSR) and POPCNT carry the bit math. All of these ship on every
// AVX2-capable CPU, so in practice this is the same population as useAVX2 —
// the extra gates are just correctness belts.
var useSkipBlocks = useAVX2 && cpu.X86.HasPCLMULQDQ && cpu.X86.HasBMI1 && cpu.X86.HasBMI2 && cpu.X86.HasPOPCNT

// skipBlocksTakesTail reports that skipBlocks consumes the whole buffer from
// pos, the final < 64 bytes included: the AVX2 body reads them as the buffer's
// last 64 bytes with the lanes before pos shifted out of every bitmap, the
// AVX-512 body with a masked load. So a caller that gets end < 0 back has hit
// the end of the input, not the end of the full blocks, and the Go
// continuation (a frame, and a maskBlock call to build that last block) is
// never reached. Both bodies are handed a buffer of at least 64 bytes: the
// AVX2 one needs them for the overlapping block, and the AVX-512 one, whose
// masked load needs none, is held to the same rule so that the two agree on
// every input — under one block the byte walk decides, and on malformed input
// it reads a stray backslash differently (see TestSkipBackslashLengthCliff).
const skipBlocksTakesTail = true

// useSkipBlocks512 selects the AVX-512 variant: one 64-byte load and a
// VPCMPEQB straight into a k-mask per class (two instructions) instead of
// AVX2's two compares + two movemasks + shift/or fold (seven).
var useSkipBlocks512 = useSkipBlocks && cpu.X86.HasAVX512BW

// skipBlocks runs the balance scan over data[pos:], starting at the given
// depth inside the container ('['/']' when isArray, else '{'/'}'): every full
// 64-byte block, and then the final < 64 bytes as one more block (see
// skipBlocksTakesTail; the AVX2 body needs len(data) >= 64 for it). If the
// container's close is found it returns its index + 1 in end; otherwise end is
// -1 and the input ran out inside the container. A caller that must resume
// in a later buffer — the streaming ValueScanner — hands it a buffer cut at
// the last full block, so that no tail is taken and ndepth/prevEscaped/
// prevInString are the state at that cut, exact for the byte after it.
//
// The three carried results are written ONLY on that failing return: once the
// close is found they are dead — skipContainerFast reads them under end < 0 and
// nowhere else — so writing them cost three stores on the path every small
// container takes. Anything reading them when end >= 0 reads the frame's
// leftovers.
//
// It IS the assembly — the AVX-512 selection is made inside it, off
// useSkipBlocks512, exactly as the SSE2/AVX2 and NEON/SVE2 scanners select
// their bodies. A Go wrapper holding the two calls was what stood here, and it
// could not inline (cost 164 against the budget of 80), so every container skip
// paid a second frame for a branch; see the comment on the TEXT symbol.
//
//go:noescape
func skipBlocks(data []byte, pos, depth int, isArray bool) (end, ndepth int, prevEscaped, prevInString uint64)

// skipBlocksAVX512 is reached only by skipBlocks' tail JMP; nothing in Go calls
// it. The declaration is still load-bearing — asmdecl validates the assembly's
// frame offsets against exactly this signature, and go vet reports assembly
// with no Go prototype — but the `unused` linter reads Go only and sees a dead
// declaration, the same blind spot the NEON routines in simd_arm64.go carry.
//
//nolint:unused // called from assembly; see above
//go:noescape
func skipBlocksAVX512(data []byte, pos, depth int, isArray bool) (end, ndepth int, prevEscaped, prevInString uint64)
