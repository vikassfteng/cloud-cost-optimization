package monitoring

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Duration of HTTP requests.",
		Buckets: prometheus.DefBuckets,
	})
)

func InitPrometheus() {
	prometheus.MustRegister(requestDuration)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9090", nil)
}
