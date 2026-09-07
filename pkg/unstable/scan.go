package unstable

// ValueScanner finds the end of ONE JSON value spread across a sequence of
// chunks. SkipValue needs the whole value in memory; this is what a caller
// that reads a document through a bounded buffer uses instead, in two places:
// to stream past a value it does not want (a sibling member on the way to a
// key path, which may be far larger than the buffer), and to learn when a
// value it DOES want has finally arrived complete, without re-scanning it from
// the start after every refill — a "refill and retry SkipValue from the value's
// first byte" loop is O(n²) in the number of chunks a value arrives in.
//
// Semantics are the scalar skip path's, byte for byte: skipObject/skipArray's
// typed bracket balance (a stray closer of the OTHER kind is stepped over, not
// counted), SkipString's escape handling, skipNumber's accept set of
// [0-9.eE+-] (it measures a number token, it does not validate one) and the
// literal arms' exact match against "true"/"false"/"null". The same errors come
// back: ErrBadNumber for a byte that starts no value, ErrInvalidJSON for a
// misspelt literal, ErrMaxDepth past MaxDepth, and ErrTruncated for input that
// ends inside a string or a container.
//
// It is NOT the same code as SkipValue and does not always agree with it on
// MALFORMED input — SkipValue may take the SIMD in-string-mask path, whose
// divergences skipfast.go enumerates. On well-formed input the two agree on
// every value, which TestValueScannerMatchesSkipValue pins across every chunk
// split of every value in the corpus.
type ValueScanner struct {
	state    uint8
	depth    int
	inString bool
	escaped  bool
	lit      string // the bytes of a literal still to match
	// kinds[d/64] bit d%64 is set when the container opened at depth d+1 is an
	// object, clear when it is an array — the type stack skipObject and
	// skipArray keep as their recursion. Never cleared: every bit is written
	// on the way down before it is read on the way up.
	kinds [MaxDepth/64 + 1]uint64
}

// The scanner's states. A string INSIDE a container is scanContainer with
// inString set; scanString is only a value that is itself a string.
const (
	scanValueStart uint8 = iota
	scanContainer
	scanString
	scanNumber
	scanLiteral
	scanDone
)

// Reset prepares the scanner for a new value. It does not clear the type
// stack, which is written before it is read.
func (s *ValueScanner) Reset() {
	s.state, s.depth, s.inString, s.escaped, s.lit = scanValueStart, 0, false, false, ""
}

// Feed advances the scan over the next chunk of the document, which must
// continue immediately after the previous one; the first chunk starts at the
// value's first byte, with any leading whitespace already skipped. final says
// no more bytes will follow, which is what lets a number or a literal at the
// very end of the input be recognised as complete.
//
// It returns how many bytes of chunk belong to the value, whether the value
// ended within them (the last byte of the value is then chunk[n-1]), and an
// error if the value is malformed. When the value has not ended and there is
// no error, the whole chunk was consumed and the caller should feed more.
func (s *ValueScanner) Feed(chunk []byte, final bool) (n int, done bool, err error) {
	i := 0
	if s.state == scanValueStart {
		if len(chunk) == 0 {
			if final {
				return 0, false, ErrTruncated
			}
			return 0, false, nil
		}
		switch c := chunk[0]; c {
		case '"':
			s.state = scanString
		case '{':
			s.state = scanContainer
			s.push(true)
		case '[':
			s.state = scanContainer
			s.push(false)
		case 't':
			s.state, s.lit = scanLiteral, "rue"
		case 'f':
			s.state, s.lit = scanLiteral, "alse"
		case 'n':
			s.state, s.lit = scanLiteral, "ull"
		default:
			if !isNumberByte(c) {
				// The byte starts no value at all, which is what SkipValue's
				// default arm reports through skipNumber's empty run.
				return 0, false, ErrBadNumber
			}
			s.state = scanNumber
		}
		i = 1
	}
	switch s.state {
	case scanNumber:
		for ; i < len(chunk); i++ {
			if !isNumberByte(chunk[i]) {
				s.state = scanDone
				return i, true, nil
			}
		}
		if final {
			s.state = scanDone
			return i, true, nil
		}
		return i, false, nil

	case scanLiteral:
		for s.lit != "" && i < len(chunk) {
			if chunk[i] != s.lit[0] {
				return i, false, ErrInvalidJSON
			}
			s.lit = s.lit[1:]
			i++
		}
		if s.lit == "" {
			s.state = scanDone
			return i, true, nil
		}
		if final {
			// SkipValue reports a literal it cannot complete as invalid
			// rather than truncated: it tests the four (or five) bytes as a
			// whole and has no separate "ran out" arm.
			return i, false, ErrInvalidJSON
		}
		return i, false, nil

	case scanString:
		end, closed := s.stringBody(chunk, i)
		if closed {
			s.state = scanDone
			return end, true, nil
		}
		if final {
			return end, false, ErrTruncated
		}
		return end, false, nil

	case scanContainer:
		return s.container(chunk, i, final)
	}
	return i, true, nil // scanDone: the caller fed past the end
}

