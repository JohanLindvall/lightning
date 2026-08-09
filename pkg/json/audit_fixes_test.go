package json

import (
	"bytes"
	stdjson "encoding/json"
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"testing"
)

// TestCheckedRejectsUnsafeKeys locks the key gate of the checked edit wrappers
// (audit S3). Set/SetMany/SetPaths write a created key raw between two quotes, so
// a key that closes its own string splices arbitrary members into the document —
// and because the forgery is well-formed JSON, the wrappers' result-side Valid
// pass waves it through. Before the gate, the payload below returned
//
//	{"user":"alice","role":"guest","x":1,"role":"admin"}
//
// with a nil error, which encoding/json reads back as role=admin (duplicate key,
// last wins).
func TestCheckedRejectsUnsafeKeys(t *testing.T) {
	in := []byte(`{"user":"alice","role":"guest"}`)
	const inject = `x":1,"role` // closes the key string and opens a balanced member

	unsafeKeys := []struct {
		name string
		key  string
	}{
		{"injection", inject},
		{"quote", `a"b`},
		{"backslash", `a\b`},
		{"control_byte", "a\x01b"},
		{"newline", "a\nb"},
	}

	for _, uk := range unsafeKeys {
		t.Run("Set/"+uk.name, func(t *testing.T) {
			res, err := SetChecked(in, nil, []byte(`"admin"`), []string{uk.key})
			if !errors.Is(err, ErrUnsafeKey) {
				t.Fatalf("SetChecked(key=%q) = %q, %v; want ErrUnsafeKey", uk.key, res, err)
			}
			if res != nil {
				t.Errorf("SetChecked(key=%q) returned %q with err %v, want nil bytes", uk.key, res, err)
			}
		})
		t.Run("SetMany/"+uk.name, func(t *testing.T) {
			res, err := SetManyChecked(in, nil, [][]byte{[]byte(`"admin"`)}, []string{uk.key})
			if !errors.Is(err, ErrUnsafeKey) {
				t.Fatalf("SetManyChecked(key=%q) = %q, %v; want ErrUnsafeKey", uk.key, res, err)
			}
		})
		t.Run("SetPaths/leaf/"+uk.name, func(t *testing.T) {
			res, err := SetPathsChecked(in, nil, [][]byte{[]byte(`"admin"`)}, [][]string{{uk.key}})
			if !errors.Is(err, ErrUnsafeKey) {
				t.Fatalf("SetPathsChecked(path=[%q]) = %q, %v; want ErrUnsafeKey", uk.key, res, err)
			}
		})
		t.Run("SetPaths/interior/"+uk.name, func(t *testing.T) {
			// An interior path element is written raw too (appendMembers builds the
			// nested objects), so it must be checked at every depth, not just the leaf.
			res, err := SetPathsChecked(in, nil, [][]byte{[]byte(`"admin"`)}, [][]string{{uk.key, "leaf"}})
			if !errors.Is(err, ErrUnsafeKey) {
				t.Fatalf("SetPathsChecked(path=[%q leaf]) = %q, %v; want ErrUnsafeKey", uk.key, res, err)
			}
		})
	}

	// The escape gate must not reject keys that are merely non-ASCII: JSON does
	// not require escaping them and Set writes them back unchanged.
	ok := []struct {
		name string
		keys []string
	}{
		{"plain", []string{"role"}},
		{"new_key", []string{"fresh"}},
		{"utf8", []string{"héllo"}},
		{"del_byte", []string{"a\x7fb"}}, // 0x7f is not a JSON escape
		{"empty", []string{""}},
	}
	for _, c := range ok {
		res, err := SetChecked(in, nil, []byte(`"admin"`), c.keys)
		if err != nil {
			t.Errorf("SetChecked(keys=%q) = %v, want success", c.keys, err)
			continue
		}
		var back map[string]any
		if err := stdjson.Unmarshal(res, &back); err != nil {
			t.Errorf("SetChecked(keys=%q) produced %q: %v", c.keys, res, err)
			continue
		}
		if back[c.keys[0]] != "admin" {
			t.Errorf("SetChecked(keys=%q) = %q; key does not read back", c.keys, res)
		}
	}

	// The gate belongs to the checked wrappers only: the unchecked Set keeps its
	// documented raw-key behaviour (this is what the wrappers exist to guard).
	if got := string(Set(in, nil, []byte(`"admin"`), []string{inject})); !strings.Contains(got, `"x":1`) {
		t.Errorf("unchecked Set changed behaviour: %q", got)
	}
}

