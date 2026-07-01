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

var setBenchCases = []struct {
	name    string
	in, raw string
	keys    []string
}{
	// Append a fresh top-level member into a populated object (the common inject case).
	{"append", `{"ClientIP":"1.2.3.4","EdgeResponseStatus":200,"Host":"example.com"}`, `"checkout"`, []string{"Account"}},
	// Append into an empty object (no leading separator).
	{"append_empty", `{}`, `"checkout"`, []string{"Account"}},
	// Replace an existing member's value in place.
	{"replace", `{"ClientIP":"1.2.3.4","EdgeResponseStatus":200,"Host":"example.com"}`, `500`, []string{"EdgeResponseStatus"}},
}

// BenchmarkSet measures Set with a reused output buffer. The append cases exercise the
// member-insertion path, which writes the new member straight into out (no per-call
// throwaway allocation); with a warm buffer they should report zero allocs/op.
func BenchmarkSet(b *testing.B) {
	for _, bc := range setBenchCases {
		b.Run(bc.name, func(b *testing.B) {
			in, raw := []byte(bc.in), []byte(bc.raw)
			out := make([]byte, 0, len(bc.in)+len(bc.raw)+16)
			b.ReportAllocs()
			var sink []byte
			for i := 0; i < b.N; i++ {
				sink = Set(in, out[:0], raw, bc.keys)
			}
			_ = sink
		})
	}
}
