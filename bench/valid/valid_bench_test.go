package valid

import (
	"os"
	"path/filepath"
	"testing"

	ljson "github.com/JohanLindvall/lightning/pkg/json"
)

// The corpus is shared with the decoder benchmarks. Downloaded fixtures may be
// absent in a fresh checkout; the self-contained shape benchmarks live in
// pkg/json and do not depend on these files.
func BenchmarkValidCorpus(b *testing.B) {
	for _, name := range []string{
		"cloudflare", "citm_catalog", "synthea_fhir", "twitter_status",
		"twitterescaped", "gsoc_2018", "marine_ik", "mesh", "canada",
		"large-json", "numbers", "time-array", "golang_source",
	} {
		b.Run(name, func(b *testing.B) {
			doc, err := os.ReadFile(filepath.Join("..", name, "input.json"))
			if os.IsNotExist(err) {
				b.Skip("corpus fixture not downloaded")
			}
			if err != nil {
				b.Fatal(err)
			}
			b.SetBytes(int64(len(doc)))
			b.ReportAllocs()
			for b.Loop() {
				if !ljson.Valid(doc) {
					b.Fatal("invalid corpus document")
				}
			}
		})
	}
}
