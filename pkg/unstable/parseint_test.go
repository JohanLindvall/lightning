package unstable

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

// parseIntCorpus is every shape the length classification can take — each
// boundary of the four folds (3/4, 8/9, 16/17, 19/20), both integer limits and
// one past them, leading zeros long enough to matter, and every spelling that
// is not an integer token.
func parseIntCorpus() []string {
	c := []string{
		"", "0", "1", "9", "10", "99", "100", "999", "1000", "9999", "10000",
		"99999", "123456", "1234567", "12345678", "123456789", "1234567890",
		"1788087600", "1788087600123", "1234567890123456", "12345678901234567",
		"123456789012345678", "1234567890123456789", "12345678901234567890",
		"9223372036854775807", "9223372036854775808", "9223372036854775809",
		"18446744073709551615", "18446744073709551616", "99999999999999999999",
		"999999999999999999999", "0000000000000000000000000000000000000001",
		"00000000000000000000", "000000000000000000000000", "007", "0123456789",
		"18446744073709551610", "10000000000000000000", "1844674407370955161",
		"+1", "+0", "+9223372036854775807", "-1", "-0", "-9223372036854775808",
		"-9223372036854775809", "-18446744073709551615", "-0000000000000000000001",
		"1.0", "1e3", "1E3", "1.5", ".5", "5.", " 1", "1 ", "1\n", "\t1",
		"null", `"1"`, "0x10", "1_000", "१२३", "abc", "1a", "a1", "+", "-", "++1",
		"--1", "+-1", "-+1", "1-", "1+", "\x00", "1\x00", "\xff",
	}
	// Every length from 1 to 24 digits, so no fold boundary goes unvisited.
	for n := 1; n <= 24; n++ {
		s := ""
		for i := 0; i < n; i++ {
			s += string(rune('1' + (i+n)%9))
		}
		c = append(c, s, "-"+s, "+"+s, "0"+s)
		// the same length with a non-digit at each position
		for p := 0; p < n; p++ {
			c = append(c, s[:p]+"x"+s[p+1:])
		}
	}
	return c
}

// wantInt is what strconv makes of s under this package's grammar: strconv's
// own, except that a leading '+' is accepted before the digits (ParseUint's one
// documented difference from strconv.ParseUint) and no underscores.
func wantUint(s string) (uint64, bool) {
	t := s
	if len(t) > 0 && t[0] == '+' {
		t = t[1:]
	}
	if t == "" {
		return 0, false
	}
	for i := 0; i < len(t); i++ {
		if t[i] < '0' || t[i] > '9' {
			return 0, false
		}
	}
	v, err := strconv.ParseUint(t, 10, 64)
	return v, err == nil
}

func wantInt(s string) (int64, bool) {
	t, neg := s, false
	if len(t) > 0 && (t[0] == '-' || t[0] == '+') {
		neg = t[0] == '-'
		t = t[1:]
	}
	if t == "" {
		return 0, false
	}
	for i := 0; i < len(t); i++ {
		if t[i] < '0' || t[i] > '9' {
			return 0, false
		}
	}
	if neg {
		t = "-" + t
	}
	v, err := strconv.ParseInt(t, 10, 64)
	return v, err == nil
}

func TestParseIntMatchesStrconvUnstable(t *testing.T) {
	for _, s := range parseIntCorpus() {
		// make gives a slice whose capacity is exactly its length, so a load
		// past the token's end would run past what the bounds checks and the
		// race detector can see as owned.
		tight := make([]byte, len(s))
		copy(tight, s)
		wu, wuok := wantUint(s)
		gu, err := ParseUint(tight)
		if (err == nil) != wuok || (wuok && gu != wu) {
			t.Errorf("ParseUint(%q) = %d, %v; want %d, ok=%v", s, gu, err, wu, wuok)
		}
		wi, wiok := wantInt(s)
		gi, err := ParseInt(tight)
		if (err == nil) != wiok || (wiok && gi != wi) {
			t.Errorf("ParseInt(%q) = %d, %v; want %d, ok=%v", s, gi, err, wi, wiok)
		}
	}
}

