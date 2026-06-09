// Package lightning is the public API for the lightning JSON toolkit. It
// re-exports the handful of scanner primitives that are useful on their own,
// independently of the generated unmarshalers.
package lightning

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

// UnescapeStringInPlace is like UnescapeString but decodes the escapes into in
// itself instead of allocating, so it never allocates. Because unescaping never
// lengthens a string, the decoded bytes always fit in the space the escaped
// bytes occupied.
//
// The returned string aliases in and in's contents are overwritten in the
// process, so the caller must not rely on in's original (escaped) bytes
// afterwards and must keep in unchanged while the result is in use.
func UnescapeStringInPlace(in []byte) (string, error) {
	return support.UnescapeStringInPlace(in)
}
