package unstable

import (
	"encoding/binary"
	"strconv"
	"time"
)

// ReadKey reads a JSON object key (a quoted string) at data[i] without
// allocating. Keys are assumed not to contain backslash escapes; if they do,
// the slow path is taken.
func ReadKey(data []byte, i int) (string, int, error) {
	if i >= len(data) || data[i] != '"' {
		return "", i, ErrInvalidJSON
	}
	i++
	rest := data[i:]
	k := indexCloseOrEscape(rest)
	if k == len(rest) {
		return "", len(data), ErrTruncated
	}
	if rest[k] == '\\' {
		return decodeStringEscaped(data, i, i+k)
	}
	return unsafeStr(rest[:k]), i + k + 1, nil
}

// ReadStringOrNull reads a JSON string (or null) at data[i], copying the bytes
// into a fresh string.
func ReadStringOrNull(data []byte, i int) (string, int, error) {
	if i >= len(data) {
		return "", i, ErrTruncated
	}
	if data[i] == 'n' {
		end, err := ExpectNull(data, i)
		return "", end, err
	}
	if data[i] != '"' {
		return "", i, ErrInvalidJSON
	}
	i++
	rest := data[i:]
	k := indexCloseOrEscape(rest)
	if k == len(rest) {
		return "", len(data), ErrTruncated
	}
	if rest[k] == '\\' {
		return decodeStringEscaped(data, i, i+k)
	}
	return string(rest[:k]), i + k + 1, nil
}

// ReadStringNoCopyOrNull is like ReadStringOrNull but, for strings without
// escapes, returns a string that aliases data rather than copying it, so the
// caller must keep data unchanged while the string is in use. Strings
// containing escapes still allocate, since they cannot be represented as a
// slice of the input.
func ReadStringNoCopyOrNull(data []byte, i int) (string, int, error) {
	if i >= len(data) {
		return "", i, ErrTruncated
	}
	if data[i] == 'n' {
		end, err := ExpectNull(data, i)
		return "", end, err
	}
	if data[i] != '"' {
		return "", i, ErrInvalidJSON
	}
	i++
	rest := data[i:]
	k := indexCloseOrEscape(rest)
	if k == len(rest) {
		return "", len(data), ErrTruncated
	}
	if rest[k] == '\\' {
		return decodeStringEscaped(data, i, i+k)
	}
	return unsafeStr(rest[:k]), i + k + 1, nil
}

// ReadStringDestructiveOrNull is ReadStringNoCopyOrNull but, for a string that
// contains escapes, unescapes it *in place* — overwriting the escaped bytes of
// data with the decoded ones — instead of allocating a scratch buffer, and aliases
// the result. The unescaped form is never longer than the escaped body, so it fits
// within the body's bytes; the rest of the body is left as overwritten garbage and
// the closing quote (which the write cursor never reaches) still bounds the value.
// This DESTROYS the input document: the bytes of every escaped string are clobbered
// and any other alias into the same region (an overlapping nocopy value) is
// invalidated. It is the //lightning:destructive counterpart of the nocopy reader,
// for callers that own the buffer and discard it after decoding. Escape-free strings
// alias the input unchanged, exactly like the nocopy reader.
func ReadStringDestructiveOrNull(data []byte, i int) (string, int, error) {
	if i >= len(data) {
		return "", i, ErrTruncated
	}
	if data[i] == 'n' {
		end, err := ExpectNull(data, i)
		return "", end, err
	}
	if data[i] != '"' {
		return "", i, ErrInvalidJSON
	}
	i++
	rest := data[i:]
	k := indexCloseOrEscape(rest)
	if k == len(rest) {
		return "", len(data), ErrTruncated
	}
	if rest[k] == '\\' {
		// Decode into data starting at the body offset i: buf aliases data[i:] with
		// length 0 and cap to the document end, so decodeEscaped's appends write
		// through into data. The write cursor trails the read cursor (unescaping only
		// shrinks), so it never clobbers a byte not yet consumed, and cap is large
		// enough that append never reallocates away from data.
		buf := data[i:i:len(data)]
		return decodeEscaped(buf, data, i, i+k, true)
	}
	return unsafeStr(rest[:k]), i + k + 1, nil
}

