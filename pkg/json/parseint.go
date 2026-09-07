package json

import "math"

// ParseInt parses the JSON integer in b as an int64. It is the integer
// counterpart of [ParseFloat] and the reader a caller wants for the value a
// [Get] or an [ObjectEach] hands back when that value is an integer, or for a
// quoted 64-bit integer once its quotes are stripped — ClickHouse, for one,
// quotes every Int64/UInt64 so it survives JavaScript's 53-bit numbers, and
// the alternative to this function is allocating a string for
// strconv.ParseInt on every such cell.
//
// b must be exactly one integer token with no surrounding whitespace: an
// optional sign followed by decimal digits. It accepts the sign forms the rest
// of this library accepts — a leading '+' as well as '-' — and leading zeros,
// like [ParseFloat]. A fraction or an exponent (1.0, 1e3), an empty input,
// stray bytes, and a magnitude outside int64 are all ErrBadNumber: this is a
// parser for an integer, not a float that happens to be whole. Nothing is
// retained or copied.
func ParseInt(b []byte) (int64, error) {
	neg, i := false, 0
	if len(b) > 0 && (b[0] == '-' || b[0] == '+') {
		neg = b[0] == '-'
		i = 1
	}
	n, err := parseDigits(b, i)
	if err != nil {
		return 0, err
	}
	// The magnitude int64 can hold is one larger on the negative side, and
	// int64(1<<63) wraps to MinInt64, whose negation is itself — so the
	// conversion below is exact at both ends.
	if neg && n > 1<<63 || !neg && n > math.MaxInt64 {
		return 0, ErrBadNumber
	}
	v := int64(n)
	if neg {
		v = -v
	}
	return v, nil
}

// ParseUint parses the JSON integer in b as a uint64, with [ParseInt]'s
// grammar and rules: a leading '+' is accepted (and only a '+' — a '-' is
// ErrBadNumber, as is any magnitude past uint64).
func ParseUint(b []byte) (uint64, error) {
	i := 0
	if len(b) > 0 && b[0] == '+' {
		i = 1
	}
	return parseDigits(b, i)
}

// parseDigits reads b[i:] as a run of decimal digits into a uint64, rejecting
// an empty run, any non-digit byte, and overflow.
func parseDigits(b []byte, i int) (uint64, error) {
	if i >= len(b) {
		return 0, ErrBadNumber
	}
	// Past cutoff a further digit cannot fit whatever it is; below it, the
	// multiply is safe and only the add can wrap.
	const cutoff = math.MaxUint64/10 + 1
	var n uint64
	for ; i < len(b); i++ {
		d := uint64(b[i] - '0')
		if d > 9 {
			return 0, ErrBadNumber
		}
		if n >= cutoff {
			return 0, ErrBadNumber
		}
		n *= 10
		if n+d < n {
			return 0, ErrBadNumber
		}
		n += d
	}
	return n, nil
}
