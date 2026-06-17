package json

import (
	"encoding/json"
	"testing"
)

func TestSet(t *testing.T) {
	cases := []struct {
		name    string
		in, raw string
		keys    []string
		want    string
	}{
		{"replace existing", `{"a":1,"b":2}`, `9`, []string{"a"}, `{"a":9,"b":2}`},
		{"replace last", `{"a":1,"b":2}`, `9`, []string{"b"}, `{"a":1,"b":9}`},
		{"insert into non-empty", `{"a":1}`, `2`, []string{"b"}, `{"a":1,"b":2}`},
		{"insert into empty", `{}`, `2`, []string{"b"}, `{"b":2}`},
		{"replace nested", `{"a":{"b":1}}`, `5`, []string{"a", "b"}, `{"a":{"b":5}}`},
		{"insert nested", `{"a":{"b":1}}`, `5`, []string{"a", "c"}, `{"a":{"b":1,"c":5}}`},
		{"create intermediate", `{"a":1}`, `5`, []string{"x", "y"}, `{"a":1,"x":{"y":5}}`},
		{"overwrite non-object", `{"a":1}`, `5`, []string{"a", "y"}, `{"a":{"y":5}}`},
		{"string value", `{"a":"x"}`, `"z"`, []string{"a"}, `{"a":"z"}`},
		{"object value", `{"a":1}`, `{"k":2}`, []string{"b"}, `{"a":1,"b":{"k":2}}`},
		{"deep create", `{}`, `1`, []string{"a", "b", "c"}, `{"a":{"b":{"c":1}}}`},
		{"no keys replaces root", `{"a":1}`, `[1,2]`, nil, `[1,2]`},
		{"whitespace preserved", `{  "a" : 1 }`, `2`, []string{"a"}, `{  "a" : 2 }`},
		{"append after whitespace", `{ "a":1 }`, `2`, []string{"b"}, `{ "a":1,"b":2 }`},
		{"empty input", ``, `1`, []string{"a"}, `{"a":1}`},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := string(Set([]byte(c.in), nil, []byte(c.raw), c.keys))
			if got != c.want {
				t.Fatalf("Set(%q, %q, %v) = %q, want %q", c.in, c.raw, c.keys, got, c.want)
			}
			if !json.Valid([]byte(got)) {
				t.Fatalf("Set(%q, %q, %v) = %q is not valid JSON", c.in, c.raw, c.keys, got)
			}
		})
	}
}

func TestSetDoesNotMutateInput(t *testing.T) {
	in := []byte(`{"a":1}`)
	orig := string(in)
	out := Set(in, make([]byte, 0, 64), []byte(`2`), []string{"b"})
	if string(in) != orig {
		t.Fatalf("input mutated: %q", in)
	}
	if string(out) != `{"a":1,"b":2}` {
		t.Fatalf("out = %q", out)
	}
}

func TestSetReusesBuffer(t *testing.T) {
	buf := make([]byte, 0, 64)
	out := Set([]byte(`{"a":1}`), buf, []byte(`2`), []string{"b"})
	if string(out) != `{"a":1,"b":2}` {
		t.Fatalf("out = %q", out)
	}
	if &buf[:1][0] != &out[:1][0] {
		t.Fatalf("expected Set to reuse the provided buffer")
	}
}
