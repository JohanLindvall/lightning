package support

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

// --- Unwrap ----------------------------------------------------------------

func TestUnwrap(t *testing.T) {
	tests := []struct {
		name    string
		in      string // a JSON string value (the outer quotes included)
		want    string // the embedded JSON it should yield
		wantErr error
	}{
		{"plain object", `"{\"x\":1}"`, `{"x":1}`, nil},
		{"plain array", `"[1,2,3]"`, `[1,2,3]`, nil},
		{"plain with leading ws", `"  {\"x\":1}"`, `  {"x":1}`, nil},
		{"plain scalar", `"42"`, `42`, nil},
		{"base64 padded", `"eyJ4IjoxfQ=="`, `{"x":1}`, nil},          // base64({"x":1})
		{"base64 unpadded", `"eyJ4IjoxfQ"`, `{"x":1}`, nil},          // RawStd
		{"base64 array", `"WzEsMiwzXQ=="`, `[1,2,3]`, nil},           // base64([1,2,3])
		{"null", `null`, "", nil},                                   // nil document
		{"not a string", `42`, "", ErrInvalidJSON},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := Unwrap([]byte(tt.in), 0)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("err = %v, want %v", err, tt.wantErr)
			}
			if tt.wantErr == nil && string(got) != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
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
func TestFastFloat(t *testing.T) {
	fastFloat := func(b []byte) (float64, bool) {
		f, end, fast, ok := scanFloat(b, 0)
		return f, ok && fast && end == len(b)
	}
	okCases := map[string]float64{
		"0":      0,
		"1":      1,
		"-1":     -1,
		"+2":     2,
		"3.5":    3.5,
		"0.001":  0.001,
		"12e2":   1200,
		"12E2":   1200,
		"1.5e-1": 0.15,
	}
	for in, want := range okCases {
		got, ok := fastFloat([]byte(in))
		if !ok {
			t.Errorf("fastFloat(%q) ok=false, want true", in)
			continue
		}
		if got != want {
			t.Errorf("fastFloat(%q) = %v, want %v", in, got, want)
		}
	}

	declined := []string{
		"",                     // empty
		"1e23",                 // exp above 22
		"1e-23",                // exp below -22
		"1e",                   // missing exp digits
		"1e+",                  // missing exp digits after sign
		"12345678901234567890", // 20 digits, > 19
		"1e99999",              // huge exponent
		".",                    // no digits
		"1x",                   // trailing junk -> not fully consumed
	}
	for _, in := range declined {
		if _, ok := fastFloat([]byte(in)); ok {
			t.Errorf("fastFloat(%q) ok=true, want false (should defer to strconv)", in)
		}
	}
}

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
			"2023-1x-02T15:04:05Z",  // non-digit in month
			"2023-01-02T15:04:6zZ",  // non-digit in seconds
			"2023-13-02T15:04:05Z",  // month out of range
			"2023-01-02T15:04:05X",  // bad zone
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

// --- ExpectNull ------------------------------------------------------------

func TestExpectNull(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		i       int
		wantEnd int
		wantErr error
	}{
		{"valid", `null`, 0, 4, nil},
		{"valid trailing", `null,`, 0, 4, nil},
		{"offset", `xxnull`, 2, 6, nil},
		{"truncated", `nul`, 0, 0, ErrInvalidJSON},
		{"misspelled", `nxll`, 0, 0, ErrInvalidJSON},
		{"not null", `true`, 0, 0, ErrInvalidJSON},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			end, err := ExpectNull([]byte(tt.in), tt.i)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("err = %v, want %v", err, tt.wantErr)
			}
			if end != tt.wantEnd {
				t.Errorf("end = %d, want %d", end, tt.wantEnd)
			}
		})
	}
}

// --- SkipValue (and its skipString/Number/Object/Array helpers) ------------

