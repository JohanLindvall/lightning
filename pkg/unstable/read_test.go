package unstable

import (
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"
)

// pad returns a string long enough (>= 32 bytes) to exercise the SIMD path of
// indexCloseOrEscape / indexStructural on amd64 and arm64.
func longUnescaped() string { return strings.Repeat("a", 40) }

// --- ReadKey ---------------------------------------------------------------

func TestReadKey(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		i       int
		want    string
		wantEnd int
		wantErr error
	}{
		{"simple", `"key":1`, 0, "key", 5, nil},
		{"empty", `"":1`, 0, "", 2, nil},
		{"with spaces inside", `"a b c":1`, 0, "a b c", 7, nil},
		{"offset", `xx"key"`, 2, "key", 7, nil},
		{"long unescaped", `"` + longUnescaped() + `":1`, 0, longUnescaped(), len(longUnescaped()) + 2, nil},
		{"escaped key", `"a\"b":1`, 0, `a"b`, 6, nil},
		{"unicode escape key", "\"\\u0041\":1", 0, "A", 8, nil},

		{"not a string", `123`, 0, "", 0, ErrInvalidJSON},
		{"index past end", `"`, 1, "", 1, ErrInvalidJSON},
		{"truncated no close", `"abc`, 0, "", 4, ErrTruncated},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, end, err := ReadKey([]byte(tt.in), tt.i)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("err = %v, want %v", err, tt.wantErr)
			}
			if err == nil && got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
			if end != tt.wantEnd {
				t.Errorf("end = %d, want %d", end, tt.wantEnd)
			}
		})
	}
}

// --- ReadStringOrNull / ReadStringNoCopyOrNull -----------------------------

func TestReadStringOrNull(t *testing.T) {
	long := longUnescaped()
	tests := []struct {
		name    string
		in      string
		want    string
		wantEnd int
		wantErr error
	}{
		{"simple", `"hello"`, "hello", 7, nil},
		{"empty", `""`, "", 2, nil},
		{"null", `null`, "", 4, nil},
		{"long unescaped (SIMD)", `"` + long + `"`, long, len(long) + 2, nil},
		{"escape quote", `"a\"b"`, `a"b`, 6, nil},
		{"escape backslash", `"a\\b"`, `a\b`, 6, nil},
		{"escape slash", `"a\/b"`, `a/b`, 6, nil},
		{"escape controls", `"\b\f\n\r\t"`, "\b\f\n\r\t", 12, nil},
		{"unicode BMP escape", "\"\\u00e9\"", "é", 8, nil},
		{"unicode BMP uppercase hex", "\"\\u00C9\"", "É", 8, nil},
		{"unicode surrogate pair", "\"\\ud83d\\ude00\"", "😀", 14, nil},

		{"truncated empty", ``, "", 0, ErrTruncated},
		{"not a string", `42`, "", 0, ErrInvalidJSON},
		{"truncated no close", `"abc`, "", 4, ErrTruncated},
		{"bad escape", `"a\xb"`, "", 3, ErrBadEscape},
		{"bad unicode", `"\u00zz"`, "", 3, ErrBadUnicode},
		{"truncated unicode", `"\u00`, "", 3, ErrTruncated},
		{"truncated after escape", `"a\`, "", 3, ErrTruncated},
		// Lone high surrogate not followed by a low surrogate: kept as the
		// replacement rune rather than combined.
		{"lone high surrogate", "\"\\ud83d\"", "�", 8, nil},
		// High surrogate followed by a malformed second \u escape: the second
		// escape's error is propagated.
		{"surrogate bad second escape", "\"\\ud83d\\uzzzz\"", "", 9, ErrBadUnicode},
		// Escaped string left unterminated after a valid escape.
		{"escaped then truncated", `"a\n`, "", 4, ErrTruncated},
		{"bad null", `nul`, "", 0, ErrInvalidJSON},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, fn := range []struct {
				label string
				f     func([]byte, int) (string, int, error)
			}{
				{"copy", ReadStringOrNull},
				{"nocopy", ReadStringNoCopyOrNull},
				{"destructive", ReadStringDestructiveOrNull},
			} {
				got, end, err := fn.f([]byte(tt.in), 0)
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("%s: err = %v, want %v", fn.label, err, tt.wantErr)
				}
				if err == nil && got != tt.want {
					t.Errorf("%s: got %q, want %q", fn.label, got, tt.want)
				}
				if end != tt.wantEnd {
					t.Errorf("%s: end = %d, want %d", fn.label, end, tt.wantEnd)
				}
			}
		})
	}
}

