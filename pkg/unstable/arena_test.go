package unstable

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"unsafe"
)

// TestArenaSliceParity locks the ...Arena readers to their plain counterparts:
// for every input — well-formed, empty, null, whitespace-ridden, trailing-comma,
// truncated, malformed — the decoded slice contents, the returned index, the
// error identity, and the partial progress left in *out on error must be
// identical. Only the backing's provenance may differ.
func TestArenaSliceParity(t *testing.T) {
	inputs := []string{
		`[1,2,3]`,
		`[1.5, -2.25e3, 0.0006988752666567719]`,
		`[ ]`,
		`[]`,
		`null`,
		`[null,4,null]`,
		"[\n  1,\n  2\n]",
		`[1,2,]`,                    // trailing comma
		`[1,2`,                      // truncated
		`[1,"x",3]`,                 // malformed element; partial progress must match
		`[999999999999999999999,1]`, // overflow wrap parity
		`x`,
		``,
	}
	// A long array crosses the arenaMaxCarve threshold (direct make fallback).
	long := "[0"
	for i := 1; i < 200; i++ {
		long += fmt.Sprintf(",%d", i)
	}
	inputs = append(inputs, long+"]")

	for _, in := range inputs {
		t.Run(in[:min(len(in), 24)], func(t *testing.T) {
			var pf, af []float64
			// An arena is typed, so a document mixing element kinds threads one
			// per kind — which is what the generated decoders declare.
			af64, ai32, au16 := NewArena[float64](), NewArena[int32](), NewArena[uint16]()
			pe, perr := DecodeFloat64Slice(&pf, []byte(in), 0)
			ae, aerr := DecodeFloat64SliceArena(&af, []byte(in), 0, &af64)
			if pe != ae || !errIs(perr, aerr) || !reflect.DeepEqual(pf, af) {
				t.Errorf("float64 %q: plain (%v,%v,%v) arena (%v,%v,%v)", in, pf, pe, perr, af, ae, aerr)
			}

			var pi, ai []int32
			pe, perr = DecodeIntSlice(&pi, []byte(in), 0)
			ae, aerr = DecodeIntSliceArena(&ai, []byte(in), 0, &ai32)
			if pe != ae || !errIs(perr, aerr) || !reflect.DeepEqual(pi, ai) {
				t.Errorf("int32 %q: plain (%v,%v,%v) arena (%v,%v,%v)", in, pi, pe, perr, ai, ae, aerr)
			}

			var pu, au []uint16
			pe, perr = DecodeUintSlice(&pu, []byte(in), 0)
			ae, aerr = DecodeUintSliceArena(&au, []byte(in), 0, &au16)
			if pe != ae || !errIs(perr, aerr) || !reflect.DeepEqual(pu, au) {
				t.Errorf("uint16 %q: plain (%v,%v,%v) arena (%v,%v,%v)", in, pu, pe, perr, au, ae, aerr)
			}
		})
	}
}

func errIs(a, b error) bool {
	return (a == nil) == (b == nil) && (a == nil || a.Error() == b.Error())
}

// TestArenaCarveExclusive locks the arena's central safety property: each carve
// owns its full capacity exclusively, so appending to one decoded slice — within
// leftover capacity or beyond it — can never touch a neighbouring slice carved
// from the same chunk. An exact scalar count also leaves len == cap, so the
// in-capacity case only arises on a miscount; the append-beyond-cap path (a heap
// reallocation) is the one exercised on well-formed input.
func TestArenaCarveExclusive(t *testing.T) {
	arena := NewArena[float64]()
	var slices [][]float64
	for k := 0; k < 300; k++ { // 900 elements over 512-element chunks: crosses boundaries
		in := fmt.Sprintf("[%d,%d,%d]", k, k+1, k+2)
		var s []float64
		if _, err := DecodeFloat64SliceArena(&s, []byte(in), 0, &arena); err != nil {
			t.Fatal(err)
		}
		if len(s) != 3 || cap(s) != 3 {
			t.Fatalf("carve %d: len %d cap %d, want 3/3 (exact count ⇒ len == cap)", k, len(s), cap(s))
		}
		slices = append(slices, s)
	}
	// Grow the first slice far past its capacity; the append must reallocate
	// (len == cap) and leave every other carve untouched.
	grown := append(slices[0], make([]float64, 100)...)
	_ = grown
	for k, s := range slices {
		if s[0] != float64(k) || s[1] != float64(k+1) || s[2] != float64(k+2) {
			t.Fatalf("carve %d corrupted after neighbour append: %v", k, s)
		}
	}
}