// ReadInt64OrNull reads a JSON integer (or null) at data[i]. Fractional and
// exponent parts are tolerated and truncated toward zero.
func ReadInt64OrNull(data []byte, i int) (int64, int, error) {
	if i >= len(data) {
		return 0, i, ErrTruncated
	}
	if data[i] == 'n' {
		end, err := ExpectNull(data, i)
		return 0, end, err
	}
	neg := false
	if data[i] == '-' {
		neg = true
		i++
		if i >= len(data) {
			return 0, i, ErrBadNumber
		}
	}
	if data[i]-'0' > 9 { // unsigned: a byte below '0' wraps high, so one compare
		return 0, i, ErrBadNumber
	}
	var n int64
	// Fold four digits per SWAR chunk for long integers (database/timestamp IDs
	// run 9-18 digits), then a scalar tail for the last 1-3. n*10000+v matches the
	// scalar n*10+d chain bit for bit, including the wrap past 2^63 on overflow.
	for i+4 <= len(data) {
		w := binary.LittleEndian.Uint32(data[i : i+4])
		if !is4Digits(w) {
			break
		}
		n = n*10000 + int64(parse4Digits(w))
		i += 4
	}
	for i < len(data) {
		d := data[i] - '0'
		if d > 9 {
			break
		}
		n = n*10 + int64(d)
		i++
	}
	if i < len(data) && (data[i] == '.' || data[i] == 'e' || data[i] == 'E') {
		for i < len(data) {
			c := data[i]
			if (c >= '0' && c <= '9') || c == '.' || c == 'e' || c == 'E' || c == '+' || c == '-' {
				i++
				continue
			}
			break
		}
	}
	if neg {
		n = -n
	}
	return n, i, nil
}

// ReadUint64OrNull reads a JSON unsigned integer (or null) at data[i].
func ReadUint64OrNull(data []byte, i int) (uint64, int, error) {
	if i >= len(data) {
		return 0, i, ErrTruncated
	}
	if data[i] == 'n' {
		end, err := ExpectNull(data, i)
		return 0, end, err
	}
	if data[i]-'0' > 9 { // unsigned: a byte below '0' wraps high, so one compare
		return 0, i, ErrBadNumber
	}
	var n uint64
	// Fold four digits per SWAR chunk (see ReadInt64OrNull); the scalar tail picks
	// up the last 1-3 digits.
	for i+4 <= len(data) {
		w := binary.LittleEndian.Uint32(data[i : i+4])
		if !is4Digits(w) {
			break
		}
		n = n*10000 + uint64(parse4Digits(w))
		i += 4
	}
	for i < len(data) {
		d := data[i] - '0'
		if d > 9 {
			break
		}
		n = n*10 + uint64(d)
		i++
	}
	if i < len(data) && (data[i] == '.' || data[i] == 'e' || data[i] == 'E') {
		for i < len(data) {
			c := data[i]
			if (c >= '0' && c <= '9') || c == '.' || c == 'e' || c == 'E' || c == '+' || c == '-' {
				i++
				continue
			}
			break
		}
	}
	return n, i, nil
}

// ReadNumberOrNull reads a JSON number (or null) at data[i] and returns its raw
// literal as a string — the bytes a json.Number holds. The token bounds come from
// the same scanner SkipValue uses, so the literal is captured verbatim without
// parsing it to a float; a JSON null yields the empty string.
func ReadNumberOrNull(data []byte, i int) (string, int, error) {
	if i >= len(data) {
		return "", i, ErrTruncated
	}
	if data[i] == 'n' {
		end, err := ExpectNull(data, i)
		return "", end, err
	}
	start := i
	end, err := skipNumber(data, i)
	if err != nil {
		return "", end, err
	}
	return string(data[start:end]), end, nil
}

// ReadNumberNoCopyOrNull is ReadNumberOrNull but returns a string that aliases
// data instead of copying it, so the caller must keep data unchanged while the
// result is in use.
func ReadNumberNoCopyOrNull(data []byte, i int) (string, int, error) {
	if i >= len(data) {
		return "", i, ErrTruncated
	}
	if data[i] == 'n' {
		end, err := ExpectNull(data, i)
		return "", end, err
	}
	start := i
	end, err := skipNumber(data, i)
	if err != nil {
		return "", end, err
	}
	return unsafeStr(data[start:end]), end, nil
}

