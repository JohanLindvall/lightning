package unstable

import (
	"encoding/base64"
	"strconv"
	"unsafe"
)

// Batched scalar-array readers: each decodes a whole JSON array of numbers in
// one call. The generated per-element loop paid a non-inlinable reader call
// per number — two frames for floats (ReadFloat64OrNull → scanFloat) — plus
// its own append/branch machinery; on number-array-heavy documents (canada,
// numbers, mesh) that dispatch is a measurable slice of the decode. Keeping
// the loop here lets it call the private scanFloat directly (one frame per
// float) and inline the SWAR digit fold (no call at all per int). Semantics
// match the generated loop exactly: a null root yields a nil slice (fixed
// arrays are left untouched), a null element a zero value, error identities
// and positions match the Read* readers, a nil *out is presized with
// CountArrayScalars just as slicePresize emitted, and a trailing comma is
// rejected by the same first-iteration flag (a ']' at the loop top on a
// non-first iteration is only reachable after a comma).
//
// The generator routes a slice or fixed-size-array field here whenever the
// element is a bare float64/int/uint kind; other element types (float32, bool,
// strings, structs, ...) keep the generated loop.
//
// The slice readers work on a local copy of the slice header, not through
// *out: the compiler cannot prove *out doesn't alias data, so appending via
// the pointer reloads ptr/len/cap and stores the new len every element. The
// local keeps the header in registers across the (call-free, for integers)
// loop. The price is that every return — errors included — must write the
// local back, to preserve the partial-progress-on-error behavior the parity
// tests lock.
//
// Each slice reader has an ...Arena twin for the //lightning:arena decoders,
// which carves the fresh presized backing from a shared chunk instead of
// allocating per slice (see Arena). Both entry points are thin inlinable
// wrappers over one shared body, so the per-element loops stay identical and
// the arena test runs once per slice on the presize path.

// intKind and uintKind are the integer element kinds the generic readers
// stencil over. rune and byte are covered as aliases of int32 and uint8.
type intKind interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type uintKind interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// DecodeFloat64Slice decodes a JSON array of numbers at data[i] into *out,
// returning the index just past the closing ']'. A JSON null sets *out to nil;
// a null element decodes as 0. When *out is nil the slice is presized from a
// vectorized comma count.
func DecodeFloat64Slice(out *[]float64, data []byte, i int) (int, error) {
	return decodeFloat64Slice(out, data, i, nil)
}

// DecodeFloat64SliceArena is DecodeFloat64Slice with the fresh backing carved
// from a instead of allocated per slice (see Arena); the //lightning:arena
// decoders call it. Semantics are identical, including reuse: a non-nil *out
// keeps its existing backing and the arena is untouched.
func DecodeFloat64SliceArena(out *[]float64, data []byte, i int, a *Arena) (int, error) {
	return decodeFloat64Slice(out, data, i, a)
}

