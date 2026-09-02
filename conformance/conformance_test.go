package conformance

import (
	"encoding/json"
	"errors"
	"fmt"
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

	// The null rule has to hold here too, including on a buffer this decoder has
	// already rewritten: a nocopy string field clears on null (the documented
	// leaf divergence — see nullGuard) and a slice field nils (parity). Worth a
	// case of its own because destructive is the one mode that mutates data while
	// decoding it, so a later field's null is read out of a buffer that no longer
	// matches the caller's original bytes.
	seeded := DestructiveDoc{Name: "keep", Tags: []string{"t"}}
	nulls := []byte(`{"name":null,"tags":null}`)
	if err := seeded.UnmarshalJSON(nulls); err != nil {
		t.Fatal(err)
	}
	if seeded.Name != "" {
		t.Errorf("null into a destructive nocopy string: %q, want the zero value", seeded.Name)
	}
	if seeded.Tags != nil {
		t.Errorf("null left a slice field at %v, want nil", seeded.Tags)
	}
	// And with a *mutating* read earlier in the same document, so the null is
	// decoded after the in-place unescape has written into the buffer.
	mixed := []byte(`{"tags":["a\tb","c\nd"],"name":null}`)
	seeded2 := DestructiveDoc{Name: "keep2"}
	if err := seeded2.UnmarshalJSON(mixed); err != nil {
		t.Fatal(err)
	}
	if seeded2.Name != "" || len(seeded2.Tags) != 2 || seeded2.Tags[0] != "a\tb" {
		t.Errorf("got %#v", seeded2)
	}
}

// arenaDocStd strips ArenaDoc's UnmarshalJSON so encoding/json decodes it by
// reflection, giving the parity baseline (ArenaKey never had a method: it is
// referenced by ArenaDoc, so its decoder is internal).
type arenaDocStd ArenaDoc

// TestArenaDirective covers //lightning:arena end to end: the decoded value must
// equal encoding/json's bit for bit (the arena changes where backings live, never
// what is decoded), covering adjacent small carves, sub-8-byte element kinds, the
// over-threshold direct-make fallback, and null/empty arrays. Then the safety
// property: growing one arena-backed slice must reallocate (exact counts leave
// len == cap) and leave every sibling carve intact. Finally, decoding again into
// the same value must replace contents, reusing backings without a fresh carve.
func TestArenaDirective(t *testing.T) {
	doc := `{
	  "name": "anim",
	  "keys": [
	    {"pos": [1.5, 2.5, 3.5], "rot": [-1, 2, -3], "cnt": [7], "time": 0.25},
	    {"pos": [4, 5], "rot": [], "cnt": [1, 2, 3, 4], "time": 0.5},
	    {"pos": null, "rot": [9], "big": [` + bigInts(100) + `], "time": 0.75}
	  ]
	}`
	var got ArenaDoc
	if err := got.UnmarshalJSON([]byte(doc)); err != nil {
		t.Fatal(err)
	}
	var want arenaDocStd
	if err := json.Unmarshal([]byte(doc), &want); err != nil {
		t.Fatal(err)
	}
	// One deliberate, pre-existing divergence from encoding/json (not an arena
	// behavior): lightning decodes a JSON [] to a nil slice, the stdlib to a
	// non-nil empty one. Normalize the baseline; everything else must match.
	want.Keys[1].Rot = nil
	if !reflect.DeepEqual(got, ArenaDoc(want)) {
		t.Fatalf("arena decode diverges from encoding/json:\n got %#v\nwant %#v", got, want)
	}

	// Sibling safety: append far past Keys[0].Pos's capacity and then mutate the
	// grown copy; no other carve may change.
	if cap(got.Keys[0].Pos) != len(got.Keys[0].Pos) {
		t.Fatalf("exact count should leave len == cap, got len %d cap %d", len(got.Keys[0].Pos), cap(got.Keys[0].Pos))
	}
	grown := append(got.Keys[0].Pos, make([]float64, 64)...)
	for i := range grown {
		grown[i] = -1
	}
	if !reflect.DeepEqual(got.Keys[0].Rot, []int32{-1, 2, -3}) ||
		!reflect.DeepEqual(got.Keys[1].Pos, []float64{4, 5}) ||
		!reflect.DeepEqual(got.Keys[1].Cnt, []uint16{1, 2, 3, 4}) {
		t.Fatalf("append to one arena slice disturbed a sibling: %#v", got.Keys)
	}

	// Reuse: decode a second document into the same value; contents must be
	// replaced (not appended) and still match encoding/json exactly.
	doc2 := `{"name":"anim2","keys":[{"pos":[9.5],"rot":[4,5],"cnt":[],"time":1}]}`
	if err := got.UnmarshalJSON([]byte(doc2)); err != nil {
		t.Fatal(err)
	}
	var want2 arenaDocStd
	if err := json.Unmarshal([]byte(doc2), &want2); err != nil {
		t.Fatal(err)
	}
	want2.Keys[0].Cnt = nil // the []-to-nil normalization again
	if !reflect.DeepEqual(got, ArenaDoc(want2)) {
		t.Fatalf("arena reuse decode diverges:\n got %#v\nwant %#v", got, want2)
	}

	// Arena on a recursive schema: the decoders thread depth AND the arena, so a
	// nested document must decode correctly (parameter-order bugs between the two
	// extra parameters would break this immediately) and still enforce the depth
	// bound.
	var tree ArenaTree
	if err := tree.UnmarshalJSON([]byte(`{"vals":[1,2],"kids":[{"vals":[3],"kids":[{"vals":[4,5,6]}]}]}`)); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(tree.Vals, []float64{1, 2}) ||
		!reflect.DeepEqual(tree.Kids[0].Kids[0].Vals, []float64{4, 5, 6}) {
		t.Fatalf("arena tree misdecoded: %#v", tree)
	}
	var deep ArenaTree
	if err := deep.UnmarshalJSON(nestTree(unstable.MaxDepth + 10)); err == nil || !errors.Is(err, unstable.ErrMaxDepth) {
		t.Fatalf("deep arena tree: want ErrMaxDepth, got %v", err)
	}
}

