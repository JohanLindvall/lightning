package unstable

import (
	"encoding/binary"
	"math"
	"math/bits"
	"strconv"
)

// pow10exact holds the powers of ten that are exactly representable as a
// float64 (10^0 .. 10^22). Used by the fast-path float parser.
var pow10exact = [...]float64{
	1e0, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6, 1e7, 1e8, 1e9, 1e10, 1e11,
	1e12, 1e13, 1e14, 1e15, 1e16, 1e17, 1e18, 1e19, 1e20, 1e21, 1e22,
}

// scanFloat scans the JSON number token beginning at data[i] in a single pass,
// returning the index just past it and, when it can resolve the value exactly,
// the parsed result with fast=true. Two fast paths run on the mantissa and
// exponent the scan extracts, so no rescan is needed: the Clinger path (an
// exactly representable mantissa < 2^53 with a decimal exponent in [-22, 22],
// converted with one multiply or divide) and, for an exact mantissa that misses
// Clinger, Eisel-Lemire (a 128-bit multiply — see eiselLemire64). Values with a
// mantissa of more than 19 significant digits, ambiguous rounding, subnormal/
// overflow results, or an exponent outside the powers-of-ten table return
// fast=false, leaving the caller to parse data[i:end] with strconv.ParseFloat.
//
// Folding the token scan and both fast-path parses into one pass spares the
// separate skipNumber scan the previous two-call form made; and because the scan
// always runs to the end of the token, the slow path no longer pays for the
// fast-path parser's full rescan-then-reject before handing off to strconv.
func scanFloat(data []byte, i int) (f float64, end int, fast, ok bool) {
	n := len(data)
	neg := false
	if i < n {
		switch data[i] {
		case '-':
			neg = true
			i++
		case '+':
			i++
		}
	}

	// Accumulate up to 19 significant digits (the most that always fit a uint64)
	// into mant; count every digit so a longer mantissa is recognized and routed
	// to strconv rather than silently truncated.
	// The digit loops test each byte with the unsigned trick d := c-'0'; d > 9,
	// which is one comparison rather than the two a '0' <= c <= '9' range needs
	// (a byte below '0' underflows to a large value, so it also fails d > 9).
	var mant uint64
	digits, mdigits := 0, 0
	for i < n {
		d := data[i] - '0'
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
	if i < n && data[i] == '.' {
		i++
		// Leading fraction zeros (the "000" of 0.000698…) are not significant
		// digits: they only shift the decimal exponent. Skip them here so they do
		// not consume the 19-digit mantissa budget — otherwise a value like
		// 0.0006988752666567719 (16 significant digits, comfortably exact) would be
		// counted as 20 digits, falsely flagged as overflow, and dropped to strconv
		// instead of taking the Clinger/Eisel-Lemire fast path. Only when no
		// significant digit has been seen yet (mant == 0); a zero between nonzero
		// digits is significant and stays in the loops below.
		if mant == 0 {
			for i < n && data[i] == '0' {
				exp--
				i++
			}
		}
		// Fold fractional digits four bytes at a time while they fit the 19-digit
		// budget, replacing per-digit mant*10+d steps with one SWAR multiply chain
		// per chunk. The 1-3 digit tail (and any digits past the 19-digit budget)
		// drops to the scalar loop below.
		for mdigits+4 <= 19 && i+4 <= n {
			w := binary.LittleEndian.Uint32(data[i : i+4])
			if !is4Digits(w) {
				break
			}
			mant = mant*10000 + uint64(parse4Digits(w))
			mdigits += 4
			digits += 4
			exp -= 4
			i += 4
		}
		for i < n {
			d := data[i] - '0'
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
	overflow := digits > mdigits // more digits than the uint64 mantissa holds
	if i < n && (data[i] == 'e' || data[i] == 'E') {
		i++
		esign := 1
		if i < n && (data[i] == '+' || data[i] == '-') {
			if data[i] == '-' {
				esign = -1
			}
			i++
		}
		ed, eval := 0, 0
		for i < n {
			d := data[i] - '0'
			if d > 9 {
				break
			}
			if eval < 1<<30 { // clamp; an exponent this large can never be exact
				eval = eval*10 + int(d)
			}
			i++
			ed++
		}
		if ed == 0 {
			return 0, i, false, false // exponent marker with no digits
		}
		exp += esign * eval
	}
	end = i
	if digits == 0 {
		return 0, end, false, false
	}
	// A well-formed number ends here. A trailing number-continuation byte means a
	// malformed token such as "1.2.3" or "1e2e3"; consume the rest of the run (as
	// skipNumber would) and reject, so the reported end and error match the slow
	// path rather than silently accepting the leading "1.2".
	if end < n {
		if c := data[end]; c == '.' || c == 'e' || c == 'E' || (c >= '0' && c <= '9') {
			for end < n {
				c = data[end]
				if (c >= '0' && c <= '9') || c == '.' || c == 'e' || c == 'E' || c == '+' || c == '-' {
					end++
					continue
				}
				break
			}
			return 0, end, false, false
		}
	}
	if overflow || mant>>53 != 0 {
		if !overflow {
			if v, ok := eiselLemire64(mant, exp, neg); ok {
				return v, end, true, true
			}
		}
		return 0, end, false, true // hand the exact token to strconv
	}

	f = float64(mant)
	switch {
	case exp == 0:
		// already exact
	case exp > 0 && exp <= 22:
		f *= pow10exact[exp]
	case exp < 0 && exp >= -22:
		f /= pow10exact[-exp]
	default:
		if v, ok := eiselLemire64(mant, exp, neg); ok {
			return v, end, true, true
		}
		return 0, end, false, true
	}
	if neg {
		f = -f
	}
	return f, end, true, true
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

// ParseFloat parses the JSON number in b as a float64. It takes the same Clinger
// fast path as the scanner — an exact mantissa with a small decimal exponent is
// converted with a single multiply or divide — and falls back to
// strconv.ParseFloat for everything else. b must be exactly one number with no
// surrounding whitespace; trailing bytes or an empty input yield ErrBadNumber.
func ParseFloat(b []byte) (float64, error) {
	f, end, fast, ok := scanFloat(b, 0)
	if !ok || end != len(b) {
		return 0, ErrBadNumber
	}
	if fast {
		return f, nil
	}
	// unsafeStr avoids copying the token; ParseFloat does not retain it.
	f, err := strconv.ParseFloat(unsafeStr(b), 64)
	if err != nil {
		return 0, ErrBadNumber
	}
	return f, nil
}

// is4Digits reports whether the four bytes packed little-endian in w are all
// ASCII digits '0'..'9' (the simdjson bit trick).
func is4Digits(w uint32) bool {
	return (w&0xF0F0F0F0)|(((w+0x06060606)&0xF0F0F0F0)>>4) == 0x33333333
}

// parse4Digits folds four ASCII digits packed little-endian in w into their
// integer value. The caller has verified the bytes are digits with is4Digits.
func parse4Digits(w uint32) uint32 {
	w -= 0x30303030
	lo := w & 0x00FF00FF
	hi := (w >> 8) & 0x00FF00FF
	w = lo*10 + hi
	return (w&0x0000FFFF)*100 + (w >> 16)
}
