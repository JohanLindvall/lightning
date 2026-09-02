package unstable

import (
	"strconv"
	"time"
)

// ReadKey reads a JSON object key (a quoted string) at data[i] without
// allocating. Keys are assumed not to contain backslash escapes; if they do,
// the slow path is taken.
func ReadKey(data []byte, i int) (string, int, error) {
	if uint(i) >= uint(len(data)) || data[i] != '"' {
		return "", i, ErrInvalidJSON
	}
	i++
	e := indexCloseOrEscapeAt(data, i)
	if uint(e) >= uint(len(data)) {
		return "", len(data), ErrTruncated
	}
	if data[e] == '\\' {
		return decodeStringEscaped(data, i, e)
	}
	return unsafeStr(data[i:e]), e + 1, nil
}

// ReadStringOrNull reads a JSON string (or null) at data[i], copying the bytes
// into a fresh string.
//
// String bytes are returned verbatim: this decodes escapes but never inspects
// the UTF-8 of the literal runs between them, so raw invalid UTF-8 in the input
// reaches the Go string unchanged ("a\xffb" stays "a\xffb"), where
// encoding/json coerces the same input to "a\uFFFDb". Passing the bytes
// through is the cheaper and the more faithful of the two — no re-encode, and a
// caller that cares can run utf8.ValidString — but it does mean a decoded
// string is not guaranteed to be well-formed UTF-8. Every string path here
// shares the property because they share these scanners —
// ReadStringNoCopyOrNull, ReadStringDestructiveOrNull, ReadKey, the dynamic
// DecodeValue, and decodeEscaped's literal runs — and
// TestStringsPassInvalidUTF8Through pins it across all of them. Only the
// \uXXXX decoder normalizes: an unpaired surrogate escape becomes U+FFFD,
// matching encoding/json, since there is no other way to encode it.
func ReadStringOrNull(data []byte, i int) (string, int, error) {
	if uint(i) >= uint(len(data)) {
		return "", i, ErrTruncated
	}
	if data[i] == 'n' {
		end, err := ExpectNull(data, i)
		return "", end, err
	}
	if data[i] != '"' {
		return "", i, ErrInvalidJSON
	}
	i++
	e := indexCloseOrEscapeAt(data, i)
	if uint(e) >= uint(len(data)) {
		return "", len(data), ErrTruncated
	}
	if data[e] == '\\' {
		return decodeStringEscaped(data, i, e)
	}
	return string(data[i:e]), e + 1, nil
}

// ReadStringNoCopyOrNull is like ReadStringOrNull but, for strings without
// escapes, returns a string that aliases data rather than copying it, so the
// caller must keep data unchanged while the string is in use. Strings
// containing escapes still allocate, since they cannot be represented as a
// slice of the input.
func ReadStringNoCopyOrNull(data []byte, i int) (string, int, error) {
	if uint(i) >= uint(len(data)) {
		return "", i, ErrTruncated
	}
	if data[i] == 'n' {
		end, err := ExpectNull(data, i)
		return "", end, err
	}
	if data[i] != '"' {
		return "", i, ErrInvalidJSON
	}
	i++
	e := indexCloseOrEscapeAt(data, i)
	if uint(e) >= uint(len(data)) {
		return "", len(data), ErrTruncated
	}
	if data[e] == '\\' {
		return decodeStringEscaped(data, i, e)
	}
	return unsafeStr(data[i:e]), e + 1, nil
}

