package json

import "testing"

func TestKindOf(t *testing.T) {
	cases := map[string]Kind{
		`null`: Null, ` null`: Null, "null\n": Null, "\t null \r\n": Null,
		`true`: Bool, `false`: Bool, ` true `: Bool,
		`0`: Number, `-1`: Number, `+5`: Number, `1.5e3`: Number, `12345678901234567890`: Number,
		`""`: String, `"a"`: String, `"a\"b"`: String, ` "x"`: String,
		`[]`: Array, `[1,2]`: Array, ` [ ] `: Array,
		`{}`: Object, `{"a":1}`: Object,
		// Not a value at all.
		``: Invalid, ` `: Invalid, "\n": Invalid,
		`nul`: Invalid, `nullx`: Invalid, `null null`: Invalid, `truex`: Invalid, `True`: Invalid, `TRUE`: Invalid,
		`fals`: Invalid, `x`: Invalid, `'a'`: Invalid, `.5`: Invalid, `NaN`: Invalid, `Infinity`: Invalid, `undefined`: Invalid,
		// Classified by the opening byte, as the scanner dispatches: a
		// malformed string, number or container is still its kind.
		`"unterminated`: String, `1x`: Number, `-`: Number, `[1,`: Array, `{"a":`: Object,
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
			return Null
		case bool:
			return Bool
		case float64:
			return Number
		case string:
			return String
		case []any:
			return Array
		case map[string]any:
			return Object
		}
		return Invalid
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
	for k, want := range map[Kind]string{Invalid: "invalid", Null: "null", Bool: "bool", Number: "number", String: "string", Array: "array", Object: "object", Kind(200): "invalid"} {
		if got := k.String(); got != want {
			t.Errorf("Kind(%d).String() = %q, want %q", k, got, want)
		}
	}
}
