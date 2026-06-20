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

	// Slices: scalar, string, float, slice-of-slice, and struct elements.
	Ints   []int     `json:"ints"`
	Strs   []string  `json:"strs"`
	Floats []float64 `json:"floats"`
	Grid   [][]int   `json:"grid"`
	Items  []Nested  `json:"items"`

	// Fixed-size array.
	Arr [3]int `json:"arr"`

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
