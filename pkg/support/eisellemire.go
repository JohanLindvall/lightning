package support

import (
	"math"
	"math/bits"
)

// eiselLemireParse parses the well-formed JSON number token b (already validated
// by scanFloat, which is why it does no error checking) and converts it with
// Eisel-Lemire, returning ok=false — so the caller falls back to
// strconv.ParseFloat — when the mantissa has more than 19 significant digits (so
// it cannot be held exactly in a uint64) or the value is one EL declines.
//
// The scanner calls this on its slow path, before strconv, so that numbers which
// merely miss the Clinger fast path (a mantissa >= 2^53 or an exponent outside
// [-22, 22]) are still converted without strconv's expensive decimal machinery.
// It lives outside scanFloat — re-scanning the handful of slow-path bytes rather
// than threading state out — so scanFloat's hot path for the common fast floats
// keeps exactly its original, compact form. The digit/exponent extraction
// mirrors scanFloat's so the two agree on what counts as an exact mantissa.
func eiselLemireParse(b []byte) (float64, bool) {
	i, n := 0, len(b)
	neg := false
	if i < n {
		switch b[i] {
		case '-':
			neg = true
			i++
		case '+':
			i++
		}
	}
	var mant uint64
	mdigits, digits := 0, 0
	for i < n {
		d := b[i] - '0'
		if d > 9 {
			break
		}
		if mdigits < 19 {
			mant = mant*10 + uint64(d)
			mdigits++
		}
		i++
		digits++
	}
	exp := 0
	if i < n && b[i] == '.' {
		i++
		for i < n {
			d := b[i] - '0'
			if d > 9 {
				break
			}
			if mdigits < 19 {
				mant = mant*10 + uint64(d)
				mdigits++
				exp--
			}
			i++
			digits++
		}
	}
	if digits > mdigits {
		return 0, false // more than 19 significant digits; defer to strconv
	}
	if i < n && (b[i] == 'e' || b[i] == 'E') {
		i++
		esign := 1
		if i < n && (b[i] == '+' || b[i] == '-') {
			if b[i] == '-' {
				esign = -1
			}
			i++
		}
		eval := 0
		for i < n {
			d := b[i] - '0'
			if d > 9 {
				break
			}
			if eval < 1<<30 {
				eval = eval*10 + int(d)
			}
			i++
		}
		exp += esign * eval
	}
	return eiselLemire64(mant, exp, neg)
}

// eiselLemire64 converts the decimal significand man (an exact integer of at most
// 19 digits, so it fits a uint64 without truncation) scaled by 10^exp10, with the
// given sign, to the correctly rounded float64. It returns ok=false for the
// inputs the algorithm cannot resolve cheaply — a halfway-ambiguous rounding, a
// subnormal or overflowing result, or an exponent outside the table — in which
// case the caller falls back to strconv.ParseFloat.
//
// This is the Eisel-Lemire algorithm (Daniel Lemire, "Number Parsing at a
// Gigabyte per Second"), the same fast path strconv.ParseFloat takes internally;
// running it on the (man, exp10) the scanner already extracted avoids strconv's
// full re-scan of the digit string, which dominates decoding of numbers that miss
// the Clinger fast path. When it returns ok the result is bit-for-bit identical
// to strconv.ParseFloat; the ok=false cases are exactly the ambiguous ones, so
// correctness never depends on this path. detailedPowersOfTen lives in the
// generated powers_table.go.
func eiselLemire64(man uint64, exp10 int, neg bool) (float64, bool) {
	if man == 0 {
		if neg {
			return math.Float64frombits(0x8000000000000000), true // -0
		}
		return 0, true
	}
	if exp10 < detailedPowersOfTenMinExp10 || detailedPowersOfTenMaxExp10 < exp10 {
		return 0, false
	}

	// Normalize the significand so its leading bit is set.
	clz := bits.LeadingZeros64(man)
	man <<= uint(clz)
	const float64ExponentBias = 1023
	// 217706/2^16 approximates log2(10); this estimates the binary exponent.
	retExp2 := uint64(217706*exp10>>16+64+float64ExponentBias) - uint64(clz)

	pow := detailedPowersOfTen[exp10-detailedPowersOfTenMinExp10]
	xHi, xLo := bits.Mul64(man, pow[1])
	// If the high product is within 1/512 of a halfway point, refine with the low
	// half of the 128-bit power of ten.
	if xHi&0x1FF == 0x1FF && xLo+man < man {
		yHi, yLo := bits.Mul64(man, pow[0])
		mergedHi, mergedLo := xHi, xLo+yHi
		if mergedLo < xLo {
			mergedHi++
		}
		if mergedHi&0x1FF == 0x1FF && mergedLo+1 == 0 && yLo+man < man {
			return 0, false // still ambiguous; defer to strconv
		}
		xHi, xLo = mergedHi, mergedLo
	}

	msb := xHi >> 63
	retMantissa := xHi >> (msb + 9)
	retExp2 -= 1 ^ msb

	// An exact halfway value (...10000000000) cannot be rounded here.
	if xLo == 0 && xHi&0x1FF == 0 && retMantissa&3 == 1 {
		return 0, false
	}

	// Round to 53 bits.
	retMantissa += retMantissa & 1
	retMantissa >>= 1
	if retMantissa>>53 > 0 {
		retMantissa >>= 1
		retExp2++
	}

	// Reject subnormal (biased exp <= 0) and overflow (biased exp >= 0x7FF)
	// results; strconv handles those exactly.
	if retExp2-1 >= 0x7FF-1 {
		return 0, false
	}

	retBits := retExp2<<52 | retMantissa&0x000FFFFFFFFFFFFF
	if neg {
		retBits |= 0x8000000000000000
	}
	return math.Float64frombits(retBits), true
}