// bigInts renders "0,1,…,n-1", a scalar array big enough to cross the arena's
// carve threshold so the direct-make fallback is exercised in a real decode.
func bigInts(n int) string {
	b := make([]byte, 0, 4*n)
	for i := 0; i < n; i++ {
		if i > 0 {
			b = append(b, ',')
		}
		b = append(b, []byte(fmt.Sprintf("%d", i))...)
	}
	return string(b)
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

// TestLaxDecoderIsolation locks the valueDecoder memo-key fix: lax value
// decoders must be per-root (prefix + directive marker) like every other
// generated decoder, so a plain root never inherits a destructive or compact
// sibling's semantics through a shared lax field type.
func TestLaxDecoderIsolation(t *testing.T) {
	// Plain root: decoding must not mutate the input even though a destructive
	// root with the identical lax field type was generated first.
	orig := `{"s":"a\nb"}`
	data := []byte(orig)
	var p LaxSharedPlain
	if err := p.UnmarshalJSON(data); err != nil {
		t.Fatal(err)
	}
	if p.S != "a\nb" {
		t.Fatalf("plain root decoded %q, want %q", p.S, "a\nb")
	}
	if string(data) != orig {
		t.Fatalf("plain root mutated its input: %q", data)
	}

	// The destructive root itself must still unescape in place.
	ddata := []byte(orig)
	var d LaxSharedDestructive
	if err := d.UnmarshalJSON(ddata); err != nil {
		t.Fatal(err)
	}
	if d.S != "a\nb" {
		t.Fatalf("destructive root decoded %q, want %q", d.S, "a\nb")
	}
	if string(ddata) == orig {
		t.Fatal("destructive root did not mutate its input (in-place unescape lost)")
	}

	// Plain root with a lax any field: whitespaced input must decode even
	// though a compact root with the identical lax field type exists.
	var ap LaxSharedAnyPlain
	if err := ap.UnmarshalJSON([]byte(`{"v": { "a" : 1 }}`)); err != nil {
		t.Fatal(err)
	}
	if ap.V == nil {
		t.Fatal("plain root left V nil: lax decoder leaked compact semantics")
	}
}

// TestUnwrapWhitespaceBody locks the bounds guard in the generated unwrap
// closure: a wrapped body that is empty or all whitespace leaves the field at
// its zero value — it must not panic (the pointer field's null probe used to
// index one past the end) and must not consume the field's zero state.
func TestUnwrapWhitespaceBody(t *testing.T) {
	cases := []string{
		`{"p":" "}`,          // single space — the original panic input
		`{"p":"   \t\n  "}`,  // longer mixed whitespace
		`{"p":""}`,           // empty body (previously safe; must stay so)
		`{"s":" "}`,          // *string field, same guard
		`{"p":" ","s":"\t"}`, // both fields
	}
	for _, in := range cases {
		var u UnwrapPtr
		if err := u.UnmarshalJSON([]byte(in)); err != nil {
			t.Errorf("UnmarshalJSON(%q): %v", in, err)
			continue
		}
		if u.P != nil || u.S != nil {
			t.Errorf("UnmarshalJSON(%q): fields not zero: %#v", in, u)
		}
	}
	// A non-empty wrapped body still decodes.
	var u UnwrapPtr
	if err := u.UnmarshalJSON([]byte(`{"p":" {\"count\": 7} ","s":"\"hi\""}`)); err != nil {
		t.Fatal(err)
	}
	if u.P == nil || u.P.Count != 7 || u.S == nil || *u.S != "hi" {
		t.Fatalf("got %#v", u)
	}
}

// ptrReuseStd strips PtrReuse's generated UnmarshalJSON so encoding/json
// decodes it by reflection, giving the stdlib-parity baseline.
type ptrReuseStd PtrReuse

// TestPointerFieldReuse locks the guarded pointer allocation: decoding into a
// target whose pointer fields are non-nil must reuse the pointees (stdlib
// semantics — pointer identity kept, omitted pointee fields left as they were)
// instead of allocating fresh ones, and null must still nil the pointer.
func TestPointerFieldReuse(t *testing.T) {
	doc := []byte(`{"p":{"name":"x"}}`)

	// Premise: this is what encoding/json does with a reused non-nil pointer.
	sp := &Nested{Name: "old", Count: 5}
	sn := 3
	std := ptrReuseStd{P: (*Nested)(sp), N: &sn}
	if err := json.Unmarshal(doc, &std); err != nil {
		t.Fatal(err)
	}
	if std.P != sp || std.P.Name != "x" || std.P.Count != 5 || std.N != &sn || *std.N != 3 {
		t.Fatalf("stdlib premise changed: %#v (pointee reuse no longer stdlib semantics?)", std)
	}

	// Lightning must match: same pointer identity, same surviving pointee state.
	lp := &Nested{Name: "old", Count: 5}
	ln := 3
	l := PtrReuse{P: lp, N: &ln}
	if err := l.UnmarshalJSON(doc); err != nil {
		t.Fatal(err)
	}
	if l.P != lp {
		t.Error("pointer field was reallocated; want pointee reuse")
	}
	if l.P.Name != "x" || l.P.Count != 5 {
		t.Errorf("pointee = %#v, want Name=x Count=5", *l.P)
	}
	if l.N != &ln || *l.N != 3 {
		t.Error("untouched pointer field must keep its pointee")
	}

	// null still nils the pointer, and a nil pointer still allocates.
	if err := l.UnmarshalJSON([]byte(`{"p":null,"n":9}`)); err != nil {
		t.Fatal(err)
	}
	if l.P != nil || l.N == nil || *l.N != 9 {
		t.Fatalf("got %#v", l)
	}
}

// byteSliceDocStd strips the generated method for the reflection baseline.
type byteSliceDocStd ByteSliceDoc

// TestByteSliceStdlibParity locks []byte decoding to encoding/json: base64
// strings and numeric arrays are both accepted (bit-identical results), null
// is nil, invalid base64 errors, and [N]byte accepts only the numeric form.
func TestByteSliceStdlibParity(t *testing.T) {
	cases := []string{
		`{"b":"AQID"}`,                      // base64 string (stdlib Marshal form)
		`{"b":[1,2,3]}`,                     // numeric array form
		`{"b":null}`,                        // null -> nil
		`{"b":""}`,                          // empty base64 -> empty slice
		`{"fixed":[7,8,9]}`,                 // [N]byte numeric
		`{"many":["AQ==","Ag==",[3],null]}`, // elements of [][]byte are []byte too
	}
	for _, in := range cases {
		var std byteSliceDocStd
		stdErr := json.Unmarshal([]byte(in), &std)
		var l ByteSliceDoc
		lErr := l.UnmarshalJSON([]byte(in))
		if (stdErr == nil) != (lErr == nil) {
			t.Errorf("%s: error disagreement: stdlib %v, lightning %v", in, stdErr, lErr)
			continue
		}
		if stdErr == nil && !reflect.DeepEqual(ByteSliceDoc(std), l) {
			t.Errorf("%s:\n stdlib    %#v\n lightning %#v", in, std, l)
		}
	}
	// A named byte-slice root gets base64 exactly like []byte.
	var blob ByteBlob
	if err := blob.UnmarshalJSON([]byte(`"aGVsbG8="`)); err != nil || string(blob) != "hello" {
		t.Errorf("ByteBlob base64 root: %q, %v", blob, err)
	}
	if err := blob.UnmarshalJSON([]byte(`[104,105]`)); err != nil || string(blob) != "hi" {
		t.Errorf("ByteBlob numeric root: %q, %v", blob, err)
	}

	// Both reject what the other rejects.
	for _, in := range []string{`{"b":"!!!not base64!!!"}`, `{"fixed":"AQID"}`} {
		var std byteSliceDocStd
		if json.Unmarshal([]byte(in), &std) == nil {
			t.Fatalf("stdlib accepted %s — premise wrong", in)
		}
		var l ByteSliceDoc
		if l.UnmarshalJSON([]byte(in)) == nil {
			t.Errorf("lightning accepted %s, stdlib rejects it", in)
		}
	}
}

// TestLaxFixedArrays locks the batched routing for lax [N]scalar fields: valid
// arrays decode exactly as non-lax ones (leading fill, short leaves zeros,
// surplus skipped), and a mistyped value is skipped leaving the zero array.
func TestLaxFixedArrays(t *testing.T) {
	var l LaxArrays
	if err := l.UnmarshalJSON([]byte(`{"f":[1.5,2.5],"i":[-3,4,99],"u":[7]}`)); err != nil {
		t.Fatal(err)
	}
	if l.F != [3]float64{1.5, 2.5, 0} || l.I != [2]int32{-3, 4} || l.U != [4]uint16{7, 0, 0, 0} {
		t.Fatalf("got %#v", l)
	}
	// lax: a mistyped value is skipped, and previously decoded content is
	// replaced by zeros only where the document writes.
	var m LaxArrays
	if err := m.UnmarshalJSON([]byte(`{"f":"nope","i":{"x":1},"u":[9]}`)); err != nil {
		t.Fatal(err)
	}
	if m.F != [3]float64{} || m.I != [2]int32{} || m.U != [4]uint16{9} {
		t.Fatalf("got %#v", m)
	}
}

// nullDocStd and nullLaxDocStd strip the generated UnmarshalJSON so encoding/json
// decodes these by reflection. Without the twin, json.Unmarshal would delegate
// straight back to lightning and the differential would compare the decoder
// against itself — see TestStdlibTwinsAreReflectionOnly.
type nullDocStd NullDoc

type nullLaxDocStd NullLaxDoc

// seedNullDoc returns a NullDoc with every field set to something distinguishable
// from its zero value, so a null that wrongly overwrites a field is visible.
func seedNullDoc() NullDoc {
	p := 42
	return NullDoc{
		Str: "keep", StrNC: "keepNC", Bool: true,
		I: 1, I8: 2, I64: 3, U: 4, U16: 5, F32: 6.5, F64: 7.5,
		Num: json.Number("8.5"), NumNC: json.Number("9.5"),
		Time:   time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC),
		Nested: Nested{Name: "n", Count: 11},
		Arr:    [3]int{1, 2, 3},
		Sl:     []int{1, 2}, Strs: []string{"a"},
		M: map[string]int{"a": 1}, P: &p, A: "any",
		R: json.RawMessage(`{"old":1}`),
	}
}

