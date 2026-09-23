package unstable

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"math/bits"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

// checkFloatRun decodes data[i:] with the SIMD decimal-run kernel on and off,
// into a fresh (presized) target and into a reused one whose capacity fills
// part-way, and reports any difference in the values — compared bit for bit,
// so a -0 against a 0 or a last-place rounding difference fails — the end, or
// the error.
func checkFloatRun(t *testing.T, data []byte, i int) {
	t.Helper()
	type res struct {
		vals []float64
		end  int
		err  error
	}
	run := func(kernel bool, reuse bool) res {
		defer func(v bool) { useFloatRun = v }(useFloatRun)
		useFloatRun = kernel && floatRunHost // never on a CPU without it
		var s []float64
		if reuse {
			s = make([]float64, 2, 5)
		}
		end, err := DecodeFloat64Slice(&s, data, i)
		return res{s, end, err}
	}
	for _, body := range floatRunBodies() {
		useFloatRunLong = body.long
		for _, reuse := range []bool{false, true} {
			a, b := run(false, reuse), run(true, reuse)
			same := a.end == b.end && errors.Is(a.err, b.err) && errors.Is(b.err, a.err) && len(a.vals) == len(b.vals)
			for k := 0; same && k < len(a.vals); k++ {
				same = math.Float64bits(a.vals[k]) == math.Float64bits(b.vals[k])
			}
			if !same {
				t.Fatalf("%s reuse=%v: kernel differs on %q at %d:\n scalar: end=%d err=%v vals=%v\n kernel: end=%d err=%v vals=%v",
					body.name, reuse, data, i, a.end, a.err, a.vals, b.end, b.err, b.vals)
			}
		}
	}
	useFloatRunLong = floatRunLongHost
}

// floatRunHost and floatRunLongHost are the host's useFloatRun and
// useFloatRunLong, which the tests flip to run the scalar loop and both kernel
// bodies — never turning on one the CPU lacks.
var floatRunHost, floatRunLongHost = useFloatRun, useFloatRunLong

// floatRunBody is one kernel body the tests can select: long says it takes
// numbers of up to 19 digits (with Eisel-Lemire), and selecting it sets
// useFloatRunLong to that — which on amd64 is also what picks the body.
type floatRunBody struct {
	name   string
	long   bool
	digits int
}

// floatRunBodies lists the kernel bodies this host can run: on amd64 the AVX2
// one (numbers of up to 15 digits) and, with AVX-512 VBMI, the VBMI one (up to
// 19, with Eisel-Lemire); on arm64 the one NEON body, which takes up to 19.
func floatRunBodies() []floatRunBody {
	if runtime.GOARCH == "arm64" {
		return []floatRunBody{{"neon", true, 19}}
	}
	b := []floatRunBody{{"avx2", false, 15}}
	if floatRunLongHost {
		b = append(b, floatRunBody{"vbmi", true, 19})
	}
	return b
}

