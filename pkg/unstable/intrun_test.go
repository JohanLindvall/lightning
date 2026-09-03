package unstable

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

// TestIntRunMatchesScalar is the SIMD integer-run kernel's contract: with
// useIntRun on and off, DecodeIntSlice/DecodeUintSlice must return the same
// values, the same end and the same error on every input, because the kernel
// only ever consumes elements the scalar loop would have read identically and
// hands everything else back at a loop-top position. The generator drives
// every digit count at every block offset (element lengths 1-12 with random
// separators shift the offsets), signs, nulls, decimals, exponents, long ids,
// every whitespace shape (none, one space, a tab, a newline-and-indent run
// longer than a block), trailing commas, missing commas, garbage bytes, a
// truncated array, arrays that end within a block of the buffer's end, and
// a reused target whose capacity fills mid-run.
func TestIntRunMatchesScalar(t *testing.T) {
	if !useIntRun {
		t.Skip("no SIMD integer-run kernel on this machine")
	}
	defer func(v bool) { useIntRun = v }(useIntRun)
	rng := rand.New(rand.NewSource(7))
	elem := func() string {
		switch r := rng.Intn(40); {
		case r < 24:
			return digits(rng, 1+rng.Intn(8))
		case r < 28:
			return digits(rng, 9+rng.Intn(11))
		case r < 31:
			return "-" + digits(rng, 1+rng.Intn(6))
		case r < 33:
			return "null"
		case r < 35:
			return digits(rng, 1+rng.Intn(4)) + "." + digits(rng, 1+rng.Intn(3))
		case r < 36:
			return digits(rng, 1+rng.Intn(3)) + "e" + digits(rng, 1)
		case r < 37:
			return "0" + digits(rng, 1+rng.Intn(6)) // leading zeros
		case r < 38:
			return []string{"x", "-", "+5", "1-2", "", "\"s\"", "tru", "]"}[rng.Intn(8)]
		default:
			return digits(rng, 4)
		}
	}
	seps := []string{",", ",", ",", ", ", " ,", " , ", ",\t", ",\n", ",\n                ", ",\r\n  ", " ", ",,", ""}
	closes := []string{"]", "]", " ]", "\n]", ",]", " , ]", "", "]]", "}"}
	var inputs []string
	// Hand-picked shapes first.
	inputs = append(inputs,
		"[]", "[ ]", "[1]", "[1,2]", "[12345678]", "[123456789]", "[0,0,-0]",
		"[1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17]",
		"[1,22,333,4444,55555,666666,7777777,88888888,999999999,1,22,333,4444]",
		"[9223372036854775807,9223372036854775808,18446744073709551615,18446744073709551616]",
		"[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16]", // a comma on the block's last byte
		"[1234567,1234567,1234567,1234567,1234567,1234567,1234567,1234567,1234567]",
		"[12345678,12345678,12345678,12345678,12345678,12345678,12345678,12345678,12345678]",
		"[1,]", "[1,2,]", "[1 2]", "[1,,2]", "[,1]", "[1", "[1,", "[1,2",
		"[1                    ,2]", "[1,                    2]", "[                    1,2]",
		"[1.5,2.5,3.5,4.5,5.5,6.5,7.5,8.5,9.5]", "[-1,-2,-3,-4,-5,-6,-7,-8,-9,-10]",
		"[null,null,null,null,null,null]", "[1,null,2,null,3,null,4,null,5]",
	)
	for len(inputs) < 4000 {
		n := rng.Intn(60)
		if rng.Intn(8) == 0 {
			n = 60 + rng.Intn(300)
		}
		var b strings.Builder
		b.WriteString([]string{"[", "[", "[ ", "[\n    "}[rng.Intn(4)])
		for k := 0; k < n; k++ {
			if k > 0 {
				b.WriteString(seps[rng.Intn(len(seps))])
			}
			b.WriteString(elem())
		}
		b.WriteString(closes[rng.Intn(len(closes))])
		inputs = append(inputs, b.String())
	}
	for _, in := range inputs {
		// The array at a random offset in a buffer with random padding after
		// it, so the kernel's lookahead limit (16 bytes on amd64, 64 on arm64)
		// falls at every distance from the array's end.
		lead := strings.Repeat(" ", rng.Intn(5))
		var trail string
		if rng.Intn(3) > 0 {
			trail = strings.Repeat(" }", rng.Intn(24))
		}
		data := []byte(lead + in + trail)
		i := len(lead)
		checkIntRunKinds(t, data, i)
	}
}

func digits(rng *rand.Rand, n int) string {
	var b strings.Builder
	for k := 0; k < n; k++ {
		b.WriteByte(byte('0' + rng.Intn(10)))
	}
	return b.String()
}

