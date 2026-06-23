package json

import (
	"testing"
)

func BenchmarkGetManyWithSkip(b *testing.B) {
	data := []byte(`{"SmartRouteColoID":994,"CacheCacheStatus":"dynamic","CacheResponseBytes":4129,"CacheResponseStatus":200,"ClientIP":"203.0.113.23","ClientASN":64500,"ClientCountry":"na","ClientDeviceType":"mobile","ClientRegionCode":"KH","ClientSSLCipher":"NONE","ClientRequestBytes":8935,"ClientRequestHost":"api-internal.example.com","ClientRequestMethod":"POST","ClientRequestPath":"/graphql","ClientRequestProtocol":"HTTP/1.1","ClientRequestReferer":"https://app.example.org/","ClientRequestScheme":"https","ClientRequestSource":"edgeWorkerFetch","ClientRequestURI":"/graphql","ClientRequestUserAgent":"Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Mobile Safari/537.36","EdgeColoCode":"CPT","EdgeColoID":113,"EdgeEndTimestamp":"2025-03-18T12:36:53Z","EdgePathingOp":"wl","EdgePathingSrc":"macro","EdgePathingStatus":"nr","EdgeRequestHost":"api-internal.example.com","EdgeResponseBodyBytes":61,"EdgeResponseCompressionRatio":1,"EdgeResponseContentType":"application/graphql-response+json; charset=utf-8","EdgeResponseBytes":1088,"EdgeServerIP":"192.0.2.85","EdgeTimeToFirstByteMs":201,"EdgeResponseStatus":599,"EdgeStartTimestamp":"2025-03-18T12:36:53Z","OriginIP":"198.51.100.65","OriginSSLProtocol":"TLSv1.3","OriginTCPHandshakeDurationMs":3,"OriginTLSHandshakeDurationMs":4,"OriginResponseDurationMs":191,"OriginResponseHeaderReceiveDurationMs":26,"OriginResponseStatus":520,"OriginResponseTime":33357744,"WorkerSubrequest":true,"RayID":"1111111111111111","ParentRayID":"2222222222222222","RequestHeaders":{"accept-language":"en-NA,en-US;q=0.9,en;q=0.8","x-domain-brand":"acme","x-domain-environment":"live","x-forwarded-host":"backend.example.com","x-location-code":"NA","x-product-areas":"sports,casino","x-worker-name":"graphql"},"ResponseHeaders":{"server":"cloudflare"}}`)
	keys := []string{"OriginTCPHandshakeDurationMs", "OriginTLSHandshakeDurationMs", "OriginResponseDurationMs"}

	b.SetBytes(int64(len(data)))
	b.ReportAllocs()
	out := make([][]byte, 0)
	for i := 0; i < b.N; i++ {
		v, err := GetMany(data, keys, out[0:])
		out = v
		if err != nil {
			b.Fatalf("got %q err %v", v, err)
		}
	}
}

func BenchmarkGetPathsWithSkip(b *testing.B) {
	data := []byte(`{"SmartRouteColoID":994,"CacheCacheStatus":"dynamic","CacheResponseBytes":4129,"CacheResponseStatus":200,"ClientIP":"203.0.113.23","ClientASN":64500,"ClientCountry":"na","ClientDeviceType":"mobile","ClientRegionCode":"KH","ClientSSLCipher":"NONE","ClientRequestBytes":8935,"ClientRequestHost":"api-internal.example.com","ClientRequestMethod":"POST","ClientRequestPath":"/graphql","ClientRequestProtocol":"HTTP/1.1","ClientRequestReferer":"https://app.example.org/","ClientRequestScheme":"https","ClientRequestSource":"edgeWorkerFetch","ClientRequestURI":"/graphql","ClientRequestUserAgent":"Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Mobile Safari/537.36","EdgeColoCode":"CPT","EdgeColoID":113,"EdgeEndTimestamp":"2025-03-18T12:36:53Z","EdgePathingOp":"wl","EdgePathingSrc":"macro","EdgePathingStatus":"nr","EdgeRequestHost":"api-internal.example.com","EdgeResponseBodyBytes":61,"EdgeResponseCompressionRatio":1,"EdgeResponseContentType":"application/graphql-response+json; charset=utf-8","EdgeResponseBytes":1088,"EdgeServerIP":"192.0.2.85","EdgeTimeToFirstByteMs":201,"EdgeResponseStatus":599,"EdgeStartTimestamp":"2025-03-18T12:36:53Z","OriginIP":"198.51.100.65","OriginSSLProtocol":"TLSv1.3","OriginTCPHandshakeDurationMs":3,"OriginTLSHandshakeDurationMs":4,"OriginResponseDurationMs":191,"OriginResponseHeaderReceiveDurationMs":26,"OriginResponseStatus":520,"OriginResponseTime":33357744,"WorkerSubrequest":true,"RayID":"1111111111111111","ParentRayID":"2222222222222222","RequestHeaders":{"accept-language":"en-NA,en-US;q=0.9,en;q=0.8","x-domain-brand":"acme","x-domain-environment":"live","x-forwarded-host":"backend.example.com","x-location-code":"NA","x-product-areas":"sports,casino","x-worker-name":"graphql"},"ResponseHeaders":{"server":"cloudflare"}}`)
	paths := [][]string{
		{"RequestHeaders", "x-domain-brand"},
		{"RequestHeaders", "x-worker-name"},
		{"OriginResponseDurationMs"},
	}

	b.SetBytes(int64(len(data)))
	b.ReportAllocs()
	out := make([][]byte, 0)
	for i := 0; i < b.N; i++ {
		v, err := GetPaths(data, paths, out[0:])
		out = v
		if err != nil {
			b.Fatalf("got %q err %v", v, err)
		}
	}
}
