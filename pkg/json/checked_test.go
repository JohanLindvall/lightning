package json

import (
	"errors"
	"testing"
)

func TestSetChecked(t *testing.T) {
	for _, tc := range []struct {
		name    string
		in      string
		rawVal  string
		keys    []string
		want    string
		wantErr error
	}{
		{"replace", `{"a":1}`, `2`, []string{"a"}, `{"a":2}`, nil},
		{"create nested", `{"a":{}}`, `3`, []string{"a", "b"}, `{"a":{"b":3}}`, nil},
		{"malformed input", `{"a":`, `2`, []string{"a"}, "", ErrInvalidJSON},
		{"malformed input trailing", `{"a":1}}`, `2`, []string{"a"}, "", ErrInvalidJSON},
		{"malformed rawVal", `{"a":1}`, `{`, []string{"a"}, "", ErrInvalidJSON},
		{"empty rawVal", `{"a":1}`, ``, []string{"a"}, "", ErrInvalidJSON},
		{"bare rawVal token", `{"a":1}`, `tru`, []string{"a"}, "", ErrInvalidJSON},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := SetChecked([]byte(tc.in), nil, []byte(tc.rawVal), tc.keys)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("err = %v, want %v", err, tc.wantErr)
			}
			if tc.wantErr != nil {
				if got != nil {
					t.Errorf("got %q on error, want nil", got)
				}
				return
			}
			if string(got) != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestSetManyChecked(t *testing.T) {
	got, err := SetManyChecked([]byte(`{"a":1,"b":2}`), nil,
		[][]byte{[]byte("9"), []byte("3")}, []string{"b", "c"})
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	if want := `{"a":1,"b":9,"c":3}`; string(got) != want {
		t.Errorf("got %q, want %q", got, want)
	}

	// Fewer values than keys is silently tolerated by SetMany (surplus keys
	// ignored); the checked form reports it.
	if _, err := SetManyChecked([]byte(`{"a":1}`), nil,
		[][]byte{[]byte("9")}, []string{"a", "b"}); !errors.Is(err, ErrValueCount) {
		t.Errorf("err = %v, want ErrValueCount", err)
	}
	if _, err := SetManyChecked([]byte(`{"a":1}`), nil,
		[][]byte{[]byte("]")}, []string{"a"}); !errors.Is(err, ErrInvalidJSON) {
		t.Errorf("err = %v, want ErrInvalidJSON for a malformed rawVal", err)
	}
}

func TestSetPathsChecked(t *testing.T) {
	got, err := SetPathsChecked([]byte(`{"a":{"b":1}}`), nil,
		[][]byte{[]byte("9"), []byte("8")}, [][]string{{"a", "b"}, {"a", "c"}})
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	if want := `{"a":{"b":9,"c":8}}`; string(got) != want {
		t.Errorf("got %q, want %q", got, want)
	}

	if _, err := SetPathsChecked([]byte(`{"a":1}`), nil,
		[][]byte{[]byte("9")}, [][]string{{"a"}, {"b"}}); !errors.Is(err, ErrValueCount) {
		t.Errorf("err = %v, want ErrValueCount", err)
	}
	if _, err := SetPathsChecked([]byte(`nope`), nil,
		[][]byte{[]byte("9")}, [][]string{{"a"}}); !errors.Is(err, ErrInvalidJSON) {
		t.Errorf("err = %v, want ErrInvalidJSON for a malformed document", err)
	}
}

func TestStripDefaultsChecked(t *testing.T) {
	defaults := [][]byte{[]byte("0"), []byte("")}

	got, err := StripDefaultsChecked([]byte(`{"a":0,"b":"x"}`), nil, defaults, nil, AssumeCompact)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	if want := `{"b":"x"}`; string(got) != want {
		t.Errorf("got %q, want %q", got, want)
	}

	// Malformed input is copied through verbatim by StripDefaults; the checked form
	// reports it instead.
	if _, err := StripDefaultsChecked([]byte(`{"a":`), nil, defaults, nil, AssumeCompact); !errors.Is(err, ErrInvalidJSON) {
		t.Errorf("err = %v, want ErrInvalidJSON", err)
	}

	// Stripping the document down to nothing is a documented outcome, not an error:
	// the caller is expected to test len() before forwarding the result.
	got, err = StripDefaultsChecked([]byte(`{"a":{"b":0}}`), nil, defaults, nil, AssumeCompact)
	if err != nil {
		t.Fatalf("fully stripped document should not be an error, got %v", err)
	}
	if len(got) != 0 {
		t.Errorf("got %q, want empty", got)
	}
}
