package unstable

import (
	"errors"
	"strings"
	"testing"
)

// This file locks the two *bounds* the skip and presize machinery must respect
// on hostile input: the recursion bound on the scalar container skip (a stack
// overflow is fatal — recover cannot catch it) and the structural bound on the
// three presize counters (their result becomes a make() capacity, so an
// unbounded count is an unbounded allocation).

// nestedObjectDoc returns a well-formed document whose deepest container nests
// exactly depth levels: depth-1 wrapper objects around a leaf {"b":1}.
func nestedObjectDoc(depth int) []byte {
	var b strings.Builder
	b.Grow(depth*6 + 8)
	for i := 1; i < depth; i++ {
		b.WriteString(`{"a":`)
	}
	b.WriteString(`{"b":1}`)
	b.WriteString(strings.Repeat("}", depth-1))
	return []byte(b.String())
}

// nestedScalarFirstArrayDoc returns the shape from the stack-overflow repro:
// [0,[[[…]]]] nested exactly depth levels. The leading scalar element is
// load-bearing — SkipValue's fast-path probe only routes an array to the
// iterative SIMD scan when its first element is '{', '[' or '"', so a leading 0
// forces the recursive scalar skipArray on every machine.
func nestedScalarFirstArrayDoc(depth int) []byte {
	n := depth - 1
	return []byte("[0," + strings.Repeat("[", n) + "1" + strings.Repeat("]", n) + "]")
}

// TestSkipDepthBound checks that the scalar container skip refuses to descend
// past MaxDepth instead of recursing until the goroutine stack dies. Before the
// bound, skipObject/skipArray recursed unconditionally: a 20 MB document of the
// nestedScalarFirstArrayDoc shape reached "goroutine stack exceeds
// 1000000000-byte limit / fatal error: stack overflow" and killed the process
// from json.Get. The test deliberately probes at MaxDepth+1 rather than at the
// depth that actually overflows — a test that needed 20 MB of brackets to prove
// the point would be slower than the whole package suite.
func TestSkipDepthBound(t *testing.T) {
	// The array shape reaches the recursive path through the public SkipValue on
	// every machine (see nestedScalarFirstArrayDoc).
	t.Run("array/SkipValue", func(t *testing.T) {
		doc := nestedScalarFirstArrayDoc(MaxDepth)
		end, err := SkipValue(doc, 0)
		if err != nil || end != len(doc) {
			t.Fatalf("at MaxDepth: SkipValue = (%d, %v), want (%d, nil)", end, err, len(doc))
		}
		doc = nestedScalarFirstArrayDoc(MaxDepth + 1)
		if _, err := SkipValue(doc, 0); !errors.Is(err, ErrMaxDepth) {
			t.Fatalf("at MaxDepth+1: SkipValue err = %v, want ErrMaxDepth", err)
		}
	})

	// The object shape is dispatched to the iterative SIMD skip wherever that is
	// available, so the recursive path is exercised directly. On a machine
	// without the SIMD skip (and on the generated decoders' unknown-field skip
	// there) SkipValue itself lands here, which the second half asserts.
	t.Run("object/scalarPath", func(t *testing.T) {
		doc := nestedObjectDoc(MaxDepth)
		end, err := skipObject(doc, 0)
		if err != nil || end != len(doc) {
			t.Fatalf("at MaxDepth: skipObject = (%d, %v), want (%d, nil)", end, err, len(doc))
		}
		doc = nestedObjectDoc(MaxDepth + 1)
		if _, err := skipObject(doc, 0); !errors.Is(err, ErrMaxDepth) {
			t.Fatalf("at MaxDepth+1: skipObject err = %v, want ErrMaxDepth", err)
		}
		if !fastSkipAvail {
			if _, err := SkipValue(doc, 0); !errors.Is(err, ErrMaxDepth) {
				t.Fatalf("at MaxDepth+1: SkipValue err = %v, want ErrMaxDepth", err)
			}
		}
	})

	// Pure nested arrays through the scalar entry point, for the same off-by-one
	// check on the other half of the mutual recursion.
	t.Run("array/scalarPath", func(t *testing.T) {
		ok := []byte(strings.Repeat("[", MaxDepth) + "1" + strings.Repeat("]", MaxDepth))
		if end, err := skipArray(ok, 0); err != nil || end != len(ok) {
			t.Fatalf("at MaxDepth: skipArray = (%d, %v), want (%d, nil)", end, err, len(ok))
		}
		deep := []byte(strings.Repeat("[", MaxDepth+1) + "1" + strings.Repeat("]", MaxDepth+1))
		if _, err := skipArray(deep, 0); !errors.Is(err, ErrMaxDepth) {
			t.Fatalf("at MaxDepth+1: skipArray err = %v, want ErrMaxDepth", err)
		}
	})

	// The repro's full document: the deep array sits in a member of an outer
	// object, which is how Get/GetPaths and unknown-field skipping reach it.
	t.Run("repro/embedded", func(t *testing.T) {
		inner := nestedScalarFirstArrayDoc(MaxDepth + 1)
		doc := append([]byte(`{"k":1,"a":`), inner...)
		doc = append(doc, `,"zz":7}`...)
		// Skipping member "a"'s value is the operation that used to recurse.
		if _, err := SkipValue(doc, len(`{"k":1,"a":`)); !errors.Is(err, ErrMaxDepth) {
			t.Fatalf("embedded: SkipValue err = %v, want ErrMaxDepth", err)
		}
	})
}

