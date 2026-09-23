package unstable

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

// genRing builds a JSON ring of points of n numbers each, drawing element
// shapes from both sides of the kernel's line.
func genRing(rng *rand.Rand, n, count int, sep string, odd bool) string {
	num := func() string {
		switch r := rng.Intn(20); {
		case r < 8:
			return strconv.FormatFloat((rng.Float64()-0.5)*400, 'f', 15, 64) // 17-18 digits
		case r < 12:
			return strconv.FormatFloat((rng.Float64()-0.5)*400, 'f', -1, 64) // shortest
		case r < 15:
			return strconv.Itoa(rng.Intn(2000) - 1000)
		case r < 16 && odd:
			return []string{"1e5", "-2.5E-3", "null", "1.", "+1", "12345678901234567890123"}[rng.Intn(6)]
		default:
			return strconv.FormatFloat(rng.Float64(), 'f', 6, 64)
		}
	}
	var b strings.Builder
	b.WriteString("[")
	for k := 0; k < count; k++ {
		if k > 0 {
			b.WriteString(sep)
		}
		m := n
		if odd && rng.Intn(15) == 0 {
			m = []int{0, 1, n + 1}[rng.Intn(3)]
		}
		if odd && rng.Intn(40) == 0 {
			b.WriteString("null")
			continue
		}
		b.WriteString("[")
		for j := 0; j < m; j++ {
			if j > 0 {
				b.WriteString([]string{",", ", ", " , "}[rng.Intn(3)])
			}
			b.WriteString(num())
		}
		b.WriteString("]")
	}
	if rng.Intn(3) == 0 {
		// Whitespace before the ring's ']', up to past the next window: the
		// walk looks for a point's separator across windows of whitespace.
		b.WriteString("\n" + strings.Repeat(" ", rng.Intn(150)))
	}
	b.WriteString("]")
	return b.String()
}

func checkPoints2(t *testing.T, data []byte, i int) {
	t.Helper()
	type res struct {
		v   [][2]float64
		end int
		err error
	}
	run := func(kernel bool) res {
		defer func(v bool) { useFloatRunVBMI = v }(useFloatRunVBMI)
		useFloatRunVBMI = kernel && floatRunVBMIHost
		var v [][2]float64
		end, err := DecodeFloat64Points(&v, data, i)
		return res{v, end, err}
	}
	a, b := run(false), run(true)
	same := a.end == b.end && errors.Is(a.err, b.err) && errors.Is(b.err, a.err) && len(a.v) == len(b.v) && (a.v == nil) == (b.v == nil)
	for k := 0; same && k < len(a.v); k++ {
		for j := 0; j < 2; j++ {
			same = same && math.Float64bits(a.v[k][j]) == math.Float64bits(b.v[k][j])
		}
	}
	if !same {
		t.Fatalf("points walk differs on %q at %d:\n off: end=%d err=%v %v\n on:  end=%d err=%v %v", data, i, a.end, a.err, a.v, b.end, b.err, b.v)
	}
}

func checkPoints3(t *testing.T, data []byte, i int) {
	t.Helper()
	type res struct {
		v   [][3]float64
		end int
		err error
	}
	run := func(kernel bool) res {
		defer func(v bool) { useFloatRunVBMI = v }(useFloatRunVBMI)
		useFloatRunVBMI = kernel && floatRunVBMIHost
		var v [][3]float64
		end, err := DecodeFloat64Points(&v, data, i)
		return res{v, end, err}
	}
	a, b := run(false), run(true)
	same := a.end == b.end && errors.Is(a.err, b.err) && errors.Is(b.err, a.err) && len(a.v) == len(b.v) && (a.v == nil) == (b.v == nil)
	for k := 0; same && k < len(a.v); k++ {
		for j := 0; j < 3; j++ {
			same = same && math.Float64bits(a.v[k][j]) == math.Float64bits(b.v[k][j])
		}
	}
	if !same {
		t.Fatalf("points walk differs on %q at %d:\n off: end=%d err=%v %v\n on:  end=%d err=%v %v", data, i, a.end, a.err, a.v, b.end, b.err, b.v)
	}
}

