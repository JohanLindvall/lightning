// generator_test.go drives the generator end to end: every case is a schema
// source that is generated into a throwaway module and then COMPILED.
//
// Compilation is the assertion that matters. The generator's failure mode is
// not a panic or a returned error — it is exiting 0 after writing a package
// that does not build: a decoder function declared twice because a name was
// reserved without consulting the reservation set, an import decided by
// substring-scanning generated text that also contains JSON key literals, a
// foreign RawMessage look-alike emitted with no import at all. Nothing short of
// handing the output to the compiler catches those, which is why this file
// spends a `go build` per case rather than asserting on generated text: an
// assertion on text pins today's formatting, while an assertion on the compiler
// pins the only property that matters.
//
// The schema sources below are raw string literals, so their struct tags are
// written as interpreted-string literals ("json:\"x\"") instead of the usual
// backticks — a backtick cannot appear inside a raw string, and Go accepts
// either literal form as a tag.
package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// probeEmpty is the default probe. The schemas declare `package main` so the
// generated decoder can be compiled by itself in a temp module, and a main
// package needs a func main to build; a case that only checks compilation has
// nothing to run.
const probeEmpty = "package main\n\nfunc main() {}\n"

type genCase struct {
	name string
	// schema is written to data.go and handed to the generator.
	schema string
	// probe, when set, replaces the empty func main with a program that decodes
	// and prints; the case then compiles AND runs it, asserting want against its
	// stdout. Anything compared against encoding/json must go through a
	// methodless twin (type rootStd Root): Root has a generated UnmarshalJSON,
	// so json.Unmarshal into a *Root delegates straight back to lightning and
	// would compare the decoder against itself.
	probe string
	want  string
	// wantErr, when set, is a substring the generator's error must contain. Such
	// a case compiles nothing — a failed run writes no decoder.
	wantErr string
	// wantWarn and wantNoWarn are substrings that must, and must not, appear on
	// the generator's diagnostic stream.
	wantWarn   []string
	wantNoWarn []string
}

