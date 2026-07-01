//go:build arm64

package unstable

// fastSkipAvail is always true on arm64: Advanced SIMD (NEON) is mandatory in the
// ARMv8-A baseline Go targets, so maskBlock's NEON routine is always present.
//
// Measured on Apple M2 (BenchmarkGetSkipHeavy / BenchmarkSkipContainer): the NEON
// container skip beats the per-string indexStructural path on dense object/record
// containers (and is flat on scalar arrays, which take the indexStructural path
// anyway). The weight-and-fold gather in maskBlock (see the .s) is what makes it
// pay. NOTE: this version processes 64 bytes / 4 masks per call (was 32 / 6); the
// M2 numbers above predate that change and want a re-measure.
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
var useSkipBlocks = true

func skipBlocks(data []byte, pos, depth int, isArray bool) (end, ndepth int, prevEscaped, prevInString uint64) {
	return skipBlocksNEON(data, pos, depth, isArray)
}

//go:noescape
func skipBlocksNEON(data []byte, pos, depth int, isArray bool) (end, ndepth int, prevEscaped, prevInString uint64)
