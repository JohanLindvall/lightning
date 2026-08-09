package json

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"strings"
	"testing"
)

// origDefaults / origKeep reproduce the hardcoded lists from the imported
// default-stripping pass so we can verify the parameterized port matches it. The
// empty entry reproduces its isEmpty case len==0: now that isDefault treats empty
// as default only when "" is listed, the empty token must be opted in explicitly.
var origDefaults = [][]byte{
	[]byte(""), []byte("0"), []byte("00"), []byte("none"),
	[]byte("false"), []byte("unknown"), []byte("noRecord"),
}
var origKeep = [][]byte{
	[]byte("WallTimeMs"), []byte("CPUTimeMs"), []byte("WorkerCPUTime"),
	[]byte("WorkerWallTimeUs"), []byte("EdgeTimeToFirstByteMs"),
}

func strip(t *testing.T, in string) string {
	t.Helper()
	return string(StripDefaults([]byte(in), nil, origDefaults, origKeep, RemoveWhitespace))
}

func TestStripDefaults(t *testing.T) {
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
		if got := strip(t, c.in); got != c.want {
			t.Errorf("StripDefaults(%q) = %q, want %q", c.in, got, c.want)
		}
		// Stripped output that is non-empty must be valid JSON.
		if got := strip(t, c.in); got != "" && !json.Valid([]byte(got)) {
			t.Errorf("StripDefaults(%q) = %q is not valid JSON", c.in, got)
		}
	}
}

func TestStripDefaultsDoesNotMutateInput(t *testing.T) {
	in := []byte(`{"a":0,"b":2}`)
	orig := append([]byte(nil), in...)
	out := StripDefaults(in, make([]byte, 0, 64), origDefaults, origKeep, RemoveWhitespace)
	if !bytes.Equal(in, orig) {
		t.Fatalf("input mutated: got %q want %q", in, orig)
	}
	if string(out) != `{"b":2}` {
		t.Fatalf("out = %q", out)
	}
}

func TestStripDefaultsInPlace(t *testing.T) {
	in := []byte(`{"a":0,"b":2,"c":0}`)
	out := StripDefaults(in, in[:0], origDefaults, origKeep, RemoveWhitespace)
	if string(out) != `{"b":2}` {
		t.Fatalf("in-place out = %q", out)
	}
}

func TestStripDefaultsGrowsOutput(t *testing.T) {
	in := []byte(`{"a":1,"b":2}`)
	out := StripDefaults(in, make([]byte, 0, 2), origDefaults, origKeep, RemoveWhitespace) // too small -> allocates
	if string(out) != `{"a":1,"b":2}` {
		t.Fatalf("out = %q", out)
	}
}