// ReadFloat64OrNull reads a JSON number (or null) at data[i] as a float64.
func ReadFloat64OrNull(data []byte, i int) (float64, int, error) {
	if i >= len(data) {
		return 0, i, ErrTruncated
	}
	if data[i] == 'n' {
		end, err := ExpectNull(data, i)
		return 0, end, err
	}
	f, end, fast, ok := scanFloat(data, i)
	if !ok {
		return 0, end, ErrBadNumber
	}
	if fast {
		return f, end, nil
	}
	// scanFloat already tried Eisel–Lemire on the extracted mantissa/exp10; reaching
	// here means it declined (e.g. >19 significant digits, ambiguous rounding,
	// subnormal/overflow, or exponent outside the table), so defer to strconv.
	// unsafeStr avoids copying the token; ParseFloat does not retain it.
	f, perr := strconv.ParseFloat(unsafeStr(data[i:end]), 64)
	if perr != nil {
		return 0, end, ErrBadNumber
	}
	return f, end, nil
}

// ReadBoolOrNull reads a JSON boolean (or null) at data[i].
func ReadBoolOrNull(data []byte, i int) (bool, int, error) {
	if i >= len(data) {
		return false, i, ErrTruncated
	}
	switch data[i] {
	case 'n':
		end, err := ExpectNull(data, i)
		return false, end, err
	case 't':
		if i+4 <= len(data) && data[i+1] == 'r' && data[i+2] == 'u' && data[i+3] == 'e' {
			return true, i + 4, nil
		}
		return false, i, ErrInvalidJSON
	case 'f':
		if i+5 <= len(data) && data[i+1] == 'a' && data[i+2] == 'l' && data[i+3] == 's' && data[i+4] == 'e' {
			return false, i + 5, nil
		}
		return false, i, ErrInvalidJSON
	default:
		return false, i, ErrInvalidJSON
	}
}

// ReadTimeOrNull reads an RFC 3339 JSON string (or null) at data[i] as a
// time.Time, matching how encoding/json decodes time.Time. The intermediate
// string aliases data (time.Parse does not retain it), so this allocates only
// the time.Time itself.
func ReadTimeOrNull(data []byte, i int) (time.Time, int, error) {
	if i >= len(data) {
		return time.Time{}, i, ErrTruncated
	}
	if data[i] == 'n' {
		end, err := ExpectNull(data, i)
		return time.Time{}, end, err
	}
	s, end, err := ReadStringNoCopyOrNull(data, i)
	if err != nil {
		return time.Time{}, end, err
	}
	if t, ok := parseRFC3339(s, false); ok {
		return t, end, nil
	}
	t, perr := time.Parse(time.RFC3339, s)
	if perr != nil {
		return time.Time{}, end, perr
	}
	return t, end, nil
}

// ReadTimeLaxOrNull reads a time.Time at data[i], accepting more shapes than
// ReadTimeOrNull: an RFC 3339 string with either a 'T' or a space separator and
// optional fractional seconds, or a Unix timestamp (in seconds, milliseconds, or
// microseconds) given as a JSON number or a numeric string. The result is
// normalized to UTC. Anything it cannot interpret returns ErrBadTime, which the
// "lax" decode path turns into a skipped value and an unset field.
func ReadTimeLaxOrNull(data []byte, i int) (time.Time, int, error) {
	if i >= len(data) {
		return time.Time{}, i, ErrTruncated
	}
	switch data[i] {
	case 'n':
		end, err := ExpectNull(data, i)
		return time.Time{}, end, err
	case '"':
		s, end, err := ReadStringNoCopyOrNull(data, i)
		if err != nil {
			return time.Time{}, end, err
		}
		if t, ok := parseJSONTS(s); ok {
			return t, end, nil
		}
		if t, ok := parseNumTS(s); ok { // numeric timestamp encoded as a string
			return t, end, nil
		}
		return time.Time{}, end, ErrBadTime
	default:
		end, err := skipNumber(data, i)
		if err != nil {
			return time.Time{}, end, err
		}
		if t, ok := parseNumTS(unsafeStr(data[i:end])); ok {
			return t, end, nil
		}
		return time.Time{}, end, ErrBadTime
	}
}
