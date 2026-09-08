package unstable

import (
	"math/bits"
	"math/rand"
	"testing"
)

// firstStructural is the definition structuralMask has to answer to: the byte
// loop indexStructuralScalar walks, over one eight-byte word.
func firstStructural(b [8]byte) int {
	for i, c := range b {
		switch c {
		case '{', '}', '[', ']', '"':
			return i
		}
	}
	return 8
}

func wordOf(b [8]byte) uint64 {
	var w uint64
	for i, c := range b {
		w |= uint64(c) << (8 * uint(i))
	}
	return w
}

// TestStructuralMaskMatchesByteScan pins the one property every caller uses:
// the LOWEST flagged lane is the first structural byte, or there is none. A
// higher lane may be flagged by a borrow out of a matching lane, so the mask as
// a whole is not compared — only TrailingZeros64 of it, which is what
// indexStructuralAt takes.
//
// The exhaustive part walks every byte value through every lane against a
// filler chosen to be adjacent to a match in exactly the ways a borrow can
// exploit: 0x00 and 0x01 (the lanes a borrow chain runs through), the four
// bracket-adjacent letters the cheap-but-approximate single-cube test would
// have claimed, and the bytes one away from each target.
func TestStructuralMaskMatchesByteScan(t *testing.T) {
	fillers := []byte{0x00, 0x01, 'a', ' ', 0x7f, 'Y', '_', 'y', 0xff, ':', ',', '0', '9', 0x21, 0x23, 0x5a, 0x5c, 0x5e, 0x7a, 0x7c, 0x7e}
	for _, f := range fillers {
		for lane := 0; lane < 8; lane++ {
			for v := 0; v < 256; v++ {
				var b [8]byte
				for i := range b {
					b[i] = f
				}
				b[lane] = byte(v)
				checkMask(t, b)
			}
		}
	}
	// Two matches in one word, at every pair of lanes, for every pair of
	// targets: the lower one must win however the borrows fall.
	targets := []byte{'"', '[', ']', '{', '}'}
	for _, t1 := range targets {
		for _, t2 := range targets {
			for i := 0; i < 8; i++ {
				for j := 0; j < 8; j++ {
					var b [8]byte
					for k := range b {
						b[k] = 1
					}
					b[i] = t1
					b[j] = t2
					checkMask(t, b)
				}
			}
		}
	}
	r := rand.New(rand.NewSource(7))
	for n := 0; n < 2_000_000; n++ {
		var b [8]byte
		for i := range b {
			b[i] = byte(r.Intn(256))
			if r.Intn(4) == 0 {
				b[i] = targets[r.Intn(len(targets))]
			}
		}
		checkMask(t, b)
	}
}

func checkMask(t *testing.T, b [8]byte) {
	t.Helper()
	want := firstStructural(b)
	m := structuralMask(wordOf(b))
	got := 8
	if m != 0 {
		got = bits.TrailingZeros64(m) >> 3
	}
	if got != want {
		t.Fatalf("structuralMask(%x): first structural lane %d, want %d (mask %016x)", b, got, want, m)
	}
}

// TestIndexStructuralAtMatchesScalar drives the real entry point across the
// prescan's word boundaries and past them into the assembly, at every start
// offset, against the byte loop it replaced.
func TestIndexStructuralAtMatchesScalar(t *testing.T) {
	r := rand.New(rand.NewSource(11))
	fill := []byte("0123456789 abcdefghij:,.eE+-_yY\x7f\xff")
	for trial := 0; trial < 4000; trial++ {
		n := 1 + r.Intn(140)
		buf := make([]byte, n)
		for i := range buf {
			buf[i] = fill[r.Intn(len(fill))]
		}
		// Sprinkle structural bytes so every distance from 0 to well past the
		// prescan is exercised.
		for k := 0; k < r.Intn(4); k++ {
			buf[r.Intn(n)] = []byte{'"', '[', ']', '{', '}'}[r.Intn(5)]
		}
		for i := 0; i <= n; i++ {
			want := i + indexStructuralScalar(buf[i:])
			if got := indexStructuralAt(buf, i); got != want {
				t.Fatalf("indexStructuralAt(%q, %d) = %d, want %d", buf, i, got, want)
			}
		}
	}
}