// --- ReadInt64OrNull -------------------------------------------------------

func TestReadInt64OrNull(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    int64
		wantEnd int
		wantErr error
	}{
		{"zero", `0`, 0, 1, nil},
		{"positive", `12345`, 12345, 5, nil},
		{"negative", `-678`, -678, 4, nil},
		{"null", `null`, 0, 4, nil},
		{"trailing", `42,`, 42, 2, nil},
		{"truncated fraction toward zero", `3.99`, 3, 4, nil},
		{"exponent tolerated", `12e3`, 12, 4, nil},
		{"negative with fraction", `-7.5`, -7, 4, nil},

		{"empty", ``, 0, 0, ErrTruncated},
		{"bare minus", `-`, 0, 1, ErrBadNumber},
		{"minus non-digit", `-x`, 0, 1, ErrBadNumber},
		{"not a number", `abc`, 0, 0, ErrBadNumber},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, end, err := ReadInt64OrNull([]byte(tt.in), 0)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("err = %v, want %v", err, tt.wantErr)
			}
			if err == nil && got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
			if end != tt.wantEnd {
				t.Errorf("end = %d, want %d", end, tt.wantEnd)
			}
		})
	}
}

// --- ReadUint64OrNull ------------------------------------------------------

func TestReadUint64OrNull(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    uint64
		wantEnd int
		wantErr error
	}{
		{"zero", `0`, 0, 1, nil},
		{"positive", `98765`, 98765, 5, nil},
		{"null", `null`, 0, 4, nil},
		{"trailing", `7]`, 7, 1, nil},
		{"fraction truncated", `9.99`, 9, 4, nil},
		{"exponent tolerated", `5E2`, 5, 3, nil},

		{"empty", ``, 0, 0, ErrTruncated},
		{"negative rejected", `-1`, 0, 0, ErrBadNumber},
		{"not a number", `x`, 0, 0, ErrBadNumber},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, end, err := ReadUint64OrNull([]byte(tt.in), 0)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("err = %v, want %v", err, tt.wantErr)
			}
			if err == nil && got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
			if end != tt.wantEnd {
				t.Errorf("end = %d, want %d", end, tt.wantEnd)
			}
		})
	}
}

// --- ReadNumberOrNull ------------------------------------------------------

func TestReadNumberOrNull(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    string
		wantEnd int
		wantErr error
	}{
		{"int", `12345`, "12345", 5, nil},
		{"negative", `-42`, "-42", 3, nil},
		{"float", `3.14159`, "3.14159", 7, nil},
		{"scientific", `-6.022e23`, "-6.022e23", 9, nil},
		{"big uint literal", `18446744073709551616`, "18446744073709551616", 20, nil},
		{"null", `null`, "", 4, nil},
		{"trailing", `7]`, "7", 1, nil},
		{"in object", `99,`, "99", 2, nil},

		{"empty", ``, "", 0, ErrTruncated},
		{"quoted rejected", `"7.5"`, "", 0, ErrBadNumber},
		{"not a number", `x`, "", 0, ErrBadNumber},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, nocopy := range []bool{false, true} {
				var got string
				var end int
				var err error
				if nocopy {
					got, end, err = ReadNumberNoCopyOrNull([]byte(tt.in), 0)
				} else {
					got, end, err = ReadNumberOrNull([]byte(tt.in), 0)
				}
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("nocopy=%v: err = %v, want %v", nocopy, err, tt.wantErr)
				}
				if err == nil && got != tt.want {
					t.Errorf("nocopy=%v: got %q, want %q", nocopy, got, tt.want)
				}
				if end != tt.wantEnd {
					t.Errorf("nocopy=%v: end = %d, want %d", nocopy, end, tt.wantEnd)
				}
			}
		})
	}
}

// --- ReadFloat64OrNull -----------------------------------------------------

