package json

import (
	"bytes"

	"github.com/JohanLindvall/lightning/pkg/unstable"
)

// WhitespaceMode selects how StripDefaults treats inter-token whitespace in the
// input. The zero value, RemoveWhitespace, is the safe default: it tolerates any
// whitespace and produces compact output.
type WhitespaceMode int

const (
	// RemoveWhitespace scans past inter-token whitespace and drops it: the output
	// is compact regardless of how the input was formatted.
	RemoveWhitespace WhitespaceMode = iota
	// AssumeCompact asserts the input has no inter-token whitespace (the form
	// compact serializers emit) and skips the whitespace scans entirely — faster,
	// but it misreads input that actually contains whitespace. Output is compact.
	AssumeCompact
	// PreserveWhitespace keeps the input's inter-token whitespace around surviving
	// content, so a pretty-printed document stays pretty-printed; only the dropped
	// members (and their whitespace) are removed. One kind of surviving whitespace
	// is not kept: the run between a value and the ',' that follows it, which is
	// dropped with the separator it precedes ({"a":1 , "b":2} becomes
	// {"a":1, "b":2}). Whitespace before a closing '}' or ']' is preserved.
	PreserveWhitespace
)

// StripDefaults copies the JSON document in input to output, dropping every
// "default" value — one byte-equal to an entry of defaults, compared against the
// bare token: the unquoted contents for a string value, the literal token for a
// number/keyword. An object member with a default value is dropped whole (key
// included); an array element that is a default is dropped from the array, which
// reindexes it ([0,1,0,2] with "0" listed becomes [1,2]). Any object or array left
// empty by that pass is then dropped from its own parent, cascading upward until
// the document itself can come out empty.
//
// A container that is *already* empty in the input is dropped the same way, and
// that is the one rule that does not consult defaults: {"a":{}} strips to nothing
// even for a nil defaults (keep still rescues such a member, as it does any
// other). No other value is special-cased — in particular the empty string is
// dropped only when an empty entry ("") is listed in defaults.
//
// A member is kept despite a default value when its (unquoted) key is byte-equal
// to one of keep; a kept member keeps its original value — including the members
// of a container value that stripped away to nothing — reformatted according to
// ws. keep names object keys, so it has no effect on array elements.
//
// output is filled from the front and the populated prefix is returned; input is
// not modified. StripDefaults never lengthens the document, so output needs room
// for at most len(input) bytes — it is grown (allocated) only when cap(output) is
// smaller, mirroring UnescapeStringInto. Pass output == input[:0] to strip in
// place. The returned slice aliases whichever buffer was written.
//
// StripDefaults is best effort and forgiving of malformed input: on the first
// byte it cannot make sense of it copies the remainder of input through verbatim
// rather than failing.
//
// String values keep their surrounding quotes and escapes, scalars keep their
// literal token. ws controls inter-token whitespace handling — see WhitespaceMode.
func StripDefaults(input, output []byte, defaults, keep [][]byte, ws WhitespaceMode) []byte {
	if cap(output) < len(input) {
		output = make([]byte, len(input))
	} else {
		output = output[:len(input)]
	}
	s := stripper{
		in:          input,
		out:         output,
		defaults:    defaults,
		keep:        keep,
		defaultLens: lenSet(defaults),
		keepLens:    lenSet(keep),
		ws:          ws,
	}
	// Whitespace at the document start is tolerated (and preserved when asked) for
	// every mode; handle's own skips honor ws thereafter.
	start := unstable.SkipWS(input, 0)
	write := 0
	if ws == PreserveWhitespace {
		write = copy(output, input[:start])
	}
	read, write := s.handle(start, write, 0)
	if ws == PreserveWhitespace {
		write += copy(output[write:], input[read:]) // trailing whitespace
	}
	return output[:write]
}

// lenSet returns a bitmask with bit n set for every entry of items whose length
// is n. Lengths >= 64 set the top bit, so a too-long token never short-circuits
// (it falls through to the scan, which still compares correctly). Used to skip
// the list scan for tokens whose length matches no candidate — the common case,
// since most JSON values and keys are longer than any default or kept key.
func lenSet(items [][]byte) uint64 {
	var m uint64
	for _, it := range items {
		n := len(it)
		if n >= 64 {
			n = 63
		}
		m |= uint64(1) << n
	}
	return m
}