// TestArenaCarveAlignmentAndKinds checks that carves of narrow element kinds are
// correctly aligned for their type and decode correctly when interleaved.
//
// The invariant here is weaker than it used to need to be, and that is the point.
// Carving out of a shared []byte chunk meant every carve had to be rounded up to
// 8 bytes by hand, because a []uint16 view was a reinterpretation of bytes that
// happened to be there. An arena is typed now — its chunks are []T — so an
// element lands at a T-aligned offset by construction, and the alignment to
// check is T's own, not a blanket 8.
func TestArenaCarveAlignmentAndKinds(t *testing.T) {
	ai32, au16, af := NewArena[int32](), NewArena[uint16](), NewArena[float64]()
	for k := 0; k < 100; k++ {
		var i32 []int32
		var u16 []uint16
		var f []float64
		if _, err := DecodeIntSliceArena(&i32, []byte(`[-1,2,-3,4,-5]`), 0, &ai32); err != nil {
			t.Fatal(err)
		}
		if _, err := DecodeUintSliceArena(&u16, []byte(`[1,2,3]`), 0, &au16); err != nil {
			t.Fatal(err)
		}
		if _, err := DecodeFloat64SliceArena(&f, []byte(`[1.5]`), 0, &af); err != nil {
			t.Fatal(err)
		}
		if uintptr(unsafe.Pointer(unsafe.SliceData(i32)))%unsafe.Alignof(i32[0]) != 0 ||
			uintptr(unsafe.Pointer(unsafe.SliceData(u16)))%unsafe.Alignof(u16[0]) != 0 ||
			uintptr(unsafe.Pointer(unsafe.SliceData(f)))%unsafe.Alignof(f[0]) != 0 {
			t.Fatalf("iteration %d: misaligned carve", k)
		}
		if !reflect.DeepEqual(i32, []int32{-1, 2, -3, 4, -5}) ||
			!reflect.DeepEqual(u16, []uint16{1, 2, 3}) ||
			!reflect.DeepEqual(f, []float64{1.5}) {
			t.Fatalf("iteration %d: wrong contents %v %v %v", k, i32, u16, f)
		}
	}
}

// TestArenaThresholdBypass pins the direct-make fallback: a backing over
// arenaMaxCarve bytes must not come from a chunk (it would burn most of one and
// widen the pinning trade-off to large slices). White-box: the arena must record
// nothing stored for the big decode, since it never reached it.
func TestArenaThresholdBypass(t *testing.T) {
	arena := NewArena[float64]()
	var small []float64
	if _, err := DecodeFloat64SliceArena(&small, []byte(`[1,2]`), 0, &arena); err != nil {
		t.Fatal(err)
	}
	storedBefore := arena.Size()
	in := "[0"
	for i := 1; i <= arenaMaxCarve/8; i++ { // one element over the threshold
		in += fmt.Sprintf(",%d", i)
	}
	var big []float64
	if _, err := DecodeFloat64SliceArena(&big, []byte(in+"]"), 0, &arena); err != nil {
		t.Fatal(err)
	}
	if arena.Size() != storedBefore {
		t.Fatalf("big carve consumed arena: stored %d -> %d", storedBefore, arena.Size())
	}
	if len(big) != arenaMaxCarve/8+1 || big[len(big)-1] != float64(arenaMaxCarve/8) {
		t.Fatalf("big slice misdecoded: len %d", len(big))
	}
	// And exactly at the threshold it still carves.
	in = "[0"
	for i := 1; i < arenaMaxCarve/8; i++ {
		in += fmt.Sprintf(",%d", i)
	}
	var edge []float64
	if _, err := DecodeFloat64SliceArena(&edge, []byte(in+"]"), 0, &arena); err != nil {
		t.Fatal(err)
	}
	if arena.Size() == storedBefore {
		t.Fatalf("threshold-sized carve did not come from the arena")
	}
}

