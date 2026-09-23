//go:build amd64

package unstable

import (
	"math"

	"golang.org/x/sys/cpu"
)

// useFloatRun gates parseFloatRunAVX2, the SIMD kernel for arrays of decimal
// numbers: AVX2 for its classification and BMI2 for SHRX/SHLX.
var useFloatRun = cpu.X86.HasAVX2 && cpu.X86.HasBMI2

// useFloatRunLong says the kernel converts numbers of up to 19 digits (with
// Eisel-Lemire) — the shape of a coordinate — which is what the points walk
// and the fixed-array reader need; the shared readers read it under that name
// on every architecture. Both amd64 bodies do, so it is useFloatRun.
var useFloatRunLong = useFloatRun

// useFloatRunVBMI selects the AVX-512 VBMI body (parseFloatRunVBMI and the
// VBMI points walk), for CPUs that have it: VPERMI2B gathers a long number's
// digits out of one 32-byte load, and BW/VL supply the byte compares into mask
// registers and the ymm broadcast from a general register. parseFloatRunAVX2
// reads the flag in its first instructions and tail-jumps, so parseFloatRun
// stays a single call.
var useFloatRunVBMI = useFloatRun && cpu.X86.HasAVX512F && cpu.X86.HasAVX512BW &&
	cpu.X86.HasAVX512VL && cpu.X86.HasAVX512VBMI

// parseFloatRunAVX2 parses a run of comma-separated decimal numbers of at
// most 19 digits and no exponent from data[i:] into out, and returns how many
// it wrote, where the scalar loop should resume, and whether the run ended at
// the array's ']' (p is then that byte's index). See floatrun_amd64.s.
//
//go:noescape
func parseFloatRunAVX2(data []byte, i int, out []float64) (n, p, closed int)

// parseFloatRunVBMI is the AVX-512 VBMI body, reached only from
// parseFloatRunAVX2's tail jump; the declaration is what asmdecl checks the
// frame against.
//
//nolint:unused // called from assembly
//go:noescape
func parseFloatRunVBMI(data []byte, i int, out []float64) (n, p, closed int)

func parseFloatRun(data []byte, i int, out []float64) (n, p, closed int) {
	return parseFloatRunAVX2(data, i, out)
}

// parseFloatRunV is parseFloatRun, for a caller that has already checked
// useFloatRunLong (the fixed-array reader): arm64 spells it the same way.
func parseFloatRunV(data []byte, i int, out []float64) (n, p, closed int) {
	return parseFloatRunAVX2(data, i, out)
}

// floatRunTab is the kernel's tables, reached through one base register:
//
//   - at 16*(16*L1+L2), the PSHUFB control for a number of L1 integer and L2
//     fraction digits (L1+L2 <= 15): it right-aligns the digits of the
//     sixteen bytes loaded at the first digit, skipping the '.', with zeros in
//     front (0x80 zeroes a lane);
//   - from 4096, the divisors: 10^k at qword k and -10^k at qword 32+k, exact
//     for every k either body uses;
//   - from 4608, at 32*(32*L1+L2), the VBMI body's VPERMI2B template for a
//     number of up to 19 digits: for each of the 24 output lanes the offset,
//     from the number's first digit, of the digit that lands there
//     (right-aligned, the '.' skipped), or 32 for a leading zero, which
//     selects the '0' register;
//   - from 25088, at qword L2, the high word of the 128-bit significand of
//     10^-L2, from 25248, at qword L2, eiselLemire64's biased binary exponent
//     estimate for it less 2 (the long conversions fold eiselLemire64's
//     "-= 1 ^ msb" and the "- 1" their assembly wants into it), and from
//     25408, at qword L2, the low word, for the refinement LONGTAIL2 makes
//     where eiselLemire64 does;
//   - at 24928+32*(4*L1+L1+L2), for 16 <= L1+L2 <= 19, the AVX2 body's
//     VPSHUFB control for a long number (LONGGATHER2): its low half picks the
//     first digits out of the sixteen bytes at the first digit, its high half
//     the last eight out of the sixteen ending at the number's end, so the
//     digits land right-aligned in lanes 0-23 with zeros in front (the index
//     4*L1+L1+L2 is at least 20, so the region starts at 25568).
var floatRunTab [28000]byte

func init() {
	for l1 := 1; l1 <= 15; l1++ {
		for l2 := 0; l1+l2 <= 15; l2++ {
			c := floatRunTab[16*(16*l1+l2):][:16]
			for t := range c {
				c[t] = 0x80
			}
			l := l1 + l2
			for k := 0; k < l; k++ {
				src := k
				if k >= l1 {
					src++ // past the '.'
				}
				c[16-l+k] = byte(src)
			}
		}
	}
	p := 1.0
	for k := 0; k < 32; k++ {
		if k > 22 {
			p = 1 // never read: the kernel's L2 is at most 14
		}
		b := math.Float64bits(p)
		n := math.Float64bits(-p)
		for j := 0; j < 8; j++ {
			floatRunTab[4096+8*k+j] = byte(b >> (8 * j))
			floatRunTab[4096+8*(32+k)+j] = byte(n >> (8 * j))
		}
		p *= 10
	}
}

