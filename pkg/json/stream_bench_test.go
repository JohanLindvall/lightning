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
