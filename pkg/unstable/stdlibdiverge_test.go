package unstable

import (
	ejson "encoding/json"
	"fmt"
	"testing"
	"time"
	"unicode/utf8"
)

// This file pins two places where these readers deliberately accept more than
// encoding/json does, in the spirit of pkg/json's TestValidDivergesFromStdlib
// and of skipdiverge_test.go: a documented disagreement is only documented for
// as long as a test holds both sides of it. Each test asserts the stdlib premise
// too, so a change on either side surfaces here rather than silently.
//
// A failure is not automatically a bug. It means the relationship moved: either
// this library changed (decide whether the new behaviour is wanted and update
// the case and the doc comment it cites) or the stdlib did (the premise
// assertions are the ones that fire).
//
// Note both tests reach encoding/json through types it owns — time.Time,
// string, map[string]any — so there is no risk of the differential-test trap
// this repo has hit before, where the target type has a lightning-generated
// UnmarshalJSON and encoding/json merely calls back into the code under test.

// bslash is a single backslash, spelled without a Go escape so that the JSON
// escape sequences built from it below read as what they are.
const bslash = `\`

// escU renders r as the JSON escape \uXXXX.
func escU(r rune) string { return bslash + fmt.Sprintf("u%04X", r) }

// TestReadTimeAcceptsEscapedTimestamps pins the one axis on which ReadTimeOrNull
// is more lenient than encoding/json's time.Time: it decodes the JSON string and
// parses the *value*, where time.Time.UnmarshalJSON parses the raw bytes between
// the quotes without unescaping them (go.dev/issue/47353). So a timestamp
// carrying any \uXXXX escape is legal JSON denoting a legal instant that the
// stdlib rejects and this accepts.
//
// Both sides are asserted: lightning must accept and produce the instant the
// escape denotes, and the stdlib must reject. Closing the gap in either
// direction — this package parsing raw bytes, or the stdlib learning to unescape
// — fails this test, which is the point.
func TestReadTimeAcceptsEscapedTimestamps(t *testing.T) {
	want, err := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	if err != nil {
		t.Fatalf("premise broken: time.Parse rejects the plain form: %v", err)
	}
	wantOffset, err := time.Parse(time.RFC3339, "2021-01-01T02:00:00+02:00")
	if err != nil {
		t.Fatalf("premise broken: time.Parse rejects the plain offset form: %v", err)
	}

	for _, tc := range []struct {
		name string
		doc  string
		want time.Time
	}{{
		name: "zone Z escaped",
		doc:  `"2021-01-01T00:00:00` + escU('Z') + `"`,
		want: want,
	}, {
		name: "digit escaped",
		doc:  `"` + escU('2') + `021-01-01T00:00:00Z"`,
		want: want,
	}, {
		name: "separator T escaped",
		doc:  `"2021-01-01` + escU('T') + `00:00:00Z"`,
		want: want,
	}, {
		name: "colon in the numeric offset escaped",
		doc:  `"2021-01-01T02:00:00+02` + escU(':') + `00"`,
		want: wantOffset,
	}} {
		t.Run(tc.name, func(t *testing.T) {
			data := []byte(tc.doc)

			got, end, err := ReadTimeOrNull(data, 0)
			if err != nil {
				t.Fatalf("ReadTimeOrNull(%s) = %v; want the instant the escape denotes", tc.doc, err)
			}
			if !got.Equal(tc.want) {
				t.Errorf("ReadTimeOrNull(%s) = %v, want %v", tc.doc, got, tc.want)
			}
			if end != len(data) {
				t.Errorf("ReadTimeOrNull(%s) end = %d, want %d", tc.doc, end, len(data))
			}
			// The lax reader unescapes through the same string reader, so it
			// inherits the leniency; its doc says so.
			if _, _, err := ReadTimeLaxOrNull(data, 0); err != nil {
				t.Errorf("ReadTimeLaxOrNull(%s) = %v; want it to inherit the leniency", tc.doc, err)
			}

			// The premise depends on the toolchain. Through Go 1.26 the stdlib's
			// time.Time refused these — not because the instant is wrong but
			// because it parsed the raw quoted bytes without unescaping them
			// (go.dev/issue/47353); Go 1.27's json/v2-backed encoding/json
			// unescapes first, closing the divergence. Either way the document is
			// well-formed JSON, and wherever the stdlib does decode it, it must
			// land on the same instant this reader does.
			if !ejson.Valid(data) {
				t.Fatalf("test bug: %s is not valid JSON", tc.doc)
			}
			var std time.Time
			if err := ejson.Unmarshal(data, &std); err == nil {
				if !std.Equal(tc.want) {
					t.Fatalf("encoding/json decodes %s as %v, want %v: the two now disagree on the instant itself", tc.doc, std, tc.want)
				}
				t.Logf("encoding/json also decodes %s: the escape divergence is closed on this toolchain", tc.doc)
			}
		})
	}

	// The control, and the other half of the doc claim: with no escape present
	// the two agree, so the divergence above really is about unescaping and not
	// about the date grammar. Run over the whole generated date corpus, since
	// "agrees on the escape-free set" is the claim ReadTimeOrNull's comment
	// makes. If this half fails while the escaped cases above still pass, the
	// likely cause is outside this repo: encoding/json's time.Time currently
	// reduces to time.Parse(time.RFC3339, ...) because its extra RFC 3339
	// strictness is compiled out pending go.dev/issue/54580, and re-enabling
	// that would make the stdlib the stricter of the two.
	for _, s := range dateCorpus() {
		doc := []byte(`"` + s + `"`)
		got, _, err := ReadTimeOrNull(doc, 0)
		var std time.Time
		serr := ejson.Unmarshal(doc, &std)
		if (err == nil) != (serr == nil) {
			t.Fatalf("escape-free %q: lightning err = %v, encoding/json err = %v; the two must agree when no escape is involved", s, err, serr)
		}
		if err == nil && !got.Equal(std) {
			t.Fatalf("escape-free %q: lightning = %v, stdlib = %v", s, got, std)
		}
	}
}

// TestStringsPassInvalidUTF8Through pins the other deliberate leniency: string
// contents are handed back byte for byte, so raw invalid UTF-8 survives the
// decode, where encoding/json coerces every invalid byte to U+FFFD. README calls
// this out ("invalid UTF-8 is passed through") and nothing tested it, so the
// property could have drifted — in either direction — unnoticed. This is
// DECODE-only: the escape direction (pkg/json's EscapeString/EscapeStringInto)
// coerces ill-formed bytes to U+FFFD like the stdlib, pinned by pkg/json's
// TestEscapeStringMatchesStdlibCoercion — the two directions differ on purpose
// (decoding is lossless, produced JSON must be valid UTF-8).
//
// Every path that could differ is covered: the copying reader, the two aliasing
// readers (nocopy and destructive), the key reader, the dynamic any decoder
// (value and map key), decodeEscaped's literal runs (reached by putting an
// escape in the same string), and UnescapeString.
func TestStringsPassInvalidUTF8Through(t *testing.T) {
	for _, tc := range []struct {
		name string
		body string
	}{
		{"lone 0xff", "a\xffb"},
		{"lone continuation byte", "a\x80b"},
		{"truncated three-byte sequence", "a\xe4\xb8b"},
		{"invalid byte at the end", "ab\xff"},
		{"surrogate half encoded raw", "a\xed\xa0\x80b"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if utf8.ValidString(tc.body) {
				t.Fatalf("test bug: %q is valid UTF-8, so it tests nothing", tc.body)
			}
			doc := `"` + tc.body + `"`

			check := func(what, got string, err error) {
				t.Helper()
				if err != nil {
					t.Fatalf("%s(%q): %v", what, doc, err)
				}
				if got != tc.body {
					t.Errorf("%s(%q) = %q, want the bytes verbatim (%q)", what, doc, got, tc.body)
				}
			}

			s, _, err := ReadStringOrNull([]byte(doc), 0)
			check("ReadStringOrNull", s, err)
			s, _, err = ReadStringNoCopyOrNull([]byte(doc), 0)
			check("ReadStringNoCopyOrNull", s, err)
			s, _, err = ReadStringDestructiveOrNull([]byte(doc), 0)
			check("ReadStringDestructiveOrNull", s, err)
			s, _, err = ReadKey([]byte(doc), 0)
			check("ReadKey", s, err)

			v, _, err := DecodeValue([]byte(doc), 0)
			if err != nil {
				t.Fatalf("DecodeValue(%q): %v", doc, err)
			}
			if got, ok := v.(string); !ok || got != tc.body {
				t.Errorf("DecodeValue(%q) = %#v, want the bytes verbatim (%q)", doc, v, tc.body)
			}

			// The object form exercises decodeAnyObject's inline key read as well
			// as the value read.
			obj := `{` + doc + `:` + doc + `}`
			v, _, err = DecodeValue([]byte(obj), 0)
			if err != nil {
				t.Fatalf("DecodeValue(%q): %v", obj, err)
			}
			m, ok := v.(map[string]any)
			if !ok {
				t.Fatalf("DecodeValue(%q) = %#v, want a map", obj, v)
			}
			if got, ok := m[tc.body]; !ok {
				t.Errorf("DecodeValue(%q) keys = %#v, want the key bytes verbatim (%q)", obj, m, tc.body)
			} else if got != any(tc.body) {
				t.Errorf("DecodeValue(%q)[%q] = %#v, want the bytes verbatim", obj, tc.body, got)
			}

			// An escape in the same string routes through decodeEscaped, whose
			// literal runs are bulk-copied — the copy that would have to change
			// for this library to start coercing.
			escBody := tc.body + bslash + "n" + tc.body
			wantEsc := tc.body + "\n" + tc.body
			escDoc := `"` + escBody + `"`
			s, _, err = ReadStringOrNull([]byte(escDoc), 0)
			if err != nil {
				t.Fatalf("ReadStringOrNull(%q): %v", escDoc, err)
			}
			if s != wantEsc {
				t.Errorf("ReadStringOrNull(%q) = %q, want %q (escape decoded, literal runs verbatim)", escDoc, s, wantEsc)
			}
			s, _, err = ReadStringDestructiveOrNull([]byte(escDoc), 0)
			if err != nil {
				t.Fatalf("ReadStringDestructiveOrNull(%q): %v", escDoc, err)
			}
			if s != wantEsc {
				t.Errorf("ReadStringDestructiveOrNull(%q) = %q, want %q", escDoc, s, wantEsc)
			}

			s, err = UnescapeString([]byte(tc.body))
			check("UnescapeString", s, err)

			// The premise, in three parts: encoding/json considers the document
			// well-formed, decodes it without error, and returns something other
			// than the input bytes — a coerced, valid-UTF-8 string.
			if !ejson.Valid([]byte(doc)) {
				t.Fatalf("premise broken: encoding/json.Valid rejects %q, so the divergence is no longer at the decode", doc)
			}
			var std string
			if err := ejson.Unmarshal([]byte(doc), &std); err != nil {
				t.Fatalf("premise broken: encoding/json rejects %q: %v", doc, err)
			}
			if std == tc.body {
				t.Fatalf("premise broken: encoding/json no longer coerces %q — the divergence has closed", tc.body)
			}
			if !utf8.ValidString(std) {
				t.Fatalf("premise broken: encoding/json returned invalid UTF-8 %q", std)
			}
			var stdMap map[string]any
			if err := ejson.Unmarshal([]byte(obj), &stdMap); err != nil {
				t.Fatalf("premise broken: encoding/json rejects %q: %v", obj, err)
			}
			if _, ok := stdMap[tc.body]; ok {
				t.Fatalf("premise broken: encoding/json no longer coerces object keys")
			}
		})
	}

	// The documented example, pinned exactly rather than by property: "a\xffb"
	// stays "a\xffb" here and becomes "a�b" in encoding/json. Both doc
	// comments (unstable.ReadStringOrNull, json.UnescapeString) quote this pair.
	const body = "a\xffb"
	got, _, err := ReadStringOrNull([]byte(`"`+body+`"`), 0)
	if err != nil {
		t.Fatalf("ReadStringOrNull: %v", err)
	}
	if got != body {
		t.Errorf("ReadStringOrNull = %q, want %q", got, body)
	}
	var std string
	if err := ejson.Unmarshal([]byte(`"`+body+`"`), &std); err != nil {
		t.Fatalf("premise broken: %v", err)
	}
	if want := "a�b"; std != want {
		t.Errorf("premise broken: encoding/json = %q, want %q", std, want)
	}
}
