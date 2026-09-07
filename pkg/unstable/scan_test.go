package unstable

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

// scannerCorpus is one value per entry, each followed by the trailing bytes a
// walker would see after it, so the scan has to stop in the right place rather
// than at the end of the input.
func scannerCorpus() []string {
	vals := []string{
		`0`, `-1`, `1788087600123`, `1.5e-7`, `-65.613617000000029`,
		`true`, `false`, `null`,
		`""`, `"a"`, `"with \"escapes\" and \\ and é"`, `"]},["`, `"\\"`, `"\\\\"`,
		`{}`, `[]`, `{"a":1}`, `[1,2,3]`, `{"a":[1,{"b":"]"}],"c":{}}`,
		`[[[[[1]]]]]`, `{"k":{"k":{"k":{"k":"v"}}}}`,
		`["a","b","c"]`, `[{"x":1},{"x":2}]`,
		`{"s":"a\"b","n":-0.5e+3,"t":true,"f":false,"z":null,"a":[],"o":{}}`,
	}
	// A long string and a deep-ish container, both past any one chunk size.
	vals = append(vals, `"`+strings.Repeat("x", 5000)+`"`)
	vals = append(vals, `[`+strings.Repeat(`{"a":"bcd"},`, 400)+`{"a":"bcd"}]`)
	vals = append(vals, `"`+strings.Repeat(`\"`, 300)+`"`)
	out := make([]string, 0, len(vals)*4)
	for _, v := range vals {
		for _, tail := range []string{"", ",", "]", "}", " ", "\n\t ", `,"next":1}`} {
			out = append(out, v+tail)
		}
	}
	return out
}

// feedInChunks runs a whole document through the scanner in fixed-size pieces,
// which is what a caller reading from a network does.
func feedInChunks(t *testing.T, doc string, size int) (int, bool, error) {
	return feedInChunksMode(t, doc, size, false)
}

// feedInChunksFast is the same through the block-scan balance.
func feedInChunksFast(t *testing.T, doc string, size int) (int, bool, error) {
	return feedInChunksMode(t, doc, size, true)
}

func feedInChunksMode(t *testing.T, doc string, size int, fast bool) (int, bool, error) {
	t.Helper()
	var s ValueScanner
	if fast {
		s.ResetFast()
	} else {
		s.Reset()
	}
	pos := 0
	for {
		hi := pos + size
		if hi > len(doc) {
			hi = len(doc)
		}
		final := hi == len(doc)
		n, done, err := s.Feed([]byte(doc[pos:hi]), final)
		if n > hi-pos {
			t.Fatalf("Feed consumed %d of a %d-byte chunk", n, hi-pos)
		}
		pos += n
		if err != nil || done {
			return pos, done, err
		}
		if final {
			t.Fatalf("Feed(%q) neither finished nor errored at the end of the input", doc)
		}
		if n != hi-pos+n { // the chunk must be fully consumed when not done
			t.Fatalf("Feed left %d bytes of a chunk unconsumed without finishing", hi-pos-n)
		}
	}
}

// TestValueScannerMatchesSkipValue is the contract: on a well-formed value the
// resumable scan ends exactly where SkipValue does, at every chunk size.
func TestValueScannerMatchesSkipValue(t *testing.T) {
	for _, doc := range scannerCorpus() {
		want, err := SkipValue([]byte(doc), 0)
		if err != nil {
			t.Fatalf("premise: SkipValue(%.40q) = %v", doc, err)
		}
		for _, size := range []int{1, 2, 3, 7, 16, 64, 4096, len(doc), len(doc) + 1} {
			if size <= 0 {
				continue
			}
			got, done, err := feedInChunks(t, doc, size)
			if err != nil || !done || got != want {
				t.Fatalf("chunk %d over %.40q: end=%d done=%v err=%v; SkipValue says %d",
					size, doc, got, done, err, want)
			}
		}
	}
}

// TestValueScannerFastMatchesSkipValue is ResetFast's contract: on a
// WELL-FORMED value the block-scan balance ends exactly where SkipValue does,
// at every chunk size, so a caller stepping past values gets the in-memory
// walkers' answer. (What the mode changes is malformed input, which ResetFast
// enumerates by pointing at skipfast.go.)
//
// The sizes either side of 64 are the point: the block scan is handed the
// depth and nothing else, so a chunk boundary that falls inside a string — or
// just after a backslash — is state it cannot carry, and the first version of
// this got a 900-byte padding string wrong at exactly those splits while
// passing at 1, 7 and "the whole value at once".
func TestValueScannerFastMatchesSkipValue(t *testing.T) {
	if !useSkipBlocks {
		t.Skip("no block scan on this machine; ResetFast is the scalar balance")
	}
	for _, doc := range scannerCorpus() {
		want, err := SkipValue([]byte(doc), 0)
		if err != nil {
			t.Fatalf("premise: SkipValue(%.40q) = %v", doc, err)
		}
		for _, size := range []int{1, 2, 3, 7, 16, 63, 64, 65, 100, 127, 128, 129, 200, 4096, len(doc), len(doc) + 1} {
			if size <= 0 {
				continue
			}
			got, done, err := feedInChunksFast(t, doc, size)
			if err != nil || !done || got != want {
				t.Fatalf("fast chunk %d over %.40q: end=%d done=%v err=%v; SkipValue says %d",
					size, doc, got, done, err, want)
			}
		}
	}
}

