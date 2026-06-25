package conformance

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
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
		{"Arr", d.Arr, [3]int{10, 20, 30}},

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

// TestEarlyExit exercises the //lightning:earlyexit directive: the decoder
// stops once every declared field has been read and fast-skips the remainder.
func TestEarlyExit(t *testing.T) {
	// All fields present, then a duplicate "a" and an unknown field. The decoder
	// must stop after a,b,c and NOT apply the later "a":99 (encoding/json's
	// last-wins would give a=99) — proving the early-exit fired.
	var d EarlyExitDoc
	if err := d.UnmarshalJSON([]byte(`{"a":1,"b":2,"c":3,"a":99,"x":[1,2,3]}`)); err != nil {
		t.Fatal(err)
	}
	if d.A != 1 || d.B != 2 || d.C != 3 {
		t.Fatalf(`got %#v, want {1 2 3}: early-exit must ignore the trailing "a":99`, d)
	}

	// Once all fields are seen, the rest is bracket-balanced, not grammar-checked:
	// trailing content a strict per-member decode would reject is tolerated.
	var d2 EarlyExitDoc
	if err := d2.UnmarshalJSON([]byte(`{"a":1,"b":2,"c":3, not even valid json here }`)); err != nil {
		t.Fatalf("early-exit should skip unchecked trailing content: %v", err)
	}
	if d2.A != 1 || d2.B != 2 || d2.C != 3 {
		t.Fatalf("got %#v", d2)
	}

	// A missing field means all-fields-seen never triggers, so the whole object is
	// decoded normally and the present fields still land.
	var d3 EarlyExitDoc
	if err := d3.UnmarshalJSON([]byte(`{"a":1,"b":2,"x":7}`)); err != nil {
		t.Fatal(err)
	}
	if d3.A != 1 || d3.B != 2 || d3.C != 0 {
		t.Fatalf("got %#v", d3)
	}
}

func BenchmarkEarlyExit(b *testing.B) {
	var sb strings.Builder
	sb.WriteString(`{"a":1,"b":2,"c":3`)
	for i := 0; i < 200; i++ {
		fmt.Fprintf(&sb, `,"field_number_%d":"some string value here %d"`, i, i)
	}
	sb.WriteString("}")
	data := []byte(sb.String())
	b.Run("normal", func(b *testing.B) {
		b.SetBytes(int64(len(data)))
		for i := 0; i < b.N; i++ {
			var d NormalDoc
			if err := d.UnmarshalJSON(data); err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("earlyexit", func(b *testing.B) {
		b.SetBytes(int64(len(data)))
		for i := 0; i < b.N; i++ {
			var d EarlyExitDoc
			if err := d.UnmarshalJSON(data); err != nil {
				b.Fatal(err)
			}
		}
	})
}