// decodeFloat64Slice is the shared body: a nil arena allocates the presized
// backing with make, exactly as before the arena existed. The arena test sits
// on the once-per-slice presize path, not in the per-element loop, so the
// non-arena entry point costs one predictable branch over the old form.
func decodeFloat64Slice(out *[]float64, data []byte, i int, a *Arena) (int, error) {
	if uint(i) >= uint(len(data)) {
		return i, ErrTruncated
	}
	if data[i] == 'n' {
		end, err := ExpectNull(data, i)
		if err != nil {
			return end, err
		}
		*out = nil
		return end, nil
	}
	if data[i] != '[' {
		return i, ErrExpectArray
	}
	// Reset the length, keeping the backing array: decoding an array into a slice
	// replaces its contents rather than appending to them, as encoding/json
	// documents ("Unmarshal resets the slice length to zero and then appends each
	// element"). Without this, decoding twice into the same value accumulated —
	// [1,2] read into a reused target became [1,2,1,2] — and a caller reusing a
	// target to avoid allocation (the reason to reuse one) grew it without bound.
	// A nil slice stays nil through [:0], so the presize below still fires for the
	// fresh case and is correctly skipped for a reused one, which already has cap.
	s := (*out)[:0]
	if s == nil {
		if n := CountArrayScalars(data, i); n > 0 {
			if a != nil {
				s = arenaCarve[float64](a, n)
			} else {
				s = make([]float64, 0, n)
			}
		}
	}
	i++
	for first := true; ; first = false {
		// Inter-token whitespace, in the shape the generator inlines: 0-1 bytes
		// resolve in one or two compares; only a real indentation run reaches the
		// SWAR SkipWSRun (which the compiler inlines here too — no call).
		if uint(i) < uint(len(data)) && data[i] <= ' ' {
			i++
			if uint(i) < uint(len(data)) && data[i] <= ' ' {
				i = SkipWSRun(data, i+1)
			}
		}
		if uint(i) >= uint(len(data)) {
			*out = s
			return i, ErrTruncated
		}
		if data[i] == ']' {
			*out = s
			if first {
				return i + 1, nil
			}
			return i, ErrInvalidJSON // trailing comma
		}
		var f float64
		if data[i] == 'n' {
			end, err := ExpectNull(data, i)
			if err != nil {
				*out = s
				return end, err
			}
			i = end
		} else {
			v, end, fast, ok := scanFloat(data, i)
			if !ok {
				*out = s
				return end, ErrBadNumber
			}
			if !fast {
				// scanFloat declined (see ReadFloat64OrNull); defer the exact
				// token to strconv. unsafeStr avoids copying the token.
				v2, perr := strconv.ParseFloat(unsafeStr(data[i:end]), 64)
				if perr != nil {
					*out = s
					return end, ErrBadNumber
				}
				v = v2
			}
			f = v
			i = end
		}
		s = append(s, f)
		// Inter-token whitespace, in the shape the generator inlines: 0-1 bytes
		// resolve in one or two compares; only a real indentation run reaches the
		// SWAR SkipWSRun (which the compiler inlines here too — no call).
		if uint(i) < uint(len(data)) && data[i] <= ' ' {
			i++
			if uint(i) < uint(len(data)) && data[i] <= ' ' {
				i = SkipWSRun(data, i+1)
			}
		}
		if uint(i) >= uint(len(data)) {
			*out = s
			return i, ErrTruncated
		}
		if data[i] == ']' {
			*out = s
			return i + 1, nil
		}
		if data[i] != ',' {
			*out = s
			return i, ErrInvalidJSON
		}
		i++
	}
}

// DecodeIntSlice decodes a JSON array of integers at data[i] into *out. The
// element parse mirrors ReadInt64OrNull byte for byte — SWAR 4-digit folds, a
// scalar tail, tolerated (truncated) fraction/exponent, overflow wrap — inlined
// into the loop so an element costs no call. The parsed int64 is converted to T
// exactly as the generated per-element code converted (wrapping, not
// saturating). A JSON null sets *out to nil; a null element decodes as 0.
func DecodeIntSlice[T intKind](out *[]T, data []byte, i int) (int, error) {
	return decodeIntSlice(out, data, i, nil)
}

// DecodeIntSliceArena is DecodeIntSlice with the fresh backing carved from a
// (see DecodeFloat64SliceArena); the //lightning:arena decoders call it.
func DecodeIntSliceArena[T intKind](out *[]T, data []byte, i int, a *Arena) (int, error) {
	return decodeIntSlice(out, data, i, a)
}

