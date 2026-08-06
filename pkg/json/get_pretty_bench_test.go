package json

import (
	"bytes"
	stdjson "encoding/json"
	"fmt"
	"testing"
)

// prettyDoc builds an indented (pretty-printed) record of nrec nested records,
// the shape whose inter-token whitespace runs (newline + indentation) the
// walkers pay for on every member. The interesting keys sit at the end so a
// lookup walks the whole document.
func prettyDoc(b testing.TB, nrec int) []byte {
	var sb bytes.Buffer
	sb.WriteString(`{`)
	for i := 0; i < nrec; i++ {
		fmt.Fprintf(&sb, `"rec%03d":{"name":"record number %03d","value":%d,"active":true,"tags":["a","b"]},`, i, i, i)
	}
	sb.WriteString(`"target":{"deep":{"leaf":42}},"last":7}`)
	var out bytes.Buffer
	if err := stdjson.Indent(&out, sb.Bytes(), "", "    "); err != nil {
		b.Fatal(err)
	}
	return out.Bytes()
}

// BenchmarkGetManyPretty and friends measure the non-compact walkers on
// pretty-printed input — the workload the *Compact variants exclude and no
// other committed benchmark covers. The compact walkers never scan whitespace,
// so any change to inter-token skipping shows up here first.
func BenchmarkGetManyPretty(b *testing.B) {
	doc := prettyDoc(b, 60)
	keys := []string{"target", "last"}
	var out [][]byte
	b.SetBytes(int64(len(doc)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var err error
		out, err = GetMany(doc, keys, out)
		if err != nil || out[1] == nil {
			b.Fatal(err, out)
		}
	}
}

func BenchmarkGetPretty(b *testing.B) {
	doc := prettyDoc(b, 60)
	b.SetBytes(int64(len(doc)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v, _, err := Get(doc, "target", "deep", "leaf")
		if err != nil || len(v) == 0 {
			b.Fatal(err)
		}
	}
}

func BenchmarkObjectEachPretty(b *testing.B) {
	doc := prettyDoc(b, 60)
	b.SetBytes(int64(len(doc)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := 0
		if err := ObjectEach(doc, func(key string, value []byte) error {
			n++
			return nil
		}); err != nil || n == 0 {
			b.Fatal(err)
		}
	}
}

func BenchmarkStripDefaultsPretty(b *testing.B) {
	doc := prettyDoc(b, 60)
	defaults := [][]byte{[]byte("0"), []byte("false")}
	var out []byte
	b.SetBytes(int64(len(doc)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out = StripDefaults(doc, out[:0], defaults, nil, RemoveWhitespace)
		if len(out) == 0 {
			b.Fatal("empty result")
		}
	}
}
