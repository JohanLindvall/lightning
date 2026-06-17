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
	"encoding/base64"
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
	ErrKeyNotFound  = errors.New("json: key path not found")
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

// Unwrap reads the JSON string at data[i] and returns the JSON document embedded
// in it — the backing of the "unwrap" field option. The string body is
// unescaped; if it is not itself JSON (its first non-whitespace byte is not the
// start of a JSON value) it is base64-decoded first, standard alphabet, with or
// without padding. A JSON null yields a nil document and no error, and the second
// return value is the offset just past the consumed string. The returned slice
// is freshly allocated and not retained by Unwrap, so values decoded out of it
// may safely alias it (the "nocopy" option).
func Unwrap(data []byte, i int) ([]byte, int, error) {
	if i >= len(data) {
		return nil, i, ErrTruncated
	}
	if data[i] == 'n' {
		end, err := ExpectNull(data, i)
		return nil, end, err
	}
	s, end, err := ReadStringOrNull(data, i)
	if err != nil {
		return nil, end, err
	}
	body := []byte(s)
	if startsJSON(body) {
		return body, end, nil
	}
	if dec, ok := decodeBase64(body); ok {
		return dec, end, nil
	}
	// Neither plainly JSON nor valid base64; hand the body back so the caller's
	// decode reports a meaningful error on it.
	return body, end, nil
}

// startsJSON reports whether b begins, after any leading whitespace, with the
// first byte of a JSON value ('{', '[', '"', '-', a digit, or 't'/'f'/'n' for
// true/false/null). Unwrap uses it to tell an embedded JSON document apart from
// a base64-encoded one, and for that the test is exact: the base64 of a JSON
// document can never start with one of these bytes.
//
// A base64 string's first character is the top six bits of the first input byte
// (byte0 >> 2). A JSON value always starts with an ASCII byte (< 0x80), so that
// index is <= 31, which base64 always renders as a letter in A-Z or a-f. Of the
// bytes startsJSON accepts, base64 cannot emit '{', '[', '"' or '-' at all, and
// the digits 0-9 would need byte0 >= 0xD0; the sole overlap is 'f' (index 31),
// which needs byte0 in 0x7C..0x7F ('|', '}', '~', DEL) — none of which can begin
// a JSON value. So no base64-encoded JSON document ever trips this test.
//
// Plenty of *other* base64 strings do start with a digit or t/f/n and so return
// true here, but those decode to a non-ASCII leading byte (not JSON), which is
// outside Unwrap's "the body is JSON or base64-encoded JSON" contract; on such
// input Unwrap still degrades gracefully by handing the body back for the
// caller's decode to reject.
func startsJSON(b []byte) bool {
	for _, c := range b {
		switch c {
		case ' ', '\t', '\n', '\r':
			continue
		case '{', '[', '"', '-', 't', 'f', 'n',
			'0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return true
		default:
			return false
		}
	}
	return false
}

