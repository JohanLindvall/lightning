package unstable

import (
	"encoding/binary"
	"math/bits"
)

// SkipValue advances past any JSON value starting at data[i].
//
// Objects and dense arrays (whose first element is itself an object, array, or
// string) are skipped with the SIMD in-string-mask balance scan
// (skipContainerFast) when it is available (AVX2 on amd64): it absorbs string
// keys/values into one bulk pass instead of a SkipString call per string, which
// is a large win on the containers Get/GetPaths and unknown-field skipping walk
// over. A scalar-element array ('[1,2,...]') keeps the indexStructural skip,
// where a single vectorized scan already reaches the closing bracket and the
// mask path would only add per-block work. The array probe is a heuristic; a
// wrong guess only costs speed, never correctness — both paths are bracket
// balancers that return the same end index for every well-formed value.
//
// Off the well-formed set the two paths are not interchangeable, so which one
// runs — and therefore what SkipValue answers on malformed input — depends on
// the host CPU. skipfast.go's header enumerates the divergence classes (three
// as of this writing) and TestSkipPathsDivergeOnMalformed pins them; that list
// is the authority and is deliberately not restated here, since a second copy
// is what let this comment go stale while claiming to be exhaustive. Every
// caller treats such input as an error or a presize miss.
func SkipValue(data []byte, i int) (int, error) {
	if i >= len(data) {
		return i, ErrTruncated
	}
	switch data[i] {
	case '"':
		return SkipString(data, i)
	case '{':
		if fastSkipAvail {
			return skipContainerFast(data, i, '{')
		}
		return skipObject(data, i)
	case '[':
		if fastSkipAvail {
			j := SkipWS(data, i+1)
			if j < len(data) {
				switch data[j] {
				case '{', '[', '"':
					return skipContainerFast(data, i, '[')
				}
			}
		}
		return skipArray(data, i)
	case 't':
		if i+4 <= len(data) && string(data[i:i+4]) == "true" {
			return i + 4, nil
		}
		return i, ErrInvalidJSON
	case 'f':
		if i+5 <= len(data) && string(data[i:i+5]) == "false" {
			return i + 5, nil
		}
		return i, ErrInvalidJSON
	case 'n':
		return ExpectNull(data, i)
	default:
		return skipNumber(data, i)
	}
}

// SkipString advances past the JSON string starting at data[i] (data[i] must be
// '"') and returns the index just past its closing quote. Escapes are honored so
// an escaped quote (\") does not end the string; the scan itself does not
// validate or decode the escapes. It returns ErrTruncated if the closing quote
// is missing before the end of data.
func SkipString(data []byte, i int) (int, error) {
	// data[i] == '"'
	i++
	for {
		e := indexCloseOrEscapeAt(data, i)
		if e == len(data) {
			return len(data), ErrTruncated
		}
		if data[e] == '"' {
			return e + 1, nil
		}
		// Skip the escape sequence. For \uXXXX we only need to skip the
		// backslash and the next char; subsequent bytes are processed on the
		// next iteration.
		i = e + 2
		if i > len(data) {
			return len(data), ErrTruncated
		}
	}
}

// skipNumber measures a number token; it does not validate one. Any non-empty
// run of [0-9.eE+-] is a token, so SkipValue([]byte("+"), 0) is (1, nil) and
// "-", "e", ".", "1.2.3" and "+-e." are number spans too — only an empty run
// (a byte that starts none of the other value kinds, e.g. 'x') is ErrBadNumber.
// That is deliberate and consistent with the rest of the skip path being a
// bracket balancer rather than a parser: SkipValue's job is to find where a
// value ends so the caller can step over it, and the readers that produce a
// value from those same bytes are the ones that decide it is a number —
// ReadFloat64OrNull/ReadNumberOrNull reject every literal above, and Valid,
// which runs the document through those readers, rejects the document. Only a
// caller that mistakes SkipValue for validation is surprised here; pkg/json's
// checked wrappers exist because the fast walkers make exactly this trade.
func skipNumber(data []byte, i int) (int, error) {
	start := i
	if i < len(data) && data[i] == '-' {
		i++
	}
	for i < len(data) {
		c := data[i]
		if (c >= '0' && c <= '9') || c == '.' || c == 'e' || c == 'E' || c == '+' || c == '-' {
			i++
			continue
		}
		break
	}
	if i == start {
		return i, ErrBadNumber
	}
	return i, nil
}