// TestFloat64PointsMatchesPerPoint holds DecodeFloat64Points with the points
// walk to the same reader without it — which is the generated per-point
// decoder, element for element — over generated rings of two- and
// three-number points: every separator and whitespace shape, numbers either
// side of the kernel's line, nulls, points of the wrong length, trailing
// commas, truncation, and rings followed by more of the document.
func TestFloat64PointsMatchesPerPoint(t *testing.T) {
	rng := rand.New(rand.NewSource(21))
	seps := []string{",", ", ", ",\n        ", " , ", ",\t", ",\n" + strings.Repeat(" ", 70), strings.Repeat(" ", 90) + ",", "\n" + strings.Repeat(" ", 130) + ", "}
	tails := []string{"", strings.Repeat(" ", 100), `, "type": "Polygon"}` + strings.Repeat(" ", 100)}
	for it := 0; it < 1500; it++ {
		count := rng.Intn(40)
		if rng.Intn(5) == 0 {
			count = 40 + rng.Intn(200)
		}
		odd := rng.Intn(3) == 0
		sep := seps[rng.Intn(len(seps))]
		for _, n := range []int{2, 3} {
			ring := genRing(rng, n, count, sep, odd)
			switch rng.Intn(8) {
			case 0:
				ring = ring[:rng.Intn(len(ring)+1)] // truncated
			case 1:
				ring = strings.TrimSuffix(ring, "]") + ",]" // trailing comma
			}
			lead := strings.Repeat(" ", rng.Intn(3))
			data := []byte(lead + ring + tails[rng.Intn(len(tails))])
			if n == 2 {
				checkPoints2(t, data, len(lead))
			} else {
				checkPoints3(t, data, len(lead))
			}
		}
	}
	for _, in := range []string{"null", "[]", "[ ]", "[[1,2]]", "[[1,2],]", "[[1,2] [3,4]]", "[[1,2],[3,4]", "[1,2]", "[[]]", "[null]", "{}", ""} {
		data := []byte(in + strings.Repeat(" ", 100))
		checkPoints2(t, data, 0)
		checkPoints3(t, data, 0)
	}
}

// TestFloat64PointsMatchesStdlib decodes rings of regular points with the
// walk on and compares with encoding/json — the premise the per-point
// differential rests on — and then walks each ring with the kernel alone
// (on a host with the VBMI body), resuming after each point it hands back:
// every value it writes must be encoding/json's, and every point it hands back
// must hold a number it declines by design (kernelDeclines). A walk that handed
// every point back would pass the differential.
func TestFloat64PointsMatchesStdlib(t *testing.T) {
	rng := rand.New(rand.NewSource(22))
	var took, total int
	for it := 0; it < 300; it++ {
		ring := genRing(rng, 2, 1+rng.Intn(300), []string{",", ", "}[rng.Intn(2)], false)
		data := []byte(ring + strings.Repeat(" ", 160))
		var want [][2]float64
		if err := json.Unmarshal([]byte(ring), &want); err != nil {
			t.Fatal(err)
		}
		var got [][2]float64
		end, err := DecodeFloat64Points(&got, data, 0)
		if err != nil || end != len(ring) || len(got) != len(want) {
			t.Fatalf("%q: end=%d err=%v len %d, want %d %d", ring, end, err, len(got), len(ring), len(want))
		}
		for k := range want {
			for j := 0; j < 2; j++ {
				if math.Float64bits(got[k][j]) != math.Float64bits(want[k][j]) {
					t.Fatalf("%q: point %d = %v, want %v", ring, k, got[k], want[k])
				}
			}
		}
		if !floatRunVBMIHost {
			continue
		}
		flat := make([]float64, 2*len(want))
		k := 0
		for pos := 1; ; {
			np, p, closed := parseFloatPoints(data, pos, flat[2*k:], 2)
			for e := k; e < k+np; e++ {
				if math.Float64bits(flat[2*e]) != math.Float64bits(want[e][0]) || math.Float64bits(flat[2*e+1]) != math.Float64bits(want[e][1]) {
					t.Fatalf("%q: the walk wrote point %d as %v, want %v", ring, e, flat[2*e:2*e+2], want[e])
				}
			}
			k += np
			took += np
			if closed != 0 {
				if k != len(want) || p != len(ring)-1 {
					t.Fatalf("%q: the walk closed at %d after %d points, want %d after %d", ring, p, k, len(ring)-1, len(want))
				}
				break
			}
			q := SkipWS(data, p)
			e, err := SkipValue(data, q)
			if err != nil || data[q] != '[' {
				t.Fatalf("%q: the walk handed back %d, not at a point", ring, p)
			}
			declined := false
			for _, num := range strings.Split(string(data[q+1:e-1]), ",") {
				declined = declined || kernelDeclines(strings.TrimSpace(num))
			}
			if !declined {
				t.Fatalf("%q: the walk handed back point %d, %s, which it converts", ring, k, data[q:e])
			}
			k++ // the reader decodes it
			if j := SkipWS(data, e); data[j] == ']' {
				break
			} else {
				pos = j + 1
			}
		}
		total += len(want)
	}
	if floatRunVBMIHost && took < total*99/100 {
		t.Fatalf("the walk took %d of %d points; hand-backs should be rare", took, total)
	}
}

