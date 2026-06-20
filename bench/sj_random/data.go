// Code generated from random.json (minio/simdjson-go testdata); root renamed Benchmark.
package bench

type Benchmark struct {
	ID      int64  `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  []struct {
		Admin     bool   `json:"admin"`
		Age       int64  `json:"age"`
		Avatar    string `json:"avatar"`
		BirthDate string `json:"birthDate"`
		Company   string `json:"company"`
		Email     string `json:"email"`
		Field     string `json:"field"`
		Friends   []struct {
			ID    int64  `json:"id"`
			Name  string `json:"name"`
			Phone string `json:"phone"`
		} `json:"friends"`
		ID    int64  `json:"id"`
		Name  string `json:"name"`
		Phone string `json:"phone"`
	} `json:"result"`
	Total int64 `json:"total"`
}
