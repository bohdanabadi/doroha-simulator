package handler

import (
	"github.com/bohdanabadi/Traffic-Simulation/api/api/service"
	"github.com/bohdanabadi/Traffic-Simulation/api/observibility"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetMetrics(c *gin.Context) {
	metricType, exists := c.GetQuery("metricType")
	if exists {
		if !validateMetricType(metricType) {
			observibility.GetMetrics().LogErrorCounter()
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid metric type data"})
			return
		}
	}

	durationString, exists := c.GetQuery("duration")
	var duration time.Duration
	var err error
	if exists {
		duration, err = parseDuration(durationString)
		if duration == 0 || err != nil {
			observibility.GetMetrics().LogErrorCounter()
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid duration data"})
			return
		}
	}

	if err != nil {
		observibility.GetMetrics().LogErrorCounter()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error initializing metric service"})
		return
	}
	metric, err := service.GetMetric(metricType, duration)
	if err != nil {
		observibility.GetMetrics().LogErrorCounter()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if metric.MetricValue == nil {
		c.JSON(http.StatusNoContent, nil)
	}

	c.JSON(http.StatusOK, metric)
}
func validateMetricType(m string) bool {
	//Check for valid metrics
	metricType := map[string]bool{
		service.AvgResponseLatency:         true,
		service.AvgDatabaseLatency:         true,
		service.TotalErrorResponse:         true,
		service.SimulationsPerMinute:       true,
		service.AvgSimulationDuration:      true,
		service.AvgSimulationDistance:      true,
		service.TotalSuccessfulSimulations: true,
		service.TotalFailedSimulations:     true,
	}

	if !metricType[m] {
		return false
	}

	return true
}

func parseDuration(durationStr string) (time.Duration, error) {
	// Get the last character
	unit := durationStr[len(durationStr)-1:]
	// Get the numerical value
	value, err := strconv.Atoi(durationStr[:len(durationStr)-1])
	if err != nil {
		return 0, err
	}

	switch unit {
	case "h":
		return time.Duration(value) * time.Hour, nil
	case "d":
		return time.Duration(value) * 24 * time.Hour, nil
	default:
		return 0, nil

	}
}
