package json

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"testing"
)

// ArrayEachIndex is ArrayEach with a counter, written out rather than
// shared so ArrayEach's per-element call stays one indirect call — which
// makes the two walkers two copies of one loop, held to the same spans and
// the same errors here over hand-picked documents and a generated corpus.
func TestArrayEachIndexMatchesArrayEach(t *testing.T) {
	docs := []string{
		`[]`, `[1]`, `[1,2,3]`, `[ "a" , {"b": [1,2]} , null , true ]`,
		`[[[[1]]]]`, `["\"]",", ["]`, `[1.5e10,-0]`, `[{}]`, `[""]`, `[0]`,
		`{"records":[{"x":"y"},2]}`, `{"a":{"b":[[1,"2"],[3,"4"]]}}`,
		// Malformed: both must fail the same way at the same point.
		`[1,]`, `[`, `[1`, `[1,`, `[{"a":1}`, `{"records":5}`, `[1 2]`, `5`,
	}
	rnd := rand.New(rand.NewSource(3))
	for range 500 {
		n := rnd.Intn(6)
		elems := make([]string, n)
		for i := range elems {
			elems[i] = []string{`1`, `"s"`, `null`, `true`, `[1,[2]]`, `{"k":"v"}`, `-2.5e3`}[rnd.Intn(7)]
		}
		sep := []string{",", " , ", ",\n\t"}[rnd.Intn(3)]
		docs = append(docs, "["+strings.Join(elems, sep)+"]")
	}
	for _, doc := range docs {
		var keys []string
		if strings.HasPrefix(doc, `{"records"`) {
			keys = []string{"records"}
		} else if strings.HasPrefix(doc, `{"a"`) {
			keys = []string{"a", "b"}
		}
		for _, compact := range []bool{false, true} {
			if compact && strings.ContainsAny(doc, " \n\t") {
				continue
			}
			var plain, indexed []string
			var idx []int
			var e1, e2 error
			if compact {
				e1 = ArrayEachCompact([]byte(doc), func(v []byte) error { plain = append(plain, string(v)); return nil }, keys...)
				e2 = ArrayEachIndexCompact([]byte(doc), func(i int, v []byte) error { idx = append(idx, i); indexed = append(indexed, string(v)); return nil }, keys...)
			} else {
				e1 = ArrayEach([]byte(doc), func(v []byte) error { plain = append(plain, string(v)); return nil }, keys...)
				e2 = ArrayEachIndex([]byte(doc), func(i int, v []byte) error { idx = append(idx, i); indexed = append(indexed, string(v)); return nil }, keys...)
			}
			if !reflect.DeepEqual(plain, indexed) || (e1 == nil) != (e2 == nil) || (e1 != nil && !errors.Is(e2, e1)) {
				t.Fatalf("%q compact=%v: ArrayEach %q/%v, ArrayEachIndex %q/%v", doc, compact, plain, e1, indexed, e2)
			}
			for i, got := range idx {
				if got != i {
					t.Fatalf("%q: element %d was handed index %d", doc, i, got)
				}
			}
		}
	}
}

// The tuple case the index exists for: a [timestamp, value] pair read by
// position, no counter in the closure.
func TestArrayEachIndexReadsATupleByPosition(t *testing.T) {
	var ts, val []byte
	err := ArrayEachIndex([]byte(`[1788087600.5,"0.25"]`), func(i int, v []byte) error {
		switch i {
		case 0:
			ts = v
		case 1:
			val = v
		default:
			return fmt.Errorf("a pair has no element %d", i)
		}
		return nil
	})
	if err != nil || string(ts) != "1788087600.5" || string(val) != `"0.25"` {
		t.Fatalf("ts %q val %q err %v", ts, val, err)
	}
}

// ErrStop ends a walk with a nil result on every walker, plain and compact,
// wrapped or not; any other error is still returned as it is.
func TestErrStopEndsAWalkWithoutAnError(t *testing.T) {
	doc := []byte(`{"a":[10,20,30,40],"b":1,"c":2}`)
	wrapped := fmt.Errorf("found it: %w", ErrStop)

	n := 0
	if err := ArrayEach(doc, func([]byte) error { n++; return ErrStop }, "a"); err != nil || n != 1 {
		t.Fatalf("ArrayEach: err %v, calls %d", err, n)
	}
	n = 0
	if err := ArrayEachCompact(doc, func([]byte) error { n++; return wrapped }, "a"); err != nil || n != 1 {
		t.Fatalf("ArrayEachCompact with a wrapped ErrStop: err %v, calls %d", err, n)
	}
	var second []byte
	err := ArrayEachIndex(doc, func(i int, v []byte) error {
		if i == 1 {
			second = v
			return ErrStop
		}
		return nil
	}, "a")
	if err != nil || string(second) != "20" {
		t.Fatalf("ArrayEachIndex: err %v, second %q", err, second)
	}
	n = 0
	if err := ObjectEach(doc, func(string, []byte) error { n++; return ErrStop }); err != nil || n != 1 {
		t.Fatalf("ObjectEach: err %v, calls %d", err, n)
	}
	n = 0
	if err := ObjectEachCompact(doc, func(string, []byte) error { n++; return ErrStop }); err != nil || n != 1 {
		t.Fatalf("ObjectEachCompact: err %v, calls %d", err, n)
	}

	// Any other error is a failure, as before.
	boom := errors.New("boom")
	if err := ArrayEach(doc, func([]byte) error { return boom }, "a"); !errors.Is(err, boom) {
		t.Fatalf("a real error must be returned, got %v", err)
	}
	if err := ObjectEach(doc, func(string, []byte) error { return boom }); !errors.Is(err, boom) {
		t.Fatalf("a real error must be returned, got %v", err)
	}
	// And a stop in the last element is indistinguishable from finishing.
	if err := ArrayEach([]byte(`[1]`), func([]byte) error { return ErrStop }); err != nil {
		t.Fatalf("stopping on the last element: %v", err)
	}
}

func BenchmarkArrayEachIndex(b *testing.B) {
	doc := []byte(`[[1788087600,"0.1"],[1788087660,"0.2"],[1788087720,"0.3"],[1788087780,"0.4"]]`)
	b.ReportAllocs()
	for range b.N {
		err := ArrayEachIndex(doc, func(_ int, pt []byte) error {
			return ArrayEachIndex(pt, func(int, []byte) error { return nil })
		})
		if err != nil {
			b.Fatal(err)
		}
	}
}
