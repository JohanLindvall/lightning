package unstable

import (
	"strings"
	"testing"
)

// --- internal scanners: scalar implementations -----------------------------

func TestIndexCloseOrEscapeScalar(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{``, 0},
		{`abc`, 3},
		{`ab"c`, 2},
		{`ab\c`, 2},
		{`a\"b`, 1}, // backslash before quote
		{`"`, 0},    // quote first
		{`\`, 0},    // backslash first
		{`abc"def`, 3},
		{strings.Repeat("x", 50) + `"`, 50}, // long, quote at end
		{strings.Repeat("x", 50) + `\`, 50}, // long, backslash at end
	}
	for _, tt := range tests {
		if got := indexCloseOrEscapeScalar([]byte(tt.in)); got != tt.want {
			t.Errorf("indexCloseOrEscapeScalar(%q) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestIndexStructuralScalar(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{``, 0},
		{`abc`, 3},
		{`ab{c`, 2},
		{`ab}c`, 2},
		{`ab[c`, 2},
		{`ab]c`, 2},
		{`ab"c`, 2},
		{`123:456`, 7}, // none present
		{strings.Repeat("x", 40) + `{`, 40},
	}
	for _, tt := range tests {
		if got := indexStructuralScalar([]byte(tt.in)); got != tt.want {
			t.Errorf("indexStructuralScalar(%q) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

// indexCloseOrEscape / indexStructural dispatch to SIMD on long inputs; verify
// they agree with the scalar reference across a range of lengths and positions.
func TestIndexFunctionsMatchScalar(t *testing.T) {
	// 320 bytes: long enough that the scanners' widest tail loop (the AVX-512
	// 64-byte kmask loop on amd64) runs several full iterations after the 32-byte
	// SSE2 head, with lengths straddling every block boundary (96 = head + one
	// 64-byte block exactly; 97/160/193 leave 1..32-byte leftovers for the
	// narrower loops; 256/320 are multi-block).
	base := strings.Repeat("abcdefgh", 40)
	for _, n := range []int{0, 1, 15, 16, 31, 32, 33, 64, 96, 97, 100, 128, 160, 192, 193, 256, 320} {
		s := base[:n]
		if got, want := indexCloseOrEscape([]byte(s)), indexCloseOrEscapeScalar([]byte(s)); got != want {
			t.Errorf("indexCloseOrEscape(len=%d) = %d, want %d", n, got, want)
		}
		if got, want := indexStructural([]byte(s)), indexStructuralScalar([]byte(s)); got != want {
			t.Errorf("indexStructural(len=%d) = %d, want %d", n, got, want)
		}
		if got, want := indexEscape([]byte(s)), indexEscapeScalar([]byte(s)); got != want {
			t.Errorf("indexEscape(len=%d) = %d, want %d", n, got, want)
		}
		// Insert a target byte at each position and confirm both paths find it.
		// The control bytes (0x00, '\n', 0x1f) exercise indexEscape's < 0x20 test;
		// 0x20 (space) sits just past that boundary and must NOT match, pinning
		// the exact <-vs-<= edge in every vector variant.
		for _, c := range []byte{'"', '\\', '{', '}', '[', ']', 0x00, '\n', 0x1f, ' '} {
			for pos := 0; pos < n; pos++ {
				b := []byte(s)
				b[pos] = c
				if got, want := indexCloseOrEscape(b), indexCloseOrEscapeScalar(b); got != want {
					t.Fatalf("indexCloseOrEscape(len=%d, %q@%d) = %d, want %d", n, c, pos, got, want)
				}
				if got, want := indexStructural(b), indexStructuralScalar(b); got != want {
					t.Fatalf("indexStructural(len=%d, %q@%d) = %d, want %d", n, c, pos, got, want)
				}
				if got, want := indexEscape(b), indexEscapeScalar(b); got != want {
					t.Fatalf("indexEscape(len=%d, %q@%d) = %d, want %d", n, c, pos, got, want)
				}
			}
		}
	}
}
