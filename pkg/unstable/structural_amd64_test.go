//go:build amd64

package unstable

import (
	"math/rand"
	"testing"
)

// TestIndexStructuralBodies holds both bodies of indexStructuralAVX2 — the AVX2
// loop and the AVX-512 VBMI one, selected by useStructural512 — to the scalar
// oracle directly, at every scan start. The live-flag tests reach only the
// host's preferred body, and only through indexStructuralAt's prescan, which
// means the assembly is never entered below 32 bytes or at a start the prescan
// did not choose. The axes are the ones each body has boundaries on: the byte
// value (the VBMI table must match exactly five of 256, the AVX2 fold exactly
// five), the position against the 32- and 64-byte steps, the overlapping AVX2
// tail's shift, the VBMI masked tail's length, and the start offset.
func TestIndexStructuralBodies(t *testing.T) {
	saved := useStructural512
	defer func() { useStructural512 = saved }()
	structural := func(c byte) bool {
		return c == '{' || c == '}' || c == '[' || c == ']' || c == '"'
	}
	for _, vbmi := range []bool{false, true} {
		if vbmi && !saved {
			continue // host lacks AVX-512 VBMI
		}
		if !useAVX2 {
			t.Skip("no AVX2")
		}
		useStructural512 = vbmi
		name := map[bool]string{false: "avx2", true: "vbmi"}[vbmi]

		// Every byte value at every position of a 131-byte buffer (two 64-byte
		// steps and a 3-byte tail; the AVX2 32-byte step and overlap both run).
		for v := 0; v < 256; v++ {
			for pos := 0; pos < 131; pos++ {
				b := make([]byte, 131)
				for k := range b {
					b[k] = 'a' + byte(k%26)
				}
				b[pos] = byte(v)
				want := len(b)
				if structural(byte(v)) {
					want = pos
				}
				if got := indexStructuralAVX2(b, 0); got != want {
					t.Fatalf("%s: byte %#x at %d: got %d, want %d", name, v, pos, got, want)
				}
			}
		}

		// Every length, every start, one structural byte at every position at or
		// after the start (and none, and one before the start that must be ignored).
		for n := 0; n <= 330; n++ {
			base := make([]byte, n)
			for k := range base {
				base[k] = '0' + byte(k%10)
			}
			for i := 0; i <= n; i++ {
				if got := indexStructuralAVX2(base, i); got != n {
					t.Fatalf("%s: len %d start %d clean: got %d, want %d", name, n, i, got, n)
				}
				if i > 0 {
					b := append([]byte(nil), base...)
					b[i-1] = '"'
					if got := indexStructuralAVX2(b, i); got != n {
						t.Fatalf("%s: len %d start %d, match before start: got %d, want %d", name, n, i, got, n)
					}
				}
				for pos := i; pos < n; pos++ {
					b := append([]byte(nil), base...)
					b[pos] = "{}[]\""[pos%5]
					if got := indexStructuralAVX2(b, i); got != pos {
						t.Fatalf("%s: len %d start %d, %q at %d: got %d", name, n, i, b[pos], pos, got)
					}
				}
			}
		}

		rng := rand.New(rand.NewSource(64))
		for iter := 0; iter < 20000; iter++ {
			n := rng.Intn(600)
			b := make([]byte, n)
			for k := range b {
				if rng.Intn(40) == 0 {
					b[k] = "{}[]\""[rng.Intn(5)]
				} else {
					b[k] = byte(rng.Intn(256))
					if structural(b[k]) {
						b[k] = 'x'
					}
				}
			}
			i := 0
			if n > 0 {
				i = rng.Intn(n + 1)
			}
			if got, want := indexStructuralAVX2(b, i), i+indexStructuralScalar(b[i:]); got != want {
				t.Fatalf("%s: random len %d start %d: got %d, want %d", name, n, i, got, want)
			}
		}
	}
}
