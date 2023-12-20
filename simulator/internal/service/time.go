package service

import (
	"time"
)

func SimulateTime(timeChannel chan<- time.Time) {
	// Define the real-world duration of the simulation
	simulationDuration := 60 * time.Minute

	// Define the simulated duration that the simulation represents
	simulatedDuration := 24 * time.Hour

	// Calculate the ratio between the simulated duration and the real-world duration
	ratio := simulatedDuration.Seconds() / simulationDuration.Seconds()

	// Define the tick interval for the simulation
	tickInterval := 2 * time.Second

	// Create a ticker
	ticker := time.NewTicker(tickInterval)
	defer ticker.Stop()

	// Start the simulation
	start := time.Now()
	for t := range ticker.C {
		// Calculate the simulated timesimulator
		simulatedTime := start.Add(time.Duration(float64(t.Sub(start)) * ratio))

		// Send the simulated timesimulator over the channel
		timeChannel <- simulatedTime
	}
}