// TestCheckedWrappersAreDocumented is the cheap stand-in for `go doc` (audit A3):
// SetChecked's doc block was separated from its declaration by allValid, so the
// prose attached to the unexported helper and `go doc ./pkg/json SetChecked`
// printed a bare signature. Parsing the file is hermetic and catches the same
// mistake — a doc comment with anything at all declared between it and the
// function it describes.
func TestCheckedWrappersAreDocumented(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "checked.go", nil, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse checked.go: %v", err)
	}
	want := map[string]bool{"SetChecked": false, "SetManyChecked": false, "SetPathsChecked": false}
	for _, d := range f.Decls {
		fn, ok := d.(*ast.FuncDecl)
		if !ok {
			continue
		}
		if _, ok := want[fn.Name.Name]; !ok {
			continue
		}
		want[fn.Name.Name] = true
		if fn.Doc == nil {
			t.Errorf("%s has no doc comment attached to its declaration", fn.Name.Name)
			continue
		}
		doc := fn.Doc.Text()
		if !strings.Contains(doc, "ErrUnsafeKey") {
			t.Errorf("%s doc does not document the key precondition:\n%s", fn.Name.Name, doc)
		}
	}
	for name, seen := range want {
		if !seen {
			t.Errorf("%s not found in checked.go", name)
		}
	}
}

// TestSetPathsFirstOccurrenceWins locks SetPaths to Set's duplicate-key rule
// (audit C6): a key that occurs twice in one object is edited at its first
// occurrence only. setObject marked a path index matched but kept applying it, so
// every duplicate was rewritten — and at the root frame the all-matched early exit
// hid it, which made the outcome for one path depend on whether an *unrelated*
// path happened to be found.
func TestSetPathsFirstOccurrenceWins(t *testing.T) {
	cases := []struct {
		name    string
		in      string
		paths   [][]string
		rawVal  [][]byte
		want    string
		sameKey []string // key path to cross-check against Set, nil to skip
	}{
		{
			name:    "nested_frame",
			in:      `{"o":{"k":1,"k":2}}`,
			paths:   [][]string{{"o", "k"}},
			rawVal:  [][]byte{[]byte("9")},
			want:    `{"o":{"k":9,"k":2}}`,
			sameKey: []string{"o", "k"},
		},
		{
			name:   "nested_frame_with_absent_path",
			in:     `{"o":{"k":1,"k":2}}`,
			paths:  [][]string{{"o", "k"}, {"zz"}},
			rawVal: [][]byte{[]byte("9"), []byte("0")},
			want:   `{"o":{"k":9,"k":2},"zz":0}`,
		},
		{
			name:    "root_frame",
			in:      `{"k":1,"k":2}`,
			paths:   [][]string{{"k"}},
			rawVal:  [][]byte{[]byte("9")},
			want:    `{"k":9,"k":2}`,
			sameKey: []string{"k"},
		},
		{
			// The absent path suppresses the all-matched early exit, which is what
			// used to make the root frame edit both duplicates.
			name:   "root_frame_with_absent_path",
			in:     `{"k":1,"k":2}`,
			paths:  [][]string{{"k"}, {"z"}},
			rawVal: [][]byte{[]byte("9"), []byte("0")},
			want:   `{"k":9,"k":2,"z":0}`,
		},
		{
			// Duplicate parent: only the first "o" is descended, matching Set.
			name:    "duplicate_parent",
			in:      `{"o":{"k":1},"o":{"k":2}}`,
			paths:   [][]string{{"o", "k"}},
			rawVal:  [][]byte{[]byte("9")},
			want:    `{"o":{"k":9},"o":{"k":2}}`,
			sameKey: []string{"o", "k"},
		},
		{
			// The first occurrence is not an object: Set replaces it with the
			// remaining path built as nested objects and never looks at the second
			// occurrence, so SetPaths must not either.
			name:    "duplicate_parent_nonobject_first",
			in:      `{"o":5,"o":{"k":1}}`,
			paths:   [][]string{{"o", "k"}},
			rawVal:  [][]byte{[]byte("9")},
			want:    `{"o":{"k":9},"o":{"k":1}}`,
			sameKey: []string{"o", "k"},
		},
		{
			// Three occurrences, deeper nesting: still only the first.
			name:    "triple",
			in:      `{"a":{"b":{"c":1,"c":2,"c":3}}}`,
			paths:   [][]string{{"a", "b", "c"}},
			rawVal:  [][]byte{[]byte("9")},
			want:    `{"a":{"b":{"c":9,"c":2,"c":3}}}`,
			sameKey: []string{"a", "b", "c"},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := string(SetPaths([]byte(c.in), nil, c.rawVal, c.paths))
			if got != c.want {
				t.Errorf("SetPaths(%q) = %q, want %q", c.in, got, c.want)
			}
			if c.sameKey != nil {
				set := string(Set([]byte(c.in), nil, c.rawVal[0], c.sameKey))
				if got != set {
					t.Errorf("SetPaths(%q) = %q but Set(%q) = %q — the multi-path form must agree with Set", c.in, got, c.in, set)
				}
			}
		})
	}

	// The fix must not cost SetPaths its zero-allocation contract (a reused out).
	out := make([]byte, 0, 256)
	in := []byte(`{"o":{"k":1,"k":2},"other":[1,2,3]}`)
	paths := [][]string{{"o", "k"}, {"zz"}}
	rawVal := [][]byte{[]byte("9"), []byte("0")}
	if n := testing.AllocsPerRun(50, func() { out = SetPaths(in, out, rawVal, paths) }); n != 0 {
		t.Errorf("SetPaths allocs/op = %v, want 0", n)
	}
}

