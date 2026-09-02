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

// scanFloatShapes are twenty-token samples of the corpus's float shapes, each
// with the trailing padding a real token has behind it so the fast path's
// wide loads apply: "slow" is float-array-slow's 17-significant-digit
// mantissa with an exponent (the shape whose value chain runs through all
// three fraction words and Eisel-Lemire), "canada" its 2-3 integer and 14-15
// fraction digits, "mesh" mesh_pretty's Python-repr fractions with a zero
// integer part, and "array" float-array's short Clinger-path numbers.
var scanFloatShapes = []struct{ name, toks string }{
	{"slow", "-4.8946386314775124e-112,-7.2392693015831473e-119,6.1146678976296868e+241,3.1350111345149310e+69,-5.6410683628159239e+264,-8.3601037640724002e-31,-3.5778426266913286e-80,9.3310978215777177e+222,-9.5892411796651425e-50,-8.1945889983848621e+125,-8.4686167425308440e+51,3.1656543701577933e+209,6.5630341237368151e-36,-7.4147654609763245e+228,3.5866491738617601e+218,-7.9287160265241082e+37,8.9235940179578557e+216,2.7204705685205802e+70,-4.3239092647361114e+113,-8.7407225558064988e+31"},
	{"canada", "-65.613616999999977,43.420273000000009,-65.619720000000029,43.418052999999986,-65.625,43.421379000000059,-65.636123999999882,43.449714999999969,-65.633515999999983,43.451641999999984,-65.638710999999938,43.453796999999953,-65.643559999999923,43.458823000000056,-65.644542999999978,43.463421000000032,-65.641649999999984,43.466385000000005,-65.635467000000013,43.469502000000011"},
	{"mesh", "0.0636837780476,2.34647130966,0.0452156066895,0.0606756210327,2.34551143646,0.0295177698135,0.0577713251114,2.35072994232,0.0302016735077,0.0611855685711,2.35236978531,0.0464453697205,0.0572466850281,2.34213256836,0.0151053667068,0.0538873374462,2.34651136398,0.0153127908707,0.0500386953354,2.35005044937"},
	{"array", "7557.557,-1.77897877,-464917.9940213,255592.74,876703.034,-0.7246173,-0.811223,-0.788,33.35,-4254.63849591,-8174.61666,-0.527,0.989514,-97.4493,-987747.2973981,412901.86656922,-0.00039,-82.7888,952582.53361895,-0.22369"},
}

// BenchmarkScanFloatShapes measures scanFloat's fast path per twenty tokens of
// each shape, and BenchmarkScanFloatSlowShapes the loop form on the same
// tokens; the pair is what sized the 2026-09 Zen 4 fast-path work (see
// CLAUDE.md), where the fast path had been slower than the loop on "slow".
func BenchmarkScanFloatShapes(b *testing.B) {
	for _, sh := range scanFloatShapes {
		bufs := scanFloatShapeBufs(sh.toks)
		b.Run(sh.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, buf := range bufs {
					scanFloat(buf, 0)
				}
			}
		})
	}
}

func BenchmarkScanFloatSlowShapes(b *testing.B) {
	for _, sh := range scanFloatShapes {
		bufs := scanFloatShapeBufs(sh.toks)
		b.Run(sh.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, buf := range bufs {
					scanFloatSlow(buf, 0)
				}
			}
		})
	}
}

func scanFloatShapeBufs(toks string) [][]byte {
	var bufs [][]byte
	for _, tok := range strings.Split(toks, ",") {
		bufs = append(bufs, []byte(tok+","+strings.Repeat(" ", 56)))
	}
	return bufs
}