var genCases = []genCase{
	{
		// G1. namedStruct reserved its decoder name by writing g.used directly
		// while every other emitter goes through g.uniq, so the anonymous struct
		// behind the "item" field — reached second, and renamed by uniq only if
		// the name is already taken — was handed the same name. Order-dependent:
		// with the fields the other way round (the case below) uniq sees the
		// collision and renames, so a field reorder flipped a working schema to
		// one that does not compile.
		name: "named_and_anonymous_decoder_name_collision",
		schema: `package main

type Item struct {
	ID int "json:\"id\""
}

type Root struct {
	Item struct {
		Label string "json:\"label\""
	} "json:\"item\""
	Items []Item "json:\"items\""
}
`,
	},
	{
		name: "anonymous_before_named_decoder_name",
		schema: `package main

type Item struct {
	ID int "json:\"id\""
}

type Root struct {
	Items []Item "json:\"items\""
	Item  struct {
		Label string "json:\"label\""
	} "json:\"item\""
}
`,
	},
	{
		// G2. The generated file's imports used to be decided by scanning the
		// generated text for "time.Time" / "json.Number" / "unsafe.Sizeof" — but
		// JSON key literals are emitted into the dispatch switch as string
		// constants, so a key that merely looks like a qualified type name
		// dragged in an import nothing used.
		name: "json_key_literals_that_look_like_types",
		schema: `package main

type Root struct {
	A int "json:\"time.Time\""
	B int "json:\"json.Number\""
	C int "json:\"json.RawMessage\""
	D int "json:\"unsafe.Sizeof\""
}
`,
	},
	{
		// G2, the other direction: a schema can *use* time.Time and
		// json.RawMessage in positions whose decode emits no qualified name at
		// all (the readers take and return the value, the raw field is a byte
		// span), so the generated file must import neither. An import claimed by
		// a flag that fires per field rather than per emission fails to compile
		// here, "imported and not used".
		name: "imported_types_that_are_never_emitted",
		schema: `package main

import (
	"encoding/json"
	"time"
)

type Root struct {
	T  time.Time       "json:\"t\""
	R  json.RawMessage "json:\"r\""
	RN json.RawMessage "json:\"rn,nocopy\""
}
`,
	},
	{
		// G2/G3. encoding/json under an alias: the generated file has to spell
		// the qualifier the schema uses and import it under that same alias.
		name: "aliased_encoding_json_import",
		schema: `package main

import ej "encoding/json"

type Root struct {
	N  ej.Number       "json:\"n\""
	R  ej.RawMessage   "json:\"r\""
	Rs []ej.RawMessage "json:\"rs\""
	Ns []ej.Number     "json:\"ns\""
	PN *ej.Number      "json:\"pn\""
	M  map[string]ej.RawMessage "json:\"m\""
}
`,
	},
	{
		name: "aliased_time_import",
		schema: `package main

import gotime "time"

type Root struct {
	T  gotime.Time            "json:\"t\""
	Ts []gotime.Time          "json:\"ts\""
	PT *gotime.Time           "json:\"pt\""
	M  map[string]gotime.Time "json:\"m\""
	L  gotime.Time            "json:\"l,lax\""
}
`,
	},
	{
		// G3. isRaw matched RawMessage/RawValue qualified by ANY package, so a
		// foreign type of that name was decoded as if it were encoding/json's —
		// generator exits 0, package does not compile (or, worse, compiles and
		// appends raw bytes to something that is not a []byte). It must be
		// reported as the unsupported type it is. Nothing is compiled here: the
		// run fails, so no decoder is written.
		name: "foreign_rawmessage_is_not_encoding_json",
		schema: `package main

import "lightningprobe/foreign"

type Root struct {
	R foreign.RawMessage "json:\"r\""
	V foreign.RawValue   "json:\"v\""
}
`,
		wantErr: "unsupported type foreign.RawMessage",
	},
	{
		// C1. encoding/json calls RawMessage's UnmarshalJSON even for a JSON
		// null (documented), so the field ends up holding the four bytes "null".
		// lightning skipped the assignment entirely, which lost the null/absent
		// distinction RawMessage exists to preserve and — worse — left the
		// PREVIOUS document's value in place when decoding into a reused target,
		// the pattern this library encourages.
		name: "raw_message_null_matches_stdlib",
		schema: `package main

import "encoding/json"

type Root struct {
	Raw   json.RawMessage "json:\"raw\""
	RawNC json.RawMessage "json:\"rawnc,nocopy\""
}

// rootStd is the methodless twin used for the stdlib comparison.
type rootStd Root
`,
		probe: `package main

import (
	"encoding/json"
	"fmt"
)

const (
	nullDoc = "{\"raw\":null,\"rawnc\":null}"
	keepDoc = "{\"raw\":{\"keep\":1},\"rawnc\":{\"keep\":1}}"
)

func main() {
	var v Root
	if err := v.UnmarshalJSON([]byte(nullDoc)); err != nil {
		panic(err)
	}
	var s rootStd
	if err := json.Unmarshal([]byte(nullDoc), &s); err != nil {
		panic(err)
	}
	fmt.Printf("fresh lightning=%q,%q stdlib=%q,%q\n", v.Raw, v.RawNC, s.Raw, s.RawNC)

	var v2 Root
	if err := v2.UnmarshalJSON([]byte(keepDoc)); err != nil {
		panic(err)
	}
	if err := v2.UnmarshalJSON([]byte(nullDoc)); err != nil {
		panic(err)
	}
	var s2 rootStd
	if err := json.Unmarshal([]byte(keepDoc), &s2); err != nil {
		panic(err)
	}
	if err := json.Unmarshal([]byte(nullDoc), &s2); err != nil {
		panic(err)
	}
	fmt.Printf("reuse lightning=%q,%q stdlib=%q,%q\n", v2.Raw, v2.RawNC, s2.Raw, s2.RawNC)
}
`,
		want: `fresh lightning="null","null" stdlib="null","null"
reuse lightning="null","null" stdlib="null","null"
`,
	},
	{
		// A1. An unrecognized struct-tag option was silently ignored, so a typo'd
		// ",nocpy" quietly turned off the aliasing the author asked for. Warned,
		// not failed: unlike a //lightning: directive (which only ever comes from
		// this generator), a json tag arrives from encoding/json, and rejecting
		// its vocabulary outright would break every migrating schema.
		// ",omitempty" is encode-only and meaningless when decoding, so it is
		// accepted in silence.
		name: "unknown_json_tag_options",
		schema: `package main

type Root struct {
	A string "json:\"a,nocpy\""
	B string "json:\"b,omitempty\""
	C int    "json:\"c,string\""
	D string "json:\"d,nocopy,lax\""
	E string "json:\"-\""
}
`,
		// The field labels are the load-bearing half of these assertions: the
		// options lightning does act on, and the one it accepts in silence, must
		// draw no diagnostic at all.
		wantWarn:   []string{`"nocpy"`, "field A", `"string"`, "field C"},
		wantNoWarn: []string{"omitempty", "field B", "field D", "field E"},
	},
	{
		// A2. A json tag name encoding/json's isValidTag rejects (it allows
		// letters, digits and one punctuation set; a quote, backslash or control
		// byte is out) makes the stdlib discard the WHOLE tag and key the field by
		// its Go field name. lightning has no such rule, so the two disagree in
		// both directions with no error on either side: lightning answers to a key
		// the stdlib never looks for, and not to the field name it falls back to.
		// Warned rather than rejected (that would fail a schema decoding correctly
		// today) and rather than adopting the fallback (that would silently change
		// which key an existing decoder answers to) — the ",string" precedent.
		//
		// The probe is the divergence itself, both directions, through the
		// methodless twin. '|' is deliberately in the schema as a NON-case: it is
		// in the stdlib's allowed set, so lightning's alternate names draw no
		// warning even though the stdlib reads them as one long key.
		name: "invalid_json_tag_names",
		schema: `package main

type Root struct {
	Q    int "json:\"a\\\"b\""
	BS   int "json:\"a\\\\b\""
	NL   int "json:\"a\\nb\""
	OK   int "json:\"ok.name-1/2\""
	Alt  int "json:\"alt|other\""
	Sp   int "json:\"with space\""
}

type rootStd Root
`,
		probe: `package main

import (
	"encoding/json"
	"fmt"
)

const (
	tagKeys   = "{\"a\\\"b\":1,\"a\\\\b\":2,\"a\\nb\":3}"
	fieldKeys = "{\"Q\":1,\"BS\":2,\"NL\":3}"
)

func main() {
	var v Root
	if err := v.UnmarshalJSON([]byte(tagKeys)); err != nil {
		panic(err)
	}
	var s rootStd
	if err := json.Unmarshal([]byte(tagKeys), &s); err != nil {
		panic(err)
	}
	fmt.Printf("tag keys:   lightning=%d,%d,%d stdlib=%d,%d,%d\n", v.Q, v.BS, v.NL, s.Q, s.BS, s.NL)

	var v2 Root
	if err := v2.UnmarshalJSON([]byte(fieldKeys)); err != nil {
		panic(err)
	}
	var s2 rootStd
	if err := json.Unmarshal([]byte(fieldKeys), &s2); err != nil {
		panic(err)
	}
	fmt.Printf("field keys: lightning=%d,%d,%d stdlib=%d,%d,%d\n", v2.Q, v2.BS, v2.NL, s2.Q, s2.BS, s2.NL)
}
`,
		want: `tag keys:   lightning=1,2,3 stdlib=0,0,0
field keys: lightning=0,0,0 stdlib=1,2,3
`,
		wantWarn:   []string{"field Q", "field BS", "field NL", "encoding/json"},
		wantNoWarn: []string{"field OK", "field Alt", "field Sp"},
	},
	{
		// The README used to say that embedding a type from another package "is
		// decoded as a single named field instead of being flattened". It is not:
		// only the three selector types the generator knows decode at all, and
		// anything else fails the run — which is the right failure mode (silently
		// decoding an opaque foreign type would be worse) but was documented as
		// the opposite. Pinned here so the doc and the behavior cannot drift apart
		// again; the known-selector half is pinned by the sibling case below.
		name: "foreign_embedded_type_is_an_error",
		schema: `package main

import "strings"

type Root struct {
	strings.Builder
	A int "json:\"a\""
}
`,
		wantErr: "unsupported type strings.Builder",
	},
	{
		// The exception the README sentence was reaching for: the three selector
		// types the generator does know decode as named fields when embedded,
		// keyed by the type name (encoding/json keys them the same way, and for
		// time.Time and json.RawMessage then hands the whole document to the
		// promoted UnmarshalJSON — see conformance's
		// TestEmbeddedUnmarshalerDivergesFromStdlib).
		name: "known_selector_embeds_decode_as_named_fields",
		schema: `package main

import (
	"encoding/json"
	"time"
)

type Root struct {
	time.Time
	json.Number
	A int "json:\"a\""
}

type Raw struct {
	json.RawMessage
	B int "json:\"b\""
}
`,
		probe: `package main

import "fmt"

func main() {
	var v Root
	if err := v.UnmarshalJSON([]byte("{\"Time\":\"2021-01-02T03:04:05Z\",\"Number\":12,\"a\":7}")); err != nil {
		panic(err)
	}
	fmt.Printf("%s %s %d\n", v.Time.Format("2006-01-02"), v.Number, v.A)
	var r Raw
	if err := r.UnmarshalJSON([]byte("{\"RawMessage\":{\"z\":1},\"b\":8}")); err != nil {
		panic(err)
	}
	fmt.Printf("%s %d\n", r.RawMessage, r.B)
}
`,
		want: "2021-01-02 12 7\n{\"z\":1} 8\n",
	},
	{
		// D1. Directive validation only looked at comments the parser had
		// attached to a type declaration, so a blank line between the directive
		// and its type — or a directive stranded above the package clause —
		// silently disabled both the directive and the typo check.
		name: "detached_directives",
		schema: `//lightning:destructive
package main

//lightning:arena

type Detached struct {
	Vals []float64 "json:\"vals\""
}

//lightning:compact
type Attached struct {
	X int "json:\"x\""
}
`,
		wantWarn:   []string{"//lightning:destructive", "//lightning:arena", "not attached to a type declaration"},
		wantNoWarn: []string{"//lightning:compact"},
	},
	{
		name: "detached_unknown_directive",
		schema: `package main

//lightning:arna

type Root struct {
	X int "json:\"x\""
}
`,
		wantWarn: []string{"//lightning:arna"},
	},
	{
		// The attached typo stays a hard error, as the 2026-08 pass made it.
		name: "unknown_directive_is_an_error",
		schema: `package main

//lightning:arna
type Root struct {
	X int "json:\"x\""
}
`,
		wantErr: "unknown directive //lightning:arna",
	},
	{
		name: "duplicate_json_name_is_an_error",
		schema: `package main

type Root struct {
	A int "json:\"x\""
	B int "json:\"x\""
}
`,
		wantErr: `json name "x" is mapped more than once`,
	},
	{
		// The ordinary schema: every supported field shape in one type, so a
		// change that breaks any of them fails to compile here.
		name: "kitchen_sink",
		schema: `package main

import (
	"encoding/json"
	"time"
)

type Nested struct {
	Name  string "json:\"name\""
	Count int    "json:\"count\""
}

type Embedded struct {
	E1 string "json:\"e1\""
}

type PtrEmbedded struct {
	E2 string "json:\"e2\""
}

type Root struct {
	Embedded
	*PtrEmbedded

	Str  string  "json:\"str\""
	NC   string  "json:\"nc,nocopy\""
	Bool bool    "json:\"bool\""
	I8   int8    "json:\"i8\""
	U16  uint16  "json:\"u16\""
	F32  float32 "json:\"f32\""
	F64  float64 "json:\"f64\""

	Num  json.Number  "json:\"num\""
	PNum *json.Number "json:\"pnum\""

	Time    time.Time  "json:\"time\""
	TimeLax time.Time  "json:\"timeLax,lax\""
	PTime   *time.Time "json:\"ptime\""

	Raw   json.RawMessage "json:\"raw\""
	RawNC json.RawMessage "json:\"rawnc,nocopy\""

	Nested Nested "json:\"nested\""
	Anon   struct {
		X int       "json:\"x\""
		W time.Time "json:\"w\""
	} "json:\"anon\""

	Ints     []int             "json:\"ints\""
	Strs     []string          "json:\"strs\""
	Grid     [][]int           "json:\"grid\""
	Items    []Nested          "json:\"items\""
	PtrItems []*Nested         "json:\"ptrItems\""
	Times    []time.Time       "json:\"times\""
	Raws     []json.RawMessage "json:\"raws\""

	Arr  [3]int    "json:\"arr\""
	UArr [3]uint32 "json:\"uarr\""
	FArr [2]byte   "json:\"farr\""

	Bytes  []byte   "json:\"bytes\""
	Chunks [][]byte "json:\"chunks\""

	M  map[string]int    "json:\"m\""
	MN map[string]Nested "json:\"mn\""

	PI *int    "json:\"pi\""
	PN *Nested "json:\"pn\""

	Any  any   "json:\"any\""
	Anys []any "json:\"anys\""

	Alt      int    "json:\"status|EdgeStatus\""
	Ignored  int    "json:\"-\""
	Embed    Nested "json:\"embedded,unwrap\""
	LaxSkip  int    "json:\"laxSkip,lax\""
	LaxSlice []int  "json:\"laxSlice,lax\""
}

type PointList []struct {
	X int "json:\"x\""
}

type ScoreMap map[string]int

type ByteBlob []byte
`,
	},
	{
		// One root per directive, each reaching a shape the directive actually
		// changes, so all four variants are compiled.
		name: "directives",
		schema: `package main

//lightning:compact
type CompactDoc struct {
	X    int      "json:\"x\""
	Strs []string "json:\"strs\""
}

//lightning:nocopy
type NoCopyMap map[string]string

//lightning:destructive
type DestructiveDoc struct {
	Name string   "json:\"name,nocopy\""
	Tags []string "json:\"tags,nocopy\""
}

//lightning:arena
type ArenaDoc struct {
	Pos []float64 "json:\"pos\""
	Rot []int32   "json:\"rot\""
	Kid *ArenaDoc "json:\"kid\""
}
`,
		wantNoWarn: []string{"no effect"},
	},
	{
		// Recursive and mutually recursive schemas: their decoders thread the
		// depth counter, so every signature and call site has to agree.
		name: "recursive_types",
		schema: `package main

type Tree struct {
	Name string  "json:\"name\""
	Kids []*Tree "json:\"kids\""
}

type RingRoot struct {
	Start *Ring1 "json:\"start\""
}

type Ring1 struct {
	Name string "json:\"name\""
	Next *Ring2 "json:\"next\""
}

type Ring2 struct {
	Count int    "json:\"count\""
	Back  *Ring1 "json:\"back\""
}
`,
	},
	{
		// A decode-and-compare over the ordinary field shapes, against the
		// methodless twin. Deliberately avoids the two places lightning is
		// documented to differ from encoding/json (an empty JSON array yields a
		// nil slice, key matching is case-sensitive).
		name: "stdlib_parity_decode",
		schema: `package main

import "encoding/json"

type Nested struct {
	Name  string "json:\"name\""
	Count int    "json:\"count\""
}

type Root struct {
	Str    string          "json:\"str\""
	I      int             "json:\"i\""
	F      float64         "json:\"f\""
	B      bool            "json:\"b\""
	Num    json.Number     "json:\"num\""
	Raw    json.RawMessage "json:\"raw\""
	Nested Nested          "json:\"nested\""
	Items  []Nested        "json:\"items\""
	Ints   []int           "json:\"ints\""
	Arr    [3]int          "json:\"arr\""
	Bytes  []byte          "json:\"bytes\""
	M      map[string]int  "json:\"m\""
	PI     *int            "json:\"pi\""
	PN     *Nested         "json:\"pn\""
	Any    any             "json:\"any\""
	Absent string          "json:\"absent\""
	Null   *Nested         "json:\"null\""
}

type rootStd Root
`,
		probe: `package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

const doc = "{\"str\":\"hi \\u00e9\",\"i\":-7,\"f\":1.5,\"b\":true,\"num\":12.25," +
	"\"raw\":{\"a\":[1,2]},\"nested\":{\"name\":\"n\",\"count\":2}," +
	"\"items\":[{\"name\":\"a\",\"count\":1},{\"name\":\"b\",\"count\":2}]," +
	"\"ints\":[1,2,3],\"arr\":[4,5],\"bytes\":\"aGk=\",\"m\":{\"k\":9}," +
	"\"pi\":11,\"pn\":{\"name\":\"p\",\"count\":3}," +
	"\"any\":{\"x\":[1,\"s\",true,null]},\"null\":null,\"extra\":{\"skipped\":[1]}}"

func main() {
	var v Root
	if err := v.UnmarshalJSON([]byte(doc)); err != nil {
		panic(err)
	}
	var s rootStd
	if err := json.Unmarshal([]byte(doc), &s); err != nil {
		panic(err)
	}
	if !reflect.DeepEqual(v, Root(s)) {
		fmt.Printf("MISMATCH\nlightning %+v\nstdlib    %+v\n", v, Root(s))
		return
	}
	fmt.Println("equal")
}
`,
		want: "equal\n",
	},
}

