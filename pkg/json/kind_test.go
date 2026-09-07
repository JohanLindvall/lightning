package json

import "testing"

func TestKindOf(t *testing.T) {
	cases := map[string]Kind{
		`null`: KindNull, ` null`: KindNull, "null\n": KindNull, "\t null \r\n": KindNull,
		`true`: KindBool, `false`: KindBool, ` true `: KindBool,
		`0`: KindNumber, `-1`: KindNumber, `+5`: KindNumber, `1.5e3`: KindNumber, `12345678901234567890`: KindNumber,
		`""`: KindString, `"a"`: KindString, `"a\"b"`: KindString, ` "x"`: KindString,
		`[]`: KindArray, `[1,2]`: KindArray, ` [ ] `: KindArray,
		`{}`: KindObject, `{"a":1}`: KindObject,
		// Not a value at all.
		``: KindInvalid, ` `: KindInvalid, "\n": KindInvalid,
		`nul`: KindInvalid, `nullx`: KindInvalid, `null null`: KindInvalid, `truex`: KindInvalid, `True`: KindInvalid, `TRUE`: KindInvalid,
		`fals`: KindInvalid, `x`: KindInvalid, `'a'`: KindInvalid, `.5`: KindInvalid, `NaN`: KindInvalid, `Infinity`: KindInvalid, `undefined`: KindInvalid,
		// Classified by the opening byte, as the scanner dispatches: a
		// malformed string, number or container is still its kind.
		`"unterminated`: KindString, `1x`: KindNumber, `-`: KindNumber, `[1,`: KindArray, `{"a":`: KindObject,
	}
	for in, want := range cases {
		if got := KindOf([]byte(in)); got != want {
			t.Errorf("KindOf(%q) = %v, want %v", in, got, want)
		}
	}
}

// Every value DecodeAny accepts is classified as the Go type DecodeAny
// produced for it — the six kinds against the six representations — and so
// is every member and element of it, read back through the walkers.
func TestKindOfAgreesWithDecodeAny(t *testing.T) {
	kindOfValue := func(v any) Kind {
		switch v.(type) {
		case nil:
			return KindNull
		case bool:
			return KindBool
		case float64:
			return KindNumber
		case string:
			return KindString
		case []any:
			return KindArray
		case map[string]any:
			return KindObject
		}
		return KindInvalid
	}
	docs := []string{
		`null`, `true`, `false`, `0`, `-1.5e2`, `""`, `"s"`, `[]`, `{}`,
		`{"n":null,"b":true,"f":false,"i":1,"x":-2.5,"s":"a\"b","a":[1,"2",null,[],{}],"o":{"k":[true]}}`,
		` [ null , true , 1 , "x" , [ ] , { } ] `,
	}
	for _, d := range docs {
		want, err := DecodeAny([]byte(d))
		if err != nil {
			t.Fatalf("%s: %v", d, err)
		}
		if got, w := KindOf([]byte(d)), kindOfValue(want); got != w {
			t.Errorf("KindOf(%s) = %v, DecodeAny holds a %v", d, got, w)
		}
		switch w := want.(type) {
		case map[string]any:
			_ = ObjectEach([]byte(d), func(k string, v []byte) error {
				if got, want := KindOf(v), kindOfValue(w[k]); got != want {
					t.Errorf("%s: member %q KindOf = %v, want %v", d, k, got, want)
				}
				return nil
			})
		case []any:
			i := 0
			_ = ArrayEach([]byte(d), func(v []byte) error {
				if got, want := KindOf(v), kindOfValue(w[i]); got != want {
					t.Errorf("%s: element %d KindOf = %v, want %v", d, i, got, want)
				}
				i++
				return nil
			})
		}
	}
}

func TestKindString(t *testing.T) {
	for k, want := range map[Kind]string{KindInvalid: "invalid", KindNull: "null", KindBool: "bool", KindNumber: "number", KindString: "string", KindArray: "array", KindObject: "object", Kind(200): "invalid"} {
		if got := k.String(); got != want {
			t.Errorf("Kind(%d).String() = %q, want %q", k, got, want)
		}
	}
}

// TestKindOfWhitespaceIsThePackagesOwn pins the rule KindOf tolerates at both
// ends of a value: this package's whitespace is every byte <= 0x20, the
// one-compare shortcut unstable.SkipWS takes and the decoder and Valid inherit
// (CLAUDE.md's Valid entry records why it is deliberate). An earlier form
// skipped leading bytes by that rule and trailing ones by the grammar's
// four-byte set, so it called `null` followed by a NUL invalid while Valid and
// DecodeAny accepted the same document.
func TestKindOfWhitespaceIsThePackagesOwn(t *testing.T) {
	for _, ws := range []string{" ", "\t", "\n", "\r", "\x00", "\x01", "\x1f", "\x20"} {
		for _, c := range []struct {
			val  string
			want Kind
		}{{"null", KindNull}, {"true", KindBool}, {"false", KindBool}, {"1", KindNumber}, {`"a"`, KindString}} {
			for _, doc := range []string{ws + c.val, c.val + ws, ws + c.val + ws} {
				got := KindOf([]byte(doc))
				if got != c.want {
					t.Errorf("KindOf(%q) = %v, want %v", doc, got, c.want)
				}
				// Whatever KindOf says of a document, Valid must agree that
				// the document is one value: that is the property the two ends
				// sharing a whitespace rule buys.
				if !Valid([]byte(doc)) {
					t.Errorf("Valid(%q) = false, but KindOf says %v", doc, got)
				}
			}
		}
	}
	// A byte above the whitespace set still ends the literal.
	for _, doc := range []string{"null!", "nullx", "true?", "false~", "null null"} {
		if got := KindOf([]byte(doc)); got != KindInvalid {
			t.Errorf("KindOf(%q) = %v, want invalid", doc, got)
		}
	}
}
