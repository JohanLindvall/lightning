package unstable

import "strconv"

// DecodeValue decodes an arbitrary JSON value at data[i] into the standard Go
// representation (nil, bool, float64, string, []any, map[string]any).
func DecodeValue(data []byte, i int) (any, int, error) {
	return decodeValue(data, i, false)
}

// DecodeValueCompact is DecodeValue for compact JSON — input with no whitespace
// between tokens — skipping the inter-token whitespace scans DecodeValue makes
// while walking objects and arrays. The generator routes the dynamic any/map
// value path here for a //lightning:compact decoder. On compact input it behaves
// identically to DecodeValue but faster; given inter-token whitespace it may
// report an error.
func DecodeValueCompact(data []byte, i int) (any, int, error) {
	return decodeValue(data, i, true)
}

func decodeValue(data []byte, i int, compact bool) (any, int, error) {
	if i >= len(data) {
		return nil, i, ErrTruncated
	}
	switch data[i] {
	case '"':
		s, end, err := ReadStringOrNull(data, i)
		return s, end, err
	case '{':
		return decodeAnyObject(data, i, compact)
	case '[':
		return decodeAnyArray(data, i, compact)
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

func decodeAnyObject(data []byte, i int, compact bool) (any, int, error) {
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
		key, ni, err := ReadKey(data, i)
		if err != nil {
			return nil, ni, err
		}
		i = SkipWSCompact(data, ni, compact)
		if i >= len(data) || data[i] != ':' {
			return nil, i, ErrExpectColon
		}
		i = SkipWSCompact(data, i+1, compact)
		val, end, err := decodeValue(data, i, compact)
		if err != nil {
			return nil, end, err
		}
		m[string([]byte(key))] = val
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

func decodeAnyArray(data []byte, i int, compact bool) (any, int, error) {
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
		val, end, err := decodeValue(data, i, compact)
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