// skipObject and skipArray are the scalar container skips SkipValue falls back
// to when the SIMD in-string-mask scan is unavailable or declines (see
// skipContainerFast). They are *mutually recursive* — one Go frame per nesting
// level — so they carry the same MaxDepth bound the other recursive walkers do
// (decodeValue/decodeAnyObject/decodeAnyArray, stripper.handle). Without it the
// recursion is unbounded and a hostile document exhausts the goroutine stack:
// measured, a 20 MB array of the [0,[[[…]]]] shape (the leading scalar makes
// SkipValue's fast-path probe decline, so every nested '[' recursed) died with
// "goroutine stack exceeds 1000000000-byte limit / fatal error: stack overflow",
// which recover *cannot* catch — the process goes down instead of an error
// coming back. The bound belongs here rather than at the call sites because this
// path is reachable from all of them: Get/Lookup/GetMany/GetPaths/Set/SetMany/
// SetPaths/ObjectEach/ArrayEach, CountArrayElements, and every generated
// decoder's unknown-field skip.
//
// depth counts enclosing containers with the outermost at 1, exactly as
// decodeAny* does, so MaxDepth levels are accepted and level MaxDepth+1 returns
// ErrMaxDepth — the same input set encoding/json accepts depth-wise. Cost is one
// compare per '{' or '[', which CLAUDE.md measured as flat (p=0.161) for the
// analogous DecodeAny bound. The two-argument entry points are kept so the
// differential oracle and benchmarks spell them unchanged; they are trivially
// inlinable, so the extra frame is compiled away.
func skipObject(data []byte, i int) (int, error) { return skipObjectDepth(data, i, 1) }

func skipArray(data []byte, i int) (int, error) { return skipArrayDepth(data, i, 1) }

func skipObjectDepth(data []byte, i, depth int) (int, error) {
	if depth > MaxDepth {
		return i, ErrMaxDepth
	}
	// data[i] == '{'
	i++
	for i < len(data) {
		// Jump to the next structural byte, skipping inert content (keys' inner
		// chars, numbers, bools, whitespace) in one vectorized pass.
		i += indexStructural(data[i:])
		if i >= len(data) {
			break
		}
		switch data[i] {
		case '}':
			return i + 1, nil
		case '{':
			end, err := skipObjectDepth(data, i, depth+1)
			if err != nil {
				return end, err
			}
			i = end
		case '[':
			end, err := skipArrayDepth(data, i, depth+1)
			if err != nil {
				return end, err
			}
			i = end
		case '"':
			end, err := SkipString(data, i)
			if err != nil {
				return end, err
			}
			i = end
		default: // stray ']' (only on malformed input); step over it
			i++
		}
	}
	return i, ErrTruncated
}

func skipArrayDepth(data []byte, i, depth int) (int, error) {
	if depth > MaxDepth {
		return i, ErrMaxDepth
	}
	// data[i] == '['
	i++
	for i < len(data) {
		i += indexStructural(data[i:])
		if i >= len(data) {
			break
		}
		switch data[i] {
		case ']':
			return i + 1, nil
		case '[':
			end, err := skipArrayDepth(data, i, depth+1)
			if err != nil {
				return end, err
			}
			i = end
		case '{':
			end, err := skipObjectDepth(data, i, depth+1)
			if err != nil {
				return end, err
			}
			i = end
		case '"':
			end, err := SkipString(data, i)
			if err != nil {
				return end, err
			}
			i = end
		default: // stray '}' (only on malformed input); step over it
			i++
		}
	}
	return i, ErrTruncated
}

// SkipWS advances past JSON whitespace at data[i]. The four JSON whitespace
// bytes (space, tab, newline, carriage return) are all <= ' ' (0x20), so a single
// compare classifies a byte with no memory load — measurably faster than a lookup
// table on every shape from a compact exit to a deep indent run. This is the
// hottest classification in the scanner, running before and after every value.
//
// The compare also treats the other control bytes (0x00..0x1f) as whitespace.
// Those are never valid JSON between tokens, so on well-formed input the result
// is identical to matching the four bytes exactly; on malformed input SkipWS
// skips such a byte rather than stopping on it, leaving the value parser to
// reject the next real token. SkipWS is not called inside strings, so control
// bytes within string contents are unaffected.
func SkipWS(data []byte, i int) int {
	for i < len(data) && data[i] <= ' ' {
		i++
	}
	return i
}

