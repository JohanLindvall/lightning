//go:build arm64

package unstable

// fastSkipAvail is always true on arm64: Advanced SIMD (NEON) is mandatory in the
// ARMv8-A baseline Go targets, so maskBlock's NEON routine is always present.
//
// Measured on Apple M2 (BenchmarkGetSkipHeavy / BenchmarkSkipContainer): the NEON
// container skip beats the per-string indexStructural path on dense object/record
// containers (and is flat on scalar arrays, which take the indexStructural path
// anyway). The weight-and-fold gather in maskBlock (see the .s) is what makes it
// pay. With the whole-loop skipBlocksNEON on top (useSkipBlocks below), the live
// dispatch runs 10-15 GB/s on the object shapes (BenchmarkSkipContainer, M2).
const fastSkipAvail = true

// maskBlock returns the character-class bitmaps for b[:64] (see skipfast.go):
// quote, backslash, and the container's own open/close brackets (`[`/`]` when
// isArray, else `{`/`}`). NEON implementation in skipfast_arm64.s.
//
//go:noescape
func maskBlock(b []byte, isArray bool) (quote, bslash, open, close uint64)

// useSkipBlocks gates the whole-loop NEON block scan (skipBlocksNEON): the
// character classes, escape/in-string bit math and bracket balancing run in
// one assembly loop with the splats and bit-weight vector loaded once and the
// carried state in registers — removing the per-block maskBlock call, its four
// results through memory, and its five per-call VDUP splats. NEON is baseline
// on arm64, so this is unconditional; it is a var (not const) so the variant
// tests can force the Go maskBlock loop for differential comparison.
//
// Measured on Apple M2 (BenchmarkSkipBlocksVariant, n=8, benchstat): the NEON
// whole-loop beats the Go maskBlock loop stringObj -34.0%, numberObj -34.7%,
// nestedMixed -28.7% (geomean -32.5%, all p=0.000).
//
// THIS LOOP STAYS NEON EVEN ON AN SVE2 CORE, unlike the four scanners in
// simd_arm64.s. It is not an oversight: this algorithm needs a 64-bit BITMAP per
// character class, because the escape and in-string masks are built with GP
// bit math (add-with-carry and a prefix XOR) and the brackets are balanced by
// walking set bits. SVE has no cheap predicate-to-general-register extraction
// before SVE2.1's PMOV, which these cores do not have — CNTP yields only a
// population count, not the mask. The one way to get the bits out is to STR the
// predicate to the stack and reload it, which at a 128-bit vector length means
// four 2-byte stores feeding one 8-byte load: that cannot store-forward on any
// core, and the stall would land directly on the loop-carried
// depth/prevEscaped/prevInString chain. That is the same trade already rejected
// for this loop's cross-domain VMOVs (see CLAUDE.md), so SVE2 would have to pay
// it four times over to remove them. Do not port this loop to SVE without a
// predicate-to-GP move.
var useSkipBlocks = true

func skipBlocks(data []byte, pos, depth int, isArray bool) (end, ndepth int, prevEscaped, prevInString uint64) {
	return skipBlocksNEON(data, pos, depth, isArray)
}

//go:noescape
func skipBlocksNEON(data []byte, pos, depth int, isArray bool) (end, ndepth int, prevEscaped, prevInString uint64)
