package json

import (
	"encoding/json"
	"reflect"
	"testing"
)

func setManyVals(s ...string) [][]byte {
	r := make([][]byte, len(s))
	for i, x := range s {
		r[i] = []byte(x)
	}
	return r
}

func setManyEqual(t *testing.T, got []byte, want string) {
	t.Helper()
	var a, b any
	if err := json.Unmarshal(got, &a); err != nil {
		t.Fatalf("got invalid JSON %q: %v", got, err)
	}
	if err := json.Unmarshal([]byte(want), &b); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSetMany(t *testing.T) {
	// Replace an existing key and insert a missing one in the same pass.
	setManyEqual(t, SetMany([]byte(`{"a":1,"b":2,"c":3}`), nil, setManyVals("9", "8"), []string{"b", "d"}),
		`{"a":1,"b":9,"c":3,"d":8}`)
	// All keys already present.
	setManyEqual(t, SetMany([]byte(`{"a":1,"b":2}`), nil, setManyVals("7", "8"), []string{"a", "b"}),
		`{"a":7,"b":8}`)
	// All keys absent.
	setManyEqual(t, SetMany([]byte(`{"x":0}`), nil, setManyVals("1", "2"), []string{"a", "b"}),
		`{"x":0,"a":1,"b":2}`)
	// Empty object.
	setManyEqual(t, SetMany([]byte(`{}`), nil, setManyVals("1"), []string{"a"}), `{"a":1}`)
	// Non-object root is replaced by a fresh object.
	setManyEqual(t, SetMany([]byte(`5`), nil, setManyVals("1", "2"), []string{"a", "b"}), `{"a":1,"b":2}`)
	// A composite rawVal is inserted verbatim.
	setManyEqual(t, SetMany([]byte(`{"a":1}`), nil, setManyVals(`{"x":[1,2]}`), []string{"obj"}),
		`{"a":1,"obj":{"x":[1,2]}}`)
	// Surplus keys (no matching rawVal) are ignored.
	setManyEqual(t, SetMany([]byte(`{"a":1}`), nil, setManyVals("2"), []string{"a", "b"}), `{"a":2}`)
}

func TestSetManyPreservesWhitespace(t *testing.T) {
	got := string(SetMany([]byte(`{ "a" : 1, "b" : 2 }`), nil, setManyVals("9"), []string{"a"}))
	if got != `{ "a" : 9, "b" : 2 }` {
		t.Errorf("got %q", got)
	}
}

func TestSetManyDoesNotMutateInputAndReusesBuffer(t *testing.T) {
	in := []byte(`{"a":1,"b":2}`)
	inCopy := append([]byte(nil), in...)
	buf := make([]byte, 0, 64)
	out := SetMany(in, buf, setManyVals("8", "9"), []string{"b", "c"})
	setManyEqual(t, out, `{"a":1,"b":8,"c":9}`)
	if !reflect.DeepEqual(in, inCopy) {
		t.Errorf("in mutated: %s", in)
	}
}

func TestSetManyMatchesSet(t *testing.T) {
	doc := []byte(`{"a":1,"b":2,"c":3}`)
	many := SetMany(doc, nil, setManyVals("5"), []string{"b"})
	one := Set(doc, nil, []byte("5"), []string{"b"})
	if !reflect.DeepEqual(many, one) {
		t.Errorf("SetMany %s != Set %s", many, one)
	}
}

func TestSetPaths(t *testing.T) {
	p := func(keys ...string) []string { return keys }

	// Replace existing nested values, sharing the descent into "a".
	setManyEqual(t, SetPaths([]byte(`{"a":{"b":1,"c":2},"d":3}`), nil,
		setManyVals("9", "8"), [][]string{p("a", "b"), p("a", "c")}),
		`{"a":{"b":9,"c":8},"d":3}`)

	// Create missing nested paths, merged under a single created "a".
	setManyEqual(t, SetPaths([]byte(`{"x":0}`), nil,
		setManyVals("1", "2"), [][]string{p("a", "b"), p("a", "c")}),
		`{"x":0,"a":{"b":1,"c":2}}`)

	// Mix: replace a.b (exists), create a.c (within a), create d (at root).
	setManyEqual(t, SetPaths([]byte(`{"a":{"b":1}}`), nil,
		setManyVals("9", "8", "7"), [][]string{p("a", "b"), p("a", "c"), p("d")}),
		`{"a":{"b":9,"c":8},"d":7}`)

	// A path continues but the existing value is not an object: replace it.
	setManyEqual(t, SetPaths([]byte(`{"a":5}`), nil,
		setManyVals("7"), [][]string{p("a", "b")}),
		`{"a":{"b":7}}`)

	// Three levels deep, two leaves sharing the a.b descent.
	setManyEqual(t, SetPaths([]byte(`{"a":{"b":{"x":1,"y":2}}}`), nil,
		setManyVals("8", "9"), [][]string{p("a", "b", "x"), p("a", "b", "y")}),
		`{"a":{"b":{"x":8,"y":9}}}`)

	// A nil/empty path replaces the whole document.
	setManyEqual(t, SetPaths([]byte(`{"a":1}`), nil, setManyVals(`[1,2,3]`), [][]string{nil}), `[1,2,3]`)

	// Non-object root becomes a fresh merged object.
	setManyEqual(t, SetPaths([]byte(`5`), nil,
		setManyVals("1", "2"), [][]string{p("a", "b"), p("c")}),
		`{"a":{"b":1},"c":2}`)
}

func TestSetPathsPreservesWhitespace(t *testing.T) {
	got := string(SetPaths([]byte(`{ "a" : { "b" : 1, "c" : 2 } }`), nil,
		setManyVals("9"), [][]string{{"a", "b"}}))
	if got != `{ "a" : { "b" : 9, "c" : 2 } }` {
		t.Errorf("got %q", got)
	}
}

func TestSetPathsDoesNotMutateInputAndReusesBuffer(t *testing.T) {
	in := []byte(`{"a":{"b":1}}`)
	inCopy := append([]byte(nil), in...)
	out := SetPaths(in, make([]byte, 0, 64), setManyVals("9", "2"),
		[][]string{{"a", "b"}, {"a", "c"}})
	setManyEqual(t, out, `{"a":{"b":9,"c":2}}`)
	if !reflect.DeepEqual(in, inCopy) {
		t.Errorf("in mutated: %s", in)
	}
}

// TestSetPathsMatchesSequentialSet checks SetPaths equals folding Set over the
// paths one at a time, for a mix of replaced and created (prefix-sharing) paths.
func TestSetPathsMatchesSequentialSet(t *testing.T) {
	doc := []byte(`{"a":{"b":1},"keep":true}`)
	paths := [][]string{{"a", "b"}, {"a", "c"}, {"e", "f"}}
	vals := setManyVals("9", "8", "7")

	multi := SetPaths(doc, nil, vals, paths)

	seq := append([]byte(nil), doc...)
	for n, pth := range paths {
		seq = Set(seq, nil, vals[n], pth)
	}
	setManyEqual(t, multi, string(seq))

	// And every set path resolves to its value in the result.
	for n, pth := range paths {
		got, _, err := Get(multi, pth...)
		if err != nil {
			t.Fatalf("Get%v on result: %v", pth, err)
		}
		if string(got) != string(vals[n]) {
			t.Errorf("path %v: result has %q, want %q", pth, got, vals[n])
		}
	}
}
