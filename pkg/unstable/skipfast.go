package unstable

import "math/bits"

// SIMD in-string-mask container skip (the sonic-rs skip_container technique),
// the production fast path SkipValue dispatches container skips to.
// skipObject/skipArray land on each structural byte with indexStructural and
// call SkipString per string; this instead streams 64-byte
// blocks, computes an in-string mask (prefix-XOR of the real quote positions,
// simdjson's find_escaped + carry handling), masks the bracket bitmaps to bytes
// *outside* strings, and balances the container. Strings are absorbed into the
// bulk scan — no per-string call. Used only on the skip path (Get/GetPaths
// sibling-skip, unknown-field skip), where there is no typed stage-2, so the
// economics that sank the full two-stage feed do not apply.
//
// 64 bytes per block (vs an earlier 32): maskBlock processes a whole 64-byte
// block per call, halving the call/return-marshaling/VZEROUPPER overhead and the
// number of find_escaped/prefix_xor invocations vs two 32-byte blocks. maskBlock
// (AVX2 asm on amd64, NEON on arm64, scalar elsewhere) returns four uint64s —
// quote, backslash, and the container's *own* open and close brackets (it is told
// which via isArray, and never computes the other type's brackets). All the bit
// math below is plain Go, unit/differential testable independent of the asm.
//
// What the two skip paths guarantee, precisely: on a *well-formed* (bracket-
// balanced) value they return the identical end index, which is what the 50k-doc
// differential fuzz and the block-boundary corpus lock. They are NOT
// interchangeable off that set, and an earlier version of this comment wrongly
// claimed they "never diverge (including on malformed input)". Three divergences
// are real, and because SkipValue picks between the paths by CPU feature, all
// make SkipValue's answer on bad input host-dependent:
//
//   - Unbalanced brackets of the *other* type. This loop ignores them entirely;
//     skipObject/skipArray *descend* into them (skipObject calls skipArray on a
//     nested '['), so the scalar path keeps looking for the inner container's
//     close. Measured: `{"a":[}` + slack gives (7, nil) here and (87,
//     ErrTruncated) there; `{"a":[}]}` gives 7 vs 9 — accept/reject and the end
//     index both differ.
//   - Nesting past MaxDepth. This loop is genuinely iterative (depth is an int,
//     no recursion, no per-level state), so it has no stack to overflow and
//     accepts any depth; the recursive scalar path must bound itself and returns
//     ErrMaxDepth past MaxDepth (see skip.go). MaxDepth is documented as a
//     library-wide limit, so this path is the lenient outlier, not the scalar
//     one.
//   - A stray backslash *outside* a string. findEscaped64 is a pure bit
//     operation on the block's backslash bitmap: it does not know the backslash
//     is not inside a string, so it masks the following byte out of the quote
//     bitmap, and a `\"` sitting between tokens silently changes where the
//     in-string regions begin. The scalar path has no such state — indexStructural
//     simply never stops on a backslash. This one diverges in *both* directions
//     and, uniquely, is sensitive to where the 64-byte grid falls, so the same
//     logical document flips verdict on padding alone. Measured:
//     `{"a":1,\"b":2,"pad":"…"}` gives (93, ErrTruncated) here and (93, nil)
//     there — the fast path rejecting what the scalar accepts, the opposite of
//     the first bullet; `{\"}` + slack gives (4, nil) vs (84, ErrTruncated); and
//     `{` + N spaces + `\}` is accepted at N=62, rejected at N=63, accepted again
//     at N=64, because at N=63 the backslash is the block's last byte, so
//     prevEscaped carries into the tail and swallows the container's own closing
//     brace. TestSkipPathsDivergeOnMalformed pins all three classes.
//
// None costs correctness today — every caller treats such input as an error
// or, for the presize counters, as a hint that missed — but do not build on the
// two paths agreeing outside well-formed values. Converging the depth case means
// clamping depth here *and* in the whole-loop assembly scans (skipBlocks*), which
// carry depth in a register and return it but never test it; that is a hot-loop
// change on three assembly implementations, so it is deliberately not done here.

// prefixXor64 returns the inclusive prefix XOR of x: bit i of the result is the
// parity (XOR) of bits 0..i of x. Run over a block's real-quote bitmap it turns
// quote positions into "inside a string" regions (set after an odd quote count).
// This is the carryless-multiply-by-all-ones simdjson uses, done in six
// shift+XOR steps so it needs no PCLMULQDQ.
func prefixXor64(x uint64) uint64 {
	x ^= x << 1
	x ^= x << 2
	x ^= x << 4
	x ^= x << 8
	x ^= x << 16
	x ^= x << 32
	return x
}

// findEscaped64 returns the bitmap of bytes that are escaped by a preceding
// backslash, given the backslash bitmap for the block. *prevEscaped carries the
// boundary state between blocks (0 or 1: whether the block's first byte is
// escaped by a backslash that ended the previous block). This is simdjson's
// branchless find_escaped: it isolates odd-length backslash runs with an
// add-carry so that, e.g., \\\" escapes the quote but \\" does not.
func findEscaped64(backslash uint64, prevEscaped *uint64) uint64 {
	const evenBits uint64 = 0x5555555555555555
	backslash &^= *prevEscaped
	followsEscape := backslash<<1 | *prevEscaped
	oddSequenceStarts := backslash &^ evenBits &^ followsEscape
	seq, carry := bits.Add64(oddSequenceStarts, backslash, 0)
	*prevEscaped = carry
	invertMask := seq << 1
	return (evenBits ^ invertMask) & followsEscape
}

