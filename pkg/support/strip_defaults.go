package support

import "bytes"

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
	// members (and their whitespace) are removed.
	PreserveWhitespace
)

// StripDefaults copies the JSON document in input to output, dropping every object
// member whose value is a "default" — byte-equal to one of defaults, compared
// against the bare token: the unquoted contents for a string value, the literal
// token for a number/keyword — and then dropping any object or array that this
// leaves empty. Empty values are not special-cased: list an empty entry ("") in
// defaults to drop them. A member is kept despite a default value when its
// (unquoted) key is byte-equal to one of keep.
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
	start := SkipWS(input, 0)
	write := 0
	if ws == PreserveWhitespace {
		write = copy(output, input[:start])
	}
	read, write := s.handle(start, write)
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

// handle strips the value beginning at in[read], appending the kept bytes at
// out[write], and returns the read and write offsets just past it. A value that
// strips away to nothing leaves write unchanged, which is how callers detect a
// dropped member or an emptied container.
func (s *stripper) handle(read, write int) (int, int) {
	in, out := s.in, s.out
	// compact: don't scan for whitespace (input asserted to have none).
	// preserve: copy surviving inter-token whitespace through to the output.
	compact := s.ws == AssumeCompact
	preserve := s.ws == PreserveWhitespace
	dataLen := len(in)
	read = skipWSc(in, read, compact)
	if read == dataLen {
		return read, write
	}
	// eject copies the unconsumed remainder of input through verbatim, the
	// best-effort response to a byte the walk cannot interpret.
	eject := func() (int, int) {
		return dataLen, write + copy(out[write:], in[read:])
	}

	switch in[read] {
	case '{':
		read++
		if read != dataLen && in[read] == '}' {
			return read + 1, write
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
			wsStart := read
			read = skipWSc(in, read, compact)
			if read >= dataLen || in[read] != '"' {
				return eject()
			}
			keyStart := read
			keyEnd, err := skipString(in, read)
			if err != nil {
				return eject()
			}
			read = skipWSc(in, keyEnd, compact)
			colonPos := read
			if read == dataLen || in[read] != ':' {
				// Missing ':' — copy the key, then eject.
				write += copy(out[write:], in[keyStart:keyEnd])
				return eject()
			}
			read++
			tmpRead := read

			read = skipWSc(in, read, compact)
			valueEmpty := true
			if read < dataLen {
				switch in[read] {
				case '"':
					valEnd, err := skipString(in, read)
					if err != nil {
						// Bad string: eject from the original position.
						read, write = tmpRead, localStartWrite
						return eject()
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
					if peek := skipWSc(in, read+1, compact); peek < dataLen && in[peek] == closeBrace {
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
					read, write = s.handle(read, write)
					if tmpWrite != write {
						valueEmpty = false
					} else {
						write = localStartWrite
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
					write = s.emitField(write, wsStart, keyStart, keyEnd, colonPos, tmpRead, read)
					written = true
				}
			} else {
				written = true
			}
			wsBeforeDelim := read
			read = skipWSc(in, read, compact)
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
		if read != dataLen && in[read] == ']' {
			return read + 1, write
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
			read = skipWSc(in, read, compact)
			if preserve {
				write += copy(out[write:], in[wsStart:read]) // element's leading whitespace
			}
			tmpWrite := write
			read, write = s.handle(read, write)
			if tmpWrite == write {
				write = localStartWrite // element stripped away; rewind
			} else {
				written = true
			}
			wsBeforeDelim := read
			read = skipWSc(in, read, compact)
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
		send, err := skipString(in, read)
		if err != nil {
			return eject()
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
