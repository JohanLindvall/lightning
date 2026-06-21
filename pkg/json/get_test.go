package json

import (
	"errors"
	"testing"
)

func TestGet(t *testing.T) {
	doc := []byte(`{
		"a": {"b": {"c": 42, "s": "hi\n", "arr": [1, 2, 3]}},
		"name": "widget",
		"flag": true,
		"nil": null
	}`)

	tests := []struct {
		name    string
		keys    []string
		want    string
		wantErr error
	}{
		{"root", nil, string(doc), nil},
		{"top scalar", []string{"name"}, `"widget"`, nil},
		{"top bool", []string{"flag"}, `true`, nil},
		{"top null", []string{"nil"}, `null`, nil},
		{"nested object", []string{"a", "b"}, `{"c": 42, "s": "hi\n", "arr": [1, 2, 3]}`, nil},
		{"nested scalar", []string{"a", "b", "c"}, `42`, nil},
		{"nested string keeps quotes+escape", []string{"a", "b", "s"}, `"hi\n"`, nil},
		{"nested array", []string{"a", "b", "arr"}, `[1, 2, 3]`, nil},
		{"missing top key", []string{"nope"}, "", ErrKeyNotFound},
		{"missing nested key", []string{"a", "b", "nope"}, "", ErrKeyNotFound},
		{"descend into non-object", []string{"name", "x"}, "", nil}, // err set, checked below
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, off, err := Get(doc, tt.keys...)
			if tt.name == "descend into non-object" {
				if err == nil {
					t.Fatalf("expected error descending into a string")
				}
				return
			}
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("err = %v, want %v", err, tt.wantErr)
			}
			if tt.wantErr != nil {
				return
			}
			if string(got) != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
			// offset must point at the returned value within doc.
			if off < 0 || off+len(got) > len(doc) || string(doc[off:off+len(got)]) != string(got) {
				t.Fatalf("offset %d does not locate the value in doc", off)
			}
		})
	}
}

func TestGetMany(t *testing.T) {
	doc := []byte(`{
		"a": 1,
		"name": "widget",
		"obj": {"x": [1, 2]},
		"flag": true,
		"nil": null,
		"s": "hi\n"
	}`)

	t.Run("order, miss, null vs missing", func(t *testing.T) {
		keys := []string{"name", "nope", "a", "nil", "obj", "s"}
		out, err := GetMany(doc, keys, nil)
		if err != nil {
			t.Fatalf("GetMany: %v", err)
		}
		if len(out) != len(keys) {
			t.Fatalf("len(out) = %d, want %d", len(out), len(keys))
		}
		want := []string{`"widget"`, "", `1`, `null`, `{"x": [1, 2]}`, `"hi\n"`}
		for n := range keys {
			if keys[n] == "nope" {
				if out[n] != nil {
					t.Fatalf("missing key %q: got %q, want nil", keys[n], out[n])
				}
				continue
			}
			if string(out[n]) != want[n] {
				t.Fatalf("key %q: got %q, want %q", keys[n], out[n], want[n])
			}
		}
	})

	t.Run("scratch reuse does not allocate", func(t *testing.T) {
		keys := []string{"a", "name", "flag"}
		scratch := make([][]byte, 0, len(keys))
		out, err := GetMany(doc, keys, scratch)
		if err != nil {
			t.Fatalf("GetMany: %v", err)
		}
		allocs := testing.AllocsPerRun(100, func() {
			out, _ = GetMany(doc, keys, out)
		})
		if allocs != 0 {
			t.Fatalf("GetMany reusing scratch allocated %.0f times, want 0", allocs)
		}
	})

	t.Run("duplicate requested key filled in both slots", func(t *testing.T) {
		out, err := GetMany(doc, []string{"a", "a"}, nil)
		if err != nil {
			t.Fatalf("GetMany: %v", err)
		}
		if string(out[0]) != "1" || string(out[1]) != "1" {
			t.Fatalf("duplicate key: got %q and %q, want both 1", out[0], out[1])
		}
	})

	t.Run("no keys", func(t *testing.T) {
		out, err := GetMany(doc, nil, nil)
		if err != nil || len(out) != 0 {
			t.Fatalf("GetMany(nil keys) = %v, %v; want empty, nil", out, err)
		}
	})

	t.Run("non-object root errors", func(t *testing.T) {
		if _, err := GetMany([]byte(`[1, 2]`), []string{"a"}, nil); err == nil {
			t.Fatal("expected an error for a non-object root")
		}
	})

	t.Run("malformed json errors", func(t *testing.T) {
		if _, err := GetMany([]byte(`{"a": 1, "b":`), []string{"x"}, nil); err == nil {
			t.Fatal("expected an error for truncated input")
		}
	})
}

