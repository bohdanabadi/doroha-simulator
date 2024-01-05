package datastructures

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	dto "simulator/internal/dto"
	"strconv"
)

var RoadMapGraph map[dto.PointNode][]dto.PointNode
var RoadMapEdgeCostGraph map[dto.Edge]float64

func LoadGeoJSONGraph() {
	fmt.Println("Loading GeoJson")
	// Load GeoJSON data from the file
	data, err := os.ReadFile("filtered_kyiv.geojson")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)

	}

	var geoJSON dto.GeoJSON
	if err = json.Unmarshal(data, &geoJSON); err != nil {
		log.Fatalf("Failed to decode JSON: %v", err)
	}

	// Create the RoadMapGraph
	RoadMapGraph = make(map[dto.PointNode][]dto.PointNode)
	RoadMapEdgeCostGraph = make(map[dto.Edge]float64)

	// Convert each GeoJSON feature to a graph edge
	for _, feature := range geoJSON.Features {
		// Determine if the road is one-way
		oneWay := true // default value if "oneway" property is missing
		if v, ok := feature.Properties["oneway"]; ok {
			oneWay = v.(string) == "yes" // if "oneway" property is present, use its value
		}

		// Get the list of points from the feature's coordinates
		points := make([]dto.PointNode, len(feature.Geometry.Coordinates))
		for i, coord := range feature.Geometry.Coordinates {
			points[i] = dto.PointNode{X: RoundToDecimal(coord[0], 9), Y: RoundToDecimal(coord[1], 9)}
		}

		// Add an edge for each pair of consecutive points in the list
		for i := 0; i < len(points)-1; i++ {
			// Compute static edge cost for the current feature
			edgeCost := computeEdgeCost(points[i], points[i+1], feature.Properties)
			RoadMapEdgeCostGraph[dto.Edge{From: points[i], To: points[i+1]}] = edgeCost

			RoadMapGraph[points[i]] = append(RoadMapGraph[points[i]], points[i+1])
			if !oneWay {
				// If the road is not one-way, also add an edge in the opposite direction
				RoadMapGraph[points[i+1]] = append(RoadMapGraph[points[i+1]], points[i])
				edgeCostReverse := computeEdgeCost(points[i+1], points[i], feature.Properties)
				RoadMapEdgeCostGraph[dto.Edge{From: points[i+1], To: points[i]}] = edgeCostReverse
			}
		}

	}
	fmt.Println("RoadMapGraph is Done")
}

func RoundToDecimal(value float64, decimalPlaces int) float64 {
	multiplier := math.Pow(10, float64(decimalPlaces))
	return math.Round(value*multiplier) / multiplier
}
func computeEdgeCost(a, b dto.PointNode, properties map[string]interface{}) float64 {
	baseDistance := computeEuclideanDistance(a, b)

	baseSpeed := 50.0 // default

	if maxspeed, exists := properties["maxspeed"].(string); exists {
		baseSpeed, _ = strconv.ParseFloat(maxspeed, 64)
	} else if highway, exists := properties["highway"].(string); exists {
		if speed, ok := highwaySpeeds[highway]; ok {
			baseSpeed = speed
		}
	}

	surfaceMultiplier := 1.0
	if surface, exists := properties["surface"].(string); exists {
		if multiplier, ok := surfaceMultipliers[surface]; ok {
			surfaceMultiplier = multiplier
		}
	}
	return baseDistance * surfaceMultiplier / baseSpeed
}

// ComputeEuclideanDistance computes the distance between two PointNode
func computeEuclideanDistance(a, b dto.PointNode) float64 {
	deltaX := b.X - a.X
	deltaY := b.Y - a.Y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

var highwaySpeeds = map[string]float64{
	"bus_stop":       10,
	"construction":   5, // extremely slow due to ongoing construction
	"cycleway":       20,
	"living_street":  20,
	"path":           10,
	"pedestrian":     5,
	"primary":        50,
	"primary_link":   45,
	"residential":    30,
	"secondary":      40,
	"secondary_link": 35,
	"service":        20,
	"services":       20, // assuming this is similar to service
	"tertiary":       30,
	"track":          20,
	"trunk":          70, // usually these are major highways
	"trunk_link":     65,
	"unclassified":   20,
}

var surfaceMultipliers = map[string]float64{
	"asphalt":         1,
	"cobblestone":     1.2,
	"compacted":       1.1,
	"concrete":        1,
	"concrete:plates": 1.05,
	"dirt":            1.4,
	"grass":           1.5,
	"gravel":          1.3,
	"ground":          1.2,
	"paved":           1,
	"paving_stones":   1.2,
	"sett":            1.3,
	"stone":           1.4,
	"unpaved":         1.4,
	"wood":            1.6,
}
