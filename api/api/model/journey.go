package model

import (
	"database/sql/driver"
	"fmt"
	openapi "github.com/bohdanabadi/Traffic-Simulation/api/api/generated/go"
)

type Journey struct {
	ID             uint          `gorm:"primaryKey"`
	StartingPointX float64       `json:"startingPointX" gorm:"type:float;not null"`
	StartingPointY float64       `json:"startingPointY" gorm:"type:float;not null"`
	EndingPointX   float64       `json:"endingPointX" gorm:"type:float;not null"`
	EndingPointY   float64       `json:"endingPointY" gorm:"type:float;not null"`
	Distance       float32       `json:"distance" gorm:"type:float;not null"`
	DateCreate     string        `json:"dateCreate" gorm:"type:date;not null"`
	Status         JourneyStatus `json:"status" gorm:"type:journey_status_enum;not null"`
	Attempts       int32         `json:"attempts" gorm:"type:int;not null"`
}

// JourneyStatus represents the status of a journey
type JourneyStatus string

// Journey status enumeration
const (
	InQueue    JourneyStatus = "IN QUEUE"
	InProgress JourneyStatus = "IN PROGRESS"
	Finished   JourneyStatus = "FINISHED"
)

// Ensure journey status implements the Valuer interface from the sql/driver package
func (status JourneyStatus) Value() (driver.Value, error) {
	return string(status), nil
}

// Scan implements the sql.Scanner interface for JourneyStatus
func (js *JourneyStatus) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		*js = JourneyStatus(v)
	case string:
		*js = JourneyStatus(v)
	default:
		return fmt.Errorf("failed to scan JourneyStatus; expect []byte or string got %T", value)
	}
	return nil
}

func MapToGormModel(journey openapi.Journey) (Journey, error) {
	//var status JourneyStatus
	//switch journey.Status {
	//case "IN QUEUE":
	//	status = InQueue
	//case "IN PROGRESS":
	//	status = InProgress
	//case "FINISHED":
	//	status = Finished
	//default:
	//	return Journey{}, apperror.NewAppError(400, "Invalid journey status", nil)
	//}

	return Journey{
		StartingPointX: journey.StartingPoint.X,
		StartingPointY: journey.StartingPoint.Y,
		EndingPointX:   journey.EndingPoint.X,
		EndingPointY:   journey.EndingPoint.Y,
		Distance:       journey.Distance,
		DateCreate:     journey.DateCreate,
		Status:         JourneyStatus(journey.Status),
		Attempts:       int32(0),
	}, nil
}

func MapToJourneyDTO(journey Journey) (openapi.Journey, error) {

	return openapi.Journey{
		Id:            int32(journey.ID),
		StartingPoint: openapi.Point{X: journey.StartingPointX, Y: journey.StartingPointY},
		EndingPoint:   openapi.Point{X: journey.EndingPointX, Y: journey.EndingPointY},
		Distance:      journey.Distance,
		DateCreate:    journey.DateCreate,
		Status:        string(journey.Status),
	}, nil
}