// SkipWSRun advances past a whitespace run from data[i] (the caller has already
// established that data[i-1] and data[i-2] were whitespace, i.e. this is a run of
// at least two). It is the out-of-line continuation the generated decoders call
// only for genuine indentation runs in pretty-printed input — short skips (zero
// or one whitespace byte, the compact and single-space-after-token cases) are
// handled inline at the call site so they never pay a call. Eight bytes are
// classified per word (see the derivation at nws below); a full-whitespace word
// is skipped whole and the first word with a structural byte locates it with one
// trailing-zeros count.
func SkipWSRun(data []byte, i int) int {
	const (
		lo  = ^uint64(0) / 255 // 0x0101...01
		hi  = lo << 7          // 0x8080...80
		sp  = lo * ' '         // 0x2020...20: a word of eight literal spaces
		low = lo * 0x5f        // 0x5f...5f: 0x80 - (' ' + 1), the lane bias below
	)
	for i+8 <= len(data) {
		// data[i:i+8], not data[i:]: the explicit end is what lets the compiler
		// prove the load in bounds from the loop condition it already tested.
		// Written data[i:] the word costs SIX more instructions per iteration —
		// a second bounds compare, the len and cap subtractions with their
		// negative clamp, and Uint64's own "at least 8 bytes" test — because the
		// open-ended reslice is a value the prove pass then has to re-derive
		// facts about. Measured 9-12% on citm-shaped indentation runs; this loop
		// is issue-bound (citm decodes at IPC 3.6 with a 0.18% branch-miss rate),
		// so instructions removed are cycles removed. The same spelling is
		// already used by every other SWAR load here (read.go, batch.go,
		// numeric.go) — this one was the exception.
		w := binary.LittleEndian.Uint64(data[i : i+8])
		// Eight literal spaces is the overwhelmingly common word inside an
		// indentation run, and equality against the splat answers it with one
		// compare instead of the five-op classify below. It is a *sufficient*
		// condition — every byte is exactly 0x20 — so this only skips work; the
		// exact SWAR still decides every other word, including the one that ends
		// the run, so nothing about which input is accepted changes.
		if w != sp {
			// A bit set per lane that is NOT whitespace — the complement of the
			// whitespace mask, so the run-terminating exit needs no XOR to invert
			// it, and the non-zero test in front of TrailingZeros64 *proves* the
			// operand non-zero, which stops the compiler materialising 64 and
			// CMOVE-ing it as the all-zero guard.
			//
			// Derivation. w&^hi leaves each lane holding byte&0x7f, at most 0x7f,
			// so adding 0x5f per lane tops out at 0xde and CANNOT carry into the
			// neighbouring lane; the sum reaches 0x80 exactly when byte&0x7f >=
			// 0x21. That marks every byte in 0x21..0x7f and 0xa1..0xff, and the
			// | w folds in the lanes whose own top bit is set (0x80..0xff), so the
			// union is exactly "> 0x20" — the complement of SkipWS's <= ' ' rule,
			// for every byte value.
			//
			// This replaces the borrow-guard subtract (g - w&^hi) &^ w with g =
			// 0xa0...a0, which computes the same mask in one more operation and a
			// deeper chain — but the real cost was the constant: 0xa0 is not an
			// AArch64 bitmask immediate, so g took a MOVD + three MOVKs to
			// materialise, INSIDE the loop and again on the back edge, while
			// 0x8080../0x7f7f.. fold into their AND as immediates. Same trap on
			// 0x5f, but it is added once where g was rebuilt twice. Interleaved
			// micro over citm-shaped indent runs (8/16/24/28 spaces, the shape the
			// corpus actually has): -10.8%, -18.6%, -24.1%, -25.1%. Dropping the
			// w != sp shortcut on top of it was measured and is WORSE (+16% at
			// indent 24) — the shortcut still pays even against the cheaper
			// classify. Equivalence is not left to the argument above:
			// TestSkipWSRunMatchesOracle drives every byte value at every lane
			// offset and FuzzSkipWSRunMatchesOracle backs it.
			nws := ((w&^hi + low) | w) & hi
			if nws != 0 {
				return i + bits.TrailingZeros64(nws)/8
			}
		}
		i += 8
	}
	for i < len(data) && data[i] <= ' ' {
		i++
	}
	return i
}

// SkipWSCompact is the compact-aware inter-token whitespace skip shared by the
// dynamic DecodeValue path and the pkg/json toolkit's compact variants (GetCompact,
// SetMany, ObjectEachCompact, …). In compact mode the input is asserted to carry no
// whitespace between tokens (the form compact JSON serializers emit), so it returns
// i unchanged; otherwise it is SkipWS. This mirrors the generator's
// //lightning:compact decoders, which elide exactly these inter-token skips.
func SkipWSCompact(data []byte, i int, compact bool) int {
	for !compact && i < len(data) && data[i] <= ' ' {
		i++
	}
	return i
}
