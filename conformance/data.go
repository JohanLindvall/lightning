// Package conformance holds a single struct that exercises every field type the
// lightning generator supports, plus the field-tag options (nocopy, lax, unwrap,
// pipe-separated alternate names, and the "-" skip). data_unmarshal.go is the
// generated decoder; regenerate it with go generate.
package conformance

import (
	"encoding/json"
	"time"
)

//go:generate go run github.com/JohanLindvall/lightning data.go

// Nested is a named struct used as a field, a slice element, a map value, and a
// pointee, so the generator's named-struct decoder is exercised in each position.
type Nested struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

// Doc covers all supported field types. The matching test.json fills every field
// and conformance_test.go asserts the decoded value of each.
type Doc struct {
	// Strings.
	Str        string `json:"str"`
	StrNoCopy  string `json:"strNoCopy,nocopy"`
	StrUnicode string `json:"strUnicode"` // exercises \uXXXX and escapes

	// Bool.
	Bool bool `json:"bool"`

	// Signed integers, every width.
	I   int   `json:"i"`
	I8  int8  `json:"i8"`
	I16 int16 `json:"i16"`
	I32 int32 `json:"i32"`
	I64 int64 `json:"i64"`
	R   rune  `json:"r"`

	// Unsigned integers, every width.
	U    uint    `json:"u"`
	U8   uint8   `json:"u8"`
	U16  uint16  `json:"u16"`
	U32  uint32  `json:"u32"`
	U64  uint64  `json:"u64"`
	Uptr uintptr `json:"uptr"`
	B    byte    `json:"b"`

	// Floats.
	F32 float32 `json:"f32"`
	F64 float64 `json:"f64"`

	// json.Number, by value and pointer.
	Num  json.Number  `json:"num"`
	PNum *json.Number `json:"pnum"`

	// Time: RFC 3339 and the lax (Unix-timestamp / space-separated) parser.
	Time    time.Time `json:"time"`
	TimeLax time.Time `json:"timeLax,lax"`

	// Raw JSON, copied and aliased.
	Raw       json.RawMessage `json:"raw"`
	RawNoCopy json.RawMessage `json:"rawNoCopy,nocopy"`

	// Nested structs: named and anonymous.
	Nested Nested `json:"nested"`
	Anon   struct {
		X int    `json:"x"`
		Y string `json:"y"`
	} `json:"anon"`

	// Slices: scalar, string, float, slice-of-slice, struct, and
	// pointer-to-struct elements (the last also exercises the []*T presize).
	Ints     []int     `json:"ints"`
	Strs     []string  `json:"strs"`
	Floats   []float64 `json:"floats"`
	Grid     [][]int   `json:"grid"`
	Items    []Nested  `json:"items"`
	PtrItems []*Nested `json:"ptrItems"`

	// Fixed-size arrays: signed and unsigned element kinds (the latter routes
	// to the batched DecodeUintArray).
	Arr  [3]int    `json:"arr"`
	UArr [3]uint32 `json:"uarr"`

	// Maps with string keys, scalar and slice values.
	M  map[string]int      `json:"m"`
	MM map[string][]uint64 `json:"mm"`

	// Pointers to scalar and struct.
	PI *int    `json:"pi"`
	PN *Nested `json:"pn"`

	// interface{}/any, scalar and composite.
	Any  any   `json:"any"`
	Anys []any `json:"anys"`

	// Pipe-separated alternate names: either key fills this field.
	Status int `json:"status|EdgeStatus"`

	// Skipped: never read from JSON regardless of input.
	Ignored int `json:"-"`

	// unwrap: the JSON value is a string whose contents are themselves JSON.
	Embedded Nested `json:"embedded,unwrap"`

	// lax on a non-time field: a type mismatch (test.json gives a string where an
	// int is expected) is skipped, leaving the field at its zero value instead of
	// failing the whole decode.
	LaxSkip int `json:"laxSkip,lax"`
}

// PointList is a named slice root type: the generator emits an UnmarshalJSON on
// it directly (array-root JSON), decoding via the element decoder. Exercised by
// TestSliceRoot.
type PointList []struct {
	X   int    `json:"x"`
	Y   int    `json:"y"`
	Tag string `json:"tag"`
}

// ScoreMap is a named map root type: the generator emits UnmarshalJSON on it
// directly (object-root JSON that's a data map). Exercised by TestMapRoot.
type ScoreMap map[string]int

//lightning:nocopy
type NoCopyMap map[string]string

// DestructiveDoc carries //lightning:destructive: its nocopy string fields unescape
// escaped values into the input buffer (aliasing and mutating it) rather than
// allocating. Exercised by TestDestructiveDirective.
//
//lightning:destructive
type DestructiveDoc struct {
	Name string   `json:"name,nocopy"`
	Tags []string `json:"tags,nocopy"`
}

// Tree is a self-referential type: its decoder and the decoder for its []*Tree
// field call each other, so decoding recurses once per level of the document.
// Because the generator detects the cycle, both carry a depth counter and refuse
// to descend past unstable.MaxDepth — without which deeply nested input would
// exhaust the goroutine stack, a fatal error recover cannot catch. Exercised by
// TestRecursiveTypeDepthLimit.
type Tree struct {
	Name string  `json:"name"`
	Kids []*Tree `json:"kids"`
}

// RingRoot is the entry point for the mutually recursive pair below. Ring1 and
// Ring2 each reference the other, so both count as "nested in another type" and
// neither gets its own UnmarshalJSON — a cycle with no member outside it needs a
// root above it. RingRoot also checks that a type which merely *reaches* a cycle
// threads the depth counter down into it. Exercised by
// TestMutuallyRecursiveTypeDepthLimit.
type RingRoot struct {
	Start *Ring1 `json:"start"`
}

// Ring1 and Ring2 form a two-type cycle rather than a self-reference, so the
// generator's cycle search must follow an edge before the cycle closes.
type Ring1 struct {
	Name string `json:"name"`
	Next *Ring2 `json:"next"`
}

type Ring2 struct {
	Count int    `json:"count"`
	Back  *Ring1 `json:"back"`
}

// LongKeys exercises the `switch len(key)` dispatch the generator emits once a
// struct has a JSON name longer than 16 bytes (cmd/compile's inline-comparison
// limit), where names are matched in <=16-byte chunks so no comparison calls
// runtime.memequal. Every tricky case for that scheme is represented:
//
//   - names of exactly 16, 17 and 33+ bytes (one, two and three chunks);
//   - SharedPrefixA/B, whose names are the same length AND share their whole first
//     16-byte chunk, so a bug that compares only the first chunk swaps them;
//   - two distinct names of the same length in one bucket;
//   - Alt, whose pipe-separated names straddle the 16-byte boundary (one bucket
//     each), the case that duplicates a field's decode code;
//   - Short, so a bucket that keeps its nested `switch key` is covered too.
//
// Exercised by TestLongKeyDispatch.
type LongKeys struct {
	Short            int    `json:"s"`
	Exactly16Bytes__ int    `json:"exactly16bytes__"`
	SeventeenBytes_1 int    `json:"seventeenBytes_17"`
	SharedPrefixA    int    `json:"sharedPrefix16xxA"`
	SharedPrefixB    int    `json:"sharedPrefix16xxB"`
	SameLenOther     int    `json:"differentButSame"`
	ThirtyThreePlus  string `json:"aKeyOfThirtyThreeBytesExactly_333"`
	Alt              int    `json:"shortAlt|aVeryMuchLongerAlternateName"`
}
