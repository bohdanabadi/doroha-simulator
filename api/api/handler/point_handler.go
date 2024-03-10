package handler

import (
	"github.com/bohdanabadi/Traffic-Simulation/api/api/service"
	"github.com/bohdanabadi/Traffic-Simulation/api/observibility"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPotentialJourneyPoints(c *gin.Context) {
	points, err := service.GetRandomPointsWithMinDistance()
	if err != nil {
		observibility.GetMetrics().LogErrorCounter()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, points)
}
