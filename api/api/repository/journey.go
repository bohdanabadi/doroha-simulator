package repository

import (
	"github.com/bohdanabadi/Traffic-Simulation/api/api/apperror"
	"github.com/bohdanabadi/Traffic-Simulation/api/api/model"
	"github.com/bohdanabadi/Traffic-Simulation/api/db"
	"gorm.io/gorm"
)

func InsertJourney(journeys []*model.Journey) error {
	if result := db.DB.Create(journeys); result.Error != nil {
		return apperror.NewAppError(500, "Failed to create journeys", result.Error)
	}

	//for _, journey := range journeys {
	//	if result := db.DB.Create(journey); result.Error != nil {
	//		return apperror.NewAppError(500, "Failed to create journey", result.Error)
	//	}
	//}
	return nil
}

func GetJourneysByStatus(status string) (*[]model.Journey, error) {
	var journeys []model.Journey
	if result := db.DB.Where("status = ?", status).Where("attempts < ?", 3).Find(&journeys); result.Error != nil {
		return nil, apperror.NewAppError(500, "Failed to fetch journeys with status : "+status, result.Error)
	}
	return &journeys, nil
}

func GetAllJourneys() (*[]model.Journey, error) {
	var journeys []model.Journey
	if result := db.DB.Find(&journeys); result.Error != nil {
		return nil, apperror.NewAppError(500, "Failed to fetch all journeys with status", result.Error)
	}
	return &journeys, nil
}
func UpdateJourneysStatus(journeyIdList []int32, status string) error {
	if len(journeyIdList) == 0 {
		return nil
	}
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.Journey{}).
			Where("id IN ?", journeyIdList).
			Update("status", status).
			//TODO Modified recently
			UpdateColumn("attempts", gorm.Expr("attempts + 1")).
			Error; err != nil {
			return apperror.NewAppError(500, "Failed update journeys status, rolling back", err)
		}
		return nil
	})
	if err != nil {
		// Handle the error
		return apperror.NewAppError(500, "Error committing transaction", err)
	}
	return nil
}
