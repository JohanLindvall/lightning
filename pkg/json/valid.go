package json

import "github.com/JohanLindvall/lightning/pkg/unstable"

// validStackWords sizes the container-kind bitset in validValue: one bit per open
// container, bounded by unstable.MaxDepth.
const validStackWords = unstable.MaxDepth/64 + 1

// Valid reports whether data is a single well-formed JSON document: exactly one
// value, optionally surrounded by whitespace, with no trailing content.
//
// Valid is the precondition check for this library's own decoders: it accepts
// exactly what DecodeAny — and therefore a generated UnmarshalJSON — accepts, so
// "Valid(data)" answers "will lightning decode this?". That is a deliberately
// different question from encoding/json.Valid, whose answer is "does this match
// the JSON grammar?", and the two differ in a few places because the decoder's
// scanners are tuned for input already known to be JSON:
//
//   - Numbers follow the scanner's arithmetic, not the JSON grammar. A leading
//     '+', a leading zero, and an empty fraction (+1, 01, 5., .5) are accepted;
//     a magnitude that overflows a float64 (1e309) is rejected, because the
//     decoder cannot represent it. encoding/json.Valid does the opposite on both
//     counts.
//   - Whitespace between tokens is every byte <= 0x20, not just the four bytes
//     JSON names, because that test is one compare instead of four on the
//     decoder's hottest loop.
//   - A raw control byte inside a string is accepted: the string scanners stop at
//     '"' and '\\' only.
//
// Everything else is checked strictly, as the decoder checks it: a trailing comma,
// a non-string object key, a missing colon, an unknown escape, a \u escape without
// four hex digits, an unterminated string, a mismatched bracket, a bad keyword and
// trailing bytes after the value are all rejected. Semantics are not checked —
// duplicate keys are accepted, and neither UTF-8 well-formedness nor surrogate
// pairing is verified (an unpaired \uD800 decodes to U+FFFD rather than failing).
//
// Nesting deeper than unstable.MaxDepth (10000, matching encoding/json's bound) is
// rejected rather than descended into.
//
// Valid neither allocates nor builds the decoded value: strings are scanned in
// place, numbers are read with the decoder's own reader and discarded, and the
// nesting state is a stack-resident bitset. It is therefore far cheaper than
// calling DecodeAny to see whether it errors, and unlike SkipValue's lenient
// bracket balancing (which would accept a trailing comma) it is a real parse.
func Valid(data []byte) bool {
	i := unstable.SkipWS(data, 0)
	i, ok := validValue(data, i)
	if !ok {
		return false
	}
	return unstable.SkipWS(data, i) == len(data)
}

