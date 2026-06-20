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