// decodeIntSlice is the shared body; a nil arena presizes with make, exactly
// as before the arena existed (see decodeFloat64Slice).
func decodeIntSlice[T intKind](out *[]T, data []byte, i int, a *Arena) (int, error) {
	if uint(i) >= uint(len(data)) {
		return i, ErrTruncated
	}
	if data[i] == 'n' {
		end, err := ExpectNull(data, i)
		if err != nil {
			return end, err
		}
		*out = nil
		return end, nil
	}
	if data[i] != '[' {
		return i, ErrExpectArray
	}
	// Reset the length, keeping the backing array: decoding an array into a slice
	// replaces its contents rather than appending to them, as encoding/json
	// documents ("Unmarshal resets the slice length to zero and then appends each
	// element"). Without this, decoding twice into the same value accumulated —
	// [1,2] read into a reused target became [1,2,1,2] — and a caller reusing a
	// target to avoid allocation (the reason to reuse one) grew it without bound.
	// A nil slice stays nil through [:0], so the presize below still fires for the
	// fresh case and is correctly skipped for a reused one, which already has cap.
	s := (*out)[:0]
	if s == nil {
		if n := CountArrayScalars(data, i); n > 0 {
			if a != nil {
				s = arenaCarve[T](a, n)
			} else {
				s = make([]T, 0, n)
			}
		}
	}
	// The SIMD run kernel's state: run is cleared by an unproductive call, so
	// an array it cannot help costs one call; tmp receives its values for
	// the element kinds it cannot write directly (narrower than 8 bytes).
	run := true
	var tmp [32]int64
	i++
	for first := true; ; first = false {
		// Inter-token whitespace, in the shape the generator inlines: 0-1 bytes
		// resolve in one or two compares; only a real indentation run reaches the
		// SWAR SkipWSRun (which the compiler inlines here too — no call).
		if uint(i) < uint(len(data)) && data[i] <= ' ' {
			i++
			if uint(i) < uint(len(data)) && data[i] <= ' ' {
				i = SkipWSRun(data, i+1)
			}
		}
		if uint(i) >= uint(len(data)) {
			*out = s
			return i, ErrTruncated
		}
		if data[i] == ']' {
			*out = s
			if first {
				return i + 1, nil
			}
			return i, ErrInvalidJSON // trailing comma
		}
		// An element that starts with a digit is where the SIMD run kernel
		// takes over (useIntRun): it writes every following short element
		// straight into the slice's spare capacity (or a scratch, for the
		// narrower kinds) and hands back at the first element it does not
		// handle — a sign, null, a fraction, 9+ digits — which the scalar
		// code below then reads exactly as it always has. One unproductive
		// call turns it off for the rest of the array, so an array of long
		// ids or decimals pays for it once. The spare capacity is also the
		// array's remaining length when the target was presized from its
		// comma count, which is what intRunMinSlots gates on: a kernel with
		// a per-call cost of its own (arm64's classifies a 64-byte block)
		// is not worth calling for the two-element arrays of a schema like
		// twitter_status's indices.
		if useIntRun && run && data[i]-'0' <= 9 && cap(s)-len(s) >= intRunMinSlots {
			var n, p, closed int
			if unsafe.Sizeof(T(0)) == 8 {
				dst := unsafe.Slice((*int64)(unsafe.Pointer(unsafe.SliceData(s))), cap(s))[len(s):]
				n, p, closed = parseIntRun(data, i, dst)
				s = s[:len(s)+n]
			} else {
				n, p, closed = parseIntRun(data, i, tmp[:])
				for _, v := range tmp[:n] {
					s = append(s, T(v))
				}
			}
			if n > 0 {
				i = p
				if closed != 0 {
					*out = s
					return i + 1, nil
				}
				continue
			}
			run = false
		}
		var n int64
		if data[i] == 'n' {
			end, err := ExpectNull(data, i)
			if err != nil {
				*out = s
				return end, err
			}
			i = end
		} else {
			neg := false
			if data[i] == '-' {
				neg = true
				i++
				if uint(i) >= uint(len(data)) {
					*out = s
					return i, ErrBadNumber
				}
			}
			if data[i]-'0' > 9 { // unsigned: a byte below '0' wraps high
				*out = s
				return i, ErrBadNumber
			}
			// Four digits per SWAR step, then a byte tail — deliberately NOT the
			// word-at-a-time digitRun the single-value readers use. The corpus's
			// array integers are 1-4 digits (marine_ik 130k of them, mesh's index
			// arrays), and back to back: a byte loop advances under a predicted
			// branch, so the next element's load issues ahead, while a word count
			// puts that address on a load→mask→count data chain. Measured on a
			// Neoverse N2: digitRun here was mesh +1.1% with fewer instructions,
			// and two guarded hybrids were worse still (CLAUDE.md, 2026-09-02).
			if uint(i)+4 <= uint(len(data)) {
				if v, ok := tryParse4Digits(load32(data, i)); ok { // in bounds: the test above
					n = int64(v)
					i += 4
					// 5+ digits (ids, timestamps): keep folding four at a time. A
					// separate loop, entered only then, so the common 1-4 digit
					// element runs straight-line: the slice header the loop forced
					// the register allocator to spill and reload around it (three
					// stores and six loads per element) now moves only on this path.
					for uint(i)+4 <= uint(len(data)) {
						v, ok := tryParse4Digits(load32(data, i))
						if !ok {
							break
						}
						n = n*10000 + int64(v)
						i += 4
					}
				}
			}
			for uint(i) < uint(len(data)) {
				d := data[i] - '0'
				if d > 9 {
					break
				}
				n = n*10 + int64(d)
				i++
			}
			if uint(i) < uint(len(data)) && (data[i] == '.' || data[i] == 'e' || data[i] == 'E') {
				for uint(i) < uint(len(data)) {
					c := data[i]
					if (c >= '0' && c <= '9') || c == '.' || c == 'e' || c == 'E' || c == '+' || c == '-' {
						i++
						continue
					}
					break
				}
			}
			if neg {
				n = -n
			}
		}
		s = append(s, T(n))
		// Inter-token whitespace, in the shape the generator inlines: 0-1 bytes
		// resolve in one or two compares; only a real indentation run reaches the
		// SWAR SkipWSRun (which the compiler inlines here too — no call).
		if uint(i) < uint(len(data)) && data[i] <= ' ' {
			i++
			if uint(i) < uint(len(data)) && data[i] <= ' ' {
				i = SkipWSRun(data, i+1)
			}
		}
		if uint(i) >= uint(len(data)) {
			*out = s
			return i, ErrTruncated
		}
		if data[i] == ']' {
			*out = s
			return i + 1, nil
		}
		if data[i] != ',' {
			*out = s
			return i, ErrInvalidJSON
		}
		i++
	}
}

