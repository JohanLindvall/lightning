package json

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDecodeAny(t *testing.T) {
	cases := []string{
		`null`, `true`, `false`, `42`, `-3.14e2`, `"hi"`, `"escé\n"`,
		`[]`, `{}`, `[1,"a",true,null,{"x":[2,3]}]`,
		`{"a":1,"b":[true,"x",2.5],"c":{"d":null,"e":"f"}}`,
	}
	for _, c := range cases {
		got, err := DecodeAny([]byte(c), false)
		if err != nil {
			t.Errorf("%q: %v", c, err)
			continue
		}
		var want any
		if err := json.Unmarshal([]byte(c), &want); err != nil {
			t.Fatalf("bad case %q: %v", c, err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%q: got %#v want %#v", c, got, want)
		}
		// compact=true must agree on these whitespace-free inputs
		gc, err := DecodeAny([]byte(c), true)
		if err != nil || !reflect.DeepEqual(gc, want) {
			t.Errorf("%q compact: got %#v err %v", c, gc, err)
		}
	}
	// leading/trailing whitespace is fine (non-compact)
	if v, err := DecodeAny([]byte("  \n [ 1 , 2 ] \t"), false); err != nil || !reflect.DeepEqual(v, []any{1.0, 2.0}) {
		t.Errorf("whitespace: got %#v err %v", v, err)
	}
	// trailing non-whitespace is rejected
	if _, err := DecodeAny([]byte(`1 2`), false); err == nil {
		t.Error("trailing data accepted")
	}
}
