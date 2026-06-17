package support

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"testing"
)

// TestParseFloatMatchesStrconv hammers the full float path (ParseFloat, which
// runs scanFloat's Clinger fast path then eiselLemire64 before strconv)
// against strconv.ParseFloat over many random and structured numbers, requiring
// bit-for-bit equality. This guards the Eisel-Lemire fast path: whenever it
// returns a value it must equal strconv's to the bit.
func TestParseFloatMatchesStrconv(t *testing.T) {
	check := func(s string) {
		want, perr := strconv.ParseFloat(s, 64)
		got, gerr := ParseFloat([]byte(s))
		if perr != nil {
			return // out of float64 range; strconv errors, EL declines
		}
		if gerr != nil {
			t.Fatalf("ParseFloat(%q) err=%v, strconv ok=%v", s, gerr, want)
		}
		if math.Float64bits(got) != math.Float64bits(want) {
			t.Fatalf("ParseFloat(%q)=%x want=%x (%v vs %v)", s, math.Float64bits(got), math.Float64bits(want), got, want)
		}
	}
	n1, n2, n3 := 700_000, 200_000, 100_000
	if testing.Short() {
		n1, n2, n3 = 100_000, 40_000, 20_000
	}
	r := rand.New(rand.NewSource(99))
	for k := 0; k < n1; k++ {
		man := r.Uint64() % 10_000_000_000_000_000_000
		exp := r.Intn(760) - 380
		sign := ""
		if r.Intn(2) == 0 {
			sign = "-"
		}
		check(fmt.Sprintf("%s%de%d", sign, man, exp))
	}
	for k := 0; k < n2; k++ {
		man := r.Uint64() % 100000000000
		frac := r.Uint64() % 1000000000
		exp := r.Intn(80) - 40
		check(fmt.Sprintf("%d.%de%d", man, frac, exp))
	}
	// 20+ digit mantissas must defer to strconv and still be correct.
	for k := 0; k < n3; k++ {
		check(fmt.Sprintf("%d%de%d", r.Uint64(), r.Uint64()%1000, r.Intn(40)-20))
	}
	// Structured edge cases: powers of ten, halfway values, the coordinate shape.
	for _, s := range []string{
		"0", "-0", "1", "1e23", "1e-23", "5e-324", "1e308", "1.7976931348623157e308",
		"-122.42200352825247", "37.808480096967251", "123456789012345678", "9007199254740993",
	} {
		check(s)
	}
}
