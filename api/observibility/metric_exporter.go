package observibility

import (
	"github.com/bohdanabadi/Traffic-Simulation/api/db"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"gorm.io/gorm"
	"sync"
	"time"
)

const AppNameSpace = "api"
const RequestLatencyMetricName = "response_latency_seconds"
const QueryLatencyMetricName = "database_latency_seconds"

const ErrorCounterMetricName = "response_error_counter"
const DurationMetricHelp = "Latency for requests"
const QueryLatencyMetricHelp = "Query Latency for database"

const ErrorCounterMetricHelp = "Error response counter"

type Metrics struct {
	RequestLatency prometheus.Summary
	DBQueryLatency prometheus.Summary
	ErrorCounter   prometheus.Counter
}

var (
	instance *Metrics
	once     sync.Once
)

func newMetrics() *Metrics {
	m := &Metrics{
		RequestLatency: prometheus.NewSummary(prometheus.SummaryOpts{
			Namespace: AppNameSpace,
			Name:      RequestLatencyMetricName,
			Help:      DurationMetricHelp,
		}),
		DBQueryLatency: prometheus.NewSummary(prometheus.SummaryOpts{
			Namespace: AppNameSpace,
			Name:      QueryLatencyMetricName,
			Help:      QueryLatencyMetricHelp,
		}),
		ErrorCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: AppNameSpace,
			Name:      ErrorCounterMetricName,
			Help:      ErrorCounterMetricHelp,
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
	reg.MustRegister(m.RequestLatency, m.DBQueryLatency, m.ErrorCounter)
}

func (m *Metrics) RequestDurationMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		startTime := time.Now()
		context.Next()
		duration := time.Since(startTime)
		m.RequestLatency.Observe(duration.Seconds())
	}
}

func (m *Metrics) ExecuteGormQueryWithLatency(query func(*gorm.DB) *gorm.DB) *gorm.DB {
	startTime := time.Now()
	result := query(db.DB)

	err := result.Error

	if err == nil {
		duration := time.Since(startTime)
		m.DBQueryLatency.Observe(duration.Seconds())
	}
	return result
}

func (m *Metrics) LogErrorCounter() {
	m.ErrorCounter.Inc()
}
