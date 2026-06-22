//go:build arm64

package unstable

// fastSkipAvail is always true on arm64: Advanced SIMD (NEON) is mandatory in the
// ARMv8-A baseline Go targets, so maskBlock's NEON routine is always present.
//
// NOTE: the arm64 NEON maskBlock is correctness-verified (under qemu, via the
// differential fuzz) but its speed is UNMEASURED on real hardware — like the
// other arm64 SIMD here, it wants a real-arm64 benchstat run to confirm the win
// before being relied on for performance.
const fastSkipAvail = true

// maskBlock returns the character-class bitmaps for b[:32] (see skipfast.go).
// NEON implementation in skipfast_arm64.s.
//
//go:noescape
func maskBlock(b []byte) (quote, bslash, lbrace, rbrace, lbrack, rbrack uint32)