// zeroNullLeaves applies lightning's one deliberate departure from
// encoding/json's null rule to a stdlib result, so the tests below can say
// "identical to encoding/json except exactly this" instead of hand-writing an
// expected struct. An explicit null stores the ZERO value in a scalar,
// json.Number or time.Time field, where encoding/json leaves it as it was; every
// other kind agrees (see nullGuard in main.go for why the parity version was
// reverted, and the README's divergence list).
func zeroNullLeaves(d *NullDoc) {
	d.Str, d.StrNC, d.Bool = "", "", false
	d.I, d.I8, d.I64, d.U, d.U16 = 0, 0, 0, 0, 0
	d.F32, d.F64 = 0, 0
	d.Num, d.NumNC = "", ""
	d.Time = time.Time{}
}

// TestNullFieldsDivergeFromStdlib pins the null rule field by field on a SEEDED
// target — the only place any of it is observable, since into a fresh target
// "leave the value alone" and "store the zero value" are the same thing. That is
// why the difference went unnoticed for so long: every benchmark, and most of
// this suite, decodes into a freshly declared value.
//
// The test is built like TestValidDivergesFromStdlib: the expectation comes from
// a methodless twin with a NAMED transformation applied, so it asserts two things
// at once — that the divergence is exactly what it claims to be, and that it
// reaches no further. A future change that widened it to, say, nested structs, or
// that quietly closed it, fails here either way. Nothing about this is a fix
// waiting to happen: it is the recorded trade in nullGuard.
func TestNullFieldsDivergeFromStdlib(t *testing.T) {
	allNull := []byte(`{"str":null,"strNC":null,"bool":null,"i":null,"i8":null,` +
		`"i64":null,"u":null,"u16":null,"f32":null,"f64":null,"num":null,` +
		`"numNC":null,"time":null,"nested":null,"arr":null,"sl":null,"strs":null,` +
		`"m":null,"p":null,"a":null,"r":null}`)

	t.Run("seeded", func(t *testing.T) {
		std := nullDocStd(seedNullDoc())
		if err := json.Unmarshal(allNull, &std); err != nil {
			t.Fatalf("stdlib: %v", err)
		}
		// Premise check: the stdlib really does keep the scalars. If this fails,
		// the rule changed and the rest of the test means nothing.
		if std.Str != "keep" || std.I != 1 || std.Time.IsZero() || std.Nested.Count != 11 || std.Arr != [3]int{1, 2, 3} {
			t.Fatalf("stdlib premise changed: %+v", std)
		}
		l := seedNullDoc()
		if err := l.UnmarshalJSON(allNull); err != nil {
			t.Fatalf("lightning: %v", err)
		}
		want := NullDoc(std)
		zeroNullLeaves(&want)
		if !reflect.DeepEqual(want, l) {
			t.Errorf("null into a seeded target is not stdlib-with-zeroed-leaves:\n want      %+v\n lightning %+v", want, l)
		}
		// Spelled out as well as derived, so the divergence is legible here and
		// not only in the helper: the leaves go to zero, the composites do not.
		if l.Str != "" || l.I != 0 || !l.Time.IsZero() || l.Num != "" {
			t.Errorf("a null leaf should store the zero value: %+v", l)
		}
		if l.Nested.Count != 11 || l.Arr != [3]int{1, 2, 3} {
			t.Errorf("a null nested struct or fixed array must be left alone: %+v %v", l.Nested, l.Arr)
		}
	})

	t.Run("fresh", func(t *testing.T) {
		// The already-correct case, kept as a regression guard: a null into a
		// fresh target must still produce the zero value, and nil rather than
		// empty for the reference kinds.
		var std nullDocStd
		if err := json.Unmarshal(allNull, &std); err != nil {
			t.Fatalf("stdlib: %v", err)
		}
		var l NullDoc
		if err := l.UnmarshalJSON(allNull); err != nil {
			t.Fatalf("lightning: %v", err)
		}
		if !reflect.DeepEqual(NullDoc(std), l) {
			t.Errorf("null into a fresh target diverges:\n stdlib    %+v\n lightning %+v", std, l)
		}
	})

	t.Run("reuse", func(t *testing.T) {
		// The shape a real caller hits, and the one the divergence actually costs
		// something: decode a full document, then one whose fields are null. The
		// second decode nils the reference kinds like the stdlib, but clears the
		// first document's scalars where the stdlib would have kept them.
		full := []byte(`{"str":"one","strNC":"oneNC","bool":true,"i":10,"i8":11,` +
			`"i64":12,"u":13,"u16":14,"f32":1.5,"f64":2.5,"num":3.5,"numNC":4.5,` +
			`"time":"2021-03-04T05:06:07Z","nested":{"name":"x","count":9},` +
			`"arr":[7,8,9],"sl":[1],"strs":["s"],"m":{"k":1},"p":5,"a":"aa","r":{"z":1}}`)
		var std nullDocStd
		if err := json.Unmarshal(full, &std); err != nil {
			t.Fatalf("stdlib: %v", err)
		}
		if err := json.Unmarshal(allNull, &std); err != nil {
			t.Fatalf("stdlib: %v", err)
		}
		var l NullDoc
		if err := l.UnmarshalJSON(full); err != nil {
			t.Fatalf("lightning: %v", err)
		}
		if err := l.UnmarshalJSON(allNull); err != nil {
			t.Fatalf("lightning: %v", err)
		}
		if l.Str != "" || l.I != 0 {
			t.Errorf("a null over a reused target should clear the scalar: %+v", l)
		}
		want := NullDoc(std)
		zeroNullLeaves(&want)
		if !reflect.DeepEqual(want, l) {
			t.Errorf("null over a reused target is not stdlib-with-zeroed-leaves:\n want      %+v\n lightning %+v", want, l)
		}
	})

	t.Run("lax", func(t *testing.T) {
		// The same rule reached through ",lax", which decodes into a scratch
		// value and commits it. That commit has to reproduce whatever the field's
		// plain form does, or json:"i" and json:"i,lax" disagree with each other
		// on {"i":null} — so a lax leaf zeroes (the divergence) while a lax
		// nested struct or fixed array is left alone (parity), exactly as above.
		laxNull := []byte(`{"str":null,"i":null,"f64":null,"num":null,"time":null,` +
			`"nested":null,"arr":null,"sl":null,"m":null,"p":null,"a":null,"r":null}`)
		p := 42
		seed := NullLaxDoc{
			Str: "keep", I: 1, F64: 2.5, Num: json.Number("3.5"),
			Time:   time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC),
			Nested: Nested{Name: "n", Count: 11}, Arr: [3]int{1, 2, 3},
			Sl: []int{1}, M: map[string]int{"a": 1}, P: &p, A: "any",
			R: json.RawMessage(`{"old":1}`),
		}
		std := nullLaxDocStd(seed)
		if err := json.Unmarshal(laxNull, &std); err != nil {
			t.Fatalf("stdlib: %v", err)
		}
		l := seed
		if err := l.UnmarshalJSON(laxNull); err != nil {
			t.Fatalf("lightning: %v", err)
		}
		want := NullLaxDoc(std)
		want.Str, want.I, want.F64, want.Num, want.Time = "", 0, 0, "", time.Time{}
		if !reflect.DeepEqual(want, l) {
			t.Errorf("null into a seeded lax target is not stdlib-with-zeroed-leaves:\n want      %+v\n lightning %+v", want, l)
		}

		// The property that makes the above worth testing at all: whatever the
		// rule is, the two spellings of the same field agree on it.
		var plain NullDoc
		plain.Str, plain.I, plain.F64 = "keep", 1, 2.5
		plain.Nested, plain.Arr = Nested{Name: "n", Count: 11}, [3]int{1, 2, 3}
		if err := plain.UnmarshalJSON([]byte(`{"str":null,"i":null,"f64":null,"nested":null,"arr":null}`)); err != nil {
			t.Fatalf("lightning: %v", err)
		}
		if plain.Str != l.Str || plain.I != l.I || plain.F64 != l.F64 ||
			plain.Nested != l.Nested || plain.Arr != l.Arr {
			t.Errorf("json:\"x\" and json:\"x,lax\" disagree on null:\n plain %+v\n lax   %+v", plain, l)
		}
	})

	t.Run("elements", func(t *testing.T) {
		// A null *element* of a slice or map decodes into a freshly zeroed slot,
		// so it must stay zero-filled and nothing about it changes. Included
		// because it is the property a careless "just skip the assignment" fix
		// breaks: it would leave a stale element behind on a reused backing.
		in := []byte(`{"sl":[1,null,3],"strs":["a",null],"m":{"a":1,"b":null},"nested":{"name":null,"count":null}}`)
		var std nullDocStd
		if err := json.Unmarshal(in, &std); err != nil {
			t.Fatalf("stdlib: %v", err)
		}
		var l NullDoc
		if err := l.UnmarshalJSON(in); err != nil {
			t.Fatalf("lightning: %v", err)
		}
		if !reflect.DeepEqual(NullDoc(std), l) {
			t.Errorf("null elements diverge:\n stdlib    %+v\n lightning %+v", std, l)
		}
		if !reflect.DeepEqual(l.Sl, []int{1, 0, 3}) || !reflect.DeepEqual(l.Strs, []string{"a", ""}) {
			t.Errorf("null elements are not zero-filled: %v %v", l.Sl, l.Strs)
		}
		if l.M["b"] != 0 {
			t.Errorf("null map value is not zero: %v", l.M)
		}

		// The same document over a SEEDED target is deliberately NOT compared
		// against the stdlib: encoding/json decodes an array element into
		// whatever the backing array already holds, so its second element keeps
		// the old 2, while lightning resets each element first. That is the
		// pre-existing, documented divergence under "Slice, array and map
		// elements are reset before being decoded" — asserted here so the null
		// guard cannot quietly change which of the two rules applies.
		s := seedNullDoc()
		if err := s.UnmarshalJSON(in); err != nil {
			t.Fatalf("lightning: %v", err)
		}
		if !reflect.DeepEqual(s.Sl, []int{1, 0, 3}) {
			t.Errorf("a null element over a seeded slice must reset, got %v", s.Sl)
		}
		// The nested struct's own fields ARE reached by the document, so their
		// nulls follow the leaf rule and clear the seed — where a null for the
		// struct as a whole would have left it untouched. Both halves asserted,
		// since the interesting part is that the two differ.
		if s.Nested.Name != "" || s.Nested.Count != 0 {
			t.Errorf("null leaves inside a seeded nested struct should clear: %+v", s.Nested)
		}
		w := seedNullDoc()
		if err := w.UnmarshalJSON([]byte(`{"nested":null}`)); err != nil {
			t.Fatalf("lightning: %v", err)
		}
		if w.Nested.Name != "n" || w.Nested.Count != 11 {
			t.Errorf("a null for the whole nested struct must leave it alone: %+v", w.Nested)
		}
	})
}

