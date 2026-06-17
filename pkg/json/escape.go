package json

import (
	"encoding/binary"
	"strings"
)

func getTable(falseValues ...int) [256]bool {
	table := [256]bool{}

	for i := 0; i < 256; i++ {
		table[i] = true
	}

	for _, v := range falseValues {
		table[v] = false
	}

	return table
}

var escapeTable = getTable(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, '"', '\\')

const hexAlphabet = "0123456789abcdef"

// SWAR helpers operating on eight packed bytes at once.
const (
	swarLo = 0x0101010101010101 // 1 in every byte lane
	swarHi = 0x8080808080808080 // high bit of every byte lane
)

// swarHasLess reports (nonzero) whether any byte lane of v is < n, for
// 1 <= n <= 128. swarHasByte reports whether any lane equals b. Both yield a
// value with the high bit set in each matching lane.
func swarHasLess(v uint64, n uint64) uint64 { return (v - swarLo*n) & ^v & swarHi }
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
	pos := firstEscape(s)
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

// firstEscape returns the index of the first byte EscapeString would escape (a
// control byte < 0x20, '"' or '\\'), or len(s) if none do, scanning eight bytes
// at a time.
func firstEscape(s []byte) int {
	i := 0
	for ; i+8 <= len(s); i += 8 {
		v := binary.LittleEndian.Uint64(s[i:])
		if swarHasLess(v, 0x20)|swarHasByte(v, '"')|swarHasByte(v, '\\') != 0 {
			break
		}
	}
	for ; i < len(s); i++ {
		if !escapeTable[s[i]] {
			return i
		}
	}
	return i
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
		// Skip over bytes that need no escaping, eight at a time where possible:
		// a word is clean unless it holds a control byte (< 0x20), a '"' or a
		// '\\'. Clean runs are copied out in bulk below.
		for i+8 <= n {
			v := binary.LittleEndian.Uint64(s[i:])
			if swarHasLess(v, 0x20)|swarHasByte(v, '"')|swarHasByte(v, '\\') != 0 {
				break
			}
			i += 8
		}
		for i < n && escapeTable[s[i]] {
			i++
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
