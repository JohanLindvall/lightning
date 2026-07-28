package json

import (
	"errors"
	"reflect"
	"testing"
)

func TestArrayEach(t *testing.T) {
	doc := []byte(`{
		"records": [
			{"a": 1, "s": "tricky \" ]} string", "n": [1, [2, 3]]},
			"plain",
			true,
			42.5,
			null,
			[]
		],
		"other": [9]
	}`)

	t.Run("nested path", func(t *testing.T) {
		var got []string
		err := ArrayEach(doc, func(value []byte) error {
			got = append(got, string(value))
			return nil
		}, "records")
		if err != nil {
			t.Fatalf("ArrayEach: %v", err)
		}
		want := []string{
			`{"a": 1, "s": "tricky \" ]} string", "n": [1, [2, 3]]}`,
			`"plain"`, // strings keep their quotes, as with Get
			`true`,
			`42.5`,
			`null`,
			`[]`,
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("elements = %q, want %q", got, want)
		}
	})

	t.Run("root array", func(t *testing.T) {
		n := 0
		err := ArrayEach([]byte(` [ 1 , 2 , 3 ] `), func(value []byte) error {
			n++
			return nil
		})
		if err != nil || n != 3 {
			t.Fatalf("n=%d err=%v", n, err)
		}
	})

	t.Run("callback error stops iteration", func(t *testing.T) {
		sentinel := errors.New("stop")
		n := 0
		err := ArrayEach(doc, func([]byte) error {
			n++
			return sentinel
		}, "records")
		if !errors.Is(err, sentinel) {
			t.Fatalf("err = %v, want sentinel", err)
		}
		if n != 1 {
			t.Fatalf("callback ran %d times, want 1", n)
		}
	})

	t.Run("not an array", func(t *testing.T) {
		err := ArrayEach([]byte(`{"records": {"a": 1}}`), func([]byte) error { return nil }, "records")
		if !errors.Is(err, ErrExpectArray) {
			t.Fatalf("err = %v, want ErrExpectArray", err)
		}
	})

	t.Run("missing key", func(t *testing.T) {
		err := ArrayEach(doc, func([]byte) error { return nil }, "nope")
		if !errors.Is(err, ErrKeyNotFound) {
			t.Fatalf("err = %v, want ErrKeyNotFound", err)
		}
	})

	t.Run("empty array", func(t *testing.T) {
		n := 0
		if err := ArrayEach([]byte(`[]`), func([]byte) error { n++; return nil }); err != nil {
			t.Fatalf("ArrayEach: %v", err)
		}
		if n != 0 {
			t.Fatalf("callback ran %d times on empty array", n)
		}
	})

	// Strictness the callers rely on: a malformed array is an error, never a
	// silently fused or dropped element.
	t.Run("missing comma", func(t *testing.T) {
		err := ArrayEach([]byte(`[{"a":1} {"b":2}]`), func([]byte) error { return nil })
		if !errors.Is(err, ErrInvalidJSON) {
			t.Fatalf("err = %v, want ErrInvalidJSON", err)
		}
	})

	t.Run("trailing comma", func(t *testing.T) {
		if err := ArrayEach([]byte(`[1,]`), func([]byte) error { return nil }); err == nil {
			t.Fatal("a trailing comma must error")
		}
	})

	t.Run("truncated", func(t *testing.T) {
		for _, in := range []string{`[`, `[1`, `[1,`, `[{"a":1}`} {
			if err := ArrayEach([]byte(in), func([]byte) error { return nil }); err == nil {
				t.Fatalf("%q: want an error", in)
			}
		}
	})
}

func TestArrayEachCompactMatchesArrayEach(t *testing.T) {
	docs := []string{
		`[1,2,3]`,
		`[{"a":1},{"b":[true,null]},"s"]`,
		`{"records":[{"x":"y"},2]}`,
		`[]`,
	}
	for _, doc := range docs {
		var keys []string
		if doc[0] == '{' {
			keys = []string{"records"}
		}
		var a, b []string
		if err := ArrayEach([]byte(doc), func(v []byte) error { a = append(a, string(v)); return nil }, keys...); err != nil {
			t.Fatalf("%s: ArrayEach: %v", doc, err)
		}
		if err := ArrayEachCompact([]byte(doc), func(v []byte) error { b = append(b, string(v)); return nil }, keys...); err != nil {
			t.Fatalf("%s: ArrayEachCompact: %v", doc, err)
		}
		if !reflect.DeepEqual(a, b) {
			t.Fatalf("%s: compact diverged: %q vs %q", doc, a, b)
		}
	}
}

// The differential the scanner readers are held to: whenever DecodeAny accepts
// the document as an array, ArrayEach must accept it too and yield raw spans
// that re-decode to exactly DecodeAny's elements.
func FuzzArrayEachMatchesDecodeAny(f *testing.F) {
	for _, seed := range []string{
		`[]`, `[1]`, `[1,2,3]`, `[ "a" , {"b": [1,2]} , null , true ]`,
		`[[[[1]]]]`, `["\"]",", ["]`, `[1.5e10,-0]`,
		`[{}]`, `[""]`, `[0]`,
	} {
		f.Add([]byte(seed))
	}
	f.Fuzz(func(t *testing.T, data []byte) {
		want, err := DecodeAny(data)
		arr, ok := want.([]any)
		if err != nil || !ok {
			return // not a valid array document; ArrayEach's behavior is unconstrained
		}
		got := []any{} // DecodeAny yields a non-nil empty slice for []
		eachErr := ArrayEach(data, func(v []byte) error {
			elem, derr := DecodeAny(v)
			if derr != nil {
				t.Fatalf("element %q does not re-decode: %v", v, derr)
			}
			got = append(got, elem)
			return nil
		})
		if eachErr != nil {
			t.Fatalf("DecodeAny accepted an array ArrayEach rejects: %v (input %q)", eachErr, data)
		}
		if len(got) != len(arr) {
			t.Fatalf("element count %d, want %d (input %q)", len(got), len(arr), data)
		}
		if !reflect.DeepEqual(got, arr) {
			t.Fatalf("elements diverge from DecodeAny (input %q)", data)
		}
	})
}
