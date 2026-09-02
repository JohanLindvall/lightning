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
	// Straight-line fast path for the shape nearly every real-world float has:
	// an optional '-', 1-7 integer digits, an optional '.' with 1-24 fraction
	// digits, an optional exponent of 1-7 digits, at most 19 significant digits.
	//
	// The whole 32-byte window after the sign is loaded up front at FIXED
	// offsets, and the fraction is aligned to a word boundary with funnel
	// shifts by the integer-digit count, so every load issues at once and the
	// digit masks and folds of the three fraction words proceed in parallel. A
	// first version loaded each word at the offset the previous word's digit
	// count produced; that turned the loop form's well-predicted control
	// dependencies into a load-to-address DATA chain (load, mask, count, add,
	// load, ... ~13 cycles a link), and a 17-digit fraction ran 18% slower
	// than the loops it replaced despite executing fewer instructions. The
	// loop form (scanFloatSlow) costs 289 instructions and 23 taken branches
	// for an 8-byte float on Meteor Lake; this path is the arithmetic alone.
	//
	// Each digit run is measured with one SWAR mask (the lowest flagged lane
	// is exact: borrows and carries only travel upward, and every lane below
	// the first non-digit is a digit) and folded with one fixed 8-digit fold
	// after shifting the run into the top lanes, the vacated low lanes reading
	// as leading zeros. Significance follows scanFloatSlow to the digit —
	// leading fraction zeros do not count against the 19-digit budget when the
	// integer part is zero — so golang_source's 0.000698…-shaped weights stay
	// here. Anything unrecognised (a '+', 8+ integer digits, 25+ fraction
	// digits, more than 19 significant digits, an 8+ digit exponent, a
	// malformed continuation, or a token within 48 bytes of the end of the
	// buffer, the window the unchecked loads need) is handed whole to
	// scanFloatSlow, whose results this path reproduces bit for bit —
	// TestScanFloatFastMatchesSlow pins that over generated tokens.
	if uint(i)+48 <= uint(len(data)) {
		const (
			hi   = swarHi
			zero = swarZero
			bias = swarBias
			one  = swarOne
		)
		start := i
		neg := data[i] == '-'
		if neg {
			i++
		}
		w0 := load64(data, i)
		w1 := load64(data, i+8)
		w2 := load64(data, i+16)
		w3 := load64(data, i+24)
		d0 := w0 - zero
		k := bits.TrailingZeros64((d0|(d0+bias))&hi) >> 3 // integer digits, 8 = all
		if uint(k-1) < 7 {                                // 1..7 (k == 0 wraps)
			mant := parse8Digits(d0 << (uint(8-k) * 8 & 63))
			p := k // offset just past the digits consumed so far
			exp := 0
			if byte(w0>>(uint(k)*8&63)) == '.' {
				// The fraction starts at lane k+1: funnel the window down by
				// that many lanes so it begins each word. k is 1..7, so the
				// shift pairs are 8..56 and 48..0 — no shift reaches 64.
				sh := uint(k) * 8 & 63
				f0 := (w0>>8)>>sh | w1<<(56-sh)
				f1 := (w1>>8)>>sh | w2<<(56-sh)
				f2 := (w2>>8)>>sh | w3<<(56-sh)
				fd0 := f0 - zero
				m0 := (fd0 | (fd0 + bias)) & hi
				var kf int
				var fv uint64
				if m0 != 0 { // 1-7 fraction digits, or none
					kf = bits.TrailingZeros64(m0) >> 3
					if kf == 0 {
						return scanFloatSlow(data, start) // "1." and "1.x": the slow path decides
					}
					fv = parse8Digits(fd0 << (uint(8-kf) * 8 & 63))
				} else {
					fd1 := f1 - zero
					m1 := (fd1 | (fd1 + bias)) & hi
					if m1 != 0 { // 8-15
						r := bits.TrailingZeros64(m1) >> 3
						kf = 8 + r
						// (x<<8)<<((7-r)*8) lands r digits in the top lanes and, for
						// r == 0, shifts the word out entirely — no 64-bit shift.
						fv = parse8Digits(fd0)*pow10u64[r] + parse8Digits((fd1<<8)<<(uint(7-r)*8&63))
					} else {
						fd2 := f2 - zero
						m2 := (fd2 | (fd2 + bias)) & hi
						if m2 == 0 {
							return scanFloatSlow(data, start) // 25+ fraction digits
						}
						r := bits.TrailingZeros64(m2) >> 3 // 16-23
						kf = 16 + r
						fv = (parse8Digits(fd0)*1e8+parse8Digits(fd1))*pow10u64[r] + parse8Digits((fd2<<8)<<(uint(7-r)*8&63))
					}
				}
				// Significant digits: the integer digits plus the fraction digits,
				// less the fraction's leading zeros when the integer part is zero
				// (the slow path skips those without counting them). Only the first
				// word's zeros are discounted, which can only send a number with 9+
				// leading zeros to the slow path early, never admit one wrongly; lz
				// is clamped to the run because a non-digit lane above it can read
				// as zero.
				sig := k + kf
				if mant == 0 {
					lz := bits.TrailingZeros64((fd0+one)&hi) >> 3
					if lz > kf {
						lz = kf
					}
					sig -= lz
				}
				if sig > 19 {
					return scanFloatSlow(data, start) // past the budget: strconv decides
				}
				mant = mant*pow10u64[kf] + fv // mant != 0 implies kf <= 18 here
				p = k + 1 + kf
				exp = -kf
			}
			// The byte after the digits: an exponent of 1-7 digits is folded the
			// same way; a digit cannot follow a measured run, so only a '.', 'e' or
			// 'E' can still make the token malformed ("1.2.3", "1e2e3"), and those
			// go to the slow path, which measures and rejects the whole run.
			switch c := data[i+p]; c {
			case 'e', 'E':
				q := i + p + 1
				esign := 1
				if c := data[q]; c == '+' || c == '-' {
					if c == '-' {
						esign = -1
					}
					q++
				}
				de := load64(data, q) - zero // q <= i+34: the window covers it
				ke := bits.TrailingZeros64((de|(de+bias))&hi) >> 3
				if uint(ke-1) >= 7 { // no digits ("1e", "1e+"), or 8+: the slow path decides
					return scanFloatSlow(data, start)
				}
				exp += esign * int(parse8Digits(de<<(uint(8-ke)*8&63)))
				q += ke
				if c := data[q]; c == '.' || c == 'e' || c == 'E' {
					return scanFloatSlow(data, start)
				}
				end = q
			case '.':
				return scanFloatSlow(data, start)
			default:
				end = i + p
			}
			if mant>>53 != 0 {
				if v, ok := eiselLemire64(mant, exp, neg); ok {
					return v, end, true, true
				}
				return 0, end, false, true // hand the exact token to strconv
			}
			f = float64(mant)
			if exp < 0 {
				if exp < -22 {
					if v, ok := eiselLemire64(mant, exp, neg); ok {
						return v, end, true, true
					}
					return 0, end, false, true
				}
				f /= pow10exact[-exp]
			} else if exp > 0 {
				if exp > 22 {
					if v, ok := eiselLemire64(mant, exp, neg); ok {
						return v, end, true, true
					}
					return 0, end, false, true
				}
				f *= pow10exact[exp]
			}
			if neg {
				f = -f
			}
			return f, end, true, true
		}
		i = start
	}
	return scanFloatSlow(data, i)
}