// hasLen reports whether n could match an entry summarized by mask m.
func hasLen(m uint64, n int) bool {
	if n >= 64 {
		n = 63
	}
	return m&(uint64(1)<<n) != 0
}

// stripper carries the read buffer (in), write buffer (out) and the caller's
// default-value and keep-key lists through the recursive walk in handle.
// defaultLens/keepLens summarize the candidate lengths (see lenSet) so a token
// of non-matching length skips the scan.
type stripper struct {
	in          []byte
	out         []byte
	defaults    [][]byte
	keep        [][]byte
	defaultLens uint64
	keepLens    uint64
	ws          WhitespaceMode
}

// isDefault reports whether a scalar value should be dropped: one byte-equal to a
// caller-supplied default. The empty token counts only when "" is among defaults
// (the length pre-filter rejects it otherwise), so empties are opt-in.
func (s *stripper) isDefault(value []byte) bool {
	n := len(value)

	if !hasLen(s.defaultLens, n) {
		return false
	}

	if n == 0 {
		return true
	}

	for _, d := range s.defaults {
		if len(d) == n && bytes.Equal(value, d) {
			return true
		}
	}
	return false
}

// keepKey reports whether a member with a default value should be kept anyway.
// key is the raw key token including its surrounding quotes; the comparison is
// against the bare name, so keep entries are unquoted (e.g. []byte("WallTimeMs")).
func (s *stripper) keepKey(key []byte) bool {
	if len(key) >= 2 {
		key = key[1 : len(key)-1]
	}
	n := len(key)
	if !hasLen(s.keepLens, n) {
		return false
	}
	for _, k := range s.keep {
		if len(k) == n && bytes.Equal(key, k) {
			return true
		}
	}
	return false
}

// emitField appends a kept object member to out at write and returns the new
// write. In PreserveWhitespace mode it copies the verbatim span in[wsStart:valEnd]
// — leading whitespace, key, the ':' and any whitespace around it, and the value
// — so the input's formatting survives. Otherwise it emits compact "key":value:
// one copy of the whole span when key, colon and value are already adjacent (the
// usual compact case), else key + a synthesized ':' + value.
func (s *stripper) emitField(write, wsStart, keyStart, keyEnd, colonPos, valStart, valEnd int) int {
	in, out := s.in, s.out
	if s.ws == PreserveWhitespace {
		return write + copy(out[write:], in[wsStart:valEnd])
	}
	if colonPos == keyEnd && valStart == keyEnd+1 {
		return write + copy(out[write:], in[keyStart:valEnd])
	}
	write += copy(out[write:], in[keyStart:keyEnd])
	out[write] = ':'
	write++
	return write + copy(out[write:], in[valStart:valEnd])
}

// emitKept appends a member the keep list rescues: one whose value is a default,
// or a container that stripped away to nothing. Its original value bytes are
// re-emitted, so valStart..valEnd is a span of *input* the walk never re-examined
// — and in RemoveWhitespace mode that span may still hold the inter-token
// whitespace the mode promises to remove ({"b":{ "x" : 0 }} kept "b" verbatim,
// interior spaces and all, which also broke the compact(preserve) == remove
// oracle). Only this path can produce such a span: every other emitField caller
// passes a single scalar or string token. So the whitespace is taken out here,
// off the hot path — the leading run after the ':' by advancing valStart, and a
// container value's interior by copying it through compactValue.
//
// PreserveWhitespace re-emits verbatim by definition, and AssumeCompact asserts
// there is no whitespace to remove (scanning for it would contradict the mode),
// so both fall through to emitField unchanged.
//
// This reads the member's original bytes out of s.in, the same buffer emitField
// re-reads: if that source ever changes (an in-place run overwrites the span
// before this path re-reads it, so a snapshot would have to be read instead),
// it has to change in both.
func (s *stripper) emitKept(write, wsStart, keyStart, keyEnd, colonPos, valStart, valEnd int) int {
	if s.ws == RemoveWhitespace {
		in := s.in
		for valStart < valEnd && in[valStart] <= ' ' {
			valStart++
		}
		if valStart < valEnd && (in[valStart] == '{' || in[valStart] == '[') {
			out := s.out
			write += copy(out[write:], in[keyStart:keyEnd])
			out[write] = ':'
			write++
			return write + compactValue(out[write:], in[valStart:valEnd])
		}
	}
	return s.emitField(write, wsStart, keyStart, keyEnd, colonPos, valStart, valEnd)
}

