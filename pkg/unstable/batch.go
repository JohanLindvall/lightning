package unstable

import (
	"encoding/binary"
	"strconv"
)

// Batched scalar-array readers: each decodes a whole JSON array of numbers in
// one call. The generated per-element loop pays a non-inlinable call (or two:
// ReadFloat64OrNull → scanFloat) per number plus its own append/branch
// machinery; on number-array-heavy documents (canada, numbers, mesh) that
// dispatch is a measurable slice of the decode. Keeping the loop here lets it
// call the private scanFloat directly — one frame per float instead of two —
// and folds the integer SWAR digit parse into the loop body, so integers cost
// no call at all. Semantics match the generated loop exactly: a null root
// yields a nil slice (fixed arrays are left untouched), a null element a zero
// value, error identities and positions match the Read* readers, and a nil
// *out is presized with CountArrayScalars just as slicePresize emitted.
//
// The generator routes a slice or fixed-size-array field here whenever the
// element is a bare float64/int/uint kind; other element types (float32, bool,
// strings, structs, ...) keep the generated loop.

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
	if i >= len(data) {
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
	if *out == nil {
		if n := CountArrayScalars(data, i); n > 0 {
			*out = make([]float64, 0, n)
		}
	}
	i++
	for {
		// Inter-token whitespace, in the shape the generator inlines: 0-1 bytes
		// resolve in one or two compares; only a real indentation run reaches the
		// SWAR SkipWSRun (which the compiler inlines here too — no call).
		if i < len(data) && data[i] <= ' ' {
			i++
			if i < len(data) && data[i] <= ' ' {
				i = SkipWSRun(data, i+1)
			}
		}
		if i >= len(data) {
			return i, ErrTruncated
		}
		if data[i] == ']' {
			return i + 1, nil
		}
		var f float64
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
				// scanFloat declined (see ReadFloat64OrNull); defer the exact
				// token to strconv. unsafeStr avoids copying the token.
				v2, perr := strconv.ParseFloat(unsafeStr(data[i:end]), 64)
				if perr != nil {
					return end, ErrBadNumber
				}
				v = v2
			}
			f = v
			i = end
		}
		*out = append(*out, f)
		// Inter-token whitespace, in the shape the generator inlines: 0-1 bytes
		// resolve in one or two compares; only a real indentation run reaches the
		// SWAR SkipWSRun (which the compiler inlines here too — no call).
		if i < len(data) && data[i] <= ' ' {
			i++
			if i < len(data) && data[i] <= ' ' {
				i = SkipWSRun(data, i+1)
			}
		}
		if i >= len(data) {
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

// DecodeIntSlice decodes a JSON array of integers at data[i] into *out. The
// element parse mirrors ReadInt64OrNull byte for byte — SWAR 4-digit folds, a
// scalar tail, tolerated (truncated) fraction/exponent, overflow wrap — inlined
// into the loop so an element costs no call. The parsed int64 is converted to T
// exactly as the generated per-element code converted (wrapping, not
// saturating). A JSON null sets *out to nil; a null element decodes as 0.
func DecodeIntSlice[T intKind](out *[]T, data []byte, i int) (int, error) {
	if i >= len(data) {
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
	if *out == nil {
		if n := CountArrayScalars(data, i); n > 0 {
			*out = make([]T, 0, n)
		}
	}
	i++
	for {
		// Inter-token whitespace, in the shape the generator inlines: 0-1 bytes
		// resolve in one or two compares; only a real indentation run reaches the
		// SWAR SkipWSRun (which the compiler inlines here too — no call).
		if i < len(data) && data[i] <= ' ' {
			i++
			if i < len(data) && data[i] <= ' ' {
				i = SkipWSRun(data, i+1)
			}
		}
		if i >= len(data) {
			return i, ErrTruncated
		}
		if data[i] == ']' {
			return i + 1, nil
		}
		var n int64
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
				if i >= len(data) {
					return i, ErrBadNumber
				}
			}
			if data[i]-'0' > 9 { // unsigned: a byte below '0' wraps high
				return i, ErrBadNumber
			}
			for i+4 <= len(data) {
				w := binary.LittleEndian.Uint32(data[i : i+4])
				if !is4Digits(w) {
					break
				}
				n = n*10000 + int64(parse4Digits(w))
				i += 4
			}
			for i < len(data) {
				d := data[i] - '0'
				if d > 9 {
					break
				}
				n = n*10 + int64(d)
				i++
			}
			if i < len(data) && (data[i] == '.' || data[i] == 'e' || data[i] == 'E') {
				for i < len(data) {
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
		*out = append(*out, T(n))
		// Inter-token whitespace, in the shape the generator inlines: 0-1 bytes
		// resolve in one or two compares; only a real indentation run reaches the
		// SWAR SkipWSRun (which the compiler inlines here too — no call).
		if i < len(data) && data[i] <= ' ' {
			i++
			if i < len(data) && data[i] <= ' ' {
				i = SkipWSRun(data, i+1)
			}
		}
		if i >= len(data) {
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

// DecodeUintSlice is DecodeIntSlice for the unsigned kinds; the element parse
// mirrors ReadUint64OrNull.
func DecodeUintSlice[T uintKind](out *[]T, data []byte, i int) (int, error) {
	if i >= len(data) {
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
	if *out == nil {
		if n := CountArrayScalars(data, i); n > 0 {
			*out = make([]T, 0, n)
		}
	}
	i++
	for {
		// Inter-token whitespace, in the shape the generator inlines: 0-1 bytes
		// resolve in one or two compares; only a real indentation run reaches the
		// SWAR SkipWSRun (which the compiler inlines here too — no call).
		if i < len(data) && data[i] <= ' ' {
			i++
			if i < len(data) && data[i] <= ' ' {
				i = SkipWSRun(data, i+1)
			}
		}
		if i >= len(data) {
			return i, ErrTruncated
		}
		if data[i] == ']' {
			return i + 1, nil
		}
		var n uint64
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
			for i+4 <= len(data) {
				w := binary.LittleEndian.Uint32(data[i : i+4])
				if !is4Digits(w) {
					break
				}
				n = n*10000 + uint64(parse4Digits(w))
				i += 4
			}
			for i < len(data) {
				d := data[i] - '0'
				if d > 9 {
					break
				}
				n = n*10 + uint64(d)
				i++
			}
			if i < len(data) && (data[i] == '.' || data[i] == 'e' || data[i] == 'E') {
				for i < len(data) {
					c := data[i]
					if (c >= '0' && c <= '9') || c == '.' || c == 'e' || c == 'E' || c == '+' || c == '-' {
						i++
						continue
					}
					break
				}
			}
		}
		*out = append(*out, T(n))
		// Inter-token whitespace, in the shape the generator inlines: 0-1 bytes
		// resolve in one or two compares; only a real indentation run reaches the
		// SWAR SkipWSRun (which the compiler inlines here too — no call).
		if i < len(data) && data[i] <= ' ' {
			i++
			if i < len(data) && data[i] <= ' ' {
				i = SkipWSRun(data, i+1)
			}
		}
		if i >= len(data) {
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

// DecodeFloat64Array decodes a JSON array of numbers at data[i] into the
// fixed-size array whose backing the caller passes as out (out = arr[:]). It
// mirrors the generated fixed-array decoder: the array is zeroed, up to
// len(out) elements are decoded, extras are skipped, a short JSON array leaves
// the tail zero, and a JSON null leaves the array untouched. This is the
// per-point call for coordinate rings ([][2]float64, [][3]float64), where the
// generated form paid an extra call frame per coordinate.
func DecodeFloat64Array(out []float64, data []byte, i int) (int, error) {
	if i >= len(data) {
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
	for {
		// Inter-token whitespace, in the shape the generator inlines: 0-1 bytes
		// resolve in one or two compares; only a real indentation run reaches the
		// SWAR SkipWSRun (which the compiler inlines here too — no call).
		if i < len(data) && data[i] <= ' ' {
			i++
			if i < len(data) && data[i] <= ' ' {
				i = SkipWSRun(data, i+1)
			}
		}
		if i >= len(data) {
			return i, ErrTruncated
		}
		if data[i] == ']' {
			return i + 1, nil
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
		if i < len(data) && data[i] <= ' ' {
			i++
			if i < len(data) && data[i] <= ' ' {
				i = SkipWSRun(data, i+1)
			}
		}
		if i >= len(data) {
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
	if i >= len(data) {
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
	for {
		// Inter-token whitespace, in the shape the generator inlines: 0-1 bytes
		// resolve in one or two compares; only a real indentation run reaches the
		// SWAR SkipWSRun (which the compiler inlines here too — no call).
		if i < len(data) && data[i] <= ' ' {
			i++
			if i < len(data) && data[i] <= ' ' {
				i = SkipWSRun(data, i+1)
			}
		}
		if i >= len(data) {
			return i, ErrTruncated
		}
		if data[i] == ']' {
			return i + 1, nil
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
					if i >= len(data) {
						return i, ErrBadNumber
					}
				}
				if data[i]-'0' > 9 {
					return i, ErrBadNumber
				}
				var n int64
				for i+4 <= len(data) {
					w := binary.LittleEndian.Uint32(data[i : i+4])
					if !is4Digits(w) {
						break
					}
					n = n*10000 + int64(parse4Digits(w))
					i += 4
				}
				for i < len(data) {
					d := data[i] - '0'
					if d > 9 {
						break
					}
					n = n*10 + int64(d)
					i++
				}
				if i < len(data) && (data[i] == '.' || data[i] == 'e' || data[i] == 'E') {
					for i < len(data) {
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
		if i < len(data) && data[i] <= ' ' {
			i++
			if i < len(data) && data[i] <= ' ' {
				i = SkipWSRun(data, i+1)
			}
		}
		if i >= len(data) {
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