// pointListStd, scoreMapStd, noCopyListStd and byteBlobStd are the methodless
// twins for the root types TestNullRootMatchesStdlib compares.
type pointListStd PointList

type scoreMapStd ScoreMap

type noCopyListStd NoCopyList

type byteBlobStd ByteBlob

// TestNullRootMatchesStdlib pins the *root* half of the same rule, which the
// generated UnmarshalJSON decides for itself: it intercepts a null document
// before calling the root decoder, so the decoder's own null arm — the one that
// nils a slice or map — never ran. A slice or map root therefore kept its
// previous contents (PointList{…}.UnmarshalJSON("null") left the element in
// place) while a *field* of that same named type was nil'd by the very arm the
// intercept skipped. A struct root is correctly a no-op and must stay one.
//
// The pre-existing conformance coverage only decoded null into an already-nil
// root, which passes either way; the seeded receiver below is what makes the bug
// visible, and the whole test fails against the unfixed generator except for the
// struct-root subtest.
func TestNullRootMatchesStdlib(t *testing.T) {
	null := []byte("null")

	t.Run("slice_root", func(t *testing.T) {
		l := PointList{{X: 1, Y: 2, Tag: "a"}}
		std := pointListStd(l)
		if err := json.Unmarshal(null, &std); err != nil {
			t.Fatalf("stdlib: %v", err)
		}
		if std != nil {
			t.Fatalf("stdlib premise changed: %v", std)
		}
		if err := l.UnmarshalJSON(null); err != nil {
			t.Fatalf("lightning: %v", err)
		}
		if l != nil {
			t.Errorf("null left a seeded slice root at %v, want nil as the stdlib gives", l)
		}
	})

	t.Run("map_root", func(t *testing.T) {
		m := ScoreMap{"a": 1}
		std := scoreMapStd(m)
		if err := json.Unmarshal(null, &std); err != nil {
			t.Fatalf("stdlib: %v", err)
		}
		if std != nil {
			t.Fatalf("stdlib premise changed: %v", std)
		}
		if err := m.UnmarshalJSON(null); err != nil {
			t.Fatalf("lightning: %v", err)
		}
		if m != nil {
			t.Errorf("null left a seeded map root at %v, want nil", m)
		}
	})

	t.Run("directive_and_byte_roots", func(t *testing.T) {
		// The //lightning:nocopy roots and the []byte root take the same path, so
		// they get the same reset.
		nl := NoCopyList{"a", "b"}
		nlStd := noCopyListStd(nl)
		if err := json.Unmarshal(null, &nlStd); err != nil {
			t.Fatalf("stdlib: %v", err)
		}
		if err := nl.UnmarshalJSON(null); err != nil {
			t.Fatalf("lightning: %v", err)
		}
		if nl != nil || nlStd != nil {
			t.Errorf("NoCopyList: lightning %v, stdlib %v; want both nil", nl, nlStd)
		}
		nm := NoCopyMap{"a": "b"}
		if err := nm.UnmarshalJSON(null); err != nil {
			t.Fatalf("lightning: %v", err)
		}
		if nm != nil {
			t.Errorf("NoCopyMap: %v, want nil", nm)
		}
		bb := ByteBlob{1, 2, 3}
		bbStd := byteBlobStd(bb)
		if err := json.Unmarshal(null, &bbStd); err != nil {
			t.Fatalf("stdlib: %v", err)
		}
		if err := bb.UnmarshalJSON(null); err != nil {
			t.Fatalf("lightning: %v", err)
		}
		if bb != nil || bbStd != nil {
			t.Errorf("ByteBlob: lightning %v, stdlib %v; want both nil", bb, bbStd)
		}
	})

	t.Run("struct_root_is_a_no_op", func(t *testing.T) {
		// The other half of the rule, and the reason the reset is not
		// unconditional: a null document leaves a struct root untouched.
		n := 7
		seed := PtrReuse{P: &Nested{Name: "keep", Count: 3}, N: &n}
		std := ptrReuseStd(seed)
		if err := json.Unmarshal(null, &std); err != nil {
			t.Fatalf("stdlib: %v", err)
		}
		if std.P == nil || std.P.Name != "keep" || std.N == nil {
			t.Fatalf("stdlib premise changed: %+v", std)
		}
		l := seed
		if err := l.UnmarshalJSON(null); err != nil {
			t.Fatalf("lightning: %v", err)
		}
		if l.P == nil || l.P.Name != "keep" || l.N == nil || *l.N != 7 {
			t.Errorf("null zeroed a struct root: %+v", l)
		}
	})
}

