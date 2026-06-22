//go:build amd64

package unstable

// fastSkipAvail reports whether skipContainerFast's maskBlock has a SIMD
// implementation on this machine. The AVX2 routine requires AVX2; without it the
// scalar maskBlock is slower than the current indexStructural skip, so SkipValue
// stays on the latter.
var fastSkipAvail = useAVX2

// maskBlock returns the character-class bitmaps for b[:32] (see skipfast.go).
// AVX2 implementation in skipfast_amd64.s. NOTE (prototype): assumes AVX2; gate
// on useAVX2 before productionizing.
//
//go:noescape
func maskBlock(b []byte) (quote, bslash, lbrace, rbrace, lbrack, rbrack uint32)
