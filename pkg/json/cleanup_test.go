package json

import (
	"bytes"
	"encoding/json"
	"testing"
)

// origDefaults / origKeep reproduce the hardcoded lists from the imported
// cleanup.go so we can verify the parameterized port matches its behavior.
var origDefaults = [][]byte{
	[]byte("0"), []byte("00"), []byte("none"),
	[]byte("false"), []byte("unknown"), []byte("noRecord"),
}
var origKeep = [][]byte{
	[]byte("WallTimeMs"), []byte("CPUTimeMs"), []byte("WorkerCPUTime"),
	[]byte("WorkerWallTimeUs"), []byte("EdgeTimeToFirstByteMs"),
}

func clean(t *testing.T, in string) string {
	t.Helper()
	return string(Cleanup([]byte(in), nil, origDefaults, origKeep, false))
}

func TestCleanup(t *testing.T) {
	cases := []struct{ in, want string }{
		{`{}`, ``},
		{`{"a":1}`, `{"a":1}`},
		{`{"a":0}`, ``},                                // default value -> field dropped -> object empty
		{`{"a":0,"b":2}`, `{"b":2}`},                   // drop only the default
		{`{"a":1,"b":0,"c":3}`, `{"a":1,"c":3}`},       // drop middle, comma handling
		{`{"a":"none","b":"x"}`, `{"b":"x"}`},          // string default
		{`{"a":"","b":"x"}`, `{"b":"x"}`},              // empty string is default
		{`{"a":false,"b":true}`, `{"b":true}`},         // false default, true kept
		{`{"a":{}}`, ``},                               // empty object dropped
		{`{"a":[]}`, ``},                               // empty array dropped
		{`{"a":{"b":0}}`, ``},                          // nested all-default collapses
		{`{"a":{"b":0,"c":5}}`, `{"a":{"c":5}}`},       // nested partial
		{`[0,1,0,2]`, `[1,2]`},                         // array drops default scalars
		{`{"WallTimeMs":0,"x":0}`, `{"WallTimeMs":0}`}, // keep-list retains default
		{`{"unknown":"unknown"}`, ``},                  // value default (key not in keep)
		{`  {  "a" : 1 , "b":0 }  `, `{"a":1}`},        // whitespace tolerated
	}
	for _, c := range cases {
		if got := clean(t, c.in); got != c.want {
			t.Errorf("Cleanup(%q) = %q, want %q", c.in, got, c.want)
		}
		// Cleaned output that is non-empty must be valid JSON.
		if got := clean(t, c.in); got != "" && !json.Valid([]byte(got)) {
			t.Errorf("Cleanup(%q) = %q is not valid JSON", c.in, got)
		}
	}
}

func TestCleanupDoesNotMutateInput(t *testing.T) {
	in := []byte(`{"a":0,"b":2}`)
	orig := append([]byte(nil), in...)
	out := Cleanup(in, make([]byte, 0, 64), origDefaults, origKeep, false)
	if !bytes.Equal(in, orig) {
		t.Fatalf("input mutated: got %q want %q", in, orig)
	}
	if string(out) != `{"b":2}` {
		t.Fatalf("out = %q", out)
	}
}

func TestCleanupInPlace(t *testing.T) {
	in := []byte(`{"a":0,"b":2,"c":0}`)
	out := Cleanup(in, in[:0], origDefaults, origKeep, false)
	if string(out) != `{"b":2}` {
		t.Fatalf("in-place out = %q", out)
	}
}

func TestCleanupGrowsOutput(t *testing.T) {
	in := []byte(`{"a":1,"b":2}`)
	out := Cleanup(in, make([]byte, 0, 2), origDefaults, origKeep, false) // too small -> allocates
	if string(out) != `{"a":1,"b":2}` {
		t.Fatalf("out = %q", out)
	}
}