// TestArenaReuseKeepsBacking mirrors TestDecodeFloat64SliceReusesBacking for the
// arena reader: a second decode into a non-nil slice must reset and reuse the
// existing (arena-carved) backing without carving again, and must replace, not
// append.
func TestArenaReuseKeepsBacking(t *testing.T) {
	arena := NewArena[float64]()
	var s []float64
	if _, err := DecodeFloat64SliceArena(&s, []byte(`[1,2,3]`), 0, &arena); err != nil {
		t.Fatal(err)
	}
	p0 := unsafe.SliceData(s)
	storedBefore := arena.Size()
	if _, err := DecodeFloat64SliceArena(&s, []byte(`[7,8]`), 0, &arena); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(s, []float64{7, 8}) {
		t.Fatalf("reuse did not replace: %v", s)
	}
	if unsafe.SliceData(s) != p0 {
		t.Fatalf("reuse reallocated the backing")
	}
	if arena.Size() != storedBefore {
		t.Fatalf("reuse carved again: stored %d -> %d", storedBefore, arena.Size())
	}
}

// TestArenaRandomizedParity drives both readers over randomly generated scalar
// arrays (sizes straddling the carve threshold, nulls sprinkled in) and demands
// bit-identical results — a cheap fuzz for the carve/count interaction.
func TestArenaRandomizedParity(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	arena := NewArena[float64]()
	for iter := 0; iter < 2000; iter++ {
		n := rng.Intn(80)
		in := "["
		for i := 0; i < n; i++ {
			if i > 0 {
				in += ","
			}
			switch rng.Intn(4) {
			case 0:
				in += "null"
			case 1:
				in += fmt.Sprintf("%d", rng.Int63n(1e12)-5e11)
			default:
				in += fmt.Sprintf("%g", (rng.Float64()-0.5)*1e6)
			}
		}
		in += "]"
		var pf, af []float64
		pe, perr := DecodeFloat64Slice(&pf, []byte(in), 0)
		ae, aerr := DecodeFloat64SliceArena(&af, []byte(in), 0, &arena)
		if pe != ae || !errIs(perr, aerr) || !reflect.DeepEqual(pf, af) {
			t.Fatalf("iter %d %q: plain (%v,%v) arena (%v,%v)", iter, in, pe, perr, ae, aerr)
		}
	}
}

// BenchmarkDecodeSmallSlices contrasts the plain and arena batch readers on the
// workload the arena exists for: a long run of 3-element float64 arrays (the
// shape of marine_ik's pos/rot/scl keys), decoded into fresh nil slices so every
// array allocates its backing. The plain/arena delta isolates the per-slice
// mallocgc cost the arena removes; it is the standing measurement for the
// //lightning:arena trade-off.
func BenchmarkDecodeSmallSlices(b *testing.B) {
	docs := make([][]byte, 512)
	for i := range docs {
		docs[i] = []byte(fmt.Sprintf("[%d.5,%d.25,%d.125]", i, i+1, i+2))
	}
	b.Run("plain", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var s []float64
			if _, err := DecodeFloat64Slice(&s, docs[i%len(docs)], 0); err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("arena", func(b *testing.B) {
		b.ReportAllocs()
		arena := NewArena[float64]()
		for i := 0; i < b.N; i++ {
			var s []float64
			if _, err := DecodeFloat64SliceArena(&s, docs[i%len(docs)], 0, &arena); err != nil {
				b.Fatal(err)
			}
		}
	})
}
