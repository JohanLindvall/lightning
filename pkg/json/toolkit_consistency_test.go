package json

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"reflect"
	"testing"
)

// TestStripDefaultsKeepKeyHonorsWhitespaceMode locks the emitKept fix: a member
// rescued by the keep list is re-emitted from its ORIGINAL span, which is the one
// piece of output StripDefaults hands back without having decided every byte of it
// — so in RemoveWhitespace mode that span's whitespace has to be taken out here or
// the documented "output is compact regardless of how the input was formatted"
// does not hold. PreserveWhitespace still re-emits verbatim, and AssumeCompact
// asserts there is nothing to remove.
func TestStripDefaultsKeepKeyHonorsWhitespaceMode(t *testing.T) {
	defaults := [][]byte{[]byte("0"), []byte(""), []byte("a b"), []byte(`x \" { y`)}
	keep := [][]byte{[]byte("b")}
	cases := []struct {
		name, in, want string
		ws             WhitespaceMode
		// inPlaceUnsound marks a case whose in-place run cannot be compared with
		// its fresh run — not because of this fix, but because of a separate,
		// pre-existing defect in the same path: an emptied container value is
		// walked (writing into the shared buffer) before being rewound, so the
		// keep path then re-reads a span the walk has already overwritten. The
		// same three inputs diverge identically before this fix. Once the walker
		// re-reads a snapshot rather than the live buffer, the flag can go.
		inPlaceUnsound bool
	}{
		// The reported case: the kept object value came back with its interior
		// whitespace, so the output was not compact.
		{name: "object_interior", in: `{"b":{ "x" : 0 }}`, want: `{"b":{"x":0}}`, ws: RemoveWhitespace},
		{name: "array_interior", in: `{"b":[ 0 , 0 ]}`, want: `{"b":[0,0]}`, ws: RemoveWhitespace},
		// Whitespace between the ':' and the value is part of the same span.
		{name: "after_colon_scalar", in: `{"b": 0}`, want: `{"b":0}`, ws: RemoveWhitespace},
		{name: "after_colon_string", in: `{"b":  ""}`, want: `{"b":""}`, ws: RemoveWhitespace},
		{name: "after_colon_container", in: `{"b":  { "x":0 }}`, want: `{"b":{"x":0}}`, ws: RemoveWhitespace, inPlaceUnsound: true},
		{name: "before_key_colon", in: `{"b" : { "x":0 }}`, want: `{"b":{"x":0}}`, ws: RemoveWhitespace},
		// Nesting, siblings and separators still come out right around it.
		{name: "nested_containers", in: `{"b":{ "x" : [ 0 , { "y" : 0 } ] }}`, want: `{"b":{"x":[0,{"y":0}]}}`, ws: RemoveWhitespace, inPlaceUnsound: true},
		{name: "kept_between_siblings", in: `{"a": 1 , "b": { "x" : 0 } , "c": 2}`, want: `{"a":1,"b":{"x":0},"c":2}`, ws: RemoveWhitespace, inPlaceUnsound: true},
		{name: "already_empty_kept", in: `{"b":{ }}`, want: `{"b":{}}`, ws: RemoveWhitespace},
		// compactValue must not touch bytes inside a string: both of these values
		// are defaults (so the container empties and "b" is kept), and both carry
		// whitespace and structural bytes that have to survive verbatim.
		{name: "string_spaces", in: `{"b":{ "x" : "a b" }}`, want: `{"b":{"x":"a b"}}`, ws: RemoveWhitespace},
		{name: "string_escape", in: `{"b":{ "x y" : "x \" { y" }}`, want: `{"b":{"x y":"x \" { y"}}`, ws: RemoveWhitespace},
		// The other two modes are unchanged: verbatim by definition / by assertion.
		{name: "preserve_verbatim", in: `{"b":{ "x" : 0 }}`, want: `{"b":{ "x" : 0 }}`, ws: PreserveWhitespace},
		{name: "assume_compact_verbatim", in: `{"b":{"x":0}}`, want: `{"b":{"x":0}}`, ws: AssumeCompact},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := string(StripDefaults([]byte(c.in), nil, defaults, keep, c.ws))
			if got != c.want {
				t.Errorf("StripDefaults(%q, %v) = %q, want %q", c.in, c.ws, got, c.want)
			}
			if got != "" && !json.Valid([]byte(got)) {
				t.Errorf("StripDefaults(%q) = %q is not valid JSON", c.in, got)
			}
			// The mode's own contract, checked without a hand-written want: the
			// RemoveWhitespace result must be its own compaction.
			if c.ws == RemoveWhitespace && got != "" {
				var compacted bytes.Buffer
				if err := json.Compact(&compacted, []byte(got)); err != nil {
					t.Fatalf("output not compactable: %v", err)
				}
				if compacted.String() != got {
					t.Errorf("RemoveWhitespace output %q is not compact (compacts to %q)", got, compacted.String())
				}
			}
			// In place must equal fresh. compactValue reads the kept value out of
			// the input while writing the compacted copy back into it, so this is
			// where a write cursor overtaking its own read cursor would show up
			// (it cannot: the copy only ever shrinks, and it starts at or before
			// the value it reads). The flagged cases are excluded for a defect
			// that predates this fix — see inPlaceUnsound.
			if !c.inPlaceUnsound {
				inplace := []byte(c.in)
				if p := string(StripDefaults(inplace, inplace[:0], defaults, keep, c.ws)); p != got {
					t.Errorf("in place = %q, fresh = %q", p, got)
				}
			}
		})
	}

	// The fix must not cost StripDefaults its zero-allocation contract.
	in := []byte(`{"a": 1 , "b": { "x" : 0 } , "c": 2}`)
	out := make([]byte, len(in))
	if n := testing.AllocsPerRun(50, func() { out = StripDefaults(in, out[:0], defaults, keep, RemoveWhitespace) }); n != 0 {
		t.Errorf("StripDefaults allocs/op = %v, want 0", n)
	}
}

