package json

import (
	ejson "encoding/json"
	"math/rand"
	"strings"
	"testing"
	"unicode/utf8"
)

// escapeReference is a deliberately simple escaper used to validate the
// SWAR/SIMD-optimized EscapeStringInto: byte-by-byte for ASCII, with non-ASCII
// bytes walked rune-by-rune so every byte that is not part of a well-formed
// UTF-8 sequence becomes U+FFFD (utf8.DecodeRune's r == RuneError && size == 1
// convention, which is also encoding/json's) and valid sequences pass through
// verbatim.
func escapeReference(s []byte) string {
	var b strings.Builder
	for i := 0; i < len(s); {
		c := s[i]
		if c >= utf8.RuneSelf {
			r, size := utf8.DecodeRune(s[i:])
			if r == utf8.RuneError && size == 1 {
				b.WriteString("�")
			} else {
				b.Write(s[i : i+size])
			}
			i += size
			continue
		}
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
		i++
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
		want := escapeReference(in)
		if got := string(EscapeStringInto(in, nil)); got != want {
			t.Fatalf("EscapeStringInto(%q) = %q, want %q", in, got, want)
		}
		// EscapeString has its own dispatch (clean-prefix probe, UTF-8 decision);
		// hold it to the same reference over the same corpus.
		var sb strings.Builder
		EscapeString(in, &sb)
		if got := sb.String(); got != want {
			t.Fatalf("EscapeString(%q) = %q, want %q", in, got, want)
		}
	}

	for v := 0; v < 256; v++ {
		check([]byte{byte(v)})
		// Embed at offsets that cross the 8-byte word boundary. A high byte here
		// is an isolated continuation/start byte between ASCII runs — ill-formed,
		// so it must come back as U+FFFD.
		buf := append([]byte("clean123"), byte(v))
		buf = append(buf, "clean456"...)
		check(buf)
	}

	// UTF-8 corners: valid multibyte sequences (which must pass through verbatim
	// and not break clean runs), ill-formed shapes DecodeRune rejects byte by byte
	// (truncations, bare continuations, overlongs, surrogates, out-of-range), and
	// both straddling the scanner boundaries — the 8-byte SWAR word, the 16/32-byte
	// SIMD blocks, and the 48-byte minVectorRun gate. Escapes on both sides of a
	// multibyte char pin the dispatch back into the valid walk.
	utf8Cases := [][]byte{
		[]byte("héllo"),
		[]byte("héllo \"wörld\"\n"),
		[]byte("日本語のテキスト"),
		[]byte("emoji \U0001F600 tail"),
		[]byte("\xc3"),                        // truncated 2-byte sequence
		[]byte("abc\xc3"),                     // truncated at end of input
		[]byte("\x80tail"),                    // bare continuation byte first
		[]byte("a\xed\xa0\x80b"),              // surrogate half (rejected per byte)
		[]byte("a\xc0\xafb"),                  // overlong encoding
		[]byte("a\xf5\x80\x80\x80b"),          // out of range
		[]byte("\xc3\x28"),                    // bad continuation
		[]byte("ok\xffbad\xfe"),               // 0xff/0xfe never appear in UTF-8
		[]byte(strings.Repeat("x", 47) + "é"), // straddles minVectorRun
		[]byte(strings.Repeat("x", 48) + "é" + strings.Repeat("y", 48)),
		[]byte(strings.Repeat("x", 63) + "\xc3\xa9"),                       // straddles a 64-byte grid
		[]byte(strings.Repeat("x", 31) + "\xc3" + strings.Repeat("y", 32)), // invalid at block edge
		[]byte(strings.Repeat("é", 40)),                                    // dense valid multibyte
		[]byte(strings.Repeat("\xff", 40)),                                 // dense invalid
		[]byte("\"quote\x80then\\slash"),                                   // escapes around invalid
		[]byte("é\"é"),                                                     // escape between valid multibyte
	}
	for _, c := range utf8Cases {
		check(c)
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

	// A second fuzz interleaving whole valid runes (multibyte included) with raw
	// high bytes and escapes, so well-formed sequences and ill-formed bytes mix in
	// the same input — the shape the per-byte fuzz above only rarely produces.
	for n := 0; n < 5000; n++ {
		var b []byte
		for k := r.Intn(16); k > 0; k-- {
			switch r.Intn(6) {
			case 0:
				b = utf8.AppendRune(b, rune(0x80+r.Intn(0x2FFFF)))
			case 1:
				b = append(b, byte(0x80+r.Intn(0x80))) // raw high byte
			case 2:
				b = append(b, []byte{'"', '\\', '\n', 0x01}[r.Intn(4)])
			default:
				b = append(b, byte(0x20+r.Intn(0x5f))) // clean ASCII
			}
		}
		check(b)
	}
}

// TestEscapeStringMatchesStdlibCoercion pins the doc claim that the escapers
// substitute U+FFFD for ill-formed UTF-8 "exactly as encoding/json replaces them
// when marshaling". The two escapers spell things differently — encoding/json
// HTML-escapes <, > and &, and writes the replacement character as the six-byte
// \ufffd escape where lightning emits its raw three UTF-8 bytes — so equality is
// checked where it is defined: the string VALUE both encodings decode back to.
// A premise check asserts the stdlib really coerced each ill-formed input, so
// the differential cannot pass vacuously on a corpus of valid strings.
func TestEscapeStringMatchesStdlibCoercion(t *testing.T) {
	corpus := [][]byte{
		[]byte("plain ascii"),
		[]byte("héllo \"wörld\"\n"),
		[]byte("日本語 emoji \U0001F600"),
		[]byte("\xc3"),               // truncated 2-byte sequence
		[]byte("abc\xc3"),            // truncated at end
		[]byte("\x80tail"),           // bare continuation byte
		[]byte("a\xed\xa0\x80b"),     // surrogate half
		[]byte("a\xc0\xafb"),         // overlong
		[]byte("a\xf5\x80\x80\x80b"), // out of range
		[]byte("ok\xffbad\xfe"),
		[]byte(strings.Repeat("\xff", 40)),
		[]byte("\"quote\x80then\\slash\tand\x01ctl"),
	}
	for _, bc := range escapeBenchCases {
		corpus = append(corpus, bc.in)
	}
	sawCoercion := false
	for _, in := range corpus {
		stdEnc, err := ejson.Marshal(string(in))
		if err != nil {
			t.Fatalf("stdlib Marshal(%q): %v", in, err)
		}
		var want string
		if err := ejson.Unmarshal(stdEnc, &want); err != nil {
			t.Fatalf("stdlib round-trip of %q: %v", in, err)
		}
		if !utf8.Valid(in) {
			if want == string(in) {
				t.Fatalf("premise broken: stdlib did not coerce ill-formed %q", in)
			}
			sawCoercion = true
		}
		doc := append(EscapeStringInto(in, []byte{'"'}), '"')
		var got string
		if err := ejson.Unmarshal(doc, &got); err != nil {
			t.Fatalf("EscapeStringInto(%q) produced %q, not a decodable JSON string: %v", in, doc, err)
		}
		if got != want {
			t.Fatalf("EscapeStringInto(%q) decodes to %q, encoding/json's escaping decodes to %q", in, got, want)
		}
	}
	if !sawCoercion {
		t.Fatal("test bug: corpus held no ill-formed input, the differential proved nothing")
	}
}

// TestEscapeStringIntoOutBuffers locks the out contract EscapeStringInto
// documents — out must not overlap s — by exercising the buffers that satisfy it:
// a nil out, an out that already holds content (the escaped form is *appended*,
// leaving what was there), and an out reused across calls with out[:0]. Each one
// must yield exactly escapeReference's answer and leave s untouched.
//
// The aliased form the doc used to invite (out == s[:0], by analogy with
// UnescapeStringInto) is deliberately not exercised: escaping lengthens, so the
// appends overrun the bytes the scan has yet to read and the result is garbage.
// There is nothing to lock about it beyond the sentence in the doc.
func TestEscapeStringIntoOutBuffers(t *testing.T) {
	const prefix = `{"k":"`
	for _, bc := range escapeBenchCases {
		t.Run(bc.name, func(t *testing.T) {
			want := escapeReference(bc.in)
			orig := string(bc.in)

			if got := string(EscapeStringInto(bc.in, nil)); got != want {
				t.Errorf("nil out: got %q, want %q", got, want)
			}

			// An out with content: EscapeStringInto appends, so the caller can
			// build a whole JSON member in one buffer.
			if got := string(EscapeStringInto(bc.in, []byte(prefix))); got != prefix+want {
				t.Errorf("out with content: got %q, want %q", got, prefix+want)
			}

			// A reused buffer, both when it has room and when it must grow.
			for _, capacity := range []int{0, 1, 6*len(bc.in) + 8} {
				buf := make([]byte, 0, capacity)
				if got := string(EscapeStringInto(bc.in, buf[:0])); got != want {
					t.Errorf("reused out (cap %d): got %q, want %q", capacity, got, want)
				}
			}

			if string(bc.in) != orig {
				t.Errorf("input mutated: %q -> %q", orig, bc.in)
			}
		})
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
	{"unicode_clean", []byte(strings.Repeat("Motörhead spelade i Köln — 日本語テキスト här. ", 4))},
	{"unicode_with_quotes", []byte("Hon sa \"hej\" på svenska,\nsen 中文字 och é och \\ till slut.")},
	{"invalid_utf8_one_byte", []byte(strings.Repeat("plain text content ", 16) + "\xff")},
	{"invalid_utf8_dense", []byte(strings.Repeat("ok\xff\xc3(bad\x80 ", 12))},
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
