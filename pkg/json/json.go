// Package json is the public API for the lightning JSON toolkit. It re-exports
// the handful of scanner primitives that are useful on their own, independently
// of the generated unmarshalers.
package json

import "github.com/JohanLindvall/lightning/pkg/unstable"

// ErrKeyNotFound is returned by Get when the requested key path does not exist
// in the document (or descends through a value that is not an object).
var ErrKeyNotFound = unstable.ErrKeyNotFound

// UnescapeString decodes the body of a JSON string (the bytes that sit between
// the surrounding quotes, with escape sequences such as \n, \" and \uXXXX still
// present) into the Go string it represents.
//
// When in contains no escapes the returned string aliases in directly without
// copying, so the caller must keep in unchanged while the result is in use.
// When escapes are present a new string is allocated.
func UnescapeString(in []byte) (string, error) {
	return unstable.UnescapeString(in)
}

// UnescapeStringInto is like UnescapeString but writes the decoded bytes
// into out instead of allocating. If in has no escapes the returned string aliases
// in and out is untouched; otherwise the decode goes into out (allocating only
// if cap(out) < len(in), since unescaping never lengthens a string) and the
// returned string aliases out.
//
// Pass out == in (e.g. in[:0]) to decode truly in place. The returned string
// aliases the buffer it was built in, which the caller must keep unchanged
// while the result is in use; when out aliases in, in's original bytes are
// overwritten.
func UnescapeStringInto(in, out []byte) (string, error) {
	return unstable.UnescapeStringInto(in, out)
}

// ParseFloat parses the JSON number in b as a float64. It takes the scanner's
// Clinger fast path — an exact mantissa with a small decimal exponent becomes a
// single multiply or divide — and falls back to strconv.ParseFloat for the rest.
// b must be exactly one number with no surrounding whitespace; trailing bytes or
// an empty input yield an error.
func ParseFloat(b []byte) (float64, error) {
	return unstable.ParseFloat(b)
}

// DecodeAny decodes a whole JSON document into the generic Go representation —
// nil, bool, float64, string, []any, map[string]any — the same shape
// encoding/json produces when unmarshaling into an interface{}. It is the
// dynamic counterpart to a generated unmarshaler, for input whose schema is not
// known ahead of time.
//
// compact assumes the document has no inter-token whitespace and skips the scans
// that would otherwise look for it; it is faster on minified input and may return
// an error if whitespace is in fact present.
//
// Trailing content after the first complete value (other than whitespace) is an
// error.
func DecodeAny(data []byte, compact bool) (any, error) {
	i := unstable.SkipWS(data, 0)
	var (
		v   any
		end int
		err error
	)
	if compact {
		v, end, err = unstable.DecodeValueCompact(data, i)
	} else {
		v, end, err = unstable.DecodeValue(data, i)
	}
	if err != nil {
		return nil, err
	}
	if unstable.SkipWS(data, end) != len(data) {
		return nil, unstable.ErrInvalidJSON
	}
	return v, nil
}
