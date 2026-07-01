package unstable

import (
	"bytes"
	"encoding/base64"
	"unicode/utf16"
	"unicode/utf8"
)

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
// but writes the decoded bytes into out instead of allocating its own buffer:
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
// sized to the string body: since unescaping never lengthens a string, the
// decoded result fits in body-length bytes with a single allocation. Sizing to
// len(data)-start instead over-allocated by the whole remaining document, which
// is catastrophic for an escaped string near the start of a large input (e.g.
// the many "\/"-escaped URLs in a multi-hundred-KB Twitter payload).
//
// The body length is the distance to the closing '"'. The next '"' is located
// with a plain IndexByte rather than an escape-aware SkipString: on densely
// escaped text (\uXXXX-per-character CJK) SkipString stops at every backslash,
// scanning the string a second time, while IndexByte sweeps to the quote in one
// pass. When that quote is not preceded by a backslash it is unescaped — the
// real close — so its offset is the exact body length. Only when it might be an
// escaped \" (rare) is the escape-aware SkipString needed to find the true end;
// the dense-escape strings that made SkipString costly never take that branch.
func decodeStringEscaped(data []byte, start, i int) (string, int, error) {
	capHint := i - start // fallback: the escape-free prefix; append grows the rest
	if rel := bytes.IndexByte(data[start:], '"'); rel >= 0 {
		if rel == 0 || data[start+rel-1] != '\\' {
			capHint = rel // unescaped quote: exact body length
		} else if end, err := SkipString(data, start-1); err == nil {
			capHint = end - 1 - start // escaped \" up front; find the true close
		}
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
		if quoted {
			// Bulk-copy the literal run up to the next escape (or the closing
			// quote) in one vectorized pass — but only when there is a run. Densely
			// escaped strings (\uXXXX-per-character CJK text) sit on a '\' every
			// other step, so checking the current byte first skips a whole
			// indexCloseOrEscape/SSE2 call per escape.
			c := data[i]
			if c != '\\' && c != '"' {
				rest := data[i:]
				k := indexCloseOrEscape(rest)
				if k == len(rest) {
					break // unterminated string
				}
				buf = append(buf, rest[:k]...)
				i += k
				c = data[i]
			}
			if c == '"' {
				// buf is freshly allocated by decodeStringEscaped (the only quoted
				// caller) and never mutated after this point, so it is handed out as
				// a string without the extra copy string(buf) would make — same as
				// the quoted=false return below.
				return unsafeStr(buf), i + 1, nil
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

// hexNibble maps a byte to its hex value (0–15), or 0xFF if it is not a hex
// digit. Four table loads replace the four 3-way comparison chains the
// digit-at-a-time parse ran, and validation falls out of the lookup: a non-hex
// byte yields 0xFF, whose high bit survives the OR in readUnicodeEscape.
var hexNibble = func() [256]uint8 {
	var t [256]uint8
	for i := range t {
		t[i] = 0xFF
	}
	for c := byte('0'); c <= '9'; c++ {
		t[c] = c - '0'
	}
	for c := byte('a'); c <= 'f'; c++ {
		t[c] = c - 'a' + 10
	}
	for c := byte('A'); c <= 'F'; c++ {
		t[c] = c - 'A' + 10
	}
	return t
}()

func readUnicodeEscape(data []byte, i int) (rune, int, error) {
	if i+4 > len(data) {
		return 0, i, ErrTruncated
	}
	a := hexNibble[data[i]]
	b := hexNibble[data[i+1]]
	c := hexNibble[data[i+2]]
	d := hexNibble[data[i+3]]
	if a|b|c|d >= 0x80 { // some byte was not a hex digit (mapped to 0xFF)
		return 0, i, ErrBadUnicode
	}
	return rune(a)<<12 | rune(b)<<8 | rune(c)<<4 | rune(d), i + 4, nil
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
	// The nocopy read hands back an alias for an escape-free string (and a
	// fresh buffer for an escaped one); the []byte conversion below is then the
	// single copy that establishes the returned slice's freshly-allocated,
	// never-retained contract. Reading with ReadStringOrNull here copied
	// escape-free bodies twice — once into the string, once into the slice.
	s, end, err := ReadStringNoCopyOrNull(data, i)
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
