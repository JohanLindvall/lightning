package unstable

import (
	"errors"
	"testing"
)

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