// cfLogLine is a representative Cloudflare HTTP-request log line — the shape the
// imported default-stripping pass was written for, dense with default/empty fields
// (0, "", "none", "false", "unknown", "noRecord", [] ) interleaved with real
// values and a few keep-list keys (WorkerCPUTime, WorkerWallTimeUs,
// EdgeTimeToFirstByteMs) whose default value must survive.
var cfLogLine = []byte(`{"Cookies":{"clientid":"00000000-0000-4000-8000-000000000001","visitid":"00000000-0000-4000-8000-000000000002"},"LeakedCredentialCheckResult":"none","ContentScanObjTypes":[],"ContentScanObjResults":[],"SecuritySources":[],"SmartRouteColoID":0,"UpperTierColoID":0,"SecurityRuleIDs":[],"SecurityRuleID":"","SecurityRuleDescription":"","SecurityActions":[],"SecurityAction":"","CacheCacheStatus":"unknown","CacheReserveUsed":false,"CacheTieredFill":false,"CacheResponseBytes":0,"CacheResponseStatus":0,"ClientIP":"198.51.100.123","ClientASN":64512,"ClientCountry":"zz","ClientDeviceType":"desktop","ClientIPClass":"noRecord","ClientMTLSAuthCertFingerprint":"","ClientMTLSAuthStatus":"unknown","ClientRegionCode":"","ClientSSLCipher":"AEAD-AES128-GCM-SHA256","ClientSSLProtocol":"TLSv1.3","ClientSrcPort":50000,"ClientTCPRTTMs":18,"ClientXRequestedWith":"","ClientRequestBytes":7769,"ClientRequestHost":"app-001--x-demo--exampleco-zz.dev.examplebackend.test","ClientRequestMethod":"POST","ClientRequestPath":"/services/api/Configuration/v2/GetExamplePollingTimes","ClientRequestProtocol":"HTTP/2","ClientRequestReferer":"https://app-001--x-demo--exampleco-zz.dev.examplebackend.test/zz/section/samples","ClientRequestScheme":"https","ClientRequestSource":"eyeball","ClientRequestURI":"/services/api/Configuration/v2/GetExamplePollingTimes","ClientRequestUserAgent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36","EdgeCFConnectingO2O":false,"EdgeColoCode":"AMS","EdgeColoID":100,"EdgeEndTimestamp":"2024-05-15T18:42:30Z","EdgePathingOp":"wl","EdgePathingSrc":"macro","EdgePathingStatus":"nr","EdgeRequestHost":"","EdgeResponseBodyBytes":143,"EdgeResponseCompressionRatio":1,"EdgeResponseContentType":"text/html","EdgeResponseBytes":2322,"EdgeServerIP":"","EdgeTimeToFirstByteMs":31,"EdgeResponseStatus":302,"EdgeStartTimestamp":"2024-05-15T18:42:30Z","OriginDNSResponseTimeMs":0,"OriginIP":"","OriginRequestHeaderSendDurationMs":0,"OriginSSLProtocol":"unknown","OriginTCPHandshakeDurationMs":0,"OriginTLSHandshakeDurationMs":0,"OriginResponseBytes":0,"OriginResponseDurationMs":0,"OriginResponseHTTPExpires":"","OriginResponseHTTPLastModified":"","OriginResponseHeaderReceiveDurationMs":0,"OriginResponseStatus":0,"OriginResponseTime":0,"WAFAttackScore":0,"WAFFlags":"0","WAFMatchedVar":"","WAFRCEAttackScore":0,"WAFSQLiAttackScore":0,"WAFXSSAttackScore":0,"WorkerCPUTime":0,"WorkerStatus":"unknown","WorkerSubrequest":false,"WorkerSubrequestCount":0,"WorkerWallTimeUs":0,"RayID":"0a1b2c3d4e5f6071","ParentRayID":"00","RequestHeaders":{"accept-language":"en-GB,en-US;q=0.9,en;q=0.8","traceparent":"00-0af1c2d3e4b5a6978899aabbccddeeff-1122334455667788-01"},"ResponseHeaders":{"server":"cloudflare"}}`)

