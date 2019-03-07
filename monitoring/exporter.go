package monitoring

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type ReadingCounter struct {
	prometheus.CounterVec
}
type ReadingCounterInterface interface {
	Increment(string, string)
}

func NewReadingCounter() ReadingCounterInterface {
	var rc ReadingCounter

	rc.CounterVec = *promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "coordinator_reading_count",
		Help: "The total number of processed readings by the coordinator"},
		[]string{"layer", "sensor"})

	return &rc
}

func (rc *ReadingCounter) Increment(layer, sensor string) {
	rc.With(prometheus.Labels{"sensor": sensor})
}

func MetricExporter() {
	http.Handle("/metrics", promhttp.Handler())
	_ = http.ListenAndServe(":2112", nil)
}
