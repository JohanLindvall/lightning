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

// TestStripDefaultsCheckedModes covers the wrapper's whole outcome space in every
// whitespace mode: a document stripped away entirely (as an object and as a bare
// scalar root, each with and without surrounding whitespace), one stripped in
// part, one nothing is stripped from, and the two inputs that hold no document at
// all. The rows are mode-independent on purpose — that is the contract, and it is
// what PreserveWhitespace used to break: keeping the input's outer whitespace left
// the consumed-document case with a whitespace-only result, which is not Valid, so
// the wrapper reported ErrInvalidJSON on input it had just validated. Those two
// rows (" {\"a\":0} " and " true\n" under PreserveWhitespace) fail against the
// unfixed code.
func TestStripDefaultsCheckedModes(t *testing.T) {
	zero := [][]byte{[]byte("0")}
	yes := [][]byte{[]byte("true")}

	rows := []struct {
		name     string
		in       string
		defaults [][]byte
		want     string
		wantErr  error
	}{
		{"fully stripped object", `{"a":0}`, zero, "", nil},
		{"fully stripped object, padded", ` {"a":0} `, zero, "", nil},
		{"fully stripped nested", `{"a":[{"b":0}]}`, zero, "", nil},
		{"fully stripped scalar root", `true`, yes, "", nil},
		{"fully stripped scalar root, padded", " true\n", yes, "", nil},
		{"partially stripped", `{"a":0,"b":1}`, zero, `{"b":1}`, nil},
		{"nothing stripped", `{"b":1}`, zero, `{"b":1}`, nil},
		{"empty input", ``, zero, "", ErrInvalidJSON},
		{"whitespace-only input", `  `, zero, "", ErrInvalidJSON},
		{"malformed input", `{"a":`, zero, "", ErrInvalidJSON},
	}

	for _, m := range stripWSModes {
		for _, row := range rows {
			t.Run(m.name+"/"+row.name, func(t *testing.T) {
				got, err := StripDefaultsChecked([]byte(row.in), nil, row.defaults, nil, m.ws)
				if !errors.Is(err, row.wantErr) {
					t.Fatalf("err = %v, want %v", err, row.wantErr)
				}
				if row.wantErr != nil {
					if got != nil {
						t.Errorf("got %q on error, want a nil slice", got)
					}
					return
				}
				if string(got) != row.want {
					t.Errorf("got %q, want %q", got, row.want)
				}
			})
		}
	}

	// The consumed-document result is empty in every mode, which is what makes
	// len() the documented test — even where the unchecked function, keeping the
	// document's outer whitespace, hands back bytes.
	if raw := StripDefaults([]byte(` {"a":0} `), nil, zero, nil, PreserveWhitespace); string(raw) != "  " {
		t.Errorf("premise: unchecked PreserveWhitespace = %q, want the two outer spaces", raw)
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
