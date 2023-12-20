package service

import (
	"math"
	"simulator/internal/datastructures"
	"simulator/internal/dto"
)

func calculateMovementFactor(currentPoint, nextPoint *dto.PointNode, cycleTimeMs int) float64 {
	edgeCost := getEdgeCost(currentPoint, nextPoint)

	// Assuming effortPerCycle is a predetermined constant that represents
	// how much 'effort' a vehicle can exert in one cycle.
	// This constant can be derived based on system averages or empirical data.
	effortPerCycle := calculateEffortPerCycle(cycleTimeMs)

	movementFactor := effortPerCycle / edgeCost

	// Ensuring the movement factor is not more than 1
	if movementFactor > 1 {
		movementFactor = 1
	}

	// Check if the fractional part of the movement factor is close to 1
	if movementFactor-math.Floor(movementFactor) >= 0.9 && movementFactor-math.Floor(movementFactor) <= 1.0 {
		movementFactor = math.Ceil(movementFactor)
	}

	return movementFactor
}

func getEdgeCost(currentNode, nextNode *dto.PointNode) float64 {
	return datastructures.RoadMapEdgeCostGraph[dto.Edge{From: *currentNode, To: *nextNode}]
}

func calculateEffortPerCycle(cycleTimeMs int) float64 {
	return 1.0 / float64(cycleTimeMs)
}
