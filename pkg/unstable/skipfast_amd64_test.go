//go:build amd64

package unstable

import (
	"encoding/json"
	"math/rand"
	"strings"
	"testing"
)

// BenchmarkSkipBlocksVariant pits the AVX2 and AVX-512 whole-loop forms (and
// the Go maskBlock loop) against each other on the standard skip shapes, to
// decide whether the AVX-512 variant earns its gate.
func BenchmarkSkipBlocksVariant(b *testing.B) {
	if !useSkipBlocks {
		b.Skip("no assembly skip loop")
	}
	sb, sb512 := useSkipBlocks, useSkipBlocks512
	defer func() { useSkipBlocks, useSkipBlocks512 = sb, sb512 }()
	variants := []struct {
		name         string
		blocks, b512 bool
	}{
		{"goloop", false, false},
		{"avx2", true, false},
	}
	if sb512 {
		variants = append(variants, struct {
			name         string
			blocks, b512 bool
		}{"avx512", true, true})
	}
	shapes := []struct {
		name string
		data []byte
	}{
		{"stringObj", bigStringObject(300)},
		{"numberObj", numberObject(300)},
		{"nestedMixed", nestedMixed(200)},
	}
	for _, va := range variants {
		for _, s := range shapes {
			b.Run(s.name+"/"+va.name, func(b *testing.B) {
				useSkipBlocks, useSkipBlocks512 = va.blocks, va.b512
				b.SetBytes(int64(len(s.data)))
				for i := 0; i < b.N; i++ {
					if _, err := skipContainerFast(s.data, 0, s.data[0]); err != nil {
						b.Fatal(err)
					}
				}
			})
		}
	}
}

// TestSkipBlocksVariants runs the boundary corpus plus a randomized document
// fuzz through every skip implementation this machine has — the AVX-512 loop,
// the AVX2 loop, and the Go maskBlock loop — by flipping the dispatch flags,
// comparing each against the scalar oracle. This is what locks the assembly
// loops (and the FP result offsets the asmdecl vet pass cannot check — see the
// maskBlock gotcha in CLAUDE.md) to the tested Go bit math.
func TestSkipBlocksVariants(t *testing.T) {
	if !useSkipBlocks {
		t.Skip("no assembly skip loop on this machine")
	}
	sb, sb512 := useSkipBlocks, useSkipBlocks512
	defer func() { useSkipBlocks, useSkipBlocks512 = sb, sb512 }()

	variants := []struct {
		name         string
		blocks, b512 bool
	}{
		{"goloop", false, false},
		{"avx2", true, false},
	}
	if sb512 {
		variants = append(variants, struct {
			name         string
			blocks, b512 bool
		}{"avx512", true, true})
	}

	docs := boundaryDocs()
	r := rand.New(rand.NewSource(7))
	for iter := 0; iter < 4000; iter++ {
		v := randValue(r, 4)
		doc, err := json.Marshal(v)
		if err != nil {
			t.Fatal(err)
		}
		if len(doc) > 0 && (doc[0] == '{' || doc[0] == '[') {
			docs = append(docs, string(doc))
		}
	}

	for _, va := range variants {
		useSkipBlocks, useSkipBlocks512 = va.blocks, va.b512
		for _, c := range docs {
			for _, pad := range []int{0, 70} {
				data := []byte(c + strings.Repeat(" ", pad))
				want, werr := refSkip(data, 0)
				got, gerr := skipContainerFast(data, 0, data[0])
				if got != want || (werr == nil) != (gerr == nil) {
					t.Fatalf("%s: %q pad %d: fast=(%d,%v) ref=(%d,%v)",
						va.name, c, pad, got, gerr, want, werr)
				}
				// Truncation safety per variant: cut anywhere, must error not panic.
				if len(data) > 2 {
					tr := data[:1+len(data)/2]
					_, e1 := refSkip(tr, 0)
					_, e2 := skipContainerFast(tr, 0, tr[0])
					if (e1 == nil) != (e2 == nil) {
						t.Fatalf("%s: %q truncated: ref err=%v fast err=%v", va.name, c, e1, e2)
					}
				}
			}
		}
	}
}
