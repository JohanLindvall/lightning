// Code generated from update_center.json (minio/simdjson-go testdata); root renamed Benchmark.
package bench

type Benchmark struct {
	ConnectionCheckUrl string `json:"connectionCheckUrl,nocopy"`
	Core               struct {
		BuildDate string `json:"buildDate,nocopy"`
		Name      string `json:"name,nocopy"`
		Sha1      string `json:"sha1,nocopy"`
		URL       string `json:"url,nocopy"`
		Version   string `json:"version,nocopy"`
	} `json:"core"`
	ID      string `json:"id,nocopy"`
	Plugins map[string]struct {
		BuildDate    string `json:"buildDate,nocopy"`
		Dependencies []struct {
			Name     string `json:"name,nocopy"`
			Optional bool   `json:"optional"`
			Version  string `json:"version,nocopy"`
		} `json:"dependencies"`
		Developers []struct {
			DeveloperId string `json:"developerId,nocopy"`
			Email       string `json:"email,nocopy"`
			Name        string `json:"name,nocopy"`
		} `json:"developers"`
		Excerpt                string   `json:"excerpt,nocopy"`
		Gav                    string   `json:"gav,nocopy"`
		Labels                 []string `json:"labels,nocopy"`
		Name                   string   `json:"name,nocopy"`
		ReleaseTimestamp       string   `json:"releaseTimestamp,nocopy"`
		RequiredCore           string   `json:"requiredCore,nocopy"`
		Scm                    string   `json:"scm,nocopy"`
		Sha1                   string   `json:"sha1,nocopy"`
		Title                  string   `json:"title,nocopy"`
		URL                    string   `json:"url,nocopy"`
		Version                string   `json:"version,nocopy"`
		Wiki                   string   `json:"wiki,nocopy"`
		PreviousTimestamp      string   `json:"previousTimestamp,nocopy"`
		PreviousVersion        string   `json:"previousVersion,nocopy"`
		CompatibleSinceVersion string   `json:"compatibleSinceVersion,nocopy"`
	} `json:"plugins"`
	Signature struct {
		Certificates     []string `json:"certificates,nocopy"`
		CorrectDigest    string   `json:"correct_digest,nocopy"`
		CorrectSignature string   `json:"correct_signature,nocopy"`
		Digest           string   `json:"digest,nocopy"`
		Signature        string   `json:"signature,nocopy"`
	} `json:"signature"`
	UpdateCenterVersion string `json:"updateCenterVersion,nocopy"`
}