// TestSkipDepthBoundKeepsShallowDocs is the regression guard for the bound
// itself: ordinary documents must be unaffected, including the stray-bracket and
// truncation branches skip_test.go covers.
func TestSkipDepthBoundKeepsShallowDocs(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{`{"a":{"b":{"c":[1,[2,[3]]]}}}`, 29},
		{`[[1],[2,[3]]]`, 13},
		{`{"a":"}]{["}`, 12},
	}
	for _, c := range cases {
		end, err := SkipValue([]byte(c.in), 0)
		if err != nil || end != c.want {
			t.Errorf("SkipValue(%q) = (%d, %v), want (%d, nil)", c.in, end, err, c.want)
		}
	}
}

// --- S2: the presize counters are hints, but they are also make() capacities ---

// TestCountHintsClampedToSpan pins the structural clamp on all three counters.
// Each returns a count derived from bytes inside the array (commas, braces, or a
// sampled density), and any of those can be inflated arbitrarily by content
// *inside a string value* — which the counters deliberately do not parse. That
// is harmless for correctness (a wrong count only mis-sizes) but not for memory:
// the result is handed straight to make(), so an unclamped count multiplies the
// input size by sizeof(element). Measured before the clamp, with a 64-byte
// element struct: a 2 MB body allocated 130 MB, linear in input.
//
// The clamp is the byte span the counter already computed: n elements need n-1
// separating commas plus at least one byte each (two for an object), so a span
// simply cannot hold more elements than that, whatever its bytes say.
func TestCountHintsClampedToSpan(t *testing.T) {
	const n = 4096

	t.Run("objects", func(t *testing.T) {
		// Braces inside a string value: the count would be n+1.
		doc := []byte(`[{"s":"` + strings.Repeat("{", n) + `","t":"x"}]`)
		span := len(doc) - 2 // content between '[' and the closing ']'
		lim := (span + 1) / 3
		got := CountArrayObjects(doc, 0)
		if got > lim {
			t.Fatalf("CountArrayObjects = %d, want <= %d (span %d)", got, lim, span)
		}
		if got >= n {
			t.Fatalf("CountArrayObjects = %d: clamp did not fire (unclamped would be %d)", got, n+1)
		}
		if got < 1 {
			t.Fatalf("CountArrayObjects = %d, want a usable hint >= 1", got)
		}
	})

	t.Run("scalars", func(t *testing.T) {
		// Commas inside a string value — reachable because []time.Time counts as
		// a "scalar" element (an RFC 3339 value holds no comma, but a hostile
		// document is not obliged to be one).
		doc := []byte(`["a` + strings.Repeat(",", n) + `"]`)
		span := len(doc) - 2
		lim := (span + 1) / 2
		got := CountArrayScalars(doc, 0)
		if got > lim {
			t.Fatalf("CountArrayScalars = %d, want <= %d (span %d)", got, lim, span)
		}
		if got >= n {
			t.Fatalf("CountArrayScalars = %d: clamp did not fire (unclamped would be %d)", got, n+1)
		}
	})

	t.Run("elements", func(t *testing.T) {
		// The sample-and-extrapolate path: 64 one-byte elements set the density,
		// then one huge element stretches the span the density is applied to.
		doc := []byte("[" + strings.Repeat("1,", countSampleCap) + `"` + strings.Repeat("x", 100000) + `"]`)
		span := len(doc) - 1 // '[' .. ']'
		lim := span / 2
		got := CountArrayElements(doc, 0)
		if got > lim {
			t.Fatalf("CountArrayElements = %d, want <= %d (span %d)", got, lim, span)
		}
	})
}

// TestCountHintsHonestCountsUnchanged proves the clamps cannot bite a real
// count: they are applied at the densest legal packing, where every element is
// as small as JSON allows, so any honest array is strictly under the bound.
func TestCountHintsHonestCountsUnchanged(t *testing.T) {
	for n := 1; n <= 64; n++ {
		// Densest object array: [{},{},…] — two bytes per element plus commas.
		objs := "[" + strings.TrimSuffix(strings.Repeat("{},", n), ",") + "]"
		if got := CountArrayObjects([]byte(objs), 0); got != n {
			t.Fatalf("CountArrayObjects(%d elems) = %d, want %d", n, got, n)
		}
		// Densest scalar array: [1,1,…] — one byte per element plus commas.
		scal := "[" + strings.TrimSuffix(strings.Repeat("1,", n), ",") + "]"
		if got := CountArrayScalars([]byte(scal), 0); got != n {
			t.Fatalf("CountArrayScalars(%d elems) = %d, want %d", n, got, n)
		}
		if got := CountArrayElements([]byte(scal), 0); got != n {
			t.Fatalf("CountArrayElements(%d elems) = %d, want %d", n, got, n)
		}
	}
	// Past the sample cap the element counter extrapolates; on the densest legal
	// array the clamp is exactly the true count, so it must not under-size.
	const big = 1000
	dense := "[" + strings.TrimSuffix(strings.Repeat("1,", big), ",") + "]"
	if got := CountArrayElements([]byte(dense), 0); got < big {
		t.Fatalf("CountArrayElements(%d dense elems) = %d, want >= %d", big, got, big)
	}
}