// TestGetPathsFirstOccurrenceWins locks GetPaths to Get's duplicate-key rule
// (audit C7). GetPaths is documented as the multi-path form of Get, but it
// descended *every* occurrence of a matched parent key, so Get(o,b) reported
// ErrKeyNotFound on {"o":{"a":1},"o":{"b":2}} while GetPaths(o,b) returned 2.
func TestGetPathsFirstOccurrenceWins(t *testing.T) {
	const dup = `{"o":{"a":1},"o":{"b":2}}`
	for _, path := range [][]string{{"o", "a"}, {"o", "b"}, {"o"}} {
		want, _, gerr := Get([]byte(dup), path...)
		got, err := GetPaths([]byte(dup), [][]string{path}, nil)
		if err != nil {
			t.Fatalf("GetPaths(%q) error: %v", path, err)
		}
		if gerr != nil {
			// Get reports a missing path as an error; GetPaths reports it as nil.
			if got[0] != nil {
				t.Errorf("Get(%q) = %v but GetPaths = %q; the two must agree", path, gerr, got[0])
			}
			continue
		}
		if !bytes.Equal(got[0], want) {
			t.Errorf("GetPaths(%q) = %q, Get = %q; the two must agree", path, got[0], want)
		}
	}

	// Both paths at once: only the first "o" is descended, so "b" stays absent.
	got, err := GetPaths([]byte(dup), [][]string{{"o", "a"}, {"o", "b"}}, nil)
	if err != nil {
		t.Fatalf("GetPaths: %v", err)
	}
	if string(got[0]) != "1" || got[1] != nil {
		t.Errorf("GetPaths([o a],[o b]) = [%q %q], want [\"1\" nil]", got[0], got[1])
	}

	// The first occurrence of the parent is not an object: Get stops there with
	// ErrExpectObject rather than trying the second, so GetPaths must report the
	// path absent instead of walking on to the later duplicate.
	if _, _, err := Get([]byte(`{"o":5,"o":{"a":1}}`), "o", "a"); err == nil {
		t.Fatal("premise: Get should refuse to descend a non-object")
	}
	got, err = GetPaths([]byte(`{"o":5,"o":{"a":1}}`), [][]string{{"o", "a"}}, nil)
	if err != nil {
		t.Fatalf("GetPaths: %v", err)
	}
	if got[0] != nil {
		t.Errorf("GetPaths([o a]) on a non-object first duplicate = %q, want nil (Get reports ErrExpectObject)", got[0])
	}

	// A duplicate leaf key takes its first occurrence, as GetMany already did.
	got, err = GetPaths([]byte(`{"o":{"k":1,"k":2}}`), [][]string{{"o", "k"}}, nil)
	if err != nil {
		t.Fatalf("GetPaths: %v", err)
	}
	if string(got[0]) != "1" {
		t.Errorf("GetPaths on duplicate leaf = %q, want \"1\"", got[0])
	}

	// Duplicate requested paths are still all filled, and a non-duplicate document
	// is unaffected by the consumed-path bookkeeping.
	plain := []byte(`{"a":{"b":1,"c":2},"d":3}`)
	paths := [][]string{{"a", "b"}, {"a", "c"}, {"d"}, {"a", "b"}, {"nope"}}
	got, err = GetPaths(plain, paths, nil)
	if err != nil {
		t.Fatalf("GetPaths: %v", err)
	}
	for n, want := range []string{"1", "2", "3", "1", ""} {
		if string(got[n]) != want {
			t.Errorf("GetPaths[%d] = %q, want %q", n, got[n], want)
		}
	}

	// The fix must not cost GetPaths its zero-allocation contract.
	out := make([][]byte, len(paths))
	if n := testing.AllocsPerRun(50, func() { out, _ = GetPaths(plain, paths, out) }); n != 0 {
		t.Errorf("GetPaths allocs/op = %v, want 0", n)
	}
}

