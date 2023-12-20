package dto

import (
	"math"
)

type JourneyPoints struct {
	StartingPoint PointNode `json:"startingPoint,omitempty"`
	EndingPoint   PointNode `json:"endingPoint,omitempty"`
	Distance      float32   `json:"distance,omitempty"`
}

type PointNode struct {
	X, Y float64
	////TODO: Dynamic weighted routing
	//Speed      int // Speed limit in km/h
	//Width      int // How many points can occupy
	//Occupation int // How many cars at a specific node
}

type NodePair struct {
	From, To *PointNode
}

type Edge struct {
	From, To PointNode
}

func (p *PointNode) Heuristic(target PointNode) float64 {
	euclideanDistance := math.Sqrt(math.Pow(target.X-p.X, 2) + math.Pow(target.Y-p.Y, 2))
	worstSpeed := 70.0
	worstSurfaceMultiplier := 1.0
	return euclideanDistance * worstSurfaceMultiplier / worstSpeed
}

func (p *JourneyPoints) ToJourneyPointsDTO(dateCreate string, status string, attempts int32) Journey {
	return Journey{
		StartingPointNode: p.StartingPoint,
		EndingPointNode:   p.EndingPoint,
		Distance:          p.Distance,
		DateCreate:        dateCreate,
		Status:            JourneyStatus(status),
		Attempts:          attempts,
	}
}
