// Code generated from github.com/go-json-experiment/jsonbench testdata; the
// canada_geometry root type, with the root renamed to Benchmark for the lightning harness.
package bench

type (
	Benchmark struct {
		Type     string `json:"type,nocopy"`
		Features []struct {
			Type       string `json:"type,nocopy"`
			Properties struct {
				Name string `json:"name,nocopy"`
			} `json:"properties"`
			Geometry struct {
				Type        string        `json:"type,nocopy"`
				Coordinates [][][2]float64 `json:"coordinates"`
			} `json:"geometry"`
		} `json:"features"`
	}
)
