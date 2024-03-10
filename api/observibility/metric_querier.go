package observibility

import (
	"encoding/json"
	"fmt"
	"github.com/bohdanabadi/Traffic-Simulation/api/api/apperror"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const PrometheusURL = "http://localhost:9090"
const PrometheusEndpointQuery = "/api/v1/query?"
const PrometheusEndpointRangeQuery = "/api/v1/query_range?"

type PrometheusQueryResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
			} `json:"metric"`
			Values []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

type PrometheusQueryRangeResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
			} `json:"metric"`
			Values [][]interface{} `json:"values"`
		} `json:"result"`
	} `json:"data"`
}

func QueryPrometheus(metricQuery string, d time.Duration) (PrometheusQueryResponse, error) {
	query := strings.Replace(metricQuery, "%duration", d.String(), -1)

	step := "1m"

	queryURL := buildPrometheusQueryURL(PrometheusEndpointQuery, query, step, "", "")

	resp, err := http.Get(queryURL)
	if err != nil {
		return PrometheusQueryResponse{}, apperror.NewAppError(500, "Error while trying to HTTP Get prometheus", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	var promResponse PrometheusQueryResponse
	err = json.Unmarshal(body, &promResponse)
	if err != nil {
		return PrometheusQueryResponse{}, apperror.NewAppError(500, "Error while trying Unmarshal query prometheus response", err)
	}

	return promResponse, nil
}

func RangeQueryPrometheus(metricQuery string, start, end int64, d time.Duration) (PrometheusQueryRangeResponse, error) {
	query := strings.Replace(metricQuery, "%duration", d.String(), -1)

	step := durationToStepMapper(d.String())

	queryURL := buildPrometheusQueryURL(PrometheusEndpointRangeQuery, query, step, strconv.FormatInt(start, 10), strconv.FormatInt(end, 10))

	resp, err := http.Get(queryURL)
	if err != nil {
		return PrometheusQueryRangeResponse{}, apperror.NewAppError(500, "Error while trying to HTTP Get prometheus", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	var promResponse PrometheusQueryRangeResponse
	err = json.Unmarshal(body, &promResponse)
	if err != nil {
		return PrometheusQueryRangeResponse{}, apperror.NewAppError(500, "Error while trying Unmarshal range query prometheus response", err)
	}
	return promResponse, nil
}

func buildPrometheusQueryURL(endpoint, query, step, start, end string) string {
	queryParam := url.Values{}
	queryParam.Set("query", query)
	queryParam.Set("step", step)
	if len(start) > 0 {
		queryParam.Set("start", start)
	}
	if len(end) > 0 {
		queryParam.Set("end", end)
	}

	return fmt.Sprintf("%s"+endpoint+"%s", PrometheusURL, queryParam.Encode())
}

func durationToStepMapper(duration string) string {
	switch duration {
	case "1h":
		return "30m"
	case "6h":
		return "3h"
	case "24h":
		return "12h"
	case "7d":
		return "84h"

	}
	return "1h"
}
