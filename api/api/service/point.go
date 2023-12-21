package service

import (
	"fmt"
	openapi "github.com/bohdanabadi/Traffic-Simulation/api/api/generated/go"
	"github.com/bohdanabadi/Traffic-Simulation/api/api/model"
	"github.com/bohdanabadi/Traffic-Simulation/api/api/repository"
)

func GetRandomPointsWithMinDistance() (openapi.PotentialJourneyPoints, error) {
	points, err := repository.GetRandomStartAndEndPoints()
	if err != nil {
		return openapi.PotentialJourneyPoints{}, fmt.Errorf("failed to get Random Start and Ending Points: %w", err)
	}
	return model.MapToPotentialJourneyPointsDTO(points), nil
}
