//go:build !amd64

package unstable

import "bytes"

// countKernel is the presize scan: the first ']' in data[i:] (rb, its offset
// from i, or -1) and the number of c bytes before it — or, with hint, the
// element count CountArrayScalars wants: the commas plus one clamped to what
// the span can hold, or for a span with no comma 1 or 0 by whether it holds a
// byte above 0x20 (this library's whitespace rule, the one SkipWS applies
// between tokens). amd64 does this in one assembly pass; here it is the two
// vectorized runtime scans the counters always made.
func countKernel(data []byte, i int, c byte, hint bool) (rb, n int) {
	if uint(i) >= uint(len(data)) {
		return -1, 0
	}
	rb = bytes.IndexByte(data[i:], ']')
	if rb < 0 {
		return -1, 0
	}
	sep := commaByte
	if c == '{' {
		sep = openBraceByte
	}
	n = bytes.Count(data[i:i+rb], sep)
	if !hint {
		return rb, n
	}
	if n > 0 {
		n++
		if lim := (rb + 1) / 2; n > lim {
			n = lim
		}
		return rb, n
	}
	for _, b := range data[i : i+rb] {
		if b > ' ' {
			return rb, 1
		}
	}
	return rb, 0
}

// The two bytes the presize counters count, as the one-byte slices
// bytes.Count takes.
var commaByte = []byte{','}
var openBraceByte = []byte{'{'}
