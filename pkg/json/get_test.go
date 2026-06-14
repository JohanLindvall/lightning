package json

import (
	"errors"
	"testing"
)

func TestGet(t *testing.T) {
	doc := []byte(`{
		"a": {"b": {"c": 42, "s": "hi\n", "arr": [1, 2, 3]}},
		"name": "widget",
		"flag": true,
		"nil": null
	}`)

	tests := []struct {
		name    string
		keys    []string
		want    string
		wantErr error
	}{
		{"root", nil, string(doc), nil},
		{"top scalar", []string{"name"}, `"widget"`, nil},
		{"top bool", []string{"flag"}, `true`, nil},
		{"top null", []string{"nil"}, `null`, nil},
		{"nested object", []string{"a", "b"}, `{"c": 42, "s": "hi\n", "arr": [1, 2, 3]}`, nil},
		{"nested scalar", []string{"a", "b", "c"}, `42`, nil},
		{"nested string keeps quotes+escape", []string{"a", "b", "s"}, `"hi\n"`, nil},
		{"nested array", []string{"a", "b", "arr"}, `[1, 2, 3]`, nil},
		{"missing top key", []string{"nope"}, "", ErrKeyNotFound},
		{"missing nested key", []string{"a", "b", "nope"}, "", ErrKeyNotFound},
		{"descend into non-object", []string{"name", "x"}, "", nil}, // err set, checked below
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, off, err := Get(doc, tt.keys...)
			if tt.name == "descend into non-object" {
				if err == nil {
					t.Fatalf("expected error descending into a string")
				}
				return
			}
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("err = %v, want %v", err, tt.wantErr)
			}
			if tt.wantErr != nil {
				return
			}
			if string(got) != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
			// offset must point at the returned value within doc.
			if off < 0 || off+len(got) > len(doc) || string(doc[off:off+len(got)]) != string(got) {
				t.Fatalf("offset %d does not locate the value in doc", off)
			}
		})
	}
}

// TestGetObjectCheck demonstrates the jsonparser.Get(... ) != Object idiom
// rebuilt on Get: navigate to the parent path and confirm it is an object.
func TestGetObjectCheck(t *testing.T) {
	doc := []byte(`{"parent": {"child": 1}, "leaf": 5}`)

	if v, _, err := Get(doc, "parent"); err != nil || len(v) == 0 || v[0] != '{' {
		t.Fatalf("expected parent to be an object, got %q err=%v", v, err)
	}
	if v, _, err := Get(doc, "leaf"); err == nil && len(v) > 0 && v[0] == '{' {
		t.Fatalf("leaf should not look like an object: %q", v)
	}
}
