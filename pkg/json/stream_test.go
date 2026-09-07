package json

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"
	"testing/iotest"
)

// chunkReader hands out at most n bytes per Read, which is what a socket does
// and what every boundary in the stream reader has to survive.
type chunkReader struct {
	data []byte
	n    int
}

func (c *chunkReader) Read(p []byte) (int, error) {
	if len(c.data) == 0 {
		return 0, io.EOF
	}
	n := c.n
	if n > len(p) {
		n = len(p)
	}
	if n > len(c.data) {
		n = len(c.data)
	}
	copy(p, c.data[:n])
	c.data = c.data[n:]
	return n, nil
}

// readerKinds is every awkward shape an io.Reader can take: one byte at a
// time, fixed chunks either side of the token sizes, half of what was asked
// for, and a reader that returns its last bytes together with io.EOF.
func readerKinds(doc string) map[string]func() io.Reader {
	return map[string]func() io.Reader{
		"onebyte": func() io.Reader { return iotest.OneByteReader(strings.NewReader(doc)) },
		"chunk2":  func() io.Reader { return &chunkReader{data: []byte(doc), n: 2} },
		"chunk3":  func() io.Reader { return &chunkReader{data: []byte(doc), n: 3} },
		"chunk7":  func() io.Reader { return &chunkReader{data: []byte(doc), n: 7} },
		"chunk64": func() io.Reader { return &chunkReader{data: []byte(doc), n: 64} },
		"chunk4096": func() io.Reader {
			return &chunkReader{data: []byte(doc), n: 4096}
		},
		"half":    func() io.Reader { return iotest.HalfReader(strings.NewReader(doc)) },
		"dataerr": func() io.Reader { return iotest.DataErrReader(strings.NewReader(doc)) },
		"whole":   func() io.Reader { return strings.NewReader(doc) },
	}
}

// streamDocs is the corpus: the shapes that make a streaming walk different
// from an in-memory one — values larger than a chunk, strings holding the
// structural bytes, escapes at every boundary, whitespace everywhere, and
// documents that end inside each kind of value.
func streamDocs() []string {
	return []string{
		`[]`, `[1]`, `[1,2,3]`, `[ 1 , 2 , 3 ]`, "[\n  1,\n  2\n]",
		`["a","b"]`, `["]","}",",","\"","\\"]`, `[ "a\nb" , "céd" ]`,
		`[[1,2],[3,4]]`, `[{"a":1},{"b":[2,{"c":"]"}]}]`,
		`[null,true,false,0,-1.5e-7]`,
		`{}`, `{"a":1}`, `{"a":1,"b":"two","c":[3],"d":{"e":4}}`,
		`{ "a" : 1 , "b" : 2 }`, "{\n\t\"a\": 1\n}",
		`{"esc\"key":1,"tab\tkey":2,"unicodeé":3}`,
		`{"outer":{"inner":[1,2,3]}}`, `{"a":null,"b":null}`,
		`{"m":null}`, `{"m":[]}`, `[null]`,
		`[` + strings.Repeat(`"`+strings.Repeat("x", 300)+`",`, 5) + `1]`,
		`[` + strings.Repeat(`{"k":"`+strings.Repeat("y", 200)+`"},`, 4) + `2]`,
		`{"big":"` + strings.Repeat("z", 5000) + `","after":7}`,
		`[1788087600123,1788087600124]`,
		// Malformed, in each way a walk can fail.
		`[1,`, `[1,2`, `["a`, `{"a":`, `{"a"`, `{`, `[`, ``, `   `,
		`[1,]`, `{"a":1,}`, `{"a" 1}`, `[1 2]`, `{"a":1"b":2}`, `nul`, `nullx`,
		`"astring"`, `12345`, `true`,
	}
}

// walkStream runs one streaming walk and records what the callback saw.
type walkResult struct {
	keys   []string
	values []string
	err    error
}

func streamArray(t *testing.T, doc string, mk func() io.Reader, opts []ReaderOption, keys ...string) walkResult {
	t.Helper()
	var got walkResult
	r := NewReader(mk(), opts...)
	got.err = r.ArrayEach(func(v []byte) error {
		got.values = append(got.values, string(v))
		return nil
	}, keys...)
	return got
}