// container is the balance loop, the flat form of skipObject and skipArray's
// mutual recursion: indexStructural jumps to the next byte that can change the
// state, and a closer of the kind the innermost container did NOT open is
// stepped over exactly as those two do.
func (s *ValueScanner) container(chunk []byte, i int, final bool) (int, bool, error) {
	for {
		if s.inString {
			end, closed := s.stringBody(chunk, i)
			if !closed {
				if final {
					return end, false, ErrTruncated
				}
				return end, false, nil
			}
			s.inString = false
			i = end
		}
		if i >= len(chunk) {
			break
		}
		i += indexStructural(chunk[i:])
		if i >= len(chunk) {
			break
		}
		switch chunk[i] {
		case '{', '[':
			if s.depth == MaxDepth {
				return i, false, ErrMaxDepth
			}
			s.push(chunk[i] == '{')
			i++
		case '}', ']':
			if s.isObject() == (chunk[i] == '}') {
				s.depth--
				if s.depth == 0 {
					s.state = scanDone
					return i + 1, true, nil
				}
			}
			i++
		default: // '"'
			s.inString = true
			i++
		}
	}
	if final {
		return i, false, ErrTruncated
	}
	return i, false, nil
}

// stringBody advances past the body of the string whose opening quote is
// already consumed, returning the index just past its closing quote and
// whether it was found within chunk. The escape state carries across chunks:
// a '\' at the end of one makes the first byte of the next inert, which is
// what makes a split \uXXXX invisible to the caller (its four hex digits are
// neither a quote nor a backslash).
func (s *ValueScanner) stringBody(chunk []byte, i int) (int, bool) {
	for {
		if s.escaped {
			if i >= len(chunk) {
				return i, false
			}
			s.escaped = false
			i++
		}
		e := indexCloseOrEscapeAt(chunk, i)
		if uint(e) >= uint(len(chunk)) {
			return len(chunk), false
		}
		if chunk[e] == '"' {
			return e + 1, true
		}
		s.escaped = true
		i = e + 1
	}
}

// push opens a container of the given kind one level deeper.
func (s *ValueScanner) push(object bool) {
	d := s.depth
	if object {
		s.kinds[d/64] |= 1 << (uint(d) % 64)
	} else {
		s.kinds[d/64] &^= 1 << (uint(d) % 64)
	}
	s.depth = d + 1
}

// isObject reports whether the innermost open container is an object.
func (s *ValueScanner) isObject() bool {
	d := s.depth - 1
	return s.kinds[d/64]&(1<<(uint(d)%64)) != 0
}

// isNumberByte is skipNumber's accept set.
func isNumberByte(c byte) bool {
	return (c >= '0' && c <= '9') || c == '.' || c == 'e' || c == 'E' || c == '+' || c == '-'
}
