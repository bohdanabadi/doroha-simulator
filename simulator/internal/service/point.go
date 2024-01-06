package service

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"simulator/internal/datastructures"
	"simulator/internal/dto"
)

func GetPossibleJourneyPoints(ch chan<- *dto.JourneyPoints) {
	request := gorequest.New()
	respDto := &dto.JourneyPoints{}
	_, _, errs := request.Get("http://localhost:8081/v1/points/random-pair").EndStruct(respDto)
	if errs != nil {
		fmt.Printf("Error making GET request to get POINTS: %v\n", errs)
		return
	}
	ch <- respDto
}

func PathExists(start, end dto.PointNode) bool {
	fmt.Printf("Validating our point before persisting : %f %v\n", start, end)
	visited := map[dto.PointNode]bool{}
	queue := []dto.PointNode{start}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == end {
			return true
		}

		for _, neighbor := range datastructures.RoadMapGraph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return false
}