// cfLogLine is a representative Cloudflare HTTP-request log line — the shape the
// imported cleanup pass was written for, dense with default/empty fields
// (0, "", "none", "false", "unknown", "noRecord", [] ) interleaved with real
// values and a few keep-list keys (WorkerCPUTime, WorkerWallTimeUs,
// EdgeTimeToFirstByteMs) whose default value must survive.
var cfLogLine = []byte(`{"Cookies":{"clientid":"00000000-0000-4000-8000-000000000001","visitid":"00000000-0000-4000-8000-000000000002"},"LeakedCredentialCheckResult":"none","ContentScanObjTypes":[],"ContentScanObjResults":[],"SecuritySources":[],"SmartRouteColoID":0,"UpperTierColoID":0,"SecurityRuleIDs":[],"SecurityRuleID":"","SecurityRuleDescription":"","SecurityActions":[],"SecurityAction":"","CacheCacheStatus":"unknown","CacheReserveUsed":false,"CacheTieredFill":false,"CacheResponseBytes":0,"CacheResponseStatus":0,"ClientIP":"198.51.100.123","ClientASN":64512,"ClientCountry":"zz","ClientDeviceType":"desktop","ClientIPClass":"noRecord","ClientMTLSAuthCertFingerprint":"","ClientMTLSAuthStatus":"unknown","ClientRegionCode":"","ClientSSLCipher":"AEAD-AES128-GCM-SHA256","ClientSSLProtocol":"TLSv1.3","ClientSrcPort":50000,"ClientTCPRTTMs":18,"ClientXRequestedWith":"","ClientRequestBytes":7769,"ClientRequestHost":"app-001--x-demo--exampleco-zz.dev.examplebackend.test","ClientRequestMethod":"POST","ClientRequestPath":"/services/api/Configuration/v2/GetExamplePollingTimes","ClientRequestProtocol":"HTTP/2","ClientRequestReferer":"https://app-001--x-demo--exampleco-zz.dev.examplebackend.test/zz/section/samples","ClientRequestScheme":"https","ClientRequestSource":"eyeball","ClientRequestURI":"/services/api/Configuration/v2/GetExamplePollingTimes","ClientRequestUserAgent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36","EdgeCFConnectingO2O":false,"EdgeColoCode":"AMS","EdgeColoID":100,"EdgeEndTimestamp":"2024-05-15T18:42:30Z","EdgePathingOp":"wl","EdgePathingSrc":"macro","EdgePathingStatus":"nr","EdgeRequestHost":"","EdgeResponseBodyBytes":143,"EdgeResponseCompressionRatio":1,"EdgeResponseContentType":"text/html","EdgeResponseBytes":2322,"EdgeServerIP":"","EdgeTimeToFirstByteMs":31,"EdgeResponseStatus":302,"EdgeStartTimestamp":"2024-05-15T18:42:30Z","OriginDNSResponseTimeMs":0,"OriginIP":"","OriginRequestHeaderSendDurationMs":0,"OriginSSLProtocol":"unknown","OriginTCPHandshakeDurationMs":0,"OriginTLSHandshakeDurationMs":0,"OriginResponseBytes":0,"OriginResponseDurationMs":0,"OriginResponseHTTPExpires":"","OriginResponseHTTPLastModified":"","OriginResponseHeaderReceiveDurationMs":0,"OriginResponseStatus":0,"OriginResponseTime":0,"WAFAttackScore":0,"WAFFlags":"0","WAFMatchedVar":"","WAFRCEAttackScore":0,"WAFSQLiAttackScore":0,"WAFXSSAttackScore":0,"WorkerCPUTime":0,"WorkerStatus":"unknown","WorkerSubrequest":false,"WorkerSubrequestCount":0,"WorkerWallTimeUs":0,"RayID":"0a1b2c3d4e5f6071","ParentRayID":"00","RequestHeaders":{"accept-language":"en-GB,en-US;q=0.9,en;q=0.8","traceparent":"00-0af1c2d3e4b5a6978899aabbccddeeff-1122334455667788-01"},"ResponseHeaders":{"server":"cloudflare"}}`)

// benchmarkCleanup runs the cleanup over the (compact) Cloudflare log line with
// inter-token whitespace skipping on or off. A reusable output buffer sized to
// the input means no per-op allocation (Cleanup never lengthens the document)
// and input is left unmodified, so the same input drives every iteration.
func benchmarkCleanup(b *testing.B, compact bool) {
	out := make([]byte, 0, len(cfLogLine))
	b.SetBytes(int64(len(cfLogLine)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out = Cleanup(cfLogLine, out[:0], origDefaults, origKeep, compact)
	}
	_ = out
}

func BenchmarkCleanup(b *testing.B)        { benchmarkCleanup(b, false) }
func BenchmarkCleanupCompact(b *testing.B) { benchmarkCleanup(b, true) }
