package service

import (
	"fmt"
	"github.com/bohdanabadi/Traffic-Simulation/api/api/apperror"
	openapi "github.com/bohdanabadi/Traffic-Simulation/api/api/generated/go"
	"github.com/bohdanabadi/Traffic-Simulation/api/cache"
	"github.com/bohdanabadi/Traffic-Simulation/api/observibility"
	"gonum.org/v1/gonum/stat"
	"math"
	"net/http"
	"strconv"
	"time"
)

const (
	AvgResponseLatency         = "avg_response_latency_seconds"
	AvgDatabaseLatency         = "avg_database_latency_seconds"
	TotalErrorResponse         = "total_response_error_counter"
	SimulationsPerMinute       = "simulations_per_minute"
	AvgSimulationDuration      = "avg_simulation_duration"
	AvgSimulationDistance      = "avg_simulation_distance"
	TotalSuccessfulSimulations = "total_successful_simulations"
	TotalFailedSimulations     = "total_failed_simulations"
)

const (
	AvgResponseLatencyQuery         = "sum(rate(api_response_latency_seconds_sum[%duration])) / sum(rate(api_response_latency_seconds_count[%duration]))"
	AvgDatabaseLatencyQuery         = "sum(rate(api_database_latency_seconds_sum[%duration])) / sum(rate(api_database_latency_seconds_count[%duration]))"
	TotalErrorResponseQuery         = "round(increase(api_response_error_counter[%duration]))-1"
	SimulatorPerMinuteQuery         = "rate(simulator_simulation_duration_seconds_count[%duration]) * 60"
	AvgSimulationDurationQuery      = "sum(rate(simulator_simulation_duration_seconds_sum[%duration])) / sum(rate(simulator_simulation_duration_seconds_count[%duration]))"
	AvgSimulationDistanceQuery      = "sum(rate(simulator_simulation_distance_meters_sum[%duration])) / sum(rate(simulator_simulation_distance_meters_count[%duration]))"
	TotalSuccessfulSimulationsQuery = "round(increase(simulator_simulation_successful_counter[%duration]))-1"
	TotalFailedSimulationQuery      = "round(increase(simulator_simulation_failed_counter[%duration]))-1"
)

var MetricQueryMap = map[string]string{
	AvgResponseLatency:         AvgResponseLatencyQuery,
	AvgDatabaseLatency:         AvgDatabaseLatencyQuery,
	TotalErrorResponse:         TotalErrorResponseQuery,
	SimulationsPerMinute:       SimulatorPerMinuteQuery,
	AvgSimulationDuration:      AvgSimulationDurationQuery,
	AvgSimulationDistance:      AvgSimulationDistanceQuery,
	TotalSuccessfulSimulations: TotalSuccessfulSimulationsQuery,
	TotalFailedSimulations:     TotalFailedSimulationQuery,
}

var LowerIsBetterMetricList = map[string]struct{}{
	AvgResponseLatency:     {},
	AvgDatabaseLatency:     {},
	TotalErrorResponse:     {},
	TotalFailedSimulations: {},
}

// Define a struct for thresholds
type Thresholds struct {
	Good float64
	Fair float64
	// No need for "Bad" since it's the default case
}

func GetMetric(metricName string, d time.Duration) (*openapi.Metric, error) {
	resp, exists := cache.AppCache.Get(cache.GenerateKey(metricName, d))
	if !exists {
		prometheus, err := observibility.QueryPrometheus(MetricQueryMap[metricName], d)
		if err != nil {
			return nil, fmt.Errorf("failed querying prometheus: %w", err)
		}

		if len(prometheus.Data.Result) == 0 {
			return &openapi.Metric{}, nil
		}

		startDate := time.Now().AddDate(0, 0, -15).Unix()
		endDate := time.Now().Unix()
		maxMetric15Days, err := observibility.RangeQueryPrometheus(MetricQueryMap[metricName], startDate, endDate, d)
		if err != nil {
			return nil, fmt.Errorf("failed querying prometheus: %w", err)
		}
		var metricHealth = "u"
		if len(maxMetric15Days.Data.Result) > 0 {
			metricHealth = getMetricHealth(metricName, prometheus.Data.Result[0].Values[1], maxMetric15Days.Data.Result[0].Values)
		}

		metric := openapi.Metric{MetricType: metricName, MetricTime: prometheus.Data.Result[0].Values[0].(float64),
			MetricHealth: metricHealth, MetricValue: roundToFourDecimalPlaces(prometheus.Data.Result[0].Values[1])}
		err = cache.AppCache.Add(cache.GenerateKey(metricName, d), &metric, cache.AppCacheExpiration)
		if err != nil {
			return nil, apperror.NewAppError(http.StatusInternalServerError, "Error while trying to add to cache", err)
		}
		return &metric, nil
	}

	return resp.(*openapi.Metric), nil
}
func roundToFourDecimalPlaces(n interface{}) string {
	floatVal, err := strconv.ParseFloat(n.(string), 64)
	if err != nil {
		return "0"
	}

	floatVal = math.Round(floatVal*1e4) / 1e4
	if floatVal < 0 {
		floatVal = 0
	}

	return strconv.FormatFloat(floatVal, 'f', -1, 64)
}

func getMetricHealth(metricName string, currentValue interface{}, max15DaysValue [][]interface{}) string {
	cValue, err := strconv.ParseFloat(currentValue.(string), 64)
	if err != nil {
		return "u"
	}

	values := extractHistoricalValues(max15DaysValue)

	ucl, lcl, cl := calculateControlLimits(values)

	lowerIsBetteThreshold := 1.1
	higherIsBetterThreshold := .9
	_, isLowerBetter := LowerIsBetterMetricList[metricName]
	status := "u"
	if isLowerBetter {
		if cValue <= (cl * lowerIsBetteThreshold) {
			status = "g"
		} else if cValue > (cl*lowerIsBetteThreshold) && cValue < ucl {
			status = "o"
		} else {
			status = "r"
		}
	} else {
		if cValue >= (cl * higherIsBetterThreshold) {
			status = "g"
		} else if cValue < (cl*higherIsBetterThreshold) && cValue > lcl {
			status = "o"
		} else {
			status = "r"
		}
	}
	return status
}

func extractHistoricalValues(max15DaysValue [][]interface{}) []float64 {
	values := make([]float64, 0, len(max15DaysValue))
	for _, v := range max15DaysValue {
		stringValue := v[1]
		value, _ := strconv.ParseFloat(stringValue.(string), 64)

		if !math.IsNaN(value) {
			if value < 0 {
				value = 0
			}
			values = append(values, value)
		}
	}
	return values
}
func calculateControlLimits(values []float64) (ucl, lcl, cl float64) {
	// Calculate the mean (CL)
	cl = stat.Mean(values, nil)

	// Calculate the standard deviation
	sd := stat.StdDev(values, nil)

	// Assuming k = 3 for control limits
	k := 3.0

	// Calculate Upper Control Limit (UCL) and Lower Control Limit (LCL)
	ucl = cl + k*sd
	lcl = cl - k*sd
	if lcl < 0 {
		lcl = 0
	}

	return ucl, lcl, cl
}
