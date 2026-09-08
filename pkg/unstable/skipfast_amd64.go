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
var useSkipBlocks = useAVX2 && cpu.X86.HasPCLMULQDQ && cpu.X86.HasBMI1 && cpu.X86.HasPOPCNT

// useSkipBlocks512 selects the AVX-512 variant: one 64-byte load and a
// VPCMPEQB straight into a k-mask per class (two instructions) instead of
// AVX2's two compares + two movemasks + shift/or fold (seven).
var useSkipBlocks512 = useSkipBlocks && cpu.X86.HasAVX512BW

// skipBlocks runs the balance scan over every full 64-byte block of
// data[pos:], starting at the given depth inside the container ('['/']' when
// isArray, else '{'/'}'). If the container's close is found it returns its
// index + 1 in end; otherwise end is -1 and ndepth/prevEscaped/prevInString
// return the carried state for the caller's scalar tail, which resumes at
// pos + (len(data)-pos)&^63.
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