// The SWAR digit constants: swarZero subtracts '0' from every lane, swarBias
// then carries any lane above 9 (0x0a..0x7f) into its top bit while a digit
// (0..9) reaches at most 0x7f, and a byte below '0' wrapped past 0x80 by the
// subtraction, so (d | (d + swarBias)) & swarHi flags exactly the non-digit
// lanes; swarOne likewise flags the digit lanes that are not zero.
const (
	swarLo   = ^uint64(0) / 255 // 0x0101...01
	swarHi   = swarLo << 7      // 0x8080...80
	swarZero = swarLo * '0'
	swarBias = swarLo * 0x76
	swarOne  = swarLo * 0x7f
)

// pow10u64 holds 10^0 .. 10^19 as integers, the scale that shifts an integer
// part past a measured run of fraction digits in scanFloat's fast path. It is
// indexed by a fraction length of up to 24, but a nonzero integer part is only
// ever scaled within the 19-digit budget; the entries past 10^19 are reached
// only with a zero multiplicand and hold zero.
var pow10u64 = [25]uint64{
	1, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6, 1e7, 1e8, 1e9, 1e10, 1e11, 1e12, 1e13, 1e14, 1e15, 1e16, 1e17, 1e18, 1e19,
}

// parse8Digits folds eight digit lanes (each already reduced by '0', lane 0 the
// most significant digit) into their value — simdjson's parse_eight_digits_
// unrolled. Lanes shifted in as zero read as leading '0' digits, which is how
// scanFloat folds a run shorter than eight with the same three multiplies.
func parse8Digits(d uint64) uint64 {
	d = d*10 + d>>8
	d = (d&0x000000FF000000FF)*(100+(1000000<<32)) + ((d>>16)&0x000000FF000000FF)*(1+(10000<<32))
	return d >> 32
}