func TestGenerate(t *testing.T) {
	t.Parallel()

	repo, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	goMod := probeGoMod(t, repo)
	goSum, err := os.ReadFile(filepath.Join(repo, "go.sum"))
	if err != nil {
		t.Fatal(err)
	}

	for _, c := range genCases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			dir := t.TempDir()
			probe := c.probe
			if probe == "" {
				probe = probeEmpty
			}
			writeFile(t, dir, "go.mod", goMod)
			writeFile(t, dir, "go.sum", string(goSum))
			writeFile(t, dir, "data.go", c.schema)
			writeFile(t, dir, "probe.go", probe)

			var warn bytes.Buffer
			err := generateTo(filepath.Join(dir, "data.go"), &warn)
			warnings := warn.String()
			t.Logf("generator diagnostics:\n%s", warnings)

			switch {
			case c.wantErr == "" && err != nil:
				t.Fatalf("generate: %v", err)
			case c.wantErr != "" && err == nil:
				t.Fatalf("generate succeeded; want error containing %q", c.wantErr)
			case c.wantErr != "" && !strings.Contains(err.Error(), c.wantErr):
				t.Fatalf("generate error = %v; want it to contain %q", err, c.wantErr)
			}
			for _, w := range c.wantWarn {
				if !strings.Contains(warnings, w) {
					t.Errorf("diagnostics missing %q; got:\n%s", w, warnings)
				}
			}
			for _, w := range c.wantNoWarn {
				if strings.Contains(warnings, w) {
					t.Errorf("diagnostics contain %q but should not; got:\n%s", w, warnings)
				}
			}
			if c.wantErr != "" {
				return // a failed run writes no decoder; nothing to compile
			}

			if c.probe == "" {
				if out, err := goTool(dir, "build", "-o", os.DevNull, "."); err != nil {
					t.Fatalf("generated package does not compile: %v\n%s\n--- generated ---\n%s",
						err, out, readFile(t, dir, "data_unmarshal.go"))
				}
				return
			}
			out, err := goTool(dir, "run", ".")
			if err != nil {
				t.Fatalf("go run: %v\n%s\n--- generated ---\n%s",
					err, out, readFile(t, dir, "data_unmarshal.go"))
			}
			if out != c.want {
				t.Errorf("probe output:\n%s\nwant:\n%s", out, c.want)
			}
		})
	}
}

