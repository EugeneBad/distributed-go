package monitoring

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func RuntimeMetricExporter() {
	http.Handle("/metrics", promhttp.Handler())
	_ = http.ListenAndServe(":2112", nil)
}