// ReadStringDestructiveOrNull is ReadStringNoCopyOrNull but, for a string that
// contains escapes, unescapes it *in place* — overwriting the escaped bytes of
// data with the decoded ones — instead of allocating a scratch buffer, and aliases
// the result. The unescaped form is never longer than the escaped body, so it fits
// within the body's bytes; the rest of the body is left as overwritten garbage and
// the closing quote (which the write cursor never reaches) still bounds the value.
// This DESTROYS the input document: the bytes of every escaped string are clobbered
// and any other alias into the same region (an overlapping nocopy value) is
// invalidated. It is the //lightning:destructive counterpart of the nocopy reader,
// for callers that own the buffer and discard it after decoding. Escape-free strings
// alias the input unchanged, exactly like the nocopy reader.
func ReadStringDestructiveOrNull(data []byte, i int) (string, int, error) {
	if uint(i) >= uint(len(data)) {
		return "", i, ErrTruncated
	}
	if data[i] == 'n' {
		end, err := ExpectNull(data, i)
		return "", end, err
	}
	if data[i] != '"' {
		return "", i, ErrInvalidJSON
	}
	i++
	e := indexCloseOrEscapeAt(data, i)
	if uint(e) >= uint(len(data)) {
		return "", len(data), ErrTruncated
	}
	if data[e] == '\\' {
		// Decode into data starting at the body offset i: buf aliases data[i:] with
		// length 0 and cap to the document end, so decodeEscaped's appends write
		// through into data. The write cursor trails the read cursor (unescaping only
		// shrinks), so it never clobbers a byte not yet consumed, and cap is large
		// enough that append never reallocates away from data.
		buf := data[i:i:len(data)]
		return decodeEscaped(buf, data, i, e, true)
	}
	return unsafeStr(data[i:e]), e + 1, nil
}

// ReadInt64OrNull reads a JSON integer (or null) at data[i]. Fractional and
// exponent parts are tolerated and truncated toward zero.
func ReadInt64OrNull(data []byte, i int) (int64, int, error) {
	if uint(i) >= uint(len(data)) {
		return 0, i, ErrTruncated
	}
	if data[i] == 'n' {
		end, err := ExpectNull(data, i)
		return 0, end, err
	}
	neg := false
	if data[i] == '-' {
		neg = true
		i++
		if uint(i) >= uint(len(data)) {
			return 0, i, ErrBadNumber
		}
	}
	if data[i]-'0' > 9 { // unsigned: a byte below '0' wraps high, so one compare
		return 0, i, ErrBadNumber
	}
	var n int64
	// The digit run, a word at a time (digitRun): this reader is called per
	// member from a decoder loop, so the address chain its count creates is
	// broken by the key read that follows; the batch array loops keep their
	// byte-loop tail for the opposite reason (see decodeIntSlice).
	v, end := digitRun(data, i)
	n = int64(v)
	i = end
	if i < len(data) && (data[i] == '.' || data[i] == 'e' || data[i] == 'E') {
		for uint(i) < uint(len(data)) {
			c := data[i]
			if (c >= '0' && c <= '9') || c == '.' || c == 'e' || c == 'E' || c == '+' || c == '-' {
				i++
				continue
			}
			break
		}
	}
	if neg {
		n = -n
	}
	return n, i, nil
}

// ReadUint64OrNull reads a JSON unsigned integer (or null) at data[i].
func ReadUint64OrNull(data []byte, i int) (uint64, int, error) {
	if uint(i) >= uint(len(data)) {
		return 0, i, ErrTruncated
	}
	if data[i] == 'n' {
		end, err := ExpectNull(data, i)
		return 0, end, err
	}
	if data[i]-'0' > 9 { // unsigned: a byte below '0' wraps high, so one compare
		return 0, i, ErrBadNumber
	}
	var n uint64
	// The digit run, a word at a time (digitRun): this reader is called per
	// member from a decoder loop, so the address chain its count creates is
	// broken by the key read that follows; the batch array loops keep their
	// byte-loop tail for the opposite reason (see decodeIntSlice).
	v, end := digitRun(data, i)
	n = uint64(v)
	i = end
	if i < len(data) && (data[i] == '.' || data[i] == 'e' || data[i] == 'E') {
		for uint(i) < uint(len(data)) {
			c := data[i]
			if (c >= '0' && c <= '9') || c == '.' || c == 'e' || c == 'E' || c == '+' || c == '-' {
				i++
				continue
			}
			break
		}
	}
	return n, i, nil
}

