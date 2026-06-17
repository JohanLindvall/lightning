package support

import "bytes"

// Cleanup copies the JSON document in input to output, dropping every object
// member whose value is a "default" — empty, or byte-equal to one of defaults
// (compared against the bare token: the unquoted contents for a string value,
// the literal token for a number/keyword) — and then dropping any object or
// array that this leaves empty. A member is kept despite a default value when
// its (unquoted) key is byte-equal to one of keep.
//
// output is filled from the front and the populated prefix is returned; input is
// not modified. Cleanup never lengthens the document, so output needs room for
// at most len(input) bytes — it is grown (allocated) only when cap(output) is
// smaller, mirroring UnescapeStringInto. Pass output == input[:0] to clean in
// place. The returned slice aliases whichever buffer was written.
//
// Cleanup is best effort and forgiving of malformed input: on the first byte it
// cannot make sense of it copies the remainder of input through verbatim rather
// than failing.
//
// Empty arrays and objects already present in input are dropped, as are members
// whose value becomes empty after cleaning; string values keep their surrounding
// quotes and escapes, scalars keep their literal token.
//
// Set compact when the input carries no whitespace between tokens — the form
// compact JSON serializers emit — to elide the inter-token whitespace skips,
// matching GetCompact/ObjectEachCompact. Leading whitespace at the document start
// is still tolerated; whitespace anywhere else makes compact misread the input.
func Cleanup(input, output []byte, defaults, keep [][]byte, compact bool) []byte {
	if cap(output) < len(input) {
		output = make([]byte, len(input))
	} else {
		output = output[:len(input)]
	}
	c := cleaner{in: input, out: output, defaults: defaults, keep: keep, compact: compact}
	c.defaultLens = lenSet(defaults)
	c.keepLens = lenSet(keep)
	// The unconditional skip here tolerates leading whitespace at the document
	// start even in compact mode; handle's own skips honor c.compact thereafter.
	_, write := c.handle(false, SkipWS(input, 0), 0)
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
		m |= uint64(1) << uint(n)
	}
	return m
}

// hasLen reports whether n could match an entry summarized by mask m.
func hasLen(m uint64, n int) bool {
	if n >= 64 {
		n = 63
	}
	return m&(uint64(1)<<uint(n)) != 0
}

// cleaner carries the read buffer (in), write buffer (out) and the caller's
// default-value and keep-key lists through the recursive walk in handle.
// defaultLens/keepLens summarize the candidate lengths (see lenSet) so a token
// of non-matching length skips the scan.
type cleaner struct {
	in          []byte
	out         []byte
	defaults    [][]byte
	keep        [][]byte
	defaultLens uint64
	keepLens    uint64
	compact     bool
}

// isDefault reports whether a scalar value should be dropped: an empty token, or
// one byte-equal to a caller-supplied default.
func (c *cleaner) isDefault(value []byte) bool {
	n := len(value)
	if n == 0 {
		return true
	}
	if !hasLen(c.defaultLens, n) {
		return false
	}
	for _, d := range c.defaults {
		if len(d) == n && bytes.Equal(value, d) {
			return true
		}
	}
	return false
}

