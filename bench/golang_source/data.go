// Code generated from github.com/go-json-experiment/jsonbench testdata; the
// golang_source root type, with the root renamed to Log for the lightning harness.
package bench

type (
	Log struct {
		Tree     *golangNode `json:"tree"`
		Username string      `json:"username"`
	}
	golangNode struct {
		Name     string       `json:"name"`
		Kids     []golangNode `json:"kids"`
		CLWeight float64      `json:"cl_weight"`
		Touches  int          `json:"touches"`
		MinT     uint64       `json:"min_t"`
		MaxT     uint64       `json:"max_t"`
		MeanT    uint64       `json:"mean_t"`
	}
)
