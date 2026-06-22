package unstable

import "math/bits"

// Prototype: SIMD in-string-mask container skip (the sonic-rs skip_container
// technique). skipObject/skipArray today land on each structural byte with
// indexStructural and call SkipString per string; this instead streams 32-byte
// blocks, computes an in-string mask (prefix-XOR of the real quote positions,
// simdjson's find_escaped + carry handling), masks the bracket bitmaps to bytes
// *outside* strings, and balances the one container's bracket pair. Strings are
// absorbed into the bulk scan — no per-string call. Used only on the skip path
// (Get/GetPaths sibling-skip, unknown-field skip), where there is no typed
// stage-2, so the economics that sank the full two-stage feed do not apply.
//
// The character-class masks come from maskBlock (AVX2 asm on amd64, scalar
// elsewhere); all the bit math below is plain Go so it is unit/differential
// testable independent of the asm.

// prefixXor32 returns the inclusive prefix XOR of x: bit i of the result is the
// parity (XOR) of bits 0..i of x. Run over a block's real-quote bitmap it turns
// quote positions into "inside a string" regions (set after an odd quote count).
// This is the carryless-multiply-by-all-ones simdjson uses, done in five
// shift+XOR steps so it needs no PCLMULQDQ.
func prefixXor32(x uint32) uint32 {
	x ^= x << 1
	x ^= x << 2
	x ^= x << 4
	x ^= x << 8
	x ^= x << 16
	return x
}

// findEscaped32 returns the bitmap of bytes that are escaped by a preceding
// backslash, given the backslash bitmap for the block. *prevEscaped carries the
// boundary state between blocks (0 or 1: whether the block's first byte is
// escaped by a backslash that ended the previous block). This is simdjson's
// branchless find_escaped, narrowed to 32 bits: it isolates odd-length backslash
// runs with an add-carry so that, e.g., \\\" escapes the quote but \\" does not.
func findEscaped32(backslash uint32, prevEscaped *uint32) uint32 {
	const evenBits uint32 = 0x55555555
	backslash &^= *prevEscaped
	followsEscape := backslash<<1 | *prevEscaped
	oddSequenceStarts := backslash &^ evenBits &^ followsEscape
	seq, carry := bits.Add32(oddSequenceStarts, backslash, 0)
	*prevEscaped = carry
	invertMask := seq << 1
	return (evenBits ^ invertMask) & followsEscape
}

// skipContainerFast skips the JSON container that opens at data[i] (data[i] is
// open, the matching close is '}' for '{' and ']' for '['), returning the index
// just past the closing bracket. It counts only that bracket pair outside
// strings; a nested array inside an object (or vice versa) is balanced within
// itself and never reaches depth zero for the outer pair, so well-formed input
// is skipped correctly. A truncated/malformed container returns ErrTruncated.
func skipContainerFast(data []byte, i int, open byte) (int, error) {
	close := byte('}')
	if open == '[' {
		close = ']'
	}
	depth := 1
	pos := i + 1
	var prevEscaped, prevInString uint32

	for pos+32 <= len(data) {
		q, bs, lbrace, rbrace, lbrack, rbrack := maskBlock(data[pos:])
		escaped := findEscaped32(bs, &prevEscaped)
		inStr := prefixXor32(q&^escaped) ^ prevInString
		prevInString = uint32(int32(inStr) >> 31) // carry: inside-string at byte 31

		var op, cl uint32
		if open == '{' {
			op, cl = lbrace, rbrace
		} else {
			op, cl = lbrack, rbrack
		}
		op &^= inStr
		cl &^= inStr

		brackets := op | cl
		for brackets != 0 {
			j := bits.TrailingZeros32(brackets)
			if op&(uint32(1)<<uint(j)) != 0 {
				depth++
			} else {
				depth--
				if depth == 0 {
					return pos + j + 1, nil
				}
			}
			brackets &= brackets - 1
		}
		pos += 32
	}

	// Tail: fewer than 32 bytes remain. Walk byte by byte, carrying the
	// inside-string / pending-escape state out of the block loop.
	inStr := prevInString != 0
	esc := prevEscaped != 0
	for ; pos < len(data); pos++ {
		c := data[pos]
		if esc {
			esc = false
			continue
		}
		if inStr {
			switch c {
			case '\\':
				esc = true
			case '"':
				inStr = false
			}
			continue
		}
		switch c {
		case '"':
			inStr = true
		case open:
			depth++
		case close:
			depth--
			if depth == 0 {
				return pos + 1, nil
			}
		}
	}
	return len(data), ErrTruncated
}