// probeGoMod builds the throwaway module's go.mod from the repo's own: the
// generated code is compiled against THIS working tree (a filesystem replace, so
// no network and no version skew), and carrying the repo's requirements over
// keeps the probe building when they change — pkg/unstable's SIMD dispatch pulls
// in golang.org/x/sys, which module-graph pruning makes the probe declare too.
func probeGoMod(t *testing.T, repo string) string {
	t.Helper()
	b, err := os.ReadFile(filepath.Join(repo, "go.mod"))
	if err != nil {
		t.Fatal(err)
	}
	var out strings.Builder
	for _, line := range strings.Split(string(b), "\n") {
		if strings.HasPrefix(strings.TrimSpace(line), "module ") {
			out.WriteString("module lightningprobe\n")
			continue
		}
		out.WriteString(line)
		out.WriteByte('\n')
	}
	fmt.Fprintf(&out, "\nrequire %s v0.0.0\n\nreplace %s => %s\n", lightningMod, lightningMod, repo)
	return out.String()
}

// lightningMod is this module's path, the one the generated code imports
// pkg/unstable from.
const lightningMod = "github.com/JohanLindvall/lightning"

func writeFile(t *testing.T, dir, name, content string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func readFile(t *testing.T, dir, name string) string {
	t.Helper()
	b, err := os.ReadFile(filepath.Join(dir, name))
	if err != nil {
		return "<" + err.Error() + ">"
	}
	return string(b)
}

// goTool runs the go command in dir with the network switched off: everything
// the probe needs is either in this working tree or already in the module cache
// (the repo builds), so a proxy fetch would mean the probe module is wrong, and
// failing loudly beats hanging.
func goTool(dir string, args ...string) (string, error) {
	cmd := exec.Command("go", args...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "GOFLAGS=-mod=mod", "GOPROXY=off")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return stdout.String() + stderr.String(), err
	}
	return stdout.String(), nil
}