// checkIntRunKinds decodes data[i:] as every integer kind with the kernel
// enabled and disabled, into fresh targets and into targets whose backing
// fills partway through, and reports any difference.
func checkIntRunKinds(t *testing.T, data []byte, i int) {
	t.Helper()
	type res struct {
		vals any
		end  int
		err  error
	}
	run := func(kernel bool, mk func() (func() (int, error), func() any)) res {
		useIntRun = kernel
		dec, get := mk()
		end, err := dec()
		return res{get(), end, err}
	}
	kinds := []struct {
		name string
		mk   func() (func() (int, error), func() any)
	}{
		{"int64", func() (func() (int, error), func() any) {
			var s []int64
			return func() (int, error) { return DecodeIntSlice(&s, data, i) }, func() any { return s }
		}},
		{"int32", func() (func() (int, error), func() any) {
			var s []int32
			return func() (int, error) { return DecodeIntSlice(&s, data, i) }, func() any { return s }
		}},
		{"int16", func() (func() (int, error), func() any) {
			var s []int16
			return func() (int, error) { return DecodeIntSlice(&s, data, i) }, func() any { return s }
		}},
		{"int", func() (func() (int, error), func() any) {
			var s []int
			return func() (int, error) { return DecodeIntSlice(&s, data, i) }, func() any { return s }
		}},
		{"uint64", func() (func() (int, error), func() any) {
			var s []uint64
			return func() (int, error) { return DecodeUintSlice(&s, data, i) }, func() any { return s }
		}},
		{"uint8", func() (func() (int, error), func() any) {
			var s []uint8
			return func() (int, error) { return DecodeUintSlice(&s, data, i) }, func() any { return s }
		}},
		// Reused targets: the kernel must respect the spare capacity and the
		// scalar loop grow past it, for both the direct and the scratch path.
		{"int64-cap3", func() (func() (int, error), func() any) {
			s := make([]int64, 1, 3)
			return func() (int, error) { return DecodeIntSlice(&s, data, i) }, func() any { return s }
		}},
		{"int32-cap5", func() (func() (int, error), func() any) {
			s := make([]int32, 2, 5)
			return func() (int, error) { return DecodeIntSlice(&s, data, i) }, func() any { return s }
		}},
		{"uint64-cap1", func() (func() (int, error), func() any) {
			s := make([]uint64, 0, 1)
			return func() (int, error) { return DecodeUintSlice(&s, data, i) }, func() any { return s }
		}},
	}
	for _, k := range kinds {
		a := run(false, k.mk)
		b := run(true, k.mk)
		if a.end != b.end || !errors.Is(b.err, a.err) || !errors.Is(a.err, b.err) || !reflect.DeepEqual(a.vals, b.vals) {
			t.Fatalf("%s: kernel differs on %q at %d:\n scalar: end=%d err=%v vals=%v\n kernel: end=%d err=%v vals=%v",
				k.name, data, i, a.end, a.err, a.vals, b.end, b.err, b.vals)
		}
	}
}

// TestParseIntRunDirect pins the kernel's own contract on a few shapes: what
// it writes, where it stops, and the closed flag.
func TestParseIntRunDirect(t *testing.T) {
	if !useIntRun {
		t.Skip("no SIMD integer-run kernel on this machine")
	}
	pad := strings.Repeat(" ", 80) // past either kernel's lookahead (16 bytes amd64, 64 arm64)
	cases := []struct {
		in     string
		i      int
		avail  int
		want   []int64
		p      int
		closed int
	}{
		{"1,22,333,4444,55555]" + pad, 0, 16, []int64{1, 22, 333, 4444, 55555}, 19, 1},
		{"1,22,333,4444,55555," + pad, 0, 16, []int64{1, 22, 333, 4444, 55555}, 20, 0},
		{"12345678,1]" + pad, 0, 16, []int64{12345678, 1}, 10, 1},
		{"123456789,1]" + pad, 0, 16, nil, 0, 0}, // 9 digits: not this kernel's
		{"1, 2 ,3 , 4]" + pad, 0, 16, []int64{1, 2, 3, 4}, 11, 1},
		{"1,-2,3]" + pad, 0, 16, []int64{1}, 2, 0},                 // stops at the sign
		{"1,null,3]" + pad, 0, 16, []int64{1}, 2, 0},               // and at null
		{"1,2.5,3]" + pad, 0, 16, []int64{1}, 2, 0},                // and at a decimal
		{"1,2,3,4,5,6,7,8,9]" + pad, 0, 3, []int64{1, 2, 3}, 6, 0}, // out full
		{"7]" + pad, 0, 16, []int64{7}, 1, 1},
		{"1,2", 0, 16, nil, 0, 0}, // fewer than a block: nothing
		{"0000,0001,0010]" + pad, 0, 16, []int64{0, 1, 10}, 14, 1},
		{"   5,6]" + pad, 0, 16, []int64{5, 6}, 6, 1},
	}
	for _, c := range cases {
		out := make([]int64, c.avail)
		n, p, closed := parseIntRun([]byte(c.in), c.i, out)
		got := out[:n]
		if len(got) == 0 {
			got = nil
		}
		if !reflect.DeepEqual(got, c.want) || p != c.p || closed != c.closed {
			t.Errorf("parseIntRun(%q): got %v p=%d closed=%d, want %v p=%d closed=%d", c.in, got, p, closed, c.want, c.p, c.closed)
		}
	}
	// Every (offset, length) shuffle-table row: an L-digit number at block
	// offset b, reached by padding with b leading spaces of a fresh call.
	for b := 0; b < 16; b++ {
		for l := 1; l <= 8 && b+l < 16; l++ {
			num := fmt.Sprintf("%0*d", l, 1234567890%pow10i(l))
			in := strings.Repeat(" ", b) + num + "," + pad
			out := make([]int64, 4)
			n, p, _ := parseIntRun([]byte(in), 0, out)
			want, err := strconv.ParseInt(num, 10, 64)
			if err != nil {
				t.Fatal(err)
			}
			if n != 1 || out[0] != want || p != b+l+1 {
				t.Errorf("b=%d l=%d %q: n=%d v=%d p=%d, want 1 %d %d", b, l, in[:b+l+1], n, out[0], p, want, b+l+1)
			}
		}
	}
}

