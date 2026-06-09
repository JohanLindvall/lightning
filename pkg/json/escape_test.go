package json

import (
	"math/rand"
	"strings"
	"testing"
)

// escapeReference is a deliberately simple byte-by-byte escaper used to
// validate the SWAR-optimized EscapeStringInto.
func escapeReference(s []byte) string {
	var b strings.Builder
	for _, c := range s {
		switch {
		case c == '"':
			b.WriteString(`\"`)
		case c == '\\':
			b.WriteString(`\\`)
		case c == '\t':
			b.WriteString(`\t`)
		case c == '\r':
			b.WriteString(`\r`)
		case c == '\n':
			b.WriteString(`\n`)
		case c < 0x20:
			b.WriteString(`\u00`)
			b.WriteByte(hexAlphabet[c>>4])
			b.WriteByte(hexAlphabet[c&0xf])
		default:
			b.WriteByte(c)
		}
	}
	return b.String()
}

// TestEscapeStringIntoReference checks EscapeStringInto against escapeReference
// over every byte value (alone and embedded in clean runs that straddle the
// 8-byte SWAR boundary) and a deterministic fuzz of byte slices biased toward
// bytes that need escaping — exercising cross-lane borrow in the SWAR check.
func TestEscapeStringIntoReference(t *testing.T) {
	check := func(in []byte) {
		t.Helper()
		got := string(EscapeStringInto(in, nil))
		if want := escapeReference(in); got != want {
			t.Fatalf("EscapeStringInto(%q) = %q, want %q", in, got, want)
		}
	}

	for v := 0; v < 256; v++ {
		check([]byte{byte(v)})
		// Embed at offsets that cross the 8-byte word boundary.
		buf := append([]byte("clean123"), byte(v))
		buf = append(buf, "clean456"...)
		check(buf)
	}

	r := rand.New(rand.NewSource(1))
	for n := 0; n < 5000; n++ {
		b := make([]byte, r.Intn(40))
		for i := range b {
			switch r.Intn(4) {
			case 0:
				b[i] = byte(r.Intn(0x20)) // control byte
			case 1:
				b[i] = []byte{'"', '\\'}[r.Intn(2)]
			default:
				b[i] = byte(0x20 + r.Intn(0xe0)) // printable / high byte
			}
		}
		check(b)
	}
}

// altEscaping names the unescapeCases whose canonical JSON body (in) uses an
// escape the escaper does not emit, so EscapeString's output differs from in:
// the escaper leaves '/' unescaped and renders other control bytes as \u00XX
// rather than \b/\f. For these, only the round-trip is checked.
var altEscaping = map[string]bool{"slash": true, "backspace formfeed": true}

func Test_unit_EscapeString(t *testing.T) {
	// Escape the decoded form of every valid UnescapeString case and check both
	// the exact output and the round-trip, exercising the escaper over the same
	// inputs as TestUnescapeString. (The wantErr cases are malformed escape
	// input, which does not apply to escaping.)
	for _, tt := range unescapeCases {
		if tt.wantErr {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			decoded := []byte(tt.want)

			var sb strings.Builder
			EscapeString(decoded, &sb)
			escaped := sb.String()

			// Exact output: for the ordinary cases EscapeString reproduces the
			// canonical body in.
			if !altEscaping[tt.name] && escaped != tt.in {
				t.Fatalf("EscapeString(%q) = %q, want %q", tt.want, escaped, tt.in)
			}

			// EscapeStringInto must agree with EscapeString.
			if into := string(EscapeStringInto(decoded, nil)); into != escaped {
				t.Fatalf("EscapeStringInto = %q, EscapeString = %q", into, escaped)
			}

			// Round-trip: unescaping the escaped form reproduces the original.
			got, err := UnescapeString([]byte(escaped))
			if err != nil {
				t.Fatalf("UnescapeString(%q): %v", escaped, err)
			}
			if got != tt.want {
				t.Fatalf("round-trip got %q, want %q (escaped to %q)", got, tt.want, escaped)
			}
		})
	}
}

// Realistic corpus of raw (unescaped) string values to escape, covering the
// no-escape fast path and a range of escape densities (quotes, backslashes,
// newlines, and control bytes that expand to the 6-byte \u00XX form).
var escapeBenchCases = []struct {
	name string
	in   []byte
}{
	{"short_clean", []byte("user@example.com")},
	{"sentence_clean", []byte("The quick brown fox jumps over the lazy dog.")},
	{"url_clean", []byte("https://example.com/path/to/resource?q=search&page=2")},
	{"log_line_clean", []byte(strings.Repeat("2026-06-09T12:00:00Z INFO request handled ", 8))},
	{"path_with_backslash", []byte(`C:\Users\jl\src\lightning\pkg\support`)},
	{"json_in_json", []byte(`{"id":42,"name":"widget","tags":["a","b"]}`)},
	{"prose_with_quotes", []byte("She said \"hello\" and then\nwalked away.")},
	{"control_bytes", []byte("line1\nline2\tcol\r\x00\x01\x02\x1f end")},
	{"mostly_clean_one_quote", []byte(strings.Repeat("plain text content ", 16) + `"`)},
}

// BenchmarkEscapeString measures escaping onto a reused strings.Builder, the
// common "write a JSON string field" path. EscapeString allocates a scratch
// slice per call, so this is not allocation-free; see BenchmarkEscapeStringInto
// for the buffer-reuse form.
func BenchmarkEscapeString(b *testing.B) {
	for _, bc := range escapeBenchCases {
		b.Run(bc.name, func(b *testing.B) {
			var sb strings.Builder
			b.SetBytes(int64(len(bc.in)))
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				sb.Reset()
				EscapeString(bc.in, &sb)
			}
			_ = sb.String()
		})
	}
}

// BenchmarkEscapeStringInto measures the allocation-free form: escaping into a
// reusable out buffer. The buffer is sized once for the worst case (every byte
// expanding to the 6-byte \u00XX form) so append never reallocates.
func BenchmarkEscapeStringInto(b *testing.B) {
	for _, bc := range escapeBenchCases {
		b.Run(bc.name, func(b *testing.B) {
			out := make([]byte, 0, 6*len(bc.in)+8)
			b.SetBytes(int64(len(bc.in)))
			b.ReportAllocs()
			var sink []byte
			for i := 0; i < b.N; i++ {
				sink = EscapeStringInto(bc.in, out[:0])
			}
			_ = sink
		})
	}
}
