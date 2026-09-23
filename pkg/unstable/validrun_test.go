package unstable

import (
	"errors"
	"math/rand"
	"strings"
	"testing"
)

// validRunHost and validRun512Host are the host's useValidRun and
// useValidRun512, which the tests flip to run the scalar walk and both bodies
// — never turning on one the CPU lacks.
var validRunHost, validRun512Host = useValidRun, useValidRun512

// validRunBodies lists the bodies this host can run.
func validRunBodies() []bool {
	if validRun512Host {
		return []bool{false, true}
	}
	return []bool{false}
}

// validNumbers draws array elements from both sides of validNumberRun's line:
// plain decimals it passes over, and everything it must hand back to the
// scalar walk — which then decides, accepting some (an exponent, a '+', a
// bare or leading '.', a number too long for its window) and rejecting others
// (a double sign, a lone '-', an overflow, a stray byte).
var validNumbers = []string{
	"0", "-0", "7", "-12", "1234567", "0.5", "-0.25", "3.14159265358979", "-65.613616999999977",
	"43.420273000000009", "00012", "-007.50", "123456789012345678901234567890",
	"1.000000000000000000000000000000000000000000001",
	"1e5", "-2.5E-3", "6.02e+23", "1e309", "-1e-400", "+1", ".5", "5.", "-.5", "--5", "-", "1.2.3",
	"1-2", "0x10", "Infinity", "NaN", "null", "true", "\"7\"", "[1]", "{}", "[]",
	strings.Repeat("9", 70), strings.Repeat("1", 400), "1" + strings.Repeat("0", 320),
}

func genValidArray(rng *rand.Rand, count int, bad bool) string {
	var b strings.Builder
	b.WriteString([]string{"[", "[ ", "[\n  ", "[\x01"}[rng.Intn(4)])
	for k := 0; k < count; k++ {
		if k > 0 {
			if bad && rng.Intn(30) == 0 {
				b.WriteString([]string{",,", " ", "", ";"}[rng.Intn(4)])
			} else {
				b.WriteString([]string{",", ", ", " , ", ",\n    ", "\t,\t", ",\x00"}[rng.Intn(6)])
			}
		}
		if bad && rng.Intn(8) == 0 {
			b.WriteString(validNumbers[rng.Intn(len(validNumbers))])
			continue
		}
		switch r := rng.Intn(6); {
		case r < 2:
			b.WriteString(validNumbers[rng.Intn(10)])
		case r < 4:
			b.WriteString(strings.Repeat("-", rng.Intn(2)) + strings.Repeat("9", 1+rng.Intn(20)) + "." + strings.Repeat("3", 1+rng.Intn(20)))
		default:
			b.WriteString(strings.Repeat("-", rng.Intn(2)) + strings.Repeat("1", 1+rng.Intn(18)))
		}
	}
	if bad && rng.Intn(10) == 0 {
		b.WriteString(",")
	}
	b.WriteString([]string{"]", " ]", "\n]", "]"}[rng.Intn(4)])
	return b.String()
}

// TestValidNumberRunMatchesScalar holds SkipValueStrict with the numeric-array
// walk to itself without it, over generated documents: arrays of every number
// shape and separator, well-formed and not, at the top level, nested in
// objects and in arrays, followed by more document, and truncated at every
// length for a share of them. The offset and the error must agree exactly.
func TestValidNumberRunMatchesScalar(t *testing.T) {
	if !validRunHost {
		t.Skip("no numeric-array validation kernel on this machine")
	}
	defer func() { useValidRun, useValidRun512 = validRunHost, validRun512Host }()
	check := func(doc string) {
		data := []byte(doc)
		i := SkipWS(data, 0)
		useValidRun, useValidRun512 = false, false // the ring walk is gated on the 512 flag alone
		we, werr := SkipValueStrict(data, i)
		useValidRun = true
		for _, b512 := range validRunBodies() {
			useValidRun512 = b512
			ge, gerr := SkipValueStrict(data, i)
			if ge != we || !errors.Is(gerr, werr) || !errors.Is(werr, gerr) {
				t.Fatalf("kernel (512=%v) (%d, %v), scalar (%d, %v) on %q", b512, ge, gerr, we, werr, doc)
			}
		}
	}
	rng := rand.New(rand.NewSource(8))
	for it := 0; it < 20000; it++ {
		bad := rng.Intn(3) == 0
		count := rng.Intn(12)
		if rng.Intn(5) == 0 {
			count = 12 + rng.Intn(300)
		}
		arr := genValidArray(rng, count, bad)
		var doc string
		switch rng.Intn(4) {
		case 0:
			doc = arr
		case 1:
			doc = `{"a":` + arr + `,"b":[` + genValidArray(rng, rng.Intn(5), bad) + `, 3]}`
		case 2:
			doc = "[" + arr + "," + genValidArray(rng, rng.Intn(40), bad) + "]"
		default:
			doc = `{"k": [` + arr + `], "t": true}`
		}
		doc += []string{"", " ", strings.Repeat(" ", 70), "x"}[rng.Intn(4)]
		check(doc)
		if rng.Intn(20) == 0 {
			for k := 0; k < len(doc) && k < 300; k++ {
				check(doc[:k])
			}
		}
	}
	// Coordinate rings, the shape validPointsRun512 walks: compact and
	// pretty-printed, regular and odd (nulls, wrong lengths, exponents,
	// strays), nested in a polygon and followed by more document, and a ring
	// whose points would sit one level past MaxDepth.
	for it := 0; it < 4000; it++ {
		odd := rng.Intn(3) == 0
		sep := []string{",", ", ", ",\n        ", " , ", ",\n" + strings.Repeat(" ", 70)}[rng.Intn(5)]
		ring := genRing(rng, 2+rng.Intn(2), rng.Intn(60), sep, odd)
		doc := []string{ring, `{"coordinates":[` + ring + `, ` + ring + `]}`, "[" + ring + "]"}[rng.Intn(3)]
		doc += []string{"", strings.Repeat(" ", 70), ",x"}[rng.Intn(3)]
		check(doc)
		if rng.Intn(10) == 0 {
			for k := 0; k < len(doc) && k < 400; k++ {
				check(doc[:k])
			}
		}
	}
	for _, shape := range []string{"canada", "geometry", "citylots"} {
		check(string(pointRing(shape, 300)))
	}
	deep := strings.Repeat("[", MaxDepth-1) + "[[1,2],[3,4]]" + strings.Repeat("]", MaxDepth-1)
	check(deep + strings.Repeat(" ", 64))
	check(strings.Repeat("[", MaxDepth-2) + "[[1,2],[3,4]]" + strings.Repeat("]", MaxDepth-2) + strings.Repeat(" ", 64))

	// Every element shape alone, and next to a long decimal run, at every
	// offset within a window.
	for _, num := range validNumbers {
		for pad := 0; pad < 70; pad++ {
			check("[" + strings.Repeat(" ", pad) + num + "]" + strings.Repeat(" ", 64))
			check("[" + strings.Repeat("1,", pad) + num + ",2]" + strings.Repeat(" ", 64))
		}
	}
}

