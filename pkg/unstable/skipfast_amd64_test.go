//go:build amd64

package unstable

import "testing"

// TestSkipBlocksVariants differentially tests every skip implementation this
// machine has — the AVX-512 loop, the AVX2 loop, and the Go maskBlock loop —
// against the scalar oracle (see testSkipVariantCorpus).
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
	for _, va := range variants {
		useSkipBlocks, useSkipBlocks512 = va.blocks, va.b512
		testSkipVariantCorpus(t, va.name)
	}
}

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

// TestArrayEndAtVariants drives the growth estimate's forward container scan
// through each of its three arms — the whole-loop assembly, the Go block loop,
// and the scalar element walk — by flipping the same dispatch flags
// TestSkipBlocksVariants does. Without this the scalar arm never runs on a
// machine with AVX2, which is every machine this corpus is measured on.
func TestArrayEndAtVariants(t *testing.T) {
	sb, fs := useSkipBlocks, fastSkipAvail
	defer func() { useSkipBlocks, fastSkipAvail = sb, fs }()
	for _, v := range []struct {
		name             string
		blocks, fastSkip bool
	}{
		{"asm", sb, fs},
		{"goloop", false, fs},
		{"scalar", false, false},
	} {
		if v.blocks && !sb {
			continue // no assembly on this host
		}
		t.Run(v.name, func(t *testing.T) {
			useSkipBlocks, fastSkipAvail = v.blocks, v.fastSkip
			checkArrayEndAt(t)
		})
	}
}
