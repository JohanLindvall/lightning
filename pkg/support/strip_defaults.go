package support

import "bytes"

// StripDefaults copies the JSON document in input to output, dropping every object
// member whose value is a "default" — empty, or byte-equal to one of defaults
// (compared against the bare token: the unquoted contents for a string value,
// the literal token for a number/keyword) — and then dropping any object or
// array that this leaves empty. A member is kept despite a default value when
// its (unquoted) key is byte-equal to one of keep.
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
// Empty arrays and objects already present in input are dropped, as are members
// whose value becomes empty after stripping; string values keep their surrounding
// quotes and escapes, scalars keep their literal token.
//
// Set compact when the input carries no whitespace between tokens — the form
// compact JSON serializers emit — to elide the inter-token whitespace skips,
// matching GetCompact/ObjectEachCompact. Leading whitespace at the document start
// is still tolerated; whitespace anywhere else makes compact misread the input.
func StripDefaults(input, output []byte, defaults, keep [][]byte, compact bool) []byte {
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
		compact:     compact,
	}
	// The unconditional skip here tolerates leading whitespace at the document
	// start even in compact mode; handle's own skips honor the compact flag.
	_, write := s.handle(SkipWS(input, 0), 0)
	return output[:write]
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
	compact     bool
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

// isDefault reports whether a scalar value should be dropped: an empty token, or
// one byte-equal to a caller-supplied default.
func (s *stripper) isDefault(value []byte) bool {
	n := len(value)
	if n == 0 {
		return true
	}
	if !hasLen(s.defaultLens, n) {
		return false
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

// emitField appends the member in[keyStart:keyEnd] + ':' + in[valStart:valEnd] to
// out at write and returns the new write. colonPos is where the ':' sits in the
// input; when the key, colon and value are adjacent there (no whitespace between
// them — always so for compact input) the whole "key":value span is moved in one
// copy, otherwise the key and value are copied separately with a synthesized ':'.
func (s *stripper) emitField(write, keyStart, keyEnd, colonPos, valStart, valEnd int) int {
	in, out := s.in, s.out
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
	in, out, compact := s.in, s.out, s.compact
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
			// Parse the key — always a quoted string.
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
						write = s.emitField(write, keyStart, keyEnd, colonPos, read, valEnd)
					}
					read = valEnd
				case '{', '[':
					// Detect an empty {} or [] without recursing.
					closeBrace := byte('}')
					if in[read] == '[' {
						closeBrace = ']'
					}
					if peek := skipWSc(in, read+1, compact); peek < dataLen && in[peek] == closeBrace {
						read = peek + 1
						break
					}
					// Non-empty nested value: write key + colon, then recurse.
					write += copy(out[write:], in[keyStart:keyEnd])
					out[write] = ':'
					write++
					tmpWrite := write
					read, write = s.handle(tmpRead, write)
					if tmpWrite != write {
						valueEmpty = false
					} else {
						write = localStartWrite
					}
				default:
					end := findDelimiter(in, read)
					if !s.isDefault(in[read:end]) {
						valueEmpty = false
						write = s.emitField(write, keyStart, keyEnd, colonPos, read, end)
					}
					read = end
				}
			}
			if valueEmpty {
				if !s.keepKey(in[keyStart:keyEnd]) {
					write = localStartWrite // rewind to before the comma
				} else {
					// Keep a default-valued member: key + colon + its raw value.
					write = s.emitField(write, keyStart, keyEnd, colonPos, tmpRead, read)
					written = true
				}
			} else {
				written = true
			}
			read = skipWSc(in, read, compact)
			if read == dataLen {
				return eject()
			}
			switch in[read] {
			case ',':
				read++
			case '}':
				read++
				if write == startWrite+1 {
					return read, startWrite // object emptied
				}
				out[write] = '}'
				write++
				return read, write
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
			tmpWrite := write
			read, write = s.handle(read, write)
			if tmpWrite == write {
				write = localStartWrite // element stripped away; rewind
			} else {
				written = true
			}
			read = skipWSc(in, read, compact)
			if read == dataLen {
				return eject()
			}
			switch in[read] {
			case ',':
				read++
			case ']':
				read++
				if write == startWrite+1 {
					return read, startWrite // array emptied
				}
				out[write] = ']'
				write++
				return read, write
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
