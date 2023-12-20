package dto

type GeoJSON struct {
	Type     string `json:"type"`
	Features []struct {
		Type     string `json:"type"`
		Geometry struct {
			Type        string      `json:"type"`
			Coordinates [][]float64 `json:"coordinates"`
		} `json:"geometry"`
		Properties map[string]interface{} `json:"properties"`
	} `json:"features"`
}