// skipContainerFast skips the JSON container that opens at data[i] (data[i] is
// open, the matching close is '}' for '{' and ']' for '['), returning the index
// just past the close. It counts only that bracket pair outside strings; a stray
// bracket of the other type is ignored rather than descended into, which is one
// of the two documented divergences from skipObject/skipArray above. A
// truncated container returns ErrTruncated. Depth is an int, not recursion, so
// arbitrarily deep input is safe here (and, unlike the scalar path, accepted)
// — except on a machine without the SIMD maskBlock, where the first line hands
// the value straight to skipObject/skipArray and their bound applies.
func skipContainerFast(data []byte, i int, open byte) (int, error) {
	if !fastSkipAvail {
		// No SIMD maskBlock here, and the scalar one is slower than the
		// indexStructural balance — which is what fastSkipAvail says. The gate
		// used to sit in SkipValue's arms; it moved here so that SkipObject can
		// be a single call and inline into a walker that has already seen the
		// brace.
		if open == '{' {
			return skipObject(data, i)
		}
		return skipArray(data, i)
	}
	pos := i + 1
	if useSkipBlocks && pos+64 <= len(data) {
		// amd64: the whole block loop runs in assembly (skipBlocks) — splats
		// loaded once, depth/escape/in-string state carried in registers, the
		// prefix XOR done with one carryless multiply — eliminating the
		// per-block maskBlock call/marshaling that dominated the Go loop. It
		// consumes every full 64-byte block: either it finds the close (end >=
		// 0) or it hands the carried state to the shared scalar tail below.
		//
		// Everything that follows the call is out of line — skipContainerBlocks
		// holds the Go block loop and the byte tail — and that is not tidiness.
		// With them inline the compiler must keep data, i, open, close, isArray
		// and depth alive ACROSS the call, which on the path that returns right
		// after it (a container whose close is in the first block: every record
		// in an array of records, every unknown-field skip a generated decoder
		// does) cost a 136-byte frame and eight spill stores before the call
		// was even made. The outlined form asks only for what a failed scan
		// needs, and the arguments it needs are this function's own.
		end, d, pe, pis := skipBlocks(data, pos, 1, open == '[')
		if end >= 0 {
			return end, nil
		}
		return skipContainerBlocks(data, pos+((len(data)-pos)&^63), open, d, pe, pis)
	}
	return skipContainerBlocks(data, pos, open, 1, 0, 0)
}

// skipContainerBlocks is skipContainerFast's continuation: the Go maskBlock
// block loop (which runs only where the whole-loop assembly is unavailable —
// see useSkipBlocks) and the byte tail both paths end in. pos is the first byte
// it has not accounted for, depth the balance there, and prevEscaped /
// prevInString the two bits that carry across a block boundary.
func skipContainerBlocks(data []byte, pos int, open byte, depth int, prevEscaped, prevInString uint64) (int, error) {
	close := byte('}')
	if open == '[' {
		close = ']'
	}
	isArray := open == '['
	for pos+64 <= len(data) {
		quote, bslash, op, cl := maskBlock(data[pos:], isArray)

		// The escaped -> inStr -> prevInString computation is a loop-carried
		// dependency chain (each block's in-string mask depends on the previous
		// block's), and the block loop is latency-bound on it — the profile puts
		// these few ALU lines well above the vector kernel. Two predictable
		// branches keep the common cases off that chain:
		//
		// Escapes are rare: a block with no backslash and none pending from the
		// previous block has nothing escaped, skipping findEscaped64's add-carry
		// steps. And a block with no (unescaped) quote cannot change the
		// in-string state — the mask is just the carried prevInString — skipping
		// the six-step prefix-XOR chain (number/bracket-dense blocks).
		var escaped uint64
		if bslash|prevEscaped != 0 {
			escaped = findEscaped64(bslash, &prevEscaped)
		}
		inStr := prevInString
		if q := quote &^ escaped; q != 0 {
			inStr = prefixXor64(q) ^ prevInString
			prevInString = uint64(int64(inStr) >> 63) // carry: inside-string at byte 63
		}

		op &^= inStr
		cl &^= inStr
		// The per-bit walk exists only to spot the depth-0 crossing. A block
		// with no close bracket cannot cross, and neither can one with fewer
		// closes than the current depth (opens only raise it, so the running
		// minimum is at least depth - popcount(cl) >= 1): both update depth with
		// popcounts in bulk. Only a block that might actually close the
		// container walks its bracket bits in order.
		if cl == 0 {
			depth += bits.OnesCount64(op)
		} else if nc := bits.OnesCount64(cl); depth > nc {
			depth += bits.OnesCount64(op) - nc
		} else {
			brackets := op | cl
			for brackets != 0 {
				j := bits.TrailingZeros64(brackets)
				if op&(uint64(1)<<uint(j)) != 0 {
					depth++
				} else {
					depth--
					if depth == 0 {
						return pos + j + 1, nil
					}
				}
				brackets &= brackets - 1
			}
		}
		pos += 64
	}

	// Tail: fewer than 64 bytes remain. Walk byte by byte, carrying the
	// inside-string / pending-escape state out of the block loop.
	inStr := prevInString != 0
	esc := prevEscaped != 0
	for ; uint(pos) < uint(len(data)); pos++ {
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

// SkipObject is SkipValue's '{' arm, exported for the same reason SkipNumber
// is: a walker that has already looked at data[i] and seen a '{' can spell the
// arm at the call site and skip SkipValue's frame and its comparison tree. It
// is kept to a single call so that it inlines — a wrapper that did not would
// only trade one frame for another — which is why the fastSkipAvail gate moved
// inside skipContainerFast.
func SkipObject(data []byte, i int) (int, error) {
	return skipContainerFast(data, i, '{')
}