// DecodeUintSlice is DecodeIntSlice for the unsigned kinds; the element parse
// mirrors ReadUint64OrNull.
func DecodeUintSlice[T uintKind](out *[]T, data []byte, i int) (int, error) {
	return decodeUintSlice(out, data, i, nil)
}

// DecodeUintSliceArena is DecodeUintSlice with the fresh backing carved from a
// (see DecodeFloat64SliceArena); the //lightning:arena decoders call it.
func DecodeUintSliceArena[T uintKind](out *[]T, data []byte, i int, a *Arena) (int, error) {
	return decodeUintSlice(out, data, i, a)
}

// decodeUintSlice is the shared body; a nil arena presizes with make, exactly
// as before the arena existed (see decodeFloat64Slice).
func decodeUintSlice[T uintKind](out *[]T, data []byte, i int, a *Arena) (int, error) {
	if uint(i) >= uint(len(data)) {
		return i, ErrTruncated
	}
	if data[i] == 'n' {
		end, err := ExpectNull(data, i)
		if err != nil {
			return end, err
		}
		*out = nil
		return end, nil
	}
	if data[i] != '[' {
		return i, ErrExpectArray
	}
	// Reset the length, keeping the backing array: decoding an array into a slice
	// replaces its contents rather than appending to them, as encoding/json
	// documents ("Unmarshal resets the slice length to zero and then appends each
	// element"). Without this, decoding twice into the same value accumulated —
	// [1,2] read into a reused target became [1,2,1,2] — and a caller reusing a
	// target to avoid allocation (the reason to reuse one) grew it without bound.
	// A nil slice stays nil through [:0], so the presize below still fires for the
	// fresh case and is correctly skipped for a reused one, which already has cap.
	s := (*out)[:0]
	if s == nil {
		if n := CountArrayScalars(data, i); n > 0 {
			if a != nil {
				s = arenaCarve[T](a, n)
			} else {
				s = make([]T, 0, n)
			}
		}
	}
	// The SIMD run kernel's state: run is cleared by an unproductive call, so
	// an array it cannot help costs one call; tmp receives its values for
	// the element kinds it cannot write directly (narrower than 8 bytes).
	run := true
	var tmp [32]int64
	i++
	for first := true; ; first = false {
		// Inter-token whitespace, in the shape the generator inlines: 0-1 bytes
		// resolve in one or two compares; only a real indentation run reaches the
		// SWAR SkipWSRun (which the compiler inlines here too — no call).
		if uint(i) < uint(len(data)) && data[i] <= ' ' {
			i++
			if uint(i) < uint(len(data)) && data[i] <= ' ' {
				i = SkipWSRun(data, i+1)
			}
		}
		if uint(i) >= uint(len(data)) {
			*out = s
			return i, ErrTruncated
		}
		if data[i] == ']' {
			*out = s
			if first {
				return i + 1, nil
			}
			return i, ErrInvalidJSON // trailing comma
		}
		// The SIMD run kernel, as in decodeIntSlice.
		if useIntRun && run && data[i]-'0' <= 9 && cap(s)-len(s) >= intRunMinSlots {
			var n, p, closed int
			if unsafe.Sizeof(T(0)) == 8 {
				dst := unsafe.Slice((*int64)(unsafe.Pointer(unsafe.SliceData(s))), cap(s))[len(s):]
				n, p, closed = parseIntRun(data, i, dst)
				s = s[:len(s)+n]
			} else {
				n, p, closed = parseIntRun(data, i, tmp[:])
				for _, v := range tmp[:n] {
					s = append(s, T(v))
				}
			}
			if n > 0 {
				i = p
				if closed != 0 {
					*out = s
					return i + 1, nil
				}
				continue
			}
			run = false
		}
		var n uint64
		if data[i] == 'n' {
			end, err := ExpectNull(data, i)
			if err != nil {
				*out = s
				return end, err
			}
			i = end
		} else {
			if data[i]-'0' > 9 {
				*out = s
				return i, ErrBadNumber
			}
			if uint(i)+4 <= uint(len(data)) {
				if v, ok := tryParse4Digits(load32(data, i)); ok { // in bounds: the test above
					n = uint64(v)
					i += 4
					// 5+ digits (ids, timestamps): keep folding four at a time. A
					// separate loop, entered only then, so the common 1-4 digit
					// element runs straight-line: the slice header the loop forced
					// the register allocator to spill and reload around it (three
					// stores and six loads per element) now moves only on this path.
					for uint(i)+4 <= uint(len(data)) {
						v, ok := tryParse4Digits(load32(data, i))
						if !ok {
							break
						}
						n = n*10000 + uint64(v)
						i += 4
					}
				}
			}
			for uint(i) < uint(len(data)) {
				d := data[i] - '0'
				if d > 9 {
					break
				}
				n = n*10 + uint64(d)
				i++
			}
			if uint(i) < uint(len(data)) && (data[i] == '.' || data[i] == 'e' || data[i] == 'E') {
				for uint(i) < uint(len(data)) {
					c := data[i]
					if (c >= '0' && c <= '9') || c == '.' || c == 'e' || c == 'E' || c == '+' || c == '-' {
						i++
						continue
					}
					break
				}
			}
		}
		s = append(s, T(n))
		// Inter-token whitespace, in the shape the generator inlines: 0-1 bytes
		// resolve in one or two compares; only a real indentation run reaches the
		// SWAR SkipWSRun (which the compiler inlines here too — no call).
		if uint(i) < uint(len(data)) && data[i] <= ' ' {
			i++
			if uint(i) < uint(len(data)) && data[i] <= ' ' {
				i = SkipWSRun(data, i+1)
			}
		}
		if uint(i) >= uint(len(data)) {
			*out = s
			return i, ErrTruncated
		}
		if data[i] == ']' {
			*out = s
			return i + 1, nil
		}
		if data[i] != ',' {
			*out = s
			return i, ErrInvalidJSON
		}
		i++
	}
}