// TestLaxRejectsSyntaxErrors is the H3 regression. ",lax" documents that "only
// type mismatches are tolerated; genuinely malformed JSON (a syntax error in the
// value) still fails", and it did not: the skip after a failed decode was
// unstable.SkipValue, a bracket balancer, so every balanced-but-invalid value
// was dropped silently and the whole decode returned nil — which also punched a
// hole through the trailing-comma rejection the rest of the decoder enforces,
// and gave a host-dependent answer on some shapes, since SkipValue picks a SIMD
// or scalar balancer by CPU feature.
//
// Every case is run through StrictKinds too, which pins the premise that the
// input really is malformed rather than merely mistyped: both must fail. Against
// the unfixed generator all the balanced rows fail with "lax accepted"; the
// truncated ones passed even before, since SkipValue could not balance them
// either.
func TestLaxRejectsSyntaxErrors(t *testing.T) {
	cases := []struct {
		name string
		in   string
	}{
		{"slice_trailing_comma", `{"sl":[1,],"x":9}`},
		{"slice_missing_commas", `{"sl":[1 2 3],"x":9}`},
		{"slice_double_comma", `{"sl":[1,,2],"x":9}`},
		{"slice_truncated", `{"sl":[1,2`},
		{"slice_unbalanced", `{"sl":[1,2},"x":9}`},
		{"slice_bad_keyword", `{"sl":[tru],"x":9}`},
		// The interesting interaction: a type mismatch is what triggers the
		// skip, and the skip then has to notice the value is also malformed.
		{"slice_mismatch_then_garbage", `{"sl":["a"e],"x":9}`},
		{"slice_mismatch_trailing_comma", `{"sl":["a",],"x":9}`},
		{"struct_trailing_comma", `{"st":{"name":"a",},"x":9}`},
		{"struct_missing_colon", `{"st":{"name" "a"},"x":9}`},
		{"struct_unquoted_key", `{"st":{name:"a"},"x":9}`},
		{"struct_truncated", `{"st":{"name":"a"`},
		{"struct_bad_escape", `{"st":{"name":"a\q"},"x":9}`},
		{"struct_unterminated_string", `{"st":{"name":"a},"x":9}`},
		{"map_trailing_comma", `{"m":{"a":1,},"x":9}`},
		{"map_missing_value", `{"m":{"a":},"x":9}`},
		{"map_truncated", `{"m":{"a":1`},
		{"array_trailing_comma", `{"arr":[1,],"x":9}`},
		{"array_missing_commas", `{"arr":[1 2],"x":9}`},
		{"time_bad_escape", `{"t":"2020-01-02T03:04:05Z\q","x":9}`},
		{"time_mismatch_trailing_comma", `{"t":["a",],"x":9}`},
		{"scalar_bad_unicode_escape", `{"n":"\uZZZZ","x":9}`},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var s StrictKinds
			if err := s.UnmarshalJSON([]byte(c.in)); err == nil {
				t.Fatalf("premise: the non-lax decoder ACCEPTED %s, so it is not a syntax error", c.in)
			}
			var l LaxKinds
			if err := l.UnmarshalJSON([]byte(c.in)); err == nil {
				t.Errorf("lax accepted malformed JSON %s (decoded %+v); lax must swallow "+
					"type mismatches only", c.in, l)
			}
		})
	}
}