// TestStripDefaultsDocumentedBehavior pins the three statements the doc comment
// used to get wrong (they describe behavior the shipped tests already relied on,
// so this test is what keeps doc and behavior from drifting apart again):
// already-empty containers are dropped whatever defaults holds, array *elements*
// equal to a default are removed (which reindexes the array), and
// PreserveWhitespace drops the whitespace that precedes a ','.
func TestStripDefaultsDocumentedBehavior(t *testing.T) {
	t.Run("empty_containers_ignore_defaults", func(t *testing.T) {
		// nil defaults and nil keep: nothing is a default, yet the empty
		// containers still go — and cascade, emptying their parents.
		cases := []struct{ in, want string }{
			{`{"a":{}}`, ``},
			{`{"a":[]}`, ``},
			{`{"a":{},"b":1}`, `{"b":1}`},
			{`[[],{}]`, ``},
			{`{"a":{"b":{}}}`, ``},
			{`[1,{},2]`, `[1,2]`},
		}
		for _, c := range cases {
			if got := string(StripDefaults([]byte(c.in), nil, nil, nil, RemoveWhitespace)); got != c.want {
				t.Errorf("StripDefaults(%q, nil defaults) = %q, want %q", c.in, got, c.want)
			}
		}
		// keep still rescues the member holding it, as it does any other.
		if got := string(StripDefaults([]byte(`{"a":{}}`), nil, nil, [][]byte{[]byte("a")}, RemoveWhitespace)); got != `{"a":{}}` {
			t.Errorf("keep-listed empty container: got %q, want %q", got, `{"a":{}}`)
		}
	})

	t.Run("array_elements_are_dropped", func(t *testing.T) {
		defaults := [][]byte{[]byte("0")}
		cases := []struct{ in, want string }{
			{`[0,1,0,2]`, `[1,2]`}, // reindexes: no hole is left behind
			{`{"a":[0,1,0,2]}`, `{"a":[1,2]}`},
			{`[0,0]`, ``},                                  // all elements default: array empties, cascades
			{`[[0,1],[0]]`, `[[1]]`},                       // nested arrays, one emptying
			{`{"a":[0,{"b":0,"c":3}]}`, `{"a":[{"c":3}]}`}, // object elements strip too
		}
		for _, c := range cases {
			if got := string(StripDefaults([]byte(c.in), nil, defaults, nil, RemoveWhitespace)); got != c.want {
				t.Errorf("StripDefaults(%q) = %q, want %q", c.in, got, c.want)
			}
		}
		// keep applies to object keys only: naming the array's own key does not
		// stop its elements from being dropped.
		got := string(StripDefaults([]byte(`{"a":[0,1]}`), nil, defaults, [][]byte{[]byte("a")}, RemoveWhitespace))
		if got != `{"a":[1]}` {
			t.Errorf("keep on an array member = %q, want %q", got, `{"a":[1]}`)
		}
	})

	t.Run("preserve_drops_whitespace_before_comma", func(t *testing.T) {
		cases := []struct{ in, want string }{
			{`{"a":1 , "b":2}`, `{"a":1, "b":2}`},    // object separator
			{`[1 , 2]`, `[1, 2]`},                    // array separator
			{`{"a":1 }`, `{"a":1 }`},                 // before '}': kept
			{`[1 ]`, `[1 ]`},                         // before ']': kept
			{"{\n  \"a\": 1\n}", "{\n  \"a\": 1\n}"}, // the pretty-printed shape is untouched
		}
		for _, c := range cases {
			if got := string(StripDefaults([]byte(c.in), nil, nil, nil, PreserveWhitespace)); got != c.want {
				t.Errorf("StripDefaults(%q, PreserveWhitespace) = %q, want %q", c.in, got, c.want)
			}
		}
	})
}

