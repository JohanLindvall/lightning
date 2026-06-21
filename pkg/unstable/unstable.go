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
)

// UnsafeStr returns a string that aliases b without copying, so the caller must
// keep b unchanged while the result is in use. Exported for the same inlined
// read path; escaped or copied results still go through ReadKey/ReadString*.
func UnsafeStr(b []byte) string { return unsafeStr(b) }

// ExpectNull consumes the literal null at data[i].
func ExpectNull(data []byte, i int) (int, error) {
	if i+4 > len(data) || data[i] != 'n' || data[i+1] != 'u' || data[i+2] != 'l' || data[i+3] != 'l' {
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
