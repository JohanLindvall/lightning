package unstable

// strictStackWords sizes the container-kind bitset in SkipValueStrict: one bit
// per open container, bounded by MaxDepth.
const strictStackWords = MaxDepth/64 + 1

// SkipValueStrict parses the single JSON value at data[i] and returns the offset
// just past it, or an error if the value is not well-formed JSON as this library
// defines it. Nothing is decoded and nothing is allocated: it is the validating
// counterpart of SkipValue, which is a bracket balancer and therefore accepts
// balanced nonsense ("[1,]", "[1 2 3]", "[1,,2]") that this rejects.
//
// It has two callers, and they want it for opposite halves of the same property.
// pkg/json.Valid is exactly this walk plus a trailing-content check. And a
// generated decoder's ",lax" field uses it as the skip that runs after a failed
// decode: lax must swallow a *type* mismatch (that is the whole point of the
// option) while still failing on a *syntax* error, and the difference between
// those two is precisely "SkipValue accepted it" versus "this accepted it".
//
// The acceptance set is this library's own, not encoding/json's — it is the set
// DecodeAny reads, which is what makes Valid a usable gate in front of these
// decoders. See pkg/json.Valid's documentation for the enumerated differences
// (numbers are checked by arithmetic, so 01/+1/.5/5. are accepted and 1e309 is
// rejected; whitespace is any byte <= 0x20; a raw control byte inside a string
// is accepted). Everything else is checked strictly: a trailing comma, a
// non-string key, a missing colon, an unknown escape, a \u without four hex
// digits, an unterminated string, a mismatched bracket, a bad keyword.
//
// It is a flat loop rather than recursion: the open containers live in a bitset
// (one bit per level, set for an object and clear for an array) instead of on
// the goroutine stack, so however deeply nested the input is, checking it costs
// bits and not stack frames — a Go stack overflow being fatal and beyond
// recover's reach. MaxDepth bounds the bitset to a fixed ~1.2 KiB local that
// never escapes, and nesting past it returns ErrMaxDepth rather than descending.
//
// The scalar cases delegate to the decoder's own readers wherever a reader
// exists that does not allocate, so the two agree by construction rather than by
// parallel reimplementation: numbers go through ReadFloat64OrNull (the exact
// tier chain decodeValue uses, overflow behavior included) — all but the plain
// decimals of an array of numbers or of coordinate points, which the SIMD
// walks at the '[' case pass over unconverted; they take only numbers that
// reader accepts and hand everything else back to it, so the agreement is
// unchanged (TestValidNumberRunMatchesScalar). Strings are the
// exception — the decoder's reader unescapes, and so allocates — and are checked
// by an inline scan and strictStringEscaped.
//
// The three labels are the parser's states: scanValue wants a value, scanKey
// wants a member's "key": prefix, and scanAfter has just finished a value and
// looks for a comma or the enclosing close bracket. All locals are declared up
// front because Go forbids a goto that jumps over a declaration.
func SkipValueStrict(data []byte, i int) (int, error) {
	var (
		stack  [strictStackWords]uint64
		depth  int
		isObj  bool
		err    error
		closed int
	)
	n := len(data)

scanValue:
	if uint(i) >= uint(n) {
		return i, ErrTruncated
	}
	switch data[i] {
	case '{':
		if depth == MaxDepth {
			return i, ErrMaxDepth
		}
		stack[depth/64] |= 1 << (depth % 64)
		depth++
		i = SkipWS(data, i+1)
		// An empty object is complete here. Handling it at the open brace is what
		// lets scanKey reject a '}' outright as a trailing comma.
		if uint(i) < uint(n) && data[i] == '}' {
			i++
			depth--
			goto scanAfter
		}
		goto scanKey
	case '[':
		if depth == MaxDepth {
			return i, ErrMaxDepth
		}
		stack[depth/64] &^= 1 << (depth % 64)
		depth++
		i = SkipWS(data, i+1)
		// As above: an empty array is complete, so a ']' after a comma is a
		// trailing comma and falls through to scanValue's failure.
		if uint(i) < uint(n) && data[i] == ']' {
			i++
			depth--
			goto scanAfter
		}
		// An array of numbers: the SIMD walk passes over its plain decimals
		// (and closes it, when they are all it holds) without converting one.
		// It takes only numbers the reader below accepts and hands back at
		// anything else, so acceptance is unchanged. An array of arrays of
		// numbers — a coordinate ring — has a walk of its own, one level down
		// (hence the depth test), which hands back any point it does not take
		// whole. The first element's first byte picks between them, so each
		// kind of array pays for its own test only.
		if uint(i) < uint(n) {
			if c := data[i]; c-'0' <= 9 || c == '-' {
				if useValidRun {
					if i, closed = validNumberRun(data, i); closed != 0 {
						i++
						depth--
						goto scanAfter
					}
					i = SkipWS(data, i)
				}
			} else if c == '[' && useValidPoints && depth < MaxDepth && uint(i+1) < uint(n) {
				if c = data[i+1]; c-'0' <= 9 || c == '-' || c <= ' ' {
					if i, closed = validPointsRun(data, i); closed != 0 {
						i++
						depth--
						goto scanAfter
					}
					i = SkipWS(data, i)
				}
			}
		}
		goto scanValue
	case '"':
		// Keep the common clean-string scan in this frame. Only an escape
		// needs the validating loop, which resumes at the backslash.
		i = indexCloseOrEscapeAt(data, i+1)
		if uint(i) >= uint(n) {
			return i, ErrTruncated
		}
		if data[i] == '"' {
			i++
		} else if i, err = strictStringEscaped(data, i); err != nil {
			return i, err
		}
	case 't':
		if !hasLiteral(data, i, "true") {
			return i, ErrInvalidJSON
		}
		i += len("true")
	case 'f':
		if !hasLiteral(data, i, "false") {
			return i, ErrInvalidJSON
		}
		i += len("false")
	case 'n':
		if !hasLiteral(data, i, "null") {
			return i, ErrInvalidJSON
		}
		i += len("null")
	case ']', '}':
		// A closing bracket where a value must start: the trailing comma of
		// [1,] or of an object member list (an *empty* container never reaches
		// here — it is completed at the open bracket). Falling through to the
		// number reader
		// would reject it too, but as ErrBadNumber; naming it ErrInvalidJSON
		// matches what the generated container loops report for the same input,
		// which matters now that a ",lax" field surfaces this error. Acceptance
		// is unchanged either way.
		return i, ErrInvalidJSON
	default:
		// The decoder's own number path: same Clinger/Eisel-Lemire/strconv tiers,
		// same span consumed, same rejection of an unrepresentable magnitude. The
		// value is discarded — this is a check, not a decode — and no allocation
		// happens either way.
		if _, i, err = ReadFloat64OrNull(data, i); err != nil {
			return i, err
		}
	}

scanAfter:
	if depth == 0 {
		// The outermost value is complete; the caller decides what may follow it
		// (Valid requires nothing but whitespace, a lax field resumes decoding).
		return i, nil
	}
	i = SkipWS(data, i)
	if uint(i) >= uint(n) {
		return i, ErrTruncated
	}
	isObj = stack[(depth-1)/64]&(1<<((depth-1)%64)) != 0
	switch data[i] {
	case ',':
		i = SkipWS(data, i+1)
		if isObj {
			goto scanKey
		}
		goto scanValue
	case '}':
		if !isObj {
			return i, ErrInvalidJSON // ']' expected: brackets must match by kind
		}
		i++
		depth--
		goto scanAfter
	case ']':
		if isObj {
			return i, ErrInvalidJSON
		}
		i++
		depth--
		goto scanAfter
	}
	return i, ErrInvalidJSON

scanKey:
	// A member needs a quoted key and a colon. i is at the first non-space byte of
	// the key, never at a '}' from an empty object (handled at the open brace), so
	// a '}' here is the trailing comma of {"a":1,} and is rejected below.
	if uint(i) >= uint(n) {
		return i, ErrTruncated
	}
	if data[i] != '"' {
		return i, ErrInvalidJSON
	}
	i = indexCloseOrEscapeAt(data, i+1)
	if uint(i) >= uint(n) {
		return i, ErrTruncated
	}
	if data[i] == '"' {
		i++
	} else if i, err = strictStringEscaped(data, i); err != nil {
		return i, err
	}
	i = SkipWS(data, i)
	if uint(i) >= uint(n) || data[i] != ':' {
		return i, ErrExpectColon
	}
	i = SkipWS(data, i+1)
	goto scanValue
}