// TestParseIntLimits spells the boundaries out one by one: an accumulator that
// wraps gets exactly these wrong and nothing else.
func TestParseIntLimits(t *testing.T) {
	cases := []struct {
		in   string
		want int64
		ok   bool
	}{
		{"9223372036854775807", math.MaxInt64, true},
		{"9223372036854775808", 0, false},
		{"-9223372036854775808", math.MinInt64, true},
		{"-9223372036854775809", 0, false},
		{"-9223372036854775807", math.MinInt64 + 1, true},
		{"00000000009223372036854775807", math.MaxInt64, true},
	}
	for _, c := range cases {
		got, err := ParseInt([]byte(c.in))
		if (err == nil) != c.ok || (c.ok && got != c.want) {
			t.Errorf("ParseInt(%q) = %d, %v; want %d, ok=%v", c.in, got, err, c.want, c.ok)
		}
	}
	ucases := []struct {
		in   string
		want uint64
		ok   bool
	}{
		{"18446744073709551615", math.MaxUint64, true},
		{"18446744073709551616", 0, false},
		{"18446744073709551614", math.MaxUint64 - 1, true},
		{"9999999999999999999", 9999999999999999999, true},
		{"10000000000000000000", 10000000000000000000, true},
		{"20000000000000000000", 0, false},
		{"000000000000000000000018446744073709551615", math.MaxUint64, true},
	}
	for _, c := range ucases {
		got, err := ParseUint([]byte(c.in))
		if (err == nil) != c.ok || (c.ok && got != c.want) {
			t.Errorf("ParseUint(%q) = %d, %v; want %d, ok=%v", c.in, got, err, c.want, c.ok)
		}
	}
}

// TestParseIntAtBufferEnd puts the token at the very end of its slice at every
// length: the folds read four and eight bytes at a time, and a load that ran
// past the token would read past the allocation here.
func TestParseIntAtBufferEnd(t *testing.T) {
	for n := 1; n <= 21; n++ {
		s := ""
		for i := 0; i < n; i++ {
			s += string(rune('0' + (i+1)%10))
		}
		buf := make([]byte, len(s))
		copy(buf, s)
		want, ok := wantUint(s)
		got, err := ParseUint(buf)
		if (err == nil) != ok || (ok && got != want) {
			t.Fatalf("ParseUint(%q) = %d, %v; want %d, ok=%v", s, got, err, want, ok)
		}
	}
}

func FuzzParseIntMatchesStrconv(f *testing.F) {
	for _, s := range []string{"1", "1788087600123", "18446744073709551615", "-9223372036854775808", "0007"} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, s string) {
		if len(s) > 64 {
			return
		}
		b := make([]byte, len(s))
		copy(b, s)
		if w, ok := wantUint(s); true {
			g, err := ParseUint(b)
			if (err == nil) != ok || (ok && g != w) {
				t.Fatalf("ParseUint(%q) = %d, %v; want %d, ok=%v", s, g, err, w, ok)
			}
		}
		if w, ok := wantInt(s); true {
			g, err := ParseInt(b)
			if (err == nil) != ok || (ok && g != w) {
				t.Fatalf("ParseInt(%q) = %d, %v; want %d, ok=%v", s, g, err, w, ok)
			}
		}
	})
}

func BenchmarkParseIntUnstable(b *testing.B) {
	for _, n := range []int{1, 3, 5, 8, 10, 13, 16, 19, 20} {
		s := ""
		for i := 0; i < n; i++ {
			s += string(rune('1' + i%9))
		}
		in := []byte(s)
		b.Run(fmt.Sprintf("d%02d", n), func(b *testing.B) {
			var v int64
			for i := 0; i < b.N; i++ {
				v, _ = ParseInt(in)
			}
			_ = v
		})
	}
}
