package json

import (
	"errors"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

// The integer parsers are held to strconv over a hand-picked corpus and a
// generated one: every accepted token must parse to the same value, and every
// token strconv refuses must be refused here — with one documented
// difference, the leading '+' that ParseUint accepts and strconv.ParseUint
// does not (ParseFloat and ParseInt both take it, so the family agrees).
func TestParseIntMatchesStrconv(t *testing.T) {
	corpus := []string{
		"0", "1", "-1", "+1", "42", "-42", "007", "-007", "+007", "-0", "+0",
		"9223372036854775807", "-9223372036854775808", // the int64 edges
		"9223372036854775808", "-9223372036854775809", // one past them
		"18446744073709551615", "18446744073709551616", // the uint64 edge and past it
		"99999999999999999999", "184467440737095516150", // 20 and 21 digits
		"", "-", "+", "--1", "+-1", "1.0", "1e3", "1.", ".5", "0x10", "1_000",
		" 1", "1 ", "\t1", "1\n", "null", "true", "\"1\"", "1,", "١", "12a", "a12",
		"00000000000000000000000000001", "-00000000000000000000000000001",
	}
	rnd := rand.New(rand.NewSource(7))
	for range 20000 {
		var sb strings.Builder
		switch rnd.Intn(6) {
		case 0:
			sb.WriteByte('-')
		case 1:
			sb.WriteByte('+')
		}
		for range rnd.Intn(3) {
			sb.WriteByte('0')
		}
		for range 1 + rnd.Intn(21) {
			sb.WriteByte(byte('0' + rnd.Intn(10)))
		}
		switch rnd.Intn(12) {
		case 0:
			sb.WriteString(".5")
		case 1:
			sb.WriteString("e2")
		case 2:
			sb.WriteByte(' ')
		}
		corpus = append(corpus, sb.String())
	}
	for _, s := range corpus {
		b := []byte(s)

		wantI, wantErr := strconv.ParseInt(s, 10, 64)
		gotI, err := ParseInt(b)
		switch {
		case (err == nil) != (wantErr == nil):
			t.Errorf("ParseInt(%q): err %v, strconv %v", s, err, wantErr)
		case err != nil && !errors.Is(err, ErrBadNumber):
			t.Errorf("ParseInt(%q): err %v, want ErrBadNumber", s, err)
		case err == nil && gotI != wantI:
			t.Errorf("ParseInt(%q) = %d, strconv %d", s, gotI, wantI)
		}

		// strconv.ParseUint takes no '+'; the family here does.
		us := strings.TrimPrefix(s, "+")
		if us != s && strings.HasPrefix(us, "+") {
			us = s // "++1": stays refused on both sides
		}
		wantU, wantErr := strconv.ParseUint(us, 10, 64)
		gotU, err := ParseUint(b)
		switch {
		case (err == nil) != (wantErr == nil):
			t.Errorf("ParseUint(%q): err %v, strconv(%q) %v", s, err, us, wantErr)
		case err != nil && !errors.Is(err, ErrBadNumber):
			t.Errorf("ParseUint(%q): err %v, want ErrBadNumber", s, err)
		case err == nil && gotU != wantU:
			t.Errorf("ParseUint(%q) = %d, strconv %d", s, gotU, wantU)
		}
	}
}

// The edges, spelled out: the two int64 limits parse exactly, one past each
// is refused, and the uint64 limit parses where the same digits overflow
// int64 — the cases a wrap in the accumulator would get wrong silently.
func TestParseIntEdges(t *testing.T) {
	if v, err := ParseInt([]byte("-9223372036854775808")); err != nil || v != math.MinInt64 {
		t.Fatalf("MinInt64: %d, %v", v, err)
	}
	if v, err := ParseInt([]byte("9223372036854775807")); err != nil || v != math.MaxInt64 {
		t.Fatalf("MaxInt64: %d, %v", v, err)
	}
	if _, err := ParseInt([]byte("9223372036854775808")); !errors.Is(err, ErrBadNumber) {
		t.Fatalf("MaxInt64+1 must be ErrBadNumber, got %v", err)
	}
	if _, err := ParseInt([]byte("-9223372036854775809")); !errors.Is(err, ErrBadNumber) {
		t.Fatalf("MinInt64-1 must be ErrBadNumber, got %v", err)
	}
	if v, err := ParseUint([]byte("18446744073709551615")); err != nil || v != math.MaxUint64 {
		t.Fatalf("MaxUint64: %d, %v", v, err)
	}
	if _, err := ParseUint([]byte("18446744073709551616")); !errors.Is(err, ErrBadNumber) {
		t.Fatalf("MaxUint64+1 must be ErrBadNumber, got %v", err)
	}
	if _, err := ParseUint([]byte("-1")); !errors.Is(err, ErrBadNumber) {
		t.Fatalf("a negative uint must be ErrBadNumber, got %v", err)
	}
}

// A token ParseFloat accepts as a whole number is not thereby an integer:
// ParseInt refuses a fraction and an exponent rather than truncating.
func TestParseIntIsNotParseFloat(t *testing.T) {
	for _, s := range []string{"1.0", "1e3", "10.", "1E0"} {
		if f, err := ParseFloat([]byte(s)); err != nil || f != math.Trunc(f) {
			t.Fatalf("ParseFloat(%q) = %v, %v: the fixture assumes a whole float", s, f, err)
		}
		if _, err := ParseInt([]byte(s)); !errors.Is(err, ErrBadNumber) {
			t.Errorf("ParseInt(%q) must refuse a float spelling, got %v", s, err)
		}
	}
}

func BenchmarkParseInt(b *testing.B) {
	in := []byte("1788087600123")
	b.ReportAllocs()
	for range b.N {
		if _, err := ParseInt(in); err != nil {
			b.Fatal(err)
		}
	}
}
