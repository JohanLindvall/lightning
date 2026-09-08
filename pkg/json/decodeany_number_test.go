package json

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
)

// DecodeAnyNumber answers what encoding/json answers under UseNumber: every
// number a json.Number holding its literal, containers and the rest as
// DecodeAny already has them.
func TestDecodeAnyNumberMatchesUseNumber(t *testing.T) {
	for _, doc := range []string{
		`0.95`, `12345678901234567890`, `-0`, `1e21`,
		`{"a":[1,2.5,{"b":3}],"c":"x","d":null,"e":true}`,
		`[]`, `{}`, `[0.1,0.2,0.30000000000000004]`,
	} {
		got, err := DecodeAnyNumber([]byte(doc))
		if err != nil {
			t.Fatalf("%s: %v", doc, err)
		}
		gotC, err := DecodeAnyNumberCompact([]byte(doc))
		if err != nil {
			t.Fatalf("%s compact: %v", doc, err)
		}
		dec := json.NewDecoder(bytes.NewReader([]byte(doc)))
		dec.UseNumber()
		var want any
		if err := dec.Decode(&want); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, want) || !reflect.DeepEqual(gotC, want) {
			t.Errorf("%s: %#v / %#v, UseNumber %#v", doc, got, gotC, want)
		}
	}
	if _, err := DecodeAnyNumber([]byte(`1e400x`)); err == nil {
		t.Error("a malformed number must still fail")
	}
}
