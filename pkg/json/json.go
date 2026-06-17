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
// the raw bytes of the value found at that path, without reporting a value type.
// It is built on the scanner primitives, so non-target object members are skipped
// without allocating and the returned slice aliases data (the caller must keep
// data unchanged while the result is in use).
//
// For a string the result includes the surrounding quotes with escapes intact;
// for an object or array it spans the whole '{'..'}' or '['..']'; for a scalar
// it is the literal token. The returned int is the offset in data at which the
// value begins. Each key must name a member of the object at the current level,
// otherwise Get returns ErrKeyNotFound; with no keys it returns the document's
// root value. To check "value exists and is an object" without inspecting a
// type, test err == nil and that the first byte is '{'.
func Get(data []byte, keys ...string) ([]byte, int, error) {
	return support.Get(data, keys...)
}

// GetCompact is Get for compact JSON — input with no whitespace between tokens,
// as compact serializers emit. It skips the inter-token whitespace scans Get
// makes while descending the key path (leading whitespace at the document start
// is still tolerated), so it is faster on compact input but may report an error
// if the input actually contains inter-token whitespace.
func GetCompact(data []byte, keys ...string) ([]byte, int, error) {
	return support.GetCompact(data, keys...)
}

// GetMany looks up several top-level members of the JSON object in data at once,
// in a single pass, returning their raw value bytes in the same order as keys —
// the batch form of Get for pulling a handful of fields out of one record. The
// results are written into out[:0] (pass a slice to reuse across calls; a nil
// out allocates) and out is returned with length len(keys): out[n] is the value
// for keys[n], aliasing data and following Get's conventions, or nil if that key
// is absent. A missing key is reported by the nil slot, not an error; a
// non-object root or malformed JSON returns a non-nil error.
func GetMany(data []byte, keys []string, out [][]byte) ([][]byte, error) {
	return support.GetMany(data, keys, out)
}

// GetManyCompact is GetMany for compact JSON — input with no whitespace between
// tokens, as compact serializers emit. It skips GetMany's inter-token whitespace
// scans (leading whitespace at the document start is still tolerated), so it is
// faster on compact input but may report an error if the input actually contains
// inter-token whitespace.
func GetManyCompact(data []byte, keys []string, out [][]byte) ([][]byte, error) {
	return support.GetManyCompact(data, keys, out)
}

// ObjectEach calls fn once for every member of the JSON object reached by the
// object-key path keys in data, without reporting a value type. fn receives the
// member's decoded key and the raw bytes of its value (both aliasing data, so the
// caller must keep data unchanged while they are in use); the value follows the same
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

// ObjectEachCompact is ObjectEach for compact JSON — input with no whitespace
// between tokens, as compact serializers emit. It skips ObjectEach's inter-token
// whitespace scans (leading whitespace at the document start is still
// tolerated), so it is faster on compact input but may report an error if the
// input actually contains inter-token whitespace.
func ObjectEachCompact(data []byte, fn func(key string, value []byte) error, keys ...string) error {
	return support.ObjectEachCompact(data, fn, keys...)
}

// WhitespaceMode selects how StripDefaults treats inter-token whitespace; see the
// constants below. The zero value (RemoveWhitespace) tolerates any whitespace and
// produces compact output.
type WhitespaceMode = support.WhitespaceMode

const (
	// RemoveWhitespace scans past and drops inter-token whitespace; output is compact.
	RemoveWhitespace = support.RemoveWhitespace
	// AssumeCompact skips the whitespace scans entirely, asserting the input has
	// none (the form compact serializers emit); faster, but misreads spaced input.
	AssumeCompact = support.AssumeCompact
	// PreserveWhitespace keeps the input's whitespace around surviving content, so
	// a pretty-printed document stays pretty-printed.
	PreserveWhitespace = support.PreserveWhitespace
)

// StripDefaults copies the JSON document in input to output, dropping every object
// member whose value is a "default" and then dropping any object or array left
// empty as a result. A value is a default when it is byte-equal to one of
// defaults (compared against the bare token — the unquoted contents for a string,
// the literal token for a number or keyword); empty values are dropped only when
// an empty entry ("") is listed in defaults. A member is kept despite a default
// value when its unquoted key is byte-equal to one of keep.
//
// output is filled from the front and the populated prefix is returned; input is
// not modified. StripDefaults never lengthens the document, so output is grown
// (allocated) only when cap(output) < len(input); pass output == input[:0] to
// strip in place. The returned slice aliases whichever buffer was written.
// StripDefaults is best effort and copies malformed input through unchanged.
//
// ws controls inter-token whitespace handling (see WhitespaceMode): produce
// compact output (RemoveWhitespace), skip the scans for known-compact input
// (AssumeCompact), or keep the input's formatting (PreserveWhitespace).
func StripDefaults(input, output []byte, defaults, keep [][]byte, ws WhitespaceMode) []byte {
	return support.StripDefaults(input, output, defaults, keep, ws)
}

// Set returns the JSON document in with the value at the object-key path keys
// replaced by the raw JSON value rawVal, written into out. When the path does
// not exist it is created: a missing member is inserted into its parent object,
// and a missing intermediate object — or a non-object value found where the path
// still needs to descend — is created as nested objects. With no keys the whole
// document is replaced by rawVal.
//
// rawVal is inserted verbatim and must be a single well-formed JSON value; newly
// created keys are written as JSON strings without escaping, so they should not
// contain characters that need escaping. Inter-token whitespace in in is
// tolerated and preserved outside the edited span.
//
// out is filled from out[:0] and returned (pass a reusable buffer to avoid
// allocation; a nil out allocates). out must not alias in.
func Set(in, out, rawVal []byte, keys []string) []byte {
	return support.Set(in, out, rawVal, keys)
}
