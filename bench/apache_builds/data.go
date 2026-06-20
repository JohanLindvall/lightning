// Code generated from apache_builds.json (minio/simdjson-go testdata); root renamed Benchmark.
package bench

type Benchmark struct {
	AssignedLabels []struct {
	} `json:"assignedLabels"`
	Description string `json:"description,nocopy"`
	Jobs        []struct {
		Color string `json:"color,nocopy"`
		Name  string `json:"name,nocopy"`
		URL   string `json:"url,nocopy"`
	} `json:"jobs"`
	Mode            string `json:"mode,nocopy"`
	NodeDescription string `json:"nodeDescription,nocopy"`
	NodeName        string `json:"nodeName,nocopy"`
	NumExecutors    int64  `json:"numExecutors"`
	OverallLoad     struct {
	} `json:"overallLoad"`
	PrimaryView struct {
		Name string `json:"name,nocopy"`
		URL  string `json:"url,nocopy"`
	} `json:"primaryView"`
	QuietingDown   bool  `json:"quietingDown"`
	SlaveAgentPort int64 `json:"slaveAgentPort"`
	UnlabeledLoad  struct {
	} `json:"unlabeledLoad"`
	UseCrumbs   bool `json:"useCrumbs"`
	UseSecurity bool `json:"useSecurity"`
	Views       []struct {
		Name string `json:"name,nocopy"`
		URL  string `json:"url,nocopy"`
	} `json:"views"`
}
