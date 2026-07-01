package unstable

import (
	"encoding/json"
	"math/rand"
	"strings"
	"testing"
)

// refSkip is the known-correct existing skip, used as the differential oracle.
func refSkip(data []byte, i int) (int, error) {
	switch data[i] {
	case '{':
		return skipObject(data, i)
	case '[':
		return skipArray(data, i)
	}
	panic("refSkip: not a container")
}

func randString(r *rand.Rand) string {
	n := r.Intn(40)
	// Deliberately tricky alphabet: quotes, backslashes, brackets/braces, colon,
	// comma, whitespace, control chars and multibyte runes — so the marshaled JSON
	// carries \" \\ \uXXXX escapes and raw structural bytes inside string bodies.
	alpha := []rune("ab\"\\{}[]:, \n\t\x01\x1f日x")
	b := make([]rune, n)
	for i := range b {
		b[i] = alpha[r.Intn(len(alpha))]
	}
	return string(b)
}

func randValue(r *rand.Rand, depth int) any {
	if depth <= 0 || r.Intn(3) == 0 {
		switch r.Intn(4) {
		case 0:
			return r.Intn(1000000)
		case 1:
			return r.Float64() * 1000
		case 2:
			return randString(r)
		default:
			return r.Intn(2) == 0
		}
	}
	if r.Intn(2) == 0 {
		m := map[string]any{}
		for k := r.Intn(6); k > 0; k-- {
			m[randString(r)] = randValue(r, depth-1)
		}
		return m
	}
	s := make([]any, r.Intn(6))
	for k := range s {
		s[k] = randValue(r, depth-1)
	}
	return s
}

func TestSkipContainerFastMatches(t *testing.T) {
	// Hand-picked tricky cases (brackets/quotes/escapes inside strings).
	cases := []string{
		`{}`, `[]`, `{"a":"b"}`, `[1,2,3]`,
		`{"a":"}"}`, `{"a":"]"}`, `{"a":"{[}]"}`,
		`{"a":"\""}`, `{"a":"\\"}`, `{"a":"\\\""}`, `{"a":"\\\\"}`,
		`{"a":{"b":[1,2,"]"]},"c":"}"}`,
		`["\"","}","{","[\\]"]`,
		`{"k":"a\\\\\\\"b}c"}`,
	}
	for _, c := range cases {
		data := []byte(c)
		want, werr := refSkip(data, 0)
		got, gerr := skipContainerFast(data, 0, data[0])
		if got != want || (werr == nil) != (gerr == nil) {
			t.Errorf("%q: fast=(%d,%v) ref=(%d,%v)", c, got, gerr, want, werr)
		}
	}

	r := rand.New(rand.NewSource(1))
	for iter := 0; iter < 50000; iter++ {
		var v any
		if r.Intn(2) == 0 {
			m := map[string]any{}
			for k := r.Intn(6) + 1; k > 0; k-- {
				m[randString(r)] = randValue(r, 3)
			}
			v = m
		} else {
			s := make([]any, r.Intn(6)+1)
			for k := range s {
				s[k] = randValue(r, 3)
			}
			v = s
		}
		doc, err := json.Marshal(v)
		if err != nil {
			t.Fatal(err)
		}
		// Test both the tail path (close at EOF) and the block path (close found
		// inside a 32-byte block, with >=32 bytes of trailing slack after it).
		for _, pad := range []int{0, 40} {
			data := append(append([]byte{}, doc...), []byte(strings.Repeat(" ", pad))...)
			want, werr := refSkip(data, 0)
			got, gerr := skipContainerFast(data, 0, data[0])
			if got != want || (werr == nil) != (gerr == nil) {
				t.Fatalf("iter %d pad %d: fast=(%d,%v) ref=(%d,%v)\ndoc=%s",
					iter, pad, got, gerr, want, werr, doc)
			}
			// SkipValue is now the production dispatch; it must agree with the
			// pre-existing skipObject/skipArray oracle on every valid value.
			dgot, derr := SkipValue(data, 0)
			if dgot != want || (werr == nil) != (derr == nil) {
				t.Fatalf("iter %d pad %d: SkipValue=(%d,%v) ref=(%d,%v)\ndoc=%s",
					iter, pad, dgot, derr, want, werr, doc)
			}
		}
		// Truncation safety: a cut-off container must fail (not panic / OOB / read
		// past end) on both the dispatch and the oracle — Get/Set rely on the error
		// to reject malformed input. Any suffix of a single top-level container
		// drops its close, so both must error.
		if len(doc) > 2 {
			t1 := doc[:1+r.Intn(len(doc)-1)]
			_, e1 := refSkip(t1, 0)
			_, e2 := SkipValue(t1, 0)
			if (e1 == nil) != (e2 == nil) {
				t.Fatalf("iter %d trunc %d: ref err=%v SkipValue err=%v\ndoc=%s",
					iter, len(t1), e1, e2, doc)
			}
		}
	}
}