// hasLiteral reports whether the keyword lit sits at data[i].
func hasLiteral(data []byte, i int, lit string) bool {
	return i+len(lit) <= len(data) && string(data[i:i+len(lit)]) == lit
}

// strictStringEscaped checks a string from its first backslash at data[i] and
// returns the offset just past its closing quote. The caller has already scanned
// the clean prefix, so escaped strings do not pay to scan that prefix again.
//
// It mirrors what the decoder's string readers accept without unescaping (and so
// without their scratch allocation): the literal runs are crossed with the same
// vectorized indexCloseOrEscape, which stops at '"' and '\\' only — hence a raw
// control byte is accepted here exactly as the decoder accepts it — and each
// escape is checked against the set decodeEscaped decodes: the eight single-byte
// forms and \u with four hex digits. Surrogate pairing is not checked, matching
// the decoder, which maps an unpaired half to U+FFFD instead of failing.
func strictStringEscaped(data []byte, i int) (int, error) {
	for {
		// data[i] == '\\'
		i++
		if uint(i) >= uint(len(data)) {
			return i, ErrTruncated
		}
		if unescByte[data[i]] != 0 {
			i++
		} else if data[i] == 'u' {
			hex := data[i+1:]
			if len(hex) < 4 {
				return i, ErrTruncated
			}
			// Only validity matters here: OR the decoder's nibble entries
			// without assembling a code point. Any non-hex byte sets bit 16.
			if hexNibble[hex[0]]|hexNibble[hex[1]]|hexNibble[hex[2]]|hexNibble[hex[3]] >= 1<<16 {
				return i, ErrBadUnicode
			}
			i += 5
		} else {
			return i, ErrBadEscape // unknown escape
		}
		// Consecutive escapes and the closing quote have no literal run.
		// Avoid entering the scanner just to find the byte already at i.
		if uint(i) < uint(len(data)) && data[i] != '\\' && data[i] != '"' {
			i = indexCloseOrEscapeAt(data, i)
		}
		if uint(i) >= uint(len(data)) {
			return i, ErrTruncated // unterminated: ran out before a closing quote
		}
		if data[i] == '"' {
			return i + 1, nil
		}
	}
}