// DecodeFloat64Array decodes a JSON array of numbers at data[i] into the
// fixed-size array whose backing the caller passes as out (out = arr[:]). It
// mirrors the generated fixed-array decoder: the array is zeroed, up to
// len(out) elements are decoded, extras are skipped, a short JSON array leaves
// the tail zero, and a JSON null leaves the array untouched. This is the
// per-point call for coordinate rings ([][2]float64, [][3]float64), where the
// generated form paid an extra call frame per coordinate.
func DecodeFloat64Array(out []float64, data []byte, i int) (int, error) {
	if uint(i) >= uint(len(data)) {
		return i, ErrTruncated
	}
	if data[i] == 'n' {
		return ExpectNull(data, i)
	}
	if data[i] != '[' {
		return i, ErrExpectArray
	}
	clear(out)
	i++
	idx := 0
	for first := true; ; first = false {
		// Inter-token whitespace, in the shape the generator inlines: 0-1 bytes
		// resolve in one or two compares; only a real indentation run reaches the
		// SWAR SkipWSRun (which the compiler inlines here too — no call).
		if uint(i) < uint(len(data)) && data[i] <= ' ' {
			i++
			if uint(i) < uint(len(data)) && data[i] <= ' ' {
				i = SkipWSRun(data, i+1)
			}
		}
		if uint(i) >= uint(len(data)) {
			return i, ErrTruncated
		}
		if data[i] == ']' {
			if first {
				return i + 1, nil
			}
			return i, ErrInvalidJSON // trailing comma
		}
		if idx < len(out) {
			if data[i] == 'n' {
				end, err := ExpectNull(data, i)
				if err != nil {
					return end, err
				}
				i = end
			} else {
				v, end, fast, ok := scanFloat(data, i)
				if !ok {
					return end, ErrBadNumber
				}
				if !fast {
					v2, perr := strconv.ParseFloat(unsafeStr(data[i:end]), 64)
					if perr != nil {
						return end, ErrBadNumber
					}
					v = v2
				}
				out[idx] = v
				i = end
			}
		} else {
			end, err := SkipValue(data, i)
			if err != nil {
				return end, err
			}
			i = end
		}
		idx++
		// Inter-token whitespace, in the shape the generator inlines: 0-1 bytes
		// resolve in one or two compares; only a real indentation run reaches the
		// SWAR SkipWSRun (which the compiler inlines here too — no call).
		if uint(i) < uint(len(data)) && data[i] <= ' ' {
			i++
			if uint(i) < uint(len(data)) && data[i] <= ' ' {
				i = SkipWSRun(data, i+1)
			}
		}
		if uint(i) >= uint(len(data)) {
			return i, ErrTruncated
		}
		if data[i] == ']' {
			return i + 1, nil
		}
		if data[i] != ',' {
			return i, ErrInvalidJSON
		}
		i++
	}
}