// BenchmarkValidNumbers is SkipValueStrict over the number shapes the corpus
// validates: a flat array of decimals (numbers), coordinate rings of short and
// long points (canada, geometry), and short arrays inside objects (marine_ik).
func BenchmarkValidNumbers(b *testing.B) {
	for _, sh := range []struct{ name, doc string }{
		{"flat", "[" + strings.Repeat("0.0636837780476,", 3999) + "0.0636837780476]"},
		{"canada", string(pointRing("canada", 1000))},
		{"geometry", string(pointRing("geometry", 1000))},
		{"objects", "[" + strings.Repeat(`{"pos":[1.5,-2.25,0.75],"rot":[0,0,0,1]},`, 499) + `{"pos":[1,2,3],"rot":[0,0,0,1]}]`},
	} {
		data := []byte(sh.doc)
		b.Run(sh.name, func(b *testing.B) {
			b.SetBytes(int64(len(data)))
			for i := 0; i < b.N; i++ {
				if _, err := SkipValueStrict(data, 0); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// TestValidNumberRunTakesArrays is the differential's premise: on arrays of
// plain decimals with a window's worth of input after them, the walk itself
// must pass over every element and close the array — a walk that handed every
// element back would pass TestValidNumberRunMatchesScalar untouched.
func TestValidNumberRunTakesArrays(t *testing.T) {
	if !validRunHost {
		t.Skip("no numeric-array validation kernel on this machine")
	}
	defer func() { useValidRun512 = validRun512Host }()
	rng := rand.New(rand.NewSource(9))
	for it := 0; it < 3000; it++ {
		arr := genValidArray(rng, 1+rng.Intn(200), false)
		data := []byte(arr + strings.Repeat(" ", 64))
		i := SkipWS(data, 1)
		if data[i] != '-' && data[i]-'0' > 9 {
			continue // opens with a control byte the generator uses as whitespace
		}
		for _, b512 := range validRunBodies() {
			useValidRun512 = b512
			p, closed := validNumberRun(data, i)
			if closed != 1 || p != len(arr)-1 {
				t.Fatalf("512=%v %q: p=%d closed=%d, want the ']' at %d", b512, arr, p, closed, len(arr)-1)
			}
		}
	}
}

// TestValidPointsRunTakesRings is the ring walk's premise, as
// TestValidNumberRunTakesArrays is the flat walk's: on rings of regular points
// with a window's worth of input after them it must take every point that fits
// a window from its '[' — resuming after any that does not, which is the only
// point it may hand back — and reach the ring's ']' itself.
func TestValidPointsRunTakesRings(t *testing.T) {
	if !validRun512Host {
		t.Skip("no AVX-512 validation walk on this machine")
	}
	rng := rand.New(rand.NewSource(10))
	rings := []string{string(pointRing("canada", 500)), string(pointRing("geometry", 500)), string(pointRing("citylots", 500))}
	for it := 0; it < 2000; it++ {
		sep := []string{",", ", ", ",\n        ", ",\n" + strings.Repeat(" ", 70)}[rng.Intn(4)]
		rings = append(rings, genRing(rng, 2+rng.Intn(2), 1+rng.Intn(80), sep, false))
	}
	for _, r := range rings {
		r = strings.TrimRight(r, " ")
		data := []byte(r + strings.Repeat(" ", 64))
		for i := SkipWS(data, 1); ; {
			p, closed := validPointsRun512(data, i)
			if closed == 1 {
				if p != len(r)-1 {
					t.Fatalf("%.100q: closed at %d, want the ']' at %d", r, p, len(r)-1)
				}
				break
			}
			q := SkipWS(data, p)
			e, err := SkipValue(data, q)
			if err != nil || data[q] != '[' || e-q <= 64 {
				t.Fatalf("%.100q: handed back %q at %d, which fits a window", r, data[q:min(e, len(data))], p)
			}
			j := SkipWS(data, e)
			if data[j] == ']' {
				break
			}
			i = SkipWS(data, j+1)
		}
	}
}