// TestLaxSwallowsTypeMismatches is the other half of the contract, and the
// reason the H3 fix could not simply propagate every error: a well-formed value
// of the wrong type — including one whose mismatch is buried inside an otherwise
// well-formed container, which is what makes lax worth having — must still be
// skipped, leaving the field at its zero value and the rest of the object
// decoding normally.
func TestLaxSwallowsTypeMismatches(t *testing.T) {
	cases := []struct {
		name string
		in   string
	}{
		{"slice_is_string", `{"sl":"nope","x":9}`},
		{"slice_is_object", `{"sl":{"a":1},"x":9}`},
		{"slice_element_is_string", `{"sl":["a"],"x":9}`},
		{"slice_element_is_array", `{"sl":[[1]],"x":9}`},
		{"slice_element_is_object", `{"sl":[{"a":1},2],"x":9}`},
		{"struct_is_number", `{"st":42,"x":9}`},
		{"struct_is_array", `{"st":[1,2],"x":9}`},
		{"struct_field_is_wrong_type", `{"st":{"name":5,"count":1},"x":9}`},
		{"struct_field_is_container", `{"st":{"name":{"deep":[1,2]}},"x":9}`},
		{"map_is_array", `{"m":[1,2],"x":9}`},
		{"map_value_is_string", `{"m":{"a":"b"},"x":9}`},
		{"map_value_is_object", `{"m":{"a":{"b":1}},"x":9}`},
		{"time_is_not_a_date", `{"t":"not a date","x":9}`},
		{"time_is_object", `{"t":{"a":1},"x":9}`},
		{"time_is_bool", `{"t":true,"x":9}`},
		{"array_is_string", `{"arr":"x","x":9}`},
		{"array_element_is_string", `{"arr":["a","b"],"x":9}`},
		{"scalar_is_string", `{"n":"7","x":9}`},
		{"scalar_is_object", `{"n":{"a":1},"x":9}`},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var s StrictKinds
			if err := s.UnmarshalJSON([]byte(c.in)); err == nil {
				t.Fatalf("premise: the non-lax decoder ACCEPTED %s, so it is not a mismatch", c.in)
			}
			var l LaxKinds
			if err := l.UnmarshalJSON([]byte(c.in)); err != nil {
				t.Fatalf("lax rejected a well-formed value of the wrong type %s: %v", c.in, err)
			}
			if l.X != 9 {
				t.Errorf("%s: decoding did not continue past the lax field (X=%d, want 9)", c.in, l.X)
			}
			want := LaxKinds{X: 9}
			if !reflect.DeepEqual(l, want) {
				t.Errorf("%s: lax left the field set: %+v", c.in, l)
			}
		})
	}
}