// boundaryDocs builds containers engineered around the 64-byte block grid: a
// backslash run of every parity crossing the block boundary at every offset
// (the prevEscaped carry between blocks), quotes landing exactly on the
// boundary, deep bracket runs (the popcount bulk path and >64 depth), close
// brackets at exact block edges, and close-heavy blocks that must take the
// per-bit walk.
func boundaryDocs() []string {
	var docs []string
	// Escape runs straddling the boundary at every alignment and parity. The
	// container scan starts at index 1, so pad lengths 40..80 sweep the run and
	// its terminating quote across bytes 56..72 of the first block.
	for pad := 40; pad <= 80; pad++ {
		for nbs := 0; nbs <= 5; nbs++ {
			body := strings.Repeat("x", pad) + strings.Repeat(`\`, nbs) + `"`
			if nbs%2 == 1 {
				body += `tail"` // the quote was escaped; close the string for real
			}
			docs = append(docs, `{"a":"`+body+`,"b":{"c":[1,"]"]}}`)
		}
	}
	// Deep pure-bracket nesting: block-spanning open runs (bulk adds), then a
	// close run that walks depth back down across blocks (bulk subtract until
	// depth <= popcount, then the bit walk finds the crossing).
	for n := 1; n <= 200; n += 7 {
		docs = append(docs, strings.Repeat("[", n)+"1"+strings.Repeat("]", n))
	}
	// Close-dense blocks: many sibling empty objects, so op/cl interleave and
	// depth stays low (the per-bit path fires every block).
	docs = append(docs, "["+strings.TrimSuffix(strings.Repeat("{},", 100), ",")+"]")
	// Containers whose close lands at/around exact block multiples.
	for _, n := range []int{62, 63, 64, 65, 126, 127, 128, 129, 191, 192, 193} {
		docs = append(docs, "["+strings.Repeat(" ", n)+"]")
		docs = append(docs, `{"k":"`+strings.Repeat("y", n)+`"}`)
	}
	return docs
}

// TestSkipContainerBoundaries runs the crafted block-boundary corpus through the
// production skipContainerFast (whatever implementation is live on this arch)
// against the scalar oracle, with and without trailing slack so both the block
// and tail paths see every case.
func TestSkipContainerBoundaries(t *testing.T) {
	for _, c := range boundaryDocs() {
		for _, pad := range []int{0, 70} {
			data := []byte(c + strings.Repeat(" ", pad))
			want, werr := refSkip(data, 0)
			got, gerr := skipContainerFast(data, 0, data[0])
			if got != want || (werr == nil) != (gerr == nil) {
				t.Fatalf("%q pad %d: fast=(%d,%v) ref=(%d,%v)", c, pad, got, gerr, want, werr)
			}
		}
	}
}

// bigStringObject builds a string-heavy object: many "key":"value" pairs, the
// shape Get/GetPaths skips over and where per-string SkipString calls dominate.
func bigStringObject(pairs int) []byte {
	var b strings.Builder
	b.WriteByte('{')
	for i := 0; i < pairs; i++ {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteString(`"key_field_number_`)
		b.WriteString(strings.Repeat("x", i%7))
		b.WriteString(`":"some string value with \"escapes\" and unicode é here"`)
	}
	b.WriteByte('}')
	return []byte(b.String())
}

// numberObject: string keys but scalar (number) values — the ambiguous case for
// a dispatch heuristic, since the keys are strings but the values are inert.
func numberObject(n int) []byte {
	var b strings.Builder
	b.WriteByte('{')
	for i := 0; i < n; i++ {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteString(`"key_field_number_`)
		b.WriteString(strings.Repeat("x", i%7))
		b.WriteString(`":123456.789`)
	}
	b.WriteByte('}')
	return []byte(b.String())
}

func numberArray(n int) []byte {
	var b strings.Builder
	b.WriteByte('[')
	for i := 0; i < n; i++ {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteString("123456.789")
	}
	b.WriteByte(']')
	return []byte(b.String())
}

// nestedMixed: small objects in an array, a mix of strings/numbers — closer to a
// real record array Get/GetPaths walks over.
func nestedMixed(n int) []byte {
	var b strings.Builder
	b.WriteByte('[')
	for i := 0; i < n; i++ {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteString(`{"id":12345,"name":"a name","tags":[1,2,3],"ok":true}`)
	}
	b.WriteByte(']')
	return []byte(b.String())
}

func BenchmarkSkipContainer(b *testing.B) {
	shapes := []struct {
		name string
		data []byte
		open byte
	}{
		{"stringObj", bigStringObject(300), '{'},
		{"numberObj", numberObject(300), '{'},
		{"numberArr", numberArray(600), '['},
		{"nestedMixed", nestedMixed(200), '['},
	}
	for _, s := range shapes {
		s := s
		// current: the pre-existing path directly (skipObject/skipArray), the A/B
		// baseline. dispatch: SkipValue, which now routes through skipContainerFast.
		b.Run(s.name+"/current", func(b *testing.B) {
			b.SetBytes(int64(len(s.data)))
			for i := 0; i < b.N; i++ {
				var err error
				if s.open == '{' {
					_, err = skipObject(s.data, 0)
				} else {
					_, err = skipArray(s.data, 0)
				}
				if err != nil {
					b.Fatal(err)
				}
			}
		})
		b.Run(s.name+"/dispatch", func(b *testing.B) {
			b.SetBytes(int64(len(s.data)))
			for i := 0; i < b.N; i++ {
				if _, err := SkipValue(s.data, 0); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
