//go:build arm64

package unstable

import (
	"math/rand"
	"strings"
	"testing"
)

// TestIndexVariantsFlip differentially tests both dispatch arms of the four
// arm64 scanners against the scalar oracle. The live-flag run in
// TestIndexFunctionsMatchScalar only exercises whichever body the machine
// selects, so on an SVE2 core the NEON routines — still the path every Apple
// M-series and Neoverse N1/V1 CPU takes — would otherwise go untested, and on a
// NEON-only core the SVE2 bodies cannot run at all. Mirrors the amd64 test of
// the same name and TestSkipBlocksVariants' flag-flip pattern.
//
// Inputs: every target byte at every offset of a 320-byte clean base, so a match
// lands at every position relative to a block boundary at any vector length
// (16-, 32- or 64-byte blocks all divide into 320), plus random mixtures. The
// boundary bytes are chosen to pin the exact edges each scanner disagrees on:
// 0x20 (space) sits just past indexEscape's control-byte test and must never
// match, and 0x7f/0x80 straddle indexEscapeNonASCII's ASCII edge.
//
// Lengths are swept separately in TestIndexArm64Lengths below; here the length
// is fixed so the position sweep stays cheap.
func TestIndexVariantsFlip(t *testing.T) {
	sve := useSVE2
	defer func() { useSVE2 = sve }()

	variants := []struct {
		name string
		sve2 bool
	}{
		{"neon", false},
		{"sve2", true},
	}
	base := strings.Repeat("abcdefgh", 40) // 320 bytes, no special bytes
	rng := rand.New(rand.NewSource(512))
	for _, v := range variants {
		if v.sve2 && !sve {
			continue // machine lacks SVE2; its instructions would trap
		}
		useSVE2 = v.sve2
		for _, c := range []byte{'"', '\\', '{', '}', '[', ']', 0x00, 0x1f, ' ', 'z', 0x7f, 0x80, 0xff} {
			for pos := 0; pos < len(base); pos++ {
				b := []byte(base)
				b[pos] = c
				if got, want := indexCloseOrEscape(b), indexCloseOrEscapeScalar(b); got != want {
					t.Fatalf("%s: indexCloseOrEscape(%q@%d) = %d, want %d", v.name, c, pos, got, want)
				}
				// The offset form, on both bodies: a nonzero start moves the
				// match relative to the peeled first block and the unrolled
				// loop, which is exactly where the two bodies differ, and is
				// how every caller actually uses the scanner.
				for _, start := range []int{1, 15, 16, 17, 31, 32, 33, 63, 64, 65, 127, 128, 129} {
					if start > len(b) {
						continue
					}
					want := start + indexCloseOrEscapeScalar(b[start:])
					if got := indexCloseOrEscapeAt(b, start); got != want {
						t.Fatalf("%s: indexCloseOrEscapeAt(%q@%d, start %d) = %d, want %d", v.name, c, pos, start, got, want)
					}
				}
				if got, want := indexEscape(b), indexEscapeScalar(b); got != want {
					t.Fatalf("%s: indexEscape(%q@%d) = %d, want %d", v.name, c, pos, got, want)
				}
				if got, want := indexEscapeNonASCII(b), indexEscapeNonASCIIScalar(b); got != want {
					t.Fatalf("%s: indexEscapeNonASCII(%q@%d) = %d, want %d", v.name, c, pos, got, want)
				}
				if got, want := indexStructural(b), indexStructuralScalar(b); got != want {
					t.Fatalf("%s: indexStructural(%q@%d) = %d, want %d", v.name, c, pos, got, want)
				}
			}
		}
		for iter := 0; iter < 5000; iter++ {
			n := rng.Intn(400)
			b := make([]byte, n)
			for i := range b {
				if rng.Intn(24) == 0 {
					b[i] = []byte{'"', '\\', '{', '}', '[', ']', 0x07, 0x1f, ' ', 0x7f, 0x80, 0xc3, 0xff}[rng.Intn(13)]
				} else {
					b[i] = byte('a' + rng.Intn(26))
				}
			}
			if got, want := indexCloseOrEscape(b), indexCloseOrEscapeScalar(b); got != want {
				t.Fatalf("%s: random len %d: indexCloseOrEscape = %d, want %d (%q)", v.name, n, got, want, b)
			}
			if got, want := indexEscape(b), indexEscapeScalar(b); got != want {
				t.Fatalf("%s: random len %d: indexEscape = %d, want %d (%q)", v.name, n, got, want, b)
			}
			if got, want := indexEscapeNonASCII(b), indexEscapeNonASCIIScalar(b); got != want {
				t.Fatalf("%s: random len %d: indexEscapeNonASCII = %d, want %d (%q)", v.name, n, got, want, b)
			}
			if got, want := indexStructural(b), indexStructuralScalar(b); got != want {
				t.Fatalf("%s: random len %d: indexStructural = %d, want %d (%q)", v.name, n, got, want, b)
			}
		}
	}
}

