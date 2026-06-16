package json

import (
	"errors"
	"testing"
)

func TestObjectEach(t *testing.T) {
	doc := []byte(`{
		"a": {"x": 1, "y": 2},
		"name": "widget",
		"flag": true,
		"esc\"key": "v"
	}`)

	t.Run("root members", func(t *testing.T) {
		got := map[string]string{}
		err := ObjectEach(doc, func(key string, value []byte) error {
			got[key] = string(value)
			return nil
		})
		if err != nil {
			t.Fatalf("ObjectEach: %v", err)
		}
		want := map[string]string{
			`a`:       `{"x": 1, "y": 2}`,
			`name`:    `"widget"`,
			`flag`:    `true`,
			`esc"key`: `"v"`, // key escapes are decoded; value keeps its quotes
		}
		if len(got) != len(want) {
			t.Fatalf("got %d members, want %d: %v", len(got), len(want), got)
		}
		for k, v := range want {
			if got[k] != v {
				t.Fatalf("member %q = %q, want %q", k, got[k], v)
			}
		}
	})

	t.Run("nested path", func(t *testing.T) {
		var keys []string
		err := ObjectEach(doc, func(key string, value []byte) error {
			keys = append(keys, key+"="+string(value))
			return nil
		}, "a")
		if err != nil {
			t.Fatalf("ObjectEach: %v", err)
		}
		if len(keys) != 2 || keys[0] != "x=1" || keys[1] != "y=2" {
			t.Fatalf("nested iteration = %v", keys)
		}
	})

	t.Run("callback error stops iteration", func(t *testing.T) {
		sentinel := errors.New("stop")
		n := 0
		err := ObjectEach(doc, func(key string, value []byte) error {
			n++
			return sentinel
		})
		if !errors.Is(err, sentinel) {
			t.Fatalf("err = %v, want sentinel", err)
		}
		if n != 1 {
			t.Fatalf("callback ran %d times, want 1", n)
		}
	})

	t.Run("not an object", func(t *testing.T) {
		if err := ObjectEach(doc, func(string, []byte) error { return nil }, "name"); err == nil {
			t.Fatal("expected error iterating a non-object value")
		}
	})

	t.Run("missing key", func(t *testing.T) {
		err := ObjectEach(doc, func(string, []byte) error { return nil }, "nope")
		if !errors.Is(err, ErrKeyNotFound) {
			t.Fatalf("err = %v, want ErrKeyNotFound", err)
		}
	})

	t.Run("empty object", func(t *testing.T) {
		n := 0
		if err := ObjectEach([]byte(`{}`), func(string, []byte) error { n++; return nil }); err != nil {
			t.Fatalf("ObjectEach: %v", err)
		}
		if n != 0 {
			t.Fatalf("callback ran %d times on empty object", n)
		}
	})
}
