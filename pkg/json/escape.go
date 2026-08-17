package json

import (
	"encoding/binary"
	"math/bits"
	"strings"
	"unicode/utf8"

	"github.com/JohanLindvall/lightning/pkg/unstable"
)

const hexAlphabet = "0123456789abcdef"

// rfffd is the raw UTF-8 encoding of U+FFFD, the replacement character every
// byte that is not part of a well-formed UTF-8 sequence is substituted with.
// Raw bytes rather than the six-byte \ufffd escape form: both are valid JSON for the
// same string, and the raw form is half the size and what utf8.AppendRune emits.
const rfffd = "\xef\xbf\xbd"

// EscapeString writes the JSON-escaped form of s to out — the string body only,
// without the surrounding quotes. The bytes JSON requires escaped (control bytes
// below 0x20, '"' and '\\') are replaced by their escape sequences: the short
// forms \t, \r, \n, \" and \\ where defined, and \u00XX for the remaining control
// bytes. Bytes that are not part of a well-formed UTF-8 sequence are replaced by
// U+FFFD, exactly as encoding/json replaces them when marshaling — a JSON text
// must be valid UTF-8 (RFC 8259), and decoders substitute U+FFFD rather than
// error, so passing such bytes through verbatim would become silent corruption
// downstream. For valid input EscapeString is the inverse of UnescapeString. A
// clean prefix — the common case of a string with no escapes — is written
// straight to out, with neither a scratch buffer nor a rescan.
func EscapeString(s []byte, out *strings.Builder) {
	pos := unstable.IndexEscapeNonASCII(s)
	if pos == len(s) {
		// Nothing to escape and pure ASCII (so trivially valid UTF-8): write the
		// bytes straight to the builder, avoiding both the scratch buffer and the
		// copy into it.
		out.Write(s)
		return
	}
	if s[pos] >= utf8.RuneSelf {
		// First non-ASCII byte. The prefix is ASCII, so s is valid UTF-8 exactly
		// when the remainder is; decide once here.
		if !utf8.Valid(s[pos:]) {
			out.Write(s[:pos])
			var buf [128]byte
			out.Write(escapeInvalidInto(s[pos:], buf[:0]))
			return
		}
		// Valid UTF-8: multibyte sequences need no escaping, so resume the
		// clean-prefix probe with the plain scanner, and the remainder — now
		// UTF-8-cleared — takes the plain escape walk below.
		pos += unstable.IndexEscape(s[pos:])
		if pos == len(s) {
			out.Write(s)
			return
		}
		out.Write(s[:pos])
		var buf [128]byte
		out.Write(escapeValidInto(s[pos:], buf[:0]))
		return
	}
	// First interesting byte is an escape byte, so the remainder has NOT had its
	// UTF-8 decision yet — it goes through EscapeStringInto's combined walk, which
	// makes that decision at the first non-ASCII byte (if any). Write the clean
	// ASCII prefix directly so it is neither re-scanned nor copied through a
	// scratch buffer. The scratch for the escaped tail is stack-backed: it never
	// escapes (the Builder copies it), so when the escaped tail fits it costs no
	// allocation; a longer tail regrows on the heap as before.
	out.Write(s[:pos])
	var buf [128]byte
	out.Write(EscapeStringInto(s[pos:], buf[:0]))
}

// EscapeStringInto appends the JSON-escaped form of s to out and returns the
// extended slice; out may be nil or a buffer reused across calls to avoid
// allocation (escaping can lengthen the input, so out still grows when its
// capacity is exceeded). It escapes the same bytes as EscapeString — control
// bytes below 0x20, '"' and '\\' — substitutes U+FFFD for every byte that is not
// part of a well-formed UTF-8 sequence (see EscapeString), and writes the string
// body only, without the surrounding quotes.
//
// out must not overlap s. There is no in-place form of escaping and there cannot
// be one: every escape writes more bytes than it consumes, so an out aliasing s
// would overrun the input the scan has not read yet and then rescan the output it
// had just written. That is the one point where this differs from
// UnescapeStringInto, whose out == in[:0] is safe precisely because unescaping
// only ever shrinks.
//
// The UTF-8 handling is free until it matters: the walk below is
// escapeValidInto's with the scan predicate widened by non-ASCII bytes (one
// extra AND+OR per SWAR word, one extra POR per SIMD block), so a pure-ASCII
// string — a log line, a label, a URL — pays nothing else. The first non-ASCII
// byte decides the rest of the input ONCE with utf8.Valid: valid input (unicode
// prose) continues under the plain scanners, and only actually ill-formed input
// takes the substituting walk.
func EscapeStringInto(s []byte, out []byte) []byte {
	n := len(s)
	p := 0 // start of the current run of bytes that need no escaping
	i := 0

	for {
		// Find the next byte to escape or non-ASCII byte; scanner choice per run
		// as in escapeValidInto (see minVectorRun there).
		if n-i >= minVectorRun {
			v := binary.LittleEndian.Uint64(s[i : i+8])
			if m := unstable.SwarNeedsEscapeOrNonASCII(v); m != 0 {
				i += bits.TrailingZeros64(m) >> 3
			} else {
				i += 8 + unstable.IndexEscapeNonASCII(s[i+8:])
			}
		} else {
			for i+8 <= n {
				v := binary.LittleEndian.Uint64(s[i : i+8])
				if m := unstable.SwarNeedsEscapeOrNonASCII(v); m != 0 {
					i += bits.TrailingZeros64(m) >> 3
					break
				}
				i += 8
			}
			// int8(s[i]) >= 0x20 is one signed compare covering both bounds:
			// as int8, non-ASCII bytes (0x80..0xFF) are negative and control
			// bytes are below 0x20 — the same three tests per byte as
			// escapeValidInto's walk.
			for i < n && int8(s[i]) >= 0x20 && s[i] != '"' && s[i] != '\\' {
				i++
			}
		}
		if i >= n {
			break
		}

		out = append(out, s[p:i]...)
		switch c := s[i]; c {
		case '\t':
			out = append(out, '\\', 't')
		case '\r':
			out = append(out, '\\', 'r')
		case '\n':
			out = append(out, '\\', 'n')
		case '\\':
			out = append(out, '\\', '\\')
		case '"':
			out = append(out, '\\', '"')
		default:
			// The widened predicate funnels exactly two kinds here: control
			// bytes and non-ASCII bytes. Testing which INSIDE the default arm
			// keeps the common escape bytes ('\\', '"', \t\r\n) on the same
			// path they took before the UTF-8 handling existed — no per-escape
			// compare.
			if c >= utf8.RuneSelf {
				// First non-ASCII byte: everything before it is ASCII, so s is
				// valid UTF-8 exactly when s[i:] is. Decide once and hand the
				// rest to the matching walk — the plain one (multibyte
				// sequences need no escaping, so they must not break its runs)
				// or the substituting one.
				if utf8.Valid(s[i:]) {
					return escapeValidInto(s[i:], out)
				}
				return escapeInvalidInto(s[i:], out)
			}
			out = append(out, '\\', 'u', '0', '0', hexAlphabet[c>>4], hexAlphabet[c&0xf])
		}

		i++
		p = i
	}

	out = append(out, s[p:]...)

	return out
}