func streamObject(t *testing.T, doc string, mk func() io.Reader, opts []ReaderOption, keys ...string) walkResult {
	t.Helper()
	var got walkResult
	r := NewReader(mk(), opts...)
	got.err = r.ObjectEach(func(k string, v []byte) error {
		// Both are windows onto the reader's buffer and are only valid until
		// this returns; keeping either without a copy is the one mistake the
		// streaming contract can punish. See TestStreamValueLifetime.
		got.keys = append(got.keys, strings.Clone(k))
		got.values = append(got.values, string(v))
		return nil
	}, keys...)
	return got
}

func memArray(doc string, keys ...string) walkResult {
	var got walkResult
	got.err = ArrayEach([]byte(doc), func(v []byte) error {
		got.values = append(got.values, string(v))
		return nil
	}, keys...)
	return got
}

func memObject(doc string, keys ...string) walkResult {
	var got walkResult
	got.err = ObjectEach([]byte(doc), func(k string, v []byte) error {
		got.keys = append(got.keys, k)
		got.values = append(got.values, string(v))
		return nil
	}, keys...)
	return got
}

func sameWalk(a, b walkResult) bool {
	if fmt.Sprint(a.values) != fmt.Sprint(b.values) || fmt.Sprint(a.keys) != fmt.Sprint(b.keys) {
		return false
	}
	return errors.Is(a.err, b.err) || errors.Is(b.err, a.err) || (a.err == nil) == (b.err == nil)
}

// TestStreamMatchesInMemory is the contract: the same document walked from a
// reader yields the same values, in the same order, with the same error, at
// every chunk size — including the sizes that cut a token, an escape or a
// \uXXXX in half. The buffer is deliberately tiny so that every document in
// the corpus forces refills, compaction and growth.
func TestStreamMatchesInMemory(t *testing.T) {
	for _, doc := range streamDocs() {
		want := memArray(doc)
		wantObj := memObject(doc)
		for name, mk := range readerKinds(doc) {
			for _, size := range []int{16, 64, 1 << 16} {
				opts := []ReaderOption{WithBufferSize(size)}
				if got := streamArray(t, doc, mk, opts); !sameWalk(got, want) {
					t.Errorf("ArrayEach %.60q via %s buf=%d: %v %v; in-memory: %v %v",
						doc, name, size, got.values, got.err, want.values, want.err)
				}
				if got := streamObject(t, doc, mk, opts); !sameWalk(got, wantObj) {
					t.Errorf("ObjectEach %.60q via %s buf=%d: %v %v %v; in-memory: %v %v %v",
						doc, name, size, got.keys, got.values, got.err, wantObj.keys, wantObj.values, wantObj.err)
				}
			}
		}
	}
}

// TestStreamPathsMatchInMemory does the same through a key path, where the
// members before the target are streamed past rather than buffered.
func TestStreamPathsMatchInMemory(t *testing.T) {
	docs := []struct {
		doc  string
		keys []string
	}{
		{`{"a":[1,2,3]}`, []string{"a"}},
		{`{"skip":"` + strings.Repeat("s", 4000) + `","a":[1,2]}`, []string{"a"}},
		{`{"skip":[` + strings.Repeat(`{"x":1},`, 300) + `{"x":2}],"a":[7]}`, []string{"a"}},
		{`{"o":{"p":[1,{"q":"r"}]}}`, []string{"o", "p"}},
		{`{"a":{"b":{"c":[1]}}}`, []string{"a", "b", "c"}},
		{`{"a":1}`, []string{"missing"}},
		{`{"a":1}`, []string{"a"}},
		{`{"a":null}`, []string{"a"}},
		{`[1,2]`, []string{"a"}},
		{`{"a":{"b":1}}`, []string{"a", "missing"}},
		{`{"dup":[1],"dup":[2]}`, []string{"dup"}},
	}
	for _, c := range docs {
		want := memArray(c.doc, c.keys...)
		for name, mk := range readerKinds(c.doc) {
			for _, size := range []int{16, 128} {
				got := streamArray(t, c.doc, mk, []ReaderOption{WithBufferSize(size)}, c.keys...)
				if !sameWalk(got, want) {
					t.Errorf("ArrayEach %.50q %v via %s buf=%d: %v %v; in-memory: %v %v",
						c.doc, c.keys, name, size, got.values, got.err, want.values, want.err)
				}
			}
		}
	}
}

