package json

import "errors"

// ErrValueCount reports that a checked edit was given fewer raw values than keys
// or paths to set. The unchecked Set/SetMany/SetPaths silently ignore the surplus.
var ErrValueCount = errors.New("json: fewer raw values than keys")

// The checked counterparts of the edit and transform operations.
//
// Set, SetMany, SetPaths and StripDefaults are best effort by design: they return
// only a []byte, walk the input with a bracket balancer rather than a parser, and
// on input they cannot make sense of they produce something (a passthrough, a
// partial rewrite) instead of complaining. That keeps them fast and allocation-free
// on the hot path — the input is usually a document the caller just produced or
// already validated — but it means malformed input, or a rawVal that is not a
// single JSON value, propagates silently into the result.
//
// The *Checked variants below close that hole for callers handling input they do
// not trust. Each one validates its arguments with Valid before doing the edit and
// validates the result afterwards, so a bad document, a bad inserted value, or a
// result that somehow came out malformed is reported as an error rather than
// returned as bytes. The cost is the extra Valid passes (no allocation, and faster
// than encoding/json.Valid); the unchecked functions are untouched, so nothing that
// does not call these pays for them.
//
// "Valid" here means what Valid means throughout this package: a document this
// library's decoders accept, which is a slightly different set from the one
// encoding/json accepts (see Valid). What these variants do not check is the
// compactness *assertion* — StripDefaults' AssumeCompact mode remains a promise the
// caller makes about the input, exactly as elsewhere in this package.

// SetChecked is Set with its arguments and result validated. It returns
// ErrInvalidJSON if in is not one well-formed JSON document, if rawVal is not one
// well-formed JSON value, or if the edit somehow produced malformed output; on
// error the returned slice is nil and no result is handed back.
//
// Use Set directly when in is a document you produced or have already validated —
// this variant is for input you do not trust.
// allValid reports whether in and every value in rawVal are well-formed JSON —
// the shared argument gate of the checked edit wrappers.
func allValid(in []byte, rawVal [][]byte) bool {
	if !Valid(in) {
		return false
	}
	for _, v := range rawVal {
		if !Valid(v) {
			return false
		}
	}
	return true
}

func SetChecked(in, out, rawVal []byte, keys []string) ([]byte, error) {
	if !Valid(in) || !Valid(rawVal) {
		return nil, ErrInvalidJSON
	}
	res := Set(in, out, rawVal, keys)
	if !Valid(res) {
		return nil, ErrInvalidJSON
	}
	return res, nil
}

// SetManyChecked is SetMany with its arguments and result validated. It returns
// ErrInvalidJSON if in or any rawVal[i] is not well-formed JSON or if the result
// is not, and ErrValueCount if rawVal has fewer entries than keys — the case
// SetMany handles by silently ignoring the surplus keys. On error the returned
// slice is nil.
func SetManyChecked(in, out []byte, rawVal [][]byte, keys []string) ([]byte, error) {
	if len(rawVal) < len(keys) {
		return nil, ErrValueCount
	}
	if !allValid(in, rawVal) {
		return nil, ErrInvalidJSON
	}
	res := SetMany(in, out, rawVal, keys)
	if !Valid(res) {
		return nil, ErrInvalidJSON
	}
	return res, nil
}

// SetPathsChecked is SetPaths with its arguments and result validated. It returns
// ErrInvalidJSON if in or any rawVal[i] is not well-formed JSON or if the result
// is not, and ErrValueCount if rawVal has fewer entries than paths. On error the
// returned slice is nil.
//
// Note that a path which is a prefix of another is still resolved by SetPaths'
// shorter-wins rule rather than reported: it is a defined outcome, not an error.
func SetPathsChecked(in, out []byte, rawVal [][]byte, paths [][]string) ([]byte, error) {
	if len(rawVal) < len(paths) {
		return nil, ErrValueCount
	}
	if !allValid(in, rawVal) {
		return nil, ErrInvalidJSON
	}
	res := SetPaths(in, out, rawVal, paths)
	if !Valid(res) {
		return nil, ErrInvalidJSON
	}
	return res, nil
}

// StripDefaultsChecked is StripDefaults with its input and result validated,
// turning the unchecked function's copy-the-rest-through response to malformed
// input into an ErrInvalidJSON. On error the returned slice is nil.
//
// One outcome that is deliberately *not* an error: stripping can consume the whole
// document. When every member of the root is a default — {"a":{"b":0}} with 0
// among defaults — the empty containers this leaves cascade up and the result is
// zero bytes, which is not valid JSON but is StripDefaults working as documented.
// This variant returns that empty slice with a nil error, so a caller who forwards
// the result should test len() before treating it as a document.
func StripDefaultsChecked(input, output []byte, defaults, keep [][]byte, ws WhitespaceMode) ([]byte, error) {
	if !Valid(input) {
		return nil, ErrInvalidJSON
	}
	res := StripDefaults(input, output, defaults, keep, ws)
	if len(res) > 0 && !Valid(res) {
		return nil, ErrInvalidJSON
	}
	return res, nil
}