// TestFloatRunMatchesScalar holds the kernel to the scalar loop over generated
// arrays mixing every element shape it takes (1-15 digits split every way
// around the '.', signs, leading zeros, integers) with every shape it must hand
// back (exponents, 16+ digits, '+', null, a '.' without digits, garbage), every
// separator and whitespace shape, and every way of ending (']' after
// whitespace, a trailing comma, truncation, a stray byte).
func TestFloatRunMatchesScalar(t *testing.T) {
	rng := rand.New(rand.NewSource(11))
	digitsN := func(n int) string {
		var b strings.Builder
		for k := 0; k < n; k++ {
			b.WriteByte(byte('0' + rng.Intn(10)))
		}
		return b.String()
	}
	elem := func() string {
		switch r := rng.Intn(40); {
		case r < 22:
			l1 := 1 + rng.Intn(8)
			l2 := rng.Intn(16 - l1)
			s := digitsN(l1)
			if l2 > 0 {
				s += "." + digitsN(l2)
			}
			if rng.Intn(3) == 0 {
				s = "-" + s
			}
			return s
		case r < 25:
			return digitsN(1+rng.Intn(3)) + "." + digitsN(14+rng.Intn(6)) // 16+ digits
		case r < 28:
			return digitsN(1+rng.Intn(3)) + "." + digitsN(1+rng.Intn(4)) + []string{"e5", "E-3", "e+12"}[rng.Intn(3)]
		case r < 30:
			return "0." + strings.Repeat("0", rng.Intn(9)) + digitsN(1+rng.Intn(6))
		case r < 31:
			return "null"
		case r < 32:
			return []string{"+1", "1.", ".5", "-", "-.5", "--1", "1.2.3", "x", "", "0x10", "\"1\"", "]"}[rng.Intn(12)]
		case r < 34:
			return "00" + digitsN(1+rng.Intn(3)) + "." + digitsN(1+rng.Intn(3)) // leading zeros
		case r < 36:
			return "-0" + []string{"", ".0", ".000"}[rng.Intn(3)]
		default:
			return digitsN(1+rng.Intn(4)) + "." + digitsN(1+rng.Intn(8))
		}
	}
	seps := []string{",", ",", ", ", " ,", " , ", ",\n", ",\n        ", ",\t", ",,", " ", ""}
	closes := []string{"]", "]", " ]", "\n]", ",]", "", "]]", "}"}
	var inputs []string
	inputs = append(inputs,
		"[]", "[ ]", "[1]", "[1.5]", "[-0]", "[-0.0]", "[0.1,0.2,0.3]",
		"[123456789012345]", "[1234567890123456]", "[12345678.1234567]", "[1.23456789012345]",
		"[-9.99999999999999,-0.00000000000001]", "[1e5]", "[1,]", "[1 2]", "[1,,2]",
	)
	for len(inputs) < 5000 {
		n := rng.Intn(60)
		if rng.Intn(6) == 0 {
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
		lead := strings.Repeat(" ", rng.Intn(5))
		var trail string
		if rng.Intn(3) > 0 {
			trail = strings.Repeat(" }", rng.Intn(50))
		}
		data := []byte(lead + in + trail)
		checkFloatRun(t, data, len(lead))
	}
}

// TestFloatRunShapes converts every (integer digits, fraction digits) split
// the kernel takes, with and without a sign, at every lane of the window
// (a leading run of whitespace moves the number), and holds the value to
// strconv.ParseFloat bit for bit: each is Clinger's exact fast path, and a
// wrong shuffle control, a wrong divisor or a lost sign changes a bit.
func TestFloatRunShapes(t *testing.T) {
	if !useFloatRun {
		t.Skip("no SIMD decimal-run kernel on this machine")
	}
	defer func() { useFloatRunLong = floatRunLongHost }()
	pad := strings.Repeat(" ", 96)
	rng := rand.New(rand.NewSource(5))
	for _, body := range floatRunBodies() {
		useFloatRunLong = body.long
		for l1 := 1; l1 <= body.digits; l1++ {
			for l2 := 0; l1+l2 <= body.digits; l2++ {
				for lane := 0; lane < 64; lane += 1 + rng.Intn(3) {
					for _, neg := range []bool{false, true} {
						var b strings.Builder
						if neg {
							b.WriteByte('-')
						}
						for k := 0; k < l1; k++ {
							b.WriteByte(byte('0' + rng.Intn(10)))
						}
						if l2 > 0 {
							b.WriteByte('.')
							for k := 0; k < l2; k++ {
								b.WriteByte(byte('0' + rng.Intn(10)))
							}
						}
						num := b.String()
						want, err := strconv.ParseFloat(num, 64)
						if err != nil {
							t.Fatal(err)
						}
						data := []byte(strings.Repeat(" ", lane) + num + "," + pad)
						out := make([]float64, 2)
						n, p, closed := parseFloatRun(data, 0, out)
						if n == 0 && body.long && l1+l2 > 15 && p == lane {
							continue // Eisel-Lemire declined it: the scalar loop's
						}
						if n != 1 || math.Float64bits(out[0]) != math.Float64bits(want) || closed != 0 || p <= lane+len(num) {
							t.Fatalf("%s: %q at lane %d: n=%d v=%v (%#x) p=%d closed=%d, want %v (%#x)",
								body.name, num, lane, n, out[0], math.Float64bits(out[0]), p, closed, want, math.Float64bits(want))
						}
						// The same number as an array's last element, closed by ']'.
						data = []byte(strings.Repeat(" ", lane) + num + "]" + pad)
						n, p, closed = parseFloatRun(data, 0, out)
						if n != 1 || math.Float64bits(out[0]) != math.Float64bits(want) || closed != 1 || p != lane+len(num) {
							t.Fatalf("%s: %q] at lane %d: n=%d v=%v p=%d closed=%d, want %v p=%d closed", body.name, num, lane, n, out[0], p, closed, want, lane+len(num))
						}
					}
				}
			}
		}
	}
}

// TestFloatRunWindows drives the kernel through long uniform arrays of every
// separator style and number shape, at every alignment of its 64-byte window
// and 48-byte step, and requires it to take the whole array: a kernel that
// stopped early would be invisible to the differential.
func TestFloatRunWindows(t *testing.T) {
	if !useFloatRun {
		t.Skip("no SIMD decimal-run kernel on this machine")
	}
	defer func() { useFloatRunLong = floatRunLongHost }()
	seps := []string{",", ", ", " ,", " , ", ",\n", ",\n        ", ",\n                        ", ",\t"}
	shapes := []struct{ l1, l2 int }{{1, 0}, {1, 1}, {1, 12}, {2, 6}, {3, 3}, {1, 14}, {8, 7}, {15, 0}}
	for _, sep := range seps {
		for _, sh := range shapes {
			var b strings.Builder
			b.WriteByte('[')
			const count = 200
			want := make([]float64, count)
			for k := 0; k < count; k++ {
				if k > 0 {
					b.WriteString(sep)
				}
				v := fmt.Sprintf("%0*d", sh.l1, (k*7919+13)%pow10i(sh.l1))
				if sh.l2 > 0 {
					v += "." + fmt.Sprintf("%0*d", sh.l2, (k*104729+7)%pow10i(sh.l2))
				}
				if k%3 == 0 {
					v = "-" + v
				}
				b.WriteString(v)
				want[k], _ = strconv.ParseFloat(v, 64)
			}
			b.WriteString("]")
			arr := b.String()
			for lead := 0; lead < 64; lead += 1 + lead/8 {
				// Half the arrays are followed by more of the document in the
				// same window, as they are inside an object: the comma after
				// the ']' is not the last element's, and the kernel must
				// still close the array itself.
				tail := strings.Repeat(" ", 96)
				if lead%2 == 1 {
					tail = `, "next": [1, 2], "x": 3}` + tail
				}
				data := []byte(strings.Repeat(" ", lead) + arr + tail)
				checkFloatRun(t, data, lead)
				for _, body := range floatRunBodies() {
					useFloatRunLong = body.long
					out := make([]float64, count)
					n, p, closed := parseFloatRun(data, lead+1, out)
					if n != count || closed != 1 || data[p] != ']' {
						t.Fatalf("%s: sep %q shape %v lead %d: kernel took n=%d p=%d closed=%d, want all %d closed at ']'", body.name, sep, sh, lead, n, p, closed, count)
					}
					for k := range want {
						if math.Float64bits(out[k]) != math.Float64bits(want[k]) {
							t.Fatalf("%s: sep %q shape %v lead %d: element %d = %v, want %v", body.name, sep, sh, lead, k, out[k], want[k])
						}
					}
				}
			}
		}
	}
}

// TestFloatRunRandomValues is a value fuzz: two million random numbers of the
// kernel's shape, in arrays, each compared bit for bit with strconv.
func TestFloatRunRandomValues(t *testing.T) {
	if !useFloatRun {
		t.Skip("no SIMD decimal-run kernel on this machine")
	}
	defer func() { useFloatRunLong = floatRunLongHost }()
	for _, body := range floatRunBodies() {
		useFloatRunLong = body.long
		t.Run(body.name, func(t *testing.T) { testFloatRunRandomValues(t) })
	}
}

func testFloatRunRandomValues(t *testing.T) {
	rng := rand.New(rand.NewSource(99))
	const per = 500
	iters := 4000
	if testing.Short() {
		iters = 200
	}
	out := make([]float64, per)
	want := make([]float64, per)
	for it := 0; it < iters; it++ {
		var b strings.Builder
		b.WriteByte('[')
		for k := 0; k < per; k++ {
			if k > 0 {
				b.WriteByte(',')
			}
			l := 1 + rng.Intn(15)
			l1 := 1 + rng.Intn(l)
			var num strings.Builder
			if rng.Intn(2) == 0 {
				num.WriteByte('-')
			}
			for d := 0; d < l1; d++ {
				num.WriteByte(byte('0' + rng.Intn(10)))
			}
			if l > l1 {
				num.WriteByte('.')
				for d := 0; d < l-l1; d++ {
					num.WriteByte(byte('0' + rng.Intn(10)))
				}
			}
			want[k], _ = strconv.ParseFloat(num.String(), 64)
			b.WriteString(num.String())
		}
		b.WriteString("]" + strings.Repeat(" ", 96))
		data := []byte(b.String())
		n, _, closed := parseFloatRun(data, 1, out)
		if n != per || closed != 1 {
			t.Fatalf("iteration %d: n=%d closed=%d", it, n, closed)
		}
		for k := 0; k < per; k++ {
			if math.Float64bits(out[k]) != math.Float64bits(want[k]) {
				t.Fatalf("iteration %d element %d: %v, want %v", it, k, out[k], want[k])
			}
		}
	}
}

// TestFloatRunStopsAtSign pins where the kernel stops on an element that
// begins with a '-': at the '-' itself, whatever made it stop — the output
// being full, an exponent, a '.' without digits, a non-digit. The kernel
// measures an element past its sign before it decides, and a stop that left
// the cursor there would hand the scalar loop the digits alone: a negative
// number read back as positive, silently. (The first version did exactly that
// when the output filled up, and only a target whose presize had failed ever
// reached it.)
func TestFloatRunStopsAtSign(t *testing.T) {
	if !useFloatRun {
		t.Skip("no SIMD decimal-run kernel on this machine")
	}
	defer func() { useFloatRunLong = floatRunLongHost }()
	for _, body := range floatRunBodies() {
		useFloatRunLong = body.long
		testFloatRunStopsAtSign(t)
	}
}

func testFloatRunStopsAtSign(t *testing.T) {
	pad := strings.Repeat(" ", 96)
	for _, c := range []struct {
		in     string
		avail  int
		wantN  int
		wantP  int
		closed int
	}{
		{"[1,-2,3]", 1, 1, 3, 0},
		{"[1, -2,3]", 1, 1, 4, 0},
		{"[1,-2]", 1, 1, 3, 0},
		{"[1, -2 ]", 1, 1, 4, 0},
		{"[-1e5,2]", 4, 0, 1, 0},
		{"[1,-1e5,2]", 4, 1, 3, 0},
		{"[1,-1.,2]", 4, 1, 3, 0},
		{"[1,-x,2]", 4, 1, 3, 0},
		{"[1,-,2]", 4, 1, 3, 0},
		{"[1,-1.5 x,2]", 4, 1, 3, 0},
		{"[1,-1.5 x]", 4, 1, 3, 0},
		{"[1,-123456789012345678901]", 4, 1, 3, 0}, // 21 digits: past either body
	} {
		out := make([]float64, c.avail)
		data := []byte(c.in + pad)
		n, p, closed := parseFloatRun(data, 1, out)
		// wantP is the element's first byte that is not whitespace, which is
		// where a kernel that measured the element stops (amd64). One that
		// stops at the element without measuring it (arm64, whose capacity
		// bound is per block) returns the start of its region instead. Both
		// are states the scalar loop resumes from, since it skips whitespace
		// first; what must not happen is a p past the '-'.
		if p < c.wantP && len(bytes.TrimLeft(data[p:c.wantP], " \t\n\r")) == 0 {
			p = c.wantP
		}
		if n != c.wantN || p != c.wantP || closed != c.closed {
			t.Errorf("%q (room %d): n=%d p=%d closed=%d, want n=%d p=%d closed=%d", c.in, c.avail, n, p, closed, c.wantN, c.wantP, c.closed)
		}
	}
}

// TestFloatRunEiselLemire is the VBMI body's own fuzz: arrays of 16- to
// 19-digit numbers — past Clinger, so nearly every one takes the Eisel-Lemire
// arm — split every way around the '.', some with leading zeros that bring the
// mantissa back under 2^53, each decoded with the kernel and compared bit for
// bit with strconv. The kernel may decline an element (one eiselLemire64
// would refine or decline), which the scalar loop then converts; the values
// must come out the same either way. It then walks each array with the kernel
// alone, and every element handed back must be one of those by the
// assembly's own conditions (kernelDeclines) — a body that declined all it saw
// would pass the comparison.
func TestFloatRunEiselLemire(t *testing.T) {
	if !floatRunLongHost {
		t.Skip("no AVX-512 VBMI body on this machine")
	}
	rng := rand.New(rand.NewSource(1234))
	const per = 400
	iters := 3000
	if testing.Short() {
		iters = 150
	}
	var took, total int
	for it := 0; it < iters; it++ {
		var b strings.Builder
		b.WriteByte('[')
		want := make([]float64, per)
		for k := 0; k < per; k++ {
			if k > 0 {
				b.WriteString([]string{",", ", "}[rng.Intn(2)])
			}
			l := 16 + rng.Intn(4)
			if rng.Intn(10) == 0 {
				l = 1 + rng.Intn(19)
			}
			l1 := 1 + rng.Intn(l)
			var num strings.Builder
			if rng.Intn(2) == 0 {
				num.WriteByte('-')
			}
			for d := 0; d < l1; d++ {
				c := byte('0' + rng.Intn(10))
				if d == 0 && l1 > 1 && rng.Intn(4) != 0 {
					c = byte('1' + rng.Intn(9))
				}
				num.WriteByte(c)
			}
			if l > l1 {
				num.WriteByte('.')
				zeros := 0
				if rng.Intn(5) == 0 {
					zeros = rng.Intn(l - l1)
				}
				for d := 0; d < l-l1; d++ {
					if d < zeros {
						num.WriteByte('0')
					} else {
						num.WriteByte(byte('0' + rng.Intn(10)))
					}
				}
			}
			want[k], _ = strconv.ParseFloat(num.String(), 64)
			b.WriteString(num.String())
		}
		b.WriteString("]" + strings.Repeat(" ", 160)) // past the last window
		data := []byte(b.String())
		var s []float64
		if _, err := DecodeFloat64Slice(&s, data, 0); err != nil {
			t.Fatal(err)
		}
		for k := 0; k < per; k++ {
			if math.Float64bits(s[k]) != math.Float64bits(want[k]) {
				t.Fatalf("iteration %d element %d: %v (%#x), want %v (%#x)", it, k, s[k], math.Float64bits(s[k]), want[k], math.Float64bits(want[k]))
			}
		}
		// What the kernel alone takes, resuming after each element it hands
		// back — and every one it hands back must be a number it declines by
		// design (kernelDeclines).
		out := make([]float64, per)
		for i := 1; ; {
			n, p, closed := parseFloatRun(data, i, out)
			took += n
			if closed != 0 {
				break
			}
			j := p
			for data[j] != ',' && data[j] != ']' {
				j++
			}
			if elem := strings.TrimSpace(string(data[p:j])); !kernelDeclines(elem) {
				t.Fatalf("iteration %d: the kernel handed back %q, which it converts", it, elem)
			}
			if data[j] == ']' {
				break
			}
			i = j + 1
		}
		total += per
	}
	// The refusals are the Eisel-Lemire cases eiselLemire64 refines. They are
	// not uniform: a value exactly representable in binary (a fraction of .5,
	// .25, .125, …) is just undershot by the truncated power of ten, which
	// leaves the product's low bits all ones — so a one-digit fraction declines
	// 9% of the time, two digits 2.5%, three 0.6%, and fifteen (a canada
	// coordinate) 0.13%. This generator's short fractions put it near 1%.
	if took < total*98/100 {
		t.Fatalf("the kernel converted %d of %d numbers; declines should be rare", took, total)
	}
}

// kernelDeclines reports whether the kernel's long conversion (the amd64 VBMI
// body's LONGCONV, the arm64 kernel's) hands the number s (of that shape: a
// sign, digits, a '.' and digits, 19 digits at most) back to the scalar loop
// by design. Both decline an exact halfway value. Where eiselLemire64 would
// refine the product with the power's low word — its low nine bits all ones
// and xLo + man wrapping — the amd64 body declines, and the arm64 one refines
// as eiselLemire64 does and declines only a product still ambiguous after it.
// The conditions are the assembly's, restated.
func kernelDeclines(s string) bool {
	s = strings.TrimPrefix(s, "-")
	ip, fp, _ := strings.Cut(s, ".")
	man, err := strconv.ParseUint(ip+fp, 10, 64)
	if err != nil || len(ip+fp) > 19 || man>>53 == 0 {
		return false // not this body's shape, or Clinger's, which never declines
	}
	man <<= bits.LeadingZeros64(man)
	pow := detailedPowersOfTen[-len(fp)-detailedPowersOfTenMinExp10]
	xHi, xLo := bits.Mul64(man, pow[1])
	if xHi&0x1FF == 0x1FF && xLo+man < man {
		if runtime.GOARCH != "arm64" {
			return true
		}
		yHi, yLo := bits.Mul64(man, pow[0])
		mergedHi, mergedLo := xHi, xLo+yHi
		if mergedLo < xLo {
			mergedHi++
		}
		if mergedHi&0x1FF == 0x1FF && mergedLo+1 == 0 && yLo+man < man {
			return true
		}
		xHi, xLo = mergedHi, mergedLo
	}
	msb := xHi >> 63
	return xLo == 0 && xHi&0x1FF == 0 && (xHi>>(msb+9))&3 == 1
}

// needsRefine reports whether eiselLemire64 refines the number s's product
// with the power's low word — the numbers the amd64 kernel declines.
func needsRefine(s string) bool {
	s = strings.TrimPrefix(s, "-")
	ip, fp, _ := strings.Cut(s, ".")
	man, err := strconv.ParseUint(ip+fp, 10, 64)
	if err != nil || len(ip+fp) > 19 || man>>53 == 0 {
		return false
	}
	man <<= bits.LeadingZeros64(man)
	xHi, xLo := bits.Mul64(man, detailedPowersOfTen[-len(fp)-detailedPowersOfTenMinExp10][1])
	return xHi&0x1FF == 0x1FF && xLo+man < man
}

// TestFloatRunRefines holds the arm64 kernel's Eisel-Lemire refinement to
// strconv: numbers of the shape canada is full of — a six-decimal value printed
// back to 17 digits, "46.851662000000033" — filtered to the ones whose product
// needs the power's low word, each of which the kernel must convert itself, in
// the flat walk and the points walk, bit for bit. A declined number would still
// decode correctly, through the scalar loop, so what this pins is that the
// kernel does not hand these back — for this shape the refinement only confirms
// the high product. The refinement's arithmetic, where it changes the result,
// is TestFloatRunEiselLemire's: dropping it, or never declining after it, each
// put a 1-ulp error there.
func TestFloatRunRefines(t *testing.T) {
	if runtime.GOARCH != "arm64" || !useFloatRun {
		t.Skip("the arm64 kernel refines; the amd64 one declines these")
	}
	rng := rand.New(rand.NewSource(46))
	var nums []string
	for tries := 0; len(nums) < 2000 && tries < 2000000; tries++ {
		// A six-decimal coordinate stored as a double and printed back with
		// fifteen fraction digits: canada's "43.420273000000009".
		v, _ := strconv.ParseFloat(fmt.Sprintf("%d.%06d", rng.Intn(360)-180, rng.Intn(1000000)), 64)
		num := strconv.FormatFloat(v, 'f', 15, 64)
		if len(strings.Trim(strings.ReplaceAll(strings.TrimPrefix(num, "-"), ".", ""), "0")) > 19 {
			continue
		}
		if needsRefine(num) && !kernelDeclines(num) {
			nums = append(nums, num)
		}
	}
	if len(nums) < 100 {
		t.Fatalf("premise: only %d numbers need the refinement", len(nums))
	}
	pad := strings.Repeat(" ", 100)
	out := make([]float64, 4)
	for _, num := range nums {
		want, err := strconv.ParseFloat(num, 64)
		if err != nil {
			t.Fatal(err)
		}
		n, _, closed := parseFloatRun([]byte("["+num+","+num+"]"+pad), 1, out)
		if n != 2 || closed != 1 || math.Float64bits(out[0]) != math.Float64bits(want) || math.Float64bits(out[1]) != math.Float64bits(want) {
			t.Fatalf("flat %s: n=%d closed=%d %v, want %v", num, n, closed, out[:n], want)
		}
		np, _, closed := parseFloatPoints([]byte("[["+num+","+num+"]]"+pad), 1, out, 2)
		if np != 1 || closed != 1 || math.Float64bits(out[0]) != math.Float64bits(want) || math.Float64bits(out[1]) != math.Float64bits(want) {
			t.Fatalf("points %s: np=%d closed=%d %v, want %v", num, np, closed, out[:2], want)
		}
	}
}

// TestFloatRunFixedArrays holds DecodeFloat64Array to itself with the kernel
// off, through every body, for fixed arrays of 1 to 5 slots over JSON arrays
// shorter than, equal to and longer than them (the extras are skipped), of
// every element shape the kernel takes or refuses: the reader hands the
// kernel the array's remaining slots, and a coordinate point is one call
// closed at its ']'.
func TestFloatRunFixedArrays(t *testing.T) {
	defer func() { useFloatRunLong = floatRunLongHost }()
	rng := rand.New(rand.NewSource(77))
	elems := []string{"1", "-2.5", "0.000123", "-65.613616999999977", "43.420273000000009", "1e5", "null", "12345678901234567890", "-0.0", "3.14159265358979"}
	for it := 0; it < 3000; it++ {
		m := rng.Intn(7)
		var b strings.Builder
		b.WriteString("[")
		for k := 0; k < m; k++ {
			if k > 0 {
				b.WriteString([]string{",", ", ", " ,\n  "}[rng.Intn(3)])
			}
			b.WriteString(elems[rng.Intn(len(elems))])
		}
		b.WriteString([]string{"]", " ]", ",]", "]]"}[rng.Intn(4)])
		tail := []string{"", strings.Repeat(" ", 80), `, [1,2], "k": 3}` + strings.Repeat(" ", 70)}[rng.Intn(3)]
		data := []byte(b.String() + tail)
		for slots := 1; slots <= 5; slots++ {
			type res struct {
				vals []float64
				end  int
				err  error
			}
			run := func(kernel bool) res {
				defer func(v bool) { useFloatRun = v }(useFloatRun)
				useFloatRun = kernel && floatRunHost
				out := make([]float64, slots)
				for k := range out {
					out[k] = 99 // the reader must clear what it does not fill
				}
				end, err := DecodeFloat64Array(out, data, 0)
				return res{out, end, err}
			}
			for _, body := range floatRunBodies() {
				useFloatRunLong = body.long
				a, c := run(false), run(true)
				same := a.end == c.end && errors.Is(a.err, c.err) && errors.Is(c.err, a.err)
				for k := 0; same && k < slots; k++ {
					same = math.Float64bits(a.vals[k]) == math.Float64bits(c.vals[k])
				}
				if !same {
					t.Fatalf("%s [%d]float64 on %q:\n scalar: end=%d err=%v %v\n kernel: end=%d err=%v %v", body.name, slots, data, a.end, a.err, a.vals, c.end, c.err, c.vals)
				}
			}
		}
	}
}

// BenchmarkFloatRunShapes is the kernel alone over 4000-element arrays of one
// number shape each: the short decimals of mesh and numbers (Clinger in the
// AVX2 body's conversion) and the 17-digit coordinates of canada and
// large-json (the VBMI body's gather and Eisel-Lemire), compact and
// ", "-separated.
func BenchmarkFloatRunShapes(b *testing.B) {
	if !useFloatRun {
		b.Skip("no SIMD decimal-run kernel on this machine")
	}
	for _, sh := range []struct{ name, num, sep string }{
		{"short", "0.0636837780476", ","},
		{"short_space", "-0.0636837780476", ", "},
		{"digits6", "0.270354", ","},
		{"long17", "-65.613616999999977", ","},
		{"long16", "37.80848009696725", ","},
	} {
		data := []byte("[" + strings.Repeat(sh.num+sh.sep, 3999) + sh.num + "]" + strings.Repeat(" ", 96))
		out := make([]float64, 4000)
		b.Run(sh.name, func(b *testing.B) {
			b.SetBytes(int64(len(data)))
			for i := 0; i < b.N; i++ {
				if n, _, _ := parseFloatRun(data, 1, out); n != 4000 {
					b.Fatal(n)
				}
			}
		})
	}
}

// TestFloatRunRoundsIntoExponent holds the VBMI body to strconv on numbers
// whose Eisel-Lemire mantissa rounds up to 2^53 — just under a power of two —
// which LONGCONV absorbs with no test of its own: the carry out of the
// mantissa lands in the exponent field (see the macro's comment). Each must
// be converted by the kernel itself, not handed back, or the carry path would
// go untested; both walks, flat and points, are checked.
func TestFloatRunRoundsIntoExponent(t *testing.T) {
	if !floatRunLongHost {
		t.Skip("no AVX-512 VBMI body on this machine")
	}
	nums := []string{
		"0.999999999999999999", "1.999999999999999999", "3.99999999999999999", "7.99999999999999999",
		"18014398509481983", "36028797018963967", "9223372036854775807", "4611686018427387903",
		"0.99999999999999999", "63.9999999999999999", "1023.99999999999999",
	}
	pad := strings.Repeat(" ", 100)
	for _, num := range nums {
		for _, s := range []string{num, "-" + num} {
			want, err := strconv.ParseFloat(s, 64)
			if err != nil {
				t.Fatal(err)
			}
			if math.Float64bits(want)&(1<<52-1) != 0 {
				t.Fatalf("%s: premise: %v is not a power of two", s, want)
			}
			if kernelDeclines(s) {
				t.Fatalf("%s: the kernel declines it, so its carry goes untested", s)
			}
			out := make([]float64, 4)
			n, _, closed := parseFloatRunV([]byte("["+s+","+s+"]"+pad), 1, out)
			if n != 2 || closed != 1 || math.Float64bits(out[0]) != math.Float64bits(want) || math.Float64bits(out[1]) != math.Float64bits(want) {
				t.Errorf("flat %s: n=%d closed=%d %v, want %v", s, n, closed, out[:n], want)
			}
			np, _, closed := parseFloatPoints([]byte("[["+s+","+s+"]]"+pad), 1, out, 2)
			if np != 1 || closed != 1 || math.Float64bits(out[0]) != math.Float64bits(want) || math.Float64bits(out[1]) != math.Float64bits(want) {
				t.Errorf("points %s: np=%d closed=%d %v, want %v", s, np, closed, out[:2*np], want)
			}
		}
	}
}
