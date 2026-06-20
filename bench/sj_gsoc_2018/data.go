// Code generated from gsoc_2018.json (minio/simdjson-go testdata); root renamed Benchmark.
package bench

type Benchmark map[string]struct {
	Context string `json:"@context,nocopy"`
	Type    string `json:"@type,nocopy"`
	Author  struct {
		Type string `json:"@type,nocopy"`
		Name string `json:"name,nocopy"`
	} `json:"author"`
	Description string `json:"description,nocopy"`
	Name        string `json:"name,nocopy"`
	Sponsor     struct {
		Type                      string `json:"@type,nocopy"`
		Description               string `json:"description,nocopy"`
		DisambiguatingDescription string `json:"disambiguatingDescription,nocopy"`
		Logo                      string `json:"logo,nocopy"`
		Name                      string `json:"name,nocopy"`
		URL                       string `json:"url,nocopy"`
	} `json:"sponsor"`
}