func TestReadFloat64OrNull(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    float64
		wantEnd int
		wantErr error
	}{
		{"integer", `42`, 42, 2, nil},
		{"decimal", `3.14`, 3.14, 4, nil},
		{"negative", `-2.5`, -2.5, 4, nil},
		{"positive exp", `1e3`, 1000, 3, nil},
		{"negative exp", `5e-2`, 0.05, 4, nil},
		{"explicit plus exp", `2e+2`, 200, 4, nil},
		{"null", `null`, 0, 4, nil},
		{"zero", `0`, 0, 1, nil},
		{"trailing", `1.5,`, 1.5, 3, nil},
		// slow-path fallthrough: > 19 digits forces strconv.
		{"many digits", `12345678901234567890.5`, 12345678901234567890.5, 22, nil},
		// big exponent forces slow path.
		{"big exponent", `1e100`, 1e100, 5, nil},
		{"exp out of fast range", `1e23`, 1e23, 4, nil},

		{"empty", ``, 0, 0, ErrTruncated},
		{"not a number", `abc`, 0, 0, ErrBadNumber},
		{"malformed", `1.2.3`, 0, 5, ErrBadNumber},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, end, err := ReadFloat64OrNull([]byte(tt.in), 0)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("err = %v, want %v", err, tt.wantErr)
			}
			if err == nil && got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
			if end != tt.wantEnd {
				t.Errorf("end = %d, want %d", end, tt.wantEnd)
			}
		})
	}
}

// scanFloat's fast path is exercised indirectly above, but verify it agrees with
// the slow path and correctly declines out-of-range tokens. A token is "taken"
// on the fast path when fast=true and the whole input was consumed.
// --- ReadBoolOrNull --------------------------------------------------------

func TestReadBoolOrNull(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    bool
		wantEnd int
		wantErr error
	}{
		{"true", `true`, true, 4, nil},
		{"false", `false`, false, 5, nil},
		{"null", `null`, false, 4, nil},
		{"true trailing", `true,`, true, 4, nil},

		{"empty", ``, false, 0, ErrTruncated},
		{"truncated true", `tru`, false, 0, ErrInvalidJSON},
		{"truncated false", `fals`, false, 0, ErrInvalidJSON},
		{"misspelled true", `trxe`, false, 0, ErrInvalidJSON},
		{"misspelled false", `falxe`, false, 0, ErrInvalidJSON},
		{"garbage", `xyz`, false, 0, ErrInvalidJSON},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, end, err := ReadBoolOrNull([]byte(tt.in), 0)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("err = %v, want %v", err, tt.wantErr)
			}
			if err == nil && got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
			if end != tt.wantEnd {
				t.Errorf("end = %d, want %d", end, tt.wantEnd)
			}
		})
	}
}

// --- ReadTimeOrNull --------------------------------------------------------

func TestReadTimeOrNull(t *testing.T) {
	const ts = "2023-01-02T15:04:05Z"
	want, _ := time.Parse(time.RFC3339, ts)

	t.Run("valid", func(t *testing.T) {
		got, end, err := ReadTimeOrNull([]byte(`"`+ts+`"`), 0)
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		if !got.Equal(want) {
			t.Errorf("got %v, want %v", got, want)
		}
		if end != len(ts)+2 {
			t.Errorf("end = %d, want %d", end, len(ts)+2)
		}
	})

	t.Run("with offset", func(t *testing.T) {
		const o = "2023-01-02T15:04:05+02:00"
		got, _, err := ReadTimeOrNull([]byte(`"`+o+`"`), 0)
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		exp, _ := time.Parse(time.RFC3339, o)
		if !got.Equal(exp) {
			t.Errorf("got %v, want %v", got, exp)
		}
	})

	t.Run("null", func(t *testing.T) {
		got, end, err := ReadTimeOrNull([]byte(`null`), 0)
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		if !got.IsZero() {
			t.Errorf("got %v, want zero time", got)
		}
		if end != 4 {
			t.Errorf("end = %d, want 4", end)
		}
	})

	t.Run("truncated", func(t *testing.T) {
		_, _, err := ReadTimeOrNull(nil, 0)
		if !errors.Is(err, ErrTruncated) {
			t.Fatalf("err = %v, want ErrTruncated", err)
		}
	})

	t.Run("not a string", func(t *testing.T) {
		_, _, err := ReadTimeOrNull([]byte(`42`), 0)
		if !errors.Is(err, ErrInvalidJSON) {
			t.Fatalf("err = %v, want ErrInvalidJSON", err)
		}
	})

	t.Run("bad format", func(t *testing.T) {
		_, _, err := ReadTimeOrNull([]byte(`"not-a-time"`), 0)
		if err == nil {
			t.Fatal("expected parse error, got nil")
		}
	})

	// The fast path (parseRFC3339 with atoi2/atoi4) must agree with time.Parse on
	// every shape, including malformed numeric fields that must be rejected so the
	// result still matches the standard library (which the decoder falls back to).
	t.Run("matches time.Parse", func(t *testing.T) {
		cases := []string{
			"2023-01-02T15:04:05Z",
			"2023-01-02T15:04:05.5Z",
			"2023-01-02T15:04:05.123456789Z",
			"2023-12-31T23:59:59Z",
			"2024-02-29T00:00:00Z", // leap day
			"2023-01-02T15:04:05+02:00",
			"2023-01-02T15:04:05-11:30",
			"2023-01-02T15:04:05.250+00:00",
			"0001-01-01T00:00:00Z",
			"2023-1x-02T15:04:05Z",     // non-digit in month
			"2023-01-02T15:04:6zZ",     // non-digit in seconds
			"2023-13-02T15:04:05Z",     // month out of range
			"2023-01-02T15:04:05X",     // bad zone
			"2023-01-02T15:04:05+2:00", // malformed offset
			"not-a-time",
		}
		for _, ts := range cases {
			want, werr := time.Parse(time.RFC3339, ts)
			got, _, gerr := ReadTimeOrNull([]byte(`"`+ts+`"`), 0)
			if (werr == nil) != (gerr == nil) {
				t.Fatalf("%q: stdlib err=%v, ours err=%v", ts, werr, gerr)
			}
			if werr == nil && !got.Equal(want) {
				t.Fatalf("%q: got %v, want %v", ts, got, want)
			}
		}
	})

	// Fuzz the civil→Unix UTC fast path (daysFromCivil) across many dates,
	// including every month, leap years, and century boundaries, against the
	// standard library which it must reproduce exactly.
	t.Run("daysFromCivil vs time.Parse", func(t *testing.T) {
		for _, y := range []int{1, 1969, 1970, 1971, 2000, 2023, 2024, 2100, 2400, 9999} {
			for m := 1; m <= 12; m++ {
				for _, d := range []int{1, 15, 28} {
					ts := fmt.Sprintf("%04d-%02d-%02dT13:14:15.500Z", y, m, d)
					want, err := time.Parse(time.RFC3339, ts)
					if err != nil {
						t.Fatalf("stdlib rejected %q: %v", ts, err)
					}
					got, _, gerr := ReadTimeOrNull([]byte(`"`+ts+`"`), 0)
					if gerr != nil || !got.Equal(want) || got.Location() != want.Location() {
						t.Fatalf("%q: got %v (err=%v), want %v", ts, got, gerr, want)
					}
				}
			}
		}
	})
}