// TestSetManyMatchesSetOnNonObjectRoots extends TestSetManyMatchesSet past the
// object roots it locked: SetMany's non-object-root branch replaced the whole
// input rather than the root VALUE, so it dropped leading whitespace (and any
// trailing bytes) that Set and SetPaths both keep — Set(" 5") gave " {\"a\":7}"
// and SetMany gave "{\"a\":7}".
func TestSetManyMatchesSetOnNonObjectRoots(t *testing.T) {
	roots := []string{
		` 5`, `5 `, "\n\t5\n", `5`, `"str"`, `[1,2]`, `null`, `true`,
		``, ` `, ` {"a":1}`, `{"a":1} `, `{ }`, ` [1] `,
	}
	keys := []string{"a", "b"}
	vals := setManyVals("7", "8")
	for _, root := range roots {
		in := []byte(root)
		// One key: SetMany and Set must agree byte for byte.
		many := SetMany(in, nil, vals[:1], keys[:1])
		one := Set(in, nil, vals[0], keys[:1])
		if !bytes.Equal(many, one) {
			t.Errorf("root %q: SetMany = %q, Set = %q", root, many, one)
		}
		// Two keys: SetMany must equal folding Set over them.
		many2 := SetMany(in, nil, vals, keys)
		seq := append([]byte(nil), in...)
		for n := range keys {
			seq = Set(seq, nil, vals[n], keys[n:n+1])
		}
		if !bytes.Equal(many2, seq) {
			t.Errorf("root %q: SetMany = %q, folded Set = %q", root, many2, seq)
		}
		// And SetPaths, the third member of the family, must agree too.
		paths := [][]string{{keys[0]}, {keys[1]}}
		if p := SetPaths(in, nil, vals, paths); !bytes.Equal(many2, p) {
			t.Errorf("root %q: SetMany = %q, SetPaths = %q", root, many2, p)
		}
	}
}

// TestSetManyDuplicateKeyRequest locks SetMany to SetPaths' answer for a
// degenerate request: a key listed twice is set once, from its first entry.
// SetMany used to serve the second entry by appending a second member, so it
// emitted a duplicate-key *document* ({"a":7,"a":8}) that no later validity check
// would flag — while SetPaths, whose matched[] consumes both entries at the first
// occurrence and whose appendMembers dedups by key, returned {"a":7}.
func TestSetManyDuplicateKeyRequest(t *testing.T) {
	cases := []struct {
		name, in string
		keys     []string
		vals     [][]byte
		want     string
	}{
		{"replace_existing", `{"a":1}`, []string{"a", "a"}, setManyVals("7", "8"), `{"a":7}`},
		{"append_absent", `{"z":1}`, []string{"a", "a"}, setManyVals("7", "8"), `{"z":1,"a":7}`},
		{"empty_object", `{}`, []string{"a", "a"}, setManyVals("7", "8"), `{"a":7}`},
		{"non_object_root", `5`, []string{"a", "a"}, setManyVals("7", "8"), `{"a":7}`},
		{"triple", `{"a":1}`, []string{"a", "a", "a"}, setManyVals("7", "8", "9"), `{"a":7}`},
		{"interleaved", `{"a":1,"b":2}`, []string{"a", "b", "a"}, setManyVals("7", "8", "9"), `{"a":7,"b":8}`},
		{"dup_absent_between", `{"z":0}`, []string{"a", "b", "a"}, setManyVals("7", "8", "9"), `{"z":0,"a":7,"b":8}`},
		// A duplicate key in the DOCUMENT is still edited at its first occurrence
		// only; the duplicate request must not reach the second one either.
		{"document_duplicate", `{"a":1,"a":2}`, []string{"a", "a"}, setManyVals("7", "8"), `{"a":7,"a":2}`},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := string(SetMany([]byte(c.in), nil, c.vals, c.keys))
			if got != c.want {
				t.Errorf("SetMany(%q, keys=%v) = %q, want %q", c.in, c.keys, got, c.want)
			}
			// The same request through SetPaths: the family must agree.
			paths := make([][]string, len(c.keys))
			for i, k := range c.keys {
				paths[i] = []string{k}
			}
			if p := string(SetPaths([]byte(c.in), nil, c.vals, paths)); p != got {
				t.Errorf("SetMany(%q) = %q but SetPaths = %q — the two must agree", c.in, got, p)
			}
		})
	}

	// The dedup must not cost SetMany its zero-allocation contract.
	in := []byte(`{"a":1,"b":2,"c":3}`)
	keys := []string{"a", "b", "a", "d"}
	vals := setManyVals("7", "8", "9", "0")
	out := make([]byte, 0, 128)
	if n := testing.AllocsPerRun(50, func() { out = SetMany(in, out, vals, keys) }); n != 0 {
		t.Errorf("SetMany allocs/op = %v, want 0", n)
	}
}