// keepKey reports whether a member with a default value should be kept anyway.
// key is the raw key token including its surrounding quotes; the comparison is
// against the bare name, so keep entries are unquoted (e.g. []byte("WallTimeMs")).
func (c *cleaner) keepKey(key []byte) bool {
	if len(key) >= 2 {
		key = key[1 : len(key)-1]
	}
	n := len(key)
	if !hasLen(c.keepLens, n) {
		return false
	}
	for _, k := range c.keep {
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
func (c *cleaner) emitField(write, keyStart, keyEnd, colonPos, valStart, valEnd int) int {
	out, in := c.out, c.in
	if colonPos == keyEnd && valStart == keyEnd+1 {
		return write + copy(out[write:], in[keyStart:valEnd])
	}
	write += copy(out[write:], in[keyStart:keyEnd])
	out[write] = ':'
	write++
	return write + copy(out[write:], in[valStart:valEnd])
}

// handle cleans the value beginning at in[read], appending the cleaned bytes at
// out[write], and returns the read and write offsets just past it. A value that
// cleans away to nothing leaves write unchanged, which is how callers detect a
// dropped member or an emptied container.
func (c *cleaner) handle(key bool, read, write int) (int, int) {
	in, out, compact := c.in, c.out, c.compact
	datalen := len(in)
	read = skipWSc(in, read, compact)
	if read == datalen {
		return read, write
	}
	// eject copies the unconsumed remainder of input through verbatim, the
	// best-effort response to a byte the walk cannot interpret.
	eject := func() (int, int) {
		copy(out[write:], in[read:])
		write += datalen - read
		return datalen, write
	}

	switch in[read] {
	case '{':
		read++
		if read != datalen && in[read] == '}' {
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
			// Parse key position — keys are always quoted strings. skipString
			// honors backslash escapes, so a key like "a\"b" is spanned correctly.
			read = skipWSc(in, read, compact)
			if read >= datalen || in[read] != '"' {
				return eject()
			}
			keyStart := read
			keyEnd, err := skipString(in, read)
			if err != nil {
				return eject()
			}
			read = skipWSc(in, keyEnd, compact)
			colonPos := read
			if read == datalen || in[read] != ':' {
				// Missing ':' — copy key then eject.
				write += copy(out[write:], in[keyStart:keyEnd])
				return eject()
			}
			read++
			tmpRead := read

			// Parse value inline for common scalar types.
			read = skipWSc(in, read, compact)
			valueEmpty := true
			if read < datalen {
				switch in[read] {
				case '"':
					valEnd, err := skipString(in, read)
					if err != nil {
						// Bad string, eject from original position.
						read = tmpRead
						write = localStartWrite
						return eject()
					}
					evalue := in[read+1 : valEnd-1]
					if !c.isDefault(evalue) {
						valueEmpty = false
						write = c.emitField(write, keyStart, keyEnd, colonPos, read, valEnd)
					}
					read = valEnd
				case '{', '[':
					// Fast path: detect empty {} or [] without recursion.
					closeBrace := byte('}')
					if in[read] == '[' {
						closeBrace = ']'
					}
					peek := skipWSc(in, read+1, compact)
					if peek < datalen && in[peek] == closeBrace {
						// Empty nested structure — drop the field.
						read = peek + 1
						break
					}
					// For nested structures, copy key+colon then recurse.
					keyLen := keyEnd - keyStart
					copy(out[write:], in[keyStart:keyEnd])
					write += keyLen
					out[write] = ':'
					write++
					tmpWrite := write
					read, write = c.handle(false, tmpRead, write)
					if tmpWrite != write {
						valueEmpty = false
					} else {
						write = localStartWrite
					}
				default:
					end := findDelimiter(in, read)
					if !c.isDefault(in[read:end]) {
						valueEmpty = false
						write = c.emitField(write, keyStart, keyEnd, colonPos, read, end)
					}
					read = end
				}
			}
			if valueEmpty {
				if c.keepKey(in[keyStart:keyEnd]) {
					// Keep a default-valued member: key + colon + its raw value.
					write = c.emitField(write, keyStart, keyEnd, colonPos, tmpRead, read)
					written = true
				} else {
					write = localStartWrite // rewind to before comma
				}
			} else {
				written = true
			}
			read = skipWSc(in, read, compact)
			if read == datalen {
				return eject()
			}
			switch in[read] {
			case ',':
				read++
			case '}':
				read++
				if write == startWrite+1 {
					// Empty
					return read, startWrite
				}
				out[write] = '}'
				write++
				return read, write
			default:
				// Bad delimiter
				return eject()
			}
		}
	case '[':
		read++
		if read != datalen && in[read] == ']' {
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
			read, write = c.handle(false, read, write)
			if tmpWrite == write {
				write = localStartWrite // rewind
			} else {
				written = true
			}
			read = skipWSc(in, read, compact)
			if read == datalen {
				// Bad JSON, no end.
				return eject()
			}
			switch in[read] {
			case ',':
				read++
			case ']':
				read++
				if write == startWrite+1 {
					// Empty
					return read, startWrite
				}
				out[write] = ']'
				write++
				return read, write
			default:
				// Bad delimiter
				return eject()
			}
		}
	case '"':
		send, err := skipString(in, read)
		if err != nil {
			// Bad string
			return eject()
		}
		if !key {
			evalue := in[read+1 : send-1]
			if c.isDefault(evalue) {
				return send, write
			}
		}

		value := in[read:send]
		copy(out[write:], value)
		write += len(value)
		return send, write
	default:
		end := findDelimiter(in, read)
		value := in[read:end]
		if c.isDefault(value) {
			return end, write
		}
		copy(out[write:], value)
		write += len(value)
		return end, write
	}
}

// delimTable maps each byte to true if it terminates a JSON scalar value:
// whitespace (<= ' '), '{', '}', '[', ']' or ','.
var delimTable = func() (t [256]bool) {
	for c := 0; c <= ' '; c++ {
		t[c] = true
	}
	t['{'] = true
	t['}'] = true
	t['['] = true
	t[']'] = true
	t[','] = true
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
