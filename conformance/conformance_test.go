package conformance

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

// TestConformance decodes test.json with the generated UnmarshalJSON and checks
// every field, exercising all supported types and the nocopy/lax/unwrap/pipe/skip
// tag options. Expected values are spelled out (not compared against
// encoding/json) because several options — unwrap, lax, pipe alternate names —
// are lightning-specific semantics encoding/json does not share.
func TestConformance(t *testing.T) {
	data, err := os.ReadFile("test.json")
	if err != nil {
		t.Fatalf("read test.json: %v", err)
	}

	var d Doc
	if err := d.UnmarshalJSON(data); err != nil {
		t.Fatalf("UnmarshalJSON: %v", err)
	}

	wantTime := time.Date(2023, 6, 15, 12, 30, 45, 0, time.UTC)
	pnum := json.Number("1.5e10")
	pi := 99

	checks := []struct {
		name string
		got  any
		want any
	}{
		{"Str", d.Str, "hello"},
		{"StrNoCopy", d.StrNoCopy, "world"},
		{"StrUnicode", d.StrUnicode, "café\n\t\"x\""},
		{"Bool", d.Bool, true},

		{"I", d.I, int(-42)},
		{"I8", d.I8, int8(-128)},
		{"I16", d.I16, int16(32000)},
		{"I32", d.I32, int32(-2000000000)},
		{"I64", d.I64, int64(9223372036854775807)},
		{"R", d.R, rune(97)},

		{"U", d.U, uint(42)},
		{"U8", d.U8, uint8(255)},
		{"U16", d.U16, uint16(65535)},
		{"U32", d.U32, uint32(4000000000)},
		{"U64", d.U64, uint64(18446744073709551615)},
		{"Uptr", d.Uptr, uintptr(12345)},
		{"B", d.B, byte(200)},

		{"F32", d.F32, float32(3.14)},
		{"F64", d.F64, float64(2.718281828459045)},

		{"Num", d.Num, json.Number("12345678901234567890")},
		{"PNum", d.PNum, &pnum},

		{"Raw", string(d.Raw), `{"a":1,"b":[2,3]}`},
		{"RawNoCopy", string(d.RawNoCopy), `[1,2,3]`},

		{"Nested", d.Nested, Nested{Name: "n1", Count: 5}},
		{"Anon.X", d.Anon.X, 7},
		{"Anon.Y", d.Anon.Y, "why"},

		{"Ints", d.Ints, []int{1, 2, 3, 4}},
		{"Strs", d.Strs, []string{"a", "b", "c"}},
		{"Floats", d.Floats, []float64{1.5, 2.5, 3.5}},
		{"Grid", d.Grid, [][]int{{1, 2}, {3, 4, 5}}},
		{"Items", d.Items, []Nested{{Name: "i1", Count: 1}, {Name: "i2", Count: 2}}},
		{"PtrItems", d.PtrItems, []*Nested{{Name: "p1", Count: 1}, nil, {Name: "p2", Count: 2}}},
		{"Arr", d.Arr, [3]int{10, 20, 30}},
		{"UArr", d.UArr, [3]uint32{7, 8, 4000000000}},

		{"M", d.M, map[string]int{"x": 1, "y": 2}},
		{"MM", d.MM, map[string][]uint64{"p": {1, 2}, "q": {3}}},

		{"PI", d.PI, &pi},
		{"PN", d.PN, &Nested{Name: "pp", Count: 8}},

		{"Any", d.Any, float64(3.5)},
		{"Anys", d.Anys, []any{true, "x", float64(2)}},

		// "EdgeStatus" is the pipe alternate name for the Status field.
		{"Status", d.Status, 503},
		// json:"-" is never filled even though test.json has an "Ignored" key.
		{"Ignored", d.Ignored, 0},
		// unwrap: the embedded JSON string decodes into the nested struct.
		{"Embedded", d.Embedded, Nested{Name: "emb", Count: 3}},
		// lax: the string value for an int field is skipped, leaving zero.
		{"LaxSkip", d.LaxSkip, 0},
	}

	for _, c := range checks {
		if !reflect.DeepEqual(c.got, c.want) {
			t.Errorf("%s = %#v, want %#v", c.name, c.got, c.want)
		}
	}

	// Time fields compared with time.Equal (wall-clock + monotonic agnostic).
	if !d.Time.Equal(wantTime) {
		t.Errorf("Time = %v, want %v", d.Time, wantTime)
	}
	if !d.TimeLax.Equal(wantTime) {
		t.Errorf("TimeLax = %v, want %v", d.TimeLax, wantTime)
	}
}