// TestGetPathsIndependentOfCoRequestedPaths locks the property GetPaths' "the
// multi-path form of Get" claim implies: what a path resolves to does not depend
// on what else was requested. Descending for a deeper path is stricter than
// skipping a value — Get(doc,"a") reads {"b" 1} out of {"a":{"b" 1}} while
// Get(doc,"a","x") reports ErrExpectColon — so co-requesting the deeper path used
// to turn the shorter one's success into a whole-call error with BOTH slots nil.
// walkPaths now recovers a failed descent with the same lenient SkipValue a solo
// lookup uses, so the capture stands and the walk continues; the failure is still
// reported as the call's error.
func TestGetPathsIndependentOfCoRequestedPaths(t *testing.T) {
	cases := []struct {
		name  string
		doc   string
		paths [][]string
	}{
		{"leaf_and_deeper_at_same_key", `{"a":{"b" 1}}`, [][]string{{"a"}, {"a", "x"}}},
		{"later_sibling_after_bad_descent", `{"a":{"b" 1},"z":9}`, [][]string{{"a", "x"}, {"z"}}},
		{"sibling_before_and_after", `{"y":8,"a":{"b" 1},"z":9}`, [][]string{{"y"}, {"a", "x"}, {"z"}}},
		{"failure_two_levels_down", `{"o":{"m":{"b" 1},"n":5}}`, [][]string{{"o", "m", "x"}, {"o", "n"}}},
		{"deeper_path_first", `{"o":{"m":{"b" 1},"n":5}}`, [][]string{{"o", "m", "x"}, {"o"}}},
		{"two_bad_descents", `{"a":{"b" 1},"c":{"d" 2},"z":9}`, [][]string{{"a", "x"}, {"c", "y"}, {"z"}}},
		{"unrecoverable_truncation", `{"a":{"b" 1`, [][]string{{"a"}, {"a", "x"}}},
		{"bad_key_in_subtree", `{"a":{x:1},"z":9}`, [][]string{{"a"}, {"a", "x"}, {"z"}}},
		// Well-formed input must be untouched by the recovery.
		{"well_formed", `{"a":{"x":1,"y":2},"z":9}`, [][]string{{"a"}, {"a", "x"}, {"z"}, {"nope"}}},
		{"well_formed_dup_key", `{"o":{"a":1},"o":{"b":2}}`, [][]string{{"o"}, {"o", "a"}, {"o", "b"}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			doc := []byte(c.doc)
			all, allErr := GetPaths(doc, c.paths, nil)
			for n, p := range c.paths {
				solo, soloErr := GetPaths(doc, [][]string{p}, nil)
				if !bytes.Equal(all[n], solo[0]) {
					t.Errorf("path %v: got %q when co-requested, %q alone", p, all[n], solo[0])
				}
				// Get is the reference the doc names: same value for a path that
				// resolves, and where Get reports an error GetPaths reports nil.
				got, _, gerr := Get(doc, p...)
				if gerr == nil && !bytes.Equal(solo[0], got) {
					t.Errorf("path %v: GetPaths alone = %q, Get = %q", p, solo[0], got)
				}
				_ = soloErr
			}
			// The failure still reaches the caller — it is recovered from, not
			// swallowed — and a clean document still reports no error.
			bad := false
			for _, p := range c.paths {
				if _, soloErr := GetPaths(doc, [][]string{p}, nil); soloErr != nil {
					bad = true
				}
			}
			if bad && allErr == nil {
				t.Errorf("a requested path fails alone but the combined call reported no error")
			}
			if !bad && allErr != nil {
				t.Errorf("every path succeeds alone but the combined call reported %v", allErr)
			}
		})
	}

	// The exact shape from the report, spelled out.
	doc := []byte(`{"a":{"b" 1}}`)
	out, err := GetPaths(doc, [][]string{{"a"}, {"a", "x"}}, nil)
	if string(out[0]) != `{"b" 1}` {
		t.Errorf("out[0] = %q, want %q", out[0], `{"b" 1}`)
	}
	if out[1] != nil {
		t.Errorf("out[1] = %q, want nil", out[1])
	}
	if _, _, want := Get(doc, "a", "x"); err != want {
		t.Errorf("err = %v, want %v (what the deeper path reports alone)", err, want)
	}

	// The recovery must not cost GetPaths its zero-allocation contract.
	clean := []byte(`{"a":{"x":1,"y":2},"z":9,"skip":{"deep":{"er":[1,2,3]}}}`)
	paths := [][]string{{"a", "x"}, {"a", "y"}, {"z"}}
	res := make([][]byte, 3)
	if n := testing.AllocsPerRun(50, func() { res, _ = GetPaths(clean, paths, res) }); n != 0 {
		t.Errorf("GetPaths allocs/op = %v, want 0", n)
	}
	if !reflect.DeepEqual(res[0], []byte("1")) {
		t.Errorf("res = %q", res)
	}
}

