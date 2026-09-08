package unstable

import "math/bits"

// ParseUint parses b, which must be exactly one JSON integer token with no
// surrounding whitespace, as a uint64: an optional leading '+' (the sign form
// the rest of this package accepts — ParseFloat takes it too) followed by
// decimal digits, leading zeros allowed. A fraction or an exponent, an empty
// input, a stray byte, a '-', and a magnitude past uint64 are all
// ErrBadNumber. Nothing is retained or copied.
//
// The token's length is known before a byte is folded, which is what lets the
// whole parse be straight-line SWAR rather than the digit loop the readers in
// read.go use (they scan a cursor into a document and cannot know where the
// run ends). Three consequences, in order of what they are worth:
//
//   - A run of at most 19 digits cannot overflow a uint64, so the per-digit
//     cutoff test the loop form pays disappears. Only a 20-digit token needs a
//     checked multiply, and only a longer one — which, leading zeros aside, is
//     out of range whatever its bytes are — needs to be rejected outright.
//   - The digits are folded eight at a time by parse8Digits' three multiplies
//     instead of one multiply-add per digit, and the folds of the separate
//     words are independent, so they issue in parallel. That is the whole win:
//     a 13-digit token has a 13-long multiply chain in the loop form.
//   - The folds decompose the token from the RIGHT: the last eight digits are
//     a whole word needing no mask, and the digits above them are the same
//     word shifted up so the vacated low lanes read as leading zeros, which
//     makes every scale a compile-time constant.
//
// Coverage is what validates: the words overlap, and between them they touch
// every byte of the token, so one digit test per word rejects a non-digit
// anywhere in it. The length dispatch is written as unsigned range tests so
// that the shortest token — the common one — is reached in a single compare
// and the degenerate lengths fall out of the same tests.
func ParseUint(b []byte) (uint64, error) {
	i := 0
	if len(b) > 0 && b[0] == '+' {
		i = 1
	}
	end := len(b)
	nd := end - i
retry:
	if uint(nd-1) < 3 {
		// Too short for a four-byte load to stay inside the token, and a
		// chain of at most two multiplies is cheaper than a fold anyway.
		//
		// This walks an index rather than ranging over b[i:]. The reslice is
		// the seven instructions IndexCloseOrEscapeAt's entry describes — a
		// bounds compare, the cap and len subtractions and the negative-length
		// clamp on the base pointer — on the arm every one-to-three-digit
		// token takes, which is the common one. The unsigned loop condition is
		// what leaves b[i] with no bounds check of its own, and that is worth
		// more than the compares: runtime.panicBounds is a CALL, so one
		// surviving check would give this whole leaf function a stack frame.
		// uint64(b[i]) - '0' widens before subtracting for the same reason the
		// walkers' number predicate does, to avoid a byte truncation.
		var n uint64
		for uint(i) < uint(len(b)) {
			d := uint64(b[i]) - '0'
			if d > 9 {
				return 0, ErrBadNumber
			}
			n *= 5 // two LEAs, not three; see ReadInt64OrNull
			n = d + n<<1
			i++
		}
		return n, nil
	}
	if uint(nd-4) < 5 {
		// Two four-byte loads, at the token's start and at its end, OR'd into
		// one eight-lane word: they overlap when nd < 8 and the overlapping
		// lanes hold the same bytes, so the OR is exact.
		n, ok := foldDigitWord(((uint64(load32(b, i)) ^ uint64(swarZero32)) << (uint(8-nd) * 8)) |
			((uint64(load32(b, end-4)) ^ uint64(swarZero32)) << 32))
		if !ok {
			return 0, ErrBadNumber
		}
		return n, nil
	}
	if uint(nd-9) < 12 {
		lo, ok := foldDigitWord(load64(b, end-8) ^ swarZero)
		if !ok {
			return 0, ErrBadNumber
		}
		if nd <= 16 {
			hi, ok := foldDigitWord((load64(b, i) ^ swarZero) << (uint(16-nd) * 8))
			if !ok {
				return 0, ErrBadNumber
			}
			return hi*100000000 + lo, nil
		}
		mid, ok := foldDigitWord(load64(b, end-16) ^ swarZero)
		if !ok {
			return 0, ErrBadNumber
		}
		top, ok := foldDigitWord((load64(b, i) ^ swarZero) << (uint(24-nd) * 8))
		if !ok {
			return 0, ErrBadNumber
		}
		n := mid*100000000 + lo
		if nd < 20 {
			return top*10000000000000000 + n, nil
		}
		// The only multiply here that can overflow: 17 to 19 digits are at
		// most 10^19-1, and the low 16 digits are exact.
		h, l := bits.Mul64(top, 10000000000000000)
		if h != 0 || l+n < l {
			return 0, ErrBadNumber
		}
		return l + n, nil
	}
	// Nothing above matched: the token is empty, or longer than 20 digits and
	// so out of range — unless it carries leading zeros, the one way a long
	// token can still hold a small number. Skipping them is worth a pass only
	// here, and it can only shorten, so the retry runs at most twice.
	if nd > 0 {
		for uint(i) < uint(len(b)) && b[i] == '0' {
			i++
		}
		if n := end - i; n != nd {
			nd = n
			if nd == 0 {
				return 0, nil // the token was all zeros
			}
			goto retry
		}
	}
	return 0, ErrBadNumber
}

