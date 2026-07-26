package conformance

import (
	"encoding/json"
	"errors"
	"os"
	"reflect"
	"testing"
	"time"
	"unsafe"

	"github.com/JohanLindvall/lightning/pkg/unstable"
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

// nestTree builds a JSON document of n nested Tree levels: {"kids":[{"kids":[…]}]}.
func nestTree(n int) []byte {
	var b []byte
	for i := 0; i < n; i++ {
		b = append(b, `{"name":"x","kids":[`...)
	}
	b = append(b, `{"name":"leaf"}`...)
	for i := 0; i < n; i++ {
		b = append(b, ']', '}')
	}
	return b
}

// TestRecursiveTypeDepthLimit covers the generator's cycle detection. Tree's
// decoder and its []*Tree slice decoder call each other, so a nested document
// recurses once per level; before the depth counter, input like the 2-million-level
// case below overflowed the goroutine stack — a *fatal* error that recover cannot
// catch, taking the process down rather than returning an error. The generator now
// detects the cycle and threads a depth parameter through exactly those decoders.
func TestRecursiveTypeDepthLimit(t *testing.T) {
	// A moderate nesting decodes normally, all the way down.
	var tr Tree
	if err := tr.UnmarshalJSON(nestTree(100)); err != nil {
		t.Fatalf("100 levels should decode: %v", err)
	}
	depth, n := 0, &tr
	for n != nil {
		depth++
		if len(n.Kids) == 0 {
			break
		}
		n = n.Kids[0]
	}
	if depth != 101 { // 100 wrappers plus the leaf
		t.Errorf("decoded depth = %d, want 101", depth)
	}
	if n.Name != "leaf" {
		t.Errorf("deepest node Name = %q, want \"leaf\"", n.Name)
	}

	// Past the bound it reports an error instead of crashing. The 2-million case is
	// the regression test: that is the shape that used to be fatal.
	for _, levels := range []int{unstable.MaxDepth + 1, 2_000_000} {
		var deep Tree
		err := deep.UnmarshalJSON(nestTree(levels))
		if !errors.Is(err, unstable.ErrMaxDepth) {
			t.Errorf("%d levels: err = %v, want ErrMaxDepth", levels, err)
		}
	}
}

// TestMutuallyRecursiveTypeDepthLimit is the two-type cycle: Ring1 -> Ring2 ->
// Ring1, reached through RingRoot, which is itself not on the cycle and so only
// passes the counter down. It checks the cycle search follows edges rather than
// only spotting a direct self-reference.
func TestMutuallyRecursiveTypeDepthLimit(t *testing.T) {
	nest := func(n int) []byte {
		var b []byte
		b = append(b, `{"start":`...)
		for i := 0; i < n; i++ {
			b = append(b, `{"name":"a","next":{"count":1,"back":`...)
		}
		b = append(b, `null`...)
		for i := 0; i < n; i++ {
			b = append(b, '}', '}')
		}
		b = append(b, '}')
		return b
	}

	var r RingRoot
	if err := r.UnmarshalJSON(nest(50)); err != nil {
		t.Fatalf("50 rings should decode: %v", err)
	}
	if r.Start == nil || r.Start.Name != "a" || r.Start.Next == nil || r.Start.Next.Count != 1 {
		t.Fatalf("got %#v", r.Start)
	}

	var deep RingRoot
	if err := deep.UnmarshalJSON(nest(1_000_000)); !errors.Is(err, unstable.ErrMaxDepth) {
		t.Errorf("1M rings: err = %v, want ErrMaxDepth", err)
	}
}

// TestNonRecursiveTypesTakeNoDepthParam is the other half of the contract: the
// depth counter is threaded *only* through decoders that can reach a cycle, so a
// schema without one generates exactly the code it did before and pays nothing.
// Doc is large and cycle-free, so its decoder must still have the three-parameter
// signature — asserted by the compiler here, since a depth parameter would make
// this call fail to build.
func TestNonRecursiveTypesTakeNoDepthParam(t *testing.T) {
	raw, err := os.ReadFile("test.json")
	if err != nil {
		t.Fatal(err)
	}
	var d Doc
	if err := d.UnmarshalJSON(raw); err != nil {
		t.Fatal(err)
	}
	// PointList and ScoreMap are cycle-free roots too.
	var pl PointList
	if err := pl.UnmarshalJSON([]byte(`[{"x":1,"y":2,"tag":"a"}]`)); err != nil {
		t.Fatal(err)
	}
	var sm ScoreMap
	if err := sm.UnmarshalJSON([]byte(`{"a":1}`)); err != nil {
		t.Fatal(err)
	}
}

// TestSliceReuseReplaces locks the rule that decoding an array into a slice
// *replaces* its contents rather than appending to them, which is what
// encoding/json documents: "Unmarshal resets the slice length to zero and then
// appends each element to the slice."
//
// It is a regression test for a real bug: every slice decoder started from
// `*out` rather than `(*out)[:0]`, so decoding a second document into the same
// value accumulated — [1,2] read twice became [1,2,1,2] — silently, with no error,
// and a caller reusing a target to avoid allocation (the whole reason to reuse one)
// grew it without bound. No benchmark caught it because every benchmark decodes
// into a freshly declared value.
//
// The backing array must still be kept, or reuse would start allocating; that half
// is checked by TestSliceReuseKeepsBacking below.
func TestSliceReuseReplaces(t *testing.T) {
	first := []byte(`{"ints":[1,2,3],"strs":["a","b","c"],"floats":[1.5,2.5],` +
		`"grid":[[1,2],[3,4]],"items":[{"name":"x","count":1},{"name":"y","count":2}],` +
		`"ptrItems":[{"name":"p","count":9}],"anys":[1,2,3]}`)
	second := []byte(`{"ints":[9],"strs":["z"],"floats":[9.5],"grid":[[9]],` +
		`"items":[{"name":"q","count":8}],"ptrItems":[{"name":"r","count":7}],"anys":[9]}`)

	var d Doc
	if err := d.UnmarshalJSON(first); err != nil {
		t.Fatal(err)
	}
	if err := d.UnmarshalJSON(second); err != nil {
		t.Fatal(err)
	}

	if got, want := d.Ints, []int{9}; !reflect.DeepEqual(got, want) {
		t.Errorf("Ints = %v, want %v (accumulated instead of replaced)", got, want)
	}
	if got, want := d.Strs, []string{"z"}; !reflect.DeepEqual(got, want) {
		t.Errorf("Strs = %v, want %v", got, want)
	}
	if got, want := d.Floats, []float64{9.5}; !reflect.DeepEqual(got, want) {
		t.Errorf("Floats = %v, want %v", got, want)
	}
	if got, want := d.Grid, [][]int{{9}}; !reflect.DeepEqual(got, want) {
		t.Errorf("Grid = %v, want %v", got, want)
	}
	if len(d.Items) != 1 || d.Items[0].Name != "q" || d.Items[0].Count != 8 {
		t.Errorf("Items = %+v, want one element {q 8}", d.Items)
	}
	if len(d.PtrItems) != 1 || d.PtrItems[0].Name != "r" {
		t.Errorf("PtrItems = %+v, want one element {r 7}", d.PtrItems)
	}
	if got, want := d.Anys, []any{9.0}; !reflect.DeepEqual(got, want) {
		t.Errorf("Anys = %v, want %v", got, want)
	}

	// A named slice root type takes the same path and the same rule.
	var pl PointList
	if err := pl.UnmarshalJSON([]byte(`[{"x":1,"y":1,"tag":"a"},{"x":2,"y":2,"tag":"b"}]`)); err != nil {
		t.Fatal(err)
	}
	if err := pl.UnmarshalJSON([]byte(`[{"x":9,"y":9,"tag":"z"}]`)); err != nil {
		t.Fatal(err)
	}
	if len(pl) != 1 || pl[0].Tag != "z" {
		t.Errorf("PointList = %+v, want one element tagged z", pl)
	}

	// A nocopy slice (DestructiveDoc.Tags) too. Its own input must stay writable,
	// so give each decode its own buffer.
	var dd DestructiveDoc
	for _, s := range []string{`{"tags":["a","b"]}`, `{"tags":["z"]}`} {
		if err := dd.UnmarshalJSON([]byte(s)); err != nil {
			t.Fatal(err)
		}
	}
	if got, want := dd.Tags, []string{"z"}; !reflect.DeepEqual(got, want) {
		t.Errorf("nocopy Tags = %v, want %v", got, want)
	}
}

// TestSliceReuseKeepsBacking is the performance half of the reuse rule: resetting
// the length must reuse the backing array, so decoding repeatedly into one value
// stays allocation-free. Throwing the array away would make the correctness fix a
// performance regression.
func TestSliceReuseKeepsBacking(t *testing.T) {
	doc := []byte(`{"ints":[1,2,3,4,5,6,7,8],"floats":[1,2,3,4]}`)
	var d Doc
	if err := d.UnmarshalJSON(doc); err != nil {
		t.Fatal(err)
	}
	intsAddr, intsCap := &d.Ints[0], cap(d.Ints)
	fltAddr, fltCap := &d.Floats[0], cap(d.Floats)

	if err := d.UnmarshalJSON(doc); err != nil {
		t.Fatal(err)
	}
	if &d.Ints[0] != intsAddr || cap(d.Ints) != intsCap {
		t.Errorf("Ints backing array replaced on reuse (cap %d -> %d)", intsCap, cap(d.Ints))
	}
	if &d.Floats[0] != fltAddr || cap(d.Floats) != fltCap {
		t.Errorf("Floats backing array replaced on reuse (cap %d -> %d)", fltCap, cap(d.Floats))
	}
	if n := testing.AllocsPerRun(50, func() {
		if err := d.UnmarshalJSON(doc); err != nil {
			t.Fatal(err)
		}
	}); n != 0 {
		t.Errorf("reuse allocated %v times per decode, want 0", n)
	}
}

// TestLongKeyDispatch covers the `switch len(key)` dispatch the generator emits for
// a struct with a JSON name longer than 16 bytes, where names are compared in
// <=16-byte chunks so no comparison becomes a runtime.memequal call.
//
// The chunking is the part that can go subtly wrong, so the cases are chosen to
// break a careless implementation rather than to look representative:
// sharedPrefix16xxA and sharedPrefix16xxB are the same length and share their entire
// first 16-byte chunk, so comparing only the first chunk would swap them; the 33-byte
// name needs three chunks (16 + 16 + 1); exactly16bytes__ sits on the boundary where
// the compiler still inlines; and Alt's pipe-separated names fall either side of it,
// so its decode code is emitted in two different length buckets.
func TestLongKeyDispatch(t *testing.T) {
	full := []byte(`{
		"s": 1,
		"exactly16bytes__": 2,
		"seventeenBytes_17": 3,
		"sharedPrefix16xxA": 4,
		"sharedPrefix16xxB": 5,
		"differentButSame": 6,
		"aKeyOfThirtyThreeBytesExactly_333": "x",
		"shortAlt": 7
	}`)
	var lk LongKeys
	if err := lk.UnmarshalJSON(full); err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		name string
		got  int
		want int
	}{
		{"Short", lk.Short, 1},
		{"Exactly16Bytes__", lk.Exactly16Bytes__, 2},
		{"SeventeenBytes_1", lk.SeventeenBytes_1, 3},
		{"SharedPrefixA", lk.SharedPrefixA, 4},
		{"SharedPrefixB", lk.SharedPrefixB, 5},
		{"SameLenOther", lk.SameLenOther, 6},
		{"Alt (short spelling)", lk.Alt, 7},
	} {
		if c.got != c.want {
			t.Errorf("%s = %d, want %d", c.name, c.got, c.want)
		}
	}
	if lk.ThirtyThreePlus != "x" {
		t.Errorf("ThirtyThreePlus = %q, want \"x\"", lk.ThirtyThreePlus)
	}

	// The long alternate spelling fills the same field, from the other bucket.
	var alt LongKeys
	if err := alt.UnmarshalJSON([]byte(`{"aVeryMuchLongerAlternateName":9}`)); err != nil {
		t.Fatal(err)
	}
	if alt.Alt != 9 {
		t.Errorf("Alt via long alternate name = %d, want 9", alt.Alt)
	}

	// Unknown keys must be skipped, not misrouted: each of these is a near miss
	// against a real name — differing in the last chunk, one byte too long or short,
	// or sharing a 16-byte prefix — and none may set any field.
	for _, doc := range []string{
		`{"sharedPrefix16xxC":1}`,                 // same length, last chunk differs
		`{"sharedPrefix16xx":1}`,                  // one byte short of two names
		`{"sharedPrefix16xxAB":1}`,                // one byte long
		`{"exactly16bytes_":1}`,                   // 15 bytes
		`{"exactly16bytes___":1}`,                 // 17 bytes
		`{"aKeyOfThirtyThreeBytesExactly_33x":1}`, // 33 bytes, last chunk differs
		`{"aVeryMuchLongerAlternateNam":1}`,       // 27 bytes
		`{"differentButSamX":1}`,                  // 16 bytes, last byte differs
	} {
		var z LongKeys
		if err := z.UnmarshalJSON([]byte(doc)); err != nil {
			t.Errorf("%s: unexpected error %v", doc, err)
			continue
		}
		if z != (LongKeys{}) {
			t.Errorf("%s: set a field (%+v), want all zero", doc, z)
		}
	}
}