// The methodless twins for the embedded-Unmarshaler comparison. Stripping the
// GENERATED UnmarshalJSON is what makes encoding/json decode these by reflection
// instead of delegating to lightning (the trap CLAUDE.md records); the embedded
// field's own method still promotes to the twin, which is the stdlib behavior
// under test.
type (
	embedTimeStd   EmbedTime
	embedRawStd    EmbedRaw
	embedNumberStd EmbedNumber
)

// TestEmbeddedUnmarshalerDivergesFromStdlib pins a deliberate departure from the
// promotion parity the README otherwise claims: when the embedded type carries
// its own UnmarshalJSON, Go promotes that METHOD to the outer struct, so
// encoding/json treats the whole struct as a json.Unmarshaler and hands it the
// entire document — the sibling fields are never decoded. lightning promotes
// fields only: the embed becomes a named field keyed by its type name and the
// siblings decode normally.
//
// Written in the spirit of TestValidDivergesFromStdlib: neither side is asserted
// to be right, both are asserted to keep doing what they do, so the disagreement
// cannot rot into an undocumented one.
func TestEmbeddedUnmarshalerDivergesFromStdlib(t *testing.T) {
	// time.Time: its UnmarshalJSON demands a JSON string, so the stdlib fails the
	// whole document — a loud divergence.
	doc := []byte(`{"Time":"2021-01-02T03:04:05Z","a":7}`)
	var std embedTimeStd
	if err := json.Unmarshal(doc, &std); err == nil {
		t.Fatalf("stdlib premise changed: %#v decoded without error; the promoted "+
			"(*time.Time).UnmarshalJSON no longer hijacks the struct?", std)
	}
	var v EmbedTime
	if err := v.UnmarshalJSON(doc); err != nil {
		t.Fatalf("lightning: %v", err)
	}
	// v.Equal is the embedded time.Time's promoted method — the same promotion
	// that makes encoding/json treat the whole struct as a json.Unmarshaler.
	if v.A != 7 || !v.Equal(time.Date(2021, 1, 2, 3, 4, 5, 0, time.UTC)) {
		t.Errorf("got %+v, want the embed decoded as a named field and A=7", v)
	}

	// json.RawMessage: its UnmarshalJSON accepts anything, so the stdlib silently
	// swallows the whole document into the embedded field and leaves B zero — the
	// same divergence without an error to notice it by.
	rdoc := []byte(`{"RawMessage":{"z":1},"b":8}`)
	var rstd embedRawStd
	if err := json.Unmarshal(rdoc, &rstd); err != nil {
		t.Fatalf("stdlib: %v", err)
	}
	if string(rstd.RawMessage) != string(rdoc) || rstd.B != 0 {
		t.Fatalf("stdlib premise changed: got %+v, want the whole document in the embed", rstd)
	}
	var r EmbedRaw
	if err := r.UnmarshalJSON(rdoc); err != nil {
		t.Fatalf("lightning: %v", err)
	}
	if string(r.RawMessage) != `{"z":1}` || r.B != 8 {
		t.Errorf("got %+v, want RawMessage={\"z\":1} B=8", r)
	}

	// Control: through Go 1.26 json.Number has no unmarshaler of its own, so
	// nothing is promoted and the two agree — the divergence is the promoted
	// method, not the foreign embed. Go 1.27's json/v2-backed encoding/json gives
	// json.Number an UnmarshalJSONFrom, which promotes exactly like time.Time's
	// method above; on that toolchain the stdlib hands the whole document to the
	// embed and the case is a third instance of the divergence instead.
	ndoc := []byte(`{"Number":12,"c":9}`)
	var n EmbedNumber
	if err := n.UnmarshalJSON(ndoc); err != nil {
		t.Fatalf("lightning: %v", err)
	}
	if n.Number != "12" || n.C != 9 {
		t.Errorf("lightning %+v, want Number=12 C=9", n)
	}
	var nstd embedNumberStd
	if err := json.Unmarshal(ndoc, &nstd); err != nil {
		t.Logf("encoding/json rejects the document (%v): json.Number carries a promoted unmarshaler on this toolchain, so the embed diverges like time.Time", err)
		return
	}
	if n != EmbedNumber(nstd) {
		t.Errorf("lightning %+v, stdlib %+v: embedding a type without an unmarshaler must agree", n, nstd)
	}
}

