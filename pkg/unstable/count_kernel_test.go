package unstable

import (
	"bytes"
	"math/rand"
	"strings"
	"testing"
)

// countBody is one body of countKernel a host can run: set selects it and
// returns what restores the host's own.
type countBody struct {
	name string
	set  func() (restore func())
}

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

// TestCountBeforeCloseBodies holds every body of countBeforeClose this host
// has (countBodies: amd64's AVX2 one and its byte loop, arm64's NEON one, the
// runtime-scan fallback elsewhere) to the IndexByte+Count reference at every
// start and every ']' position of every length to 200, with the counted byte
// sprinkled before, at and after the close, and on random buffers. The vector
// bodies' boundaries are their step (32 bytes on amd64, 64 on arm64), the
// overlapping tail's shift (a function of where the cursor sits relative to
// the buffer's end) and the mask that keeps only the counted lanes below the
// close; arm64's per-lane counters add a flush every 63 steps, which the long
// buffers at the end cross.
func TestCountBeforeCloseBodies(t *testing.T) {
	for _, body := range countBodies() {
		restore := body.set()
		avx := body.name
		for _, c := range []byte{',', '{'} {
			for n := 0; n <= 200; n++ {
				base := make([]byte, n)
				for k := range base {
					base[k] = "0123456789, {"[k%13]
				}
				for i := 0; i <= n; i++ {
					if gr, gn := countBeforeClose(base, i, c); gr != -1 || gn != 0 {
						t.Fatalf("%s c=%q len %d start %d no close: got (%d,%d)", avx, c, n, i, gr, gn)
					}
					for pos := i; pos < n; pos++ {
						b := append([]byte(nil), base...)
						b[pos] = ']'
						wr, wn := refCountBeforeClose(b, i, c)
						if gr, gn := countBeforeClose(b, i, c); gr != wr || gn != wn {
							t.Fatalf("%s c=%q len %d start %d ']'@%d: got (%d,%d) want (%d,%d)", avx, c, n, i, pos, gr, gn, wr, wn)
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
								t.Fatalf("%s hint fill %#x len %d start %d ']'@%d: got (%d,%d) want (%d,%d)", avx, fill, n, i, pos, gr, gn, wr, wn)
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
				t.Fatalf("%s random len %d start %d c=%q: got (%d,%d) want (%d,%d)", avx, n, i, c, gr, gn, wr, wn)
			}
			wr, wn = refScalarsHint(b, i)
			if gr, gn := countKernel(b, i, ',', true); gr != wr || gn != wn {
				t.Fatalf("%s random hint len %d start %d: got (%d,%d) want (%d,%d)", avx, n, i, gr, gn, wr, wn)
			}
		}
		// Long spans: every flush boundary of the per-lane counters (63 steps
		// of 64 bytes) and a ']' at every offset around the last blocks.
		for _, n := range []int{4031, 4032, 4033, 4096, 8064, 8065, 20000} {
			b := []byte(strings.Repeat("1,", n/2+1))[:n]
			for _, pos := range []int{n - 1, n - 2, n - 63, n - 64, n - 65, n - 100, 4031, 4032, 4033} {
				if pos < 0 || pos >= n {
					continue
				}
				bb := append([]byte(nil), b...)
				bb[pos] = ']'
				for _, i := range []int{0, 1, 63, 64, 65} {
					wr, wn := refCountBeforeClose(bb, i, ',')
					if gr, gn := countBeforeClose(bb, i, ','); gr != wr || gn != wn {
						t.Fatalf("%s long len %d ']'@%d start %d: got (%d,%d) want (%d,%d)", avx, n, pos, i, gr, gn, wr, wn)
					}
				}
			}
		}
		restore()
	}
}
