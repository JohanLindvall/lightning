package unstable

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
		f, end, err := ReadFloat64OrNull(data, i)
		return f, end, err
	}
}

func decodeAnyObject(data []byte, i int, compact bool) (any, int, error) {
	// data[i] == '{'
	i++
	m := map[string]any{}
	for {
		i = SkipWSCompact(data, i, compact)
		if i >= len(data) {
			return nil, i, ErrTruncated
		}
		if data[i] == '}' {
			return m, i + 1, nil
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
	for {
		i = SkipWSCompact(data, i, compact)
		if i >= len(data) {
			return nil, i, ErrTruncated
		}
		if data[i] == ']' {
			return a, i + 1, nil
		}
		val, end, err := decodeValue(data, i, compact)
		if err != nil {
			return nil, end, err
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
