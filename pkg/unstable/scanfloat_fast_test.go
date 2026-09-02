package unstable

import (
	"math"
	"math/rand"
	"strings"
	"testing"
)

// TestScanFloatFastMatchesSlow pins scanFloat's straight-line fast path to the
// loop form it fronts: every token shape, at every position in a padded buffer
// (so both the wide-load path and the near-end-of-buffer fallback are hit), must
// return the same (value bits, end, fast, ok).
func TestScanFloatFastMatchesSlow(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	digits := func(n int) string {
		var b strings.Builder
		for i := 0; i < n; i++ {
			b.WriteByte(byte('0' + rng.Intn(10)))
		}
		return b.String()
	}
	var toks []string
	// Hand-picked corners.
	toks = append(toks,
		"0", "-0", "+0", "1", "-1", "1.", "1.0", "-0.0", ".5", "-.5", "1e", "1e5", "1E-5", "1.5e3",
		"1.2.3", "1e2e3", "1-2", "01", "007.5", "0.270354", "-65.613616999999977", "43.420273000000009",
		"0.0006988752666567719", "12345678", "123456789", "1234567.8", "0.12345678", "0.123456789",
		"0.1234567890123456", "0.12345678901234567", "9007199254740993", "9007199254740992.5",
		"1234567890123456789", "12345678901234567890", "0.0000000000000000001", "-", "+", "x", "",
		"1x", "1.x", "0.1x", "1e+", "1e-", "-1.5E+10", "9999999.9999999999999", "1.7976931348623157e308",
		"5e-324", "1e309", "0.00000000000000000001", "123456789012.1234567", "1234567.123456789012",
	)
	// Generated: sign x int digits x fraction digits x exponent x trailing byte.
	trail := []string{"", ",", "]", "}", " ", "\n", "e", "E", ".", "x", "5", "-", "+"}
	for n := 0; n < 20000; n++ {
		var b strings.Builder
		switch rng.Intn(6) {
		case 0:
			b.WriteByte('-')
		case 1:
			b.WriteByte('+')
		}
		b.WriteString(digits(rng.Intn(11)))
		if rng.Intn(3) != 0 {
			b.WriteByte('.')
			if rng.Intn(4) == 0 {
				b.WriteString(strings.Repeat("0", rng.Intn(12)))
			}
			b.WriteString(digits([]int{0, 1, 7, 8, 9, 15, 16, 17, 23, 24, 25, rng.Intn(27)}[rng.Intn(12)]))
		}
		if rng.Intn(6) == 0 {
			b.WriteString([]string{"e", "E"}[rng.Intn(2)])
			b.WriteString([]string{"", "+", "-"}[rng.Intn(3)])
			b.WriteString(digits(rng.Intn(4)))
		}
		b.WriteString(trail[rng.Intn(len(trail))])
		toks = append(toks, b.String())
	}
	pads := []int{0, 1, 5, 12, 26, 40, 48, 60}
	for _, tok := range toks {
		for _, pad := range pads {
			for _, lead := range []string{"", "[", "[1,"} {
				buf := []byte(lead + tok + strings.Repeat("]", pad))
				i := len(lead)
				f1, e1, fast1, ok1 := scanFloat(buf, i)
				f2, e2, fast2, ok2 := scanFloatSlow(buf, i)
				if math.Float64bits(f1) != math.Float64bits(f2) || e1 != e2 || fast1 != fast2 || ok1 != ok2 {
					t.Fatalf("scanFloat(%q@%d) = (%v,%d,%v,%v); slow = (%v,%d,%v,%v)", buf, i, f1, e1, fast1, ok1, f2, e2, fast2, ok2)
				}
			}
		}
	}
}

func TestParse8Digits(t *testing.T) {
	rng := rand.New(rand.NewSource(2))
	for n := 0; n < 200000; n++ {
		v := uint64(rng.Intn(100000000))
		var b [8]byte
		x := v
		for j := 7; j >= 0; j-- {
			b[j] = byte('0' + x%10)
			x /= 10
		}
		w := uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 | uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
		if got := parse8Digits(w - 0x3030303030303030); got != v {
			t.Fatalf("parse8Digits(%q) = %d, want %d", b[:], got, v)
		}
	}
}
