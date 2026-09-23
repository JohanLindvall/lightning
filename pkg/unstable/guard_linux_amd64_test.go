//go:build linux && amd64

package unstable

import (
	"strings"
	"testing"
)

// TestAssemblyStaysInBounds runs every amd64 scanner body over buffers flush
// against a guard page at their end and at their start, at every length up to
// 300, clean (so each scan runs to the end) and with a match in the last byte.
func TestAssemblyStaysInBounds(t *testing.T) {
	page := guardedPage(t)
	savedS, savedA := useStructural512, useAVX2
	savedB, savedB512, savedC := useSkipBlocks, useSkipBlocks512, useCountAVX2
	defer func() {
		useStructural512, useAVX2 = savedS, savedA
		useSkipBlocks, useSkipBlocks512, useCountAVX2 = savedB, savedB512, savedC
	}()

	type body struct {
		name       string
		avx2, vbmi bool
	}
	bodies := []body{{"sse2", false, false}}
	if savedA {
		bodies = append(bodies, body{"avx2", true, false})
	}
	if savedS {
		bodies = append(bodies, body{"vbmi", true, true})
	}
	for _, bd := range bodies {
		useAVX2, useStructural512 = bd.avx2, bd.vbmi
		for n := 0; n <= 300; n++ {
			for _, atEnd := range []bool{true, false} {
				var b []byte
				if atEnd {
					b = page[len(page)-n:]
				} else {
					b = page[:n:n]
				}
				for _, last := range []byte{'a', '"', '}', '\\', 0x01} {
					for k := range b {
						b[k] = 'a'
					}
					if n > 0 {
						b[n-1] = last
					}
					_ = indexQuoteOrBackslashSSE2(b, 0)
					_ = indexEscapeSSE2(b)
					_ = indexEscapeNonASCIISSE2(b)
					if bd.avx2 {
						for _, i := range []int{0, n / 2, n} {
							_ = indexStructuralAVX2(b, i)
						}
					}
					if savedA {
						_ = indexStructural(b)
					}
				}
				if bd.avx2 && savedC {
					useCountAVX2 = true
					for _, i := range []int{0, n / 2, n} {
						_, _ = countBeforeClose(b, i, ',')
					}
					useCountAVX2 = savedC
				}
				// The container skip over an unterminated container, so every
				// body scans to the end: the AVX2 tail reads the last 64 bytes,
				// the AVX-512 one a masked load, and neither may read past the
				// end or (for the overlapping block) before the buffer's start.
				if savedB && n >= 2 {
					for k := range b {
						b[k] = 'a'
					}
					b[0] = '['
					for _, v := range []struct{ blocks, b512 bool }{{true, false}, {true, true}} {
						if v.b512 && !savedB512 {
							continue
						}
						useSkipBlocks, useSkipBlocks512 = v.blocks, v.b512
						for _, i := range []int{0, n / 2, n - 1} {
							b[i] = '['
							if _, err := skipContainerFast(b, i, '['); err == nil {
								t.Fatalf("n=%d i=%d: unterminated container accepted", n, i)
							}
						}
					}
					useSkipBlocks, useSkipBlocks512 = savedB, savedB512
				}
			}
		}
	}
}

// TestNumberKernelsStayInBounds runs the integer- and float-array kernels —
// every float body, both points walks and both validation walks, flat and
// ring — over arrays flush against a guard page at their end (unterminated,
// so each walk runs until its window check stops it) and at their start, from
// every start position within reach of the end. The kernels read whole windows
// of 64 to 96 bytes and load sixteen or 32 bytes at any lane of one, so each
// is a claim that the window check covers the farthest byte its body touches.
func TestNumberKernelsStayInBounds(t *testing.T) {
	page := guardedPage(t)
	defer restoreKernels()
	patterns := []string{
		"1234,-5,67890123,0,",
		"1.25,-3.5, 17.123456789012345,-0.000123 ,",
		"-65.613616999999977,43.420273000000009,1234567890.123456789,",
		"[-65.613616999999977,43.420273000000009],[1.5, 2],",
		"[1.5,-2.25]" + strings.Repeat(" ", 70) + ",", // a separator windows away
		// A number late in its window behind whitespace, and one that closes
		// the array: the elements the 80-byte bound exists for — a sixteen-byte
		// load at a first digit near lane 63 — and a long number whose loads
		// end at its delimiter near lane 63.
		"1.5," + strings.Repeat(" ", 45) + "-2.25," + strings.Repeat("\n", 57) + "3]",
		strings.Repeat(" ", 58) + "12345678.9012345,",
		strings.Repeat(" ", 44) + "-12345678.90123456789,",
		"[" + strings.Repeat(" ", 40) + "-65.613616999999977, 43.420273000000009]" + strings.Repeat(" ", 20) + ",",
		// The same for a point: a short number starting late in the point's
		// window, whose sixteen-byte load is what needs 80 bytes.
		"[" + strings.Repeat(" ", 50) + "1.5,2],",
		"[1.5," + strings.Repeat(" ", 48) + "-2.25],",
	}
	ints := make([]int64, 1024)
	floats := make([]float64, 1024)
	for _, pat := range patterns {
		for n := 0; n <= 300; n++ {
			for _, atEnd := range []bool{true, false} {
				var b []byte
				if atEnd {
					b = page[len(page)-n:]
				} else {
					b = page[:n:n]
				}
				for k := range b {
					b[k] = pat[k%len(pat)]
				}
				if n > 0 {
					b[0] = '['
				}
				for i := max(0, n-140); i <= n; i++ {
					if useIntRun {
						_, _, _ = parseIntRunAVX2(b, i, ints)
					}
					for _, body := range validRunBodies() {
						body.use()
						_, _ = validNumberRun(b, i)
						_, _ = validPointsRun(b, i)
					}
					for _, body := range floatRunBodies() {
						body.use()
						_, _, _ = parseFloatRun(b, i, floats)
						_, _, _ = parseFloatPoints(b, i, floats, 2)
						_, _, _ = parseFloatPoints(b, i, floats, 3)
					}
					restoreKernels()
				}
			}
		}
	}
}
