// Package unstable provides the JSON scanning primitives the lightning code
// generator's output (and the pkg/json toolkit) call into.
//
// It is NOT a stable API. The package is exported only because generated
// *_unmarshal.go files — which live in other modules — must import it; its
// exported functions, signatures, and error sentinels may change or be removed
// in any release. Do not import it directly. The name says so on purpose.
//
// Generated decoders call the exported Read*/Skip*/Decode*/ExpectNull functions
// and the Err* sentinels; the unexported helpers are internal to this package.
// The scanner is index based and avoids allocation on the common paths
// (unescaped strings, integers, object keys).
package unstable

import (
	"errors"
	"unsafe"
)

// Errors returned by the scanner and by generated decoders.
var (
	ErrInvalidJSON  = errors.New("json: invalid JSON")
	ErrTruncated    = errors.New("json: truncated JSON")
	ErrBadEscape    = errors.New("json: invalid string escape")
	ErrBadUnicode   = errors.New("json: invalid unicode escape")
	ErrBadNumber    = errors.New("json: invalid number")
	ErrExpectColon  = errors.New("json: expected ':'")
	ErrExpectObject = errors.New("json: expected '{'")
	ErrExpectArray  = errors.New("json: expected '['")
	ErrBadTime      = errors.New("json: invalid time")
	ErrKeyNotFound  = errors.New("json: key path not found")
	ErrMaxDepth     = errors.New("json: exceeded max depth")
)

// MaxDepth bounds how deeply the recursive walkers — DecodeValue and the
// validator — will descend into nested objects and arrays before giving up with
// ErrMaxDepth. Those walkers recurse once per nesting level, so without a bound a
// document of a few hundred thousand brackets exhausts the goroutine stack, and a
// Go stack overflow is a *fatal* error that recover cannot catch: hostile input
// would take the process down rather than return an error.
//
// The limit matches encoding/json's, so input this package accepts is input
// encoding/json would also accept depth-wise.
const MaxDepth = 10000

// UnsafeStr returns a string that aliases b without copying, so the caller must
// keep b unchanged while the result is in use. Exported for the same inlined
// read path; escaped or copied results still go through ReadKey/ReadString*.
func UnsafeStr(b []byte) string { return unsafeStr(b) }

// ExpectNull consumes the literal null at data[i]. The constant-string compare
// compiles to a single word load and compare (no allocation, no memequal call
// for constants <= 16 bytes) instead of four byte compares; a partial literal
// ("nul" at end of input) fails the bounds test and returns i, exactly as the
// byte-at-a-time form did.
func ExpectNull(data []byte, i int) (int, error) {
	if i+4 > len(data) || string(data[i:i+4]) != "null" {
		return i, ErrInvalidJSON
	}
	return i + 4, nil
}

func unsafeStr(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return unsafe.String(unsafe.SliceData(b), len(b))
}

// SameBuffer reports whether a and b are backed by the same array, which is how
// the pkg/json rewriters recognize an in-place call (out passed as in[:0]) and
// take the extra care that requires. It answers "same backing array", not
// "overlapping": a slice into the middle of another reports false, so a caller
// must treat false as "not proven separate" and stay on the safe path.
func SameBuffer(a, b []byte) bool {
	return cap(a) > 0 && cap(b) > 0 && &a[:1][0] == &b[:1][0]
}