// decodeBase64 decodes b as standard-alphabet base64 and reports whether it
// succeeded. Trailing '=' padding is stripped first so a single unpadded decode
// (base64.RawStdEncoding) covers both the padded and unpadded forms — one
// allocation and one pass, rather than attempting each encoding in turn.
func decodeBase64(b []byte) ([]byte, bool) {
	for len(b) > 0 && b[len(b)-1] == '=' {
		b = b[:len(b)-1]
	}
	dst := make([]byte, base64.RawStdEncoding.DecodedLen(len(b)))
	n, err := base64.RawStdEncoding.Decode(dst, b)
	if err != nil {
		return nil, false
	}
	return dst[:n], true
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
// escapes, terminated by the closing quote in data. It allocates a new buffer
// sized to the string body: the closing quote is located up front (data[start-1]
// is the opening quote) and, since unescaping never lengthens a string, the
// decoded result fits in body-length bytes with a single allocation. Sizing to
// len(data)-start instead over-allocated by the whole remaining document, which
// is catastrophic for an escaped string near the start of a large input (e.g.
// the many "\/"-escaped URLs in a multi-hundred-KB Twitter payload).
func decodeStringEscaped(data []byte, start, i int) (string, int, error) {
	capHint := i - start // fallback: the escape-free prefix; append grows the rest
	if end, err := skipString(data, start-1); err == nil {
		capHint = end - 1 - start // end is past the closing quote
	}
	return decodeEscaped(make([]byte, 0, capHint), data, start, i, true)
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

// scanFloat scans the JSON number token beginning at data[i] in a single pass,
// returning the index just past it and, when it can resolve the value exactly,
// the parsed result with fast=true. Two fast paths run on the mantissa and
// exponent the scan extracts, so no rescan is needed: the Clinger path (an
// exactly representable mantissa < 2^53 with a decimal exponent in [-22, 22],
// converted with one multiply or divide) and, for an exact mantissa that misses
// Clinger, Eisel-Lemire (a 128-bit multiply — see eiselLemire64). Only a mantissa
// of more than 19 significant digits or a value Eisel-Lemire finds ambiguous
// returns fast=false, leaving the caller to parse data[i:end] with
// strconv.ParseFloat. ok=false marks a token with no digits (or an exponent
// marker with no digits), in which case end points at the offending byte.
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

// parseRFC3339 parses an RFC 3339 timestamp on the fast path shared by the time
// decoders, returning the same instant as time.Parse(time.RFC3339Nano, s). It
// returns ok=false — so the caller falls back to time.Parse — for any shape it
// does not handle exactly: a wrong-length or malformed token, an out-of-range
// field, or a timezone that is neither 'Z' nor ±HH:MM. allowSpace additionally
// permits a space in place of the 'T' date/time separator (the lax variant).
//
// It sidesteps time.Parse's reference-layout machinery, which dominates decoding
// of timestamp-heavy documents, by reading the fixed-offset fields directly.
func parseRFC3339(s string, allowSpace bool) (time.Time, bool) {
	// Shortest accepted form is "2006-01-02T15:04:05Z" (20 bytes).
	if len(s) < 20 {
		return time.Time{}, false
	}
	if sep := s[10]; sep != 'T' && !(allowSpace && sep == ' ') {
		return time.Time{}, false
	}
	if s[4] != '-' || s[7] != '-' || s[13] != ':' || s[16] != ':' {
		return time.Time{}, false
	}
	year, ok0 := atoi4(s, 0)
	month, ok1 := atoi2(s, 5)
	day, ok2 := atoi2(s, 8)
	hour, ok3 := atoi2(s, 11)
	min, ok4 := atoi2(s, 14)
	sec, ok5 := atoi2(s, 17)
	if !(ok0 && ok1 && ok2 && ok3 && ok4 && ok5) {
		return time.Time{}, false
	}
	// Reject out-of-range fields so leap seconds and the like defer to time.Parse.
	if month < 1 || month > 12 || day < 1 || day > 31 || hour > 23 || min > 59 || sec > 59 {
		return time.Time{}, false
	}

	i := 19
	nsec := 0
	if s[i] == '.' {
		i++
		fd := 0
		for i < len(s) && s[i] >= '0' && s[i] <= '9' {
			if fd < 9 { // nanosecond precision; extra digits are ignored
				nsec = nsec*10 + int(s[i]-'0')
				fd++
			}
			i++
		}
		if fd == 0 {
			return time.Time{}, false
		}
		for ; fd < 9; fd++ {
			nsec *= 10
		}
	}

	if i >= len(s) {
		return time.Time{}, false
	}
	var loc *time.Location
	switch s[i] {
	case 'Z':
		if i+1 != len(s) {
			return time.Time{}, false
		}
		loc = time.UTC
	case '+', '-':
		if i+6 != len(s) || s[i+3] != ':' {
			return time.Time{}, false
		}
		oh, oka := atoi2(s, i+1)
		om, okb := atoi2(s, i+4)
		if !oka || !okb || oh > 23 || om > 59 {
			return time.Time{}, false
		}
		off := (oh*60 + om) * 60
		if s[i] == '-' {
			off = -off
		}
		if off == 0 {
			loc = time.UTC
		} else {
			loc = time.FixedZone("", off)
		}
	default:
		return time.Time{}, false
	}
	if loc == time.UTC {
		// The fields are already range-checked, so build the instant directly and
		// skip time.Date's field-normalization machinery (dateToAbsDays and the
		// surrounding range handling were ~18% of timestamp-array decoding).
		secs := daysFromCivil(year, month, day)*86400 + int64(hour)*3600 + int64(min)*60 + int64(sec)
		return time.Unix(secs, int64(nsec)).UTC(), true
	}
	return time.Date(year, time.Month(month), day, hour, min, sec, nsec, loc), true
}

// daysFromCivil returns the number of days from 1970-01-01 to the given
// proleptic-Gregorian date using Howard Hinnant's algorithm, which is branch-
// light and valid across time.Time's full year range. month is 1-12 and day
// 1-31; the caller has already range-checked them. It lets the RFC 3339 fast
// path construct UTC timestamps via time.Unix.
func daysFromCivil(y, m, d int) int64 {
	yy := int64(y)
	if m <= 2 {
		yy-- // treat Jan/Feb as months 13/14 of the previous year
	}
	var era int64
	if yy >= 0 {
		era = yy / 400
	} else {
		era = (yy - 399) / 400
	}
	yoe := yy - era*400 // year of era, [0, 399]
	mm := int64(m)
	mp := mm - 3 // March-based month index
	if mm <= 2 {
		mp = mm + 9
	}
	doy := (153*mp+2)/5 + int64(d) - 1      // day of year, [0, 365]
	doe := yoe*365 + yoe/4 - yoe/100 + doy  // day of era, [0, 146096]
	return era*146097 + doe - 719468        // 719468 = days from 0000-03-01 to 1970-01-01
}

// atoi2 parses exactly two decimal digits of s at off into a non-negative int,
// returning false if either byte is not an ASCII digit. The branchless digit
// test and the absence of a loop (n is fixed) make it noticeably cheaper than a
// general atoiN on the RFC 3339 fast path, which parses six such fields per
// timestamp. The caller guarantees off+2 <= len(s).
func atoi2(s string, off int) (int, bool) {
	d0 := s[off] - '0'
	d1 := s[off+1] - '0'
	return int(d0)*10 + int(d1), d0 <= 9 && d1 <= 9
}

// atoi4 parses exactly four decimal digits of s at off (the RFC 3339 year).
// The caller guarantees off+4 <= len(s).
func atoi4(s string, off int) (int, bool) {
	d0 := s[off] - '0'
	d1 := s[off+1] - '0'
	d2 := s[off+2] - '0'
	d3 := s[off+3] - '0'
	return int(d0)*1000 + int(d1)*100 + int(d2)*10 + int(d3), d0 <= 9 && d1 <= 9 && d2 <= 9 && d3 <= 9
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
	if t, ok := parseRFC3339(ts, true); ok {
		return t.UTC(), true
	}
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

// CountArrayElements returns the number of top-level elements in the JSON array
// beginning at data[i] (data[i] must be '['), so a destination slice can be
// allocated once instead of grown by repeated append. It returns 0 for an empty
// array or when the count cannot be determined cheaply (a truncated or malformed
// array); the caller then simply falls back to append-driven growth, so an
// imperfect count is only ever a missed optimization, never a correctness
// problem.
//
// Each element is skipped whole with SkipValue rather than walked byte by byte:
// SkipValue uses the SIMD indexStructural scanner for nested arrays/objects and
// indexCloseOrEscape for strings, so a structurally dense element — a nested
// coordinate array of many numbers, say — is jumped over in vectorized strides
// instead of one byte at a time. This is what makes presizing slices of arrays,
// objects, or strings cheap.
func CountArrayElements(data []byte, i int) int {
	if i >= len(data) || data[i] != '[' {
		return 0
	}
	i = SkipWS(data, i+1)
	if i >= len(data) || data[i] == ']' {
		return 0
	}
	n := 1
	for {
		end, err := SkipValue(data, i)
		if err != nil {
			return 0
		}
		i = SkipWS(data, end)
		if i >= len(data) {
			return 0
		}
		switch data[i] {
		case ']':
			return n
		case ',':
			n++
			i = SkipWS(data, i+1)
			if i >= len(data) {
				return 0
			}
		default:
			return 0 // malformed; let the caller grow on demand
		}
	}
}

// CountArrayScalars counts the elements of a JSON array of scalar values
// (numbers, booleans, or null) beginning at data[i] (data[i] must be '['). Such
// elements never contain a quote, comma, or bracket, so the closing ']' is the
// first ']' in the input and the element count is one more than the number of
// commas before it — both found with vectorized byte scans. It is therefore much
// cheaper than CountArrayElements but valid only when the element type is known
// (by the generator) to be a scalar. It returns 0 for an empty array.
func CountArrayScalars(data []byte, i int) int {
	rb := bytes.IndexByte(data[i+1:], ']')
	if rb < 0 {
		return 0
	}
	seg := data[i+1 : i+1+rb]
	for _, c := range seg {
		if c != ' ' && c != '\t' && c != '\n' && c != '\r' {
			return bytes.Count(seg, commaByte) + 1
		}
	}
	return 0 // empty (whitespace-only) array
}

var commaByte = []byte{','}

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

// isWS is a lookup table of the four JSON whitespace bytes (space, tab, newline,
// carriage return). A single indexed load replaces the four-way comparison
// chain; since the index is a byte it is provably in range, so the load carries
// no bounds check. This is the hottest classification in the scanner — SkipWS
// runs before and after every value — and the table form is faster on the common
// compact case (the next byte is structural, so the loop exits immediately).
var isWS = [256]bool{' ': true, '\t': true, '\n': true, '\r': true}

// SkipWS advances past JSON whitespace at data[i].
func SkipWS(data []byte, i int) int {
	for i < len(data) && isWS[data[i]] {
		i++
	}
	return i
}

// skipWSc is the inter-token whitespace skip used by the compact-aware Get,
// GetMany and ObjectEach. In compact mode the input is asserted to carry no
// whitespace between tokens (the form compact JSON serializers emit), so it
// returns i unchanged; otherwise it is SkipWS. This mirrors the generator's
// //lightning:compact decoders, which elide exactly these inter-token skips.
func skipWSc(data []byte, i int, compact bool) int {
	if compact {
		return i
	}
	return SkipWS(data, i)
}

func unsafeStr(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return unsafe.String(unsafe.SliceData(b), len(b))
}

// GetMany looks up several top-level members of the JSON object in data at once,
// in a single pass over the object, and returns their raw value bytes in the
// same order as keys. Where Get walks one key path and N separate Get calls
// would rescan the object N times, GetMany scans it once and distributes each
// member to the matching requested key, so extracting a handful of fields from a
// wide record costs one walk rather than one per field.
//
// The results are written into out[:0] — pass a slice to reuse (a nil out
// allocates) — and out is returned with length len(keys): out[n] is the value
// for keys[n], or nil if that key is absent from the object. Each value aliases
// data and follows Get's conventions (quotes kept for strings, the full span for
// objects and arrays, the literal token for scalars); a present member whose
// value is JSON null yields the four bytes "null", distinct from a missing key's
// nil. A duplicate key in the document takes its first occurrence; a key in keys
// not present in the object stays nil without being an error.
//
// A missing key is not an error — that is what the nil slot reports. Any other
// failure (a non-object root, malformed or truncated JSON) returns a non-nil
// error together with out as filled so far. Like Get, GetMany returns as soon as
// every requested key has been found and does not validate the rest of the
// object.
func GetMany(data []byte, keys []string, out [][]byte) ([][]byte, error) {
	return getMany(data, keys, out, false)
}

// GetManyCompact is GetMany for compact JSON — input with no whitespace between
// tokens, as compact serializers emit — skipping the inter-token whitespace
// scans GetMany makes (leading whitespace at the document start is still
// tolerated). On such input it behaves identically to GetMany but faster; given
// input that does contain inter-token whitespace it may report an error.
func GetManyCompact(data []byte, keys []string, out [][]byte) ([][]byte, error) {
	return getMany(data, keys, out, true)
}

func getMany(data []byte, keys []string, out [][]byte, compact bool) ([][]byte, error) {
	out = out[:0]
	for range keys {
		out = append(out, nil)
	}
	i := SkipWS(data, 0)
	if i >= len(data) {
		return out, ErrTruncated
	}
	if data[i] != '{' {
		return out, ErrExpectObject
	}
	i++
	found := 0
	for {
		i = skipWSc(data, i, compact)
		if i >= len(data) {
			return out, ErrTruncated
		}
		if data[i] == '}' {
			return out, nil
		}
		key, ni, err := ReadKey(data, i)
		if err != nil {
			return out, err
		}
		i = skipWSc(data, ni, compact)
		if i >= len(data) || data[i] != ':' {
			return out, ErrExpectColon
		}
		i = skipWSc(data, i+1, compact)
		start := i
		end, err := SkipValue(data, i)
		if err != nil {
			return out, err
		}
		// Distribute this member to every requested key it matches that is still
		// unset, so the first occurrence in the document wins and duplicate
		// entries in keys are all filled. Key sets are small, so the linear scan
		// is cheaper than building a map.
		for n, k := range keys {
			if out[n] == nil && k == key {
				out[n] = data[start:end]
				found++
			}
		}
		if found == len(keys) {
			return out, nil // every requested key found; skip the rest
		}
		i = skipWSc(data, end, compact)
		if i >= len(data) {
			return out, ErrTruncated
		}
		switch data[i] {
		case '}':
			return out, nil
		case ',':
			i++
		default:
			return out, ErrInvalidJSON
		}
	}
}

// Get walks the object-key path keys into the JSON document data and returns
// the raw bytes of the value found at that path, mirroring
// github.com/buger/jsonparser's Get without reporting a value type. The
// returned slice aliases data — for a string it includes the surrounding
// quotes (escapes left intact), for an object or array it spans the matching
// '{'..'}' or '['..']', and for a scalar it is the literal token.
//
// Each key descends one level: at every step the current value must be a JSON
// object that contains the key, or Get returns ErrKeyNotFound (and, for a
// non-object where descent was attempted, the index is left at that value).
// With no keys Get returns the whole value at the document root. The second
// return value is the offset in data at which the returned value begins, and
// leading whitespace is tolerated at every level.
func Get(data []byte, keys ...string) ([]byte, int, error) {
	return get(data, false, keys...)
}

// GetCompact is Get for compact JSON — input with no whitespace between tokens,
// as compact serializers emit — skipping the inter-token whitespace scans Get
// makes while descending the key path (leading whitespace at the document start
// is still tolerated). On such input it behaves identically to Get but faster;
// given input that does contain inter-token whitespace it may report an error.
func GetCompact(data []byte, keys ...string) ([]byte, int, error) {
	return get(data, true, keys...)
}

func get(data []byte, compact bool, keys ...string) ([]byte, int, error) {
	i := SkipWS(data, 0)
	for _, key := range keys {
		var err error
		i, err = objectField(data, i, key, compact)
		if err != nil {
			return nil, i, err
		}
	}
	end, err := SkipValue(data, i)
	if err != nil {
		return nil, i, err
	}
	return data[i:end], i, nil
}

// ObjectEach calls fn once for every member of the JSON object reached by the
// object-key path keys in data, modeled on github.com/buger/jsonparser's
// ObjectEach but without reporting a value type. fn receives the member's
// decoded key and the raw bytes of its value; both alias data (so the caller
// must keep data unchanged while they are in use), and the value follows the
// same conventions as Get — quotes kept for strings, the full span for objects
// and arrays, the literal token for scalars.
//
// With no keys ObjectEach iterates the document's root object; otherwise each
// key descends one level and the value at the end of the path must itself be an
// object (ErrExpectObject if not, ErrKeyNotFound if a key is missing). If fn
// returns a non-nil error, iteration stops and that error is returned.
// Non-target members along the path are skipped without allocating.
func ObjectEach(data []byte, fn func(key string, value []byte) error, keys ...string) error {
	return objectEach(data, fn, false, keys...)
}

// ObjectEachCompact is ObjectEach for compact JSON — input with no whitespace
// between tokens, as compact serializers emit — skipping the inter-token
// whitespace scans ObjectEach makes (leading whitespace at the document start is
// still tolerated). On such input it behaves identically to ObjectEach but
// faster; given input that does contain inter-token whitespace it may report an
// error.
func ObjectEachCompact(data []byte, fn func(key string, value []byte) error, keys ...string) error {
	return objectEach(data, fn, true, keys...)
}

func objectEach(data []byte, fn func(key string, value []byte) error, compact bool, keys ...string) error {
	i := SkipWS(data, 0)
	for _, key := range keys {
		var err error
		i, err = objectField(data, i, key, compact)
		if err != nil {
			return err
		}
	}
	i = skipWSc(data, i, compact)
	if i >= len(data) {
		return ErrTruncated
	}
	if data[i] != '{' {
		return ErrExpectObject
	}
	i++
	for {
		i = skipWSc(data, i, compact)
		if i >= len(data) {
			return ErrTruncated
		}
		if data[i] == '}' {
			return nil
		}
		key, ni, err := ReadKey(data, i)
		if err != nil {
			return err
		}
		i = skipWSc(data, ni, compact)
		if i >= len(data) || data[i] != ':' {
			return ErrExpectColon
		}
		i = skipWSc(data, i+1, compact)
		start := i
		end, err := SkipValue(data, i)
		if err != nil {
			return err
		}
		if err := fn(key, data[start:end]); err != nil {
			return err
		}
		i = skipWSc(data, end, compact)
		if i >= len(data) {
			return ErrTruncated
		}
		switch data[i] {
		case '}':
			return nil
		case ',':
			i++
		default:
			return ErrInvalidJSON
		}
	}
}

// objectField scans the JSON object at data[i] (after any leading whitespace)
// for the member named key and returns the index of its value, with the value's
// own leading whitespace already skipped. It returns ErrExpectObject if data[i]
// is not an object and ErrKeyNotFound if the object has no such key. It reuses
// the scanner primitives (ReadKey, SkipValue, SkipWS) so non-target members are
// skipped without allocating.
func objectField(data []byte, i int, key string, compact bool) (int, error) {
	i = skipWSc(data, i, compact)
	if i >= len(data) {
		return i, ErrTruncated
	}
	if data[i] != '{' {
		return i, ErrExpectObject
	}
	i++
	for {
		i = skipWSc(data, i, compact)
		if i >= len(data) {
			return i, ErrTruncated
		}
		if data[i] == '}' {
			return i, ErrKeyNotFound
		}
		k, ni, err := ReadKey(data, i)
		if err != nil {
			return ni, err
		}
		i = skipWSc(data, ni, compact)
		if i >= len(data) || data[i] != ':' {
			return i, ErrExpectColon
		}
		i = skipWSc(data, i+1, compact)
		if k == key {
			return i, nil
		}
		end, err := SkipValue(data, i)
		if err != nil {
			return end, err
		}
		i = skipWSc(data, end, compact)
		if i >= len(data) {
			return i, ErrTruncated
		}
		switch data[i] {
		case '}':
			return i, ErrKeyNotFound
		case ',':
			i++
		default:
			return i, ErrInvalidJSON
		}
	}
}

// DecodeValue decodes an arbitrary JSON value at data[i] into the standard Go
// representation (nil, bool, float64, string, []any, map[string]any).
func DecodeValue(data []byte, i int) (any, int, error) {
	return decodeValue(data, i, false)
}

// DecodeValueCompact is DecodeValue for compact JSON — input with no whitespace
// between tokens — skipping the inter-token whitespace scans DecodeValue makes
// while walking objects and arrays. The generator routes the dynamic any/map
// value path here for a //lightning:compact type. On compact input it behaves
// identically to DecodeValue but faster; given inter-token whitespace it may
// report an error.
func DecodeValueCompact(data []byte, i int) (any, int, error) {
	return decodeValue(data, i, true)
}

func decodeValue(data []byte, i int, compact bool) (any, int, error) {
	if i >= len(data) {
		return nil, i, ErrTruncated
	}
	switch data[i] {
	case '"':
		s, end, err := ReadStringOrNull(data, i)
		return s, end, err
	case '{':
		return decodeAnyObject(data, i, compact)
	case '[':
		return decodeAnyArray(data, i, compact)
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

func decodeAnyObject(data []byte, i int, compact bool) (any, int, error) {
	// data[i] == '{'
	i++
	m := map[string]any{}
	for {
		i = skipWSc(data, i, compact)
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
		i = skipWSc(data, ni, compact)
		if i >= len(data) || data[i] != ':' {
			return nil, i, ErrExpectColon
		}
		i = skipWSc(data, i+1, compact)
		val, end, err := decodeValue(data, i, compact)
		if err != nil {
			return nil, end, err
		}
		m[string([]byte(key))] = val
		i = skipWSc(data, end, compact)
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

func decodeAnyArray(data []byte, i int, compact bool) (any, int, error) {
	// data[i] == '['
	i++
	a := []any{}
	for {
		i = skipWSc(data, i, compact)
		if i >= len(data) {
			return nil, i, ErrTruncated
		}
		if data[i] == ']' {
			return a, i + 1, nil
		}
		val, end, err := decodeValue(data, i, compact)
		if err != nil {
			return nil, end, err
		}
		a = append(a, val)
		i = skipWSc(data, end, compact)
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
