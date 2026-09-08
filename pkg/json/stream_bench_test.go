package json

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
)

// matrixDoc is the shape the streaming reader exists for: a Prometheus-style
// matrix, an array of series each holding its labels and a long array of
// [timestamp, value] pairs. The real ones are hundreds of megabytes; this is
// the same shape at a size a benchmark can hold twice.
func matrixDoc(series, points int) []byte {
	var b bytes.Buffer
	b.WriteString(`{"status":"success","data":{"resultType":"matrix","result":[`)
	for s := 0; s < series; s++ {
		if s > 0 {
			b.WriteByte(',')
		}
		fmt.Fprintf(&b, `{"metric":{"__name__":"http_requests_total","instance":"host-%04d:9100","job":"node","path":"/api/v1/query"},"values":[`, s)
		for p := 0; p < points; p++ {
			if p > 0 {
				b.WriteByte(',')
			}
			fmt.Fprintf(&b, `[%d,"%d.%03d"]`, 1788087600+p*15, p, s)
		}
		b.WriteString(`]}`)
	}
	b.WriteString(`]}}`)
	return b.Bytes()
}

// BenchmarkStreamMatrix is the comparison the streaming reader is for:
// io.ReadAll and then walk, against walking as it arrives. The allocation
// columns are the point — the in-memory form's B/op is the document, the
// stream's is its buffer — and ns/op says what that costs.
func BenchmarkStreamMatrix(b *testing.B) {
	doc := matrixDoc(200, 240)
	b.Logf("document: %.1f MB", float64(len(doc))/(1<<20))
	b.ResetTimer()
	count := func(v []byte) error { streamSink += len(v); return nil }

	b.Run("readall+walk", func(b *testing.B) {
		b.SetBytes(int64(len(doc)))
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			data, err := io.ReadAll(&chunkReader{data: doc, n: 32 << 10})
			if err != nil {
				b.Fatal(err)
			}
			if err := ArrayEach(data, count, "data", "result"); err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("stream", func(b *testing.B) {
		b.SetBytes(int64(len(doc)))
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := NewReader(&chunkReader{data: doc, n: 32 << 10})
			if err := r.ArrayEach(count, "data", "result"); err != nil {
				b.Fatal(err)
			}
		}
	})
	// One reader across a stream of documents, which is what a client that
	// makes the same query repeatedly does: the buffer is paid for once.
	b.Run("stream_reused", func(b *testing.B) {
		b.SetBytes(int64(len(doc)))
		b.ReportAllocs()
		r := NewReader(&chunkReader{})
		for i := 0; i < b.N; i++ {
			r.Reset(&chunkReader{data: doc, n: 32 << 10})
			if err := r.ArrayEach(count, "data", "result"); err != nil {
				b.Fatal(err)
			}
		}
	})
	// The nested walk a caller actually writes: every point of every series.
	b.Run("stream_points", func(b *testing.B) {
		b.SetBytes(int64(len(doc)))
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := NewReader(&chunkReader{data: doc, n: 32 << 10})
			if err := r.ArrayEach(func(series []byte) error {
				return ArrayEach(series, count, "values")
			}, "data", "result"); err != nil {
				b.Fatal(err)
			}
		}
	})
}

var streamSink int