// DecodeIntArray is DecodeFloat64Array for the integer kinds; the element
// parse mirrors ReadInt64OrNull (inlined, as in DecodeIntSlice).
func DecodeIntArray[T intKind](out []T, data []byte, i int) (int, error) {
	if uint(i) >= uint(len(data)) {
		return i, ErrTruncated
	}
	if data[i] == 'n' {
		return ExpectNull(data, i)
	}
	if data[i] != '[' {
		return i, ErrExpectArray
	}
	clear(out)
	i++
	idx := 0
	for first := true; ; first = false {
		// Inter-token whitespace, in the shape the generator inlines: 0-1 bytes
		// resolve in one or two compares; only a real indentation run reaches the
		// SWAR SkipWSRun (which the compiler inlines here too — no call).
		if uint(i) < uint(len(data)) && data[i] <= ' ' {
			i++
			if uint(i) < uint(len(data)) && data[i] <= ' ' {
				i = SkipWSRun(data, i+1)
			}
		}
		if uint(i) >= uint(len(data)) {
			return i, ErrTruncated
		}
		if data[i] == ']' {
			if first {
				return i + 1, nil
			}
			return i, ErrInvalidJSON // trailing comma
		}
		if idx < len(out) {
			if data[i] == 'n' {
				end, err := ExpectNull(data, i)
				if err != nil {
					return end, err
				}
				i = end
			} else {
				neg := false
				if data[i] == '-' {
					neg = true
					i++
					if uint(i) >= uint(len(data)) {
						return i, ErrBadNumber
					}
				}
				if data[i]-'0' > 9 {
					return i, ErrBadNumber
				}
				var n int64
				if uint(i)+4 <= uint(len(data)) {
					if v, ok := tryParse4Digits(load32(data, i)); ok { // in bounds: the test above
						n = int64(v)
						i += 4
						// 5+ digits (ids, timestamps): keep folding four at a time. A
						// separate loop, entered only then, so the common 1-4 digit
						// element runs straight-line: the slice header the loop forced
						// the register allocator to spill and reload around it (three
						// stores and six loads per element) now moves only on this path.
						for uint(i)+4 <= uint(len(data)) {
							v, ok := tryParse4Digits(load32(data, i))
							if !ok {
								break
							}
							n = n*10000 + int64(v)
							i += 4
						}
					}
				}
				for uint(i) < uint(len(data)) {
					d := data[i] - '0'
					if d > 9 {
						break
					}
					n = n*10 + int64(d)
					i++
				}
				if uint(i) < uint(len(data)) && (data[i] == '.' || data[i] == 'e' || data[i] == 'E') {
					for uint(i) < uint(len(data)) {
						c := data[i]
						if (c >= '0' && c <= '9') || c == '.' || c == 'e' || c == 'E' || c == '+' || c == '-' {
							i++
							continue
						}
						break
					}
				}
				if neg {
					n = -n
				}
				out[idx] = T(n)
			}
		} else {
			end, err := SkipValue(data, i)
			if err != nil {
				return end, err
			}
			i = end
		}
		idx++
		// Inter-token whitespace, in the shape the generator inlines: 0-1 bytes
		// resolve in one or two compares; only a real indentation run reaches the
		// SWAR SkipWSRun (which the compiler inlines here too — no call).
		if uint(i) < uint(len(data)) && data[i] <= ' ' {
			i++
			if uint(i) < uint(len(data)) && data[i] <= ' ' {
				i = SkipWSRun(data, i+1)
			}
		}
		if uint(i) >= uint(len(data)) {
			return i, ErrTruncated
		}
		if data[i] == ']' {
			return i + 1, nil
		}
		if data[i] != ',' {
			return i, ErrInvalidJSON
		}
		i++
	}
}

