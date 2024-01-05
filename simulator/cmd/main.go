package main

import (
	"fmt"
	"log"
	"os"
	"simulator/internal/datastructures"
	"simulator/internal/dto"
	"simulator/internal/service"
	"simulator/internal/util"
	"time"
)

var newJourneyChannel = make(chan *dto.Journey)
var WebsocketSendDataUrl string

func init() {
	env := os.Getenv("ENV")
	if env == "production" {
		WebsocketSendDataUrl = "ws://localhost:8081/v1/ws/simulation/path"
	} else {
		WebsocketSendDataUrl = "ws://localhost:8080/v1/ws/simulation/path"
	}
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
	log.Printf("About to defer websocket")
	defer websocketManager.Close()
	log.Printf("Defered success")
	// Start the timesimulator simulation in a separate goroutine
	go service.SimulateTime(timeChannel)
	log.Printf("Simulate Time after")
	// Start the API calls in a separate goroutine
	go service.ScheduleJourneyAPICalls(timeChannel)
	// Start the Polling for valid Journeys
	go service.PollJourneys(newJourneyChannel)
	// Start the car movement
	go service.RunSimulation(newJourneyChannel, websocketManager)
	// Wait for a user input before exiting
	fmt.Scanln()
}
