package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"simulator/internal/datastructures"
	"simulator/internal/dto"
	"simulator/internal/service"
	"simulator/internal/util"
	"simulator/observibility"
	"time"
)

var newJourneyChannel = make(chan *dto.Journey)
var WebsocketSendDataUrl string

func init() {
	//env := os.Getenv("ENV")
	//if env == "production" {
	//	WebsocketSendDataUrl = "ws://localhost:8081/v1/ws/simulation/path"
	//} else {
	//	WebsocketSendDataUrl = "ws://localhost:8081/v1/ws/simulation/path"
	//}
}
func main() {

	datastructures.LoadGeoJSONGraph()

	// Create a channel to communicate the simulated timesimulator
	timeChannel := make(chan time.Time)

	websocketManager := util.NewWebSocketManager(WebsocketSendDataUrl, 5*time.Second, 4)
	err := websocketManager.Connect()
	if err != nil {
		log.Fatal("Could not connect to server:", err)
	}
	defer websocketManager.Close()
	// Start the timesimulator simulation in a separate goroutine
	go service.SimulateTime(timeChannel)
	// Start the API calls in a separate goroutine
	go service.ScheduleJourneyAPICalls(timeChannel)
	// Start the Polling for valid Journeys
	go service.PollJourneys(newJourneyChannel)
	// Start the car movement
	go service.RunSimulation(newJourneyChannel, websocketManager)

	// Setup signal handling
	reg := prometheus.NewRegistry()
	m := observibility.GetMetrics()
	m.Register(reg)
	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg})
	http.Handle("/metrics", promHandler)
	err = http.ListenAndServe(":8080", nil)
	// Exit
}
