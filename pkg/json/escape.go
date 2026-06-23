package json

import (
	"encoding/binary"
	"math/bits"
	"strings"

	"github.com/JohanLindvall/lightning/pkg/unstable"
)

const hexAlphabet = "0123456789abcdef"

// SWAR (SIMD-within-a-register) helpers operating on eight packed bytes at once,
// used by EscapeStringInto's first-run probe. swarHasLess reports (high bit per
// lane) whether any byte lane of v is < n (1 <= n <= 128); swarHasByte reports
// whether any lane equals b.
const (
	swarLo = 0x0101010101010101 // 1 in every byte lane
	swarHi = 0x8080808080808080 // high bit of every byte lane
)

func swarHasLess(v, n uint64) uint64 { return (v - swarLo*n) & ^v & swarHi }
func swarHasByte(v uint64, b byte) uint64 {
	x := v ^ (swarLo * uint64(b))
	return (x - swarLo) & ^x & swarHi
}

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
		// Find the next byte to escape, choosing the scanner by how much input is
		// left in the run — decided once per run, so there is no per-word bookkeeping
		// to tax the common short/escape-dense cases.
		//
		// minVectorRun: below this the vector scanner cannot amortize its per-call
		// setup, so short strings and the short runs between escapes (JSON-in-a-
		// string, Windows paths, prose) walk words with SWAR — exact offset via
		// TrailingZeros, no vector call. At or above it the run is long enough that
		// the SIMD pass wins (log lines, URLs, mostly-clean text): probe the first
		// word with SWAR (so an immediate escape still avoids setup) and hand the
		// clean bulk to the scanner.
		const minVectorRun = 48
		if n-i >= minVectorRun {
			v := binary.LittleEndian.Uint64(s[i:])
			if m := swarHasLess(v, 0x20) | swarHasByte(v, '"') | swarHasByte(v, '\\'); m != 0 {
				i += bits.TrailingZeros64(m) >> 3
			} else {
				i += 8 + unstable.IndexEscape(s[i+8:])
			}
		} else {
			for i+8 <= n {
				v := binary.LittleEndian.Uint64(s[i:])
				if m := swarHasLess(v, 0x20) | swarHasByte(v, '"') | swarHasByte(v, '\\'); m != 0 {
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
