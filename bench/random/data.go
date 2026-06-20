// Code generated from random.json (minio/simdjson-go testdata); root renamed Benchmark.
package bench

type Benchmark struct {
	ID      int64  `json:"id"`
	Jsonrpc string `json:"jsonrpc,nocopy"`
	Result  []struct {
		Admin     bool   `json:"admin"`
		Age       int64  `json:"age"`
		Avatar    string `json:"avatar,nocopy"`
		BirthDate string `json:"birthDate,nocopy"`
		Company   string `json:"company,nocopy"`
		Email     string `json:"email,nocopy"`
		Field     string `json:"field,nocopy"`
		Friends   []struct {
			ID    int64  `json:"id"`
			Name  string `json:"name,nocopy"`
			Phone string `json:"phone,nocopy"`
		} `json:"friends"`
		ID    int64  `json:"id"`
		Name  string `json:"name,nocopy"`
		Phone string `json:"phone,nocopy"`
	} `json:"result"`
	Total int64 `json:"total"`
}