// --- ReadTimeLaxOrNull -----------------------------------------------------

func TestReadTimeLaxOrNull(t *testing.T) {
	utc := func(s string) time.Time {
		ts, err := time.Parse(time.RFC3339, s)
		if err != nil {
			t.Fatalf("setup parse %q: %v", s, err)
		}
		return ts.UTC()
	}

	tests := []struct {
		name string
		in   string
		want time.Time
	}{
		{"rfc3339 T", `"2023-01-02T15:04:05Z"`, utc("2023-01-02T15:04:05Z")},
		{"rfc3339 nanos", `"2023-01-02T15:04:05.123456789Z"`, utc("2023-01-02T15:04:05.123456789Z")},
		{"rfc3339 offset", `"2023-01-02T15:04:05+02:00"`, utc("2023-01-02T15:04:05+02:00")},
		{"space separator", `"2023-01-02 15:04:05Z"`, utc("2023-01-02T15:04:05Z")},
		{"unix seconds", `1700000000`, time.Unix(1700000000, 0).UTC()},
		{"unix millis", `1700000000000`, time.UnixMilli(1700000000000).UTC()},
		{"unix micros", `1700000000000000`, time.UnixMicro(1700000000000000).UTC()},
		{"numeric string", `"1700000000"`, time.Unix(1700000000, 0).UTC()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, end, err := ReadTimeLaxOrNull([]byte(tt.in), 0)
			if err != nil {
				t.Fatalf("unexpected err: %v", err)
			}
			if !got.Equal(tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
			if end != len(tt.in) {
				t.Errorf("end = %d, want %d", end, len(tt.in))
			}
		})
	}

	t.Run("null", func(t *testing.T) {
		got, end, err := ReadTimeLaxOrNull([]byte(`null`), 0)
		if err != nil || !got.IsZero() || end != 4 {
			t.Fatalf("got %v, end %d, err %v", got, end, err)
		}
	})

	// Values it cannot interpret return ErrBadTime (which lax decoding swallows).
	for _, in := range []string{`"not-a-time"`, `"99"`, `42`, `999999999999999999999`} {
		t.Run("bad "+in, func(t *testing.T) {
			if _, _, err := ReadTimeLaxOrNull([]byte(in), 0); !errors.Is(err, ErrBadTime) {
				t.Fatalf("err = %v, want ErrBadTime", err)
			}
		})
	}
}
