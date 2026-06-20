// Code generated from canada.json (minio/simdjson-go testdata); root renamed Benchmark.
package bench

type Benchmark struct {
	Features []struct {
		Geometry struct {
			Coordinates [][][2]float64 `json:"coordinates"`
			Type        string         `json:"type"`
		} `json:"geometry"`
		Properties struct {
			Name string `json:"name"`
		} `json:"properties"`
		Type string `json:"type"`
	} `json:"features"`
	Type string `json:"type"`
}
