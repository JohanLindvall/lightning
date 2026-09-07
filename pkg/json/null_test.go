package json

import (
	"errors"
	"testing"
)

// A null where the container would be is a container with nothing in it:
// the callback is never called and the walk succeeds — for both walkers, at
// the root and behind a path, in the plain and the compact forms. This is
// what unmarshalling null into a map or a slice does, and what a wire format
// means by `"metric":null` or `"values":null`.
func TestAContainerThatIsNullIsEmpty(t *testing.T) {
	type walker struct {
		name string
		run  func(data []byte, keys ...string) (calls int, err error)
	}
	walkers := []walker{
		{"ObjectEach", func(data []byte, keys ...string) (int, error) {
			n := 0
			err := ObjectEach(data, func(string, []byte) error { n++; return nil }, keys...)
			return n, err
		}},
		{"ObjectEachCompact", func(data []byte, keys ...string) (int, error) {
			n := 0
			err := ObjectEachCompact(data, func(string, []byte) error { n++; return nil }, keys...)
			return n, err
		}},
		{"ArrayEach", func(data []byte, keys ...string) (int, error) {
			n := 0
			err := ArrayEach(data, func([]byte) error { n++; return nil }, keys...)
			return n, err
		}},
		{"ArrayEachCompact", func(data []byte, keys ...string) (int, error) {
			n := 0
			err := ArrayEachCompact(data, func([]byte) error { n++; return nil }, keys...)
			return n, err
		}},
	}
	for _, w := range walkers {
		for _, c := range []struct {
			doc  string
			keys []string
		}{
			{`null`, nil},
			{` null `, nil},
			{`{"a":null}`, []string{"a"}},
			{`{"a":{"b":null}}`, []string{"a", "b"}},
			{`{"metric":null,"values":null}`, []string{"values"}},
		} {
			calls, err := w.run([]byte(c.doc), c.keys...)
			if err != nil || calls != 0 {
				t.Errorf("%s(%s, %v): calls %d, err %v; want no calls and no error", w.name, c.doc, c.keys, calls, err)
			}
		}
	}
}

// Only the literal null is empty: a misspelling, a truncation, and every
// other non-container value stay the errors they were.
func TestOnlyTheLiteralNullIsAnEmptyContainer(t *testing.T) {
	for _, doc := range []string{`nul`, `nullx`, `nu`, `n`, `NULL`, `"null"`, `0`, `true`, `false`} {
		if err := ObjectEach([]byte(doc), func(string, []byte) error { return nil }); err == nil {
			t.Errorf("ObjectEach(%q) must still error", doc)
		} else if len(doc) >= 4 && !errors.Is(err, ErrExpectObject) && !errors.Is(err, ErrTruncated) {
			t.Errorf("ObjectEach(%q): %v, want ErrExpectObject", doc, err)
		}
		if err := ArrayEach([]byte(doc), func([]byte) error { return nil }); err == nil {
			t.Errorf("ArrayEach(%q) must still error", doc)
		} else if len(doc) >= 4 && !errors.Is(err, ErrExpectArray) && !errors.Is(err, ErrTruncated) {
			t.Errorf("ArrayEach(%q): %v, want ErrExpectArray", doc, err)
		}
	}
	// The wrong container is still the wrong container.
	if err := ArrayEach([]byte(`{}`), func([]byte) error { return nil }); !errors.Is(err, ErrExpectArray) {
		t.Errorf("ArrayEach({}) = %v, want ErrExpectArray", err)
	}
	if err := ObjectEach([]byte(`[]`), func(string, []byte) error { return nil }); !errors.Is(err, ErrExpectObject) {
		t.Errorf("ObjectEach([]) = %v, want ErrExpectObject", err)
	}
}

// A null ELEMENT or MEMBER is a value like any other and is still handed to
// the callback as the four bytes "null" — the change is only to the container
// the walk is asked to iterate.
func TestANullMemberIsStillDelivered(t *testing.T) {
	var got []string
	if err := ArrayEach([]byte(`[null, 1, null]`), func(v []byte) error { got = append(got, string(v)); return nil }); err != nil {
		t.Fatal(err)
	}
	if len(got) != 3 || got[0] != "null" || got[2] != "null" {
		t.Fatalf("elements = %q, want the nulls delivered", got)
	}
	got = nil
	if err := ObjectEach([]byte(`{"a":null}`), func(k string, v []byte) error { got = append(got, k+"="+string(v)); return nil }); err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 || got[0] != "a=null" {
		t.Fatalf("members = %q, want a=null delivered", got)
	}
}

// TestNullContainerWhitespaceIsThePackagesOwn holds isNullToken's boundary set
// to the rule the rest of the package uses — every byte <= 0x20 is whitespace
// — so that a document Valid accepts as a null is an empty container here too,
// whatever whitespace follows it.
func TestNullContainerWhitespaceIsThePackagesOwn(t *testing.T) {
	for _, ws := range []string{" ", "\t", "\n", "\r", "\x00", "\x01", "\x1f", "\x20"} {
		doc := []byte("null" + ws)
		if !Valid(doc) {
			t.Fatalf("premise: Valid(%q) is false", doc)
		}
		called := false
		if err := ObjectEach(doc, func(string, []byte) error { called = true; return nil }); err != nil || called {
			t.Errorf("ObjectEach(%q) = %v, called=%v; want nil, false", doc, err, called)
		}
		if err := ArrayEach(doc, func([]byte) error { called = true; return nil }); err != nil || called {
			t.Errorf("ArrayEach(%q) = %v, called=%v; want nil, false", doc, err, called)
		}
	}
	// A byte above the whitespace set still makes it a misspelling.
	for _, doc := range []string{"null!", "nullx", "null~"} {
		if err := ObjectEach([]byte(doc), func(string, []byte) error { return nil }); err == nil {
			t.Errorf("ObjectEach(%q) = nil, want an error", doc)
		}
	}
}
