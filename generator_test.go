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
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"slices"
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
		want:       tagProbeWant(),
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
		// H2a. The collect loop switched on ts.Type and never looked at
		// ts.TypeParams, so a generic declaration was collected like any other
		// root: generator exits 0, then "cannot use generic type Root[T any]
		// without instantiation". Warned and skipped rather than failed — a
		// schema file may hold a generic helper that was never meant as a root.
		// Here it is the only type, so the "nothing to generate" error fires,
		// which is also the assertion that the skip really happened.
		name: "generic_struct_root_is_skipped",
		schema: `package main

type Root[T any] struct {
	A int "json:\"a\""
}
`,
		wantErr:  "no top-level struct, slice or map types found",
		wantWarn: []string{"generic type Root", "type parameters"},
	},
	{
		name: "generic_slice_root_is_skipped",
		schema: `package main

type Root[T any] []int
`,
		wantErr:  "no top-level struct, slice or map types found",
		wantWarn: []string{"generic type Root", "type parameters"},
	},
	{
		// H2b. ts.Assign went unread too, so an alias TypeSpec was collected as
		// a defined type and got a method: "invalid receiver type *Root".
		name: "alias_to_struct_root_is_skipped",
		schema: `package main

type Root = struct {
	A int "json:\"a\""
}
`,
		wantErr:  "no top-level struct, slice or map types found",
		wantWarn: []string{"type alias Root", "defined type"},
	},
	{
		name: "alias_to_slice_root_is_skipped",
		schema: `package main

type Root = []int
`,
		wantErr:  "no top-level struct, slice or map types found",
		wantWarn: []string{"type alias Root", "defined type"},
	},
	{
		// The skips are per declaration, not per file: a generic helper and a
		// compat alias alongside a real root leave the root generating (and
		// compiling) exactly as it would alone.
		name: "generic_and_alias_beside_a_real_root",
		schema: `package main

type Pair[T any] struct {
	A T "json:\"a\""
}

type Legacy = Root

type Root struct {
	A int "json:\"a\""
}
`,
		wantWarn: []string{"generic type Pair", "type alias Legacy"},
	},
	{
		// Skipping an alias must not cost it the one thing it CAN do. `type
		// Legacy = struct{...}` is not a defined type, so no method attaches to
		// it — but a decoder taking a *Legacy is legal, and such an alias
		// already decoded as a field type; the collect loop keeps it resolvable
		// (registered, but not a root) so this generates the same decoder it
		// always did.
		name: "alias_to_struct_still_decodes_as_a_field",
		schema: `package main

type Legacy = struct {
	A int "json:\"a\""
}

type Root struct {
	L  Legacy   "json:\"l\""
	Ls []Legacy "json:\"ls\""
}
`,
		probe: `package main

import "fmt"

func main() {
	var v Root
	if err := v.UnmarshalJSON([]byte("{\"l\":{\"a\":1},\"ls\":[{\"a\":2},{\"a\":3}]}")); err != nil {
		panic(err)
	}
	fmt.Printf("%d %d %d\n", v.L.A, v.Ls[0].A, v.Ls[1].A)
}
`,
		want:     "1 2 3\n",
		wantWarn: []string{"type alias Legacy"},
	},
	{
		// H2c. Every *ast.InterfaceType routed to the any decoder, so a field
		// whose type has a method assigned an `any` to it: exits 0, does not
		// compile. The test is on the interface's CONTENT — interface{ any } is
		// still `any` and must keep working (the case below).
		name: "interface_with_method_is_rejected",
		schema: `package main

type Root struct {
	E interface{ Foo() } "json:\"e\""
}
`,
		wantErr: "unsupported type interface{ Foo() } for v.E",
	},
	{
		name: "interface_with_embedded_name_is_rejected",
		schema: `package main

type Root struct {
	E interface{ error } "json:\"e\""
}
`,
		wantErr: "unsupported type interface{ error } for v.E",
	},
	{
		// The spellings of `any` that must keep decoding, in every position the
		// walk reaches an interface through: a field, a slice element, a map
		// value, a pointee and a lax field.
		name: "empty_interface_spellings_decode",
		schema: `package main

type Root struct {
	A any                       "json:\"a\""
	B interface{}               "json:\"b\""
	C interface{ any }          "json:\"c\""
	D interface{ interface{} }  "json:\"d\""
	S []interface{}             "json:\"s\""
	M map[string]interface{}    "json:\"m\""
	P *interface{}              "json:\"p\""
	L interface{}               "json:\"l,lax\""
}
`,
		probe: `package main

import "fmt"

const doc = "{\"a\":1,\"b\":\"x\",\"c\":true,\"d\":null,\"s\":[1,2],\"m\":{\"k\":3},\"p\":4,\"l\":[5]}"

func main() {
	var v Root
	if err := v.UnmarshalJSON([]byte(doc)); err != nil {
		panic(err)
	}
	fmt.Printf("%v %v %v %v %v %v %v %v\n", v.A, v.B, v.C, v.D, v.S, v.M, *v.P, v.L)
}
`,
		want: "1 x true <nil> [1 2] map[k:3] 4 [5]\n",
	},
	{
		// H2d. The generated bodies declare their own parameters and locals
		// around the schema's type names, so a type named `data` was captured by
		// the []byte parameter ("data (parameter) is not a type", "cannot use
		// new(data) (value of type *[]byte)") — again exiting 0 on a package
		// that does not compile. Reported as an error naming the type instead;
		// renaming the generated locals would have churned every committed
		// generated file for nothing.
		name: "schema_type_named_like_a_generated_local",
		schema: `package main

type data struct {
	X int "json:\"x\""
}

type Root struct {
	Items []data "json:\"items\""
	P     *data  "json:\"p\""
}
`,
		wantErr: `type name "data" collides with an identifier used by the generated code`,
	},
	{
		// Rejected even where this particular schema would have compiled: a
		// lone `data` root happens to name the type only in signatures, where
		// the parameter is not yet in scope. The set is a property of the name,
		// not of how far a given schema gets — which is what makes it
		// maintainable.
		name: "schema_type_named_data_alone",
		schema: `package main

type data struct {
	X int "json:\"x\""
}
`,
		wantErr: `type name "data" collides with an identifier used by the generated code`,
	},
	{
		// The import-name half: a package-level `unstable` and the scanner
		// import cannot coexist ("unstable already declared through import").
		name: "schema_type_named_like_the_scanner_import",
		schema: `package main

type unstable struct {
	X int "json:\"x\""
}

type Root struct {
	U unstable "json:\"u\""
}
`,
		wantErr: `type name "unstable" collides with an identifier used by the generated code`,
	},
	{
		// And the predeclared-identifier half: a package-level `max` shadows the
		// builtin the un-presized slice's capacity hint calls. The element nests
		// a slice so slicePresize declines to count it, which is what selects
		// that hint.
		name: "schema_type_named_like_a_builtin",
		schema: `package main

type max struct {
	Xs []int "json:\"xs\""
}

type Root struct {
	Items []max "json:\"items\""
}
`,
		wantErr: `type name "max" collides with an identifier used by the generated code`,
	},
	{
		// M4. `referenced` marks both members of a mutually recursive pair, so
		// both were skipped as "nested in another type" and the pair got NO
		// method and no decoder at all — silently, exit 0. The old rescue (emit
		// everything when every type is referenced) only fired when the cycle
		// was the whole file, so one unrelated type alongside brought the hole
		// back. The fixpoint promotes the whole strongly-connected component of
		// a cycle no emitted type enters, so BOTH members are decodable however
		// the file is ordered; the probe checks that, decodes, and checks the
		// depth bound still applies to the cycle.
		name: "mutually_recursive_pair_beside_another_type",
		schema: `package main

type Other struct {
	X int "json:\"x\""
}

type MutA struct {
	Name string "json:\"name\""
	B    *MutB  "json:\"b\""
}

type MutB struct {
	Count int   "json:\"count\""
	A     *MutA "json:\"a\""
}
`,
		probe: `package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/JohanLindvall/lightning/pkg/unstable"
)

func main() {
	var o Other
	if err := o.UnmarshalJSON([]byte("{\"x\":7}")); err != nil {
		panic(err)
	}
	var v MutA
	if err := v.UnmarshalJSON([]byte("{\"name\":\"a\",\"b\":{\"count\":2,\"a\":{\"name\":\"deep\"}}}")); err != nil {
		panic(err)
	}
	// Both cycle members are decodable, not just the one declared first.
	var w MutB
	if err := w.UnmarshalJSON([]byte("{\"count\":5,\"a\":{\"name\":\"z\"}}")); err != nil {
		panic(err)
	}
	fmt.Printf("%d %s %d %s %d %s\n", o.X, v.Name, v.B.Count, v.B.A.Name, w.Count, w.A.Name)

	// The cycle is depth-bounded, not a stack overflow. The nesting has to
	// alternate b/a to keep descending — a repeated "b" would be an unknown key
	// inside MutB and get skipped, never recursing.
	n := unstable.MaxDepth/2 + 10
	deep := strings.Repeat("{\"b\":{\"a\":", n) + "null" + strings.Repeat("}}", n)
	var d MutA
	fmt.Println(errors.Is(d.UnmarshalJSON([]byte(deep)), unstable.ErrMaxDepth))
}
`,
		want: "7 a 2 deep 5 z\ntrue\n",
	},
	{
		// The degenerate file the old special case existed for: a cycle with
		// nothing else in the file. It keeps working through the general rule,
		// and keeps its per-member methods — promoting the SCC rather than one
		// member is what makes the general rule a superset of the special case
		// it replaced, instead of a silent downgrade on upgrade.
		name: "mutually_recursive_pair_alone",
		schema: `package main

type MutA struct {
	Name string "json:\"name\""
	B    *MutB  "json:\"b\""
}

type MutB struct {
	Count int   "json:\"count\""
	A     *MutA "json:\"a\""
}
`,
		probe: `package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var v MutA
	if err := v.UnmarshalJSON([]byte("{\"name\":\"a\",\"b\":{\"count\":2}}")); err != nil {
		panic(err)
	}
	_, aHas := any(&MutA{}).(json.Unmarshaler)
	_, bHas := any(&MutB{}).(json.Unmarshaler)
	fmt.Printf("%s %d %v %v\n", v.Name, v.B.Count, aHas, bHas)
}
`,
		want: "a 2 true true\n",
	},
	{
		// A record hanging off a cycle stays nested. This is the case the SCC
		// rule has to get right BY CONSTRUCTION: Leaf is reachable from the
		// cycle but does not reach back into it, so it is outside the component
		// and keeps only an internal decoder — where promoting the whole
		// uncovered set would have handed it a method and pulled a
		// `type leafStd Leaf` reflection baseline into the generated decoder.
		name: "type_below_a_cycle_stays_nested",
		schema: `package main

type CycA struct {
	B    *CycB "json:\"b\""
	Leaf Leaf  "json:\"leaf\""
}

type CycB struct {
	A *CycA "json:\"a\""
}

type Leaf struct {
	N int "json:\"n\""
}
`,
		probe: `package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var v CycA
	if err := v.UnmarshalJSON([]byte("{\"leaf\":{\"n\":3},\"b\":{\"a\":{\"leaf\":{\"n\":4}}}}")); err != nil {
		panic(err)
	}
	// Both cycle members get the method; Leaf keeps an internal decoder, which
	// is what preserves the reflection baselines — a json.Unmarshaler on it
	// would make encoding/json delegate to lightning and compare it with
	// itself.
	_, aHas := any(&CycA{}).(json.Unmarshaler)
	_, bHas := any(&CycB{}).(json.Unmarshaler)
	_, leafHas := any(&Leaf{}).(json.Unmarshaler)
	fmt.Printf("%d %d %v %v %v\n", v.Leaf.N, v.B.A.Leaf.N, aHas, bHas, leafHas)
}
`,
		want: "3 4 true true false\n",
	},
	{
		// The same shape with the hanging record declared FIRST, which is the
		// order that broke: with no entry type in the file nothing is covered
		// on the first round, so scanning for the first uncovered type in
		// source order picked Leaf — whose component is itself — and handed it
		// a method, where declaring it last left it covered and nested.
		// Requiring the candidate's component to be a SOURCE of the uncovered
		// subgraph is what makes the two orders agree. Keep both cases: they
		// differ only in declaration order and must produce the same answer.
		name: "type_below_a_cycle_stays_nested_declared_first",
		schema: `package main

type Leaf struct {
	N int "json:\"n\""
}

type CycA struct {
	B    *CycB "json:\"b\""
	Leaf Leaf  "json:\"leaf\""
}

type CycB struct {
	A *CycA "json:\"a\""
}
`,
		probe: `package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var v CycA
	if err := v.UnmarshalJSON([]byte("{\"leaf\":{\"n\":3},\"b\":{\"a\":{\"leaf\":{\"n\":4}}}}")); err != nil {
		panic(err)
	}
	_, aHas := any(&CycA{}).(json.Unmarshaler)
	_, bHas := any(&CycB{}).(json.Unmarshaler)
	_, leafHas := any(&Leaf{}).(json.Unmarshaler)
	fmt.Printf("%d %d %v %v %v\n", v.Leaf.N, v.B.A.Leaf.N, aHas, bHas, leafHas)
}
`,
		want: "3 4 true true false\n",
	},
	{
		// A cycle longer than two: every member is mutually reachable with
		// every other, so the component is all three and each gets a method.
		name: "three_member_cycle",
		schema: `package main

type CycA struct {
	N int   "json:\"n\""
	B *CycB "json:\"b\""
}

type CycB struct {
	N int   "json:\"n\""
	C *CycC "json:\"c\""
}

type CycC struct {
	N int   "json:\"n\""
	A *CycA "json:\"a\""
}
`,
		probe: `package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var v CycB
	if err := v.UnmarshalJSON([]byte("{\"n\":1,\"c\":{\"n\":2,\"a\":{\"n\":3}}}")); err != nil {
		panic(err)
	}
	_, a := any(&CycA{}).(json.Unmarshaler)
	_, b := any(&CycB{}).(json.Unmarshaler)
	_, c := any(&CycC{}).(json.Unmarshaler)
	fmt.Printf("%d %d %d %v %v %v\n", v.N, v.C.N, v.C.A.N, a, b, c)
}
`,
		want: "1 2 3 true true true\n",
	},
	{
		// Two disjoint entryless cycles in one file: the fixpoint has to
		// promote a component, recompute, and find the second one still
		// uncovered. All four members end up decodable.
		name: "two_disjoint_cycles",
		schema: `package main

type OneA struct {
	B *OneB "json:\"b\""
}

type OneB struct {
	A *OneA "json:\"a\""
}

type TwoA struct {
	B *TwoB "json:\"b\""
}

type TwoB struct {
	A *TwoA "json:\"a\""
}
`,
		probe: `package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var v TwoB
	if err := v.UnmarshalJSON([]byte("{\"a\":{\"b\":{}}}")); err != nil {
		panic(err)
	}
	_, a1 := any(&OneA{}).(json.Unmarshaler)
	_, b1 := any(&OneB{}).(json.Unmarshaler)
	_, a2 := any(&TwoA{}).(json.Unmarshaler)
	_, b2 := any(&TwoB{}).(json.Unmarshaler)
	fmt.Printf("%v %v %v %v %v\n", v.A.B != nil, a1, b1, a2, b2)
}
`,
		want: "true true true true true\n",
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

// tagProbeStd mirrors invalid_json_tag_names' Root for the three names whose
// treatment by encoding/json depends on the toolchain, so the probe's stdlib
// columns are computed by the encoding/json the test runs against rather than
// pinned to one version: through Go 1.26 an invalid name made the stdlib
// discard the whole tag and key the field by its Go name; the json/v2-backed
// decoder of Go 1.27 ignores a field whose name holds a quote or backslash
// (neither key matches it) and accepts the newline name as written. The
// lightning columns are fixed — it matches every name as written — and the
// divergence itself is still required: if the stdlib ever matched all three
// tag names too, the want string says so and the case fails.
type tagProbeStd struct {
	Q  int `json:"a\"b"` //nolint:staticcheck // deliberately a name encoding/json's tag rule rejects
	BS int `json:"a\\b"` //nolint:staticcheck // deliberately a name encoding/json's tag rule rejects
	NL int `json:"a\nb"`
}

func tagProbeWant() string {
	var s, s2 tagProbeStd
	_ = json.Unmarshal([]byte("{\"a\\\"b\":1,\"a\\\\b\":2,\"a\\nb\":3}"), &s)
	_ = json.Unmarshal([]byte("{\"Q\":1,\"BS\":2,\"NL\":3}"), &s2)
	if s.Q == 1 && s.BS == 2 && s.NL == 3 {
		return "premise broken: encoding/json now matches every invalid tag name as written; the divergence has closed\n"
	}
	return fmt.Sprintf("tag keys:   lightning=1,2,3 stdlib=%d,%d,%d\nfield keys: lightning=0,0,0 stdlib=%d,%d,%d\n",
		s.Q, s.BS, s.NL, s2.Q, s2.BS, s2.NL)
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

// unmarshalerRe finds the receiver type of each generated UnmarshalJSON method,
// which is the observable form of entryTypes' decision: a type with a method is
// a root, a type with only an internal decoder is nested inside one.
var unmarshalerRe = regexp.MustCompile(`func \(v \*(\w+)\) UnmarshalJSON`)

// TestEntryTypesOrderIndependent is a property test over declaration order: the
// set of types that receive an UnmarshalJSON must not depend on the order the
// types are declared in.
//
// It exists because the table above cannot see this class of bug. Each case
// there fixes one order, so a rule that picks its answer by source position
// passes every one of them and still moves a method when a user reorders two
// declarations. That is exactly what happened: an earlier fixpoint took the
// first UNCOVERED type in source order as its promotion candidate, which is
// correct whenever the file has an entry type (the hanger-on is covered before
// the scan reaches it) and wrong when it has none — a record hanging off a
// cycle, declared before that cycle, was promoted and given a method, pulling a
// `type recordStd Record` reflection baseline into the generated decoder.
//
// Only the generator runs here, not the compiler: the property is about which
// methods are emitted, and skipping the per-permutation `go build` is what
// makes exhausting 5! orders cheap.
func TestEntryTypesOrderIndependent(t *testing.T) {
	t.Parallel()

	cycleA := `type CycA struct {
	B *CycB "json:\"b\""
}`
	cycleB := `type CycB struct {
	A    *CycA "json:\"a\""
	Leaf *Leaf "json:\"leaf\""
}`
	leaf := `type Leaf struct {
	N int "json:\"n\""
}`
	oneA := `type OneA struct {
	B *OneB "json:\"b\""
}`
	oneB := `type OneB struct {
	A    *OneA "json:\"a\""
	Hang *Hang "json:\"hang\""
}`
	twoA := `type TwoA struct {
	B *TwoB "json:\"b\""
}`
	twoB := `type TwoB struct {
	A *TwoA "json:\"a\""
}`
	hang := `type Hang struct {
	N int "json:\"n\""
}`

	cases := []struct {
		name  string
		decls []string
		// want is the set of types that must get an UnmarshalJSON, in every
		// declaration order. The hangers-on (Leaf, Hang) are deliberately
		// absent: reachable from a cycle, never reaching back into it, so they
		// belong to no source component and stay nested.
		want []string
	}{
		{
			name:  "cycle_with_hanging_record",
			decls: []string{cycleA, cycleB, leaf},
			want:  []string{"CycA", "CycB"},
		},
		{
			name:  "two_disjoint_cycles_with_hanging_record",
			decls: []string{oneA, oneB, twoA, twoB, hang},
			want:  []string{"OneA", "OneB", "TwoA", "TwoB"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			dir := t.TempDir()
			perms := permutations(c.decls)
			for _, perm := range perms {
				schema := "package main\n\n" + strings.Join(perm, "\n\n") + "\n"
				writeFile(t, dir, "data.go", schema)
				var warn bytes.Buffer
				if err := generateTo(filepath.Join(dir, "data.go"), &warn); err != nil {
					t.Fatalf("generate: %v\n--- schema ---\n%s", err, schema)
				}
				var got []string
				for _, m := range unmarshalerRe.FindAllStringSubmatch(readFile(t, dir, "data_unmarshal.go"), -1) {
					got = append(got, m[1])
				}
				slices.Sort(got)
				if !slices.Equal(got, c.want) {
					t.Fatalf("declaration order changed which types get UnmarshalJSON:\ngot  %v\nwant %v\n--- schema ---\n%s",
						got, c.want, schema)
				}
			}
			t.Logf("%d declaration orders, all yielding %v", len(perms), c.want)
		})
	}
}

// permutations returns every ordering of in. The inputs are a handful of type
// declarations, so the factorial blowup is bounded by the caller's shapes (6 and
// 120 orders); anything larger should sample instead.
func permutations(in []string) [][]string {
	if len(in) <= 1 {
		return [][]string{slices.Clone(in)}
	}
	var out [][]string
	for i := range in {
		rest := make([]string, 0, len(in)-1)
		rest = append(rest, in[:i]...)
		rest = append(rest, in[i+1:]...)
		for _, p := range permutations(rest) {
			out = append(out, append([]string{in[i]}, p...))
		}
	}
	return out
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