// ParseInt is ParseUint's grammar as an int64: an optional '-' or '+' and
// then the digits, with a magnitude outside int64 reported as ErrBadNumber.
//
// The body below is ParseUint's, written out rather than called: the shared
// call measured 0.7-1.0 ns against a parse that is under 3 ns for a short
// token, a fifth of it, and the two are held to each other and to strconv by
// TestParseIntMatchesStrconvUnstable and FuzzParseIntMatchesStrconv. Every
// comment in ParseUint applies here. Two things differ beyond the sign: each
// case returns where it folds rather than joining a common tail (a merge point
// cost a one-digit parse 30%), and the range check is written only in the arm
// that can fail it — 16 digits are at most 10^16, two orders under MaxInt64,
// so every shorter token is in range by construction.
func ParseInt(b []byte) (int64, error) {
	i, neg := 0, false
	if len(b) > 0 {
		if c := b[0]; c == '-' {
			neg, i = true, 1
		} else if c == '+' {
			i = 1
		}
	}
	end := len(b)
	nd := end - i
retry:
	if uint(nd-1) < 3 {
		var n uint64
		for uint(i) < uint(len(b)) {
			d := uint64(b[i]) - '0'
			if d > 9 {
				return 0, ErrBadNumber
			}
			n *= 5 // two LEAs, not three; see ReadInt64OrNull
			n = d + n<<1
			i++
		}
		return withSign(n, neg), nil
	}
	if uint(nd-4) < 5 {
		n, ok := foldDigitWord(((uint64(load32(b, i)) ^ uint64(swarZero32)) << (uint(8-nd) * 8)) |
			((uint64(load32(b, end-4)) ^ uint64(swarZero32)) << 32))
		if !ok {
			return 0, ErrBadNumber
		}
		return withSign(n, neg), nil
	}
	if uint(nd-9) < 12 {
		lo, ok := foldDigitWord(load64(b, end-8) ^ swarZero)
		if !ok {
			return 0, ErrBadNumber
		}
		if nd <= 16 {
			hi, ok := foldDigitWord((load64(b, i) ^ swarZero) << (uint(16-nd) * 8))
			if !ok {
				return 0, ErrBadNumber
			}
			return withSign(hi*100000000+lo, neg), nil
		}
		mid, ok := foldDigitWord(load64(b, end-16) ^ swarZero)
		if !ok {
			return 0, ErrBadNumber
		}
		top, ok := foldDigitWord((load64(b, i) ^ swarZero) << (uint(24-nd) * 8))
		if !ok {
			return 0, ErrBadNumber
		}
		n := mid*100000000 + lo
		if nd < 20 {
			n += top * 10000000000000000
		} else {
			h, l := bits.Mul64(top, 10000000000000000)
			if h != 0 || l+n < l {
				return 0, ErrBadNumber
			}
			n = l + n
		}
		// int64 holds one more magnitude on the negative side, and
		// int64(1<<63) is MinInt64, whose negation is itself, so withSign is
		// exact at both ends.
		if n > 1<<63 || (n == 1<<63 && !neg) {
			return 0, ErrBadNumber
		}
		return withSign(n, neg), nil
	}
	if nd > 0 {
		for uint(i) < uint(len(b)) && b[i] == '0' {
			i++
		}
		if n := end - i; n != nd {
			nd = n
			if nd == 0 {
				return 0, nil
			}
			goto retry
		}
	}
	return 0, ErrBadNumber
}

// withSign applies a parsed sign. The negation is written as a branch because
// the compiler turns it into a conditional move.
func withSign(n uint64, neg bool) int64 {
	v := int64(n)
	if neg {
		v = -v
	}
	return v
}

// foldDigitWord checks that all eight lanes of d — a word already reduced by
// '0', lane 0 the most significant digit — are digits and folds them into
// their value. It is small enough to inline, which is the point: the fold is
// then written at each site with no call.
func foldDigitWord(d uint64) (uint64, bool) {
	if ((d+swarSix)|d)&swarNib != 0 {
		return 0, false
	}
	return parse8Digits(d), true
}

// swarZero32 is swarZero's four-lane form, for the loads that cover a token
// too short for an eight-byte one.
const swarZero32 = uint32(0x01010101) * '0'
