package unstable

import (
	"encoding/binary"
	"math/bits"
)

// SkipValue advances past any JSON value starting at data[i].
func SkipValue(data []byte, i int) (int, error) {
	if i >= len(data) {
		return i, ErrTruncated
	}
	switch data[i] {
	case '"':
		return SkipString(data, i)
	case '{':
		return skipObject(data, i)
	case '[':
		return skipArray(data, i)
	case 't':
		if i+4 <= len(data) && data[i+1] == 'r' && data[i+2] == 'u' && data[i+3] == 'e' {
			return i + 4, nil
		}
		return i, ErrInvalidJSON
	case 'f':
		if i+5 <= len(data) && data[i+1] == 'a' && data[i+2] == 'l' && data[i+3] == 's' && data[i+4] == 'e' {
			return i + 5, nil
		}
		return i, ErrInvalidJSON
	case 'n':
		return ExpectNull(data, i)
	default:
		return skipNumber(data, i)
	}
}

func SkipString(data []byte, i int) (int, error) {
	// data[i] == '"'
	i++
	for {
		rest := data[i:]
		k := indexCloseOrEscape(rest)
		if k == len(rest) {
			return len(data), ErrTruncated
		}
		if rest[k] == '"' {
			return i + k + 1, nil
		}
		// Skip the escape sequence. For \uXXXX we only need to skip the
		// backslash and the next char; subsequent bytes are processed on the
		// next iteration.
		i += k + 2
		if i > len(data) {
			return len(data), ErrTruncated
		}
	}
}

func skipNumber(data []byte, i int) (int, error) {
	start := i
	if i < len(data) && data[i] == '-' {
		i++
	}
	for i < len(data) {
		c := data[i]
		if (c >= '0' && c <= '9') || c == '.' || c == 'e' || c == 'E' || c == '+' || c == '-' {
			i++
			continue
		}
		break
	}
	if i == start {
		return i, ErrBadNumber
	}
	return i, nil
}

func skipObject(data []byte, i int) (int, error) {
	// data[i] == '{'
	i++
	for i < len(data) {
		// Jump to the next structural byte, skipping inert content (keys' inner
		// chars, numbers, bools, whitespace) in one vectorized pass.
		i += indexStructural(data[i:])
		if i >= len(data) {
			break
		}
		switch data[i] {
		case '}':
			return i + 1, nil
		case '{':
			end, err := skipObject(data, i)
			if err != nil {
				return end, err
			}
			i = end
		case '[':
			end, err := skipArray(data, i)
			if err != nil {
				return end, err
			}
			i = end
		case '"':
			end, err := SkipString(data, i)
			if err != nil {
				return end, err
			}
			i = end
		default: // stray ']' (only on malformed input); step over it
			i++
		}
	}
	return i, ErrTruncated
}

func skipArray(data []byte, i int) (int, error) {
	// data[i] == '['
	i++
	for i < len(data) {
		i += indexStructural(data[i:])
		if i >= len(data) {
			break
		}
		switch data[i] {
		case ']':
			return i + 1, nil
		case '[':
			end, err := skipArray(data, i)
			if err != nil {
				return end, err
			}
			i = end
		case '{':
			end, err := skipObject(data, i)
			if err != nil {
				return end, err
			}
			i = end
		case '"':
			end, err := SkipString(data, i)
			if err != nil {
				return end, err
			}
			i = end
		default: // stray '}' (only on malformed input); step over it
			i++
		}
	}
	return i, ErrTruncated
}

// SkipWS advances past JSON whitespace at data[i]. The four JSON whitespace
// bytes (space, tab, newline, carriage return) are all <= ' ' (0x20), so a single
// compare classifies a byte with no memory load — measurably faster than a lookup
// table on every shape from a compact exit to a deep indent run. This is the
// hottest classification in the scanner, running before and after every value.
//
// The compare also treats the other control bytes (0x00..0x1f) as whitespace.
// Those are never valid JSON between tokens, so on well-formed input the result
// is identical to matching the four bytes exactly; on malformed input SkipWS
// skips such a byte rather than stopping on it, leaving the value parser to
// reject the next real token. SkipWS is not called inside strings, so control
// bytes within string contents are unaffected.
func SkipWS(data []byte, i int) int {
	for i < len(data) && data[i] <= ' ' {
		i++
	}
	return i
}

// SkipWSRun advances past a whitespace run from data[i] (the caller has already
// established that data[i-1] and data[i-2] were whitespace, i.e. this is a run of
// at least two). It is the out-of-line continuation the generated decoders call
// only for genuine indentation runs in pretty-printed input — short skips (zero
// or one whitespace byte, the compact and single-space-after-token cases) are
// handled inline at the call site so they never pay a call. Eight bytes are
// classified per word: ws sets a lane's high bit exactly when that byte is <= ' '
// (clearing each lane's top bit into w&^hi keeps the per-lane subtract g-... from
// borrowing across lanes, so it is exact for every byte value, and &^ w rejects
// the >= 0x80 bytes malformed input may carry). A full-whitespace word equals the
// all-high mask and is skipped whole; the first word with a structural byte
// locates it with one trailing-zeros count.
func SkipWSRun(data []byte, i int) int {
	const (
		lo = ^uint64(0) / 255 // 0x0101...01
		hi = lo << 7          // 0x8080...80
		g  = lo*' ' | hi      // 0xA0...A0: ' ' with a per-lane borrow guard
	)
	for i+8 <= len(data) {
		w := binary.LittleEndian.Uint64(data[i:])
		ws := (g - w&^hi) &^ w & hi
		if ws != hi {
			return i + bits.TrailingZeros64(ws^hi)/8
		}
		i += 8
	}
	for i < len(data) && data[i] <= ' ' {
		i++
	}
	return i
}

// SkipWSCompact is the compact-aware inter-token whitespace skip shared by the
// dynamic DecodeValue path and the pkg/json toolkit's compact variants (GetCompact,
// SetMany, ObjectEachCompact, …). In compact mode the input is asserted to carry no
// whitespace between tokens (the form compact JSON serializers emit), so it returns
// i unchanged; otherwise it is SkipWS. This mirrors the generator's
// //lightning:compact decoders, which elide exactly these inter-token skips.
func SkipWSCompact(data []byte, i int, compact bool) int {
	for !compact && i < len(data) && data[i] <= ' ' {
		i++
	}
	return i
}
