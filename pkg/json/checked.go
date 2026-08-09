package json

import (
	"errors"

	"github.com/JohanLindvall/lightning/pkg/unstable"
)

// ErrValueCount reports that a checked edit was given fewer raw values than keys
// or paths to set. The unchecked Set/SetMany/SetPaths silently ignore the surplus.
var ErrValueCount = errors.New("json: fewer raw values than keys")

// ErrUnsafeKey reports that a checked edit was given a key (or path element) the
// setters cannot write: Set/SetMany/SetPaths emit a created key raw between two
// quotes — no escaping pass, which is part of what keeps them allocation-free —
// so a key containing '"', '\\' or a control byte is not written back as itself.
//
// This is a security check, not a tidiness one. A key crafted to close its own
// string and open a balanced member (x":1,"role) splices attacker-chosen members
// into the document, and because the forgery is well-formed JSON the wrappers'
// result-side Valid pass accepts it — the injected {"role":"admin"} duplicate then
// wins in every last-key-wins reader, encoding/json included. Keys are therefore
// checked before the edit, not after it.
var ErrUnsafeKey = errors.New("json: key needs escaping")

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
// "Valid" here means what Valid means throughout this package: a document
// DecodeAny accepts, which is a slightly different set from the one encoding/json
// accepts, and only a guide to what a *generated* decoder accepts (see Valid, which
// spells out both directions). Each edit wrapper additionally checks the
// keys it is asked to set (see ErrUnsafeKey), since a created key is written raw
// and so is the one argument the result-side Valid pass cannot vet. What these
// variants do not check is the compactness *assertion* — StripDefaults'
// AssumeCompact mode remains a promise the caller makes about the input, exactly
// as elsewhere in this package.

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

// keysSafe reports whether every key can be written raw between two quotes, i.e.
// whether escaping would leave it unchanged. It asks the scanner rather than
// spelling the byte set out again: unstable.IndexEscape is what EscapeString
// itself scans with, over unstable.SwarNeedsEscape, the repo's single definition
// of "JSON must escape this byte". So a key is safe exactly when EscapeString
// would copy it through — bytes JSON does not require escaped (0x7f, any UTF-8
// sequence) are left alone, and the check rejects only keys Set would mangle.
func keysSafe(keys []string) bool {
	for _, k := range keys {
		if unstable.IndexEscape([]byte(k)) != len(k) {
			return false
		}
	}
	return true
}

// pathsSafe is keysSafe over every element of every path: an interior path
// element is written raw too (appendMembers builds the missing intermediate
// objects), so every depth is checked, not just the leaf.
func pathsSafe(paths [][]string) bool {
	for _, p := range paths {
		if !keysSafe(p) {
			return false
		}
	}
	return true
}

// SetChecked is Set with its arguments and result validated. It returns
// ErrInvalidJSON if in is not one well-formed JSON document, if rawVal is not one
// well-formed JSON value, or if the edit somehow produced malformed output, and
// ErrUnsafeKey if any key contains a byte that would have to be escaped ('"',
// '\\' or a control byte) — Set writes created keys raw, so such a key does not
// round-trip and a crafted one injects members into the result. On error the
// returned slice is nil and no result is handed back.
//
// Use Set directly when in is a document you produced or have already validated —
// this variant is for input you do not trust.
func SetChecked(in, out, rawVal []byte, keys []string) ([]byte, error) {
	if !keysSafe(keys) {
		return nil, ErrUnsafeKey
	}
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
// is not, ErrValueCount if rawVal has fewer entries than keys — the case SetMany
// handles by silently ignoring the surplus keys — and ErrUnsafeKey if any key
// contains a byte that would have to be escaped (created keys are written raw;
// see ErrUnsafeKey). On error the returned slice is nil.
func SetManyChecked(in, out []byte, rawVal [][]byte, keys []string) ([]byte, error) {
	if len(rawVal) < len(keys) {
		return nil, ErrValueCount
	}
	if !keysSafe(keys) {
		return nil, ErrUnsafeKey
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
// is not, ErrValueCount if rawVal has fewer entries than paths, and ErrUnsafeKey
// if any element of any path — interior as well as leaf, since the intermediate
// objects are built from those elements — contains a byte that would have to be
// escaped (see ErrUnsafeKey). On error the returned slice is nil.
//
// Note that a path which is a prefix of another is still resolved by SetPaths'
// shorter-wins rule rather than reported: it is a defined outcome, not an error.
func SetPathsChecked(in, out []byte, rawVal [][]byte, paths [][]string) ([]byte, error) {
	if len(rawVal) < len(paths) {
		return nil, ErrValueCount
	}
	if !pathsSafe(paths) {
		return nil, ErrUnsafeKey
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
// PreserveWhitespace reaches the same outcome with the document's outer whitespace
// still in hand, so its result is whitespace rather than nothing; both are reported
// the same way here, as an empty slice (the whitespace is dropped) and a nil error,
// so len() stays the test in every mode for a caller who forwards the result.
func StripDefaultsChecked(input, output []byte, defaults, keep [][]byte, ws WhitespaceMode) ([]byte, error) {
	if !Valid(input) {
		return nil, ErrInvalidJSON
	}
	res := StripDefaults(input, output, defaults, keep, ws)
	// A result holding no token at all is the consumed-document case above, not a
	// malformed one — Valid rejects it either way, so it has to be recognized
	// before the result check. The whitespace test is unstable.SkipWS, the same
	// (deliberately lenient, <= 0x20) notion of whitespace the stripper copied
	// those bytes through with and Valid skips them with.
	if unstable.SkipWS(res, 0) == len(res) {
		return res[:0], nil
	}
	if !Valid(res) {
		return nil, ErrInvalidJSON
	}
	return res, nil
}
