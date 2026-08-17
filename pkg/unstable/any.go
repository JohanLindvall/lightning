package unstable

import "strconv"

// DecodeValue decodes an arbitrary JSON value at data[i] into the standard Go
// representation (nil, bool, float64, string, []any, map[string]any).
//
// It reads through this package's readers, so it inherits their accept set
// rather than the JSON grammar's: numbers take scanFloat's superset (a leading
// '+', leading zeros, an empty integer part or fraction — DecodeValue("+5") is
// float64(5), and so is the element of "[+5]"; see ParseFloat for the full
// list), and strings come back with their bytes verbatim, invalid UTF-8
// included, where encoding/json coerces to U+FFFD (see ReadStringOrNull). Both
// leniencies are shared with Valid, which is what lets Valid promise it accepts
// exactly what this does.
func DecodeValue(data []byte, i int) (any, int, error) {
	return decodeValue(data, i, false, 0)
}

// DecodeValueCompact is DecodeValue for compact JSON — input with no whitespace
// between tokens — skipping the inter-token whitespace scans DecodeValue makes
// while walking objects and arrays. The generator routes the dynamic any/map
// value path here for a //lightning:compact decoder. On compact input it behaves
// identically to DecodeValue but faster; given inter-token whitespace it may
// report an error.
func DecodeValueCompact(data []byte, i int) (any, int, error) {
	return decodeValue(data, i, true, 0)
}

// decodeValue decodes the value at data[i]. depth is the number of enclosing
// objects/arrays; the two container cases pass depth+1 and refuse to descend past
// MaxDepth, which is what keeps a deeply nested document from exhausting the
// stack (see MaxDepth). Only container entry pays for the bound — one compare per
// '{' or '[', nothing on the scalar paths.
func decodeValue(data []byte, i int, compact bool, depth int) (any, int, error) {
	if i >= len(data) {
		return nil, i, ErrTruncated
	}
	switch data[i] {
	case '"':
		// Inline no-escape string fast path (the readKey inline trick, applied
		// by hand): decodeValue is recursive and never inlines, so the added
		// body is free, while the common clean string skips the non-inlined
		// ReadStringOrNull call — only the SIMD scan itself remains
		// (indexCloseOrEscape inlines here). The any path always copies, so
		// string(data[i+1:e]) builds the same value; an escaped or truncated
		// string falls back to ReadStringOrNull, whose error identities and
		// positions are unchanged (it re-checks data[i] == '"' and takes its
		// existing escaped/truncated paths).
		if e := indexCloseOrEscapeAt(data, i+1); e < len(data) && data[e] == '"' {
			return string(data[i+1 : e]), e + 1, nil
		}
		s, end, err := ReadStringOrNull(data, i)
		return s, end, err
	case '{':
		return decodeAnyObject(data, i, compact, depth+1)
	case '[':
		return decodeAnyArray(data, i, compact, depth+1)
	case 't', 'f':
		b, end, err := ReadBoolOrNull(data, i)
		return b, end, err
	case 'n':
		end, err := ExpectNull(data, i)
		return nil, end, err
	default:
		// The number path calls scanFloat directly rather than through
		// ReadFloat64OrNull: that reader is not inlinable, so going through it
		// cost two frames per number (ReadFloat64OrNull → scanFloat) plus
		// re-checks of the truncation and 'n' cases this switch already
		// excludes — the same mechanism as the batched array readers. The
		// strconv fallback mirrors ReadFloat64OrNull byte for byte (see there
		// for why scanFloat may decline).
		f, end, fast, ok := scanFloat(data, i)
		if !ok {
			return nil, end, ErrBadNumber
		}
		if !fast {
			f2, perr := strconv.ParseFloat(unsafeStr(data[i:end]), 64)
			if perr != nil {
				return nil, end, ErrBadNumber
			}
			f = f2
		}
		return f, end, nil
	}
}

