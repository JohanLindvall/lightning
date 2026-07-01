//go:build arm64

package unstable

import "testing"

// TestSkipBlocksVariants differentially tests the NEON whole-loop skip and the
// Go maskBlock loop against the scalar oracle (see testSkipVariantCorpus).
func TestSkipBlocksVariants(t *testing.T) {
	sb := useSkipBlocks
	defer func() { useSkipBlocks = sb }()
	for _, v := range []struct {
		name string
		on   bool
	}{
		{"goloop", false},
		{"neon", true},
	} {
		useSkipBlocks = v.on
		testSkipVariantCorpus(t, v.name)
	}
}

// BenchmarkSkipBlocksVariant pits the NEON whole-loop form against the Go
// maskBlock loop on the standard skip shapes — the number that decides whether
// the arm64 assembly loop pays off (run on real hardware; qemu only checks
// correctness).
func BenchmarkSkipBlocksVariant(b *testing.B) {
	sb := useSkipBlocks
	defer func() { useSkipBlocks = sb }()
	variants := []struct {
		name string
		on   bool
	}{
		{"goloop", false},
		{"neon", true},
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
				useSkipBlocks = va.on
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
