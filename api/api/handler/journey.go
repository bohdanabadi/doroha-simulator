package handler

import (
	"errors"
	"github.com/bohdanabadi/Traffic-Simulation/api/api/apperror"
	openapi "github.com/bohdanabadi/Traffic-Simulation/api/api/generated/go"
	"github.com/bohdanabadi/Traffic-Simulation/api/api/service"
	"github.com/bohdanabadi/Traffic-Simulation/api/observibility"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateJourney(c *gin.Context) {
	var journey []openapi.Journey
	if err := c.ShouldBindJSON(&journey); err != nil {
		observibility.GetMetrics().LogErrorCounter()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the journey fields
	if !validateJourney(journey) {
		observibility.GetMetrics().LogErrorCounter()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid journey data"})
		return
	}

	// Call your journey creation service here
	if err := service.CreateJourney(&journey); err != nil {
		observibility.GetMetrics().LogErrorCounter()
		var appErr *apperror.AppError
		if errors.As(err, &appErr) {
			c.JSON(appErr.Code, gin.H{"error": appErr.Description})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Journey created successfully"})
}

func GetJourney(c *gin.Context) {
	status, exists := c.GetQuery("status")

	if exists {
		if !validateJourneyStatus(status) {
			observibility.GetMetrics().LogErrorCounter()
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
			return
		}
	}
	journeys, err := service.GetJourneys(status)
	if err != nil {
		observibility.GetMetrics().LogErrorCounter()
		var appErr *apperror.AppError
		if errors.As(err, &appErr) {
			c.JSON(appErr.Code, gin.H{"error": appErr.Description})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"journeys": journeys})
}

func UpdateJourneyStatus(c *gin.Context) {
	var journeyToPatch openapi.JourneyListStatus
	if err := c.ShouldBindJSON(&journeyToPatch); err != nil {
		observibility.GetMetrics().LogErrorCounter()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !validateJourneyStatus(journeyToPatch.Status) {
		observibility.GetMetrics().LogErrorCounter()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}
	if err := service.UpdateJourneysStatus(&journeyToPatch); err != nil {
		observibility.GetMetrics().LogErrorCounter()
		var appErr *apperror.AppError
		if errors.As(err, &appErr) {
			c.JSON(appErr.Code, gin.H{"error": appErr.Description})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "Journeys patched"})
}

func validateJourney(journeys []openapi.Journey) bool {
	for _, journey := range journeys {
		if journey.StartingPoint.Y == 0 || journey.StartingPoint.X == 0 ||
			journey.EndingPoint.X == 0 || journey.EndingPoint.Y == 0 {
			return false
		}

		// Check for valid distance
		if journey.Distance <= 0 {
			return false
		}

		// Check for valid date (implement this based on your date format)
		// ...

		// Check for valid status
		if !validateJourneyStatus(journey.Status) {
			return false
		}
	}
	return true
}

func validateJourneyStatus(status string) bool {
	// Check for valid status
	validStatuses := map[string]bool{
		"IN QUEUE":    true,
		"IN PROGRESS": true,
		"FINISHED":    true,
	}
	if !validStatuses[status] {
		return false
	}
	return true

}
