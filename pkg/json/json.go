// Package json is the public API for the lightning JSON toolkit. It re-exports
// the handful of scanner primitives that are useful on their own, independently
// of the generated unmarshalers.
package json

import "github.com/JohanLindvall/lightning/pkg/support"

// ErrKeyNotFound is returned by Get when the requested key path does not exist
// in the document (or descends through a value that is not an object).
var ErrKeyNotFound = support.ErrKeyNotFound

// UnescapeString decodes the body of a JSON string (the bytes that sit between
// the surrounding quotes, with escape sequences such as \n, \" and \uXXXX still
// present) into the Go string it represents.
//
// When in contains no escapes the returned string aliases in directly without
// copying, so the caller must keep in unchanged while the result is in use.
// When escapes are present a new string is allocated.
func UnescapeString(in []byte) (string, error) {
	return support.UnescapeString(in)
}

// UnescapeStringInto is like UnescapeString but writes the decoded bytes
// into out instead of allocating, mirroring the Unescape(in, out) convention of
// github.com/buger/jsonparser. If in has no escapes the returned string aliases
// in and out is untouched; otherwise the decode goes into out (allocating only
// if cap(out) < len(in), since unescaping never lengthens a string) and the
// returned string aliases out.
//
// Pass out == in (e.g. in[:0]) to decode truly in place. The returned string
// aliases the buffer it was built in, which the caller must keep unchanged
// while the result is in use; when out aliases in, in's original bytes are
// overwritten.
func UnescapeStringInto(in, out []byte) (string, error) {
	return support.UnescapeStringInto(in, out)
}

// ParseFloat parses the JSON number in b as a float64. It takes the scanner's
// Clinger fast path — an exact mantissa with a small decimal exponent becomes a
// single multiply or divide — and falls back to strconv.ParseFloat for the rest.
// b must be exactly one number with no surrounding whitespace; trailing bytes or
// an empty input yield an error.
func ParseFloat(b []byte) (float64, error) {
	return support.ParseFloat(b)
}

// Get walks the object-key path keys into the JSON document data and returns
// the raw bytes of the value found at that path, modeled on
// github.com/buger/jsonparser's Get but without reporting a value type. It is
// built on the scanner primitives, so non-target object members are skipped
// without allocating and the returned slice aliases data (the caller must keep
// data unchanged while the result is in use).
//
// For a string the result includes the surrounding quotes with escapes intact;
// for an object or array it spans the whole '{'..'}' or '['..']'; for a scalar
// it is the literal token. The returned int is the offset in data at which the
// value begins. Each key must name a member of the object at the current level,
// otherwise Get returns ErrKeyNotFound; with no keys it returns the document's
// root value. To replicate jsonparser's "value exists and is an object" check
// without inspecting a type, test err == nil and that the first byte is '{'.
func Get(data []byte, keys ...string) ([]byte, int, error) {
	return support.Get(data, keys...)
}

// ObjectEach calls fn once for every member of the JSON object reached by the
// object-key path keys in data, modeled on github.com/buger/jsonparser's
// ObjectEach but without reporting a value type. fn receives the member's
// decoded key and the raw bytes of its value (both aliasing data, so the caller
// must keep data unchanged while they are in use); the value follows the same
// conventions as Get — quotes kept for strings, the full span for objects and
// arrays, the literal token for scalars.
//
// With no keys ObjectEach iterates the document's root object; otherwise each
// key descends one level and the value at the end of the path must itself be an
// object (ErrKeyNotFound if a key is missing). If fn returns a non-nil error,
// iteration stops and that error is returned. It is built on the scanner
// primitives, so members are visited without allocating.
func ObjectEach(data []byte, fn func(key string, value []byte) error, keys ...string) error {
	return support.ObjectEach(data, fn, keys...)
}
