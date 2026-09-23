package unstable

import "unsafe"

// DecodeFloat64Points decodes the JSON array at data[i] of fixed-size numeric
// arrays — a ring of coordinate points, [][2]float64 or [][3]float64, the
// shape of every GeoJSON document (canada, large-json) — into *out. T must be
// an array of float64; the generator routes only [][N]float64 fields here.
//
// It is the generated decoder for such a slice, element for element (the
// same reset, the same first-append capacity hint and growth, each point
// through DecodeFloat64Array, the same errors and the same partial result on
// one), with one addition: where a point begins, the SIMD kernel's points
// walk (parseFloatPoints, the VBMI body) takes as many points as it can in
// one call, straight into the slice's spare slots. The generated form made a
// Go call, an append and a kernel call per point, and on canada those were a
// third of what was left once the numbers themselves were cheap. A point the
// walk does not take — a null, a point of the wrong length, a number it does
// not convert — is decoded here exactly as before, and the walk resumes after.
func DecodeFloat64Points[T any](out *[]T, data []byte, i int) (int, error) {
	var zero T
	n := int(unsafe.Sizeof(zero) / 8)
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
	if len(*out) != 0 {
		*out = (*out)[:0]
	}
	// hold gives the point a productive call stopped at to the per-point path,
	// and run counts down to off on unproductive calls — 1 at first, 2 once
	// the walk has taken something, the batch slice readers' rule — so the
	// ring's first call, or the second in a row after the walk has taken
	// something, switches it off. A ring the walk cannot take costs one call;
	// a ring it has been taking survives a refusal at the first point of a
	// later call, which on canada, before the kernel refined its Eisel-Lemire
	// products, sent 8,573 of 55,563 points through the per-point path.
	run, hold := 0, false
	if useFloatRunLong {
		run = 1
	}
	i++
	for first := true; ; first = false {
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
			return i, ErrInvalidJSON
		}
		if run != 0 && !hold && data[i] == '[' {
			s := *out
			if s == nil {
				s = make([]T, 0, max(4, 256/max(1, int(unsafe.Sizeof(zero)))))
			} else if len(s) == cap(s) {
				s = GrowSlice(s)
			}
			flat := unsafe.Slice((*float64)(unsafe.Pointer(unsafe.SliceData(s))), cap(s)*n)[len(s)*n:]
			np, p, closed := parseFloatPoints(data, i, flat, n)
			*out = s[:len(s)+np]
			if np > 0 {
				i = p
				if closed != 0 {
					return i + 1, nil
				}
				// The walk stopped at a point it does not take: that point is
				// decoded below, and the walk resumes after it.
				hold, run = true, 2
				continue
			}
			run--
		}
		hold = false
		if *out == nil {
			*out = make([]T, 1, max(4, 256/max(1, int(unsafe.Sizeof(zero)))))
		} else {
			if len(*out) == cap(*out) {
				*out = GrowSlice(*out)
			}
			*out = append(*out, zero)
		}
		elem := unsafe.Slice((*float64)(unsafe.Pointer(&(*out)[len(*out)-1])), n)
		end, err := DecodeFloat64Array(elem, data, i)
		if err != nil {
			return end, err
		}
		i = end
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
