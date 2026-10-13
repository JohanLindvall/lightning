package main

import "fmt"

// Reusable forms of Overcast's post-generation edits: field readers, [] versus
// null, and go vet's rejection of self-assignments in map[string]any decoders.
var genOverrideCases = append([]genCase{
	{
		name: "custom_reader_fields",
		schema: `package main
type Root struct {
	*Promoted
	Plain string "json:\"plain\""
	Date string "json:\"date|longAlternateDateKey,reader=readText\""
	Stamp Stamp "json:\"stamp,reader=readStamp\""
	Ptr *string "json:\"ptr,reader=readPointer\""
	Values []int "json:\"values,reader=readInts\""
	Own Own "json:\"own,reader=readOwn\""
	Soft string "json:\"soft,reader=readText,lax\""
	Wrapped string "json:\"wrapped,reader=readText,unwrap\""
	Both string "json:\"both,reader=readText,lax,unwrap\""
	Rows []Row "json:\"rows\""
	Mapped map[string]Row "json:\"mapped\""
}
type Stamp string
type Own struct { Text string }
func (*Own) UnmarshalJSON([]byte) error { panic("custom reader must win") }
`,
		extra: map[string]string{
			"rows.go": `package main
type Row struct { Date string "json:\"date,reader=readText\"" }
type Promoted struct { Extra string "json:\"extra,reader=readText\"" }
`,
			"readers.go": `package main
import (
	"errors"
	"github.com/JohanLindvall/lightning/pkg/unstable"
)
var errReader = errors.New("reader failed")
var calls int
func readText(data []byte, i int) (string, int, error) {
	calls++
	s, end, err := unstable.ReadStringOrNull(data, i)
	if err != nil { return "bad", end, err }
	if s == "fail" { return "bad", end, errReader }
	if data[i] == 'n' { return "custom:null", end, nil }
	return "custom:" + s, end, nil
}
func readStamp(data []byte, i int) (Stamp, int, error) {
	s, end, err := readText(data, i)
	return Stamp(s), end, err
}
func readPointer(data []byte, i int) (*string, int, error) {
	s, end, err := readText(data, i)
	return &s, end, err
}
func readInts(data []byte, i int) ([]int, int, error) {
	var values []int
	end, err := unstable.DecodeIntSlice(&values, data, i)
	return values, end, err
}
func readOwn(data []byte, i int) (Own, int, error) {
	s, end, err := readText(data, i)
	return Own{Text:s}, end, err
}
`,
		},
		probe: `package main
import (
	"errors"
	"fmt"
	"reflect"
)
func main() {
	raw := []byte("{\"extra\":\"e\",\"plain\":\"p\",\"date\":\"first\",\"longAlternateDateKey\": \"2026-10-13\",\"stamp\":\"s\",\"ptr\":null,\"values\":[1,2],\"own\":\"o\",\"wrapped\":\"\\\"w\\\"\",\"rows\":[{\"date\":\"r\"}],\"mapped\":{\"k\":{\"date\":\"m\"}}}")
	var v Root
	if err := v.UnmarshalJSON(raw); err != nil { panic(err) }
	for i := range raw { raw[i] = 'x' }
	if v.Extra != "custom:e" || v.Plain != "p" || v.Date != "custom:2026-10-13" || v.Stamp != "custom:s" || *v.Ptr != "custom:null" || !reflect.DeepEqual(v.Values, []int{1,2}) || v.Own.Text != "custom:o" || v.Wrapped != "custom:w" || v.Rows[0].Date != "custom:r" || v.Mapped["k"].Date != "custom:m" || calls != 9 {
		panic(fmt.Sprintf("unexpected fields: %+v, calls=%d", v, calls))
	}
	if err := v.UnmarshalJSON([]byte("{}")); err != nil || calls != 9 { panic("absent reader called") }
	if err := v.UnmarshalJSON([]byte("{\"date\":null}")); err != nil || v.Date != "custom:null" { panic("reader null lost") }
	v.Date = "kept"
	end, err := lightningLightningprobeRootdecodeRoot(&v, []byte("{\"date\":\"fail\"}"), 0)
	if !errors.Is(err, errReader) || end != 14 || v.Date != "kept" { panic(fmt.Sprintf("error contract: %d %v %q", end, err, v.Date)) }
	v.Soft = "kept"
	if err := v.UnmarshalJSON([]byte("{\"soft\":42,\"plain\":\"after\"}")); err != nil || v.Soft != "kept" || v.Plain != "after" { panic("lax mismatch") }
	if err := v.UnmarshalJSON([]byte("{\"soft\":null}")); err != nil || v.Soft != "custom:null" { panic("lax null") }
	if err := v.UnmarshalJSON([]byte("{\"both\":\"42\"}")); err != nil || v.Both != "" { panic("lax unwrap") }
	for _, doc := range []string{
		"{\"date\":42}", "{\"date\":\"bad\\q\"}", "{\"soft\":[1,]}",
		"{\"wrapped\":\"\\\"ok\\\" false\"}", "{\"both\":\"42 false\"}",
	} {
		if err := v.UnmarshalJSON([]byte(doc)); err == nil { panic("accepted " + doc) }
	}
	fmt.Println("custom readers preserve values, offsets, errors and tag composition")
}
`,
		want: "custom readers preserve values, offsets, errors and tag composition\n",
	},
	{
		name: "custom_reader_imported_and_embedded",
		schema: `package main
import (
	helper "lightningprobe/helper"
	"github.com/JohanLindvall/lightning/pkg/unstable"
)
var _ = helper.Read
var _ = unstable.ReadStringOrNull
type Root struct {
	Embedded "json:\",reader=readEmbedded\""
	Text string "json:\"text,reader=helper.Read\""
	Raw string "json:\"raw,reader=unstable.ReadStringOrNull\""
}
type Embedded struct { Value string }
func readEmbedded(data []byte, i int) (Embedded, int, error) {
	s, end, err := helper.Read(data, i)
	return Embedded{Value: s}, end, err
}
`,
		extra: map[string]string{
			"helper/read.go": `package helper
import "github.com/JohanLindvall/lightning/pkg/unstable"
func Read(data []byte, i int) (string, int, error) {
	s, end, err := unstable.ReadStringOrNull(data, i)
	return "imported:"+s, end, err
}
`,
		},
		probe: `package main
import "fmt"
func main() {
	var v Root
	err := v.UnmarshalJSON([]byte("{\"Embedded\":\"e\",\"text\":\"t\",\"raw\":\"r\"}"))
	fmt.Println(err, v.Value, v.Text, v.Raw)
}
`,
		want: "<nil> imported:e imported:t r\n",
	},
	{
		name: "empty_slice_directive_all_decoders",
		schema: `package main
import "encoding/json"
//lightning:emptyslice
type Root struct {
	Strings []string "json:\"strings\""
	Bools []bool "json:\"bools\""
	Small []float32 "json:\"small\""
	Floats []float64 "json:\"floats\""
	Ints []int "json:\"ints\""
	Uints []uint64 "json:\"uints\""
	Bytes []byte "json:\"bytes\""
	Points [][2]float64 "json:\"points\""
	Rows []Row "json:\"rows\""
	Nested [][]string "json:\"nested\""
	Map map[string][]int "json:\"map\""
	Ptr *[]string "json:\"ptr\""
	Array [2][]string "json:\"array\""
	Named Items "json:\"named\""
	NamedInts Numbers "json:\"namedInts\""
	Lax []int "json:\"lax,lax\""
	Wrapped []float64 "json:\"wrapped,unwrap\""
	Dynamic any "json:\"dynamic\""
	Raw json.RawMessage "json:\"raw\""
}
type Row struct { Children []Row "json:\"children\"" }
type Plain struct { Rows []Row "json:\"rows\""; Ints []int "json:\"ints\"" }
//lightning:emptyslice
type Items []Row
//lightning:emptyslice
//lightning:arena
type Numbers []int
//lightning:emptyslice
//lightning:arena
//lightning:compact
//lightning:strict
type Arena Root
//lightning:emptyslice
//lightning:destructive
type Destructive Root
`,
		probe: `package main
import (
	"fmt"
	"reflect"
)
func must(err error) { if err != nil { panic(err) } }
func main() {
	var v Root
	fields := []string{"strings","bools","small","floats","ints","uints","bytes","points","rows","nested","named","namedInts","lax"}
	for _, field := range fields { must(v.UnmarshalJSON([]byte("{\""+field+"\":[]}"))) }
	rv := reflect.ValueOf(v)
	for _, i := range []int{0,1,2,3,4,5,6,7,8,9,13,14,15} {
		if rv.Field(i).IsNil() || rv.Field(i).Len() != 0 { panic(fmt.Sprintf("nil empty field %s", rv.Type().Field(i).Name)) }
	}
	if v.Map != nil || v.Ptr != nil || v.Wrapped != nil { panic("absent changed") }
	must(v.UnmarshalJSON([]byte("{\"rows\":[{\"children\":[]}],\"nested\":[[],null],\"map\":{\"e\":[],\"n\":null},\"ptr\":[],\"array\":[[],null],\"wrapped\":\"[]\",\"dynamic\":[],\"raw\":[]}")))
	if v.Rows[0].Children == nil || v.Nested[0] == nil || v.Nested[1] != nil || v.Map["e"] == nil || v.Map["n"] != nil || *v.Ptr == nil || v.Array[0] == nil || v.Array[1] != nil || v.Wrapped == nil || v.Dynamic.([]any) == nil || string(v.Raw) != "[]" { panic("nested empty slices") }
	for _, doc := range []string{"{\"strings\":[\"a\"],\"ints\":[1,2]}","{\"strings\":[],\"ints\":[]}"} { must(v.UnmarshalJSON([]byte(doc))) }
	if v.Strings == nil || cap(v.Strings) == 0 || v.Ints == nil || cap(v.Ints) == 0 { panic("lost reused backing") }
	stringsCap, intsCap := cap(v.Strings), cap(v.Ints)
	must(v.UnmarshalJSON([]byte("{\"strings\":[],\"ints\":[]}")))
	if cap(v.Strings) != stringsCap || cap(v.Ints) != intsCap { panic("empty reallocated") }
	must(v.UnmarshalJSON([]byte("{\"strings\":[],\"strings\":null,\"ints\":null,\"ints\":[]}")))
	if v.Strings != nil || v.Ints == nil { panic("duplicate key order") }
	for _, field := range fields { must(v.UnmarshalJSON([]byte("{\""+field+"\":null}"))) }
	rv = reflect.ValueOf(v)
	for _, i := range []int{0,1,2,3,4,5,6,7,8,9,13,14,15} { if !rv.Field(i).IsNil() { panic("null materialized") } }
	for _, doc := range []string{"{\"strings\":[}","{\"strings\":[\"x\",]}","{\"ints\":[1,]}","{\"points\":[[1,2],]}"} {
		if v.UnmarshalJSON([]byte(doc)) == nil { panic("accepted " + doc) }
	}
	var items Items
	must(items.UnmarshalJSON([]byte("[]")))
	if items == nil { panic("named root") }
	must(items.UnmarshalJSON([]byte("null")))
	if items != nil { panic("named null") }
	var nums Numbers
	must(nums.UnmarshalJSON([]byte("[]")))
	if nums == nil { panic("arena root") }
	must(nums.UnmarshalJSON([]byte("[1,2]")))
	if !reflect.DeepEqual(nums, Numbers{1,2}) { panic("arena values") }
	var a Arena
	must(a.UnmarshalJSON([]byte("{\"ints\":[],\"floats\":[],\"bytes\":[],\"points\":[],\"namedInts\":[],\"lax\":[],\"rows\":[{\"children\":[]}]}")))
	if a.Ints == nil || a.Floats == nil || a.Bytes == nil || a.Points == nil || a.NamedInts == nil || a.Lax == nil || a.Rows[0].Children == nil { panic("arena wrappers") }
	var d Destructive
	must(d.UnmarshalJSON([]byte("{\"strings\":[],\"ints\":[]}")))
	if d.Strings == nil || d.Ints == nil { panic("destructive variant") }
	var p Plain
	must(p.UnmarshalJSON([]byte("{\"rows\":[{\"children\":[]}],\"ints\":[]}")))
	if p.Rows[0].Children != nil || p.Ints != nil { panic("directive leaked") }
	fmt.Println("empty slices preserve null, absence, reuse and variant isolation")
}
`,
		want: "empty slices preserve null, absence, reuse and variant isolation\n",
	},
	{
		name: "dynamic_map_values_pass_vet",
		schema: `package main
type Root struct {
	Values map[string]any "json:\"values\""
	Nested map[string]map[string]interface{} "json:\"nested\""
}
//lightning:compact
type Compact Root
type Dynamic map[string]any
`,
		probe: `package main
import (
	"encoding/json"
	"fmt"
	"reflect"
)
type standard Root
func main() {
	doc := []byte("{\"values\":{\"str\":\"x\",\"num\":42,\"bool\":true,\"null\":null,\"array\":[1,2],\"object\":{}},\"nested\":{\"k\":{\"v\":3}}}")
	var got Root
	var compact Compact
	var want standard
	if err := got.UnmarshalJSON(doc); err != nil { panic(err) }
	if err := compact.UnmarshalJSON(doc); err != nil { panic(err) }
	if err := json.Unmarshal(doc, &want); err != nil { panic(err) }
	fmt.Println(reflect.DeepEqual(got, Root(want)), reflect.DeepEqual(got, Root(compact)))
}
`,
		want: "true true\n",
		vet:  true,
	},
}, invalidReaderCases()...)

func invalidReaderCases() []genCase {
	var cases []genCase
	for i, tc := range []struct{ option, err string }{
		{"reader=", "invalid reader"},
		{"reader", "invalid reader"},
		{"reader=readText,reader=other", "multiple reader options"},
		{"reader=readText()", "invalid reader"},
		{"reader=missing.Read", "invalid reader"},
		{"reader=unstable.ReadStringOrNull", "invalid reader"},
		{"reader=data", "invalid reader"},
		{"reader=_", "invalid reader"},
		{"reader=readText,nocopy", "cannot be combined"},
		{"reader=readText,number", "cannot be combined"},
	} {
		cases = append(cases, genCase{
			name:    fmt.Sprintf("invalid_reader_%d", i),
			schema:  fmt.Sprintf("package main\ntype Root struct { Value string %q }\n", `json:"value,`+tc.option+`"`),
			wantErr: tc.err,
		})
	}
	return cases
}
