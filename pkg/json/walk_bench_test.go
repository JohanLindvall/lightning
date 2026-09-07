package json

import (
	"bytes"
	"fmt"
	"testing"
)

// The walker benchmarks. ObjectEach and ArrayEach had no committed benchmark
// on compact input before — only BenchmarkObjectEachPretty, whose whitespace
// runs hide the per-member costs — so any change to the member loop (a null
// probe on the container, an ErrStop test on the callback's error) was
// measurable only indirectly. These are the shapes callers actually walk: a
// wide flat log record, an array of small records, an array of bare scalars,
// and the [timestamp, value] series pair.

var walkRecord = []byte(`{"SmartRouteColoID":994,"CacheCacheStatus":"dynamic","CacheResponseBytes":4129,"CacheResponseStatus":200,"ClientIP":"203.0.113.23","ClientASN":64500,"ClientCountry":"na","ClientDeviceType":"mobile","ClientRegionCode":"KH","ClientSSLCipher":"NONE","ClientRequestBytes":8935,"ClientRequestHost":"api-internal.example.com","ClientRequestMethod":"POST","ClientRequestPath":"/graphql","ClientRequestProtocol":"HTTP/1.1","ClientRequestReferer":"https://app.example.org/","ClientRequestScheme":"https","ClientRequestSource":"edgeWorkerFetch","ClientRequestURI":"/graphql","ClientRequestUserAgent":"Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Mobile Safari/537.36","EdgeColoCode":"CPT","EdgeColoID":113,"EdgeEndTimestamp":"2025-03-18T12:36:53Z","EdgePathingOp":"wl","EdgePathingSrc":"macro","EdgePathingStatus":"nr","EdgeRequestHost":"api-internal.example.com","EdgeResponseBodyBytes":61,"EdgeResponseCompressionRatio":1,"EdgeResponseContentType":"application/graphql-response+json; charset=utf-8","EdgeResponseBytes":1088,"EdgeServerIP":"192.0.2.85","EdgeTimeToFirstByteMs":201,"EdgeResponseStatus":599,"EdgeStartTimestamp":"2025-03-18T12:36:53Z","OriginIP":"198.51.100.65","OriginSSLProtocol":"TLSv1.3","OriginTCPHandshakeDurationMs":3,"OriginTLSHandshakeDurationMs":4,"OriginResponseDurationMs":191,"OriginResponseHeaderReceiveDurationMs":26,"OriginResponseStatus":520,"OriginResponseTime":33357744,"WorkerSubrequest":true,"RayID":"1111111111111111","ParentRayID":"2222222222222222","RequestHeaders":{"accept-language":"en-NA,en-US;q=0.9,en;q=0.8","x-domain-brand":"acme","x-domain-environment":"live","x-forwarded-host":"backend.example.com","x-location-code":"NA","x-product-areas":"sports,casino","x-worker-name":"graphql"},"ResponseHeaders":{"server":"cloudflare"}}`)

func walkRecordArray(n int) []byte {
	var sb bytes.Buffer
	sb.WriteByte('[')
	for i := 0; i < n; i++ {
		if i > 0 {
			sb.WriteByte(',')
		}
		fmt.Fprintf(&sb, `{"name":"record number %03d","value":%d,"active":true}`, i, i*37)
	}
	sb.WriteByte(']')
	return sb.Bytes()
}

func walkScalarArray(n int) []byte {
	var sb bytes.Buffer
	sb.WriteByte('[')
	for i := 0; i < n; i++ {
		if i > 0 {
			sb.WriteByte(',')
		}
		fmt.Fprintf(&sb, "%d", 1788087600+i*60)
	}
	sb.WriteByte(']')
	return sb.Bytes()
}

func walkSeries(n int) []byte {
	var sb bytes.Buffer
	sb.WriteByte('[')
	for i := 0; i < n; i++ {
		if i > 0 {
			sb.WriteByte(',')
		}
		fmt.Fprintf(&sb, `[%d,"%d.%d"]`, 1788087600+i*60, i, i)
	}
	sb.WriteByte(']')
	return sb.Bytes()
}

func walkStringArray(n int) []byte {
	var sb bytes.Buffer
	sb.WriteByte('[')
	for i := 0; i < n; i++ {
		if i > 0 {
			sb.WriteByte(',')
		}
		fmt.Fprintf(&sb, `"entity-service-%05d"`, i)
	}
	sb.WriteByte(']')
	return sb.Bytes()
}

var walkSink int

func BenchmarkArrayEachStrings(b *testing.B) {
	doc := walkStringArray(100)
	b.SetBytes(int64(len(doc)))
	b.ReportAllocs()
	n := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := ArrayEachCompact(doc, func(v []byte) error {
			n += len(v)
			return nil
		}); err != nil {
			b.Fatal(err)
		}
	}
	walkSink = n
}

func BenchmarkObjectEachRecord(b *testing.B) {
	doc := walkRecord
	b.SetBytes(int64(len(doc)))
	b.ReportAllocs()
	n := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := ObjectEach(doc, func(k string, v []byte) error {
			n += len(k) + len(v)
			return nil
		}); err != nil {
			b.Fatal(err)
		}
	}
	walkSink = n
}

func BenchmarkObjectEachRecordCompact(b *testing.B) {
	doc := walkRecord
	b.SetBytes(int64(len(doc)))
	b.ReportAllocs()
	n := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := ObjectEachCompact(doc, func(k string, v []byte) error {
			n += len(k) + len(v)
			return nil
		}); err != nil {
			b.Fatal(err)
		}
	}
	walkSink = n
}

func BenchmarkObjectEachNested(b *testing.B) {
	doc := walkRecord
	b.SetBytes(int64(len(doc)))
	b.ReportAllocs()
	n := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := ObjectEachCompact(doc, func(k string, v []byte) error {
			n += len(k) + len(v)
			return nil
		}, "RequestHeaders"); err != nil {
			b.Fatal(err)
		}
	}
	walkSink = n
}

func BenchmarkArrayEachRecords(b *testing.B) {
	doc := walkRecordArray(50)
	b.SetBytes(int64(len(doc)))
	b.ReportAllocs()
	n := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := ArrayEachCompact(doc, func(v []byte) error {
			n += len(v)
			return nil
		}); err != nil {
			b.Fatal(err)
		}
	}
	walkSink = n
}

func BenchmarkArrayEachScalars(b *testing.B) {
	doc := walkScalarArray(200)
	b.SetBytes(int64(len(doc)))
	b.ReportAllocs()
	n := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := ArrayEachCompact(doc, func(v []byte) error {
			n += len(v)
			return nil
		}); err != nil {
			b.Fatal(err)
		}
	}
	walkSink = n
}

func BenchmarkArrayEachSeries(b *testing.B) {
	doc := walkSeries(64)
	b.SetBytes(int64(len(doc)))
	b.ReportAllocs()
	n := 0
	b.ResetTimer()
	// Both closures are made once: an inner closure built per element would
	// allocate per element and measure that instead of the walk.
	inner := func(v []byte) error {
		n += len(v)
		return nil
	}
	outer := func(pt []byte) error { return ArrayEachCompact(pt, inner) }
	for i := 0; i < b.N; i++ {
		if err := ArrayEachCompact(doc, outer); err != nil {
			b.Fatal(err)
		}
	}
	walkSink = n
}
