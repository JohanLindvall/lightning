//go:build amd64

package unstable

import (
	"bytes"
	"math/rand"
	"testing"
)

// refCountBeforeClose is what the presize counters computed before
// countBeforeClose: the first ']' with IndexByte, then Count over the span.
func refCountBeforeClose(data []byte, i int, c byte) (int, int) {
	if uint(i) >= uint(len(data)) {
		return -1, 0
	}
	rb := bytes.IndexByte(data[i:], ']')
	if rb < 0 {
		return -1, 0
	}
	return rb, bytes.Count(data[i:i+rb], []byte{c})
}

// refScalarsHint is CountArrayScalars' element count from the reference scan:
// commas plus one clamped to (rb+1)/2, or for a comma-free span 1 unless every
// byte is <= 0x20.
func refScalarsHint(data []byte, i int) (int, int) {
	rb, commas := refCountBeforeClose(data, i, ',')
	if rb < 0 {
		return -1, 0
	}
	if commas > 0 {
		n := commas + 1
		if lim := (rb + 1) / 2; n > lim {
			n = lim
		}
		return rb, n
	}
	for _, c := range data[i : i+rb] {
		if c > ' ' {
			return rb, 1
		}
	}
	return rb, 0
}

// TestCountBeforeCloseBodies holds both bodies of countBeforeClose — the AVX2
// one and the byte loop — to the IndexByte+Count reference at every start and
// every ']' position of every length to 200, with the counted byte sprinkled
// before, at and after the close, and on random buffers. The AVX2 body's
// boundaries are the 32-byte step, the overlapping tail's shift (a function of
// where the cursor sits relative to the buffer's end) and the BZHI that keeps
// only the counted lanes below the close.
func TestCountBeforeCloseBodies(t *testing.T) {
	saved := useCountAVX2
	defer func() { useCountAVX2 = saved }()
	for _, avx := range []bool{false, true} {
		if avx && !saved {
			continue
		}
		useCountAVX2 = avx
		for _, c := range []byte{',', '{'} {
			for n := 0; n <= 200; n++ {
				base := make([]byte, n)
				for k := range base {
					base[k] = "0123456789, {"[k%13]
				}
				for i := 0; i <= n; i++ {
					if gr, gn := countBeforeClose(base, i, c); gr != -1 || gn != 0 {
						t.Fatalf("avx=%v c=%q len %d start %d no close: got (%d,%d)", avx, c, n, i, gr, gn)
					}
					for pos := i; pos < n; pos++ {
						b := append([]byte(nil), base...)
						b[pos] = ']'
						wr, wn := refCountBeforeClose(b, i, c)
						if gr, gn := countBeforeClose(b, i, c); gr != wr || gn != wn {
							t.Fatalf("avx=%v c=%q len %d start %d ']'@%d: got (%d,%d) want (%d,%d)", avx, c, n, i, pos, gr, gn, wr, wn)
						}
						// The hint mode over the same buffer, and over a blank
						// span (all whitespace, then all control bytes) where only
						// the byte walk decides.
						for _, fill := range []byte{0, ' ', 0x01} {
							bb := b
							if fill != 0 {
								bb = append([]byte(nil), b...)
								for k := i; k < pos; k++ {
									bb[k] = fill
								}
							}
							wr, wn := refScalarsHint(bb, i)
							if gr, gn := countKernel(bb, i, ',', true); gr != wr || gn != wn {
								t.Fatalf("avx=%v hint fill %#x len %d start %d ']'@%d: got (%d,%d) want (%d,%d)", avx, fill, n, i, pos, gr, gn, wr, wn)
							}
						}
					}
				}
			}
		}
		r := rand.New(rand.NewSource(3))
		for it := 0; it < 50000; it++ {
			n := r.Intn(700)
			b := make([]byte, n)
			for k := range b {
				switch r.Intn(10) {
				case 0:
					b[k] = ','
				case 1:
					b[k] = '{'
				default:
					b[k] = byte(r.Intn(256))
				}
				if b[k] == ']' && r.Intn(50) != 0 {
					b[k] = 'x'
				}
			}
			i := 0
			if n > 0 {
				i = r.Intn(n + 1)
			}
			c := []byte{',', '{'}[r.Intn(2)]
			wr, wn := refCountBeforeClose(b, i, c)
			if gr, gn := countBeforeClose(b, i, c); gr != wr || gn != wn {
				t.Fatalf("avx=%v random len %d start %d c=%q: got (%d,%d) want (%d,%d)", avx, n, i, c, gr, gn, wr, wn)
			}
			wr, wn = refScalarsHint(b, i)
			if gr, gn := countKernel(b, i, ',', true); gr != wr || gn != wn {
				t.Fatalf("avx=%v random hint len %d start %d: got (%d,%d) want (%d,%d)", avx, n, i, gr, gn, wr, wn)
			}
		}
	}
}