// validValue checks the single JSON value at data[i] and returns the offset just
// past it. It is a flat loop rather than recursion: the open containers live in a
// bitset (one bit per level, set for an object and clear for an array) instead of
// on the goroutine stack, so however deeply nested the input is, validating it
// costs bits and not stack frames — a Go stack overflow being fatal and beyond
// recover's reach. MaxDepth bounds the bitset to a fixed ~1.2 KiB local that
// never escapes.
//
// The scalar cases delegate to the decoder's own readers wherever a reader exists
// that does not allocate, so the two agree by construction rather than by parallel
// reimplementation: numbers go through ReadFloat64OrNull (the exact tier chain
// decodeValue uses, overflow behavior included). Strings are the exception — the
// decoder's reader unescapes, and so allocates — and are checked by validString.
//
// The three labels are the parser's states: scanValue wants a value, scanKey wants
// a member's "key": prefix, and scanAfter has just finished a value and looks for a
// comma or the enclosing close bracket. All locals are declared up front because Go
// forbids a goto that jumps over a declaration.
func validValue(data []byte, i int) (int, bool) {
	var (
		stack [validStackWords]uint64
		depth int
		ok    bool
		isObj bool
		err   error
	)
	n := len(data)

scanValue:
	if i >= n {
		return i, false
	}
	switch data[i] {
	case '{':
		if depth == unstable.MaxDepth {
			return i, false
		}
		stack[depth/64] |= 1 << (depth % 64)
		depth++
		i = unstable.SkipWS(data, i+1)
		// An empty object is complete here. Handling it at the open brace is what
		// lets scanKey reject a '}' outright as a trailing comma.
		if i < n && data[i] == '}' {
			i++
			depth--
			goto scanAfter
		}
		goto scanKey
	case '[':
		if depth == unstable.MaxDepth {
			return i, false
		}
		stack[depth/64] &^= 1 << (depth % 64)
		depth++
		i = unstable.SkipWS(data, i+1)
		// As above: an empty array is complete, so a ']' after a comma is a
		// trailing comma and falls through to scanValue's failure.
		if i < n && data[i] == ']' {
			i++
			depth--
			goto scanAfter
		}
		goto scanValue
	case '"':
		if i, ok = validString(data, i); !ok {
			return i, false
		}
	case 't':
		if !hasLiteral(data, i, "true") {
			return i, false
		}
		i += len("true")
	case 'f':
		if !hasLiteral(data, i, "false") {
			return i, false
		}
		i += len("false")
	case 'n':
		if !hasLiteral(data, i, "null") {
			return i, false
		}
		i += len("null")
	default:
		// The decoder's own number path: same Clinger/Eisel-Lemire/strconv tiers,
		// same span consumed, same rejection of an unrepresentable magnitude. The
		// value is discarded — this is a check, not a decode — and no allocation
		// happens either way.
		if _, i, err = unstable.ReadFloat64OrNull(data, i); err != nil {
			return i, false
		}
	}

scanAfter:
	if depth == 0 {
		// The outermost value is complete; Valid checks for trailing content.
		return i, true
	}
	i = unstable.SkipWS(data, i)
	if i >= n {
		return i, false
	}
	isObj = stack[(depth-1)/64]&(1<<((depth-1)%64)) != 0
	switch data[i] {
	case ',':
		i = unstable.SkipWS(data, i+1)
		if isObj {
			goto scanKey
		}
		goto scanValue
	case '}':
		if !isObj {
			return i, false // ']' expected: brackets must match by kind
		}
		i++
		depth--
		goto scanAfter
	case ']':
		if isObj {
			return i, false
		}
		i++
		depth--
		goto scanAfter
	}
	return i, false

scanKey:
	// A member needs a quoted key and a colon. i is at the first non-space byte of
	// the key, never at a '}' from an empty object (handled at the open brace), so
	// a '}' here is the trailing comma of {"a":1,} and is rejected below.
	if i >= n || data[i] != '"' {
		return i, false
	}
	if i, ok = validString(data, i); !ok {
		return i, false
	}
	i = unstable.SkipWS(data, i)
	if i >= n || data[i] != ':' {
		return i, false
	}
	i = unstable.SkipWS(data, i+1)
	goto scanValue
}

// hasLiteral reports whether the keyword lit sits at data[i].
func hasLiteral(data []byte, i int, lit string) bool {
	return i+len(lit) <= len(data) && string(data[i:i+len(lit)]) == lit
}

// validString checks the quoted string starting at data[i] (which is '"') and
// returns the offset just past its closing quote.
//
// It mirrors what the decoder's string readers accept without unescaping (and so
// without their scratch allocation): the literal runs are crossed with the same
// vectorized unstable.IndexCloseOrEscape, which stops at '"' and '\\' only — hence
// a raw control byte is accepted here exactly as the decoder accepts it — and each
// escape is checked against the set decodeEscaped decodes: the eight single-byte
// forms and \u with four hex digits. Surrogate pairing is not checked, matching the
// decoder, which maps an unpaired half to U+FFFD instead of failing.
func validString(data []byte, i int) (int, bool) {
	i++ // opening quote
	for {
		i += unstable.IndexCloseOrEscape(data[i:])
		if i >= len(data) {
			return i, false // unterminated: ran out before a closing quote
		}
		if data[i] == '"' {
			return i + 1, true
		}
		// data[i] == '\\'
		i++
		if i >= len(data) {
			return i, false
		}
		switch data[i] {
		case '"', '\\', '/', 'b', 'f', 'n', 'r', 't':
			i++
		case 'u':
			if i+5 > len(data) {
				return i, false
			}
			for _, c := range data[i+1 : i+5] {
				if !isHex(c) {
					return i, false
				}
			}
			i += 5
		default:
			return i, false // unknown escape
		}
	}
}

func isHex(c byte) bool {
	return c-'0' < 10 || (c|0x20)-'a' < 6
}