// scanNumberToken measures the JSON number token at data[i] and reports whether
// this library's own number decoder would accept it, returning the index just
// past the token. It mirrors ReadFloat64OrNull byte for byte — scanFloat, then
// strconv for the tokens scanFloat declines to resolve — and discards only the
// value.
//
// Written out rather than shared with ReadFloat64OrNull on purpose: that reader
// is the hot float path (see CLAUDE.md's float tiers) and is left byte-identical,
// the same reason any.go inlines this fallback at its own call site instead of
// routing through it. The mirror is not left to inspection — the accept sets are
// compared over a generated corpus by TestReadNumberAcceptSetMatchesFloat64.
//
// It is deliberately not skipNumber, whose span is a strict superset of the
// grammar: skipNumber consumes any run of [0-9.eE+-], so it happily measured
// "1.2.3", "--1" and "1e" as tokens.
func scanNumberToken(data []byte, i int) (int, error) {
	_, end, fast, ok := scanFloat(data, i)
	if !ok {
		return end, ErrBadNumber
	}
	if fast {
		return end, nil
	}
	// scanFloat extracted a well-formed token but declined to resolve it (>19
	// significant digits, ambiguous rounding, or an exponent outside the table);
	// strconv decides, and rejects an overflowing magnitude such as 1e309.
	if _, perr := strconv.ParseFloat(unsafeStr(data[i:end]), 64); perr != nil {
		return end, ErrBadNumber
	}
	return end, nil
}

// ReadNumberOrNull reads a JSON number (or null) at data[i] and returns its raw
// literal as a string — the bytes a json.Number holds — copied out verbatim, with
// no value produced; a JSON null yields the empty string.
//
// The literal is validated before it is captured. Without that check the token
// bounds came from a scanner that only measures a run of number bytes, so a
// malformed literal ("1.2.3", "-", "--1", "1e") was stored intact and failed much
// later and far from the decode, at the first .Float64()/.Int64() — and, worse,
// this package's own Valid rejected documents the decoder had accepted.
//
// The acceptance rule is agreement with ReadFloat64OrNull, not with
// encoding/json, because Valid checks numbers by running them through that reader
// (see pkg/json/valid.go): matching it is what makes "Valid accepts exactly what
// these decoders accept" true for json.Number fields as well. Two consequences
// are deliberate and pinned by TestReadNumberAcceptSetMatchesFloat64: "01" and
// ".5" stay accepted though encoding/json rejects them (a pre-existing, separate
// divergence — narrowing it here would only move the disagreement), and a literal
// whose magnitude overflows float64 ("1e309") is rejected though its digits are
// well-formed.
func ReadNumberOrNull(data []byte, i int) (string, int, error) {
	if uint(i) >= uint(len(data)) {
		return "", i, ErrTruncated
	}
	if data[i] == 'n' {
		end, err := ExpectNull(data, i)
		return "", end, err
	}
	end, err := scanNumberToken(data, i)
	if err != nil {
		return "", end, err
	}
	return string(data[i:end]), end, nil
}

// ReadNumberNoCopyOrNull is ReadNumberOrNull but returns a string that aliases
// data instead of copying it, so the caller must keep data unchanged while the
// result is in use.
func ReadNumberNoCopyOrNull(data []byte, i int) (string, int, error) {
	if uint(i) >= uint(len(data)) {
		return "", i, ErrTruncated
	}
	if data[i] == 'n' {
		end, err := ExpectNull(data, i)
		return "", end, err
	}
	end, err := scanNumberToken(data, i)
	if err != nil {
		return "", end, err
	}
	return unsafeStr(data[i:end]), end, nil
}

// ReadFloat64OrNull reads a JSON number (or null) at data[i] as a float64.
func ReadFloat64OrNull(data []byte, i int) (float64, int, error) {
	if uint(i) >= uint(len(data)) {
		return 0, i, ErrTruncated
	}
	if data[i] == 'n' {
		end, err := ExpectNull(data, i)
		return 0, end, err
	}
	f, end, fast, ok := scanFloat(data, i)
	if !ok {
		return 0, end, ErrBadNumber
	}
	if fast {
		return f, end, nil
	}
	// scanFloat already tried Eisel–Lemire on the extracted mantissa/exp10; reaching
	// here means it declined (e.g. >19 significant digits, ambiguous rounding,
	// subnormal/overflow, or exponent outside the table), so defer to strconv.
	// unsafeStr avoids copying the token; ParseFloat does not retain it.
	f, perr := strconv.ParseFloat(unsafeStr(data[i:end]), 64)
	if perr != nil {
		return 0, end, ErrBadNumber
	}
	return f, end, nil
}

