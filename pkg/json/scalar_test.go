package json

import (
	"encoding/json"
	"errors"
	"testing"
	"unsafe"
)

// String is held to encoding/json's decode of the same token on every
// escape the grammar has, and refuses everything that is not a string token
// with ErrExpectString rather than a decode error or a silent empty.
func TestStringMatchesTheJSONDecode(t *testing.T) {
	tokens := []string{
		`""`, `"a"`, `"api"`, `"a b"`, `"line\nbreak"`, `"quote\"inside"`, `"back\\slash"`,
		`"tab\there"`, `"é"`, `"☃ snow"`, `"😀"`, `"\/slash"`, `"\b\f\r"`,
		`"snow ☃"`, `"trailing "`, `"  "`, `"\"\""`, `"\\\\"`, `"1"`, `"null"`, `"true"`,
	}
	for _, s := range tokens {
		got, err := String([]byte(s))
		if err != nil {
			t.Errorf("String(%s): %v", s, err)
			continue
		}
		var want string
		if err := json.Unmarshal([]byte(s), &want); err != nil {
			t.Fatalf("oracle %s: %v", s, err)
		}
		if got != want {
			t.Errorf("String(%s) = %q, encoding/json %q", s, got, want)
		}
	}
	for _, s := range []string{``, `"`, `"abc`, `abc"`, `1`, `null`, `true`, `[]`, `{}`, `["a"]`, ` "a"`, `"a" `, `'a'`} {
		if _, err := String([]byte(s)); !errors.Is(err, ErrExpectString) {
			t.Errorf("String(%q): err %v, want ErrExpectString", s, err)
		}
	}
	// A bad escape inside a well-delimited token is the unescaper's error,
	// not "not a string".
	for _, s := range []string{`"\x"`, `"\u12"`, `"\"`} {
		_, err := String([]byte(s))
		if err == nil || errors.Is(err, ErrExpectString) {
			t.Errorf("String(%q): err %v, want a decode error", s, err)
		}
	}
}

// The aliasing rule is UnescapeString's and is part of the contract: an
// escape-free token costs no copy, so the result points into raw; an escaped
// one is fresh. Pinned so a change to either half is a deliberate one.
func TestStringAliasesExactlyWhenUnescapeStringDoes(t *testing.T) {
	raw := []byte(`"plain"`)
	s, err := String(raw)
	if err != nil {
		t.Fatal(err)
	}
	if unsafe.StringData(s) != &raw[1] {
		t.Fatal("an escape-free token must alias raw, as UnescapeString does")
	}
	raw = []byte(`"esc\naped"`)
	s, err = String(raw)
	if err != nil {
		t.Fatal(err)
	}
	if p := unsafe.StringData(s); p != nil && uintptr(unsafe.Pointer(p)) >= uintptr(unsafe.Pointer(&raw[0])) && uintptr(unsafe.Pointer(p)) < uintptr(unsafe.Pointer(&raw[0]))+uintptr(len(raw)) {
		t.Fatal("an escaped token must be decoded into a fresh string")
	}
}

func TestBool(t *testing.T) {
	for s, want := range map[string]bool{"true": true, "false": false} {
		got, err := Bool([]byte(s))
		if err != nil || got != want {
			t.Errorf("Bool(%s) = %v, %v", s, got, err)
		}
	}
	for _, s := range []string{``, `null`, `1`, `0`, `"true"`, `tru`, `truex`, `True`, `TRUE`, ` true`, `false `} {
		if _, err := Bool([]byte(s)); !errors.Is(err, ErrExpectBool) {
			t.Errorf("Bool(%q): err %v, want ErrExpectBool", s, err)
		}
	}
}

// The readers take what the walkers give: every string and bool an
// ObjectEach hands back reads through them to the values DecodeAny holds.
func TestScalarReadersAgreeWithDecodeAny(t *testing.T) {
	doc := []byte(`{"s":"a\"b","e":"","u":"é","t":true,"f":false,"n":null,"i":7,"o":{},"a":[]}`)
	want, err := DecodeAny(doc)
	if err != nil {
		t.Fatal(err)
	}
	m := want.(map[string]any)
	err = ObjectEach(doc, func(k string, v []byte) error {
		switch w := m[k].(type) {
		case string:
			got, err := String(v)
			if err != nil || got != w {
				t.Errorf("%s: String = %q, %v; DecodeAny %q", k, got, err, w)
			}
			if _, err := Bool(v); !errors.Is(err, ErrExpectBool) {
				t.Errorf("%s: Bool must refuse a string", k)
			}
		case bool:
			got, err := Bool(v)
			if err != nil || got != w {
				t.Errorf("%s: Bool = %v, %v; DecodeAny %v", k, got, err, w)
			}
			if _, err := String(v); !errors.Is(err, ErrExpectString) {
				t.Errorf("%s: String must refuse a bool", k)
			}
		default:
			if _, err := String(v); !errors.Is(err, ErrExpectString) {
				t.Errorf("%s: String must refuse %s", k, v)
			}
			if _, err := Bool(v); !errors.Is(err, ErrExpectBool) {
				t.Errorf("%s: Bool must refuse %s", k, v)
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func BenchmarkString(b *testing.B) {
	in := []byte(`"entity-service-20260903-1"`)
	b.ReportAllocs()
	for range b.N {
		if _, err := String(in); err != nil {
			b.Fatal(err)
		}
	}
}
