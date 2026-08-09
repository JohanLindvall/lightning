package conformance

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"
	"unsafe"
)

// compactDocStd and compactPlainStd strip the generated UnmarshalJSON so
// encoding/json decodes these by reflection. See TestStdlibTwinsAreReflectionOnly
// for why every stdlib differential in this package must go through such a twin.
type (
	compactDocStd   CompactDoc
	compactPlainStd CompactPlain
)

// compactSample is a whitespace-free document that fills CompactDoc/CompactPlain
// completely: every container the generator emits an inter-token SkipWS into is
// present and non-empty (nested object, generated slice loop, map loop,
// fixed-array loop, dynamic `any` object and array), plus one empty map to cover
// the loop-top-closer branch. Every scalar token is a single character, so a
// space inserted anywhere outside a string literal is an inter-token space and
// never splits a token.
const compactSample = `{"n":{"name":"a","vals":["x","y"],"m":{"k":1},"a":["p","q"],"any":{"q":1}},` +
	`"items":[{"name":"b","vals":["z"],"m":{},"a":["r","s"],"any":[1,2]}]}`

// TestCompactDirective is //lightning:compact's only executed coverage. The
// directive works by compile-time elision — g.skipWS emits nothing between
// tokens, and g.anyValue picks unstable.DecodeValueCompact — which changes
// emitted code at every call site in every decoder the root reaches, so a
// regression there is broad and silent. Nothing in the repository ran against it
// before this test: the only compact schema outside conformance is
// bench/cloudflare-compact, and bench/ is a separate module every runner enters
// with -run='^$'.
//
// The sweep below covers both mechanisms, and separates them: of its 77
// positions, 69 are guarded by the skipWS elision and the 8 inside the `any`
// values by the DecodeValueCompact selection (measured by disabling each in
// turn), so neither can regress behind the other.
//
// The load-bearing assertion is the *rejection* sweep, not the decode. Decoding
// compact input successfully proves nothing — a plain, whitespace-skipping
// decoder does that too. So the test inserts a single space at every position of
// compactSample that lies outside a string literal (i.e. at every inter-token
// boundary) and requires that
//
//	CompactDoc  — the //lightning:compact root — REJECTS it, and
//	CompactPlain — the same shape without the directive — ACCEPTS it.
//
// Both halves are needed, and together they give the test two distinct failure
// modes:
//
//   - If the elision stops happening (g.skipWS emits the whitespace block
//     regardless of g.compact, or the directive stops being parsed/plumbed into
//     g.compact), CompactDoc becomes whitespace-tolerant and every rejection
//     fails.
//   - If the compact and plain decoders for the shared CompactShared type
//     collapse onto one another — the memo key at each emitter is
//     g.prefix + g.cmark() + kind, and losing either component merges them —
//     then one root inherits the other's semantics: whichever variant is
//     generated first wins, so either CompactDoc accepts whitespace inside the
//     shared type or CompactPlain rejects it. One of the two halves fails in
//     either direction, which is why CompactPlain is not merely decorative.
//
// Leading and trailing whitespace around the whole document is deliberately not
// swept: the generated UnmarshalJSON trims the document ends with
// unstable.SkipWS even under the directive (only *inter-token* skips are
// elided), so those two positions are accepted by design.
func TestCompactDirective(t *testing.T) {
	// The compact document itself must decode, and decode to exactly what
	// encoding/json produces for the same shape.
	var got CompactDoc
	if err := got.UnmarshalJSON([]byte(compactSample)); err != nil {
		t.Fatalf("compact root rejected compact input: %v", err)
	}
	var want compactDocStd
	if err := json.Unmarshal([]byte(compactSample), &want); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, CompactDoc(want)) {
		t.Fatalf("compact decode diverges from encoding/json:\n got %#v\nwant %#v", got, want)
	}

	// The plain twin decodes the same input to the same value: the two roots
	// differ only in whitespace tolerance.
	var plain CompactPlain
	if err := plain.UnmarshalJSON([]byte(compactSample)); err != nil {
		t.Fatalf("plain root rejected compact input: %v", err)
	}
	if !reflect.DeepEqual(CompactShared(plain.N), CompactShared(got.N)) ||
		!reflect.DeepEqual(plain.Items, got.Items) {
		t.Fatalf("plain and compact roots disagree on compact input:\n compact %#v\n plain   %#v", got, plain)
	}

	// The sweep.
	points := interTokenPositions(compactSample)
	if len(points) < 40 {
		t.Fatalf("sweep covers only %d positions; compactSample lost its shape", len(points))
	}
	for _, pos := range points {
		in := compactSample[:pos] + " " + compactSample[pos:]

		// Premise: the insertion produced a still-valid JSON document. A space
		// between tokens cannot make one invalid, so a failure here means the
		// position was misclassified, not that lightning is wrong.
		if !json.Valid([]byte(in)) {
			t.Fatalf("position %d: inserted space produced invalid JSON %q", pos, in)
		}

		var c CompactDoc
		if err := c.UnmarshalJSON([]byte(in)); err == nil {
			t.Errorf("compact root ACCEPTED whitespace at %d: %q\n"+
				"\t(//lightning:compact is not in effect there — the SkipWS elision, "+
				"or DecodeValueCompact inside an `any`)", pos, in)
		}
		var p CompactPlain
		if err := p.UnmarshalJSON([]byte(in)); err != nil {
			t.Errorf("plain root REJECTED whitespace at %d: %q: %v\n"+
				"\t(the plain root inherited compact semantics)", pos, in, err)
		}
	}

	// The other side of the boundary the sweep skips: the generated
	// UnmarshalJSON trims the document's own ends with unstable.SkipWS whatever
	// the directive says, so whitespace *around* a compact document is fine.
	// Pinned rather than merely commented, since the sweep's exclusion of
	// positions 0 and len(doc) is only sound while this holds.
	var ends CompactDoc
	if err := ends.UnmarshalJSON([]byte(" \t\n" + compactSample + "\n ")); err != nil {
		t.Errorf("compact root rejected whitespace around the document: %v", err)
	} else if !reflect.DeepEqual(ends, got) {
		t.Errorf("document-end whitespace changed the compact decode:\n got %#v\nwant %#v", ends, got)
	}
}

