package json

import (
	"bytes"
	stdjson "encoding/json"
	"strconv"
	"testing"
)

// Keep short and long arrays together: saving growth on a long array must not
// add allocations to the much more common small arrays nested in records.
func perfArray(n int, value string) []byte {
	doc := []byte{'['}
	for i := 0; i < n; i++ {
		if i != 0 {
			doc = append(doc, ',')
		}
		doc = append(doc, value...)
	}
	return append(doc, ']')
}

func BenchmarkDecodeAnyArrays(b *testing.B) {
	for _, shape := range []struct{ name, value string }{
		{"numbers", "12345.5"},
		{"strings", `"a short string"`},
		{"records", `{"id":123,"tags":["a","b"],"active":true}`},
	} {
		for _, n := range []int{0, 1, 4, 16, 17, 256, 4096} {
			doc := perfArray(n, shape.value)
			b.Run(shape.name+"/"+strconv.Itoa(n), func(b *testing.B) {
				b.SetBytes(int64(len(doc)))
				b.ReportAllocs()
				for b.Loop() {
					if _, err := DecodeAnyCompact(doc); err != nil {
						b.Fatal(err)
					}
				}
			})
		}
	}
}

func BenchmarkValidShapes(b *testing.B) {
	for _, shape := range []struct {
		name string
		doc  []byte
	}{
		{"scalar", []byte("12345.5")},
		{"empty_object", []byte("{}")},
		{"record", walkRecord},
		{"records", walkRecordArray(50)},
		{"strings", walkStringArray(100)},
		{"numbers", walkScalarArray(200)},
		{"escapes", perfArray(100, `"line\nbreak\t\u1234\ud83d\ude00"`)},
		{"deep", append(append(bytes.Repeat([]byte{'['}, 128), '0'), bytes.Repeat([]byte{']'}, 128)...)},
	} {
		b.Run(shape.name, func(b *testing.B) {
			b.SetBytes(int64(len(shape.doc)))
			b.ReportAllocs()
			for b.Loop() {
				if !Valid(shape.doc) {
					b.Fatal("invalid document")
				}
			}
		})
		if shape.name == "record" {
			var pretty bytes.Buffer
			if err := stdjson.Indent(&pretty, shape.doc, "", "  "); err != nil {
				b.Fatal(err)
			}
			b.Run("pretty_record", func(b *testing.B) {
				doc := pretty.Bytes()
				b.SetBytes(int64(len(doc)))
				b.ReportAllocs()
				for b.Loop() {
					if !Valid(doc) {
						b.Fatal("invalid document")
					}
				}
			})
		}
	}
}