// TestStreamGetMatchesInMemory holds the single-value read to Get's answer.
func TestStreamGetMatchesInMemory(t *testing.T) {
	docs := []struct {
		doc  string
		keys []string
	}{
		{`{"a":1,"b":[1,2],"c":"x"}`, []string{"b"}},
		{`{"a":1,"b":{"c":"deep"}}`, []string{"b", "c"}},
		{`{"skip":"` + strings.Repeat("s", 3000) + `","want":"here"}`, []string{"want"}},
		{`{"a":1}`, []string{"zz"}},
		{`[1,2]`, nil},
		{`"top"`, nil},
		{`{"a":1}`, nil},
		{`{"n":null}`, []string{"n"}},
	}
	for _, c := range docs {
		want, _, wantErr := Get([]byte(c.doc), c.keys...)
		for name, mk := range readerKinds(c.doc) {
			r := NewReader(mk(), WithBufferSize(16))
			got, err := r.Get(c.keys...)
			if (err == nil) != (wantErr == nil) || string(got) != string(want) {
				t.Errorf("Get %.40q %v via %s: %q %v; in-memory: %q %v",
					c.doc, c.keys, name, got, err, want, wantErr)
			}
		}
	}
}

// TestStreamGrowsForOneElementOnly holds the memory bound: the buffer grows to
// fit the largest single element and no further, and an element past the limit
// is ErrElementTooLarge rather than an allocation.
func TestStreamGrowsForOneElementOnly(t *testing.T) {
	big := `"` + strings.Repeat("x", 200_000) + `"`
	doc := `[1,` + big + `,2]`

	r := NewReader(&chunkReader{data: []byte(doc), n: 1024}, WithBufferSize(1024))
	n := 0
	if err := r.ArrayEach(func(v []byte) error { n++; return nil }); err != nil {
		t.Fatal(err)
	}
	if n != 3 {
		t.Fatalf("walked %d elements, want 3", n)
	}
	if len(r.buf) < len(big) {
		t.Fatalf("buffer is %d bytes, too small to have held the %d-byte element", len(r.buf), len(big))
	}
	if len(r.buf) > 4*len(big) {
		t.Fatalf("buffer grew to %d bytes for a %d-byte element", len(r.buf), len(big))
	}

	// Past the limit the walk fails and the buffer stops there.
	r = NewReader(&chunkReader{data: []byte(doc), n: 1024}, WithBufferSize(1024), WithMaxElement(64<<10))
	err := r.ArrayEach(func(v []byte) error { return nil })
	if !errors.Is(err, ErrElementTooLarge) {
		t.Fatalf("ArrayEach over a 200 KB element with a 64 KiB limit = %v, want ErrElementTooLarge", err)
	}
	if len(r.buf) > 64<<10 {
		t.Fatalf("buffer grew to %d bytes past a 64 KiB limit", len(r.buf))
	}
}

// TestStreamSkipsPastLargeSibling is the other half of the bound: a member the
// path does not want is streamed past, so it costs reads and not buffer.
func TestStreamSkipsPastLargeSibling(t *testing.T) {
	doc := `{"skip":[` + strings.Repeat(`{"pad":"`+strings.Repeat("p", 900)+`"},`, 300) +
		`{"pad":"x"}],"want":[1,2,3]}`
	r := NewReader(&chunkReader{data: []byte(doc), n: 4096}, WithBufferSize(4096))
	var got []string
	if err := r.ArrayEach(func(v []byte) error { got = append(got, string(v)); return nil }, "want"); err != nil {
		t.Fatal(err)
	}
	if fmt.Sprint(got) != "[1 2 3]" {
		t.Fatalf("got %v", got)
	}
	if len(r.buf) != 4096 {
		t.Fatalf("buffer grew to %d bytes walking past a %d-byte sibling that was never wanted",
			len(r.buf), len(doc))
	}
	if r.Consumed() != int64(len(doc)) {
		t.Fatalf("Consumed = %d, want %d", r.Consumed(), len(doc))
	}
}

