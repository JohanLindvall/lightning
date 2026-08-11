// Package json is the public, stable API for the lightning JSON toolkit — a set
// of allocation-light operations that work directly on raw JSON bytes, without
// unmarshaling into a Go value first. It complements the generated
// json.Unmarshaler implementations produced by the lightning command: use a
// generated decoder when the whole document maps onto a known struct, and this
// package to read, edit, or reshape arbitrary JSON when it does not.
//
// The toolkit falls into four groups:
//
//   - Read — extract raw values by key or path without decoding the rest of the
//     document: [Get]/[GetCompact], [Lookup]/[LookupCompact] (Get without its
//     offset return, for when the value's position in the input is not needed —
//     a missing key is still ErrKeyNotFound), [GetMany]/[GetManyCompact]
//     (several top-level keys in one pass), [GetPaths]/[GetPathsCompact] (several
//     nested paths in one prefix-sharing pass), and [ObjectEach]/[ObjectEachCompact]
//     and [ArrayEach]/[ArrayEachCompact] (iterate an object's members or an
//     array's elements). Returned values alias the input, so the caller must keep
//     it unchanged while they are in use.
//   - Edit — splice raw values into a document, creating any missing path:
//     [Set], [SetMany], [SetPaths]. These write into a caller-provided buffer and
//     allocate nothing when it is reused.
//   - Transform — [StripDefaults] drops object members equal to a set of default
//     values; [EscapeString]/[EscapeStringInto] and
//     [UnescapeString]/[UnescapeStringInto] convert between a Go string and a JSON
//     string body (escaping coerces ill-formed UTF-8 to U+FFFD as encoding/json
//     does when marshaling; unescaping passes it through, see each function);
//     [DecodeAny] decodes a whole document into the generic
//     nil/bool/float64/string/[]any/map[string]any representation; [ParseFloat]
//     parses a single JSON number.
//   - Check — [Valid] reports whether a document is one well-formed JSON value,
//     without allocating or building the decoded value. It accepts exactly what
//     [DecodeAny] accepts — that equivalence is the one this package fuzz-tests —
//     which differs in a few documented ways from encoding/json.Valid. For a
//     generated UnmarshalJSON it is a guide rather than an equivalence, in both
//     directions; see [Valid].
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
	// ErrKeyNotFound is returned by Get when a requested key does not exist in
	// the object it was looked up in. Descending through a value that is not an
	// object reports ErrExpectObject instead.
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
	// ErrBadNumber reports a malformed or unrepresentable JSON number (returned
	// by ParseFloat, DecodeAny, Valid and generated decoders; 1e309 is
	// unrepresentable in a float64 and so rejected).
	ErrBadNumber = unstable.ErrBadNumber
	// ErrBadEscape reports an invalid backslash escape in a JSON string.
	ErrBadEscape = unstable.ErrBadEscape
	// ErrBadUnicode reports a malformed \uXXXX escape in a JSON string.
	ErrBadUnicode = unstable.ErrBadUnicode
	// ErrBadTime reports a time value a generated time.Time field could not
	// parse (neither RFC 3339 nor a Unix timestamp).
	ErrBadTime = unstable.ErrBadTime
)

// MaxDepth is how deeply nested a document may be before DecodeAny, Valid,
// StripDefaults and the generated decoders for recursive schemas give up
// (with ErrMaxDepth; StripDefaults, having no error return, stops stripping
// and copies the remainder through verbatim). It matches encoding/json's limit.
//
// The bound exists because these walks recurse once per nesting level, and a Go
// stack overflow is a fatal error that recover cannot catch: without it, deeply
// nested input would end the process rather than return an error. Get, Lookup
// and GetMany walk iteratively; GetPaths and Set recurse only per requested
// path segment, so their depth is bounded by the caller's path length rather
// than by the document. None of those five need the bound.
const MaxDepth = unstable.MaxDepth

// UnescapeString decodes the body of a JSON string (the bytes that sit between
// the surrounding quotes, with escape sequences such as \n, \" and \uXXXX still
// present) into the Go string it represents.
//
// When in contains no escapes the returned string aliases in directly without
// copying, so the caller must keep in unchanged while the result is in use.
// When escapes are present a new string is allocated.
//
// Only escapes are decoded: the bytes between them are copied through as they
// are, so raw invalid UTF-8 survives into the result where encoding/json would
// coerce it to U+FFFD. Call utf8.ValidString if that matters to you. (An
// unpaired \u surrogate escape is the one thing normalized to U+FFFD, because
// no byte sequence encodes it.)
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
// overwritten. In place is safe here only because unescaping never lengthens a
// string, so the write cursor always trails the read cursor. The reverse
// direction has no such form: EscapeStringInto's out must not overlap s.
func UnescapeStringInto(in, out []byte) (string, error) {
	return unstable.UnescapeStringInto(in, out)
}

// ParseFloat parses the number in b as a float64. It takes the scanner's
// fast paths — Clinger (an exact mantissa with a small decimal exponent becomes
// a single multiply or divide), then Eisel-Lemire for longer mantissas and
// larger exponents — and falls back to strconv.ParseFloat for the rest.
// b must be exactly one number with no surrounding whitespace; trailing bytes or
// an empty input yield an error.
//
// It accepts the numbers the rest of this library accepts rather than exactly
// the JSON grammar's: a leading '+' as well as '-' (ParseFloat([]byte("+5")) is
// 5, nil), leading zeros, an empty integer part and an empty fraction are all
// read, while a magnitude no float64 can hold (1e309) is ErrBadNumber. [Valid]
// documents that accept set and why it is shared.
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
//
// It accepts exactly the documents [Valid] accepts, so the leniencies listed
// there are DecodeAny's too: numbers follow the scanner's grammar rather than
// JSON's (+5 decodes to float64(5)), and string contents are handed back
// verbatim — invalid UTF-8 included, where encoding/json coerces it to U+FFFD.
// Only an unpaired \u surrogate escape is normalized to U+FFFD, as it is there.
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