// compactValue copies the JSON value src into dst with its inter-token whitespace
// removed, returning the number of bytes written. Whitespace is the same
// deliberately lenient "<= ' '" the rest of this walker uses, and a string token
// is copied whole so bytes inside it are never touched. Output can only be
// shorter than src, so dst — the tail of an output buffer sized to the whole
// input — always has room.
//
// A malformed string (unterminated) ends the scan: the rest of src is copied
// through verbatim, the same best-effort response handle's eject gives.
func compactValue(dst, src []byte) int {
	n := 0
	for i := 0; i < len(src); {
		switch c := src[i]; {
		case c <= ' ':
			i++
		case c == '"':
			end, err := unstable.SkipString(src, i)
			if err != nil {
				return n + copy(dst[n:], src[i:])
			}
			n += copy(dst[n:], src[i:end])
			i = end
		default:
			dst[n] = c
			n++
			i++
		}
	}
	return n
}

// handle strips the value beginning at in[read], appending the kept bytes at
// out[write], and returns the read and write offsets just past it. A value that
// strips away to nothing leaves write unchanged, which is how callers detect a
// dropped member or an emptied container.
//
// depth is the number of enclosing containers. handle recurses once per nesting
// level, so the bound below is what keeps a deeply nested document from
// exhausting the stack — a Go stack overflow being fatal and unrecoverable. Past
// unstable.MaxDepth it stops stripping and ejects, which is the same best-effort
// response this walker already gives malformed input.
func (s *stripper) handle(read, write, depth int) (int, int) {
	in, out := s.in, s.out
	// compact: don't scan for whitespace (input asserted to have none).
	// preserve: copy surviving inter-token whitespace through to the output.
	compact := s.ws == AssumeCompact
	preserve := s.ws == PreserveWhitespace
	dataLen := len(in)
	if !compact && read < len(in) && in[read] <= ' ' {
		read++
		if read < len(in) && in[read] <= ' ' {
			read = unstable.SkipWSRun(in, read+1)
		}
	}
	if read == dataLen {
		return read, write
	}
	// eject copies the unconsumed remainder of input through verbatim, the
	// best-effort response to a byte the walk cannot interpret.
	eject := func() (int, int) {
		return dataLen, write + copy(out[write:], in[read:])
	}
	if depth > unstable.MaxDepth {
		return eject()
	}

	switch in[read] {
	case '{':
		read++
		// Peek past inter-token whitespace for the close, exactly as the member-value
		// container branch below does. Without the peek "{ }" was not recognized as
		// empty, so the member loop hit '}' where it wanted a key and fell into
		// eject() — which by contract copies the *whole remaining input* through
		// verbatim, leaving every later member unstripped (and, in RemoveWhitespace
		// mode, unminified) behind output that still looks like valid JSON. read is
		// deliberately left where it was when the container is not empty: the member
		// loop needs the leading whitespace to preserve it. (The skip is expanded
		// here rather than factored into a helper for the reason CLAUDE.md records:
		// a helper wrapping the two-compare fast path plus SkipWSRun costs more than
		// the inliner's budget, so it would cost a call frame per container.)
		peekObj := read
		if !compact && peekObj < dataLen && in[peekObj] <= ' ' {
			peekObj++
			if peekObj < dataLen && in[peekObj] <= ' ' {
				peekObj = unstable.SkipWSRun(in, peekObj+1)
			}
		}
		if peekObj < dataLen && in[peekObj] == '}' {
			return peekObj + 1, write
		}
		startWrite := write
		out[write] = '{'
		write++
		written := false
		for {
			localStartWrite := write
			if written {
				out[write] = ','
				write++
			}
			// Where the member's own bytes begin — just past the separator
			// comma. The container value branch below rewinds an emptied value
			// here, not to localStartWrite: the keepKey path may still emit the
			// member, and the comma belongs with it (rewinding past the comma
			// produced {"a":1"b":{...}} — invalid JSON). Dropping the member
			// entirely still rewinds to localStartWrite, comma included.
			postComma := write
			wsStart := read
			if !compact && read < len(in) && in[read] <= ' ' {
				read++
				if read < len(in) && in[read] <= ' ' {
					read = unstable.SkipWSRun(in, read+1)
				}
			}
			if read >= dataLen || in[read] != '"' {
				return eject()
			}
			keyStart := read
			// Inline the no-escape close-quote scan (one SIMD pass) so the common
			// unescaped key skips SkipString's non-inlinable call frame; escaped or
			// truncated keys fall back to SkipString. Same trick on the value reads below.
			krest := in[read+1:]
			kk := unstable.IndexCloseOrEscape(krest)
			keyEnd := read + kk + 2
			if kk >= len(krest) || krest[kk] != '"' {
				var err error
				if keyEnd, err = unstable.SkipString(in, read); err != nil {
					return eject()
				}
			}
			read = keyEnd
			if !compact && read < len(in) && in[read] <= ' ' {
				read++
				if read < len(in) && in[read] <= ' ' {
					read = unstable.SkipWSRun(in, read+1)
				}
			}
			colonPos := read
			if read == dataLen || in[read] != ':' {
				// Missing ':' — copy the key, then eject.
				write += copy(out[write:], in[keyStart:keyEnd])
				return eject()
			}
			read++
			tmpRead := read

			if !compact && read < len(in) && in[read] <= ' ' {
				read++
				if read < len(in) && in[read] <= ' ' {
					read = unstable.SkipWSRun(in, read+1)
				}
			}
			valueEmpty := true
			if read < dataLen {
				switch in[read] {
				case '"':
					vrest := in[read+1:]
					vk := unstable.IndexCloseOrEscape(vrest)
					valEnd := read + vk + 2
					if vk >= len(vrest) || vrest[vk] != '"' {
						var err error
						if valEnd, err = unstable.SkipString(in, read); err != nil {
							// Bad string: eject from the original position.
							read, write = tmpRead, localStartWrite
							return eject()
						}
					}
					if !s.isDefault(in[read+1 : valEnd-1]) {
						valueEmpty = false
						write = s.emitField(write, wsStart, keyStart, keyEnd, colonPos, read, valEnd)
					}
					read = valEnd
				case '{', '[':
					closeBrace := byte('}')
					if in[read] == '[' {
						closeBrace = ']'
					}
					peek := read + 1
					if !compact && peek < dataLen && in[peek] <= ' ' {
						peek++
						if peek < dataLen && in[peek] <= ' ' {
							peek = unstable.SkipWSRun(in, peek+1)
						}
					}
					if peek < dataLen && in[peek] == closeBrace {
						read = peek + 1 // empty nested container — drop the member
						break
					}
					// Non-empty nested value: write key + colon (or the verbatim
					// "ws key : ws" prefix when preserving), then recurse.
					if preserve {
						write += copy(out[write:], in[wsStart:read])
					} else {
						write += copy(out[write:], in[keyStart:keyEnd])
						out[write] = ':'
						write++
					}
					tmpWrite := write
					read, write = s.handle(read, write, depth+1)
					if tmpWrite != write {
						valueEmpty = false
					} else {
						write = postComma // rewind the key+colon; keepKey may still emit
					}
				default:
					end := findDelimiter(in, read)
					if !s.isDefault(in[read:end]) {
						valueEmpty = false
						write = s.emitField(write, wsStart, keyStart, keyEnd, colonPos, read, end)
					}
					read = end
				}
			}
			if valueEmpty {
				if !s.keepKey(in[keyStart:keyEnd]) {
					write = localStartWrite // rewind: drop the member (and its whitespace/comma)
				} else {
					write = s.emitKept(write, wsStart, keyStart, keyEnd, colonPos, tmpRead, read)
					written = true
				}
			} else {
				written = true
			}
			wsBeforeDelim := read
			if !compact && read < len(in) && in[read] <= ' ' {
				read++
				if read < len(in) && in[read] <= ' ' {
					read = unstable.SkipWSRun(in, read+1)
				}
			}
			if read == dataLen {
				return eject()
			}
			switch in[read] {
			case ',':
				read++
			case '}':
				if write == startWrite+1 {
					return read + 1, startWrite // object emptied
				}
				if preserve {
					write += copy(out[write:], in[wsBeforeDelim:read])
				}
				out[write] = '}'
				write++
				return read + 1, write
			default:
				return eject()
			}
		}
	case '[':
		read++
		// The same whitespace peek as the object case above. "[ ]" reached the right
		// answer without it — the element loop's lone whitespace-only "element" is an
		// empty scalar token that emits nothing, so the array still came out empty
		// and was dropped — but only by a longer route through findDelimiter, and
		// only as long as that route holds. Deciding emptiness here makes the two
		// container cases agree by construction.
		peekArr := read
		if !compact && peekArr < dataLen && in[peekArr] <= ' ' {
			peekArr++
			if peekArr < dataLen && in[peekArr] <= ' ' {
				peekArr = unstable.SkipWSRun(in, peekArr+1)
			}
		}
		if peekArr < dataLen && in[peekArr] == ']' {
			return peekArr + 1, write
		}
		startWrite := write
		out[write] = '['
		write++
		written := false
		for {
			localStartWrite := write
			if written {
				out[write] = ','
				write++
			}
			wsStart := read
			if !compact && read < len(in) && in[read] <= ' ' {
				read++
				if read < len(in) && in[read] <= ' ' {
					read = unstable.SkipWSRun(in, read+1)
				}
			}
			if preserve {
				write += copy(out[write:], in[wsStart:read]) // element's leading whitespace
			}
			tmpWrite := write
			read, write = s.handle(read, write, depth+1)
			if tmpWrite == write {
				write = localStartWrite // element stripped away; rewind
			} else {
				written = true
			}
			wsBeforeDelim := read
			if !compact && read < len(in) && in[read] <= ' ' {
				read++
				if read < len(in) && in[read] <= ' ' {
					read = unstable.SkipWSRun(in, read+1)
				}
			}
			if read == dataLen {
				return eject()
			}
			switch in[read] {
			case ',':
				read++
			case ']':
				if write == startWrite+1 {
					return read + 1, startWrite // array emptied
				}
				if preserve {
					write += copy(out[write:], in[wsBeforeDelim:read])
				}
				out[write] = ']'
				write++
				return read + 1, write
			default:
				return eject()
			}
		}
	case '"':
		srest := in[read+1:]
		sk := unstable.IndexCloseOrEscape(srest)
		send := read + sk + 2
		if sk >= len(srest) || srest[sk] != '"' {
			var err error
			if send, err = unstable.SkipString(in, read); err != nil {
				return eject()
			}
		}
		if s.isDefault(in[read+1 : send-1]) {
			return send, write
		}
		return send, write + copy(out[write:], in[read:send])
	default:
		end := findDelimiter(in, read)
		if s.isDefault(in[read:end]) {
			return end, write
		}
		return end, write + copy(out[write:], in[read:end])
	}
}

// delimTable maps each byte to true if it terminates a JSON scalar value:
// whitespace (<= ' '), '{', '}', '[', ']' or ','.
var delimTable = func() (t [256]bool) {
	for c := 0; c <= ' '; c++ {
		t[c] = true
	}
	for _, c := range []byte{'{', '}', '[', ']', ','} {
		t[c] = true
	}
	return
}()

// findDelimiter returns the index of the first scalar-terminating byte at or
// after offset (per delimTable), or len(input) if none remains.
func findDelimiter(input []byte, offset int) int {
	for end := len(input); offset < end; offset++ {
		if delimTable[input[offset]] {
			break
		}
	}
	return offset
}
