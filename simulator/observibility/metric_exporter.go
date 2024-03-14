package observibility

import (
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"sync"
	"time"
)

const AppNameSpace = "simulator"
const SimulationDurationMetricName = "simulation_duration_seconds"
const SimulationDistanceMetricName = "simulation_distance_meters"
const SimulationSuccessfulCounter = "simulation_successful_counter"
const SimulationFailedCounter = "simulation_failed_counter"
const SimulationMetricHelp = "Duration of journey simulations"
const SimulationDistanceHelp = "Journey distance in meters"
const SimulationSuccessfulCounterHelp = "Number of successful simulated journey"
const SimulationFailedCounterHelp = "Number of failed simulated journey"

type Metrics struct {
	SimulationDuration                   prometheus.Summary
	SimulationDistance                   prometheus.Summary
	SuccessfullySimulatedJourneysCounter prometheus.Counter
	FailedSimulatedJourneyCounter        prometheus.Counter
}

var (
	instance *Metrics
	once     sync.Once
)

func newMetrics() *Metrics {
	m := &Metrics{
		SimulationDuration: prometheus.NewSummary(prometheus.SummaryOpts{
			Namespace: AppNameSpace,
			Name:      SimulationDurationMetricName,
			Help:      SimulationMetricHelp,
		}),
		SimulationDistance: prometheus.NewSummary(prometheus.SummaryOpts{
			Namespace: AppNameSpace,
			Name:      SimulationDistanceMetricName,
			Help:      SimulationDistanceHelp,
		}),
		SuccessfullySimulatedJourneysCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: AppNameSpace,
			Name:      SimulationSuccessfulCounter,
			Help:      SimulationSuccessfulCounterHelp,
		}),
		FailedSimulatedJourneyCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: AppNameSpace,
			Name:      SimulationFailedCounter,
			Help:      SimulationFailedCounterHelp,
		}),
	}
	return m
}

func GetMetrics() *Metrics {
	once.Do(func() {
		instance = newMetrics()
	})
	return instance
}

func (m *Metrics) Register(reg prometheus.Registerer) {
	reg.MustRegister(m.SimulationDuration, m.SimulationDistance, m.SuccessfullySimulatedJourneysCounter,
		m.FailedSimulatedJourneyCounter)
}
func (m *Metrics) LogTimeToFinishSimulation(t time.Duration) {
	m.SimulationDuration.Observe(t.Seconds())
}

func (m *Metrics) LogJourneyDistance(d float32) {
	m.SimulationDistance.Observe(float64(d))
}

func (m *Metrics) LogSuccessfulSimulatedJourney() {
	log.Printf("Counter increasing success counter")
	m.SuccessfullySimulatedJourneysCounter.Inc()
}

func (m *Metrics) LogFailedSimulatedJourney() {
	m.FailedSimulatedJourneyCounter.Inc()
}
