package json

import (
	ejson "encoding/json"
	"errors"
	"math/rand"
	"strings"
	"testing"
)

// validSeeds covers each rule Valid enforces, on both sides of the line. They
// double as the mutation base for the differential tests below.
var validSeeds = []string{
	// well-formed
	`{}`, `[]`, `null`, `true`, `false`, `0`, `-0`, `1`, `1.5`, `1e5`, `1E5`,
	`1e+5`, `1e-5`, `1e05`, `0e0`, `0.0e-0`, `-1.5e-5`, `123456789012345678901234567890`,
	`""`, `"x"`, `"a b"`, `"\""`, `"\\"`, `"\/"`, `"\b\f\n\r\t"`,
	`" "`, `"\uD800"`, `"￿"`, `"😀"`, `{"a":1}`, `{"":0}`,
	`{"a":1,"b":2}`, `[1,2,3]`, `[[[]]]`, `{"a":{"b":[1,{"c":null}]}}`,
	`  {  "a"  :  1  }  `, "\n\t[1,\r2]\n", `[{"a":[]},{"b":{}}]`, `{"a":1,"a":2}`,
	// Accepted by lightning's scanners though encoding/json rejects them; see
	// TestValidDivergesFromStdlib for the explicit list and the reasons.
	`01`, `00`, `-01`, `.5`, `5.`, `1.`, `+1`, `[0123]`, "\"\t\"", "\"\x00\"",

	// malformed
	``, ` `, `,`, `:`, `{`, `}`, `[`, `]`, `x`, `nul`, `nulll`, `tru`, `TRUE`, `True`,
	`-`, `--1`, `1e`, `1e+`, `1.2.3`, `1..2`, `0x1`, `1 2`, `1e309`,
	`{"a":1}{"b":2}`, `[1][2]`,
	`{"a":1,}`, `[1,]`, `{,}`, `[,]`, `[1,,2]`, `{"a":}`, `{:1}`, `{"a" 1}`,
	`{a:1}`, `{1:2}`, `{"a":1,"b"}`, `[1 2]`, `"unterminated`, `"\q"`, `"\u12"`,
	`"\uZZZZ"`, `"\u"`, `"\`, `{"a":[}`, `[{]`, `{]`, `[}`, `[[]`, `{"a":{}`,
	`"a"x`, `NaN`, `Infinity`,
}

// mutateSeeds calls check on every seed plus per byte-level mutations of it —
// substitute, insert, delete and truncate — over an alphabet of the bytes that
// matter to the grammar. This is what catches an off-by-one in the number or
// escape handling.
func mutateSeeds(per int, check func(string)) {
	alphabet := []byte("{}[]\":,\\0123456789.eE+-truefalsn \t\n\r\x00\x1f/uxZ\xff")
	rng := rand.New(rand.NewSource(1))
	for _, seed := range validSeeds {
		check(seed)
		for n := 0; n < per; n++ {
			b := []byte(seed)
			for m := rng.Intn(3) + 1; m > 0; m-- {
				c := alphabet[rng.Intn(len(alphabet))]
				switch rng.Intn(4) {
				case 0:
					if len(b) > 0 {
						b[rng.Intn(len(b))] = c
					}
				case 1:
					p := rng.Intn(len(b) + 1)
					b = append(b[:p:p], append([]byte{c}, b[p:]...)...)
				case 2:
					if len(b) > 0 {
						p := rng.Intn(len(b))
						b = append(b[:p:p], b[p+1:]...)
					}
				case 3:
					if len(b) > 0 {
						b = b[:rng.Intn(len(b))]
					}
				}
			}
			check(string(b))
		}
	}
}

// TestValidMatchesDecodeAny is the contract Valid promises: it accepts exactly the
// documents DecodeAny accepts. DecodeAny is the reference because Valid is built on
// the same readers, so the two can be held to a genuine equivalence.
//
// It is only that. Sharing the scanners with a generated UnmarshalJSON does not make
// this test say anything about one: the schema adds strictness Valid has no view of,
// and the generated decoder is looser both where it skips (unknown fields take
// SkipValue's bracket balancing) and where it reads (the integer readers accept
// "1.2.3" as 1, which Valid rejects). An audit found the earlier wording here —
// that a divergence would mean "Valid lying about whether a decode will succeed" —
// asserting an equivalence nothing tests; Valid's own doc comment now states both
// directions.
func TestValidMatchesDecodeAny(t *testing.T) {
	mutateSeeds(400, func(s string) {
		_, err := DecodeAny([]byte(s))
		if got, want := Valid([]byte(s)), err == nil; got != want {
			t.Errorf("Valid(%q) = %v, but DecodeAny err = %v", s, got, err)
		}
	})
}

// TestValidDivergesFromStdlib pins the places Valid deliberately disagrees with
// encoding/json.Valid, so a change on either side surfaces as a failure rather than
// a silent shift. Each case follows from the decoder's scanners being tuned for
// input already known to be JSON — see Valid's doc comment.
func TestValidDivergesFromStdlib(t *testing.T) {
	for _, tc := range []struct {
		doc  string
		why  string
		want bool // what Valid says; encoding/json.Valid says the opposite
	}{
		{`01`, "leading zero: scanFloat does not check the integer grammar", true},
		{`00`, "leading zero", true},
		{`-01`, "leading zero after sign", true},
		{`.5`, "empty integer part", true},
		{`5.`, "empty fraction", true},
		{`1.`, "empty fraction", true},
		{`+1`, "leading plus", true},
		{`[0123]`, "leading zero inside a container", true},
		{"\"\t\"", `raw control byte: the string scanners stop at '"' and '\' only`, true},
		{"\"\x00\"", "raw control byte", true},
		{"{\"a\":\x001}", "every byte <= 0x20 counts as inter-token whitespace", true},
		{`1e309`, "overflows float64, so the decoder cannot represent it", false},
		{`-1e309`, "overflows float64", false},
	} {
		if got := Valid([]byte(tc.doc)); got != tc.want {
			t.Errorf("Valid(%q) = %v, want %v (%s)", tc.doc, got, tc.want, tc.why)
		}
		if got := ejson.Valid([]byte(tc.doc)); got == tc.want {
			t.Errorf("encoding/json.Valid(%q) = %v, expected it to disagree with Valid (%s)",
				tc.doc, got, tc.why)
		}
	}
}

// TestValidRejectsDeepNesting locks the depth bound. Without it these inputs
// recursed per level and a Go stack overflow is fatal — recover cannot catch it —
// so hostile input took the process down instead of being reported.
func TestValidRejectsDeepNesting(t *testing.T) {
	for _, tc := range []struct {
		name  string
		depth int
		want  bool
	}{
		{"at limit", 10000, true},
		{"past limit", 10001, false},
		{"far past limit", 2_000_000, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			doc := strings.Repeat("[", tc.depth) + strings.Repeat("]", tc.depth)
			if got := Valid([]byte(doc)); got != tc.want {
				t.Errorf("Valid(%d-deep array) = %v, want %v", tc.depth, got, tc.want)
			}
			obj := strings.Repeat(`{"a":`, tc.depth) + "1" + strings.Repeat("}", tc.depth)
			if got := Valid([]byte(obj)); got != tc.want {
				t.Errorf("Valid(%d-deep object) = %v, want %v", tc.depth, got, tc.want)
			}
		})
	}
}

// TestValidZeroAlloc locks the property that motivated the rewrite: Valid checks
// the document in place instead of building (and discarding) the decoded value.
func TestValidZeroAlloc(t *testing.T) {
	doc := []byte(`{"a":[1,2,3,{"b":"x\ny","c":[true,false,null]}],"d":1.5e10}`)
	if !Valid(doc) {
		t.Fatal("seed document must be valid")
	}
	if n := testing.AllocsPerRun(100, func() {
		if !Valid(doc) {
			t.Fatal("invalid")
		}
	}); n != 0 {
		t.Errorf("Valid allocated %v times per run, want 0", n)
	}
}

// FuzzValidMatchesDecodeAny is TestValidMatchesDecodeAny driven by the native
// fuzzer, so the corpus grows past the hand-written mutations.
func FuzzValidMatchesDecodeAny(f *testing.F) {
	for _, s := range validSeeds {
		f.Add([]byte(s))
	}
	f.Fuzz(func(t *testing.T, data []byte) {
		_, err := DecodeAny(data)
		if got, want := Valid(data), err == nil; got != want {
			t.Fatalf("Valid(%q) = %v, but DecodeAny err = %v", data, got, err)
		}
	})
}

func BenchmarkValid(b *testing.B) {
	doc := []byte(`{"RayID":"5b3f2c1d9e7a4b0c","ClientIP":"203.0.113.23",` +
		`"EdgeResponseStatus":599,"Tags":["a","b","c"],"Nested":{"x":[1.5,2.5,3.5],` +
		`"y":{"z":null}},"Escaped":"line\nbreak é and \"quotes\""}`)
	if !Valid(doc) {
		b.Fatal("seed document must be valid")
	}
	b.SetBytes(int64(len(doc)))
	b.ReportAllocs()
	for b.Loop() {
		if !Valid(doc) {
			b.Fatal("invalid")
		}
	}
}

// TestErrMaxDepthReExported checks the sentinel and the bound are reachable from
// this package, so a caller can match the error and reason about the limit without
// importing pkg/unstable — which the package doc tells them not to do.
func TestErrMaxDepthReExported(t *testing.T) {
	doc := []byte(strings.Repeat("[", MaxDepth+1) + strings.Repeat("]", MaxDepth+1))
	if _, err := DecodeAny(doc); !errors.Is(err, ErrMaxDepth) {
		t.Errorf("DecodeAny err = %v, want ErrMaxDepth", err)
	}
	if Valid(doc) {
		t.Error("Valid = true for input past MaxDepth")
	}
}
