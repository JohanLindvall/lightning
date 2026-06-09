package json

import (
	"math"
	"strconv"
	"testing"
)

func TestParseFloat(t *testing.T) {
	good := []string{
		"0", "1", "-1", "3.14", "-3.14", "100", "1e3", "1.5e-3", "-2.25E2",
		"0.0001", "12345678901234567890", // many digits -> strconv fallback
		"1e308", "1e-300", "123456789012345.6789",
	}
	for _, s := range good {
		got, err := ParseFloat([]byte(s))
		if err != nil {
			t.Errorf("ParseFloat(%q) error: %v", s, err)
			continue
		}
		want, _ := strconv.ParseFloat(s, 64)
		if got != want {
			t.Errorf("ParseFloat(%q) = %v, want %v", s, got, want)
		}
	}
	for _, s := range []string{"", "abc", "1.2.3", "12x", " 1", "1 ", "0x10"} {
		if f, err := ParseFloat([]byte(s)); err == nil {
			t.Errorf("ParseFloat(%q) = %v, want error", s, f)
		}
	}
	if f, _ := ParseFloat([]byte("1e3")); f != 1000 || math.IsNaN(f) {
		t.Errorf("fast path wrong: %v", f)
	}
}
