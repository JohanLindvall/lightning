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

// ArenaDoc carries //lightning:arena: the presized backings of its small numeric
// slice fields are carved from per-decode arena chunks by the Decode*SliceArena
// readers instead of allocated one make() per slice. Its shape covers every carve
// path: several small slices per element (adjacent carves in one chunk), int and
// uint element kinds narrower than 8 bytes (carve-size rounding), a slice big
// enough to exceed the carve threshold (direct make fallback), plus null/empty
// arrays and non-slice fields riding along. ArenaKey is a *named* element struct
// so a decoder under the arena variant threads the arena through an intermediate
// slice-of-named-struct decoder. Exercised by TestArenaDirective.
//
//lightning:arena
type ArenaDoc struct {
	Keys []ArenaKey `json:"keys"`
	Name string     `json:"name"`
}

type ArenaKey struct {
	Pos  []float64 `json:"pos"`
	Rot  []int32   `json:"rot"`
	Cnt  []uint16  `json:"cnt"`
	Big  []float64 `json:"big"`
	Time float64   `json:"time"`
}

// ArenaTree combines //lightning:arena with a recursive schema, so its decoders
// thread *both* the depth counter and the arena pointer — the two extra
// parameters must agree in order between every signature and call site, which
// only a type needing both exercises. Decoded in TestArenaDirective.
//
//lightning:arena
type ArenaTree struct {
	Vals []float64    `json:"vals"`
	Kids []*ArenaTree `json:"kids"`
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

// LaxSharedDestructive and LaxSharedPlain share the exact same ",nocopy,lax"
// string field type, with the destructive root declared FIRST: before the
// valueDecoder memo key carried the per-root prefix and directive marker, both
// roots shared whichever lax value decoder was generated first, so the plain
// root inherited the destructive in-place unescape and silently mutated the
// caller's buffer. Exercised by TestLaxDecoderIsolation.
//
//lightning:destructive
type LaxSharedDestructive struct {
	S string `json:"s,nocopy,lax"`
}

type LaxSharedPlain struct {
	S string `json:"s,nocopy,lax"`
}

// LaxSharedCompact and LaxSharedAnyPlain cover the same leak for the compact
// directive: pre-fix the shared lax `any` decoder called DecodeValueCompact, so
// the plain root decoding whitespaced input had lax swallow the error and left
// V nil. Exercised by TestLaxDecoderIsolation.
//
//lightning:compact
type LaxSharedCompact struct {
	V any `json:"v,lax"`
}

type LaxSharedAnyPlain struct {
	V any `json:"v,lax"`
}

// CompactShared is the type both roots below reach. It holds one of every
// container the generator emits an inter-token SkipWS into — an object (itself),
// a generated slice loop, a map loop, a fixed-size array loop — plus a dynamic
// `any` field, which is the one field kind whose *reader* has a compact variant
// (unstable.DecodeValueCompact). //lightning:compact elides the skip at every
// one of those sites, so a decoder built under the directive must reject
// whitespace at each of them.
//
// The slice and array element types are deliberately strings, and the map value
// an int reached through the map loop, rather than a []int/[N]int field: those
// route to the batched pkg/unstable readers, which do their own whitespace
// skipping and stay whitespace-tolerant under the directive. A batched field
// here would silently weaken TestCompactDirective's rejection assertions.
type CompactShared struct {
	Name string         `json:"name"`
	Vals []string       `json:"vals"`
	M    map[string]int `json:"m"`
	A    [2]string      `json:"a"`
	Any  any            `json:"any"`
}

// CompactDoc carries //lightning:compact: every inter-token SkipWS in its
// decoders — and in the decoders of every type it reaches — is elided at compile
// time, so it decodes whitespace-free input and *rejects* input with whitespace
// between tokens. CompactPlain below is the same shape without the directive,
// and both reach the same CompactShared, so the pair pins the two halves of the
// directive at once: the elision itself, and the fact that the compact and plain
// decoders for a shared type stay distinct (memo keys carry g.prefix +
// g.cmark()). It is declared FIRST, like LaxSharedDestructive above, so a memo
// key that lost that distinction would generate the compact variant first and
// hand it to the plain root.
//
// Exercised by TestCompactDirective.
//
//lightning:compact
type CompactDoc struct {
	N     CompactShared   `json:"n"`
	Items []CompactShared `json:"items"`
}

// CompactPlain is CompactDoc without the directive: same fields, same shared
// element type, ordinary whitespace-skipping decoders. It is the control for
// every rejection TestCompactDirective asserts of CompactDoc — without it, a
// "compact rejects whitespace" test could pass because the input was malformed
// for some unrelated reason.
type CompactPlain struct {
	N     CompactShared   `json:"n"`
	Items []CompactShared `json:"items"`
}

// NoCopyList is the slice half of //lightning:nocopy (NoCopyMap above is the map
// half): a slice root whose string elements alias the input instead of being
// copied. Exercised by TestNoCopyDirective.
//
//lightning:nocopy
type NoCopyList []string

// RawNullDoc pins json.RawMessage's treatment of a JSON null to encoding/json's:
// RawMessage implements json.Unmarshaler, and Unmarshal calls UnmarshalJSON
// "including when the input is a JSON null", so the field ends up holding the
// four bytes "null" — it is NOT left untouched. The generator used to skip the
// assignment for a null, which was invisible on a fresh target (nil either way)
// and wrong on a reused one, where the field silently kept the previous
// document's value. Exercised by TestRawMessageNull.
type RawNullDoc struct {
	R    json.RawMessage   `json:"r"`
	NC   json.RawMessage   `json:"nc,nocopy"`
	Many []json.RawMessage `json:"many"`
}

// UnwrapPtr covers the unwrap tag on a *pointer* field: the pointer null probe
// in the generated inner decoder reads data[i] unguarded (safe everywhere else
// because the enclosing loop guarantees a value byte), so the unwrap closure
// must bail out before it on a body that is entirely whitespace. Exercised by
// TestUnwrapWhitespaceBody.
type UnwrapPtr struct {
	P *Nested `json:"p,unwrap"`
	S *string `json:"s,unwrap"`
}

// PtrReuse exercises decoding into a reused target with non-nil pointer
// fields: the generated allocation is guarded (if p == nil { p = new(T) }), so
// a non-nil pointee is decoded into rather than replaced — encoding/json's
// documented pointer semantics, and zero allocations on reuse. Exercised by
// TestPointerFieldReuse.
type PtrReuse struct {
	P *Nested `json:"p"`
	N *int    `json:"n"`
}

// ByteSliceDoc covers []byte fields, which follow encoding/json: a base64
// string (the form the stdlib marshals) or a numeric array both decode, null
// is nil, and a fixed-size [N]byte stays numeric-only. Exercised by
// TestByteSliceStdlibParity.
type ByteSliceDoc struct {
	B     []byte   `json:"b"`
	Fixed [3]byte  `json:"fixed"`
	Many  [][]byte `json:"many"`
}

// ByteBlob is a named byte-slice root type; base64 applies to it exactly as to
// []byte (encoding/json keys on the underlying type). A named slice type is
// only decodable as a root, so it is tested as one.
type ByteBlob []byte

// NullDoc pins encoding/json's rule for an explicit JSON null: "unmarshaling a
// JSON null into any other Go type has no effect on the value and produces no
// error". The rule is only observable on a target that is not already zero — the
// seeded or reused decode target this library encourages — which is why
// TestNullFieldsMatchStdlib decodes into a pre-filled value rather than a fresh
// one. The emitted code got it wrong because the *OrNull readers signal a null
// by returning the ZERO value with a nil error, and the assignment was
// unconditional, so every scalar field was wiped instead of kept.
//
// One field of every kind is present, because the rule splits by kind: a scalar,
// json.Number, time.Time, a struct and a fixed-size array are left ALONE, while a
// slice, map, pointer and any become nil and a json.RawMessage takes the four
// bytes "null". The nocopy spellings ride along because they use different
// readers.
type NullDoc struct {
	Str    string          `json:"str"`
	StrNC  string          `json:"strNC,nocopy"`
	Bool   bool            `json:"bool"`
	I      int             `json:"i"`
	I8     int8            `json:"i8"`
	I64    int64           `json:"i64"`
	U      uint            `json:"u"`
	U16    uint16          `json:"u16"`
	F32    float32         `json:"f32"`
	F64    float64         `json:"f64"`
	Num    json.Number     `json:"num"`
	NumNC  json.Number     `json:"numNC,nocopy"`
	Time   time.Time       `json:"time"`
	Nested Nested          `json:"nested"`
	Arr    [3]int          `json:"arr"`
	Sl     []int           `json:"sl"`
	Strs   []string        `json:"strs"`
	M      map[string]int  `json:"m"`
	P      *int            `json:"p"`
	A      any             `json:"a"`
	R      json.RawMessage `json:"r"`
}

// NullLaxDoc is NullDoc's ",lax" twin. A lax field decodes into a fresh scratch
// value and then commits it, so the null rule has to be applied a second time at
// that commit — and per kind, since committing is right for the kinds a null
// assigns (slice, map, pointer, any, RawMessage) and wrong for the kinds it
// leaves alone. Without that, json:"i" and json:"i,lax" disagreed with each
// other on {"i":null}: the plain field kept its value and the lax one was
// zeroed. Exercised by TestNullFieldsMatchStdlib.
type NullLaxDoc struct {
	Str    string          `json:"str,lax"`
	I      int             `json:"i,lax"`
	F64    float64         `json:"f64,lax"`
	Num    json.Number     `json:"num,lax"`
	Time   time.Time       `json:"time,lax"`
	Nested Nested          `json:"nested,lax"`
	Arr    [3]int          `json:"arr,lax"`
	Sl     []int           `json:"sl,lax"`
	M      map[string]int  `json:"m,lax"`
	P      *int            `json:"p,lax"`
	A      any             `json:"a,lax"`
	R      json.RawMessage `json:"r,lax"`
}

// LaxKinds and StrictKinds are the same shape with and without ",lax" on every
// field, so one input can be run through both. That pairing is what makes the
// lax contract testable in both directions: a *type mismatch* must be swallowed
// by lax and rejected by strict, while a *syntax error* must be rejected by
// both.
//
// The second half was false until the lax skip became unstable.SkipValueStrict.
// It used to be unstable.SkipValue, a bracket balancer, so any balanced but
// invalid value — [1,] and [1 2 3] and [1,,2] — was silently dropped and the whole
// decode returned nil, punching a hole through the trailing-comma rejection the
// rest of the decoder enforces (and giving a host-dependent answer on some
// shapes, since SkipValue picks a SIMD or scalar balancer by CPU feature). The
// field kinds are chosen to cover every branch of the lax value decoder: a
// slice, a named struct, a map, a time.Time (whose lax reader is a different
// reader, not just a different error path), a fixed-size array routed through
// the batched reader, and a bare scalar. X follows the lax field in every test
// document so a test can tell "swallowed and kept decoding" from "failed".
type LaxKinds struct {
	Sl  []int          `json:"sl,lax"`
	St  Nested         `json:"st,lax"`
	M   map[string]int `json:"m,lax"`
	T   time.Time      `json:"t,lax"`
	Arr [2]int         `json:"arr,lax"`
	N   int            `json:"n,lax"`
	X   int            `json:"x"`
}

type StrictKinds struct {
	Sl  []int          `json:"sl"`
	St  Nested         `json:"st"`
	M   map[string]int `json:"m"`
	T   time.Time      `json:"t"`
	Arr [2]int         `json:"arr"`
	N   int            `json:"n"`
	X   int            `json:"x"`
}

// LaxArrays covers fixed-size scalar arrays behind the lax tag option: the
// value decoder routes them through the batched pkg/unstable array readers
// (the same ones field-level [N]T uses), and lax semantics still apply — a
// mistyped value is skipped, leaving the zero array. Exercised by
// TestLaxFixedArrays.
type LaxArrays struct {
	F [3]float64 `json:"f,lax"`
	I [2]int32   `json:"i,lax"`
	U [4]uint16  `json:"u,lax"`
}

// EmbedTime and EmbedRaw embed a type that carries its own UnmarshalJSON, which
// is where lightning's field promotion parts company with encoding/json. Go
// promotes the embedded method to the OUTER struct, so encoding/json sees the
// whole struct as a json.Unmarshaler and hands it the entire document — the
// sibling fields are never looked at. lightning promotes fields, not methods: the
// embed decodes as a named field keyed by its type name and the siblings decode
// normally. Both are pinned by TestEmbeddedUnmarshalerDivergesFromStdlib.
type EmbedTime struct {
	time.Time
	A int `json:"a"`
}

// EmbedRaw is the same shape with json.RawMessage, whose promoted UnmarshalJSON
// succeeds rather than failing — so the stdlib silently swallows the whole
// document into the embedded field instead of reporting anything.
type EmbedRaw struct {
	json.RawMessage
	B int `json:"b"`
}

// EmbedNumber is the control: json.Number is a defined string type with NO
// UnmarshalJSON of its own, so nothing is promoted and both decoders treat the
// embed as a named field. It is what keeps the divergence above attributable to
// the promoted method rather than to embedding a foreign type as such.
type EmbedNumber struct {
	json.Number
	C int `json:"c"`
}

// UnwrapTrailing covers the unwrap option's trailing-content rule: the wrapped
// string holds a whole JSON document, so content after its top-level value is an
// error — the check a root UnmarshalJSON makes. The generated closure used to
// stop at the end of the first value and ignore the rest. L pairs the option
// with lax, whose contract the check has to respect from the inside: a wrapped
// value of the wrong TYPE is still swallowed, while trailing content is
// malformed JSON and so still fails. Exercised by
// TestUnwrapRejectsTrailingContent.
type UnwrapTrailing struct {
	W Nested `json:"w,unwrap"`
	L Nested `json:"l,unwrap,lax"`
	X int    `json:"x"`
}

// UnwrapRoot has Nested's shape but is referenced by nothing, so it is a root
// with its own UnmarshalJSON: the same bytes decoded directly, which is the
// behavior TestUnwrapRejectsTrailingContent holds the wrapped decode to.
type UnwrapRoot struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}
