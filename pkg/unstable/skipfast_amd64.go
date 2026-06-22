//go:build amd64

package unstable

// fastSkipAvail reports whether skipContainerFast's maskBlock has a SIMD
// implementation on this machine. The AVX2 routine requires AVX2; without it the
// scalar maskBlock is slower than the current indexStructural skip, so SkipValue
// stays on the latter.
var fastSkipAvail = useAVX2

// maskBlock returns the character-class bitmaps for b[:64] (see skipfast.go):
// quote, backslash, and the container's own open/close brackets (`[`/`]` when
// isArray, else `{`/`}`). AVX2 implementation in skipfast_amd64.s; used only when
// fastSkipAvail (AVX2 present).
//
//go:noescape
func maskBlock(b []byte, isArray bool) (quote, bslash, open, close uint64)