func decodeAnyObject(data []byte, i int, compact bool, depth int) (any, int, error) {
	if depth > MaxDepth {
		return nil, i, ErrMaxDepth
	}
	// data[i] == '{'
	i++
	m := map[string]any{}
	// The loop top is reached only after '{' or after a comma, and a '}' after
	// a member returns from the post-value check below — so '}' here on a
	// non-first iteration is a trailing comma ({"a":1,}), rejected as
	// encoding/json does (the first-iteration flag matches the generated
	// decoders).
	for first := true; ; first = false {
		i = SkipWSCompact(data, i, compact)
		if i >= len(data) {
			return nil, i, ErrTruncated
		}
		if data[i] == '}' {
			if first {
				return m, i + 1, nil
			}
			return nil, i, ErrInvalidJSON
		}
		// Inline no-escape key fast path (the readKey inline trick, applied by
		// hand): decodeAnyObject is recursive and never inlines, so the added
		// body is free, while the common clean key skips the non-inlined
		// ReadKey call. The map insert needs a copied key either way (ReadKey's
		// clean result aliases data), so the fast path copies with
		// string(rest[:k]) up front and the insert below drops the defensive
		// string([]byte(key)) re-copy; an escaped key, truncated input, or a
		// non-string falls back to ReadKey, whose error identities and
		// positions are unchanged.
		var key string
		ni := -1
		if data[i] == '"' {
			if e := indexCloseOrEscapeAt(data, i+1); e < len(data) && data[e] == '"' {
				key = string(data[i+1 : e])
				ni = e + 1
			}
		}
		if ni < 0 {
			akey, aend, err := ReadKey(data, i)
			if err != nil {
				return nil, aend, err
			}
			// Reached only for an escaped key (the inline scan above catches
			// every clean one, and a truncated key errors), and ReadKey's
			// escaped path returns a fresh never-reused buffer — safe to own
			// without another copy.
			key, ni = akey, aend
		}
		i = SkipWSCompact(data, ni, compact)
		if i >= len(data) || data[i] != ':' {
			return nil, i, ErrExpectColon
		}
		i = SkipWSCompact(data, i+1, compact)
		val, end, err := decodeValue(data, i, compact, depth)
		if err != nil {
			return nil, end, err
		}
		m[key] = val
		i = SkipWSCompact(data, end, compact)
		if i >= len(data) {
			return nil, i, ErrTruncated
		}
		if data[i] == '}' {
			return m, i + 1, nil
		}
		if data[i] != ',' {
			return nil, i, ErrInvalidJSON
		}
		i++
	}
}

func decodeAnyArray(data []byte, i int, compact bool, depth int) (any, int, error) {
	if depth > MaxDepth {
		return nil, i, ErrMaxDepth
	}
	// data[i] == '['
	i++
	a := []any{}
	// Trailing commas ([1,]) are rejected by the first-iteration flag, as in
	// decodeAnyObject.
	for first := true; ; first = false {
		i = SkipWSCompact(data, i, compact)
		if i >= len(data) {
			return nil, i, ErrTruncated
		}
		if data[i] == ']' {
			if first {
				return a, i + 1, nil
			}
			return nil, i, ErrInvalidJSON
		}
		val, end, err := decodeValue(data, i, compact, depth)
		if err != nil {
			return nil, end, err
		}
		if cap(a) == 0 {
			// First element: the same ~256-byte static first-append capacity
			// hint the generated decoders use for un-presized slices (16-byte
			// `any` elements → 16), instead of append growing 1→2→4→…. `[]`
			// never reaches this point, so it still returns the non-nil empty
			// slice above; a longer array regrows exactly as before.
			a = make([]any, 0, 16)
		}
		a = append(a, val)
		i = SkipWSCompact(data, end, compact)
		if i >= len(data) {
			return nil, i, ErrTruncated
		}
		if data[i] == ']' {
			return a, i + 1, nil
		}
		if data[i] != ',' {
			return nil, i, ErrInvalidJSON
		}
		i++
	}
}
