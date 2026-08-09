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
// place: the walk never writes past its own read cursor, so it only ever
// overwrites bytes it has already consumed, and the result is byte-identical to
// the one a separate output buffer produces. The returned slice aliases
// whichever buffer was written.
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
// of non-matching length skips the scan. scratch holds the one member shape that
// has to be re-read after out has been written over it — see handle's container
// case; it stays nil unless such a member appears, so the common walk allocates
// nothing.
type stripper struct {
	in          []byte
	out         []byte
	defaults    [][]byte
	keep        [][]byte
	defaultLens uint64
	keepLens    uint64
	ws          WhitespaceMode
	scratch     []byte // last: the fields above keep the offsets the walk loads from
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

// emitFieldSnap is emitField reading the member's bytes from the snapshot in
// s.scratch instead of from in, where scratch[i-base] is in[i]. It re-emits a
// member handle has already written over; nothing else needs it.
//
// It is a second function rather than a src/base pair threaded through emitField
// because the parameterized single form costs 88 against the inliner's budget of
// 80: emitField would stop inlining and every kept member's emit — the hot one —
// would pay a call frame to serve this rare one. This copy is out of line (cost
// 93) and belongs there. Same reasoning as the whitespace-skip blocks this file
// expands by hand rather than factoring into a helper.
func (s *stripper) emitFieldSnap(base, write, wsStart, keyStart, keyEnd, colonPos, valStart, valEnd int) int {
	snap, out := s.scratch, s.out
	if s.ws == PreserveWhitespace {
		return write + copy(out[write:], snap[wsStart-base:valEnd-base])
	}
	if colonPos == keyEnd && valStart == keyEnd+1 {
		return write + copy(out[write:], snap[keyStart-base:valEnd-base])
	}
	write += copy(out[write:], snap[keyStart-base:keyEnd-base])
	out[write] = ':'
	write++
	return write + copy(out[write:], snap[valStart-base:valEnd-base])
}

// The two non-offset values of handle's per-member snapBase (see there): no early
// keep decision was taken, and one was taken and came out negative.
const (
	snapNone    = -1
	snapDropped = -2
)

// snapshotMember takes the early drop-or-keep decision for a member whose value
// is the container at in[read], and, when the key is kept, the copy of the
// member's original bytes that handle will need if that container strips to
// nothing. It returns the snapBase handle carries: snapDropped, or the base
// offset (wsStart) of the copy in s.scratch.
//
// It is deliberately out of line. The work is per container member and reached
// through a cold branch, while handle is a large function whose object-member
// loop is the hot one in this package's profile — expanding this into it grew the
// loop's body and cost the pretty-printed benchmark a few percent for nothing.
func (s *stripper) snapshotMember(read, wsStart, keyStart, keyEnd int) int {
	in := s.in
	if !s.keepKey(in[keyStart:keyEnd]) {
		return snapDropped
	}
	// The copy has to reach as far as the walk will, and the walk stops at the
	// end of the value — which SkipValue finds over the same bytes with the same
	// string handling. Malformed input is the one place the two can disagree, and
	// there the walk may run to the end of the document, so cover that far.
	snapEnd, err := unstable.SkipValue(in, read)
	if err != nil {
		snapEnd = len(in)
	}
	s.scratch = append(s.scratch[:0], in[wsStart:snapEnd]...)
	return wsStart
}

// finishEarly finishes a member whose value was a container that stripped to
// nothing and whose keep decision was therefore taken early, in snapshotMember:
// it drops the member, or re-emits it, and reports whether it emitted anything.
// Out of line for the same reason snapshotMember is — both arms are cold, and
// expanding them into handle's member loop only lengthens it.
func (s *stripper) finishEarly(snapBase, write, localStartWrite, wsStart, keyStart, keyEnd, colonPos, valStart, valEnd int) (int, bool) {
	switch {
	case snapBase == snapDropped:
		return localStartWrite, false // rewind: drop the member (and its whitespace/comma)
	case valEnd-snapBase <= len(s.scratch):
		// The kept member, re-emitted from the snapshot: its own bytes in the
		// input have been written over by the speculative write and the
		// recursion whenever output aliases input.
		return s.emitFieldSnap(snapBase, write, wsStart, keyStart, keyEnd, colonPos, valStart, valEnd), true
	default:
		// The walk ran past the end SkipValue found over the same bytes, so the
		// snapshot stops short of the member. That takes malformed input — the
		// two agree everywhere else — and leaves the input as the only source;
		// it is the intact one whenever output is a buffer of its own.
		return s.emitField(write, wsStart, keyStart, keyEnd, colonPos, valStart, valEnd), true
	}
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
			// A container value is the one member shape whose bytes reach out
			// before the member's fate is known, so that branch has to decide
			// early and hand the answer down to the drop-or-keep block below;
			// see the '{', '[' case. One int carries it, to keep this loop's
			// live set where it was: snapNone (the ordinary path — nothing has
			// been written for this member, so that block can still read the key
			// and the value out of the input), snapDropped (decided early, not a
			// kept key), or the base offset in s.scratch of the snapshot of the
			// member's original bytes (decided early, kept).
			snapBase := snapNone
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
					// Everything below this point writes speculatively: the key
					// and colon (or the preserved prefix), then whatever the
					// recursion emits, all rewound if the value strips to
					// nothing. out may alias in — StripDefaults documents
					// output == input[:0] — so those writes land on this
					// member's own bytes, and both things the drop-or-keep
					// decision below needs would then be read back clobbered:
					// the key (silently flipping the decision) and the original
					// span a kept member is re-emitted from. Take them here,
					// while the input is certainly intact. The keep test is a
					// pure function of the key, so it can simply move; the bytes
					// need a copy, taken only when the key is kept so that the
					// ordinary member costs nothing and allocates nothing.
					//
					// One scratch serves every nesting level: a deeper frame
					// overwrites it only by snapshotting a kept member of its
					// own, and such a member is then always emitted (stripped or
					// verbatim), which leaves every container enclosing it
					// non-empty — so no outer frame ever reaches back for a
					// snapshot a deeper one replaced.
					//
					// Only a kept key needs any of that, and only a key whose
					// length matches a keep entry can be kept — keepKey's own
					// length pre-filter, which is exact in the negative
					// direction. Leading with it keeps the ordinary container
					// member (empty keep list, or no candidate of that length)
					// at one compare instead of a call; a container-dense
					// document has as many of those as it has members, and a
					// call each measured ~4% on the pretty benchmark. The key
					// span is the quotes plus the name, and both key reads above
					// land past a closing quote, so it is never below 2 and the
					// subtraction never hands hasLen a negative shift.
					snapBase = snapDropped
					if hasLen(s.keepLens, keyEnd-keyStart-2) {
						snapBase = s.snapshotMember(read, wsStart, keyStart, keyEnd)
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
				switch {
				case snapBase != snapNone:
					// Decided early because the value was a container: both
					// outcomes live in finishEarly, out of line. written only
					// ever goes true, so a dropped member must not clear it.
					var emitted bool
					write, emitted = s.finishEarly(snapBase, write, localStartWrite, wsStart, keyStart, keyEnd, colonPos, tmpRead, read)
					written = written || emitted
				case s.keepKey(in[keyStart:keyEnd]):
					write = s.emitField(write, wsStart, keyStart, keyEnd, colonPos, tmpRead, read)
					written = true
				default:
					write = localStartWrite // rewind: drop the member (and its whitespace/comma)
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
