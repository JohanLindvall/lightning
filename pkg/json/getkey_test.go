package json

import "testing"

// TestEscapedKeysAllReaders covers the fallback half of the inlined key read in
// get.go. The fast path handles a key with no backslash by aliasing it out of the
// input; anything escaped drops to unstable.ReadKey, which decodes it. A bug in
// that split would be silent — the key simply would not match — so every reader is
// checked, with the escaped keys placed after a plain one so the fast path runs
// first and the fallback is genuinely reached mid-object.
func TestEscapedKeysAllReaders(t *testing.T) {
	// `esc"key` needs a real escape; "tab\there" decodes to a control byte; and
	// `sl/ash` uses \/, an escape whose decoded byte needs none.
	doc := []byte(`{"plain":1,"esc\"key":2,"tab\there":3,"sl\/ash":4,` +
		`"nest":{"in\"ner":5}}`)

	t.Run("Get and Lookup", func(t *testing.T) {
		for _, tc := range []struct{ key, want string }{
			{"plain", "1"}, {`esc"key`, "2"}, {"tab\there", "3"}, {"sl/ash", "4"},
		} {
			got, _, err := Get(doc, tc.key)
			if err != nil || string(got) != tc.want {
				t.Errorf("Get(%q) = %s, %v; want %s", tc.key, got, err, tc.want)
			}
			if got, err := Lookup(doc, tc.key); err != nil || string(got) != tc.want {
				t.Errorf("Lookup(%q) = %s, %v; want %s", tc.key, got, err, tc.want)
			}
		}
		// An escaped key one level down, reached through objectField's descent.
		if got, err := Lookup(doc, "nest", `in"ner`); err != nil || string(got) != "5" {
			t.Errorf(`Lookup(nest, in"ner) = %s, %v; want 5`, got, err)
		}
	})

	t.Run("GetMany", func(t *testing.T) {
		keys := []string{"plain", `esc"key`, "tab\there", "sl/ash", "absent"}
		vals, err := GetMany(doc, keys, nil)
		if err != nil {
			t.Fatal(err)
		}
		for n, w := range []string{"1", "2", "3", "4", ""} {
			if string(vals[n]) != w {
				t.Errorf("GetMany[%d] (%q) = %q, want %q", n, keys[n], vals[n], w)
			}
		}
	})

	t.Run("GetPaths", func(t *testing.T) {
		paths := [][]string{{`esc"key`}, {"nest", `in"ner`}, {"sl/ash"}}
		vals, err := GetPaths(doc, paths, nil)
		if err != nil {
			t.Fatal(err)
		}
		for n, w := range []string{"2", "5", "4"} {
			if string(vals[n]) != w {
				t.Errorf("GetPaths[%d] (%v) = %q, want %q", n, paths[n], vals[n], w)
			}
		}
	})

	t.Run("ObjectEach", func(t *testing.T) {
		got := map[string]string{}
		if err := ObjectEach(doc, func(k string, v []byte) error {
			got[k] = string(v)
			return nil
		}); err != nil {
			t.Fatal(err)
		}
		for k, w := range map[string]string{
			"plain": "1", `esc"key`: "2", "tab\there": "3", "sl/ash": "4",
		} {
			if got[k] != w {
				t.Errorf("ObjectEach[%q] = %q, want %q", k, got[k], w)
			}
		}
	})
}

// TestMalformedKeysAllReaders checks the inlined fast path does not swallow a bad
// key: the close-quote test must fail for a truncated or badly escaped key and hand
// off to unstable.ReadKey, which reports it. Aliasing past the end instead would be
// a memory-safety bug, so this is the important half.
func TestMalformedKeysAllReaders(t *testing.T) {
	// The looked-up key is absent from every document, so each reader has to walk
	// all the way onto the malformed key rather than stopping early. (Get and
	// GetMany deliberately return as soon as the requested keys are found and do not
	// validate the rest of the object, so asking for a key that appears *before* the
	// damage would succeed by design.)
	const absent = "zzz"
	for _, bad := range []string{
		`{"unterminated`,      // no close quote at all
		`{"esc\`,              // ends inside an escape
		`{"a\q":1}`,           // invalid escape in the key
		`{"a":1,"unterminat`,  // truncated on a later member: fast path ran first
		`{"a":1,"b\u00":2}`,   // short \u escape on a later member
		`{"ok":1,"bad\uZZZZ"`, // bad hex on a later member
	} {
		if _, _, err := Get([]byte(bad), absent); err == nil {
			t.Errorf("Get(%q) succeeded, want an error", bad)
		}
		if _, err := GetMany([]byte(bad), []string{absent}, nil); err == nil {
			t.Errorf("GetMany(%q) succeeded, want an error", bad)
		}
		if _, err := GetPaths([]byte(bad), [][]string{{absent}}, nil); err == nil {
			t.Errorf("GetPaths(%q) succeeded, want an error", bad)
		}
		if err := ObjectEach([]byte(bad), func(string, []byte) error { return nil }); err == nil {
			t.Errorf("ObjectEach(%q) succeeded, want an error", bad)
		}
	}
}
