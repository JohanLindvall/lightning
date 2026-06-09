// Package support provides the JSON scanning primitives shared by the
// unmarshalers produced by the lightning code generator (see ../../program).
//
// Generated decoders call the exported Read*/Skip*/Decode*/ExpectNull
// functions and the exported Err* sentinels; the unexported helpers are
// internal to this package. The scanner is index based and avoids allocation
// on the common paths (unescaped strings, integers, object keys).
package support

import (
	"bytes"
	"errors"
	"strconv"
	"time"
	"unicode/utf16"
	"unicode/utf8"
	"unsafe"
)

// Errors returned by the scanner and by generated decoders.
var (
	ErrInvalidJSON  = errors.New("json: invalid JSON")
	ErrTruncated    = errors.New("json: truncated JSON")
	ErrBadEscape    = errors.New("json: invalid string escape")
	ErrBadUnicode   = errors.New("json: invalid unicode escape")
	ErrBadNumber    = errors.New("json: invalid number")
	ErrExpectColon  = errors.New("json: expected ':'")
	ErrExpectObject = errors.New("json: expected '{'")
	ErrExpectArray  = errors.New("json: expected '['")
	ErrBadTime      = errors.New("json: invalid time")
)

// rfc3339NanoSpaceLayout is RFC 3339 with the date and time separated by a space
// instead of a 'T', a common variant emitted by databases and log pipelines.
const rfc3339NanoSpaceLayout = "2006-01-02 15:04:05.999999999Z07:00"

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

// indexCloseOrEscapeScalar is the portable fallback for indexCloseOrEscape: it
// returns the index of the first '"' or '\\' in b, or len(b) if neither is
// present, using the runtime's (already SIMD-optimized) bytes.IndexByte.
func indexCloseOrEscapeScalar(b []byte) int {
	q := bytes.IndexByte(b, '"')
	if q < 0 {
		if bs := bytes.IndexByte(b, '\\'); bs >= 0 {
			return bs
		}
		return len(b)
	}
	if bs := bytes.IndexByte(b[:q], '\\'); bs >= 0 {
		return bs
	}
	return q
}

// indexStructuralScalar is the portable fallback for indexStructural: it returns
// the index of the first '{', '}', '[', ']' or '"' in b, or len(b) if none.
func indexStructuralScalar(b []byte) int {
	for i := 0; i < len(b); i++ {
		switch b[i] {
		case '{', '}', '[', ']', '"':
			return i
		}
	}
	return len(b)
}

// UnescapeString decodes the body of a JSON string (the bytes that sit between
// the surrounding quotes, with escape sequences such as \n, \" and \uXXXX still
// present) into the Go string it represents. When in contains no escapes the
// returned string aliases in directly without copying, so the caller must keep
// in unchanged while the result is in use; when escapes are present a new
// string is allocated. It shares its slow path with the JSON scanner.
func UnescapeString(in []byte) (string, error) {
	k := bytes.IndexByte(in, '\\')
	if k < 0 {
		return unsafeStr(in), nil
	}
	s, _, err := decodeEscaped(make([]byte, 0, len(in)), in, 0, k, false)
	return s, err
}

// UnescapeStringInto decodes the body of a JSON string like UnescapeString,
// but writes the decoded bytes into out instead of allocating its own buffer,
// mirroring the Unescape(in, out) convention of github.com/buger/jsonparser:
//
//   - if in contains no escapes, the returned string aliases in and out is
//     untouched;
//   - otherwise the decode is written into out and the returned string aliases
//     out. When cap(out) >= len(in) this allocates nothing (unescaping never
//     lengthens a string, so the result always fits); a shorter out is grown,
//     which allocates.
//
// Pass out == in (e.g. in[:0]) to decode truly in place: that is safe because
// the write cursor never overtakes the read cursor (each escape consumes at
// least as many bytes as it produces) and append uses an overlap-safe memmove.
// Either way the returned string aliases the buffer it was built in, which the
// caller must keep unchanged while the result is in use; when out aliases in,
// in's original (escaped) bytes are overwritten.
func UnescapeStringInto(in, out []byte) (string, error) {
	k := bytes.IndexByte(in, '\\')
	if k < 0 {
		return unsafeStr(in), nil
	}
	s, _, err := decodeEscaped(out[:0], in, 0, k, false)
	return s, err
}

