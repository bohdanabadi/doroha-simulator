package service

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"log"
	"math/rand"
	"simulator/internal/dto"
	"sync"
	"time"
)

func ScheduleJourneyAPICalls(timeChannel <-chan time.Time) {
	for simulatedTime := range timeChannel {
		// Determine the frequency of the API calls based on the simulated timesimulator
		hour := simulatedTime.Hour()
		var numCalls int
		// Make API calls more frequently
		if hour >= 7 && hour < 10 || hour > 16 && hour < 18 {
			fmt.Printf("Making API calls more frequently at hour: %d\n", hour)
			numCalls = 0 + rand.Intn(2)
		} else {
			// Make API calls less frequently
			fmt.Printf("Making API calls less frequently: %d\n\n", hour)
			numCalls = 0 + rand.Intn(2)
		}
		fmt.Printf("Number of calls: %d\n", numCalls)
		retrieveAndValidateJourneys(numCalls)

	}
}

func retrieveAndValidateJourneys(numCalls int) {
	var wg sync.WaitGroup
	ch := make(chan *dto.JourneyPoints, numCalls)
	for i := 0; i < numCalls; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			GetPossibleJourneyPoints(ch)
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	var journeys []dto.Journey
	for point := range ch {
		start := point.StartingPoint
		end := point.EndingPoint
		valid := PathExists(start, end)
		if valid {
			journeys = append(journeys, point.ToJourneyPointsDTO(time.Now().Format(time.RFC3339), string(dto.InQueue), 0))
		}
	}
	if len(journeys) > 0 {
		persistJourney(&journeys)
	}
}
func persistJourney(journey *[]dto.Journey) {
	request := gorequest.New()
	resp, _, errs := request.Post("http://localhost:8081/v1/journeys").Send(journey).End()
	if errs != nil {
		log.Fatalf("Request failed: %v", errs)
	}

	if resp.StatusCode != 201 {
		log.Printf("Unexpected status code: %d", resp.StatusCode)
	}
	//log.Printf("Response Code: %s\n", resp.StatusCode)
}

func PollJourneys(newJourneyChannel chan<- *dto.Journey) {
	interval := 5 * time.Second
	ticker := time.NewTicker(interval)
	for range ticker.C {
		request := gorequest.New()
		var journeyResponse dto.JourneysResponse
		resp, _, errs := request.Get("http://localhost:8081/v1/journeys?status=IN QUEUE").EndStruct(&journeyResponse)
		if errs != nil || resp.StatusCode != 200 {
			fmt.Printf("Error making GET request to get Valid Journeys: %v\n", errs)
		}
		var ids []int32
		for _, journey := range journeyResponse.Journeys {
			ids = append(ids, journey.Id)
		}
		if err := updateStatusForJourney(ids, string(dto.InProgress)); err != nil {
			fmt.Printf("Error making Updating Journeys Status: %v\n", errs)
		} else {
			for _, journey := range journeyResponse.Journeys {
				go initiateActiveJourney(journey, newJourneyChannel)
			}
		}
	}
}

func updateStatusForJourney(journeyIds []int32, newStatus string) error {
	request := gorequest.New()
	journeyToPatch := &dto.JourneyListStatusUpdate{Ids: journeyIds, Status: newStatus}
	resp, _, errs := request.Patch("http://localhost:8081/v1/journeys/status").Send(journeyToPatch).End()
	if errs != nil || resp.StatusCode != 204 {
		return fmt.Errorf("Error making PATCH request Update Journeys Status: %v\n", errs)
	}
	return nil
}

// InitiateActiveJourney initiates the active journey by creating a new journey with the path and total cost
// of the inQueueJourney and sending it to the newJourneyChannel to be added to the active journeys list
func initiateActiveJourney(inQueueJourney dto.Journey, newJourneyChannel chan<- *dto.Journey) {
	//TODO start here
	path, totalCost, err := AStar(&inQueueJourney) // Define the path of the inQueueJourney

	if err != nil {
		fmt.Printf("Not good another bug ? ")
	}

	newJourney := dto.NewJourney(
		dto.WithId(inQueueJourney.Id),
		dto.WithStartingPointNode(inQueueJourney.StartingPointNode),
		dto.WithEndingPointNode(inQueueJourney.EndingPointNode),
		dto.WithCurrentPointNode(inQueueJourney.StartingPointNode),
		dto.WithDistance(inQueueJourney.Distance),
		dto.WithPath(*path),
		dto.WithTotalTripCost(totalCost),
		dto.WithStatus(inQueueJourney.Status),
		dto.WithDateCreate(time.Now().Format(time.RFC3339)),
	)
	newJourneyChannel <- newJourney

}