// BenchmarkStreamShapes is the throughput of the streaming walk against the
// in-memory one over the element kinds whose skip differs, with the reader
// handing out realistic block sizes.
func BenchmarkStreamShapes(b *testing.B) {
	docs := map[string][]byte{
		"scalars": walkScalarArray(2000),
		"strings": walkStringArray(1000),
		"records": walkRecordArray(500),
	}
	count := func(v []byte) error { streamSink += len(v); return nil }
	for name, doc := range docs {
		b.Run(name+"/inmemory", func(b *testing.B) {
			b.SetBytes(int64(len(doc)))
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				if err := ArrayEach(doc, count); err != nil {
					b.Fatal(err)
				}
			}
		})
		b.Run(name+"/stream", func(b *testing.B) {
			b.SetBytes(int64(len(doc)))
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				r := NewReader(&chunkReader{data: doc, n: 32 << 10})
				if err := r.ArrayEach(count); err != nil {
					b.Fatal(err)
				}
			}
		})
		// The same walk with the Reader reused, which is what separates the
		// WALK from the setup: a fresh Reader allocates the buffer, and on a
		// document of a few kilobytes that allocation is most of the row above
		// — 3.3 us of the 3.5 a NewReader costs is the 64 KiB make, two thirds
		// of which is the GC work an allocation rate buys. Read a fresh-Reader
		// row against its in-memory twin only with this one beside it.
		b.Run(name+"/stream_reused", func(b *testing.B) {
			b.SetBytes(int64(len(doc)))
			b.ReportAllocs()
			r := NewReader(&chunkReader{})
			for i := 0; i < b.N; i++ {
				r.Reset(&chunkReader{data: doc, n: 32 << 10})
				if err := r.ArrayEach(count); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// BenchmarkStreamSkipToKey measures the descent: the wanted key sits after a
// megabyte of members the walk never buffers.
func BenchmarkStreamSkipToKey(b *testing.B) {
	var sb strings.Builder
	sb.WriteString(`{"skip":[`)
	for i := 0; i < 4000; i++ {
		if i > 0 {
			sb.WriteByte(',')
		}
		fmt.Fprintf(&sb, `{"pad":"%s","n":%d}`, strings.Repeat("p", 200), i)
	}
	sb.WriteString(`],"want":{"deep":[1,2,3]}}`)
	doc := []byte(sb.String())
	b.SetBytes(int64(len(doc)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r := NewReader(&chunkReader{data: doc, n: 32 << 10})
		if err := r.ArrayEach(func(v []byte) error { streamSink += len(v); return nil }, "want", "deep"); err != nil {
			b.Fatal(err)
		}
	}
}

// recordDoc is a flat record of the width an API response has: the shape a
// member walk pays a key read and a separator for, once each per member.
func recordDoc(members int) []byte {
	var b bytes.Buffer
	b.WriteByte('{')
	for i := 0; i < members; i++ {
		if i > 0 {
			b.WriteByte(',')
		}
		fmt.Fprintf(&b, `"field_number_%02d":"value-%02d"`, i, i)
	}
	b.WriteByte('}')
	return b.Bytes()
}

// descentDoc is a record whose wanted key is last, with members that are
// themselves small objects: every one of them is read and stepped over, and
// every one of them is buffered — which is what separates this from
// [BenchmarkStreamSkipToKey], whose single skipped sibling is far larger than
// the buffer.
func descentDoc(members int) []byte {
	var b bytes.Buffer
	b.WriteByte('{')
	for i := 0; i < members; i++ {
		fmt.Fprintf(&b, `"field_%02d":{"id":%d,"name":"value-%02d","tags":["a","b"]},`, i, i*7919, i)
	}
	b.WriteString(`"want":{"deep":[1,2,3]}}`)
	return b.Bytes()
}

// largeElementDoc holds elements several times the buffer, which is the case
// the resumable scanner exists for: each one has to be assembled across many
// refills before its callback can see it.
func largeElementDoc(elems, rows int) []byte {
	var b bytes.Buffer
	b.WriteByte('[')
	for i := 0; i < elems; i++ {
		if i > 0 {
			b.WriteByte(',')
		}
		fmt.Fprintf(&b, `{"id":%d,"rows":[`, i)
		for j := 0; j < rows; j++ {
			if j > 0 {
				b.WriteByte(',')
			}
			fmt.Fprintf(&b, `{"t":%d,"v":"%d.%03d","tag":"host-%04d"}`, 1788087600+j, j, i, j)
		}
		b.WriteString(`]}`)
	}
	b.WriteByte(']')
	return b.Bytes()
}

// BenchmarkStreamObjectEach is the member walk against the in-memory one over
// the same record. The reader is reused, as a client making the same request
// repeatedly does, so what is measured is the walk rather than the buffer —
// [BenchmarkStreamMatrix] measures a reader per document as well.
func BenchmarkStreamObjectEach(b *testing.B) {
	doc := recordDoc(45)
	fn := func(k string, v []byte) error { streamSink += len(k) + len(v); return nil }
	b.Run("stream", func(b *testing.B) {
		b.SetBytes(int64(len(doc)))
		b.ReportAllocs()
		r := NewReader(&chunkReader{})
		for i := 0; i < b.N; i++ {
			r.Reset(&chunkReader{data: doc, n: 32 << 10})
			if err := r.ObjectEach(fn); err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("inmemory", func(b *testing.B) {
		b.SetBytes(int64(len(doc)))
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			if err := ObjectEach(doc, fn); err != nil {
				b.Fatal(err)
			}
		}
	})
}

// BenchmarkStreamDescent is the path descent every Get and every keyed walk
// begins with: the members before the wanted key are read and stepped over.
func BenchmarkStreamDescent(b *testing.B) {
	doc := descentDoc(60)
	count := func(v []byte) error { streamSink += len(v); return nil }
	b.Run("get", func(b *testing.B) {
		b.SetBytes(int64(len(doc)))
		b.ReportAllocs()
		r := NewReader(&chunkReader{})
		for i := 0; i < b.N; i++ {
			r.Reset(&chunkReader{data: doc, n: 32 << 10})
			v, err := r.Get("want", "deep")
			if err != nil {
				b.Fatal(err)
			}
			streamSink += len(v)
		}
	})
	b.Run("arrayeach", func(b *testing.B) {
		b.SetBytes(int64(len(doc)))
		b.ReportAllocs()
		r := NewReader(&chunkReader{})
		for i := 0; i < b.N; i++ {
			r.Reset(&chunkReader{data: doc, n: 32 << 10})
			if err := r.ArrayEach(count, "want", "deep"); err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("inmemory", func(b *testing.B) {
		b.SetBytes(int64(len(doc)))
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			if err := ArrayEach(doc, count, "want", "deep"); err != nil {
				b.Fatal(err)
			}
		}
	})
}

// BenchmarkStreamLargeElements walks elements far larger than the buffer, so
// every one of them is assembled by the resumable scanner across refills. It
// is the shape that says whether a value the walk DOES want is scanned once or
// several times over.
func BenchmarkStreamLargeElements(b *testing.B) {
	doc := largeElementDoc(8, 3000)
	b.Logf("element: %.0f KB, buffer: 64 KB", float64(len(doc))/8/1024)
	b.SetBytes(int64(len(doc)))
	b.ReportAllocs()
	b.ResetTimer()
	r := NewReader(&chunkReader{})
	for i := 0; i < b.N; i++ {
		r.Reset(&chunkReader{data: doc, n: 32 << 10})
		if err := r.ArrayEach(func(v []byte) error { streamSink += len(v); return nil }); err != nil {
			b.Fatal(err)
		}
	}
}
