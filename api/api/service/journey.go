package service

import (
	"fmt"
	openapi "github.com/bohdanabadi/Traffic-Simulation/api/api/generated/go"
	"github.com/bohdanabadi/Traffic-Simulation/api/api/model"
	"github.com/bohdanabadi/Traffic-Simulation/api/api/repository"
)

func CreateJourney(journeys *[]openapi.Journey) error {
	var journeysModel []*model.Journey

	for _, journey := range *journeys {
		journey, err := model.MapToGormModel(journey)
		if err != nil {
			return fmt.Errorf("failed to map DTO journey to model: %w", err)
		}
		journeysModel = append(journeysModel, &journey)
	}
	if err := repository.InsertJourney(journeysModel); err != nil {
		return fmt.Errorf("failed to insert journey: %w", err)
	}
	return nil
}

func GetJourneys(status string) (journeys *[]openapi.Journey, error error) {
	var journeysModel *[]model.Journey
	err := error
	if status == "" {
		journeysModel, err = repository.GetAllJourneys()
	} else {
		journeysModel, err = repository.GetJourneysByStatus(status)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to insert journey: %w", err)
	}
	var journeysDTO []openapi.Journey
	for _, journeyModel := range *journeysModel {
		journey, err := model.MapToJourneyDTO(journeyModel)
		if err != nil {
			return nil, fmt.Errorf("failed to map model journey to DTO: %w", err)
		}
		journeysDTO = append(journeysDTO, journey)
	}
	return &journeysDTO, nil
}

func UpdateJourneysStatus(journeysListStatus *openapi.JourneyListStatus) error {
	var journeysListId []int32

	for _, journeyIdS := range journeysListStatus.Ids {
		journeysListId = append(journeysListId, journeyIdS)
	}
	if err := repository.UpdateJourneysStatus(journeysListId, journeysListStatus.Status); err != nil {
		return fmt.Errorf("failed to update journeys status: %w", err)
	}
	return nil
}
