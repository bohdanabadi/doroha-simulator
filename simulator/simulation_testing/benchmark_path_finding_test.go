package simulation_testing

import (
	"simulator/internal/datastructures"
	"simulator/internal/dto"
	"simulator/internal/service"
	"testing"
)

func BenchmarkAStar(b *testing.B) {
	datastructures.LoadGeoJSONGraph()
	journey := dto.Journey{StartingPointX: 30.542297,
		StartingPointY: 50.43449,
		EndingPointX:   30.545498,
		EndingPointY:   50.42468,
		Distance:       1114.8798,
	}

	b.ResetTimer() // Resets the timer to exclude the setup time

	for i := 0; i < b.N; i++ {
		service.AStar(&journey)
	}
}

func BenchmarkDijkstra(b *testing.B) {
	datastructures.LoadGeoJSONGraph()
	journey := dto.Journey{StartingPointX: 30.542297,
		StartingPointY: 50.43449,
		EndingPointX:   30.545498,
		EndingPointY:   50.42468,
		Distance:       1114.8798,
	}

	b.ResetTimer() // Resets the timer to exclude the setup time

	for i := 0; i < b.N; i++ {
		service.Dijkstra(&journey)
	}
}
