// Code generated from marine_ik.json (minio/simdjson-go testdata); root renamed Benchmark.
package bench

type Benchmark struct {
	Animations []struct {
		Fps    int64  `json:"fps"`
		Name   string `json:"name"`
		Tracks []any  `json:"tracks"`
	} `json:"animations"`
	Geometries []struct {
		Data struct {
			Animations []struct {
				Fps       int64 `json:"fps"`
				Hierarchy []struct {
					Keys []struct {
						Pos  []float64 `json:"pos"`
						Rot  []float64 `json:"rot"`
						Scl  []float64 `json:"scl"`
						Time float64   `json:"time"`
					} `json:"keys"`
					Parent int64 `json:"parent"`
				} `json:"hierarchy"`
				Length float64 `json:"length"`
				Name   string  `json:"name"`
			} `json:"animations"`
			Bones []struct {
				Name   string    `json:"name"`
				Parent int64     `json:"parent"`
				Pos    []float64 `json:"pos"`
				Rotq   []float64 `json:"rotq"`
				Scl    []int64   `json:"scl"`
			} `json:"bones"`
			Faces               []int64 `json:"faces"`
			InfluencesPerVertex int64   `json:"influencesPerVertex"`
			Metadata            struct {
				Bones     int64  `json:"bones"`
				Faces     int64  `json:"faces"`
				Generator string `json:"generator"`
				Normals   int64  `json:"normals"`
				Uvs       int64  `json:"uvs"`
				Version   int64  `json:"version"`
				Vertices  int64  `json:"vertices"`
			} `json:"metadata"`
			Name        string      `json:"name"`
			Normals     []float64   `json:"normals"`
			SkinIndices []int64     `json:"skinIndices"`
			SkinWeights []float64   `json:"skinWeights"`
			Uvs         [][]float64 `json:"uvs"`
			Vertices    []float64   `json:"vertices"`
		} `json:"data"`
		Type string `json:"type"`
		Uuid string `json:"uuid"`
	} `json:"geometries"`
	Images []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
		Uuid string `json:"uuid"`
	} `json:"images"`
	Materials []struct {
		Blending     string `json:"blending"`
		Color        int64  `json:"color"`
		DepthTest    bool   `json:"depthTest"`
		DepthWrite   bool   `json:"depthWrite"`
		Emissive     int64  `json:"emissive"`
		Map          string `json:"map"`
		Name         string `json:"name"`
		Shininess    int64  `json:"shininess"`
		Specular     int64  `json:"specular"`
		Transparent  bool   `json:"transparent"`
		Type         string `json:"type"`
		Uuid         string `json:"uuid"`
		VertexColors int64  `json:"vertexColors"`
	} `json:"materials"`
	Metadata struct {
		Generator  string  `json:"generator"`
		SourceFile string  `json:"sourceFile"`
		Type       string  `json:"type"`
		Version    float64 `json:"version"`
	} `json:"metadata"`
	Object struct {
		Children []struct {
			CastShadow    bool    `json:"castShadow"`
			Geometry      string  `json:"geometry"`
			Material      string  `json:"material"`
			Matrix        []int64 `json:"matrix"`
			Name          string  `json:"name"`
			ReceiveShadow bool    `json:"receiveShadow"`
			Type          string  `json:"type"`
			Uuid          string  `json:"uuid"`
			Visible       bool    `json:"visible"`
		} `json:"children"`
		Matrix []int64 `json:"matrix"`
		Type   string  `json:"type"`
		Uuid   string  `json:"uuid"`
	} `json:"object"`
	Textures []struct {
		Anisotropy int64   `json:"anisotropy"`
		Image      string  `json:"image"`
		MagFilter  int64   `json:"magFilter"`
		Mapping    int64   `json:"mapping"`
		MinFilter  int64   `json:"minFilter"`
		Name       string  `json:"name"`
		Repeat     []int64 `json:"repeat"`
		Uuid       string  `json:"uuid"`
		Wrap       []int64 `json:"wrap"`
	} `json:"textures"`
}