// scanFloatSlow is the general loop form of scanFloat (see there): it handles
// every token shape and is the oracle the fast path is pinned against.
func scanFloatSlow(data []byte, i int) (f float64, end int, fast, ok bool) {
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
			// ok4, not ok: this function's result is a named `ok` return, and a
			// `:=` here would shadow it inside the loop body.
			v, ok4 := tryParse4Digits(binary.LittleEndian.Uint32(data[i : i+4]))
			if !ok4 {
				break
			}
			mant = mant*10000 + uint64(v)
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

// ParseFloat parses the number in b as a float64. It takes the same Clinger
// fast path as the scanner — an exact mantissa with a small decimal exponent is
// converted with a single multiply or divide — and falls back to
// strconv.ParseFloat for everything else. b must be exactly one number with no
// surrounding whitespace; trailing bytes or an empty input yield ErrBadNumber.
//
// What it accepts is scanFloat's grammar, which is deliberately a superset of
// RFC 8259's number: a leading '+' as well as '-' (ParseFloat("+5") is 5, nil),
// leading zeros ("01"), an empty integer part (".5") and an empty fraction
// ("1."). That is not laxness for its own sake — it is the accept set of every
// number reader in this package, so ParseFloat, Valid and a generated decoder
// agree on which documents are numbers; TestValidDivergesFromStdlib pins the
// same list from Valid's side. In the other direction it is narrower than the
// JSON grammar in one place: a magnitude no float64 can represent (1e309) is
// ErrBadNumber, since there is no value to return.
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
// ASCII digits '0'..'9' (the simdjson bit trick). Kept for the differential
// test that pins tryParse4Digits' accept set to it; the SWAR digit loops all
// call tryParse4Digits, which decides and folds in one pass.
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

// tryParse4Digits folds the four bytes packed little-endian in w into the
// integer they spell, reporting false (and leaving the value unspecified) if any
// of them is not an ASCII digit. It is the digit step of every SWAR fold here —
// ReadInt64OrNull/ReadUint64OrNull, batch.go's four array readers, and
// scanFloat's fraction loop — replacing an is4Digits test followed by a
// parse4Digits fold.
//
// It exists because the split form paid for its constants twice over on arm64.
// AArch64 folds an immediate into a logical op only when it is a "bitmask
// immediate" (a rotated run of ones, replicated), and ADD/SUB take a 12-bit
// immediate and nothing wider — so is4Digits' 0x06060606 (an ADD operand) and
// 0x33333333 (a CMP operand) each cost a MOVD plus three MOVKs, IN the loop and
// again on the back edge, and parse4Digits then materialised 0x30303030 for its
// own subtraction. Deciding on the SAME d = w - 0x30303030 the fold needs
// removes one constant outright and trades the other two for one.
//
// The test is the classic packed-digit form: d = w - 0x30303030 wraps a byte
// below '0' to >= 0xd0, so its own top bit flags it, and adding 0x76 to a byte
// in 0x0a..0x7f carries into the top bit while a real digit (0..9) reaches only
// 0x7f. It is EXACTLY is4Digits, not an approximation. A lane can only be
// disturbed by a carry out of the lane below it, and a lane that carries has
// already set its own flag (d >= 0x8a there), so the OR is nonzero either way;
// the proof runs in both directions and TestTryParse4DigitsMatchesIs4Digits
// checks it over every digit-adjacent byte combination plus random words.
func tryParse4Digits(w uint32) (uint32, bool) {
	d := w - 0x30303030
	if (d|(d+0x76767676))&0x80808080 != 0 {
		return 0, false
	}
	lo := d & 0x00FF00FF
	hi := (d >> 8) & 0x00FF00FF
	x := lo*10 + hi
	return (x&0x0000FFFF)*100 + (x >> 16), true
}
