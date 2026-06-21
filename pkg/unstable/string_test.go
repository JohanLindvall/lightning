package unstable

import (
	"errors"
	"testing"
)

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
		{"base64 padded", `"eyJ4IjoxfQ=="`, `{"x":1}`, nil}, // base64({"x":1})
		{"base64 unpadded", `"eyJ4IjoxfQ"`, `{"x":1}`, nil}, // RawStd
		{"base64 array", `"WzEsMiwzXQ=="`, `[1,2,3]`, nil},  // base64([1,2,3])
		{"null", `null`, "", nil},                           // nil document
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