// minVectorRun: below this the vector scanner cannot amortize its per-call
// setup, so short strings and the short runs between escapes (JSON-in-a-string,
// Windows paths, prose) walk words with SWAR — exact offset via TrailingZeros,
// no vector call. At or above it the run is long enough that the SIMD pass wins
// (log lines, URLs, mostly-clean text): probe the first word with SWAR (so an
// immediate escape still avoids setup) and hand the clean bulk to the scanner.
const minVectorRun = 48

// escapeValidInto is the escape walk for input known to be valid UTF-8 (or
// already past its UTF-8 decision): EscapeStringInto's body with the plain
// escape predicate, so multibyte sequences flow through clean runs untouched.
// EscapeStringInto dispatches here the moment the input's first non-ASCII byte
// proves valid, and escapeInvalidInto uses it for the well-formed runs between
// bad bytes.
func escapeValidInto(s []byte, out []byte) []byte {
	n := len(s)
	p := 0 // start of the current run of bytes that need no escaping
	i := 0

	for {
		// Find the next byte to escape, choosing the scanner by how much input is
		// left in the run — decided once per run, so there is no per-word bookkeeping
		// to tax the common short/escape-dense cases (see minVectorRun).
		if n-i >= minVectorRun {
			v := binary.LittleEndian.Uint64(s[i : i+8])
			if m := unstable.SwarNeedsEscape(v); m != 0 {
				i += bits.TrailingZeros64(m) >> 3
			} else {
				i += 8 + unstable.IndexEscape(s[i+8:])
			}
		} else {
			for i+8 <= n {
				v := binary.LittleEndian.Uint64(s[i : i+8])
				if m := unstable.SwarNeedsEscape(v); m != 0 {
					i += bits.TrailingZeros64(m) >> 3
					break
				}
				i += 8
			}
			for i < n && s[i] >= 0x20 && s[i] != '"' && s[i] != '\\' {
				i++
			}
		}
		if i >= n {
			break
		}

		out = append(out, s[p:i]...)
		switch c := s[i]; c {
		case '\t':
			out = append(out, '\\', 't')
		case '\r':
			out = append(out, '\\', 'r')
		case '\n':
			out = append(out, '\\', 'n')
		case '\\':
			out = append(out, '\\', '\\')
		case '"':
			out = append(out, '\\', '"')
		default:
			out = append(out, '\\', 'u', '0', '0', hexAlphabet[c>>4], hexAlphabet[c&0xf])
		}

		i++
		p = i
	}

	out = append(out, s[p:]...)

	return out
}

// escapeInvalidInto is the substituting walk for input that failed utf8.Valid:
// every byte that is not part of a well-formed UTF-8 sequence becomes U+FFFD
// (the r == RuneError && size == 1 convention of utf8.DecodeRune, which is also
// encoding/json's), and the well-formed runs between such bytes — valid by
// construction, since they are made of successfully decoded runes — are escaped
// by escapeValidInto. It runs only for input that is actually ill-formed, so the
// per-rune DecodeRune walk is off every hot path.
func escapeInvalidInto(s []byte, out []byte) []byte {
	run := 0
	for i := 0; i < len(s); {
		// ASCII advances on a byte compare; DecodeRune (a real call) runs only
		// at non-ASCII positions. Escape bytes need no attention here — the
		// escapeValidInto call on each run expands them.
		if s[i] < utf8.RuneSelf {
			i++
			continue
		}
		r, size := utf8.DecodeRune(s[i:])
		if r == utf8.RuneError && size == 1 {
			out = escapeValidInto(s[run:i], out)
			out = append(out, rfffd...)
			i++
			run = i
			continue
		}
		i += size
	}
	return escapeValidInto(s[run:], out)
}