// pointRing builds a ring of count points in one of the corpus's shapes:
// canada's compact 17-digit pairs, canada_geometry's shortest forms, and
// large-json's pretty-printed triples (a point a line, indented, with a
// trailing 0).
func pointRing(shape string, count int) []byte {
	rng := rand.New(rand.NewSource(5))
	var b strings.Builder
	switch shape {
	case "canada", "geometry":
		b.WriteString("[")
		for k := 0; k < count; k++ {
			if k > 0 {
				b.WriteString(",")
			}
			x, y := -65-rng.Float64(), 43+rng.Float64()
			if shape == "canada" {
				fmt.Fprintf(&b, "[%s,%s]", strconv.FormatFloat(x, 'f', 15, 64), strconv.FormatFloat(y, 'f', 15, 64))
			} else {
				fmt.Fprintf(&b, "[%s,%s]", strconv.FormatFloat(x, 'g', -1, 64), strconv.FormatFloat(y, 'g', -1, 64))
			}
		}
		b.WriteString("]")
	case "citylots":
		b.WriteString("[\n")
		for k := 0; k < count; k++ {
			if k > 0 {
				b.WriteString(",\n")
			}
			fmt.Fprintf(&b, "            [%s, %s, 0]", strconv.FormatFloat(-122.4-rng.Float64()/10, 'g', -1, 64), strconv.FormatFloat(37.8+rng.Float64()/10, 'g', -1, 64))
		}
		b.WriteString("\n          ]")
	}
	return []byte(b.String() + strings.Repeat(" ", 100))
}

// BenchmarkFloat64Points decodes rings of the corpus's point shapes through
// DecodeFloat64Points into a reused slice, reporting the cost per point;
// ringN/ is a ring of N points, so a short ring measures what each call and
// each ring end cost.
func BenchmarkFloat64Points(b *testing.B) {
	for _, shape := range []string{"canada", "geometry", "citylots"} {
		for _, count := range []int{6, 1000} {
			data := pointRing(shape, count)
			b.Run(fmt.Sprintf("%s/ring%d", shape, count), func(b *testing.B) {
				var v2 [][2]float64
				var v3 [][3]float64
				for i := 0; i < b.N; i++ {
					var err error
					if shape == "citylots" {
						_, err = DecodeFloat64Points(&v3, data, 0)
					} else {
						_, err = DecodeFloat64Points(&v2, data, 0)
					}
					if err != nil {
						b.Fatal(err)
					}
				}
				b.ReportMetric(float64(b.Elapsed().Nanoseconds())/float64(b.N*count), "ns/point")
			})
		}
	}
}
