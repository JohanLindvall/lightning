// Code generated from github.com/go-json-experiment/jsonbench testdata; the
// canada_geometry root type, with the root renamed to Log for the lightning harness.
package bench

type (
	Log struct {
		Type     string `json:"type"`
		Features []struct {
			Type       string `json:"type"`
			Properties struct {
				Name string `json:"name"`
			} `json:"properties"`
			Geometry struct {
				Type        string        `json:"type"`
				Coordinates [][][]float64 `json:"coordinates"`
			} `json:"geometry"`
		} `json:"features"`
	}
)
