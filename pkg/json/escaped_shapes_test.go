package json

import (
	"strings"
	"testing"
)

func escapedStringShapes() []struct{ name, body string } {
	return []struct{ name, body string }{
		{"clean", strings.Repeat("ordinary text ", 64)},
		{"single", strings.Repeat(`\n\t\"\\\/\b\f\r`, 32)},
		{"unicode", strings.Repeat(`\u4e16\u754c`, 64)},
		{"surrogates", strings.Repeat(`\ud83d\ude00`, 64)},
		{"mixed", strings.Repeat(`words\ntext\u4e16\u754c`, 32)},
		{"sparse", strings.Repeat("ordinary text ", 64) + `\nend`},
	}
}

func BenchmarkValidEscapedStrings(b *testing.B) {
	for _, shape := range escapedStringShapes() {
		for _, key := range []bool{false, true} {
			name := shape.name + "/value"
			doc := []byte(`"` + shape.body + `"`)
			if key {
				name = shape.name + "/key"
				doc = []byte(`{"` + shape.body + `":true}`)
			}
			b.Run(name, func(b *testing.B) {
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

func BenchmarkDecodeAnyEscapedStrings(b *testing.B) {
	for _, shape := range escapedStringShapes() {
		doc := []byte(`"` + shape.body + `"`)
		b.Run(shape.name, func(b *testing.B) {
			b.SetBytes(int64(len(doc)))
			b.ReportAllocs()
			for b.Loop() {
				if _, err := DecodeAny(doc); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
