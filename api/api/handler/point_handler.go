package handler

import (
	"github.com/bohdanabadi/Traffic-Simulation/api/api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPotentialJourneyPoints(c *gin.Context) {
	points, err := service.GetRandomPointsWithMinDistance()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, points)
}