// benchmarkStripDefaults runs StripDefaults over the (compact) Cloudflare log line with
// inter-token whitespace skipping on or off. A reusable output buffer sized to
// the input means no per-op allocation (StripDefaults never lengthens the document)
// and input is left unmodified, so the same input drives every iteration.
func benchmarkStripDefaults(b *testing.B, ws WhitespaceMode) {
	out := make([]byte, 0, len(cfLogLine))
	b.SetBytes(int64(len(cfLogLine)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out = StripDefaults(cfLogLine, out[:0], origDefaults, origKeep, ws)
	}
	_ = out
}

func BenchmarkStripDefaults(b *testing.B)        { benchmarkStripDefaults(b, RemoveWhitespace) }
func BenchmarkStripDefaultsCompact(b *testing.B) { benchmarkStripDefaults(b, AssumeCompact) }

// TestStripDefaultsWhitespaceModes checks the three WhitespaceMode behaviors on
// pretty-printed input: AssumeCompact misreads it (documented), RemoveWhitespace
// yields compact output, and PreserveWhitespace keeps the formatting while
// stripping the same fields. The oracle: PreserveWhitespace output must be valid
// JSON, must still contain whitespace, and must compact to exactly the
// RemoveWhitespace result (same data, same fields dropped).
//
// The keep_container document is the shape the oracle used to fail on: a
// keep-listed member whose container value strips to empty is re-emitted from its
// original span, so before the emitKept fix RemoveWhitespace handed back that
// span's interior whitespace and the two outputs disagreed.
func TestStripDefaultsWhitespaceModes(t *testing.T) {
	docs := []struct {
		name   string
		pretty []byte
	}{
		{"plain", []byte("{\n  \"a\": 1,\n  \"b\": 0,\n  \"c\": {\n    \"d\": 0,\n    \"e\": 5\n  },\n  \"f\": [ 0, 7, 0 ]\n}")},
		{"keep_container", []byte("{\n  \"a\": 1,\n  \"WallTimeMs\": {\n    \"d\": 0,\n    \"e\": \"none\"\n  },\n  \"CPUTimeMs\": [ 0, 0 ],\n  \"WorkerCPUTime\": 0,\n  \"z\": { }\n}")},
	}
	for _, d := range docs {
		t.Run(d.name, func(t *testing.T) {
			rw := string(StripDefaults(d.pretty, nil, origDefaults, origKeep, RemoveWhitespace))
			if !json.Valid([]byte(rw)) {
				t.Fatalf("RemoveWhitespace output invalid: %s", rw)
			}
			if strings.ContainsAny(rw, " \n\t") {
				t.Errorf("RemoveWhitespace output still has whitespace: %q", rw)
			}

			pw := StripDefaults(d.pretty, nil, origDefaults, origKeep, PreserveWhitespace)
			if !json.Valid(pw) {
				t.Fatalf("PreserveWhitespace output invalid: %s", pw)
			}
			if !bytes.ContainsAny(pw, "\n") {
				t.Errorf("PreserveWhitespace output lost its formatting: %q", pw)
			}
			// Oracle: compacting the preserved output must equal the compact strip.
			var compacted bytes.Buffer
			if err := json.Compact(&compacted, pw); err != nil {
				t.Fatalf("preserved output not compactable: %v", err)
			}
			if compacted.String() != rw {
				t.Errorf("preserve vs remove differ in data:\n  preserve(compacted)=%q\n  remove           =%q", compacted.String(), rw)
			}
			t.Logf("RemoveWhitespace:   %s", rw)
			t.Logf("PreserveWhitespace: %q", pw)
		})
	}
}

// TestStripDefaultsWhitespaceOnly verifies that with empty defaults/keep nothing
// is treated as default — so RemoveWhitespace only strips inter-token whitespace
// and the result equals encoding/json's own compaction. The inputs include "" and
// 0 to confirm the updated isDefault keeps them when they aren't in defaults
// (empty values are no longer unconditionally stripped). Empty {}/[] are omitted
// because the empty-container rule drops those regardless of defaults.
func TestStripDefaultsWhitespaceOnly(t *testing.T) {
	inputs := []string{
		`{ "a" : 1 , "b" : "x" , "z" : 0 , "e" : "" , "c" : [ 1 , 2 , 3 ] , "d" : { "k" : true } }`,
		"{\n  \"name\": \"value\",\n  \"zero\": 0,\n  \"blank\": \"\",\n  \"nums\": [ 1, 2, 3 ],\n  \"nested\": { \"x\": false }\n}",
		`  [ 1 , "two" , { "k" : 3.5 } , "" , 0 ]  `,
	}
	for _, in := range inputs {
		got := StripDefaults([]byte(in), nil, nil, nil, RemoveWhitespace)
		var want bytes.Buffer
		if err := json.Compact(&want, []byte(in)); err != nil {
			t.Fatalf("json.Compact(%q): %v", in, err)
		}
		if string(got) != want.String() {
			t.Errorf("StripDefaults(%q, RemoveWhitespace) =\n  %q\nwant\n  %q", in, got, want.String())
		}
	}
}

// TestStripDefaultsKeepKeyContainer locks the rewind fix in handle's object
// branch: a keep-listed member whose container value strips to empty must be
// re-emitted WITH its separator comma. The empty-container rewind used to back
// up past the comma written at the loop top, producing {"a":1"b":{...}} —
// invalid JSON.
func TestStripDefaultsKeepKeyContainer(t *testing.T) {
	defaults := [][]byte{[]byte("0")}
	keep := [][]byte{[]byte("b")}
	cases := []struct {
		in, want string
		ws       WhitespaceMode
	}{
		// keep-key member after a kept member: comma must survive.
		{`{"a":1,"b":{"x":0}}`, `{"a":1,"b":{"x":0}}`, RemoveWhitespace},
		// keep-key member first: no leading comma.
		{`{"b":{"x":0},"a":1}`, `{"b":{"x":0},"a":1}`, RemoveWhitespace},
		// array value stripping empty behaves the same.
		{`{"a":1,"b":[0]}`, `{"a":1,"b":[0]}`, RemoveWhitespace},
		// non-kept sibling dropped around the kept container.
		{`{"a":1,"c":0,"b":{"x":0}}`, `{"a":1,"b":{"x":0}}`, RemoveWhitespace},
		{`{"a": 1, "b": {"x": 0}}`, `{"a": 1, "b": {"x": 0}}`, PreserveWhitespace},
	}
	for _, c := range cases {
		got := string(StripDefaults([]byte(c.in), nil, defaults, keep, c.ws))
		if got != c.want {
			t.Errorf("StripDefaults(%q) = %q, want %q", c.in, got, c.want)
		}
		if got != "" && !json.Valid([]byte(got)) {
			t.Errorf("StripDefaults(%q) = %q is not valid JSON", c.in, got)
		}
	}
}

// stripFresh and stripInPlace are the two ways StripDefaults' output buffer can
// be supplied: a buffer of its own, and the input itself (output == input[:0],
// the documented in-place mode). They must agree byte for byte on every input —
// the walk only ever writes behind its own read cursor, so nothing it still needs
// can have been overwritten — which is the property TestStripDefaultsInPlace*
// below check, and the one a keep-listed member whose container value strips to
// empty used to break: that member's key and original bytes are re-read after the
// speculative write of the key, the colon and the recursion's own output.
func stripFresh(in string, defaults, keep [][]byte, ws WhitespaceMode) string {
	return string(StripDefaults([]byte(in), nil, defaults, keep, ws))
}

func stripInPlace(in string, defaults, keep [][]byte, ws WhitespaceMode) string {
	buf := []byte(in)
	return string(StripDefaults(buf, buf[:0], defaults, keep, ws))
}

var stripWSModes = []struct {
	ws   WhitespaceMode
	name string
}{
	{RemoveWhitespace, "RemoveWhitespace"},
	{AssumeCompact, "AssumeCompact"},
	{PreserveWhitespace, "PreserveWhitespace"},
}

// TestStripDefaultsInPlaceMatchesFresh pins the in-place result against the
// fresh-buffer one on the shapes that reach the keep-listed-container path, and
// pins the fresh result itself against a written-out expectation so the fix
// cannot quietly redefine the semantics it was preserving. want is the compact
// (RemoveWhitespace) output; the other two modes are checked against the fresh
// run only, since AssumeCompact misreads spaced input by contract.
//
// Before the fix every non-AssumeCompact case whose input carries a left shift
// (something dropped, or whitespace removed, ahead of the kept member) produced a
// different in-place result — the member vanished, or came out garbled — while
// the fresh result was already right.
func TestStripDefaultsInPlaceMatchesFresh(t *testing.T) {
	d0 := [][]byte{[]byte("0"), []byte("")}
	kb := [][]byte{[]byte("b"), []byte("keepMe"), []byte(`esc\"key`)}

	cases := []struct{ name, in, want string }{
		// The reported shapes: a kept container after a member whose whitespace
		// (or whole self) is dropped, so the write cursor lags the read cursor.
		{"lag from removed whitespace", `{"q": 1,"b":{"x":0},"t":2}`, `{"q":1,"b":{"x":0},"t":2}`},
		{"lag from leading whitespace", `{ "b":{"x":0}}`, `{"b":{"x":0}}`},
		{"lag from dropped member", `{"":0,"b":{"x":{"y":0}}}`, `{"b":{"x":{"y":0}}}`},
		{"large lag", `{"d1":0,"d2":0,"d3":0,"d4":0,"d5":0,"d6":0,"b":{"x":0,"y":0}}`, `{"b":{"x":0,"y":0}}`},

		// No lag at all: the kept member is first and nothing ahead of it is
		// dropped, so every write is a self-copy and in-place was already right.
		{"zero lag, only member", `{"b":{"x":0}}`, `{"b":{"x":0}}`},
		{"zero lag, kept first", `{"b":{"x":0},"a":1}`, `{"b":{"x":0},"a":1}`},
		{"kept last", `{"a":1,"b":{"x":0}}`, `{"a":1,"b":{"x":0}}`},
		{"kept array value", `{"a":0,"b":[0,0]}`, `{"b":[0,0]}`},

		// A kept member whose value does NOT strip away keeps the stripped value,
		// not the original — the semantics the snapshot must not disturb.
		{"kept, value survives", `{"a":0,"b":{"x":0,"y":9}}`, `{"b":{"y":9}}`},
		{"kept, value partly survives", `{"z":0,"b":[0,5,0]}`, `{"b":[5]}`},

		// Nested kept containers, which is what makes one shared snapshot buffer
		// enough: an inner kept member replaces the outer member's snapshot, but
		// it is itself always emitted — verbatim in the first case, stripped in
		// the second — which leaves every container above it non-empty, so the
		// outer member never reaches back for the snapshot it lost.
		{"nested kept containers", `{"a":0,"b":{"keepMe":{"x":0}}}`, `{"b":{"keepMe":{"x":0}}}`},
		{"inner emits from its snapshot", `{"pad":0,"b":{"q":0,"keepMe":{"x":0}},"z":0}`, `{"b":{"keepMe":{"x":0}}}`},
		{"inner emits a stripped value", `{"p":0,"b":{"keepMe":{"x":0,"y":1}}}`, `{"b":{"keepMe":{"y":1}}}`},
		{"kept inside dropped-looking parent", `{"a":0,"b":{"q":0,"keepMe":[0]}}`, `{"b":{"keepMe":[0]}}`},
		{"kept container in array element", `{"a":0,"z":[{"b":{"x":0}}]}`, `{"z":[{"b":{"x":0}}]}`},

		// Escapes in keys and values, including a kept key that is itself escaped.
		{"escaped key kept", `{"a":0,"esc\"key":{"x":0}}`, `{"esc\"key":{"x":0}}`},
		{"escaped value in kept container", `{"a":0,"b":{"x":0,"e":"q\"\\z"}}`, `{"b":{"e":"q\"\\z"}}`},
		{"escaped value strips to empty", `{"a":0,"b":{"e":""}}`, `{"b":{"e":""}}`},

		// Several kept members in one object: the scratch is reused across them.
		{"two kept members", `{"a":0,"b":{"x":0},"c":0,"keepMe":{"y":0}}`, `{"b":{"x":0},"keepMe":{"y":0}}`},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := stripFresh(c.in, d0, kb, RemoveWhitespace); got != c.want {
				t.Errorf("fresh RemoveWhitespace = %q, want %q", got, c.want)
			}
			for _, m := range stripWSModes {
				fresh := stripFresh(c.in, d0, kb, m.ws)
				if got := stripInPlace(c.in, d0, kb, m.ws); got != fresh {
					t.Errorf("%s: in-place = %q, fresh = %q", m.name, got, fresh)
				}
				if fresh != "" && m.ws != AssumeCompact && !json.Valid([]byte(fresh)) {
					t.Errorf("%s: output %q is not valid JSON", m.name, fresh)
				}
			}
		})
	}
}