// TestTrailingCommaRejected locks the rotated member/element loops: a trailing
// comma before a closing bracket is an error in every decoder shape — object,
// slice (generated and batched), fixed array, map, and the dynamic any path —
// exactly as encoding/json rejects it.
func TestTrailingCommaRejected(t *testing.T) {
	cases := []string{
		`{"str":"x",}`,                        // object member loop
		`{"ints":[1,2,]}`,                     // batched int slice
		`{"floats":[1.5,]}`,                   // batched float slice
		`{"strs":["a",]}`,                     // generated slice loop
		`{"items":[{"name":"a","count":1},]}`, // struct-element slice loop
		`{"arr":[10,20,30,]}`,                 // fixed-size array loop
		`{"uarr":[1,2,3,]}`,                   // batched uint fixed array
		`{"m":{"x":1,}}`,                      // map member loop
		`{"grid":[[1,2],]}`,                   // nested slice loop
		`{"anys":[1,]}`,                       // dynamic array
		`{"any":{"k":1,}}`,                    // dynamic object
	}
	for _, in := range cases {
		if json.Unmarshal([]byte(in), new(map[string]any)) == nil {
			t.Fatalf("stdlib accepted %q — test premise wrong", in)
		}
		var d Doc
		if err := d.UnmarshalJSON([]byte(in)); err == nil {
			t.Errorf("UnmarshalJSON(%q) accepted a trailing comma", in)
		}
	}
	// Empty containers are unaffected by the loop rotation.
	for _, in := range []string{`{}`, `{"ints":[]}`, `{"strs":[]}`, `{"m":{}}`, `{"anys":[]}`, `{"arr":[]}`} {
		var d Doc
		if err := d.UnmarshalJSON([]byte(in)); err != nil {
			t.Errorf("UnmarshalJSON(%q): %v", in, err)
		}
	}
}

// TestSliceRoot covers a named slice root type (array-root JSON): the generator
// emits UnmarshalJSON on PointList itself.
func TestSliceRoot(t *testing.T) {
	var got PointList
	if err := got.UnmarshalJSON([]byte(`[{"x":1,"y":2,"tag":"a"},{"x":3,"y":4,"tag":"b"}]`)); err != nil {
		t.Fatal(err)
	}
	want := PointList{{X: 1, Y: 2, Tag: "a"}, {X: 3, Y: 4, Tag: "b"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v, want %#v", got, want)
	}
	// null -> nil slice
	var n PointList
	if err := n.UnmarshalJSON([]byte(`null`)); err != nil || n != nil {
		t.Fatalf("null: got %#v err %v, want nil", n, err)
	}
}

// TestMapRoot covers a named map root type (object-root data map).
func TestMapRoot(t *testing.T) {
	var got ScoreMap
	if err := got.UnmarshalJSON([]byte(`{"a":1,"b":2,"c":3}`)); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, ScoreMap{"a": 1, "b": 2, "c": 3}) {
		t.Fatalf("got %#v", got)
	}
	var n ScoreMap
	if err := n.UnmarshalJSON([]byte(`null`)); err != nil || n != nil {
		t.Fatalf("null: got %#v err %v", n, err)
	}
}

// TestNoCopyDirective covers //lightning:nocopy on a map root: the keys and
// string values alias the input instead of being copied.
func TestNoCopyDirective(t *testing.T) {
	data := []byte(`{"alpha":"one","beta":"two"}`)
	var m NoCopyMap
	if err := m.UnmarshalJSON(data); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(m, NoCopyMap{"alpha": "one", "beta": "two"}) {
		t.Fatalf("got %#v", m)
	}
	// A key must alias data (its bytes live inside the input), not be a copy.
	base := uintptr(unsafe.Pointer(&data[0]))
	for k := range m {
		kp := uintptr(unsafe.Pointer(unsafe.StringData(k)))
		if kp < base || kp >= base+uintptr(len(data)) {
			t.Errorf("key %q does not alias input (not nocopy)", k)
		}
	}
}

// TestDestructiveDirective covers //lightning:destructive: an escaped string value is
// unescaped into the input buffer (aliasing and mutating it), an escape-free value
// aliases the input unchanged, and both decode to the same string the slow path
// would produce.
func TestDestructiveDirective(t *testing.T) {
	orig := `{"name":"a\/b\tc","tags":["plain","x\ny"]}`
	data := []byte(orig)
	var d DestructiveDoc
	if err := d.UnmarshalJSON(data); err != nil {
		t.Fatal(err)
	}
	if d.Name != "a/b\tc" || len(d.Tags) != 2 || d.Tags[0] != "plain" || d.Tags[1] != "x\ny" {
		t.Fatalf("got %#v", d)
	}
	// The escaped name was unescaped in place, so it must alias the input...
	base := uintptr(unsafe.Pointer(&data[0]))
	np := uintptr(unsafe.Pointer(unsafe.StringData(d.Name)))
	if np < base || np >= base+uintptr(len(data)) {
		t.Errorf("name %q does not alias input (not decoded in place)", d.Name)
	}
	// ...and the input must have been mutated by the in-place unescape.
	if string(data) == orig {
		t.Errorf("input was not mutated; in-place unescape did not write through")
	}
}