// decodeStringEscaped is the slow path for quoted strings containing backslash
// escapes, terminated by the closing quote in data. It allocates a new buffer.
func decodeStringEscaped(data []byte, start, i int) (string, int, error) {
	return decodeEscaped(make([]byte, 0, len(data)-start), data, start, i, true)
}

// decodeEscaped decodes a backslash-escaped string body, starting from the
// already-validated prefix data[start:i], appending the decoded bytes to buf.
// With quoted=true it stops at the first unescaped '"' and returns the index
// just past it (the JSON scanner path); with quoted=false the whole of data is
// the body and decoding runs to len(data) (the UnescapeString path). buf may
// alias data (decode in place) provided it starts empty at offset 0.
//
// Literal runs between escapes are copied in bulk: a vectorized scan locates
// the next escape (or, when quoted, the closing quote) so the common case of
// long stretches of ordinary characters costs one memmove rather than a
// per-byte append.
func decodeEscaped(buf, data []byte, start, i int, quoted bool) (string, int, error) {
	buf = append(buf, data[start:i]...)
	for i < len(data) {
		// Bulk-copy the literal run up to the next escape (quoted mode also
		// stops at the closing quote) in a single vectorized pass.
		rest := data[i:]
		var k int
		if quoted {
			k = indexCloseOrEscape(rest)
			if k == len(rest) {
				break // unterminated string
			}
			buf = append(buf, rest[:k]...)
			i += k
			if data[i] == '"' {
				return string(buf), i + 1, nil
			}
			i++ // step over the backslash
		} else {
			j := i
			for j < len(data) && data[j] != '\\' {
				j++
			}
			buf = append(buf, data[i:j]...)
			if j >= len(data) {
				break
			}
			i = j + 1 // step over the backslash
		}
		if i >= len(data) {
			return "", i, ErrTruncated
		}
		switch data[i] {
		case '"':
			buf = append(buf, '"')
			i++
		case '\\':
			buf = append(buf, '\\')
			i++
		case '/':
			buf = append(buf, '/')
			i++
		case 'b':
			buf = append(buf, '\b')
			i++
		case 'f':
			buf = append(buf, '\f')
			i++
		case 'n':
			buf = append(buf, '\n')
			i++
		case 'r':
			buf = append(buf, '\r')
			i++
		case 't':
			buf = append(buf, '\t')
			i++
		case 'u':
			r, ni, err := readUnicodeEscape(data, i+1)
			if err != nil {
				return "", ni, err
			}
			i = ni
			if utf16.IsSurrogate(r) {
				if i+1 < len(data) && data[i] == '\\' && data[i+1] == 'u' {
					r2, ni2, err := readUnicodeEscape(data, i+2)
					if err != nil {
						return "", ni2, err
					}
					if dec := utf16.DecodeRune(r, r2); dec != utf8.RuneError {
						r = dec
						i = ni2
					}
				}
			}
			buf = utf8.AppendRune(buf, r)
		default:
			return "", i, ErrBadEscape
		}
	}
	if quoted {
		return "", i, ErrTruncated
	}
	// buf is freshly allocated and never mutated after this point, so it can be
	// handed out as a string without the extra copy string(buf) would make.
	return unsafeStr(buf), i, nil
}