// stripFuzzKeys / stripFuzzScalars are the alphabets randomStripDoc draws from:
// short so keys repeat and collide with the keep list, and carrying escapes (in
// both keys and values) so the escape-aware scans are exercised. The key bodies
// are written as they appear between the quotes, which is also how keep entries
// are compared, so `esc\"key` here is the keep entry []byte(`esc\"key`).
var stripFuzzKeys = []string{"a", "b", "", "keepMe", `esc\"key`, "x", "longer_key_name", `sl\/ash`}

// (Every entry has to be valid to encoding/json as well as to this package, since
// the fuzz uses json.Valid as its oracle for the output: `00`, which Valid here
// accepts and encoding/json rejects, would fail that check on the input's own
// bytes rather than on anything the stripper did.)
var stripFuzzScalars = []string{
	`0`, `1`, `-2.5e3`, `true`, `false`, `null`,
	`""`, `"x"`, `"none"`, `"a\"b"`, `"é tail"`, `"0"`,
}

// randomStripDoc builds a random well-formed JSON document from those alphabets.
// Duplicate keys and empty containers are deliberately possible: both are shapes
// the stripper has its own handling for.
func randomStripDoc(r *rand.Rand, depth int) string {
	if depth <= 0 || r.Intn(3) == 0 {
		return stripFuzzScalars[r.Intn(len(stripFuzzScalars))]
	}
	n := r.Intn(5)
	parts := make([]string, n)
	if r.Intn(2) == 0 {
		for i := range parts {
			parts[i] = `"` + stripFuzzKeys[r.Intn(len(stripFuzzKeys))] + `":` + randomStripDoc(r, depth-1)
		}
		return "{" + strings.Join(parts, ",") + "}"
	}
	for i := range parts {
		parts[i] = randomStripDoc(r, depth-1)
	}
	return "[" + strings.Join(parts, ",") + "]"
}

