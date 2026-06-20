// Code generated from mesh_pretty.json (minio/simdjson-go testdata); root renamed Benchmark.
package bench

type Benchmark struct {
	Batches []struct {
		IndexRange  []int64 `json:"indexRange"`
		UsedBones   []int64 `json:"usedBones"`
		VertexRange []int64 `json:"vertexRange"`
	} `json:"batches"`
	Colors       []int64     `json:"colors"`
	Indices      []int64     `json:"indices"`
	Influences   [][]float64 `json:"influences"`
	MorphTargets struct {
	} `json:"morphTargets"`
	Normals   []float64 `json:"normals"`
	Positions []float64 `json:"positions"`
	Tex0      []float64 `json:"tex0"`
}
