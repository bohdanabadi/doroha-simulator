package service

import (
	"container/heap"
	"errors"
	"fmt"
	"github.com/wk8/go-ordered-map/v2"
	"math"
	"simulator/internal/datastructures"
	"simulator/internal/dto"
)

func AStar(journey *dto.Journey) (*orderedmap.OrderedMap[string, dto.PointNode], float64, error) {
	start := &dto.PointNode{X: datastructures.RoundToDecimal(journey.StartingPointNode.X, 9), Y: datastructures.RoundToDecimal(journey.StartingPointNode.Y, 9)}
	end := &dto.PointNode{X: datastructures.RoundToDecimal(journey.EndingPointNode.X, 9), Y: datastructures.RoundToDecimal(journey.EndingPointNode.Y, 9)}
	costFromStart := make(map[dto.PointNode]float64) // Cost from start to each node
	sumOfAllCosts := make(map[dto.PointNode]float64) // Total cost (gScore + heuristic) to each node
	//
	for _, keys := range datastructures.RoadMapGraph {
		for _, key := range keys {
			costFromStart[key] = math.Inf(1) // gScore
			sumOfAllCosts[key] = math.Inf(1) // fScore
		}
	}
	costFromStart[*start] = 0
	sumOfAllCosts[*start] = start.Heuristic(*end)

	openSet := &datastructures.PointNodeWrapperHeap{}
	heap.Init(openSet)
	nodeWrapper := &datastructures.PointNodeWrapper{}
	nodeWrapper.SetValues(start, costFromStart[*start], sumOfAllCosts[*start], costFromStart[*start]+sumOfAllCosts[*start], -1, nil)
	heap.Push(openSet, nodeWrapper)

	closedSet := make(map[dto.PointNode]bool)

	for openSet.Len() > 0 {
		current := heap.Pop(openSet).(*datastructures.PointNodeWrapper)

		if current.Node.Y == end.Y && current.Node.X == end.X {
			//Sum the total cost of the path
			return reconstructPath(current), current.TotalCostFromStart, nil
		}

		closedSet[*current.Node] = true

		for _, n := range datastructures.RoadMapGraph[*current.Node] {
			neighbor := n
			if _, exists := closedSet[neighbor]; exists {
				continue
			}
			edge := dto.Edge{From: *current.Node, To: neighbor}
			actualCost, exists := datastructures.RoadMapEdgeCostGraph[edge]
			if !exists {
				fmt.Println("Edge Cost doest not exist")
			}

			tentativeCostFromStart := costFromStart[*current.Node] + actualCost
			if tentativeCostFromStart < costFromStart[neighbor] {
				costFromStart[neighbor] = tentativeCostFromStart
				sumOfAllCosts[neighbor] = tentativeCostFromStart + neighbor.Heuristic(*end)

				neighborWrapper := datastructures.PointNodeWrapper{}
				neighborWrapper.SetValues(&neighbor,
					costFromStart[neighbor],
					neighbor.Heuristic(*end),
					sumOfAllCosts[neighbor],
					-1,
					current)

				heap.Push(openSet, &neighborWrapper)
			}
		}
	}
	return nil, 0, errors.New("path not found")
}

// Reconstruct the optimal path by following the cameFrom map from the goal node back to the start
func reconstructPath(nodeWrapper *datastructures.PointNodeWrapper) *orderedmap.OrderedMap[string, dto.PointNode] {
	oMap := orderedmap.New[string, dto.PointNode]()
	var stackKeys []string
	for nodeWrapper != nil {
		key := fmt.Sprintf("%f,%f", nodeWrapper.Node.X, nodeWrapper.Node.Y) // Assuming X and Y are coordinates of the node
		oMap.Set(key, *nodeWrapper.Node)
		stackKeys = append(stackKeys, key)
		nodeWrapper = nodeWrapper.CameFrom

		//path = append(path, *nodeWrapper.Node)
		//nodeWrapper = nodeWrapper.CameFrom
	}
	return reverse(*oMap, stackKeys)
}
func reverse(oMap orderedmap.OrderedMap[string, dto.PointNode], stackKeys []string) *orderedmap.OrderedMap[string, dto.PointNode] {
	reversedMap := orderedmap.New[string, dto.PointNode]()
	for len(stackKeys) > 0 {
		// Pop the top key from the stack
		topIndex := len(stackKeys) - 1
		key := stackKeys[topIndex]
		stackKeys = stackKeys[:topIndex] // update stack

		value, exists := oMap.Get(key)
		if exists {
			reversedMap.Set(key, value)
		}
	}
	return reversedMap
}