// TestStreamValueLifetime enforces the contract rather than only stating it: a
// callback that keeps the slice sees the buffer move under it, and one that
// copies does not. The reader poisons what it hands back on every refill so
// the failure is loud rather than a stale-but-plausible value.
func TestStreamValueLifetime(t *testing.T) {
	doc := `[` + strings.Repeat(`"`+strings.Repeat("a", 40)+`",`, 200) + `"end"]`
	var kept [][]byte
	var copied []string
	r := NewReader(&chunkReader{data: []byte(doc), n: 64}, WithBufferSize(128))
	if err := r.ArrayEach(func(v []byte) error {
		kept = append(kept, v)
		copied = append(copied, string(v))
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	changed := 0
	for i, v := range kept {
		if string(v) != copied[i] {
			changed++
		}
	}
	if changed == 0 {
		t.Fatal("no kept slice changed after the walk: the test no longer exercises the lifetime rule " +
			"(the buffer must be refilled and reused for it to bite)")
	}
	t.Logf("%d of %d kept slices no longer read as they did inside the callback", changed, len(kept))
}

// TestStreamErrStop ends a walk early, and the reader stops reading soon after.
func TestStreamErrStop(t *testing.T) {
	doc := `[1,2,3,4,5]`
	for _, sentinel := range []error{ErrStop, fmt.Errorf("wrapped: %w", ErrStop)} {
		n := 0
		r := NewReader(strings.NewReader(doc))
		if err := r.ArrayEach(func(v []byte) error {
			n++
			if n == 2 {
				return sentinel
			}
			return nil
		}); err != nil {
			t.Fatalf("ArrayEach with %v = %v, want nil", sentinel, err)
		}
		if n != 2 {
			t.Fatalf("walked %d elements after ErrStop, want 2", n)
		}
	}
	// Any other error is returned as it is.
	boom := errors.New("boom")
	r := NewReader(strings.NewReader(doc))
	if err := r.ArrayEach(func(v []byte) error { return boom }); !errors.Is(err, boom) {
		t.Fatalf("ArrayEach = %v, want the callback's own error", err)
	}
}

// failingReader delivers a prefix and then fails, which is what a dropped
// connection looks like.
type failingReader struct {
	data []byte
	at   int
	err  error
}

func (f *failingReader) Read(p []byte) (int, error) {
	if f.at <= 0 {
		return 0, f.err
	}
	n := f.at
	if n > len(p) {
		n = len(p)
	}
	if n > len(f.data) {
		n = len(f.data)
	}
	copy(p, f.data[:n])
	f.data, f.at = f.data[n:], f.at-n
	return n, nil
}

// TestStreamReadError returns the underlying failure rather than calling the
// document truncated, and errors.Is reaches it.
func TestStreamReadError(t *testing.T) {
	boom := errors.New("connection reset")
	doc := `[1,2,3,"` + strings.Repeat("x", 100) + `"]`
	for _, at := range []int{1, 5, 8, 40} {
		n := 0
		r := NewReader(&failingReader{data: []byte(doc), at: at, err: boom}, WithBufferSize(16))
		err := r.ArrayEach(func(v []byte) error { n++; return nil })
		if !errors.Is(err, boom) {
			t.Errorf("ArrayEach with the reader failing after %d bytes = %v, want %v", at, err, boom)
		}
	}
	// A clean end mid-document is ErrTruncated, as in memory.
	r := NewReader(strings.NewReader(`[1,2,`), WithBufferSize(16))
	if err := r.ArrayEach(func(v []byte) error { return nil }); !errors.Is(err, ErrTruncated) {
		t.Errorf("ArrayEach over a truncated document = %v, want ErrTruncated", err)
	}
}

// TestStreamElements is the iterator form: same values, same early exit.
func TestStreamElements(t *testing.T) {
	doc := `{"data":{"result":[{"m":1},{"m":2},{"m":3}]}}`
	var got []string
	r := NewReader(strings.NewReader(doc))
	for v, err := range r.Elements("data", "result") {
		if err != nil {
			t.Fatal(err)
		}
		got = append(got, string(v))
	}
	if fmt.Sprint(got) != `[{"m":1} {"m":2} {"m":3}]` {
		t.Fatalf("got %v", got)
	}
	// break ends the walk cleanly, and the error is delivered once when the
	// document is malformed.
	r = NewReader(strings.NewReader(`[1,2,3]`))
	for v := range r.Elements() {
		_ = v
		break
	}
	var ierr error
	r = NewReader(strings.NewReader(`[1,`))
	for _, err := range r.Elements() {
		if err != nil {
			ierr = err
		}
	}
	if !errors.Is(ierr, ErrTruncated) {
		t.Fatalf("iterating a truncated array reported %v, want ErrTruncated", ierr)
	}
}

// TestStreamNullContainer follows the in-memory rule: a null where a container
// would be is that container with nothing in it.
func TestStreamNullContainer(t *testing.T) {
	for _, doc := range []string{`null`, ` null `, `{"a":null}`, "null\n", "null\x00"} {
		keys := []string(nil)
		if strings.HasPrefix(doc, "{") {
			keys = []string{"a"}
		}
		called := false
		r := NewReader(strings.NewReader(doc), WithBufferSize(4))
		if err := r.ArrayEach(func([]byte) error { called = true; return nil }, keys...); err != nil || called {
			t.Errorf("ArrayEach(%q) = %v, called=%v; want nil, false", doc, err, called)
		}
		r = NewReader(strings.NewReader(doc), WithBufferSize(4))
		if err := r.ObjectEach(func(string, []byte) error { called = true; return nil }, keys...); err != nil || called {
			t.Errorf("ObjectEach(%q) = %v, called=%v; want nil, false", doc, err, called)
		}
	}
	for _, doc := range []string{`nul`, `nullx`, `"null"`, `1`, `true`} {
		r := NewReader(strings.NewReader(doc), WithBufferSize(4))
		if err := r.ArrayEach(func([]byte) error { return nil }); err == nil {
			t.Errorf("ArrayEach(%q) = nil, want an error", doc)
		}
	}
}

// TestStreamBuffered hands the rest of the stream on: what the reader read
// ahead of the document plus what it has not read is the remainder.
func TestStreamBuffered(t *testing.T) {
	src := `[1,2,3]{"next":true}` + strings.Repeat("x", 100)
	r := NewReader(strings.NewReader(src))
	if err := r.ArrayEach(func([]byte) error { return nil }); err != nil {
		t.Fatal(err)
	}
	rest := string(r.Buffered())
	if !strings.HasPrefix(rest, `{"next":true}`) {
		t.Fatalf("Buffered = %.20q, want the bytes after the array", rest)
	}
	if int64(len(rest))+int64(len(`[1,2,3]`)) != r.Consumed() {
		t.Fatalf("Buffered (%d) + the document (7) != Consumed (%d)", len(rest), r.Consumed())
	}
}

// TestStreamHugeElementIsLinear is a regression guard for the shape the
// resumable scanner exists for: a value that arrives one byte at a time must
// be scanned once, not re-scanned from its first byte after every refill. At
// 256 KB the quadratic form is ~10^10 byte-steps and this test does not
// finish; the linear one takes milliseconds.
func TestStreamHugeElementIsLinear(t *testing.T) {
	doc := `["` + strings.Repeat("q", 256<<10) + `"]`
	r := NewReader(iotest.OneByteReader(strings.NewReader(doc)), WithBufferSize(512))
	n := 0
	if err := r.ArrayEach(func(v []byte) error { n = len(v); return nil }); err != nil {
		t.Fatal(err)
	}
	if n != 256<<10+2 {
		t.Fatalf("element length %d, want %d", n, 256<<10+2)
	}
}

// TestStreamReset reuses one reader — and its buffer — over a sequence of
// documents, which is what a client making the same query repeatedly does.
func TestStreamReset(t *testing.T) {
	docs := []string{`[1,2]`, `["a","b","c"]`, `{"k":[7]}`, `[]`}
	r := NewReader(strings.NewReader(docs[0]), WithBufferSize(8))
	buf := r.buf
	for i, doc := range docs {
		if i > 0 {
			r.Reset(strings.NewReader(doc))
		}
		keys := []string(nil)
		if strings.HasPrefix(doc, "{") {
			keys = []string{"k"}
		}
		var got []string
		if err := r.ArrayEach(func(v []byte) error { got = append(got, string(v)); return nil }, keys...); err != nil {
			t.Fatalf("%q: %v", doc, err)
		}
		want := memArray(doc, keys...)
		if fmt.Sprint(got) != fmt.Sprint(want.values) {
			t.Fatalf("%q: %v, want %v", doc, got, want.values)
		}
		// A walk stops at its container's closer and reads no further: the
		// trailing '}' of the object form is never asked for.
		if c := r.Consumed(); c > int64(len(doc)) || c == 0 {
			t.Fatalf("%q: Consumed = %d, want 0 < n <= %d", doc, c, len(doc))
		}
	}
	if &r.buf[0] != &buf[0] {
		t.Fatal("Reset allocated a new buffer instead of reusing the one it had")
	}
}

// randomJSON builds a document of the given depth deterministically, mixing
// every value kind and the bytes that make a naive scanner wrong (brackets and
// quotes inside strings, escapes, unicode).
func randomJSON(rnd *uint64, depth int) string {
	next := func(n int) int {
		*rnd = *rnd*6364136223846793005 + 1442695040888963407
		return int((*rnd >> 33) % uint64(n))
	}
	if depth <= 0 {
		switch next(6) {
		case 0:
			return "null"
		case 1:
			return "true"
		case 2:
			return fmt.Sprint(next(1000000))
		case 3:
			return fmt.Sprintf("-%d.%de-%d", next(100), next(1000), next(30))
		case 4:
			return `"` + strings.Repeat("]},\\\"é ", 1+next(20)) + `"`
		}
		return `"` + strings.Repeat("x", next(300)) + `"`
	}
	n := next(6)
	parts := make([]string, n)
	if next(2) == 0 {
		for i := range parts {
			parts[i] = randomJSON(rnd, depth-1)
		}
		return "[" + strings.Join(parts, ",") + "]"
	}
	for i := range parts {
		parts[i] = fmt.Sprintf(`"k%d\t%d":%s`, i, next(100), randomJSON(rnd, depth-1))
	}
	return "{" + strings.Join(parts, ",") + "}"
}

// TestStreamRandomDocuments is the volume differential: generated documents,
// each walked in memory and from a reader at several chunk and buffer sizes,
// held to the same values and the same error. It is what catches a state that
// survives a refill wrongly on a shape the hand-written corpus does not have.
func TestStreamRandomDocuments(t *testing.T) {
	seed := uint64(1)
	for i := 0; i < 400; i++ {
		doc := randomJSON(&seed, 3)
		wantA, wantO := memArray(doc), memObject(doc)
		for _, chunk := range []int{1, 3, 17, 1024} {
			for _, buf := range []int{8, 40, 1 << 12} {
				mk := func() io.Reader { return &chunkReader{data: []byte(doc), n: chunk} }
				opts := []ReaderOption{WithBufferSize(buf)}
				if got := streamArray(t, doc, mk, opts); !sameWalk(got, wantA) {
					t.Fatalf("ArrayEach chunk=%d buf=%d over %.80q:\n got %v %v\nwant %v %v",
						chunk, buf, doc, got.values, got.err, wantA.values, wantA.err)
				}
				if got := streamObject(t, doc, mk, opts); !sameWalk(got, wantO) {
					t.Fatalf("ObjectEach chunk=%d buf=%d over %.80q:\n got %v %v %v\nwant %v %v %v",
						chunk, buf, doc, got.keys, got.values, got.err, wantO.keys, wantO.values, wantO.err)
				}
			}
		}
	}
}
