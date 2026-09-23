//go:build arm64

package unstable

import (
	"math/rand"
	"testing"
)

// TestIndexStructuralBodies holds both arm64 bodies of the structural scan —
// the SVE2 one and the NEON one — to the scalar oracle directly, at every scan
// start, as the amd64 test of the same name does for its two. The live-flag
// tests reach only the host's body, and only through indexStructuralAt's
// prescan, which enters the assembly at a start the prescan chose. The axes are
// the ones each body has boundaries on: the byte value (MATCH's set and the
// NEON nibble tables must each select exactly five of 256), the position
// against the 16-byte block and the SVE2 two- and four-vector steps, the
// ragged end (WHILELO's predicate, the NEON scalar tail), and the start.
func TestIndexStructuralBodies(t *testing.T) {
	structural := func(c byte) bool {
		return c == '{' || c == '}' || c == '[' || c == ']' || c == '"'
	}
	bodies := []struct {
		name string
		f    func([]byte, int) int
	}{{"neon", indexStructuralNEON}}
	if useSVE2 {
		bodies = append(bodies, struct {
			name string
			f    func([]byte, int) int
		}{"sve2", indexStructuralSVE2})
	}
	for _, body := range bodies {
		// Every byte value at every position of a 131-byte buffer.
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
				if got := body.f(b, 0); got != want {
					t.Fatalf("%s: byte %#x at %d: got %d, want %d", body.name, v, pos, got, want)
				}
			}
		}

		// Every length, every start, one structural byte at every position at
		// or after the start (and none, and one before the start that must be
		// ignored).
		for n := 0; n <= 330; n++ {
			base := make([]byte, n)
			for k := range base {
				base[k] = '0' + byte(k%10)
			}
			for i := 0; i <= n; i++ {
				if got := body.f(base, i); got != n {
					t.Fatalf("%s: len %d start %d clean: got %d, want %d", body.name, n, i, got, n)
				}
				if i > 0 {
					b := append([]byte(nil), base...)
					b[i-1] = '"'
					if got := body.f(b, i); got != n {
						t.Fatalf("%s: len %d start %d, match before start: got %d, want %d", body.name, n, i, got, n)
					}
				}
				for pos := i; pos < n; pos++ {
					b := append([]byte(nil), base...)
					b[pos] = "{}[]\""[pos%5]
					if got := body.f(b, i); got != pos {
						t.Fatalf("%s: len %d start %d, %q at %d: got %d", body.name, n, i, b[pos], pos, got)
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
			if got, want := body.f(b, i), i+indexStructuralScalar(b[i:]); got != want {
				t.Fatalf("%s: random len %d start %d: got %d, want %d", body.name, n, i, got, want)
			}
		}
	}
}
