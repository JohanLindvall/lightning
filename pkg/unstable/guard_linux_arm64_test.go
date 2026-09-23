//go:build linux && arm64

package unstable

import (
	"strings"
	"testing"
)

// TestAssemblyStaysInBounds runs every arm64 scanner body — NEON and, where
// the core has it, SVE2 — over buffers flush against a guard page at their end
// and at their start, at every length up to 300, clean (so each scan runs to
// the end) and with a match in the last byte; and the presize count and the
// container skip, whose final < 64 bytes are the buffer's LAST 64 read as an
// overlapping block, over buffers with no ']' or no close at all.
func TestAssemblyStaysInBounds(t *testing.T) {
	page := guardedPage(t)
	savedSVE, savedB := useSVE2, useSkipBlocks
	defer func() { useSVE2, useSkipBlocks = savedSVE, savedB }()
	for _, sve := range []bool{false, true} {
		if sve && !savedSVE {
			continue
		}
		useSVE2 = sve
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
					for _, i := range []int{0, n / 2, n} {
						_ = indexCloseOrEscapeAt(b, i)
						_ = indexStructuralAt(b, i)
						_ = indexStructuralNEON(b, i)
						if sve {
							_ = indexStructuralSVE2(b, i)
						}
					}
					_ = indexEscape(b)
					_ = indexEscapeNonASCII(b)
				}
				for k := range b {
					b[k] = "1, "[k%3]
				}
				for _, i := range []int{0, n / 2, n} {
					_, _ = countBeforeClose(b, i, ',')
					_, _ = countKernel(b, i, ',', true)
				}
				// The container skip over an unterminated container, so the
				// block loop scans to the end and takes the tail block.
				if n >= 2 {
					for k := range b {
						b[k] = 'a'
					}
					for _, on := range []bool{false, true} {
						useSkipBlocks = on
						for _, i := range []int{0, n / 2, n - 1} {
							b[i] = '['
							if _, err := skipContainerFast(b, i, '['); err == nil {
								t.Fatalf("n=%d i=%d: unterminated container accepted", n, i)
							}
						}
					}
					useSkipBlocks = savedB
				}
			}
		}
	}
}

// TestNumberKernelsStayInBounds runs the integer- and float-array kernels —
// the flat walk, the points walk and both validation walks — over arrays
// flush against a guard page at their end (unterminated, so each walk runs
// until its block check stops it) and at their start, from every start
// position within reach of the end. The float walk loads sixteen bytes at a
// number's first digit anywhere in a 64-byte block and sixteen ending at its
// last byte, so each is a claim that the 80-byte block check covers the
// farthest byte a conversion touches (and that the second load never reaches
// before the block).
func TestNumberKernelsStayInBounds(t *testing.T) {
	page := guardedPage(t)
	patterns := []string{
		"1234,-5,67890123,0,",
		"1.25,-3.5, 17.123456789012345,-0.000123 ,",
		"-65.613616999999977,43.420273000000009,1234567890.123456789,",
		"[-65.613616999999977,43.420273000000009],[1.5, 2],",
		"[1.5,-2.25]" + strings.Repeat(" ", 70) + ",", // a separator windows away
		// A number late in its block behind whitespace, and one that closes
		// the array: the elements the 80-byte bound exists for, a gather at a
		// first digit near lane 63.
		"1.5," + strings.Repeat(" ", 45) + "-2.25," + strings.Repeat("\n", 57) + "3]",
		strings.Repeat(" ", 58) + "12345678.9012345,",
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
						_, _, _ = parseIntRunNEON(b, i, ints)
					}
					if useFloatRun {
						_, _, _ = parseFloatRunNEON(b, i, floats)
						_, _, _ = parseFloatPointsNEON(b, i, floats, 2)
						_, _, _ = parseFloatPointsNEON(b, i, floats, 3)
						_, _ = validNumberRunNEON(b, i)
						_, _ = validPointsRunNEON(b, i)
					}
				}
			}
		}
	}
}