// DecodeUintArray is DecodeIntArray for the unsigned kinds; the element parse
// mirrors ReadUint64OrNull (inlined, as in DecodeUintSlice).
func DecodeUintArray[T uintKind](out []T, data []byte, i int) (int, error) {
	if uint(i) >= uint(len(data)) {
		return i, ErrTruncated
	}
	if data[i] == 'n' {
		return ExpectNull(data, i)
	}
	if data[i] != '[' {
		return i, ErrExpectArray
	}
	clear(out)
	i++
	idx := 0
	for first := true; ; first = false {
		// Inter-token whitespace, in the shape the generator inlines: 0-1 bytes
		// resolve in one or two compares; only a real indentation run reaches the
		// SWAR SkipWSRun (which the compiler inlines here too — no call).
		if uint(i) < uint(len(data)) && data[i] <= ' ' {
			i++
			if uint(i) < uint(len(data)) && data[i] <= ' ' {
				i = SkipWSRun(data, i+1)
			}
		}
		if uint(i) >= uint(len(data)) {
			return i, ErrTruncated
		}
		if data[i] == ']' {
			if first {
				return i + 1, nil
			}
			return i, ErrInvalidJSON // trailing comma
		}
		if idx < len(out) {
			if data[i] == 'n' {
				end, err := ExpectNull(data, i)
				if err != nil {
					return end, err
				}
				i = end
			} else {
				if data[i]-'0' > 9 {
					return i, ErrBadNumber
				}
				var n uint64
				if uint(i)+4 <= uint(len(data)) {
					if v, ok := tryParse4Digits(load32(data, i)); ok { // in bounds: the test above
						n = uint64(v)
						i += 4
						// 5+ digits (ids, timestamps): keep folding four at a time. A
						// separate loop, entered only then, so the common 1-4 digit
						// element runs straight-line: the slice header the loop forced
						// the register allocator to spill and reload around it (three
						// stores and six loads per element) now moves only on this path.
						for uint(i)+4 <= uint(len(data)) {
							v, ok := tryParse4Digits(load32(data, i))
							if !ok {
								break
							}
							n = n*10000 + uint64(v)
							i += 4
						}
					}
				}
				for uint(i) < uint(len(data)) {
					d := data[i] - '0'
					if d > 9 {
						break
					}
					n = n*10 + uint64(d)
					i++
				}
				if uint(i) < uint(len(data)) && (data[i] == '.' || data[i] == 'e' || data[i] == 'E') {
					for uint(i) < uint(len(data)) {
						c := data[i]
						if (c >= '0' && c <= '9') || c == '.' || c == 'e' || c == 'E' || c == '+' || c == '-' {
							i++
							continue
						}
						break
					}
				}
				out[idx] = T(n)
			}
		} else {
			end, err := SkipValue(data, i)
			if err != nil {
				return end, err
			}
			i = end
		}
		idx++
		// Inter-token whitespace, in the shape the generator inlines: 0-1 bytes
		// resolve in one or two compares; only a real indentation run reaches the
		// SWAR SkipWSRun (which the compiler inlines here too — no call).
		if uint(i) < uint(len(data)) && data[i] <= ' ' {
			i++
			if uint(i) < uint(len(data)) && data[i] <= ' ' {
				i = SkipWSRun(data, i+1)
			}
		}
		if uint(i) >= uint(len(data)) {
			return i, ErrTruncated
		}
		if data[i] == ']' {
			return i + 1, nil
		}
		if data[i] != ',' {
			return i, ErrInvalidJSON
		}
		i++
	}
}

