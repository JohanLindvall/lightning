package json

import "errors"

// The scalar readers take a value exactly as [Get], [Lookup], [GetMany],
// [GetPaths], [ObjectEach] and [ArrayEach] hand it back — a string with its
// quotes on and its escapes intact, a literal token for the rest — and
// answer the Go value. Without them every caller of the walkers wrote the
// same three steps for a string (check the quotes, look for a backslash,
// unescape) and the same literal comparisons for a bool, in every codebase
// that adopted the toolkit; two copies existed in the first one converted.
// The number readers are [ParseFloat] and its siblings.

var (
	// ErrExpectString reports that a string token was required and the value
	// was something else — a number, a literal, a container, or nothing.
	ErrExpectString = errors.New("json: expected a string")
	// ErrExpectBool reports that true or false was required and the value was
	// something else, null included.
	ErrExpectBool = errors.New("json: expected true or false")
)

// String decodes a JSON string token — `"…"`, quotes on, escapes intact,
// which is what the read functions return for a string — into the Go string
// it denotes. It follows [UnescapeString]'s contract exactly: a token with no
// escapes yields a string that ALIASES raw with no copy, so keep raw unchanged
// while the result is in use (or clone it when it outlives the input); a
// token with escapes yields a fresh string. Invalid UTF-8 passes through and
// an unpaired surrogate escape becomes U+FFFD, as there.
//
// raw must be exactly the token, with no surrounding whitespace. Anything
// that does not start and end with a quote — a number, null, a container, an
// unterminated string, an empty input — is ErrExpectString; a bad escape
// inside the quotes is ErrBadEscape or ErrBadUnicode.
func String(raw []byte) (string, error) {
	if len(raw) < 2 || raw[0] != '"' || raw[len(raw)-1] != '"' {
		return "", ErrExpectString
	}
	return UnescapeString(raw[1 : len(raw)-1])
}

// Bool reads the JSON literal true or false. Anything else — null, a number,
// a quoted "true", trailing bytes, an empty input — is ErrExpectBool.
func Bool(raw []byte) (bool, error) {
	switch string(raw) {
	case "true":
		return true, nil
	case "false":
		return false, nil
	}
	return false, ErrExpectBool
}
