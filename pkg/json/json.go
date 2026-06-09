// Package json is the public API for the lightning JSON toolkit. It re-exports
// the handful of scanner primitives that are useful on their own, independently
// of the generated unmarshalers.
package json

import "github.com/JohanLindvall/lightning/pkg/support"

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