// randomSubset returns a random non-empty-ish subset of items as byte slices.
func randomSubset(r *rand.Rand, items []string) [][]byte {
	var out [][]byte
	for _, it := range items {
		if r.Intn(3) == 0 {
			out = append(out, []byte(it))
		}
	}
	return out
}

// TestStripDefaultsInPlaceFuzz is the post-condition of the in-place fix as a
// property: for every document, every defaults/keep pool and every whitespace
// mode, stripping in place produces exactly the bytes stripping into a separate
// buffer does. It also checks the fresh run leaves its input alone and that a
// non-empty result is valid JSON whenever the mode's contract was honored — so a
// fix that made the two agree by breaking both would still fail here.
//
// Against the unfixed code this fails within the first few dozen documents:
// in-place loses or garbles a keep-listed member whose container value strips to
// empty, in all three whitespace modes.
func TestStripDefaultsInPlaceFuzz(t *testing.T) {
	r := rand.New(rand.NewSource(20260809))
	defaultPool := []string{"0", "", "false", "null", "x", "none", `a\"b`, "1", "00"}

	const iterations = 20000
	for n := 0; n < iterations; n++ {
		compact := randomStripDoc(r, 1+r.Intn(4))
		defaults := randomSubset(r, defaultPool)
		keep := randomSubset(r, stripFuzzKeys)

		// Both a compact and an indented spelling of the same document, so the
		// whitespace-skipping and whitespace-preserving paths both see real runs.
		docs := []struct {
			in       string
			isCompac bool
		}{{compact, true}}
		var pretty bytes.Buffer
		if err := json.Indent(&pretty, []byte(compact), strings.Repeat(" ", r.Intn(3)), "\t"); err == nil {
			docs = append(docs, struct {
				in       string
				isCompac bool
			}{pretty.String(), false})
		}

		for _, d := range docs {
			for _, m := range stripWSModes {
				orig := d.in
				fresh := stripFresh(d.in, defaults, keep, m.ws)
				if d.in != orig {
					t.Fatalf("fresh run mutated its input: %q -> %q", orig, d.in)
				}
				if got := stripInPlace(d.in, defaults, keep, m.ws); got != fresh {
					t.Fatalf("in-place != fresh\n  doc      = %q\n  defaults = %q\n  keep     = %q\n  mode     = %s\n  in-place = %q\n  fresh    = %q",
						d.in, defaults, keep, m.name, got, fresh)
				}
				// AssumeCompact is only answerable for compact input, by contract.
				if (m.ws != AssumeCompact || d.isCompac) && fresh != "" && !json.Valid([]byte(fresh)) {
					t.Fatalf("%s output is not valid JSON\n  doc      = %q\n  defaults = %q\n  keep     = %q\n  out      = %q",
						m.name, d.in, defaults, keep, fresh)
				}
			}
		}
	}
}