func init() {
	for l1 := 1; l1 <= 19; l1++ {
		for l2 := 0; l1+l2 <= 19; l2++ {
			tpl := floatRunTab[4608+32*(32*l1+l2):][:32]
			l := l1 + l2
			for t := range tpl {
				tpl[t] = 32 // a leading zero: the '0' register
				if k := t - (24 - l); t < 24 && k >= 0 {
					src := k
					if k >= l1 {
						src++ // past the '.'
					}
					tpl[t] = byte(src)
				}
			}
		}
	}
	for l2 := 0; l2 <= 19; l2++ {
		exp10 := -l2
		pow := detailedPowersOfTen[exp10-detailedPowersOfTenMinExp10]
		expBase := uint64(217706*exp10>>16+64+1023) - 2
		for j := 0; j < 8; j++ {
			floatRunTab[25088+8*l2+j] = byte(pow[1] >> (8 * j))
			floatRunTab[25088+160+8*l2+j] = byte(expBase >> (8 * j))
			floatRunTab[25088+320+8*l2+j] = byte(pow[0] >> (8 * j))
		}
	}
	for l1 := 1; l1 <= 19; l1++ {
		for l := max(16, l1); l <= 19; l++ {
			l2 := l - l1
			span := l // from the first digit to the last, the '.' included
			if l2 > 0 {
				span++
			}
			c := floatRunTab[24928+32*(4*l1+l):][:32]
			for t := range c {
				c[t] = 0x80 // a leading zero, and lanes 24-31
			}
			for t := 24 - l; t < 24; t++ {
				src := t - (24 - l) // the digit's index, then its offset
				if src >= l1 {
					src++ // past the '.'
				}
				if t < 16 {
					c[t] = byte(src) // from the sixteen bytes at the first digit
				} else {
					c[t] = byte(src - (span - 16)) // from the sixteen ending at the end
				}
			}
		}
	}
}

// useValidRun gates validNumberRun, the decimal-array walk as a check, which
// SkipValueStrict hands an array of numbers: AVX2 for the classification and
// BMI2 for SHRX (BMI1's TZCNT/BLSR come with it).
var useValidRun = cpu.X86.HasAVX2 && cpu.X86.HasBMI2

// useValidPoints gates validPointsRun, the coordinate-ring walk
// SkipValueStrict hands an array of arrays of numbers (the shared name); it
// has an AVX2 body, so it is useValidRun.
var useValidPoints = useValidRun

// useValid512 selects the walks' AVX-512 bodies, which classify each window
// with compares into mask registers (BW for the byte compares):
// validNumberRun reads it in its first instructions and tail-jumps, and
// validPointsRun here.
var useValid512 = useValidRun && cpu.X86.HasAVX512F && cpu.X86.HasAVX512BW

// validNumberRun passes over the run of plain decimal numbers from data[i:] —
// an array's elements — converting nothing, and returns where the scalar walk
// resumes, or closed = 1 with p at the array's ']'. See floatrun_amd64.s.
//
//go:noescape
func validNumberRun(data []byte, i int) (p, closed int)

// validPointsRun passes over a run of points — an array's elements that are
// themselves flat arrays of plain decimals, a coordinate ring — from the
// point at data[i], converting nothing, and returns where the scalar walk
// resumes, or closed = 1 with p at the ring's ']'. It is called once a ring,
// so the body is chosen here rather than in the assembly.
func validPointsRun(data []byte, i int) (p, closed int) {
	if useValid512 {
		return validPointsRun512(data, i)
	}
	return validPointsRunAVX2(data, i)
}

//go:noescape
func validPointsRun512(data []byte, i int) (p, closed int)

//go:noescape
func validPointsRunAVX2(data []byte, i int) (p, closed int)

// validNumberRun512 is reached only by validNumberRun's tail jump; the
// declaration is what asmdecl checks the frame against.
//
//nolint:unused // called from assembly
//go:noescape
func validNumberRun512(data []byte, i int) (p, closed int)

// parseFloatPointsVBMI and parseFloatPointsAVX2 convert a run of points — an
// array of fixed-size numeric arrays, n numbers each — for
// DecodeFloat64Points. See the assembly.
//
//go:noescape
func parseFloatPointsVBMI(data []byte, i int, out []float64, n int) (np, p, closed int)

//go:noescape
func parseFloatPointsAVX2(data []byte, i int, out []float64, n int) (np, p, closed int)

// parseFloatPoints is called once a ring (and again after each point it hands
// back), so the body is chosen here rather than in the assembly.
func parseFloatPoints(data []byte, i int, out []float64, n int) (np, p, closed int) {
	if useFloatRunVBMI {
		return parseFloatPointsVBMI(data, i, out, n)
	}
	return parseFloatPointsAVX2(data, i, out, n)
}
