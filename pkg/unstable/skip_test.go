package unstable

import (
	"errors"
	"testing"
)

// --- SkipValue (and its SkipString/Number/Object/Array helpers) ------------

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
