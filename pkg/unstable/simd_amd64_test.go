//go:build amd64

package unstable

import (
	"math/rand"
	"strings"
	"testing"
)

// TestIndexVariantsFlip differentially tests both dispatch arms of the three
// string scanners against the scalar oracle — the live-flag run in
// TestIndexFunctionsMatchScalar only exercises the widest path the machine
// supports, so the pure-SSE2 long-string loop (taken when useAVX2 is false)
// would otherwise go untested on any modern box. Mirrors TestSkipBlocksVariants'
// flag-flip pattern. Inputs: every target byte at every offset of a 320-byte
// clean base (the AVX2 tail runs several full iterations, with leftovers for
// the 16-byte and scalar tails), plus random mixtures whose special bytes land
// anywhere relative to the 32-byte block boundaries. 0x20 (space) sits just
// past indexEscape's control-byte boundary and must never match.
//
// An AVX-512BW 64-byte kmask tail for these loops was implemented, verified,
// and reverted: interleaved A/B measured string_unicode +1.55% (p=0.040, n=16)
// and gsoc_2018 flat — Zen 4 double-pumps zmm ops on a 256-bit datapath and the
// loop is load-port-bound, so unlike skipBlocksAVX512 (which cut 7 instructions
// per class to 2) the wide form removes almost nothing while the added code
// shifts alignment. See CLAUDE.md's tried-and-rejected entry.
func TestIndexVariantsFlip(t *testing.T) {
	avx2 := useAVX2
	defer func() { useAVX2 = avx2 }()

	variants := []struct {
		name string
		avx2 bool
	}{
		{"sse2-only", false},
		{"avx2", true},
	}
	base := strings.Repeat("abcdefgh", 40) // 320 bytes, no special bytes
	rng := rand.New(rand.NewSource(512))
	for _, v := range variants {
		if v.avx2 && !avx2 {
			continue // machine lacks AVX2
		}
		useAVX2 = v.avx2
		for _, c := range []byte{'"', '\\', 0x00, 0x1f, ' ', 'z', 0x7f, 0x80, 0xff} {
			for pos := 0; pos < len(base); pos++ {
				b := []byte(base)
				b[pos] = c
				if got, want := indexQuoteOrBackslashSSE2(b, 0), indexCloseOrEscapeScalar(b); got != want {
					t.Fatalf("%s: indexQuoteOrBackslash(%q@%d) = %d, want %d", v.name, c, pos, got, want)
				}
				if got, want := indexEscapeSSE2(b), indexEscapeScalar(b); got != want {
					t.Fatalf("%s: indexEscape(%q@%d) = %d, want %d", v.name, c, pos, got, want)
				}
				if got, want := indexEscapeNonASCIISSE2(b), indexEscapeNonASCIIScalar(b); got != want {
					t.Fatalf("%s: indexEscapeNonASCII(%q@%d) = %d, want %d", v.name, c, pos, got, want)
				}
			}
		}
		for iter := 0; iter < 5000; iter++ {
			n := rng.Intn(400)
			b := make([]byte, n)
			for i := range b {
				if rng.Intn(24) == 0 {
					b[i] = []byte{'"', '\\', 0x07, 0x1f, ' ', 0x7f, 0x80, 0xc3, 0xff}[rng.Intn(9)]
				} else {
					b[i] = byte('a' + rng.Intn(26))
				}
			}
			if got, want := indexQuoteOrBackslashSSE2(b, 0), indexCloseOrEscapeScalar(b); got != want {
				t.Fatalf("%s: random len %d: indexQuoteOrBackslash = %d, want %d (%q)", v.name, n, got, want, b)
			}
			if got, want := indexEscapeSSE2(b), indexEscapeScalar(b); got != want {
				t.Fatalf("%s: random len %d: indexEscape = %d, want %d (%q)", v.name, n, got, want, b)
			}
			if got, want := indexEscapeNonASCIISSE2(b), indexEscapeNonASCIIScalar(b); got != want {
				t.Fatalf("%s: random len %d: indexEscapeNonASCII = %d, want %d (%q)", v.name, n, got, want, b)
			}
		}
	}
}