// interTokenPositions returns every index of a whitespace-free JSON document at
// which a space may be inserted without landing inside a string literal, and
// excluding the two document ends (which the generated UnmarshalJSON trims with
// unstable.SkipWS even in compact mode). Because every scalar token in the input
// it is used on is a single character, every position it returns is a genuine
// inter-token boundary.
//
// It assumes the document holds no backslash escapes (compactSample has none);
// an escaped quote would need real string scanning, so it panics rather than
// silently misclassifying positions.
func interTokenPositions(doc string) []int {
	if strings.ContainsRune(doc, '\\') {
		panic("interTokenPositions: input must not contain escapes")
	}
	var out []int
	inString := false
	for i := 0; i < len(doc); i++ {
		if doc[i] == '"' {
			// Inserting at an opening quote is outside the string; at a closing
			// quote it would land before the quote, i.e. inside the body.
			if !inString {
				if i != 0 {
					out = append(out, i)
				}
			}
			inString = !inString
			continue
		}
		if inString || i == 0 {
			continue
		}
		out = append(out, i)
	}
	return out
}

// TestNoCopyListDirective covers the slice half of //lightning:nocopy —
// TestNoCopyDirective covers the map half. Under the directive a slice root's
// string elements alias the input rather than being copied, which is only
// observable through the element's data pointer.
func TestNoCopyListDirective(t *testing.T) {
	data := []byte(`["alpha","beta","gamma"]`)
	var l NoCopyList
	if err := l.UnmarshalJSON(data); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(l, NoCopyList{"alpha", "beta", "gamma"}) {
		t.Fatalf("got %#v", l)
	}
	base := uintptr(unsafe.Pointer(&data[0]))
	for _, s := range l {
		p := uintptr(unsafe.Pointer(unsafe.StringData(s)))
		if p < base || p >= base+uintptr(len(data)) {
			t.Errorf("element %q does not alias input (not nocopy)", s)
		}
	}
}

// rawNullDocStd strips RawNullDoc's generated UnmarshalJSON for the reflection
// baseline. json.RawMessage keeps its own UnmarshalJSON, which is the point:
// that method *is* encoding/json's null semantics, and it is what lightning has
// to match.
type rawNullDocStd RawNullDoc

