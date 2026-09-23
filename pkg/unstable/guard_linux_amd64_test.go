//go:build linux && amd64

package unstable

import (
	"strings"
	"testing"

	"golang.org/x/sys/unix"
)

// guardedPage maps three pages with the outer two inaccessible and returns the
// middle one: a slice placed flush against either end of it faults on any read
// past that end, which is how the tests below prove that no assembly body reads
// outside its buffer. The masked tail loads rely on fault suppression for the
// lanes they switch off, and the overlapping tails on reading only inside the
// slice; both are claims about memory that a normal Go allocation, always
// followed by more heap, can never falsify.
func guardedPage(t *testing.T) []byte {
	t.Helper()
	ps := unix.Getpagesize()
	mem, err := unix.Mmap(-1, 0, 3*ps, unix.PROT_NONE, unix.MAP_PRIVATE|unix.MAP_ANON)
	if err != nil {
		t.Skipf("mmap: %v", err)
	}
	t.Cleanup(func() { _ = unix.Munmap(mem) })
	mid := mem[ps : 2*ps : 2*ps]
	if err := unix.Mprotect(mid, unix.PROT_READ|unix.PROT_WRITE); err != nil {
		t.Skipf("mprotect: %v", err)
	}
	return mid
}

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
// both float bodies, the points walk and the validation walk — over arrays flush against a guard
// page at their end (unterminated, so each walk runs until its window check
// stops it) and at their start, from every start position within reach of the
// end. The kernels read whole windows of 64 to 96 bytes and gather 32 bytes at
// any lane of one, so each is a claim that the window check covers the
// farthest byte its body touches.
func TestNumberKernelsStayInBounds(t *testing.T) {
	page := guardedPage(t)
	savedV := useFloatRunVBMI
	defer func() { useFloatRunVBMI = savedV }()
	patterns := []string{
		"1234,-5,67890123,0,",
		"1.25,-3.5, 17.123456789012345,-0.000123 ,",
		"[-65.613616999999977,43.420273000000009],[1.5, 2],",
		"[1.5,-2.25]" + strings.Repeat(" ", 70) + ",", // a separator windows away
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
					if useValidRun {
						for _, b512 := range validRunBodies() {
							useValidRun512 = b512
							_, _ = validNumberRun(b, i)
						}
						useValidRun512 = validRun512Host
					}
					if floatRunHost {
						useFloatRunVBMI = false
						_, _, _ = parseFloatRunAVX2(b, i, floats)
						useFloatRunVBMI = savedV
					}
					if savedV {
						_, _, _ = parseFloatRunVBMI(b, i, floats)
						_, _, _ = parseFloatPoints(b, i, floats, 2)
						_, _, _ = parseFloatPoints(b, i, floats, 3)
					}
				}
			}
		}
	}
}
