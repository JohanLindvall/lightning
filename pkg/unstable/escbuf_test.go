package unstable

import (
	"strings"
	"sync"
	"testing"
)

// BenchmarkEscapeScratch is the number the escbuf.go comment quotes: what one
// escaped string's backing costs carved from a chunk versus made on its own.
func BenchmarkEscapeScratch(b *testing.B) {
	doc := make([]byte, 64<<10)
	b.Run("carve", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			buf, chunk, off := escapeScratch(doc, 64)
			escapeRelease(chunk, off, 64, 64)
			escapeSink = buf
		}
	})
	b.Run("make", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			escapeSink = make([]byte, 0, 64)
		}
	})
}

var escapeSink []byte

// TestEscapeScratchExclusive is the property the whole design rests on: two
// carves never overlap, so one decoded string can never be written over by the
// next one, and a carve's capacity stops at its own end so a caller's append
// reallocates instead of reaching its neighbour.
//
// The check is on CONTENTS, not addresses: every buffer is filled with a mark
// unique to its carve and kept alive to the end, so an overlap shows up as a
// wrong byte. Comparing addresses instead is wrong and looks right — a buffer
// the test does not retain is garbage, and a later chunk allocated at the same
// address then reads as an "overlap" that never happened.
func TestEscapeScratchExclusive(t *testing.T) {
	doc := make([]byte, 32<<10)
	const carves = 4000
	held := make([][]byte, 0, carves)
	for i := 0; i < carves; i++ {
		n := 1 + i%300
		buf, chunk, off := escapeScratch(doc, n)
		if cap(buf) != n {
			t.Fatalf("carve %d: cap = %d, want %d", i, cap(buf), n)
		}
		if len(buf) != 0 {
			t.Fatalf("carve %d: len = %d, want 0", i, len(buf))
		}
		for j := 0; j < n; j++ {
			buf = append(buf, byte(i))
		}
		if cap(buf) != n {
			t.Fatalf("carve %d: filling it grew cap to %d, so it reached past its own end", i, cap(buf))
		}
		escapeRelease(chunk, off, n, n)
		held = append(held, buf)
	}
	for i, b := range held {
		for j, c := range b {
			if c != byte(i) {
				t.Fatalf("carve %d byte %d = %d, want %d: a later carve overlapped it", i, j, c, byte(i))
			}
		}
	}
}

// TestEscapeScratchChunkBound pins the retention rule: a chunk is never larger
// than the document it was made for, so decoding a small document cannot leave
// a string pinning a large one.
func TestEscapeScratchChunkBound(t *testing.T) {
	for _, docLen := range []int{16, 300, 4000, 100000} {
		doc := make([]byte, docLen)
		// Drain whatever chunk the pool holds so this document's size decides.
		for i := 0; i < 64; i++ {
			buf, chunk, off := escapeScratch(doc, escapeMaxCarve)
			escapeRelease(chunk, off, escapeMaxCarve, escapeMaxCarve)
			escapeSink = buf
		}
		n := 8
		if n > docLen {
			n = docLen
		}
		if n == 0 {
			continue
		}
		_, chunk, off := escapeScratch(doc, n)
		if chunk == nil {
			t.Fatalf("docLen %d: no chunk", docLen)
		}
		want := docLen
		if want > escapeChunkMax {
			want = escapeChunkMax
		} else if want < escapeChunkMin {
			want = escapeChunkMin
		}
		if cap(chunk.buf) > want {
			t.Errorf("docLen %d: chunk cap %d exceeds the document bound %d", docLen, cap(chunk.buf), want)
		}
		escapeRelease(chunk, off, n, 0)
	}
}

// TestEscapedStringsSurviveChunkReuse decodes many escaped strings out of one
// document and checks every one of them still reads correctly afterwards — the
// end-to-end form of the exclusivity property, through the real reader.
func TestEscapedStringsSurviveChunkReuse(t *testing.T) {
	var sb strings.Builder
	var want []string
	for i := 0; i < 2000; i++ {
		body := strings.Repeat("a\\\"b", 1+i%40)
		sb.WriteString(`"` + body + `"`)
		want = append(want, strings.ReplaceAll(body, `\"`, `"`))
	}
	doc := []byte(sb.String())
	var got []string
	for i := 0; i < len(doc); {
		s, end, err := ReadStringOrNull(doc, i)
		if err != nil {
			t.Fatalf("at %d: %v", i, err)
		}
		got = append(got, s)
		i = end
	}
	if len(got) != len(want) {
		t.Fatalf("got %d strings, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("string %d = %q, want %q", i, got[i], want[i])
		}
	}
}

// TestEscapeScratchConcurrent is why the chunk lives in a sync.Pool and not in a
// package variable: run under -race, two goroutines carving at once out of a
// package variable report a data race, and two carving out of the pool must not.
// It also re-checks the values, since a lost bump would hand the same bytes to
// two decodes and the corruption would be silent.
func TestEscapeScratchConcurrent(t *testing.T) {
	const goroutines, each = 8, 4000
	doc := make([]byte, 32<<10)
	var wg sync.WaitGroup
	errs := make(chan string, goroutines)
	for g := 0; g < goroutines; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			held := make([][]byte, 0, each)
			for i := 0; i < each; i++ {
				n := 1 + (g*7+i)%200
				buf, chunk, off := escapeScratch(doc, n)
				for j := 0; j < n; j++ {
					buf = append(buf, byte(g))
				}
				escapeRelease(chunk, off, n, n)
				held = append(held, buf)
			}
			for i, b := range held {
				for j, c := range b {
					if c != byte(g) {
						errs <- "goroutine's carve was overwritten"
						_ = i
						_ = j
						return
					}
				}
			}
		}(g)
	}
	wg.Wait()
	close(errs)
	for e := range errs {
		t.Fatal(e)
	}
}
