//go:build arm64

package unstable

// fastSkipAvail is always true on arm64: Advanced SIMD (NEON) is mandatory in the
// ARMv8-A baseline Go targets, so maskBlock's NEON routine is always present.
//
// Measured on Apple M2 (BenchmarkGetSkipHeavy / BenchmarkSkipContainer): the NEON
// container skip beats the per-string indexStructural path by ~30–40% on dense
// object/record containers (and is flat on scalar arrays, which take the
// indexStructural path anyway). The weight-and-fold gather in maskBlock (see the
// .s) is what makes it pay.
const fastSkipAvail = true

// maskBlock returns the character-class bitmaps for b[:32] (see skipfast.go).
// NEON implementation in skipfast_arm64.s.
//
//go:noescape
func maskBlock(b []byte) (quote, bslash, lbrace, rbrace, lbrack, rbrack uint32)
