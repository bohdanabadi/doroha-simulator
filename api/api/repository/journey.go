package repository

import (
	"github.com/bohdanabadi/Traffic-Simulation/api/api/apperror"
	"github.com/bohdanabadi/Traffic-Simulation/api/api/model"
	"github.com/bohdanabadi/Traffic-Simulation/api/db"
	"github.com/bohdanabadi/Traffic-Simulation/api/observibility"
	"gorm.io/gorm"
	"time"
)

func InsertJourney(journeys []*model.Journey) error {
	m := observibility.GetMetrics()

	if result := m.ExecuteGormQueryWithLatency(func(db *gorm.DB) *gorm.DB {
		return db.Create(journeys)
	}); result.Error != nil {
		return apperror.NewAppError(500, "Failed to create journeys", result.Error)
	}

	return nil
}

func GetJourneysByStatus(status string) (*[]model.Journey, error) {
	m := observibility.GetMetrics()

	var journeys []model.Journey
	if result := m.ExecuteGormQueryWithLatency(func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", status).Where("attempts < ?", 3).Find(&journeys)
	}); result.Error != nil {
		return nil, apperror.NewAppError(500, "Failed to fetch journeys with status : "+status, result.Error)
	}
	return &journeys, nil
}

func GetAllJourneys() (*[]model.Journey, error) {
	m := observibility.GetMetrics()

	var journeys []model.Journey

	if result := m.ExecuteGormQueryWithLatency(func(db *gorm.DB) *gorm.DB {
		return db.Find(&journeys)
	}); result.Error != nil {
		return nil, apperror.NewAppError(500, "Failed to fetch all journeys with status", result.Error)
	}
	return &journeys, nil
}
func UpdateJourneysStatus(journeyIdList []int32, status string) error {
	m := observibility.GetMetrics()

	if len(journeyIdList) == 0 {
		return nil
	}
	startTime := time.Now()

	err := db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.Journey{}).
			Where("id IN ?", journeyIdList).
			Update("status", status).
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
	duration := time.Since(startTime)

	m.DBQueryLatency.Observe(duration.Seconds())

	return nil
}
