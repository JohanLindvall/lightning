//go:build arm64

package unstable

import (
	"math"

	"golang.org/x/sys/cpu"
)

// useFloatRun gates parseFloatRunNEON, the arm64 SIMD kernel for arrays of
// decimal numbers, on the dot-product instruction its digit fold uses (UDOT,
// ARMv8.2 DotProd — the gate useIntRun already uses, for the same reason).
// A var rather than a const so the tests can drive the batch loops with the
// kernel on and off.
var useFloatRun = cpu.ARM64.HasASIMDDP

// useFloatRunLong says the kernel converts numbers of up to 19 digits, with
// Eisel-Lemire — the shape of a coordinate — which is what the points walk and
// the fixed-array reader want (the shared readers read it under that name).
// The arm64 kernel has one body and it does, so this is useFloatRun.
var useFloatRunLong = useFloatRun

// parseFloatRunNEON parses a run of comma-separated decimal numbers of at
// most 19 digits and no exponent from data[i:] into out, and returns how many
// it wrote, where the scalar loop should resume, and whether the run ended at
// the array's ']' (p is then that byte's index). Same contract as the amd64
// kernel; see floatrun_arm64.s.
//
//go:noescape
func parseFloatRunNEON(data []byte, i int, out []float64) (n, p, closed int)

func parseFloatRun(data []byte, i int, out []float64) (n, p, closed int) {
	return parseFloatRunNEON(data, i, out)
}

// parseFloatRunV is parseFloatRun: amd64 has a second body to go straight
// to, arm64 only the one.
func parseFloatRunV(data []byte, i int, out []float64) (n, p, closed int) {
	return parseFloatRunNEON(data, i, out)
}

// parseFloatPointsNEON converts a run of points — an array of fixed-size
// numeric arrays, n numbers each — for DecodeFloat64Points. See the assembly.
//
//go:noescape
func parseFloatPointsNEON(data []byte, i int, out []float64, n int) (np, p, closed int)

func parseFloatPoints(data []byte, i int, out []float64, n int) (np, p, closed int) {
	return parseFloatPointsNEON(data, i, out, n)
}

// useValidRun gates validNumberRunNEON, the flat walk as a check, which
// SkipValueStrict hands an array of numbers; useValidPoints gates
// validPointsRunNEON, the points walk as a check, which it hands a coordinate
// ring. Both need only NEON; they are the conversion kernel's gate so that one
// flag turns every walk on or off in the tests.
var useValidRun, useValidPoints = useFloatRun, useFloatRun

// validNumberRunNEON passes over the run of plain decimal numbers from
// data[i:] — an array's elements — converting nothing, and returns where the
// scalar walk resumes, or closed = 1 with p at the array's ']'. See
// floatrun_arm64.s.
//
//go:noescape
func validNumberRunNEON(data []byte, i int) (p, closed int)

// validPointsRunNEON passes over a run of points — a coordinate ring's
// elements, flat arrays of plain decimals — from the point at data[i],
// converting nothing, and returns where the scalar walk resumes, or closed =
// 1 with p at the ring's ']'. See floatrun_arm64.s.
//
//go:noescape
func validPointsRunNEON(data []byte, i int) (p, closed int)

func validNumberRun(data []byte, i int) (p, closed int) {
	return validNumberRunNEON(data, i)
}

func validPointsRun(data []byte, i int) (p, closed int) {
	return validPointsRunNEON(data, i)
}

// floatRunTab is the kernel's tables, reached through one base register and
// laid out so that the two lookups every short number makes are each one
// scaled register-offset load:
//
//   - at qword 16*neg+L2 (L2 = 0..15), the short divisor ±10^L2, the sign
//     taken from the '-' (qwords 0-31: the slots the (L1, L2) controls below
//     leave free, L1 being at least 1);
//   - at 16*(16*L1+L2), for a number of L1 integer and L2 fraction digits
//     with L1+L2 <= 15, the TBL control that right-aligns its digits out of
//     the sixteen bytes loaded at the first digit, skipping the '.', with
//     zeros in front (0xff selects nothing, which TBL reads as zero);
//   - from 4096, the long divisors: 10^k at qword 512+k and -10^k at qword
//     544+k, exact for every k the kernel uses (k <= 19; the powers are exact
//     through 10^22);
//   - at qword 576+L2 (L2 = 0..19), the high word of the 128-bit significand
//     of 10^-L2, and at qword 596+L2 eiselLemire64's biased binary exponent
//     estimate for it less 2 (the "-= 1 ^ msb" and the "- 1" the assembly
//     wants folded in, as in the amd64 table);
//   - at qword 616+L2, the low word of the same significand, for the
//     refinement eiselLemire64 makes when the high product is ambiguous;
//   - at 32*(160+32*L1+L2), for 16 <= L1+L2 <= 19, two controls: the first
//     right-aligns the number's last sixteen digits out of the two 16-byte
//     loads at its first digit and ending at its last byte (a two-register
//     TBL: an index of 16 or more selects the second load), and the second
//     puts its first L1+L2-16 digits, one to three of them, right-aligned in
//     lanes 1-3, where the top-digit UDOT (weights 0, 100, 10, 1) folds them.
var floatRunTab [25600]byte

func init() {
	put64 := func(off int, v uint64) {
		for j := 0; j < 8; j++ {
			floatRunTab[off+j] = byte(v >> (8 * j))
		}
	}
	p := 1.0
	for k := 0; k < 32; k++ {
		if k > 22 {
			p = 1 // never read: L2 is at most 19
		}
		if k < 16 {
			put64(8*k, math.Float64bits(p))
			put64(8*(16+k), math.Float64bits(-p))
		}
		put64(8*(512+k), math.Float64bits(p))
		put64(8*(544+k), math.Float64bits(-p))
		p *= 10
	}
	for l2 := 0; l2 <= 19; l2++ {
		exp10 := -l2
		pow := detailedPowersOfTen[exp10-detailedPowersOfTenMinExp10]
		put64(8*(576+l2), pow[1])
		put64(8*(596+l2), uint64(217706*exp10>>16+64+1023)-2)
		put64(8*(616+l2), pow[0])
	}
	for l1 := 1; l1 <= 15; l1++ {
		for l2 := 0; l1+l2 <= 15; l2++ {
			c := floatRunTab[16*(16*l1+l2):][:16]
			l := l1 + l2
			for t := range c {
				c[t] = 0xff
				if j := t - (16 - l); j >= 0 {
					src := j
					if j >= l1 {
						src++ // past the '.'
					}
					c[t] = byte(src)
				}
			}
		}
	}
	for l1 := 1; l1 <= 19; l1++ {
		for l2 := 0; l1+l2 <= 19; l2++ {
			l := l1 + l2
			if l < 16 {
				continue
			}
			span := l1
			if l2 > 0 {
				span += 1 + l2
			}
			off := func(j int) int { // the byte offset of digit j from the first digit
				if j >= l1 {
					return j + 1
				}
				return j
			}
			lo := floatRunTab[32*(160+32*l1+l2):][:16]
			for k := range lo {
				o := off(l - 16 + k)
				if o < 16 {
					lo[k] = byte(o)
				} else {
					lo[k] = byte(16 + o - (span - 16))
				}
			}
			top := floatRunTab[32*(160+32*l1+l2)+16:][:16]
			for t := range top {
				top[t] = 0xff
			}
			for j := 0; j < l-16; j++ {
				top[4-(l-16)+j] = byte(off(j))
			}
		}
	}
}
