package unstable

import "bytes"

// CountArrayElements returns the number of top-level elements in the JSON array
// beginning at data[i] (data[i] must be '['), so a destination slice can be
// allocated once instead of grown by repeated append. It returns 0 for an empty
// array or when the count cannot be determined cheaply (a truncated or malformed
// array); the caller then simply falls back to append-driven growth, so an
// imperfect count is only ever a missed optimization, never a correctness
// problem.
//
// Each element is skipped whole with SkipValue rather than walked byte by byte:
// SkipValue uses the SIMD indexStructural scanner for nested arrays/objects and
// indexCloseOrEscape for strings, so a structurally dense element — a nested
// coordinate array of many numbers, say — is jumped over in vectorized strides
// instead of one byte at a time. This is what makes presizing slices of arrays,
// objects, or strings cheap.
// countSampleCap bounds the per-element walk: a huge array (apache_builds' 875
// job objects) need not be skipped in full just to size a slice. After this many
// elements the total is extrapolated from the bytes the sample spans, turning an
// O(array) skip into O(countSampleCap).
const countSampleCap = 64

func CountArrayElements(data []byte, i int) int {
	if i >= len(data) || data[i] != '[' {
		return 0
	}
	open := i
	i = SkipWS(data, i+1)
	if i >= len(data) || data[i] == ']' {
		return 0
	}
	n := 1
	for {
		end, err := SkipValue(data, i)
		if err != nil {
			return 0
		}
		i = SkipWS(data, end)
		if i >= len(data) {
			return 0
		}
		switch data[i] {
		case ']':
			return n
		case ',':
			n++
			i = SkipWS(data, i+1)
			if i >= len(data) {
				return 0
			}
			if n == countSampleCap {
				// Sample large enough: stop skipping element-by-element and
				// estimate the total from the array's byte span. The closing ']'
				// is located with a plain IndexByte — the first ']' is at or
				// before the true close (a ']' inside a string only moves it
				// earlier), so for bracket-free elements the estimate matches and
				// it is a presize hint regardless: a wrong count mis-sizes the
				// slice, never misdecodes.
				rel := bytes.IndexByte(data[i:], ']')
				if rel < 0 {
					return n
				}
				span := i + rel - open // '[' .. approximate ']'
				sampled := i - open    // bytes covering n elements
				est := n * span / sampled
				if est > span { // every element spans at least one byte
					est = span
				}
				return est
			}
		default:
			return 0 // malformed; let the caller grow on demand
		}
	}
}

// CountArrayScalars counts the elements of a JSON array of scalar values
// (numbers, booleans, or null) beginning at data[i] (data[i] must be '['). Such
// elements never contain a quote, comma, or bracket, so the closing ']' is the
// first ']' in the input and the element count is one more than the number of
// commas before it — both found with vectorized byte scans. It is therefore much
// cheaper than CountArrayElements but valid only when the element type is known
// (by the generator) to be a scalar. It returns 0 for an empty array.
func CountArrayScalars(data []byte, i int) int {
	rb := bytes.IndexByte(data[i+1:], ']')
	if rb < 0 {
		return 0
	}
	seg := data[i+1 : i+1+rb]
	for _, c := range seg {
		if c != ' ' && c != '\t' && c != '\n' && c != '\r' {
			return bytes.Count(seg, commaByte) + 1
		}
	}
	return 0 // empty (whitespace-only) array
}

var commaByte = []byte{','}
var openBraceByte = []byte{'{'}

// CountArrayObjects counts the elements of a JSON array of "bracket-free" objects
// beginning at data[i] (data[i] must be '['). A bracket-free object has only
// number/bool fields — its JSON ({"a":1,"b":2}) holds no string, '[' or ']', and
// no nested '{' — so the array's closing ']' is the first ']' in the input and the
// element count is exactly the number of '{' before it, both found with vectorized
// byte scans. Like CountArrayScalars it is far cheaper than CountArrayElements (no
// per-element SkipValue) and valid only for the element shape the generator
// vouches for; it sizes a slice of flat numeric records — citm_catalog's price
// entries — without re-scanning every struct. As a presize hint a miscount on
// unexpected input only over- or under-allocates, never misdecodes. Returns 0 for
// an empty array.
func CountArrayObjects(data []byte, i int) int {
	if i >= len(data) {
		return 0
	}
	rb := bytes.IndexByte(data[i+1:], ']')
	if rb < 0 {
		return 0
	}
	return bytes.Count(data[i+1:i+1+rb], openBraceByte)
}
