package json

import (
	"strings"
	"testing"
)

// Shape-parameterised benchmarks for the readers added on top of the walkers.
// A single shape hides where the cost is: an integer's cost is a function of
// its digit count, a string's of its length and whether it holds an escape,
// KindOf's of which arm of its dispatch a value takes.

var (
	intSink    int64
	uintSink   uint64
	strSink    string
	boolSink   bool
	kindSink   Kind
	errSinkNew error
)

func BenchmarkParseIntShapes(b *testing.B) {
	cases := []struct{ name, in string }{
		{"1digit", "7"},
		{"3digit", "200"},
		{"5digit", "64500"},
		{"10digit", "1788087600"},
		{"13digit", "1788087600123"},
		{"16digit", "1788087600123456"},
		{"19digit", "9223372036854775807"},
		{"neg10digit", "-1788087600"},
		{"20digit_overflow", "99999999999999999999"},
		{"notanint", "1.5"},
	}
	for _, c := range cases {
		in := []byte(c.in)
		b.Run(c.name, func(b *testing.B) {
			b.ReportAllocs()
			var v int64
			var err error
			for i := 0; i < b.N; i++ {
				v, err = ParseInt(in)
			}
			intSink, errSinkNew = v, err
		})
	}
}

func BenchmarkParseUintShapes(b *testing.B) {
	cases := []struct{ name, in string }{
		{"3digit", "200"},
		{"10digit", "1788087600"},
		{"13digit", "1788087600123"},
		{"20digit", "18446744073709551615"},
	}
	for _, c := range cases {
		in := []byte(c.in)
		b.Run(c.name, func(b *testing.B) {
			b.ReportAllocs()
			var v uint64
			var err error
			for i := 0; i < b.N; i++ {
				v, err = ParseUint(in)
			}
			uintSink, errSinkNew = v, err
		})
	}
}

func BenchmarkStringShapes(b *testing.B) {
	cases := []struct{ name, in string }{
		{"empty", `""`},
		{"short", `"dynamic"`},
		{"medium", `"entity-service-20260903-1"`},
		{"long", `"` + strings.Repeat("abcdefgh", 16) + `"`},
		{"escaped", `"a\nbéc"`},
		{"notastring", `12345`},
	}
	for _, c := range cases {
		in := []byte(c.in)
		b.Run(c.name, func(b *testing.B) {
			b.ReportAllocs()
			var v string
			var err error
			for i := 0; i < b.N; i++ {
				v, err = String(in)
			}
			strSink, errSinkNew = v, err
		})
	}
}

func BenchmarkBoolShapes(b *testing.B) {
	for _, c := range []struct{ name, in string }{{"true", "true"}, {"false", "false"}, {"null", "null"}} {
		in := []byte(c.in)
		b.Run(c.name, func(b *testing.B) {
			b.ReportAllocs()
			var v bool
			var err error
			for i := 0; i < b.N; i++ {
				v, err = Bool(in)
			}
			boolSink, errSinkNew = v, err
		})
	}
}

func BenchmarkKindOfShapes(b *testing.B) {
	cases := []struct{ name, in string }{
		{"string", `"entity-service-20260903-1"`},
		{"number", `1788087600123`},
		{"object", `{"a":1,"b":2}`},
		{"array", `[1,2,3]`},
		{"null", `null`},
		{"true", `true`},
		{"invalid", `nul`},
		{"ws_number", "   1788087600123"},
	}
	for _, c := range cases {
		in := []byte(c.in)
		b.Run(c.name, func(b *testing.B) {
			b.ReportAllocs()
			var v Kind
			for i := 0; i < b.N; i++ {
				v = KindOf(in)
			}
			kindSink = v
		})
	}
}

func BenchmarkUnescapeCopyShapes(b *testing.B) {
	cases := []struct{ name, in string }{
		{"short", `dynamic`},
		{"medium", `entity-service-20260903-1`},
		{"long", strings.Repeat("abcdefgh", 16)},
		{"escaped", `a\nbéc`},
		{"long_late_escape", strings.Repeat("abcdefgh", 16) + `\n`},
	}
	for _, c := range cases {
		in := []byte(c.in)
		b.Run(c.name, func(b *testing.B) {
			b.ReportAllocs()
			var v string
			var err error
			for i := 0; i < b.N; i++ {
				v, err = UnescapeStringCopy(in)
			}
			strSink, errSinkNew = v, err
		})
	}
}

// BenchmarkArrayEachIndexVsEach holds the indexed walker against the plain one
// on the same document: the two are separate code, so this is what says the
// index costs nothing.
func BenchmarkArrayEachIndexVsEach(b *testing.B) {
	doc := walkScalarArray(200)
	n := 0
	plain := func(v []byte) error { n += len(v); return nil }
	indexed := func(_ int, v []byte) error { n += len(v); return nil }
	b.Run("each", func(b *testing.B) {
		b.SetBytes(int64(len(doc)))
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			if err := ArrayEachCompact(doc, plain); err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("index", func(b *testing.B) {
		b.SetBytes(int64(len(doc)))
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			if err := ArrayEachIndexCompact(doc, indexed); err != nil {
				b.Fatal(err)
			}
		}
	})
	walkSink = n
}

// BenchmarkErrStop measures the early exit: a walk that stops at the first
// element against one that runs to the end.
func BenchmarkErrStop(b *testing.B) {
	doc := walkScalarArray(200)
	var first []byte
	stop := func(v []byte) error { first = v; return ErrStop }
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if err := ArrayEachCompact(doc, stop); err != nil {
			b.Fatal(err)
		}
	}
	walkSink = len(first)
}
