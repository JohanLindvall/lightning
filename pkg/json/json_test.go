package json

import (
	stdjson "encoding/json"
	"strings"
	"testing"
)

// unescapeCases is the shared corpus of UnescapeString inputs: in is a JSON
// string body (escapes intact), want is its decoded form, and wantErr marks
// malformed input. Test_unit_EscapeString reuses the valid cases to check the
// escaper and to round-trip escape against unescape over these inputs.
var unescapeCases = []struct {
	name    string
	in      string
	want    string
	wantErr bool
}{
	{"empty", "", "", false},
	{"plain", "hello world", "hello world", false},
	{"quote", `say \"hi\"`, `say "hi"`, false},
	{"backslash", `a\\b`, `a\b`, false},
	{"slash", `a\/b`, "a/b", false},
	{"controls", `tab\tnl\ncr\r`, "tab\tnl\ncr\r", false},
	{"backspace formfeed", `\b\f`, "\b\f", false},
	{"unicode", `Aé`, "Aé", false},
	{"surrogate pair", `😀`, "\U0001F600", false},
	{"escape at end", `abc\n`, "abc\n", false},
	{"escape at start", `\nabc`, "\nabc", false},
	{"escaped quote at start", `\"xyz`, `"xyz`, false},
	{"escaped backslash", `\\n`, `\n`, false},
	{"long unescaped", strings.Repeat("x", 256), strings.Repeat("x", 256), false},
	{"long with one escape", strings.Repeat("x", 256) + `\n`, strings.Repeat("x", 256) + "\n", false},

	{"bad escape", `a\xb`, "", true},
	{"truncated escape", `abc\`, "", true},
	{"bad unicode", `\u00zz`, "", true},
	{"truncated unicode", `\u00`, "", true},
}

func TestUnescapeString(t *testing.T) {
	for _, tt := range unescapeCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnescapeString([]byte(tt.in))
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr = %v", err, tt.wantErr)
			}
			if err == nil && got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestUnescapeStringInto(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    string
		wantErr bool
	}{
		{"empty", "", "", false},
		{"plain", "hello world", "hello world", false},
		{"quote", `say \"hi\"`, `say "hi"`, false},
		{"backslash", `a\\b`, `a\b`, false},
		{"controls", `tab\tnl\ncr\r`, "tab\tnl\ncr\r", false},
		{"unicode", `Aé`, "Aé", false},
		{"surrogate pair", `😀`, "\U0001F600", false},
		{"dense escapes", `\"a\"b\"c\"`, `"a"b"c"`, false},
		{"escape at start", `\nabc`, "\nabc", false},
		{"long with one escape", strings.Repeat("x", 256) + `\n`, strings.Repeat("x", 256) + "\n", false},

		{"bad escape", `a\xb`, "", true},
		{"truncated escape", `abc\`, "", true},
		{"bad unicode", `\u00zz`, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := []byte(tt.in)                     // copy: decoding in place mutates it
			got, err := UnescapeStringInto(buf, buf) // out == in: true in place
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr = %v", err, tt.wantErr)
			}
			if err == nil && got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}

// TestUnescapeStringIntoMatchesUnescapeString runs both decoders over the
// same inputs and requires identical results.
func TestUnescapeStringIntoMatchesUnescapeString(t *testing.T) {
	inputs := []string{
		"",
		"no escapes here",
		`a\"b\\c\/d`,
		`tabs\tand\nlines`,
		`☃ snowman 😀 emoji`,
		`\"` + strings.Repeat(`x\"`, 64),
		strings.Repeat("clean ", 100) + `\t` + strings.Repeat(" more", 50),
	}
	for _, s := range inputs {
		want, err := UnescapeString([]byte(s))
		if err != nil {
			t.Fatalf("UnescapeString(%q): %v", s, err)
		}
		buf := []byte(s)
		got, err := UnescapeStringInto(buf, buf)
		if err != nil {
			t.Fatalf("UnescapeStringInto(%q): %v", s, err)
		}
		if got != want {
			t.Fatalf("input %q: in-place got %q, want %q", s, got, want)
		}
	}
}

// TestUnescapeStringMatchesStdlib cross-checks the decoder against
// encoding/json for a range of inputs by wrapping each body in quotes.
func TestUnescapeStringMatchesStdlib(t *testing.T) {
	bodies := []string{
		"",
		"plain ascii text",
		`with a \" quote and \\ slash`,
		`tabs\tand\nnewlines`,
		`unicode ☃ snowman and 😀 emoji`,
		strings.Repeat(`word `, 50) + `\tend`,
	}
	for _, body := range bodies {
		var want string
		if err := stdjson.Unmarshal([]byte(`"`+body+`"`), &want); err != nil {
			t.Fatalf("stdlib rejected %q: %v", body, err)
		}
		got, err := UnescapeString([]byte(body))
		if err != nil {
			t.Fatalf("UnescapeString(%q) error: %v", body, err)
		}
		if got != want {
			t.Fatalf("body %q: got %q, want %q", body, got, want)
		}
	}
}

// Realistic benchmark corpus: string values of the kind that actually show up
// in JSON payloads, covering the no-escape fast path (which aliases the input)
// and the allocating escape path at a few representative escape densities.
var benchCases = []struct {
	name string
	in   []byte
}{
	{"short_clean", []byte("user@example.com")},
	{"sentence_clean", []byte("The quick brown fox jumps over the lazy dog.")},
	{"url_clean", []byte("https://example.com/path/to/resource?q=search&page=2")},
	{"log_line_clean", []byte(strings.Repeat("2026-06-09T12:00:00Z INFO request handled ", 8))},
	{"path_escaped", []byte(`C:\\Users\\jl\\src\\lightning\\pkg\\support`)},
	{"json_in_json", []byte(`{\"id\":42,\"name\":\"widget\",\"tags\":[\"a\",\"b\"]}`)},
	{"prose_with_quotes", []byte(`She said \"hello\" and then\nwalked away.`)},
	{"unicode_heavy", []byte(`café naïve ☃ résumé 😀`)},
	{"mostly_clean_one_escape", []byte(strings.Repeat("plain text content ", 16) + `\n`)},
}

func BenchmarkUnescapeString(b *testing.B) {
	for _, bc := range benchCases {
		b.Run(bc.name, func(b *testing.B) {
			b.SetBytes(int64(len(bc.in)))
			b.ReportAllocs()
			var sink string
			for i := 0; i < b.N; i++ {
				s, err := UnescapeString(bc.in)
				if err != nil {
					b.Fatal(err)
				}
				sink = s
			}
			_ = sink
		})
	}
}

// BenchmarkUnescapeStringInto measures the allocation-free decoder writing into
// a reusable out buffer (sized once with cap >= len(in)). in is left unchanged,
// so no per-iteration restore is needed.
func BenchmarkUnescapeStringInto(b *testing.B) {
	for _, bc := range benchCases {
		b.Run(bc.name, func(b *testing.B) {
			out := make([]byte, len(bc.in))
			b.SetBytes(int64(len(bc.in)))
			b.ReportAllocs()
			var sink string
			for i := 0; i < b.N; i++ {
				s, err := UnescapeStringInto(bc.in, out)
				if err != nil {
					b.Fatal(err)
				}
				sink = s
			}
			_ = sink
		})
	}
}