// TestUnwrapRejectsTrailingContent locks the unwrap option's "parsed as a fresh
// input" contract at the end of the body as well as the start: content after the
// wrapped document's top-level value is ErrInvalidJSON, exactly as it is for the
// same bytes decoded as a root. Before the check the wrapped decode stopped at
// the first value and ignored the rest, so `"{\"count\":7} garbage"` decoded
// happily while the identical root document did not.
func TestUnwrapRejectsTrailingContent(t *testing.T) {
	bad := []string{
		`{"count":7} trailing garbage`,
		`{"count":7}{"count":8}`,
		`{"count":7} null`,
		`{"count":7},`,
		`{"count":7} }`,
	}
	for _, body := range bad {
		// The root decode of the same bytes is the reference behavior.
		var root UnwrapRoot
		rootErr := root.UnmarshalJSON([]byte(body))
		if !errors.Is(rootErr, unstable.ErrInvalidJSON) {
			t.Fatalf("root decode of %q: err = %v, want ErrInvalidJSON (reference behavior changed)", body, rootErr)
		}
		var u UnwrapTrailing
		err := u.UnmarshalJSON([]byte(wrap(body)))
		if !errors.Is(err, unstable.ErrInvalidJSON) {
			t.Errorf("unwrap of %q: err = %v, want ErrInvalidJSON like the root decode", body, err)
		}
	}

	// Trailing (and leading) whitespace is not trailing content — the root
	// accepts it, so the wrapped body must too, and the whole document decodes.
	for _, body := range []string{`{"count":7}`, "  {\"count\":7}\t\n ", "\n{\"count\":7}"} {
		var root UnwrapRoot
		if err := root.UnmarshalJSON([]byte(body)); err != nil {
			t.Fatalf("root decode of %q: %v (reference behavior changed)", body, err)
		}
		var u UnwrapTrailing
		if err := u.UnmarshalJSON([]byte(wrap(body))); err != nil {
			t.Errorf("unwrap of %q: %v", body, err)
			continue
		}
		if u.W.Count != 7 || u.X != 9 {
			t.Errorf("unwrap of %q: got %+v, want W.Count=7 X=9", body, u)
		}
	}

	// The empty and all-whitespace bodies keep their own rule — no value, no
	// error, field left zero (TestUnwrapWhitespaceBody covers the pointer form).
	for _, body := range []string{``, `   `, "\t\n"} {
		var u UnwrapTrailing
		if err := u.UnmarshalJSON([]byte(wrap(body))); err != nil {
			t.Errorf("unwrap of empty body %q: %v", body, err)
			continue
		}
		if u.W != (Nested{}) || u.X != 9 {
			t.Errorf("unwrap of empty body %q: got %+v, want the zero field and X=9", body, u)
		}
	}

	// unwrap + lax: the check must not turn lax into strict. A wrapped value of
	// the wrong type is still swallowed (field unset, rest of the object
	// decoded); trailing content is malformed JSON, which lax has never
	// tolerated, so it still fails.
	var lax UnwrapTrailing
	if err := lax.UnmarshalJSON([]byte(wrapKey("l", `[1,2]`))); err != nil {
		t.Errorf("lax unwrap of a mistyped value: %v, want it swallowed", err)
	} else if lax.L != (Nested{}) || lax.X != 9 {
		t.Errorf("lax unwrap of a mistyped value: got %+v, want the zero field and X=9", lax)
	}
	for _, body := range []string{`[1,2] trailing`, `{"count":7} trailing`} {
		var u UnwrapTrailing
		if err := u.UnmarshalJSON([]byte(wrapKey("l", body))); !errors.Is(err, unstable.ErrInvalidJSON) {
			t.Errorf("lax unwrap of %q: err = %v, want ErrInvalidJSON (lax tolerates mismatches, not malformed JSON)", body, err)
		}
	}
}

// wrap builds the outer document whose "w" member is body carried as a JSON
// string, the shape the unwrap option reads.
func wrap(body string) string { return wrapKey("w", body) }

// wrapKey is wrap for a chosen member, so the same bodies can be aimed at the
// plain unwrap field and at the lax one.
func wrapKey(key, body string) string {
	q, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	return `{"` + key + `":` + string(q) + `,"x":9}`
}