func pow10i(n int) int {
	r := 1
	for ; n > 0; n-- {
		r *= 10
	}
	return r
}

// intRunArray is a JSON array of n four-digit-or-shorter integers with sep
// between them, padded so neither kernel's lookahead runs out.
func intRunArray(n int, sep string) []byte {
	var b strings.Builder
	b.WriteByte('[')
	for i := 0; i < n; i++ {
		if i > 0 {
			b.WriteString(sep)
		}
		b.WriteString(strconv.Itoa((i * 7919) % 10000))
	}
	b.WriteString("]")
	b.WriteString(strings.Repeat(" ", 80))
	return []byte(b.String())
}

// BenchmarkDecodeIntSliceRun measures DecodeIntSlice on a 4000-element array
// of short integers, compact and ", "-separated, with the SIMD run kernel on
// and (where it exists) off; the kernel-only row is the assembly alone. This
// is the micro that sized the kernel (CLAUDE.md, 2026-09-02).
func BenchmarkDecodeIntSliceRun(b *testing.B) {
	for _, sep := range []string{",", ", "} {
		data := intRunArray(4000, sep)
		modes := []bool{false}
		if useIntRun {
			modes = append(modes, true)
		}
		for _, kernel := range modes {
			name := "scalar"
			if kernel {
				name = "kernel"
			}
			b.Run(name+"/sep"+strconv.Quote(sep), func(b *testing.B) {
				defer func(v bool) { useIntRun = v }(useIntRun)
				useIntRun = kernel
				b.SetBytes(int64(len(data)))
				var s []int64
				for i := 0; i < b.N; i++ {
					s = s[:0]
					if _, err := DecodeIntSlice(&s, data, 0); err != nil {
						b.Fatal(err)
					}
				}
			})
		}
	}
	if useIntRun {
		data := intRunArray(4000, ",")
		out := make([]int64, 4000)
		b.Run("kernel-only", func(b *testing.B) {
			b.SetBytes(int64(len(data)))
			for i := 0; i < b.N; i++ {
				if n, _, _ := parseIntRun(data, 1, out); n < 3999 {
					b.Fatal(n)
				}
			}
		})
	}
}

// BenchmarkParseIntRunShapes is the kernel alone over arrays of one element
// length each, 4000 elements a run: the instrument that shaped the arm64
// kernel (CLAUDE.md, 2026-09-02). Two elements a block that never straddled
// a 16-byte block against two that always did exposed the block-to-block
// address chain, and the per-element counts under perf stat then said what
// the walk was bound by after each change. The scalar loop is not run here;
// DecodeIntSliceRun has that comparison.
func BenchmarkParseIntRunShapes(b *testing.B) {
	if !useIntRun {
		b.Skip("no SIMD integer-run kernel on this machine")
	}
	for _, sh := range []struct{ num, sep string }{
		{"1", ","}, {"12", ","}, {"123", ","}, {"1234", ","}, {"123456", ","}, {"1234567", ","}, {"1234", ", "},
	} {
		// 4000 elements, the last one closed by ']' with no separator after
		// it (a separator there would be a trailing comma, which the kernel
		// rightly stops in front of).
		data := []byte(strings.Repeat(sh.num+sh.sep, 3999) + sh.num + "]" + strings.Repeat(" ", 80))
		out := make([]int64, 4000)
		b.Run(strconv.Quote(sh.num+sh.sep), func(b *testing.B) {
			b.SetBytes(int64(len(data)))
			for i := 0; i < b.N; i++ {
				if n, _, _ := parseIntRun(data, 0, out); n != 4000 {
					b.Fatal(n)
				}
			}
		})
	}
}
