package json

import (
	"errors"
	"testing"
)

const envelope = `{"status":"success","data":{"resultType":"matrix","result":[{"m":1},{"m":2},{"m":3}],"stats":{"rows":42}},"warnings":["w1","w2"],"infos":[]}`

// A later call on a Reader continues from where the previous one stopped:
// the envelope's members are read in document order, the array walked in the
// middle of it, and the members after it — inside data and after data —
// reached without a second reader. Every answer is the in-memory function's.
func TestStreamContinuesToTheNextPath(t *testing.T) {
	for name, src := range readerKinds(envelope) {
		t.Run(name, func(t *testing.T) {
			r := NewReader(src(), WithBufferSize(16))
			want := func(keys ...string) string {
				v, _, err := Get([]byte(envelope), keys...)
				if err != nil {
					t.Fatal(err)
				}
				return string(v)
			}
			get := func(keys ...string) string {
				v, err := r.Get(keys...)
				if err != nil {
					t.Fatalf("Get %v: %v", keys, err)
				}
				return string(v)
			}
			if got := get("status"); got != want("status") {
				t.Fatalf("status %q", got)
			}
			if got := get("data", "resultType"); got != want("data", "resultType") {
				t.Fatalf("resultType %q", got)
			}
			var elems []string
			if err := r.ArrayEach(func(v []byte) error { elems = append(elems, string(v)); return nil }, "data", "result"); err != nil {
				t.Fatal(err)
			}
			if len(elems) != 3 || elems[2] != `{"m":3}` {
				t.Fatalf("elements %q", elems)
			}
			if got := get("data", "stats", "rows"); got != "42" {
				t.Fatalf("stats.rows %q", got)
			}
			if got := get("warnings"); got != want("warnings") {
				t.Fatalf("warnings %q", got)
			}
			if got := get("infos"); got != "[]" {
				t.Fatalf("infos %q", got)
			}
			if _, err := r.Get("nothing"); !errors.Is(err, ErrKeyNotFound) {
				t.Fatalf("past the last member: %v", err)
			}
		})
	}
}

// Forward only: a member the walk has passed cannot be reached again, and
// neither can the inside of a value already read whole.
func TestStreamContinuationIsForwardOnly(t *testing.T) {
	r := NewReader(readerKinds(envelope)["chunk3"]())
	if _, err := r.Get("warnings"); err != nil {
		t.Fatal(err)
	}
	if _, err := r.Get("status"); !errors.Is(err, ErrKeyNotFound) {
		t.Fatalf("a key behind the cursor: %v", err)
	}
	r = NewReader(readerKinds(envelope)["chunk3"]())
	if _, err := r.Get("data"); err != nil {
		t.Fatal(err)
	}
	if _, err := r.Get("data", "stats"); !errors.Is(err, ErrKeyNotFound) {
		t.Fatalf("inside a value already read: %v", err)
	}
	// The miss ran the root to its brace looking for a second "data", so
	// the members it passed on the way are behind the cursor too.
	if _, err := r.Get("warnings"); !errors.Is(err, ErrKeyNotFound) {
		t.Fatalf("a miss spends the object it searched: %v", err)
	}
}

// A walk a callback ended with ErrStop is finished — the rest of its
// container skipped — before the next path is looked for, from an array and
// from an object alike; and a root-level walk spends the document.
func TestStreamContinuesPastAStoppedWalk(t *testing.T) {
	for name, src := range readerKinds(envelope) {
		t.Run(name, func(t *testing.T) {
			r := NewReader(src(), WithBufferSize(16))
			n := 0
			if err := r.ArrayEach(func([]byte) error { n++; return ErrStop }, "data", "result"); err != nil || n != 1 {
				t.Fatalf("stop: n=%d err=%v", n, err)
			}
			if v, err := r.Get("data", "stats"); err != nil || string(v) != `{"rows":42}` {
				t.Fatalf("after a stopped array: %q %v", v, err)
			}
			if v, err := r.Get("warnings"); err != nil || string(v) != `["w1","w2"]` {
				t.Fatalf("after leaving data: %q %v", v, err)
			}
		})
	}
	r := NewReader(readerKinds(envelope)["chunk7"](), WithBufferSize(16))
	keys := 0
	if err := r.ObjectEach(func(string, []byte) error { keys++; return ErrStop }, "data"); err != nil || keys != 1 {
		t.Fatalf("object stop: %d %v", keys, err)
	}
	if v, err := r.Get("warnings"); err != nil || string(v) != `["w1","w2"]` {
		t.Fatalf("after a stopped object: %q %v", v, err)
	}
	r = NewReader(readerKinds(`[1,2,3]`)["onebyte"]())
	if err := r.ArrayEach(func([]byte) error { return nil }); err != nil {
		t.Fatal(err)
	}
	if _, err := r.Get("x"); !errors.Is(err, ErrKeyNotFound) {
		t.Fatalf("after a root walk: %v", err)
	}
	r.Reset(readerKinds(envelope)["chunk7"]())
	if v, err := r.Get("status"); err != nil || string(v) != `"success"` {
		t.Fatalf("Reset starts over: %q %v", v, err)
	}
}
