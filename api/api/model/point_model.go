package model

import (
	openapi "github.com/bohdanabadi/Traffic-Simulation/api/api/generated/go"
	"gorm.io/gorm"
)

type Point struct {
	gorm.Model
	StartingPointX float64 `json:"startingPointX" gorm:"type:float;not null"`
	StartingPointY float64 `json:"startingPointY" gorm:"type:float;not null"`
	EndingPointX   float64 `json:"endingPointX" gorm:"type:float;not null"`
	EndingPointY   float64 `json:"endingPointY" gorm:"type:float;not null"`
	Distance       float32 `json:"distance" gorm:"type:float;not null"`
}

func MapToPotentialJourneyPointsDTO(point Point) openapi.PotentialJourneyPoints {
	return openapi.PotentialJourneyPoints{
		StartingPoint: openapi.Point{X: point.StartingPointX, Y: point.StartingPointY},
		EndingPoint:   openapi.Point{X: point.EndingPointX, Y: point.EndingPointY},
		Distance:      point.Distance,
	}
}
