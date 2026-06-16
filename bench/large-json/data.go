package bench

// Log is the root of the large-json benchmark: the San Francisco "citylots"
// GeoJSON FeatureCollection (~10k parcel polygons, ~8 MB), decoded from
// data.json. The dominant cost is the deeply nested coordinate arrays
// ([][][]float64) — thousands of float pairs per document — so this case
// stresses slice growth and float parsing at scale, unlike the small
// object-heavy records in the other cases.
//
// Every nested type is an anonymous struct so that only Log carries a generated
// UnmarshalJSON; that keeps the `type logStd Log` baseline used by the stdlib
// and sonic benchmarks purely reflection-based.
//
// The property strings carry the "nocopy" option so they alias the input buffer
// instead of being copied out — the input is retained for the lifetime of the
// benchmark, so this is safe and removes the nine per-feature string allocations
// (the bulk of this case's allocs).
type Log struct {
	Type     string `json:"type"`
	Features []struct {
		Type       string `json:"type"`
		Properties struct {
			MapBlkLot string `json:"MAPBLKLOT,nocopy"`
			BlkLot    string `json:"BLKLOT,nocopy"`
			BlockNum  string `json:"BLOCK_NUM,nocopy"`
			LotNum    string `json:"LOT_NUM,nocopy"`
			FromSt    string `json:"FROM_ST,nocopy"`
			ToSt      string `json:"TO_ST,nocopy"`
			Street    string `json:"STREET,nocopy"`
			StType    string `json:"ST_TYPE,nocopy"`
			OddEven   string `json:"ODD_EVEN,nocopy"`
		} `json:"properties"`
		Geometry struct {
			Type        string        `json:"type"`
			Coordinates [][][]float64 `json:"coordinates"`
		} `json:"geometry"`
	} `json:"features"`
}
