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
		if got, want := indexEscapeNonASCII([]byte(s)), indexEscapeNonASCIIScalar([]byte(s)); got != want {
			t.Errorf("indexEscapeNonASCII(len=%d) = %d, want %d", n, got, want)
		}
		// Insert a target byte at each position and confirm both paths find it.
		// The control bytes (0x00, '\n', 0x1f) exercise indexEscape's < 0x20 test;
		// 0x20 (space) sits just past that boundary and must NOT match, pinning
		// the exact <-vs-<= edge in every vector variant. 0x7f/0x80/0xff pin
		// indexEscapeNonASCII's ASCII edge the same way — 0x7f must NOT match, the
		// smallest and largest high-bit bytes must — while also confirming high
		// bytes stay invisible to the three plain scanners.
		for _, c := range []byte{'"', '\\', '{', '}', '[', ']', 0x00, '\n', 0x1f, ' ', 0x7f, 0x80, 0xff} {
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
				if got, want := indexEscapeNonASCII(b), indexEscapeNonASCIIScalar(b); got != want {
					t.Fatalf("indexEscapeNonASCII(len=%d, %q@%d) = %d, want %d", n, c, pos, got, want)
				}
			}
		}
	}
}

// TestIndexEscapeNonASCIIScalarOracle pins the SWAR scalar (the reference the
// SIMD variants are tested against) to a naive byte loop over every byte value,
// alone and at every lane offset of an 8-byte word.
func TestIndexEscapeNonASCIIScalarOracle(t *testing.T) {
	naive := func(b []byte) int {
		for i, c := range b {
			if c < 0x20 || c >= 0x80 || c == '"' || c == '\\' {
				return i
			}
		}
		return len(b)
	}
	for v := 0; v < 256; v++ {
		b := []byte{byte(v)}
		if got, want := indexEscapeNonASCIIScalar(b), naive(b); got != want {
			t.Fatalf("indexEscapeNonASCIIScalar(%#x) = %d, want %d", v, got, want)
		}
		for pos := 0; pos < 16; pos++ {
			buf := []byte("abcdefghABCDEFGH")
			buf[pos] = byte(v)
			if got, want := indexEscapeNonASCIIScalar(buf), naive(buf); got != want {
				t.Fatalf("indexEscapeNonASCIIScalar(%#x@%d) = %d, want %d", v, pos, got, want)
			}
		}
	}
}

// TestIndexCloseOrEscapeAtOffsets sweeps the scan START, which is the axis the
// offset form of the scanner added and the axis no other test covers: the
// generated decoders and every Read* here call indexCloseOrEscapeAt(data, i)
// with i deep inside the document, so the assembly's peeled first block and its
// position recovery both work from a nonzero index. A body that forgot to add
// the offset back — or added it twice, or peeled from the base pointer instead
// of from base+i — still passes every len-0-start test.
//
// The oracle is the scalar scanner over the equivalent subslice, and every
// length/start/match-position combination up to two full blocks past the widest
// vector body is tried, so the peeled block, the two-vector step, the unrolled
// loop and the predicated remainder are each entered at every alignment.
func TestIndexCloseOrEscapeAtOffsets(t *testing.T) {
	for n := 0; n <= 160; n++ {
		base := make([]byte, n)
		for i := range base {
			base[i] = byte('a' + i%26)
		}
		for _, c := range []byte{'"', '\\'} {
			for pos := -1; pos < n; pos++ { // -1 = no match anywhere
				b := append([]byte(nil), base...)
				if pos >= 0 {
					b[pos] = c
				}
				for start := 0; start <= n; start++ {
					want := start + indexCloseOrEscapeScalar(b[start:])
					if got := indexCloseOrEscapeAt(b, start); got != want {
						t.Fatalf("indexCloseOrEscapeAt(len %d, %q@%d, start %d) = %d, want %d",
							n, c, pos, start, got, want)
					}
				}
			}
		}
	}
}

// TestIndexCloseOrEscapeAtPastEnd pins the out-of-span answer. Nothing in the
// tree passes i > len(b) today, but the two assembly bodies reach that case by
// different routes — arm64 falls into its predicated remainder, amd64 computes a
// negative byte count — and on amd64 an unguarded negative count walks off the
// end of the buffer. len(b) is the answer either way, which is also what a scan
// that found nothing returns.
func TestIndexCloseOrEscapeAtPastEnd(t *testing.T) {
	for _, n := range []int{0, 1, 8, 16, 33, 100} {
		b := make([]byte, n)
		for i := range b {
			b[i] = '"'
		}
		for _, start := range []int{n, n + 1, n + 7, n + 64} {
			if got := indexCloseOrEscapeAt(b, start); got != n {
				t.Fatalf("indexCloseOrEscapeAt(len %d, start %d) = %d, want %d", n, start, got, n)
			}
		}
	}
}
