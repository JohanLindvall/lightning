// Package json is the public, stable API for the lightning JSON toolkit — a set
// of allocation-light operations that work directly on raw JSON bytes, without
// unmarshaling into a Go value first. It complements the generated
// json.Unmarshaler implementations produced by the lightning command: use a
// generated decoder when the whole document maps onto a known struct, and this
// package to read, edit, or reshape arbitrary JSON when it does not.
//
// The toolkit falls into three groups:
//
//   - Read — extract raw values by key or path without decoding the rest of the
//     document: [Get]/[GetCompact], [GetMany]/[GetManyCompact] (several top-level
//     keys in one pass), [GetPaths]/[GetPathsCompact] (several nested paths in one
//     prefix-sharing pass), and [ObjectEach]/[ObjectEachCompact] (iterate an
//     object's members). Returned values alias the input, so the caller must keep
//     it unchanged while they are in use.
//   - Edit — splice raw values into a document, creating any missing path:
//     [Set], [SetMany], [SetPaths]. These write into a caller-provided buffer and
//     allocate nothing when it is reused.
//   - Transform — [StripDefaults] drops object members equal to a set of default
//     values; [EscapeString]/[EscapeStringInto] and
//     [UnescapeString]/[UnescapeStringInto] convert between a Go string and a JSON
//     string body; [DecodeAny] decodes a whole document into the generic
//     nil/bool/float64/string/[]any/map[string]any representation; [ParseFloat]
//     parses a single JSON number.
//   - Check — [Valid] reports whether a document is one well-formed JSON value,
//     without allocating or building the decoded value. It accepts exactly what
//     this library's decoders accept, which differs in a few documented ways from
//     encoding/json.Valid.
//
// The edit and transform operations are best effort: they return only a []byte and
// pass input they cannot interpret through rather than failing. For untrusted input
// each has a checked counterpart that validates its arguments and result and
// returns an error — [SetChecked], [SetManyChecked], [SetPathsChecked],
// [StripDefaultsChecked].
//
// The *Compact variants assume the input has no inter-token whitespace (the form
// compact serializers emit) and skip the scans that look for it — faster on
// minified input, but they may report an error if whitespace is in fact present.
//
// Key matching is exact and case-sensitive, unlike encoding/json's
// case-insensitive fallback.
//
// This package is implemented on the scanning primitives in pkg/unstable; that
// package is not a stable API and should not be imported directly.
package json

import "github.com/JohanLindvall/lightning/pkg/unstable"

// Errors reported by this package. They are re-exported from pkg/unstable so that
// callers of the stable API can match them with errors.Is without importing the
// unstable package. Any of these may accompany a malformed, truncated, or
// mistyped document; the individual functions document which they can return.
var (
	// ErrKeyNotFound is returned by Get when the requested key path does not
	// exist in the document (or descends through a value that is not an object).
	ErrKeyNotFound = unstable.ErrKeyNotFound
	// ErrInvalidJSON reports syntactically invalid JSON (e.g. a missing separator
	// or an unexpected byte where a value was expected).
	ErrInvalidJSON = unstable.ErrInvalidJSON
	// ErrTruncated reports input that ends in the middle of a value.
	ErrTruncated = unstable.ErrTruncated
	// ErrExpectObject reports a '{' was required (a key descent or object
	// iteration) but the value was not an object.
	ErrExpectObject = unstable.ErrExpectObject
	// ErrExpectArray reports a '[' was required but the value was not an array.
	ErrExpectArray = unstable.ErrExpectArray
	// ErrExpectColon reports a ':' was expected after an object key.
	ErrExpectColon = unstable.ErrExpectColon
	// ErrMaxDepth reports input nested deeper than MaxDepth.
	ErrMaxDepth = unstable.ErrMaxDepth
)

// MaxDepth is how deeply nested a document may be before DecodeAny, Valid and the
// generated decoders for recursive schemas give up with ErrMaxDepth. It matches
// encoding/json's limit.
//
// The bound exists because these walks recurse once per nesting level, and a Go
// stack overflow is a fatal error that recover cannot catch: without it, deeply
// nested input would end the process rather than return an error. Get, Lookup,
// GetMany, GetPaths and Set walk iteratively and are not subject to it.
const MaxDepth = unstable.MaxDepth

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
// Trailing content after the first complete value (other than whitespace) is an
// error.
func DecodeAny(data []byte) (any, error) {
	return decodeAny(data, false)
}

// DecodeAnyCompact is DecodeAny for compact JSON — input with no whitespace
// between tokens, as compact serializers emit — skipping the inter-token
// whitespace scans DecodeAny makes (leading whitespace at the document start is
// still tolerated). On such input it behaves identically to DecodeAny but faster;
// given input that does contain inter-token whitespace it may report an error.
func DecodeAnyCompact(data []byte) (any, error) {
	return decodeAny(data, true)
}

func decodeAny(data []byte, compact bool) (any, error) {
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

// Valid lives in valid.go, which checks the JSON grammar directly rather than
// decoding the document.