func TestGetPaths(t *testing.T) {
	doc := []byte(`{
		"actor": {"login": "octocat", "id": 583231},
		"repo":  {"name": "Hello-World", "owner": {"login": "octo-org"}},
		"payload": {"ref": "refs/heads/main", "commits": [{"sha": "abc"}]},
		"public": true
	}`)

	t.Run("nested paths, shared prefix, miss, root capture", func(t *testing.T) {
		paths := [][]string{
			{"actor", "login"},         // octocat
			{"actor", "id"},            // 583231 (shares the actor descent)
			{"repo", "owner", "login"}, // octo-org (two levels deep)
			{"payload", "ref"},         // refs/heads/main
			{"repo"},                   // whole repo object (ends where another path descends)
			{"missing", "x"},           // absent
			{"actor", "name"},          // absent (key not present)
			nil,                        // whole document root
		}
		out, err := GetPaths(doc, paths, nil)
		if err != nil {
			t.Fatalf("GetPaths: %v", err)
		}
		want := []string{
			`"octocat"`,
			`583231`,
			`"octo-org"`,
			`"refs/heads/main"`,
			`{"name": "Hello-World", "owner": {"login": "octo-org"}}`,
			"", // missing.x
			"", // actor.name
			string(doc),
		}
		for n := range paths {
			if string(out[n]) != want[n] {
				t.Errorf("path %v: got %q, want %q", paths[n], out[n], want[n])
			}
		}
	})

	t.Run("agrees with Get for every path", func(t *testing.T) {
		paths := [][]string{
			{"actor", "login"}, {"repo", "owner", "login"}, {"payload", "ref"}, {"public"},
		}
		out, err := GetPaths(doc, paths, nil)
		if err != nil {
			t.Fatalf("GetPaths: %v", err)
		}
		for n, p := range paths {
			g, _, gerr := Get(doc, p...)
			if gerr != nil {
				t.Fatalf("Get%v: %v", p, gerr)
			}
			if string(out[n]) != string(g) {
				t.Errorf("path %v: GetPaths %q != Get %q", p, out[n], g)
			}
		}
	})

	t.Run("scratch reuse does not allocate beyond the active set", func(t *testing.T) {
		paths := [][]string{{"public"}, {"actor", "login"}}
		out, _ := GetPaths(doc, paths, make([][]byte, 0, len(paths)))
		// One descent (into actor) builds one small recurse slice; assert it stays small.
		allocs := testing.AllocsPerRun(100, func() { out, _ = GetPaths(doc, paths, out) })
		if allocs > 1 {
			t.Fatalf("GetPaths reusing scratch allocated %.0f times, want <= 1", allocs)
		}
	})

	t.Run("compact variant agrees", func(t *testing.T) {
		compact := []byte(`{"a":{"b":1,"c":2},"d":3}`)
		paths := [][]string{{"a", "b"}, {"a", "c"}, {"d"}}
		out, err := GetPathsCompact(compact, paths, nil)
		if err != nil {
			t.Fatalf("GetPathsCompact: %v", err)
		}
		want := []string{"1", "2", "3"}
		for n := range paths {
			if string(out[n]) != want[n] {
				t.Errorf("path %v: got %q, want %q", paths[n], out[n], want[n])
			}
		}
	})

	t.Run("non-object root errors when a path needs to descend", func(t *testing.T) {
		if _, err := GetPaths([]byte(`[1,2]`), [][]string{{"a"}}, nil); err == nil {
			t.Fatal("expected an error for a non-object root")
		}
	})
}