func TestSkipValue(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		wantEnd int
		wantErr error
	}{
		{"string", `"abc"rest`, 5, nil},
		{"string with escape", `"a\"b"x`, 6, nil},
		{"long string SIMD", `"` + longUnescaped() + `"x`, len(longUnescaped()) + 2, nil},
		{"number", `123,`, 3, nil},
		{"negative number", `-1.5e3]`, 6, nil},
		{"true", `true,`, 4, nil},
		{"false", `false,`, 5, nil},
		{"null", `null,`, 4, nil},
		{"empty object", `{}`, 2, nil},
		{"empty array", `[]`, 2, nil},
		{"object", `{"a":1,"b":[2,3]}x`, 17, nil},
		{"nested", `{"a":{"b":[1,{"c":2}]}}`, 23, nil},
		{"array of strings", `["x","y","z"]`, 13, nil},
		{"object with long key (SIMD)", `{"` + longUnescaped() + `":1}`, len(longUnescaped()) + 6, nil},
		{"array with whitespace", `[ 1 , 2 ]`, 9, nil},

		{"empty", ``, 0, ErrTruncated},
		{"bad number", `x`, 0, ErrBadNumber},
		{"truncated true", `tru`, 0, ErrInvalidJSON},
		{"truncated false", `fal`, 0, ErrInvalidJSON},
		{"truncated string", `"abc`, 4, ErrTruncated},
		{"truncated string mid-escape", `"a\`, 3, ErrTruncated},
		{"truncated object", `{"a":1`, 6, ErrTruncated},
		{"truncated array", `[1,2`, 4, ErrTruncated},

		// Exercise the recursive descent and stray-bracket branches in
		// skipObject / skipArray.
		{"array of arrays", `[[1],[2,[3]]]x`, 13, nil},
		{"array of objects", `[{"a":1},{"b":2}]x`, 17, nil},
		{"object of arrays", `{"a":[1,2],"b":[]}x`, 18, nil},
		{"object of objects", `{"a":{"b":{}}}x`, 14, nil},
		{"stray close brace in array", `[}]x`, 3, nil},
		{"stray close bracket in object", `{]}x`, 3, nil},
		{"nested array truncated", `[[1,2`, 5, ErrTruncated},
		{"nested object truncated", `{"a":{"b":1`, 11, ErrTruncated},
		{"array nested object truncated", `[{"a":1`, 7, ErrTruncated},
		{"object nested array truncated", `{"a":[1,2`, 9, ErrTruncated},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			end, err := SkipValue([]byte(tt.in), 0)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("err = %v, want %v", err, tt.wantErr)
			}
			if end != tt.wantEnd {
				t.Errorf("end = %d, want %d", end, tt.wantEnd)
			}
		})
	}
}

// --- SkipWS ----------------------------------------------------------------

func TestSkipWS(t *testing.T) {
	tests := []struct {
		name string
		in   string
		i    int
		want int
	}{
		{"no ws", `abc`, 0, 0},
		{"leading spaces", `   abc`, 0, 3},
		{"all ws kinds", " \t\n\rx", 0, 4},
		{"all whitespace", "   ", 0, 3},
		{"from offset", `a   b`, 1, 4},
		{"at end", `abc`, 3, 3},
		{"empty", ``, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SkipWS([]byte(tt.in), tt.i); got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

// --- DecodeValue (and decodeAnyObject/decodeAnyArray) ----------------------

func TestDecodeValue(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		v, end, err := DecodeValue([]byte(`"hi"`), 0)
		if err != nil || v != "hi" || end != 4 {
			t.Fatalf("got %v, %d, %v", v, end, err)
		}
	})
	t.Run("number", func(t *testing.T) {
		v, _, err := DecodeValue([]byte(`3.5`), 0)
		if err != nil || v != 3.5 {
			t.Fatalf("got %v, %v", v, err)
		}
	})
	t.Run("true", func(t *testing.T) {
		v, _, err := DecodeValue([]byte(`true`), 0)
		if err != nil || v != true {
			t.Fatalf("got %v, %v", v, err)
		}
	})
	t.Run("false", func(t *testing.T) {
		v, _, err := DecodeValue([]byte(`false`), 0)
		if err != nil || v != false {
			t.Fatalf("got %v, %v", v, err)
		}
	})
	t.Run("null", func(t *testing.T) {
		v, end, err := DecodeValue([]byte(`null`), 0)
		if err != nil || v != nil || end != 4 {
			t.Fatalf("got %v, %d, %v", v, end, err)
		}
	})
	t.Run("empty truncated", func(t *testing.T) {
		_, _, err := DecodeValue(nil, 0)
		if !errors.Is(err, ErrTruncated) {
			t.Fatalf("err = %v, want ErrTruncated", err)
		}
	})

	t.Run("object", func(t *testing.T) {
		v, end, err := DecodeValue([]byte(`{"a":1,"b":"x","c":true}`), 0)
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		m, ok := v.(map[string]any)
		if !ok {
			t.Fatalf("got %T, want map", v)
		}
		if m["a"] != float64(1) || m["b"] != "x" || m["c"] != true {
			t.Errorf("unexpected map: %v", m)
		}
		if end != 24 {
			t.Errorf("end = %d, want 24", end)
		}
	})

	t.Run("empty object", func(t *testing.T) {
		v, end, err := DecodeValue([]byte(`{}`), 0)
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		m := v.(map[string]any)
		if len(m) != 0 || end != 2 {
			t.Errorf("got %v, end %d", m, end)
		}
	})

	t.Run("object with whitespace", func(t *testing.T) {
		v, _, err := DecodeValue([]byte(`{ "a" : 1 , "b" : 2 }`), 0)
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		m := v.(map[string]any)
		if m["a"] != float64(1) || m["b"] != float64(2) {
			t.Errorf("unexpected map: %v", m)
		}
	})

	t.Run("nested array", func(t *testing.T) {
		v, _, err := DecodeValue([]byte(`[1,"two",[3,4],{"k":5}]`), 0)
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		a := v.([]any)
		if len(a) != 4 {
			t.Fatalf("len = %d, want 4", len(a))
		}
		if a[0] != float64(1) || a[1] != "two" {
			t.Errorf("unexpected array head: %v", a)
		}
		inner := a[2].([]any)
		if inner[0] != float64(3) || inner[1] != float64(4) {
			t.Errorf("unexpected inner array: %v", inner)
		}
		obj := a[3].(map[string]any)
		if obj["k"] != float64(5) {
			t.Errorf("unexpected inner object: %v", obj)
		}
	})

	t.Run("empty array", func(t *testing.T) {
		v, end, err := DecodeValue([]byte(`[]`), 0)
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		if len(v.([]any)) != 0 || end != 2 {
			t.Errorf("got %v, end %d", v, end)
		}
	})

	// Error paths for the object/array decoders.
	objErrs := []struct {
		name    string
		in      string
		wantErr error
	}{
		{"obj truncated after brace", `{`, ErrTruncated},
		{"obj bad key", `{1:2}`, ErrInvalidJSON},
		{"obj missing colon", `{"a" 1}`, ErrExpectColon},
		{"obj truncated after value", `{"a":1`, ErrTruncated},
		{"obj missing comma", `{"a":1 "b":2}`, ErrInvalidJSON},
		{"obj bad value", `{"a":}`, ErrBadNumber},
		{"obj truncated value", `{"a":`, ErrTruncated},
		{"arr truncated after bracket", `[`, ErrTruncated},
		{"arr truncated after value", `[1`, ErrTruncated},
		{"arr missing comma", `[1 2]`, ErrInvalidJSON},
		{"arr bad value", `[}]`, ErrBadNumber},
	}
	for _, tt := range objErrs {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := DecodeValue([]byte(tt.in), 0)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("err = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

// --- internal scanners: scalar implementations -----------------------------

func TestIndexCloseOrEscapeScalar(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{``, 0},
		{`abc`, 3},
		{`ab"c`, 2},
		{`ab\c`, 2},
		{`a\"b`, 1},  // backslash before quote
		{`"`, 0},     // quote first
		{`\`, 0},     // backslash first
		{`abc"def`, 3},
		{strings.Repeat("x", 50) + `"`, 50}, // long, quote at end
		{strings.Repeat("x", 50) + `\`, 50}, // long, backslash at end
	}
	for _, tt := range tests {
		if got := indexCloseOrEscapeScalar([]byte(tt.in)); got != tt.want {
			t.Errorf("indexCloseOrEscapeScalar(%q) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestIndexStructuralScalar(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{``, 0},
		{`abc`, 3},
		{`ab{c`, 2},
		{`ab}c`, 2},
		{`ab[c`, 2},
		{`ab]c`, 2},
		{`ab"c`, 2},
		{`123:456`, 7}, // none present
		{strings.Repeat("x", 40) + `{`, 40},
	}
	for _, tt := range tests {
		if got := indexStructuralScalar([]byte(tt.in)); got != tt.want {
			t.Errorf("indexStructuralScalar(%q) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

// indexCloseOrEscape / indexStructural dispatch to SIMD on long inputs; verify
// they agree with the scalar reference across a range of lengths and positions.
func TestIndexFunctionsMatchScalar(t *testing.T) {
	base := strings.Repeat("abcdefgh", 16) // 128 bytes, no special chars
	for _, n := range []int{0, 1, 15, 16, 31, 32, 33, 64, 100, 128} {
		s := base[:n]
		if got, want := indexCloseOrEscape([]byte(s)), indexCloseOrEscapeScalar([]byte(s)); got != want {
			t.Errorf("indexCloseOrEscape(len=%d) = %d, want %d", n, got, want)
		}
		if got, want := indexStructural([]byte(s)), indexStructuralScalar([]byte(s)); got != want {
			t.Errorf("indexStructural(len=%d) = %d, want %d", n, got, want)
		}
		// Insert a target byte at each position and confirm both paths find it.
		for _, c := range []byte{'"', '\\', '{', '}', '[', ']'} {
			for pos := 0; pos < n; pos++ {
				b := []byte(s)
				b[pos] = c
				if got, want := indexCloseOrEscape(b), indexCloseOrEscapeScalar(b); got != want {
					t.Fatalf("indexCloseOrEscape(len=%d, %q@%d) = %d, want %d", n, c, pos, got, want)
				}
				if got, want := indexStructural(b), indexStructuralScalar(b); got != want {
					t.Fatalf("indexStructural(len=%d, %q@%d) = %d, want %d", n, c, pos, got, want)
				}
			}
		}
	}
}

// --- unsafeStr -------------------------------------------------------------

func TestUnsafeStr(t *testing.T) {
	if got := unsafeStr(nil); got != "" {
		t.Errorf("unsafeStr(nil) = %q, want empty", got)
	}
	if got := unsafeStr([]byte{}); got != "" {
		t.Errorf("unsafeStr(empty) = %q, want empty", got)
	}
	if got := unsafeStr([]byte("hello")); got != "hello" {
		t.Errorf("unsafeStr = %q, want hello", got)
	}
}

// --- error sentinels are distinct -----------------------------------------

func TestErrorSentinels(t *testing.T) {
	errs := []error{
		ErrInvalidJSON, ErrTruncated, ErrBadEscape, ErrBadUnicode,
		ErrBadNumber, ErrExpectColon, ErrExpectObject, ErrExpectArray,
	}
	for i := range errs {
		if errs[i] == nil {
			t.Fatalf("sentinel %d is nil", i)
		}
		for j := range errs {
			if i != j && errors.Is(errs[i], errs[j]) {
				t.Errorf("sentinel %d and %d compare equal", i, j)
			}
		}
	}
}

// TestCountArrayElements covers the SkipValue-based element counter: nested
// arrays, strings with structural bytes inside, objects, whitespace, and the
// malformed/empty cases that must return 0 (so the caller falls back to append).
func TestCountArrayElements(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{`[]`, 0},
		{`[1]`, 1},
		{`[1,2,3]`, 3},
		{` [ 1 , 2 , 3 ] `, 3}, // leading ws before '[' is not consumed by the caller; data[i]=='['
		{`["a","b"]`, 2},
		{`["a,b","c]d"]`, 2},        // commas/brackets inside strings are not counted
		{`[{"x":1},{"y":[2,3]}]`, 2}, // object elements
		{`[[1,2],[3,4],[5,6]]`, 3},   // nested arrays (the coordinate shape)
		{`[[[1,2],[3,4]]]`, 1},       // one outer element
		{`[null,true,false]`, 3},
		{`[`, 0},        // truncated
		{`[1,2`, 0},     // unterminated
		{`[1,2,]`, 0},   // trailing comma -> bail to append
		{`{}`, 0},       // not an array
		{``, 0},         // empty input
	}
	for _, tt := range tests {
		// The caller positions i at '[' (after its own SkipWS); mirror that.
		i := 0
		for i < len(tt.in) && (tt.in[i] == ' ' || tt.in[i] == '\t') {
			i++
		}
		if got := CountArrayElements([]byte(tt.in), i); got != tt.want {
			t.Errorf("CountArrayElements(%q) = %d, want %d", tt.in, got, tt.want)
		}
	}
}
