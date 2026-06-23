package json

import (
	"strings"

	"github.com/JohanLindvall/lightning/pkg/unstable"
)

const hexAlphabet = "0123456789abcdef"

// EscapeString writes the JSON-escaped form of s to out — the string body only,
// without the surrounding quotes. The bytes JSON requires escaped (control bytes
// below 0x20, '"' and '\\') are replaced by their escape sequences: the short
// forms \t, \r, \n, \" and \\ where defined, and \u00XX for the remaining control
// bytes. It is the inverse of UnescapeString. A clean prefix — the common case of
// a string with no escapes — is written straight to out, with neither a scratch
// buffer nor a rescan.
func EscapeString(s []byte, out *strings.Builder) {
	pos := unstable.IndexEscape(s)
	if pos == len(s) {
		// Nothing to escape: write the bytes straight to the builder, avoiding
		// both the scratch buffer and the copy into it.
		out.Write(s)
		return
	}
	// Write the clean prefix directly and escape only the remainder, so the
	// prefix is neither re-scanned nor copied through a scratch buffer.
	out.Write(s[:pos])
	out.Write(EscapeStringInto(s[pos:], make([]byte, 0, len(s)-pos)))
}

// EscapeStringInto appends the JSON-escaped form of s to out and returns the
// extended slice; out may be nil or a buffer reused across calls to avoid
// allocation (escaping can lengthen the input, so out still grows when its
// capacity is exceeded). It escapes the same bytes as EscapeString — control
// bytes below 0x20, '"' and '\\' — and writes the string body only, without the
// surrounding quotes, mirroring the in/out convention of UnescapeStringInto.
func EscapeStringInto(s []byte, out []byte) []byte {
	n := len(s)
	p := 0 // start of the current run of bytes that need no escaping
	i := 0

	for {
		// Skip the next run of bytes that need no escaping (a control byte < 0x20,
		// '"' or '\\') with the SIMD scanner — a long clean run (e.g. a log line) is
		// found in a few vector blocks instead of byte/word steps. The clean run is
		// copied out in bulk below.
		i += unstable.IndexEscape(s[i:])
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
