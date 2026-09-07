package json

import (
	"errors"

	"github.com/JohanLindvall/lightning/pkg/unstable"
)

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
	n := len(raw)
	if n < 2 || raw[0] != '"' || raw[n-1] != '"' {
		return "", ErrExpectString
	}
	// Deciding that a short token holds no escape costs a call to
	// bytes.IndexByte — 1.9 ns of this function's 4.0 ns, whatever the length
	// — and a pair of word loads answers for every token up to 32 bytes
	// without one. The token is passed whole rather than its body: a quote is
	// not a backslash, so the two extra bytes cost nothing and they are what
	// makes the windows land inside the input. See unstable.NoBackslash8; the
	// bounds there are what the windows COVER. A longer token, and any token
	// with an escape, goes the ordinary way.
	//
	// Two shapes here are deliberate. One length test admits every token the
	// windows can decide, the degenerate lengths falling out of the switch's
	// first arm, so a token too long for them pays one compare rather than a
	// pair. And each arm returns where it decides rather than setting a flag
	// the code below re-tests, which would cost the CSET that materialises the
	// flag and the branch that reads it on the path every clean short string
	// takes.
	if uint(n) < 33 {
		switch {
		case n < 4:
			// Two or three bytes: an empty or a one-byte body, below the
			// shortest window.
			if n == 2 {
				return "", nil
			}
			if raw[1] != '\\' {
				return unstable.UnsafeStr(raw[1:2]), nil
			}
		case n <= 8:
			if unstable.NoBackslash4(raw) {
				return unstable.UnsafeStr(raw[1 : n-1]), nil
			}
		case n <= 16:
			if unstable.NoBackslash8(raw) {
				return unstable.UnsafeStr(raw[1 : n-1]), nil
			}
		default:
			if unstable.NoBackslash8(raw) && unstable.NoBackslash16(raw) {
				return unstable.UnsafeStr(raw[1 : n-1]), nil
			}
		}
		// The windows covered every byte of the token, so falling out of the
		// switch means it really does hold a backslash. UnescapeString would
		// run the identical tests on the body and be told the same thing, so
		// go straight to its scan.
		return unstable.UnescapeStringScan(raw[1 : n-1])
	}
	return UnescapeString(raw[1 : n-1])
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