// TestCompactVariants checks that GetCompact, GetManyCompact and
// ObjectEachCompact agree with their non-compact counterparts on whitespace-free
// input, descend nested paths, and still tolerate leading framing whitespace.
func TestCompactVariants(t *testing.T) {
	// No whitespace between tokens (the form a compact serializer emits).
	compact := []byte(`{"a":{"b":{"c":42,"s":"hi"}},"name":"widget","flag":true,"nums":[1,2,3]}`)

	t.Run("GetCompact matches Get", func(t *testing.T) {
		paths := [][]string{nil, {"name"}, {"flag"}, {"nums"}, {"a", "b"}, {"a", "b", "c"}, {"a", "b", "s"}}
		for _, keys := range paths {
			want, woff, werr := Get(compact, keys...)
			got, goff, gerr := GetCompact(compact, keys...)
			if werr != nil {
				t.Fatalf("%v: Get returned error %v", keys, werr)
			}
			if gerr != nil || goff != woff || string(got) != string(want) {
				t.Fatalf("%v: GetCompact=(%q,%d,%v), want (%q,%d,nil)", keys, got, goff, gerr, want, woff)
			}
		}
		// Missing key path still reports ErrKeyNotFound.
		if _, _, err := GetCompact(compact, "nope"); !errors.Is(err, ErrKeyNotFound) {
			t.Fatalf("GetCompact missing key: err=%v, want ErrKeyNotFound", err)
		}
	})

	t.Run("GetManyCompact matches GetMany", func(t *testing.T) {
		keys := []string{"name", "nope", "flag", "nums"}
		want, werr := GetMany(compact, keys, nil)
		got, gerr := GetManyCompact(compact, keys, nil)
		if werr != nil || gerr != nil {
			t.Fatalf("GetMany err=%v, GetManyCompact err=%v", werr, gerr)
		}
		for n := range keys {
			if string(got[n]) != string(want[n]) {
				t.Fatalf("key %q: GetManyCompact %q, GetMany %q", keys[n], got[n], want[n])
			}
		}
	})

	t.Run("ObjectEachCompact matches ObjectEach", func(t *testing.T) {
		collect := func(each func([]byte, func(string, []byte) error, ...string) error, keys ...string) []string {
			var out []string
			if err := each(compact, func(k string, v []byte) error {
				out = append(out, k+"="+string(v))
				return nil
			}, keys...); err != nil {
				t.Fatalf("%v: ObjectEach* error: %v", keys, err)
			}
			return out
		}
		for _, keys := range [][]string{nil, {"a", "b"}} {
			want := collect(ObjectEach, keys...)
			got := collect(ObjectEachCompact, keys...)
			if len(got) != len(want) {
				t.Fatalf("%v: got %d members, want %d", keys, len(got), len(want))
			}
			for n := range got {
				if got[n] != want[n] {
					t.Fatalf("%v member %d: got %q, want %q", keys, n, got[n], want[n])
				}
			}
		}
	})

	t.Run("leading whitespace tolerated", func(t *testing.T) {
		framed := append([]byte("\n\t  "), compact...)
		if v, _, err := GetCompact(framed, "name"); err != nil || string(v) != `"widget"` {
			t.Fatalf("GetCompact framed: %q, %v", v, err)
		}
		if out, err := GetManyCompact(framed, []string{"flag"}, nil); err != nil || string(out[0]) != "true" {
			t.Fatalf("GetManyCompact framed: %q, %v", out, err)
		}
		n := 0
		if err := ObjectEachCompact(framed, func(string, []byte) error { n++; return nil }); err != nil || n == 0 {
			t.Fatalf("ObjectEachCompact framed: n=%d, %v", n, err)
		}
	})
}

// TestGetObjectCheck demonstrates the "value exists and is an object" idiom on
// Get: navigate to the parent path and confirm err == nil and it starts with '{'.
func TestGetObjectCheck(t *testing.T) {
	doc := []byte(`{"parent": {"child": 1}, "leaf": 5}`)

	if v, _, err := Get(doc, "parent"); err != nil || len(v) == 0 || v[0] != '{' {
		t.Fatalf("expected parent to be an object, got %q err=%v", v, err)
	}
	if v, _, err := Get(doc, "leaf"); err == nil && len(v) > 0 && v[0] == '{' {
		t.Fatalf("leaf should not look like an object: %q", v)
	}
}