// TestGetPathsFillsEveryEmptyPath locks audit C8: GetPaths documents that
// duplicate paths are all filled and that an empty path selects the root, but the
// root capture assigned only the first empty path and left the rest nil.
func TestGetPathsFillsEveryEmptyPath(t *testing.T) {
	data := []byte(`{"a":1}`)
	got, err := GetPaths(data, [][]string{{}, {}}, nil)
	if err != nil {
		t.Fatalf("GetPaths: %v", err)
	}
	for n := range got {
		if string(got[n]) != `{"a":1}` {
			t.Errorf("GetPaths(empty,empty)[%d] = %q, want the root value", n, got[n])
		}
	}

	// Mixed with a real path, and with nil as well as zero-length spellings.
	got, err = GetPaths(data, [][]string{nil, {"a"}, {}}, nil)
	if err != nil {
		t.Fatalf("GetPaths: %v", err)
	}
	if string(got[0]) != `{"a":1}` || string(got[1]) != "1" || string(got[2]) != `{"a":1}` {
		t.Errorf("GetPaths(nil,[a],empty) = [%q %q %q]", got[0], got[1], got[2])
	}
}

// TestStripDefaultsWhitespaceOnlyContainer locks audit C9: handle's top-level
// empty-object test lacked the whitespace peek its sibling (the object-member
// container branch) has, so "{ }" in array-element position fell into eject(),
// whose best-effort contract copies the whole remainder of the input through
// verbatim — silently leaving every later member unstripped, and in
// RemoveWhitespace mode unminified, behind valid-looking JSON.
func TestStripDefaultsWhitespaceOnlyContainer(t *testing.T) {
	defaults := [][]byte{[]byte("1")}
	cases := []struct {
		name, in, want string
	}{
		{"array_element_object", `[{ },{"a":1,"b":2}]`, `[{"b":2}]`},
		{"array_element_array", `[[ ],{"a":1,"b":2}]`, `[{"b":2}]`},
		{"nested_array", `[[{ }],{"a":1,"b":2}]`, `[{"b":2}]`},
		{"trailing_element", `[{"a":1,"b":2},{ }]`, `[{"b":2}]`},
		{"whitespace_run", "[{\n  \t},{\"a\":1,\"b\":2}]", `[{"b":2}]`},
		{"object_field", `{"x":{ },"a":1,"b":2}`, `{"b":2}`},
		{"deep", `{"o":[{ },{"a":1,"c":3}]}`, `{"o":[{"c":3}]}`},
		// The root case, for consistency with the compact "{}"/"[]" spelling,
		// which the empty-container rule already drops entirely.
		{"root_object", `{ }`, ``},
		{"root_array", `[ ]`, ``},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := string(StripDefaults([]byte(c.in), nil, defaults, nil, RemoveWhitespace))
			if got != c.want {
				t.Errorf("StripDefaults(%q, RemoveWhitespace) = %q, want %q", c.in, got, c.want)
			}
			if got != "" && !stdjson.Valid([]byte(got)) {
				t.Errorf("StripDefaults(%q) = %q is not valid JSON", c.in, got)
			}
		})
	}

	// PreserveWhitespace must reach the same data — the invariant the existing
	// preserve test uses — which before the fix it did not: the ejected remainder
	// kept "a":1 in both modes.
	for _, in := range []string{`[{ }, {"a": 1, "b": 2}]`, `{"x": { }, "a": 1, "b": 2}`} {
		pw := StripDefaults([]byte(in), nil, defaults, nil, PreserveWhitespace)
		rw := StripDefaults([]byte(in), nil, defaults, nil, RemoveWhitespace)
		var compacted bytes.Buffer
		if len(pw) > 0 {
			if err := stdjson.Compact(&compacted, pw); err != nil {
				t.Fatalf("preserve output %q is not valid JSON: %v", pw, err)
			}
		}
		if compacted.String() != string(rw) {
			t.Errorf("StripDefaults(%q): preserve(compacted) = %q, remove = %q", in, compacted.String(), rw)
		}
		if strings.Contains(compacted.String(), `"a":1`) {
			t.Errorf("StripDefaults(%q, PreserveWhitespace) = %q kept the default member", in, pw)
		}
	}
}