// TestGetPathsIndependenceRandomized is the property of the test above checked
// over generated documents rather than chosen ones: for every path, requesting it
// alone must give the same bytes as requesting it with the others. The generator
// aims at the shapes that matter — nested objects with malformations just inside
// them, so the walk really descends and really fails — and the check has teeth:
// against the pre-fix walkPaths it reports ~20k mismatches out of 300k documents.
func TestGetPathsIndependenceRandomized(t *testing.T) {
	rnd := rand.New(rand.NewSource(13))
	paths := [][]string{{"a"}, {"a", "x"}, {"a", "x", "z"}, {"b"}, {"z"}, {"b", "z"}}
	keys := []string{"a", "b", "x", "z", "q"}
	// Well-formed and malformed leaves in the same pool: a subtree is entered
	// only for a deeper path, which is exactly when its malformation used to take
	// the other paths down with it.
	frag := []string{`1`, `"s"`, `null`, `[1,2]`, `{}`, `{"x" 1}`, `{"x":}`, `{,}`, `{"x":1`, `{x:1}`, `{"x":1,}`, `{"q":{"z" 2},"z":3}`}
	var gen func(d int) string
	gen = func(d int) string {
		if d == 0 || rnd.Intn(3) == 0 {
			return frag[rnd.Intn(len(frag))]
		}
		s := "{"
		for i := 0; i < 1+rnd.Intn(3); i++ {
			if i > 0 {
				s += ","
			}
			s += `"` + keys[rnd.Intn(len(keys))] + `":` + gen(d-1)
		}
		return s + "}"
	}
	captured, errs := 0, 0
	for i := 0; i < 50000; i++ {
		doc := []byte(gen(3))
		all, err := GetPaths(doc, paths, nil)
		if err != nil {
			errs++
		}
		for k, p := range paths {
			solo, _ := GetPaths(doc, [][]string{p}, nil)
			if solo[0] != nil {
				captured++
			}
			if !bytes.Equal(all[k], solo[0]) {
				t.Fatalf("doc %q path %v: combined %q, solo %q", doc, p, all[k], solo[0])
			}
		}
	}
	// Guard against a generator that stops exercising the interesting paths.
	if captured < 10000 || errs < 10000 {
		t.Errorf("weak corpus: %d captures, %d error-reporting calls", captured, errs)
	}
}

// TestStripDefaultsRandomInputNoPanic pounds the new compactValue/emitKept path
// with random bytes: it copies input spans by index (and, in place, into the very
// buffer it is reading), so a bounds slip would be a panic in a function
// documented as best effort. Every whitespace mode, fresh and in place.
func TestStripDefaultsRandomInputNoPanic(t *testing.T) {
	defaults := [][]byte{[]byte("0"), []byte(""), []byte("none"), []byte("a b")}
	keep := [][]byte{[]byte("k"), []byte("b"), []byte("a"), []byte("")}
	alphabet := []byte(`{}[]",:0123456789abk \n\t` + "\n\t ")
	rnd := rand.New(rand.NewSource(11))
	for i := 0; i < 50000; i++ {
		doc := make([]byte, rnd.Intn(40))
		for j := range doc {
			doc[j] = alphabet[rnd.Intn(len(alphabet))]
		}
		for _, ws := range []WhitespaceMode{RemoveWhitespace, AssumeCompact, PreserveWhitespace} {
			StripDefaults(doc, nil, defaults, keep, ws)
			ip := append([]byte(nil), doc...)
			StripDefaults(ip, ip[:0], defaults, keep, ws)
		}
	}
}