func readUnicodeEscape(data []byte, i int) (rune, int, error) {
	if i+4 > len(data) {
		return 0, i, ErrTruncated
	}
	var r rune
	for k := 0; k < 4; k++ {
		c := data[i+k]
		switch {
		case c >= '0' && c <= '9':
			r = r<<4 | rune(c-'0')
		case c >= 'a' && c <= 'f':
			r = r<<4 | rune(c-'a'+10)
		case c >= 'A' && c <= 'F':
			r = r<<4 | rune(c-'A'+10)
		default:
			return 0, i, ErrBadUnicode
		}
	}
	return r, i + 4, nil
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
	c := data[i]
	if c < '0' || c > '9' {
		return 0, i, ErrBadNumber
	}
	var n int64
	for i < len(data) {
		c = data[i]
		if c < '0' || c > '9' {
			break
		}
		n = n*10 + int64(c-'0')
		i++
	}
	if i < len(data) && (data[i] == '.' || data[i] == 'e' || data[i] == 'E') {
		for i < len(data) {
			c = data[i]
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
	c := data[i]
	if c < '0' || c > '9' {
		return 0, i, ErrBadNumber
	}
	var n uint64
	for i < len(data) {
		c = data[i]
		if c < '0' || c > '9' {
			break
		}
		n = n*10 + uint64(c-'0')
		i++
	}
	if i < len(data) && (data[i] == '.' || data[i] == 'e' || data[i] == 'E') {
		for i < len(data) {
			c = data[i]
			if (c >= '0' && c <= '9') || c == '.' || c == 'e' || c == 'E' || c == '+' || c == '-' {
				i++
				continue
			}
			break
		}
	}
	return n, i, nil
}

// pow10exact holds the powers of ten that are exactly representable as a
// float64 (10^0 .. 10^22). Used by the fast-path float parser.
var pow10exact = [...]float64{
	1e0, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6, 1e7, 1e8, 1e9, 1e10, 1e11,
	1e12, 1e13, 1e14, 1e15, 1e16, 1e17, 1e18, 1e19, 1e20, 1e21, 1e22,
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
	end, err := skipNumber(data, i)
	if err != nil {
		return 0, end, err
	}
	if f, ok := fastFloat(data[i:end]); ok {
		return f, end, nil
	}
	// unsafeStr avoids copying the token; ParseFloat does not retain it.
	f, perr := strconv.ParseFloat(unsafeStr(data[i:end]), 64)
	if perr != nil {
		return 0, end, ErrBadNumber
	}
	return f, end, nil
}

// ParseFloat parses the JSON number in b as a float64. It takes the same Clinger
// fast path as the scanner — an exact mantissa with a small decimal exponent is
// converted with a single multiply or divide — and falls back to
// strconv.ParseFloat for everything else. b must be exactly one number with no
// surrounding whitespace; trailing bytes or an empty input yield ErrBadNumber.
func ParseFloat(b []byte) (float64, error) {
	if f, ok := fastFloat(b); ok {
		return f, nil
	}
	// unsafeStr avoids copying the token; ParseFloat does not retain it.
	f, err := strconv.ParseFloat(unsafeStr(b), 64)
	if err != nil {
		return 0, ErrBadNumber
	}
	return f, nil
}

// fastFloat parses a JSON number using the Clinger fast path: when the decimal
// mantissa fits exactly in a float64 (< 2^53) and the decimal exponent is in
// [-22, 22], a single float multiply or divide by an exact power of ten yields
// the correctly rounded result. It returns ok=false for anything outside that
// range (too many digits, large exponent, malformed token), which the caller
// hands to strconv.ParseFloat. b is the already-validated numeric token.
func fastFloat(b []byte) (float64, bool) {
	n := len(b)
	if n == 0 {
		return 0, false
	}
	i := 0
	neg := false
	switch b[0] {
	case '-':
		neg = true
		i++
	case '+':
		i++
	}

	var mant uint64
	digits := 0
	for i < n && b[i] >= '0' && b[i] <= '9' {
		mant = mant*10 + uint64(b[i]-'0')
		i++
		digits++
	}
	exp := 0
	if i < n && b[i] == '.' {
		i++
		for i < n && b[i] >= '0' && b[i] <= '9' {
			mant = mant*10 + uint64(b[i]-'0')
			i++
			digits++
			exp--
		}
	}
	if digits == 0 || digits > 19 { // >19 digits may overflow the uint64 mantissa
		return 0, false
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
		if i >= n || b[i] < '0' || b[i] > '9' {
			return 0, false
		}
		eval := 0
		for i < n && b[i] >= '0' && b[i] <= '9' {
			eval = eval*10 + int(b[i]-'0')
			if eval > 1000 {
				return 0, false // huge exponent; let strconv handle over/underflow
			}
			i++
		}
		exp += esign * eval
	}
	if i != n || mant>>53 != 0 {
		return 0, false
	}

	f := float64(mant)
	switch {
	case exp == 0:
		// already exact
	case exp > 0 && exp <= 22:
		f *= pow10exact[exp]
	case exp < 0 && exp >= -22:
		f /= pow10exact[-exp]
	default:
		return 0, false
	}
	if neg {
		f = -f
	}
	return f, true
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

// parseJSONTS parses an RFC 3339 timestamp, tolerating a space in place of the
// 'T' date/time separator, and normalizes the result to UTC.
func parseJSONTS(ts string) (time.Time, bool) {
	pattern := time.RFC3339Nano
	if len(ts) > 11 && ts[10] == ' ' {
		pattern = rfc3339NanoSpaceLayout
	}
	t, err := time.Parse(pattern, ts)
	if err == nil {
		t = t.UTC()
	}
	return t, err == nil
}

// parseNumTS parses a Unix timestamp expressed as a decimal integer, inferring
// the unit (seconds, milliseconds, or microseconds) from its magnitude. Values
// outside the recognized ranges are rejected.
func parseNumTS(ts string) (time.Time, bool) {
	val, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return time.Time{}, false
	}
	if val > 1_000_000_000_000_000 && val < 10_000_000_000_000_000 {
		return time.UnixMicro(val).UTC(), true
	}
	if val > 1_000_000_000_000 && val < 10_000_000_000_000 {
		return time.UnixMilli(val).UTC(), true
	}
	if val > 1_000_000_000 && val < 10_000_000_000 {
		return time.Unix(val, 0).UTC(), true
	}
	return time.Time{}, false
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

// ExpectNull consumes the literal null at data[i].
func ExpectNull(data []byte, i int) (int, error) {
	if i+4 > len(data) || data[i] != 'n' || data[i+1] != 'u' || data[i+2] != 'l' || data[i+3] != 'l' {
		return i, ErrInvalidJSON
	}
	return i + 4, nil
}

// SkipValue advances past any JSON value starting at data[i].
func SkipValue(data []byte, i int) (int, error) {
	if i >= len(data) {
		return i, ErrTruncated
	}
	switch data[i] {
	case '"':
		return skipString(data, i)
	case '{':
		return skipObject(data, i)
	case '[':
		return skipArray(data, i)
	case 't':
		if i+4 <= len(data) && data[i+1] == 'r' && data[i+2] == 'u' && data[i+3] == 'e' {
			return i + 4, nil
		}
		return i, ErrInvalidJSON
	case 'f':
		if i+5 <= len(data) && data[i+1] == 'a' && data[i+2] == 'l' && data[i+3] == 's' && data[i+4] == 'e' {
			return i + 5, nil
		}
		return i, ErrInvalidJSON
	case 'n':
		return ExpectNull(data, i)
	default:
		return skipNumber(data, i)
	}
}

func skipString(data []byte, i int) (int, error) {
	// data[i] == '"'
	i++
	for {
		rest := data[i:]
		k := indexCloseOrEscape(rest)
		if k == len(rest) {
			return len(data), ErrTruncated
		}
		if rest[k] == '"' {
			return i + k + 1, nil
		}
		// Skip the escape sequence. For \uXXXX we only need to skip the
		// backslash and the next char; subsequent bytes are processed on the
		// next iteration.
		i += k + 2
		if i > len(data) {
			return len(data), ErrTruncated
		}
	}
}

func skipNumber(data []byte, i int) (int, error) {
	start := i
	if i < len(data) && data[i] == '-' {
		i++
	}
	for i < len(data) {
		c := data[i]
		if (c >= '0' && c <= '9') || c == '.' || c == 'e' || c == 'E' || c == '+' || c == '-' {
			i++
			continue
		}
		break
	}
	if i == start {
		return i, ErrBadNumber
	}
	return i, nil
}

func skipObject(data []byte, i int) (int, error) {
	// data[i] == '{'
	i++
	for i < len(data) {
		// Jump to the next structural byte, skipping inert content (keys' inner
		// chars, numbers, bools, whitespace) in one vectorized pass.
		i += indexStructural(data[i:])
		if i >= len(data) {
			break
		}
		switch data[i] {
		case '}':
			return i + 1, nil
		case '{':
			end, err := skipObject(data, i)
			if err != nil {
				return end, err
			}
			i = end
		case '[':
			end, err := skipArray(data, i)
			if err != nil {
				return end, err
			}
			i = end
		case '"':
			end, err := skipString(data, i)
			if err != nil {
				return end, err
			}
			i = end
		default: // stray ']' (only on malformed input); step over it
			i++
		}
	}
	return i, ErrTruncated
}

func skipArray(data []byte, i int) (int, error) {
	// data[i] == '['
	i++
	for i < len(data) {
		i += indexStructural(data[i:])
		if i >= len(data) {
			break
		}
		switch data[i] {
		case ']':
			return i + 1, nil
		case '[':
			end, err := skipArray(data, i)
			if err != nil {
				return end, err
			}
			i = end
		case '{':
			end, err := skipObject(data, i)
			if err != nil {
				return end, err
			}
			i = end
		case '"':
			end, err := skipString(data, i)
			if err != nil {
				return end, err
			}
			i = end
		default: // stray '}' (only on malformed input); step over it
			i++
		}
	}
	return i, ErrTruncated
}

// SkipWS advances past JSON whitespace at data[i].
func SkipWS(data []byte, i int) int {
	for i < len(data) {
		c := data[i]
		if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
			i++
			continue
		}
		return i
	}
	return i
}

func unsafeStr(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return unsafe.String(unsafe.SliceData(b), len(b))
}

// DecodeValue decodes an arbitrary JSON value at data[i] into the standard Go
// representation (nil, bool, float64, string, []interface{},
// map[string]interface{}).
func DecodeValue(data []byte, i int) (interface{}, int, error) {
	if i >= len(data) {
		return nil, i, ErrTruncated
	}
	switch data[i] {
	case '"':
		s, end, err := ReadStringOrNull(data, i)
		return s, end, err
	case '{':
		return decodeAnyObject(data, i)
	case '[':
		return decodeAnyArray(data, i)
	case 't', 'f':
		b, end, err := ReadBoolOrNull(data, i)
		return b, end, err
	case 'n':
		end, err := ExpectNull(data, i)
		return nil, end, err
	default:
		f, end, err := ReadFloat64OrNull(data, i)
		return f, end, err
	}
}

func decodeAnyObject(data []byte, i int) (interface{}, int, error) {
	// data[i] == '{'
	i++
	m := map[string]interface{}{}
	for {
		i = SkipWS(data, i)
		if i >= len(data) {
			return nil, i, ErrTruncated
		}
		if data[i] == '}' {
			return m, i + 1, nil
		}
		key, ni, err := ReadKey(data, i)
		if err != nil {
			return nil, ni, err
		}
		i = SkipWS(data, ni)
		if i >= len(data) || data[i] != ':' {
			return nil, i, ErrExpectColon
		}
		i = SkipWS(data, i+1)
		val, end, err := DecodeValue(data, i)
		if err != nil {
			return nil, end, err
		}
		m[string([]byte(key))] = val
		i = SkipWS(data, end)
		if i >= len(data) {
			return nil, i, ErrTruncated
		}
		if data[i] == '}' {
			return m, i + 1, nil
		}
		if data[i] != ',' {
			return nil, i, ErrInvalidJSON
		}
		i++
	}
}

func decodeAnyArray(data []byte, i int) (interface{}, int, error) {
	// data[i] == '['
	i++
	a := []interface{}{}
	for {
		i = SkipWS(data, i)
		if i >= len(data) {
			return nil, i, ErrTruncated
		}
		if data[i] == ']' {
			return a, i + 1, nil
		}
		val, end, err := DecodeValue(data, i)
		if err != nil {
			return nil, end, err
		}
		a = append(a, val)
		i = SkipWS(data, end)
		if i >= len(data) {
			return nil, i, ErrTruncated
		}
		if data[i] == ']' {
			return a, i + 1, nil
		}
		if data[i] != ',' {
			return nil, i, ErrInvalidJSON
		}
		i++
	}
}