// DecodeByteSlice decodes the JSON value at data[i] into *out with
// encoding/json's []byte semantics, which accept two shapes: a string value is
// base64 (StdEncoding, decoded with the same base64.Decode call the stdlib
// uses, so accepted inputs and error identities match), and an array is the
// numeric element form (each element parsed like any uint kind, overflow wrap
// included) shared with every other uint slice. null yields a nil slice. Like
// the other slice readers it replaces *out's contents, reusing its backing when
// the decoded bytes fit.
//
// On a decode error *out holds the prefix that did decode, the partial-progress
// semantics the readers here share — and here it is not merely a convention:
// reusing the backing means a failed base64 decode has already overwritten the
// caller's previous bytes, so leaving *out alone would report a stale length
// over rewritten data. encoding/json, which always decodes into a fresh buffer,
// leaves its target untouched instead; a caller that must keep the old value
// across a failed decode has to keep its own copy.
func DecodeByteSlice(out *[]byte, data []byte, i int) (int, error) {
	return decodeByteSlice(out, data, i, nil)
}

// DecodeByteSliceArena is DecodeByteSlice with the numeric-array form's backing
// carved from a (see DecodeFloat64SliceArena); the //lightning:arena decoders
// call it. The base64 form keeps a plain make: its length comes from the string
// body, not a presize count, and Decode fills every byte, so there is no
// zeroing to save.
func DecodeByteSliceArena(out *[]byte, data []byte, i int, a *Arena) (int, error) {
	return decodeByteSlice(out, data, i, a)
}

// decodeByteSlice is the shared body (see decodeFloat64Slice for the pattern).
func decodeByteSlice(out *[]byte, data []byte, i int, a *Arena) (int, error) {
	if uint(i) >= uint(len(data)) {
		return i, ErrTruncated
	}
	if data[i] != '"' {
		// null, '[', or an error — all exactly the shared uint-slice reader.
		return decodeUintSlice(out, data, i, a)
	}
	s, end, err := ReadStringNoCopyOrNull(data, i)
	if err != nil {
		return end, err
	}
	// A nil *out always goes through make so an empty base64 string yields the
	// non-nil empty slice encoding/json produces (null, not "", is what maps to
	// a nil slice).
	b := *out
	if need := base64.StdEncoding.DecodedLen(len(s)); b == nil || cap(b) < need {
		b = make([]byte, need)
	} else {
		b = b[:need]
	}
	// Publish the decode on every return, like every other reader in this file.
	// base64.Decode fills b quantum by quantum and only then reports a bad one, so
	// when b is *out's reused backing the caller's previous bytes are already gone
	// by the time the error surfaces — returning with *out's old length left it
	// describing bytes that had been overwritten, a silent mix of the previous
	// document and this one. b[:n] instead says exactly what was decoded.
	n, derr := base64.StdEncoding.Decode(b, unsafeBytes(s))
	*out = b[:n]
	return end, derr
}
