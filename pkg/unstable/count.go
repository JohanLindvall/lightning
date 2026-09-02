package unstable

import "bytes"

// countSampleCap bounds CountArrayElements' per-element walk: a huge array
// (apache_builds' 875 job objects) need not be skipped in full just to size a
// slice. After this many elements the total is extrapolated from the bytes the
// sample spans, turning an O(array) skip into O(countSampleCap).
const countSampleCap = 64

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
func CountArrayElements(data []byte, i int) int {
	if uint(i) >= uint(len(data)) || data[i] != '[' {
		return 0
	}
	open := i
	i = SkipWS(data, i+1)
	if uint(i) >= uint(len(data)) || data[i] == ']' {
		return 0
	}
	n := 1
	for {
		end, err := SkipValue(data, i)
		if err != nil {
			return 0
		}
		i = SkipWS(data, end)
		if uint(i) >= uint(len(data)) {
			return 0
		}
		switch data[i] {
		case ']':
			return n
		case ',':
			n++
			i = SkipWS(data, i+1)
			if uint(i) >= uint(len(data)) {
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
				// Clamp to what the span can structurally hold. The result is a
				// make() capacity, so an estimate the document's bytes cannot
				// justify is not merely a bad hint — it is an allocation linear
				// in the document with a sizeof(element) multiplier. Between the
				// '[' and the ']' there are span-1 content bytes, and n elements
				// need n-1 separating commas plus at least one byte each, so
				// span-1 >= 2n-1, i.e. n <= span/2. That bound is exact at the
				// densest legal packing ([1,1,…]), so it can never clip an
				// honest count; it only replaces the old est > span guard, which
				// was twice as loose (and, as the same arithmetic shows,
				// unreachable — est is at most about span/2 + span/sampled).
				//
				// Be precise about what this does and does not do. The clamp
				// binds only when the sampled elements are near the densest legal
				// byte, so it does *not* rescue the extrapolation from
				// unrepresentative leading elements: 64 leading `""` or `{}`
				// elements followed by one huge one extrapolate to the same value
				// with or without it (measured: a 1 MB document whose true count
				// is 65 estimates 336 907 either way). That is deliberate and
				// safe, because the property being defended is not "the estimate
				// is accurate" — no cheap counter can promise that — but "the
				// estimate is one a legal document of this byte span could
				// actually have produced". Both the honest and the hostile
				// document of a given size land at the same ceiling, which is the
				// most a hint-based presize can guarantee: measured on a 144-byte
				// element, a 2 MB brace-bomb allocates 98 MB and a 2 MB *honest*
				// densest-legal array of the same type allocates 96 MB, so a
				// crafted input buys 1.02x over simply sending real data. Bounding
				// it below that would mean under-sizing honest documents.
				if lim := span / 2; est > lim {
					est = lim
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
			n := bytes.Count(seg, commaByte) + 1
			// Clamp to the element count seg can structurally hold: n elements
			// need n-1 separating commas plus at least one byte each, so
			// len(seg) >= 2n-1, i.e. n <= (len(seg)+1)/2. Exact at the densest
			// legal packing ([1,2,3]), so no honest count is ever clipped —
			// this only bounds a comma count inflated by commas that are not
			// separators. The element type is vouched scalar by the generator,
			// but []time.Time is in that set, and its quoted values can carry
			// arbitrary bytes on hostile input; the count becomes a make()
			// capacity, so an unbounded comma count is an unbounded allocation.
			// One compare on an already-computed length, never taken on real
			// input.
			if lim := (rb + 1) / 2; n > lim {
				n = lim
			}
			return n
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
	if uint(i) >= uint(len(data)) {
		return 0
	}
	rb := bytes.IndexByte(data[i+1:], ']')
	if rb < 0 {
		return 0
	}
	n := bytes.Count(data[i+1:i+1+rb], openBraceByte)
	// Clamp to the element count the span can structurally hold: an object
	// element is at least '{}' (2 bytes) and n of them need n-1 separating
	// commas, so rb >= 3n-1, i.e. n <= (rb+1)/3. Exact at the densest legal
	// packing ([{},{},…]), so an honest count is never clipped. What it bounds
	// is a brace count inflated by braces inside a *string* value — the shape
	// isFlatScalarStringStruct admits, and the reason this counter is only a
	// hint. Unclamped it is not a harmless hint: the result becomes a make()
	// capacity, so `[{"s":"{{{{…"}]` allocated 65x the input (2 MB in -> 130 MB
	// allocated with a 64-byte element struct), which at scale is an
	// unrecoverable out-of-memory rather than a mis-sized slice.
	if lim := (rb + 1) / 3; n > lim {
		n = lim
	}
	return n
}