// ReadBoolOrNull reads a JSON boolean (or null) at data[i]. The literals are
// matched with constant-string compares (a word load + compare each, see
// ExpectNull) rather than byte at a time; a partial literal returns i, as
// before.
func ReadBoolOrNull(data []byte, i int) (bool, int, error) {
	if uint(i) >= uint(len(data)) {
		return false, i, ErrTruncated
	}
	switch data[i] {
	case 'n':
		end, err := ExpectNull(data, i)
		return false, end, err
	case 't':
		if i+4 <= len(data) && string(data[i:i+4]) == "true" {
			return true, i + 4, nil
		}
		return false, i, ErrInvalidJSON
	case 'f':
		if i+5 <= len(data) && string(data[i:i+5]) == "false" {
			return false, i + 5, nil
		}
		return false, i, ErrInvalidJSON
	default:
		return false, i, ErrInvalidJSON
	}
}

// ReadTimeOrNull reads an RFC 3339 JSON string (or null) at data[i] as a
// time.Time. Its authority for the grammar is time.Parse(time.RFC3339, ...):
// the fast path in date.go is only ever allowed to be more conservative, and
// everything it declines is handed to time.Parse itself, which
// TestReadTimeMatchesStdlibAcceptance locks over a generated date corpus. That
// is also what encoding/json's time.Time reduces to today — its extra RFC 3339
// strictness is compiled out pending go.dev/issue/54580 — so on an escape-free
// timestamp the two accept the same set and produce the same instant.
//
// The parity stops at the JSON string layer, and only in the lenient direction:
// this reads the string *value* (escapes decoded) and parses that, where
// time.Time.UnmarshalJSON through Go 1.26 parsed the raw bytes between the
// quotes without unescaping them (go.dev/issue/47353). So a timestamp written
// with any \uXXXX escape — legal JSON denoting a legal instant, such as
// "2021-01-01T00:00:00\u005A" — decodes here and was rejected by that stdlib;
// Go 1.27's json/v2-backed encoding/json unescapes first and agrees.
// TestReadTimeAcceptsEscapedTimestamps pins both halves on either toolchain.
// ReadTimeLaxOrNull inherits the same leniency by construction.
//
// The intermediate string aliases data — safe because time.Parse retains it in
// neither its result nor its error (the stdlib copies into ParseError; locked by
// TestReadTimeErrorRetainsNoAlias) — so this allocates only the time.Time.
func ReadTimeOrNull(data []byte, i int) (time.Time, int, error) {
	if uint(i) >= uint(len(data)) {
		return time.Time{}, i, ErrTruncated
	}
	if data[i] == 'n' {
		end, err := ExpectNull(data, i)
		return time.Time{}, end, err
	}
	s, end, err := ReadStringNoCopyOrNull(data, i)
	if err != nil {
		return time.Time{}, end, err
	}
	if t, ok := parseRFC3339(s, false); ok {
		return t, end, nil
	}
	t, perr := time.Parse(time.RFC3339, s)
	if perr != nil {
		return time.Time{}, end, perr
	}
	return t, end, nil
}

// ReadTimeLaxOrNull reads a time.Time at data[i], accepting more shapes than
// ReadTimeOrNull: an RFC 3339 string with either a 'T' or a space separator and
// optional fractional seconds, or a Unix timestamp (in seconds, milliseconds, or
// microseconds) given as a JSON number or a numeric string. The result is
// normalized to UTC. Anything it cannot interpret returns ErrBadTime, which the
// "lax" decode path turns into a skipped value and an unset field.
func ReadTimeLaxOrNull(data []byte, i int) (time.Time, int, error) {
	if uint(i) >= uint(len(data)) {
		return time.Time{}, i, ErrTruncated
	}
	switch data[i] {
	case 'n':
		end, err := ExpectNull(data, i)
		return time.Time{}, end, err
	case '"':
		s, end, err := ReadStringNoCopyOrNull(data, i)
		if err != nil {
			return time.Time{}, end, err
		}
		if t, ok := parseJSONTS(s); ok {
			return t, end, nil
		}
		if t, ok := parseNumTS(s); ok { // numeric timestamp encoded as a string
			return t, end, nil
		}
		return time.Time{}, end, ErrBadTime
	default:
		end, err := skipNumber(data, i)
		if err != nil {
			return time.Time{}, end, err
		}
		if t, ok := parseNumTS(unsafeStr(data[i:end])); ok {
			return t, end, nil
		}
		return time.Time{}, end, ErrBadTime
	}
}