// TestStripDefaultsSnapshotShortOfMember exercises finishEarly's last branch, the
// one where the snapshot of a kept member stops short of where the walk ended.
// The snapshot is sized by SkipValue over the same bytes, and the two agree on
// every well-formed value; they part on a bare token holding a quote (`q"w`),
// where SkipValue opens a string the walk does not and so ends up *outside* a
// string where the walk is *inside* one — it then takes the `}` inside the key
// "a}b" for the container's close and stops early.
//
// What is locked here is that the branch is a fallback, not a cliff: the walk
// must not read past the snapshot, the fresh result is exactly what it was before
// the snapshot existed (the member re-emitted from the input), and the in-place
// run — the one case where the input is no longer intact to re-emit from — still
// returns rather than panicking.
func TestStripDefaultsSnapshotShortOfMember(t *testing.T) {
	keep := [][]byte{[]byte("b")}
	defaults := [][]byte{[]byte(`q"w`), []byte("0")}
	cases := []struct{ in, want string }{
		{`{"b":{"k":q"w,"a}b":0}}`, `{"b":{"k":q"w,"a}b":0}}`},
		{`{"z":0,"b":{"k":q"w,"a}b":0}}`, `{"b":{"k":q"w,"a}b":0}}`},
		{`{"b":{"k":q"w,"a}b":0},"t":1}`, `{"b":{"k":q"w,"a}b":0},"t":1}`},
	}
	for _, c := range cases {
		if got := stripFresh(c.in, defaults, keep, RemoveWhitespace); got != c.want {
			t.Errorf("fresh StripDefaults(%q) = %q, want %q", c.in, got, c.want)
		}
		stripInPlace(c.in, defaults, keep, RemoveWhitespace) // must not panic
	}
}

