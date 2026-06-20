// Code generated from marine_ik.json (minio/simdjson-go testdata); root renamed Benchmark.
package bench

type Benchmark struct {
	Animations []struct {
		Fps    int64  `json:"fps"`
		Name   string `json:"name,nocopy"`
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
				Name   string  `json:"name,nocopy"`
			} `json:"animations"`
			Bones []struct {
				Name   string    `json:"name,nocopy"`
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
				Generator string `json:"generator,nocopy"`
				Normals   int64  `json:"normals"`
				Uvs       int64  `json:"uvs"`
				Version   int64  `json:"version"`
				Vertices  int64  `json:"vertices"`
			} `json:"metadata"`
			Name        string      `json:"name,nocopy"`
			Normals     []float64   `json:"normals"`
			SkinIndices []int64     `json:"skinIndices"`
			SkinWeights []float64   `json:"skinWeights"`
			Uvs         [][]float64 `json:"uvs"`
			Vertices    []float64   `json:"vertices"`
		} `json:"data"`
		Type string `json:"type,nocopy"`
		Uuid string `json:"uuid,nocopy"`
	} `json:"geometries"`
	Images []struct {
		Name string `json:"name,nocopy"`
		URL  string `json:"url,nocopy"`
		Uuid string `json:"uuid,nocopy"`
	} `json:"images"`
	Materials []struct {
		Blending     string `json:"blending,nocopy"`
		Color        int64  `json:"color"`
		DepthTest    bool   `json:"depthTest"`
		DepthWrite   bool   `json:"depthWrite"`
		Emissive     int64  `json:"emissive"`
		Map          string `json:"map,nocopy"`
		Name         string `json:"name,nocopy"`
		Shininess    int64  `json:"shininess"`
		Specular     int64  `json:"specular"`
		Transparent  bool   `json:"transparent"`
		Type         string `json:"type,nocopy"`
		Uuid         string `json:"uuid,nocopy"`
		VertexColors int64  `json:"vertexColors"`
	} `json:"materials"`
	Metadata struct {
		Generator  string  `json:"generator,nocopy"`
		SourceFile string  `json:"sourceFile,nocopy"`
		Type       string  `json:"type,nocopy"`
		Version    float64 `json:"version"`
	} `json:"metadata"`
	Object struct {
		Children []struct {
			CastShadow    bool    `json:"castShadow"`
			Geometry      string  `json:"geometry,nocopy"`
			Material      string  `json:"material,nocopy"`
			Matrix        []int64 `json:"matrix"`
			Name          string  `json:"name,nocopy"`
			ReceiveShadow bool    `json:"receiveShadow"`
			Type          string  `json:"type,nocopy"`
			Uuid          string  `json:"uuid,nocopy"`
			Visible       bool    `json:"visible"`
		} `json:"children"`
		Matrix []int64 `json:"matrix"`
		Type   string  `json:"type,nocopy"`
		Uuid   string  `json:"uuid,nocopy"`
	} `json:"object"`
	Textures []struct {
		Anisotropy int64   `json:"anisotropy"`
		Image      string  `json:"image,nocopy"`
		MagFilter  int64   `json:"magFilter"`
		Mapping    int64   `json:"mapping"`
		MinFilter  int64   `json:"minFilter"`
		Name       string  `json:"name,nocopy"`
		Repeat     []int64 `json:"repeat"`
		Uuid       string  `json:"uuid,nocopy"`
		Wrap       []int64 `json:"wrap"`
	} `json:"textures"`
}