// TestIndexArm64Lengths sweeps every buffer length up to 200 with the target
// byte at every position within it, for both dispatch arms. Length is the axis
// the SVE2 bodies changed most: they have no scalar tail and no short-input
// entry at all, covering the ragged end of the buffer with the WHILELO predicate
// instead, so a length that leaves 1..VL-1 bytes in the final block is exactly
// where a predicate mistake would show — including len 0, which must return 0
// without loading anything.
func TestIndexArm64Lengths(t *testing.T) {
	sve := useSVE2
	defer func() { useSVE2 = sve }()

	for _, v := range []struct {
		name string
		sve2 bool
	}{{"neon", false}, {"sve2", true}} {
		if v.sve2 && !sve {
			continue
		}
		useSVE2 = v.sve2
		for n := 0; n <= 200; n++ {
			base := make([]byte, n)
			for i := range base {
				base[i] = byte('a' + i%26)
			}
			if got, want := indexCloseOrEscape(base), indexCloseOrEscapeScalar(base); got != want {
				t.Fatalf("%s: len %d clean: indexCloseOrEscape = %d, want %d", v.name, n, got, want)
			}
			if got, want := indexEscape(base), indexEscapeScalar(base); got != want {
				t.Fatalf("%s: len %d clean: indexEscape = %d, want %d", v.name, n, got, want)
			}
			if got, want := indexEscapeNonASCII(base), indexEscapeNonASCIIScalar(base); got != want {
				t.Fatalf("%s: len %d clean: indexEscapeNonASCII = %d, want %d", v.name, n, got, want)
			}
			if got, want := indexStructural(base), indexStructuralScalar(base); got != want {
				t.Fatalf("%s: len %d clean: indexStructural = %d, want %d", v.name, n, got, want)
			}
			for pos := 0; pos < n; pos++ {
				for _, c := range []byte{'"', '\\', '{', ']', 0x00, 0x1f, 0x80} {
					b := append([]byte(nil), base...)
					b[pos] = c
					if got, want := indexCloseOrEscape(b), indexCloseOrEscapeScalar(b); got != want {
						t.Fatalf("%s: len %d %q@%d: indexCloseOrEscape = %d, want %d", v.name, n, c, pos, got, want)
					}
					if got, want := indexEscape(b), indexEscapeScalar(b); got != want {
						t.Fatalf("%s: len %d %q@%d: indexEscape = %d, want %d", v.name, n, c, pos, got, want)
					}
					if got, want := indexEscapeNonASCII(b), indexEscapeNonASCIIScalar(b); got != want {
						t.Fatalf("%s: len %d %q@%d: indexEscapeNonASCII = %d, want %d", v.name, n, c, pos, got, want)
					}
					if got, want := indexStructural(b), indexStructuralScalar(b); got != want {
						t.Fatalf("%s: len %d %q@%d: indexStructural = %d, want %d", v.name, n, c, pos, got, want)
					}
				}
			}
		}
	}
}
