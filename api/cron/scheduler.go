package cron

import (
	"github.com/bohdanabadi/Traffic-Simulation/api/api/model"
	"github.com/bohdanabadi/Traffic-Simulation/api/observibility"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
	"log"
)

func SetupDBCleanupJob() {
	c := cron.New()

	_, err := c.AddFunc("@daily", func() {
		deleteFinishedRecords()
	})

	if err != nil {
		log.Println("Could not setup cron job", err)
	}

	c.Start()

}

func deleteFinishedRecords() {

	m := observibility.GetMetrics()

	if result := m.ExecuteGormQueryWithLatency(func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", "FINISHED").Delete(&model.Journey{})
	}); result.Error != nil {
		log.Println("Failed to delete finished journeys: %w", result.Error)
	}
}
