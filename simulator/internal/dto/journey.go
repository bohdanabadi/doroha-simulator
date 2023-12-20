package dto

import (
	"database/sql/driver"
	"fmt"
	"github.com/wk8/go-ordered-map/v2"
)

type Journey struct {
	Id                  int32                                    `json:"id,omitempty"`
	StartingPointNode   PointNode                                `json:"startingPoint,omitempty"`
	EndingPointNode     PointNode                                `json:"endingPoint,omitempty"`
	PrevPointNode       PointNode                                `json:"prevPoint,omitempty"`
	CurrentPointNode    PointNode                                `json:"currentPoint,omitempty"`
	Path                orderedmap.OrderedMap[string, PointNode] `json:"path,omitempty"`
	TotalTripCost       float64
	CostFromTheStart    float64
	AccumulatedMovement float64
	Progress            float32       `json:"progress,omitempty"`
	Distance            float32       `json:"distance,omitempty"`
	DateCreate          string        `json:"dateCreate,omitempty"`
	Status              JourneyStatus `json:"status,omitempty"`
	Attempts            int32         `json:"attempts,omitempty"`
}

type JourneyOption func(*Journey)

func WithId(id int32) JourneyOption {
	return func(j *Journey) {
		j.Id = id
	}
}

func WithStartingPointNode(startingPointNode PointNode) JourneyOption {
	return func(j *Journey) {
		j.StartingPointNode = startingPointNode
	}
}

func WithEndingPointNode(endingPointNode PointNode) JourneyOption {
	return func(j *Journey) {
		j.EndingPointNode = endingPointNode
	}
}

func WithPrevPointNode(prevPointNode PointNode) JourneyOption {
	return func(j *Journey) {
		j.PrevPointNode = prevPointNode
	}
}

func WithCurrentPointNode(currentPointNode PointNode) JourneyOption {
	return func(j *Journey) {
		j.CurrentPointNode = currentPointNode
	}
}

func WithPath(path orderedmap.OrderedMap[string, PointNode]) JourneyOption {
	return func(j *Journey) {
		j.Path = path
	}
}

func WithTotalTripCost(totalTripCost float64) JourneyOption {
	return func(j *Journey) {
		j.TotalTripCost = totalTripCost
	}
}

func WithProgress(progress float32) JourneyOption {
	return func(j *Journey) {
		j.Progress = progress
	}
}

func WithDistance(distance float32) JourneyOption {
	return func(j *Journey) {
		j.Distance = distance
	}
}

func WithDateCreate(dateCreate string) JourneyOption {
	return func(j *Journey) {
		j.DateCreate = dateCreate
	}
}

func WithStatus(status JourneyStatus) JourneyOption {
	return func(j *Journey) {
		j.Status = status
	}
}

func WithAttempts(attempts int32) JourneyOption {
	return func(j *Journey) {
		j.Attempts = attempts
	}
}

// NewJourney creates a new instance of Journey With initial values for its fields.
func NewJourney(options ...JourneyOption) *Journey {
	j := &Journey{}
	for _, option := range options {
		option(j)
	}
	return j
}

// New Write documentation
type JourneysResponse struct {
	Journeys []Journey `json:"journeys"`
}
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

// Ensure journey status implements the Scanner interface from the sql/driver package
func (status *JourneyStatus) Scan(value interface{}) error {
	if valueBytes, ok := value.([]byte); ok {
		*status = JourneyStatus(valueBytes)
		return nil
	}
	return fmt.Errorf("failed to scan journey status")
}

type JourneyListStatusUpdate struct {
	Ids    []int32 `json:"ids,omitempty"`
	Status string  `json:"status,omitempty"`
}
