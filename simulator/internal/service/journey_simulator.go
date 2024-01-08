package service

import (
	"encoding/json"
	"fmt"
	"log"
	"simulator/internal/datastructures"
	"simulator/internal/dto"
	"simulator/internal/util"
	"time"
)

func RunSimulation(newJourneyChannel chan *dto.Journey, websocketManager *util.WebSocketManager) {
	var activeJourneys []*dto.Journey
	simulationTick := time.NewTicker(time.Millisecond * 400) // adjust duration as needed

	for {
		select {
		case newJourney := <-newJourneyChannel:
			activeJourneys = append(activeJourneys, newJourney)
		case <-simulationTick.C:
			// Simulate movement for all active journeys
			var movedJourneys []*dto.Journey // Assuming Journey is the type of the journeys
			for _, j := range activeJourneys {
				// Update journey position...

				processJourney(j)
				message, err := json.Marshal(j)
				if err != nil {
					log.Println("Failed to marshal buffer:", err)
					return
				}
				err = websocketManager.Send(message)
				if err != nil {
					log.Println("Failed to send via websocket, error stack:", err)
				} else if reachedEnd(j) {
					// Skip appending this journey to updatedJourneys
					//err := updateStatusForJourney([]int32{j.Id}, string(dto.Finished))
					if err != nil {
						log.Println("Failed to update journey status:", err)
						return
					}
					log.Printf("Journey with id %d has reach its end\n", j.Id)
					continue
				}
				movedJourneys = append(movedJourneys, j)
			}
			// Remove completed journeys or update list as needed...
			activeJourneys = movedJourneys
		}
	}
}

func processJourney(j *dto.Journey) {

	updateVehiclePosition(j)
	updateProgress(j)
	checkAndUpdateStatus(j)
}
func updateVehiclePosition(j *dto.Journey) {
	//TODO Movement factor should be calculated based on the speed of the vehicle
	key := fmt.Sprintf("%f,%f", j.CurrentPointNode.X, j.CurrentPointNode.Y)
	nextCurrentPair := j.Path.GetPair(key).Next()
	movementFactor := calculateMovementFactor(&j.CurrentPointNode, &nextCurrentPair.Value, 200)
	j.AccumulatedMovement += movementFactor

	if nextCurrentPair != nil && j.AccumulatedMovement >= 1 {
		moveVehicleToNextPoint(j, nextCurrentPair.Value)
		updateCost(j)
		j.AccumulatedMovement -= 1
	}
}

func moveVehicleToNextPoint(j *dto.Journey, nextNode dto.PointNode) {
	j.PrevPointNode = j.CurrentPointNode
	j.CurrentPointNode = nextNode
}

func updateCost(j *dto.Journey) {
	edge := dto.Edge{From: j.PrevPointNode, To: j.CurrentPointNode}
	j.CostFromTheStart += datastructures.RoadMapEdgeCostGraph[edge]
}

func updateProgress(j *dto.Journey) {
	j.Progress = float32(j.CostFromTheStart / j.TotalTripCost)
}

func checkAndUpdateStatus(j *dto.Journey) {
	switch j.Status {
	case dto.InQueue:
		j.Status = dto.InProgress
	case dto.InProgress:
		if reachedEnd(j) {
			j.Status = dto.Finished
			log.Printf("Journey is done ID %d\n", j.Id)
		}
	}
}

func reachedEnd(j *dto.Journey) bool {
	return j.CurrentPointNode.X == j.EndingPointNode.X && j.CurrentPointNode.Y == j.EndingPointNode.Y
}