// TestRawMessageNull pins json.RawMessage's null handling to encoding/json's.
// RawMessage implements json.Unmarshaler, and Unmarshal calls UnmarshalJSON
// "including when the input is a JSON null" (literalStore dispatches to the
// Unmarshaler before it ever looks at the literal), so the field ends up holding
// the four bytes "null".
//
// The generator used to skip the assignment entirely when the value started with
// 'n'. On a fresh target that is invisible — a nil RawMessage either way — which
// is why it survived; it is only visible on a *reused* target, where the field
// kept the previous document's value, and in a slice, where a null element came
// out nil while the stdlib gives "null". Both are checked below.
func TestRawMessageNull(t *testing.T) {
	for _, in := range []string{
		`{"r":null}`,
		`{"r":null,"nc":null}`,
		`{"r":{"a":1},"nc":[1,2],"many":[null,{"b":2},null]}`,
		`{"many":[null]}`,
		`{}`, // absent key: both leave the field alone
	} {
		var std rawNullDocStd
		if err := json.Unmarshal([]byte(in), &std); err != nil {
			t.Fatalf("%s: stdlib: %v", in, err)
		}
		var l RawNullDoc
		if err := l.UnmarshalJSON([]byte(in)); err != nil {
			t.Fatalf("%s: lightning: %v", in, err)
		}
		if !reflect.DeepEqual(RawNullDoc(std), l) {
			t.Errorf("%s:\n stdlib    %#v\n lightning %#v", in, std, l)
		}
	}

	// Reuse is where the old behaviour was observable: the second document's
	// null must overwrite the first document's value, not be ignored.
	doc1, doc2 := []byte(`{"r":{"a":1},"nc":"keep"}`), []byte(`{"r":null,"nc":null}`)

	var std rawNullDocStd
	if err := json.Unmarshal(doc1, &std); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(doc2, &std); err != nil {
		t.Fatal(err)
	}
	if string(std.R) != "null" || string(std.NC) != "null" {
		t.Fatalf("stdlib premise changed: R=%q NC=%q, want both \"null\"", std.R, std.NC)
	}

	var l RawNullDoc
	if err := l.UnmarshalJSON(doc1); err != nil {
		t.Fatal(err)
	}
	if err := l.UnmarshalJSON(doc2); err != nil {
		t.Fatal(err)
	}
	if string(l.R) != "null" || string(l.NC) != "null" {
		t.Errorf("reused RawMessage kept the previous document's value: R=%q NC=%q, want both \"null\"", l.R, l.NC)
	}
	if !reflect.DeepEqual(RawNullDoc(std), l) {
		t.Errorf("reuse diverges:\n stdlib    %#v\n lightning %#v", std, l)
	}
}

// TestStdlibTwinsAreReflectionOnly guards the methodology every stdlib
// differential in this package depends on. A generated UnmarshalJSON makes the
// type a json.Unmarshaler, so json.Unmarshal(doc, &Doc{}) DELEGATES to
// lightning's own decoder and compares it against itself — a tautology that
// looks exactly like a passing parity test. The defence is a defined twin,
// `type fooStd Foo`, which drops Foo's method set.
//
// Dropping the root's methods is not sufficient, though: a twin's *field* types
// keep theirs. If Nested or ArenaKey ever grew an UnmarshalJSON (they would, if
// they stopped being reachable only as nested types and became roots), the
// twins below would quietly start delegating again. So this walks each twin's
// reflect graph and fails on any lightning-generated Unmarshaler found in it.
//
// json.RawMessage and time.Time are the deliberate exceptions: their
// UnmarshalJSON methods belong to the standard library, not to lightning, so
// they are the behaviour the differentials measure lightning *against* rather
// than a way for the comparison to short-circuit into lightning's own decoder.
// (time.Time.UnmarshalJSON is itself part of what nullDocStd pins: on the four
// bytes "null" it returns nil without touching the receiver, which is exactly
// the "no effect" rule under test.)
func TestStdlibTwinsAreReflectionOnly(t *testing.T) {
	twins := []any{
		compactDocStd{}, compactPlainStd{}, rawNullDocStd{},
		arenaDocStd{}, ptrReuseStd{}, byteSliceDocStd{},
		nullDocStd{}, nullLaxDocStd{}, pointListStd{}, scoreMapStd{},
		noCopyListStd{}, byteBlobStd{},
	}
	unmarshaler := reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
	stdlibOwned := map[reflect.Type]bool{
		reflect.TypeOf(json.RawMessage(nil)): true,
		reflect.TypeOf(time.Time{}):          true,
	}

	seen := map[reflect.Type]bool{}
	var walk func(t *testing.T, path string, ty reflect.Type)
	walk = func(t *testing.T, path string, ty reflect.Type) {
		if seen[ty] {
			return
		}
		seen[ty] = true
		if stdlibOwned[ty] {
			return // the standard library's own Unmarshaler, deliberately in play
		}
		if ty.Implements(unmarshaler) || reflect.PointerTo(ty).Implements(unmarshaler) {
			t.Errorf("%s has type %s, which implements json.Unmarshaler: "+
				"encoding/json would delegate to lightning's decoder and the "+
				"differential would compare it against itself", path, ty)
			return
		}
		switch ty.Kind() {
		case reflect.Struct:
			for i := 0; i < ty.NumField(); i++ {
				walk(t, path+"."+ty.Field(i).Name, ty.Field(i).Type)
			}
		case reflect.Slice, reflect.Array, reflect.Pointer:
			walk(t, path+"[elem]", ty.Elem())
		case reflect.Map:
			walk(t, path+"[key]", ty.Key())
			walk(t, path+"[val]", ty.Elem())
		}
	}
	for _, tw := range twins {
		ty := reflect.TypeOf(tw)
		walk(t, ty.Name(), ty)
	}
}