// TestValueScannerFastStraddlesStrings walks a chunk boundary across every
// byte of a value whose strings hold quotes, backslashes and brackets, which
// is every position the carried state can be wrong in.
func TestValueScannerFastStraddlesStrings(t *testing.T) {
	if !useSkipBlocks {
		t.Skip("no block scan on this machine")
	}
	body := `{"a":"x]}\"y","b":["p\\","q\"r",{"c":"[[["}],"d":1}`
	for _, pad := range []int{0, 1, 7, 31, 60, 61, 62, 63, 64, 65, 100} {
		doc := `{"pad":"` + strings.Repeat("z", pad) + `","v":` + body + `},`
		want, err := SkipValue([]byte(doc), 0)
		if err != nil {
			t.Fatalf("premise: SkipValue(%.40q) = %v", doc, err)
		}
		for size := 1; size <= len(doc)+1; size++ {
			got, done, err := feedInChunksFast(t, doc, size)
			if err != nil || !done || got != want {
				t.Fatalf("fast chunk %d pad %d over %.60q: end=%d done=%v err=%v; want %d",
					size, pad, doc, got, done, err, want)
			}
		}
	}
}

// TestValueScannerChunkingIsInvisible holds every split of every input —
// malformed ones included — to the answer the scanner gives when it is handed
// the whole input at once. Where the two could differ is exactly the state
// that has to survive a refill: an escape or a \uXXXX cut in half, a token cut
// in half, a container's depth.
func TestValueScannerChunkingIsInvisible(t *testing.T) {
	docs := append(scannerCorpus(),
		// Malformed, in each of the ways the scan can fail.
		`"unterminated`, `"trailing escape\`, `{"a":1`, `[1,2`, `tru`, `truex`, `nul`, `fals`,
		`x`, ``, `,`, `:`, `}`, `]`, `{]`, `[}`, `{"a":[}`, `[{]`, `"\`, `"\u00`,
		`{"a":"b"`, `[[[`, `-`, `.`, `e`, `+`, `1.2.3`, `--`, `[1,]`, `{,}`,
	)
	for _, doc := range docs {
		wantN, wantDone, wantErr := feedInChunks(t, doc, len(doc)+1)
		for _, size := range []int{1, 2, 3, 5, 8, 13, 64} {
			gotN, gotDone, gotErr := feedInChunks(t, doc, size)
			if gotN != wantN || gotDone != wantDone || !errors.Is(gotErr, wantErr) {
				t.Fatalf("chunk %d over %.40q: (%d,%v,%v); whole input gives (%d,%v,%v)",
					size, doc, gotN, gotDone, gotErr, wantN, wantDone, wantErr)
			}
		}
	}
}

// TestValueScannerErrors pins the sentinel for each way a value can be
// rejected, since a caller distinguishes them.
func TestValueScannerErrors(t *testing.T) {
	cases := []struct {
		in   string
		want error
	}{
		{`x`, ErrBadNumber},
		{`:`, ErrBadNumber},
		{`}`, ErrBadNumber},
		{``, ErrTruncated},
		{`"abc`, ErrTruncated},
		{`{"a":1`, ErrTruncated},
		{`[1,2`, ErrTruncated},
		{`tru`, ErrInvalidJSON},
		{`truX`, ErrInvalidJSON},
		{`nulX`, ErrInvalidJSON},
		{`falsX`, ErrInvalidJSON},
	}
	for _, c := range cases {
		_, done, err := feedInChunks(t, c.in, 1)
		if done || !errors.Is(err, c.want) {
			t.Errorf("scan(%q) = done=%v err=%v, want %v", c.in, done, err, c.want)
		}
	}
	// The depth bound is the recursive skips' exactly: MaxDepth containers
	// scan, and the one past that is ErrMaxDepth. Both sides are checked
	// against the scalar skip so the boundary cannot drift by one.
	for _, d := range []int{MaxDepth - 1, MaxDepth, MaxDepth + 1, MaxDepth + 2} {
		doc := strings.Repeat("[", d) + strings.Repeat("]", d)
		wantEnd, wantErr := scalarSkipValue([]byte(doc), 0)
		gotEnd, done, gotErr := feedInChunks(t, doc, 64)
		if wantErr != nil {
			if done || !errors.Is(gotErr, wantErr) {
				t.Errorf("scan at depth %d = (%d,%v,%v); the scalar skip says %v", d, gotEnd, done, gotErr, wantErr)
			}
			continue
		}
		if !done || gotErr != nil || gotEnd != wantEnd {
			t.Errorf("scan at depth %d = (%d,%v,%v), want (%d,true,nil)", d, gotEnd, done, gotErr, wantEnd)
		}
	}
}

func BenchmarkValueScanner(b *testing.B) {
	docs := map[string]string{
		"records": `[` + strings.Repeat(`{"name":"record","value":12345,"active":true},`, 200) + `{"a":1}]`,
		"strings": `[` + strings.Repeat(`"entity-service-20260903-1",`, 400) + `"x"]`,
		"numbers": `[` + strings.Repeat(`1788087600123,`, 600) + `1]`,
	}
	for name, doc := range docs {
		d := []byte(doc)
		b.Run(fmt.Sprintf("%s/whole", name), func(b *testing.B) {
			b.SetBytes(int64(len(d)))
			var s ValueScanner
			for i := 0; i < b.N; i++ {
				s.Reset()
				if _, done, err := s.Feed(d, true); !done || err != nil {
					b.Fatal(done, err)
				}
			}
		})
		b.Run(fmt.Sprintf("%s/4k", name), func(b *testing.B) {
			b.SetBytes(int64(len(d)))
			var s ValueScanner
			for i := 0; i < b.N; i++ {
				s.Reset()
				pos := 0
				for {
					hi := pos + 4096
					if hi > len(d) {
						hi = len(d)
					}
					n, done, err := s.Feed(d[pos:hi], hi == len(d))
					pos += n
					if err != nil {
						b.Fatal(err)
					}
					if done {
						break
					}
				}
			}
		})
	}
}

// scalarSkipValue is SkipValue with the scalar container skips forced, which
// is the semantics the resumable scanner mirrors. SkipValue itself may take
// the SIMD in-string-mask path, whose answers on MALFORMED input differ from
// the scalar path's in the three ways skipfast.go enumerates; this oracle is
// what lets the scanner be held to a definite answer on such input rather than
// to whichever path the host CPU selects.
func scalarSkipValue(data []byte, i int) (int, error) {
	if uint(i) >= uint(len(data)) {
		return i, ErrTruncated
	}
	switch data[i] {
	case '"':
		return SkipString(data, i)
	case '{':
		return skipObject(data, i)
	case '[':
		return skipArray(data, i)
	case 't':
		if i+4 <= len(data) && string(data[i:i+4]) == "true" {
			return i + 4, nil
		}
		return i, ErrInvalidJSON
	case 'f':
		if i+5 <= len(data) && string(data[i:i+5]) == "false" {
			return i + 5, nil
		}
		return i, ErrInvalidJSON
	case 'n':
		return ExpectNull(data, i)
	}
	return SkipNumber(data, i)
}

// TestValueScannerMatchesScalarSkip is the malformed-input half of the
// contract: the balance is TYPED, so a closer of the kind the innermost
// container did not open is stepped over rather than counted, exactly as
// skipObject and skipArray do it through their recursion. Nothing in a
// well-formed document exercises that, which is why these inputs are all
// broken on purpose.
func TestValueScannerMatchesScalarSkip(t *testing.T) {
	docs := []string{
		`[}`, `{]`, `[}]`, `{]}`, `{"a":[}]}`, `[{]}]`, `[[}]]`, `{{]}}`,
		`[1,}`, `{"a":1]`, `["]"]`, `{"}":1}`, `[{"a":[1,2}]}]`,
		`{"a":{"b":]}}`, `[]]`, `{}}`, `[[]`, `{{}`, `["a"}`, `{"a"]`,
		`{"a":1}}`, `[1]]`, `[,]`, `{,}`, `[:]`, `{"a"::1}`,
		`"unterminated`, `{"a":"b`, `[1,2`, `tru`, `nul`, `fals`, `truex`,
		`[1,]`, `{"a":1,}`, `1x`, `-`, `.`, `+`, `e`, `1.2.3`,
	}
	for _, doc := range docs {
		wantEnd, wantErr := scalarSkipValue([]byte(doc), 0)
		var s ValueScanner
		s.Reset()
		gotEnd, done, gotErr := s.Feed([]byte(doc), true)
		if wantErr == nil {
			if !done || gotErr != nil || gotEnd != wantEnd {
				t.Errorf("scan(%q) = (%d,%v,%v); the scalar skip ends at %d", doc, gotEnd, done, gotErr, wantEnd)
			}
			continue
		}
		if done || !errors.Is(gotErr, wantErr) {
			t.Errorf("scan(%q) = (%d,%v,%v); the scalar skip says %v", doc, gotEnd, done, gotErr, wantErr)
		}
	}
}