// TestStripDefaultsInPlaceFuzzMalformed is the same property over deliberately
// broken input, where StripDefaults is only promised to be best effort: the point
// is that in-place stays as safe as the fresh path — no panic, no read past the
// buffer — not that the two agree, since a truncated container can leave the walk
// and the snapshot's own SkipValue disagreeing about where the value ended.
func TestStripDefaultsInPlaceFuzzMalformed(t *testing.T) {
	r := rand.New(rand.NewSource(4242))
	defaults := [][]byte{[]byte("0"), []byte("")}
	keep := [][]byte{[]byte("b"), []byte("keepMe"), []byte("a")}

	for n := 0; n < 20000; n++ {
		doc := randomStripDoc(r, 1+r.Intn(3))
		b := []byte(doc)
		switch r.Intn(3) {
		case 0: // truncate
			b = b[:r.Intn(len(b)+1)]
		case 1: // corrupt one byte
			if len(b) > 0 {
				b[r.Intn(len(b))] = []byte(`{}[]",:\ `)[r.Intn(9)]
			}
		default: // splice in a stray byte
			at := r.Intn(len(b) + 1)
			b = append(b[:at:at], append([]byte{[]byte(`{}[]",:\`)[r.Intn(8)]}, b[at:]...)...)
		}
		for _, m := range stripWSModes {
			_ = stripFresh(string(b), defaults, keep, m.ws)
			_ = stripInPlace(string(b), defaults, keep, m.ws)
		}
	}
}
