package bootstrap

import (
	prom "github.com/devexps/go-micro/metrics/prometheus/v2"
	"github.com/devexps/go-micro/v2/middleware/metrics"
	"github.com/devexps/go-micro/v2/transport/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	_metricRequests *prometheus.CounterVec
	_metricSeconds  *prometheus.HistogramVec
)

// NewMetrics create and register prometheus' metrics
func NewMetrics() error {
	_metricRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "client",
		Subsystem: "requests",
		Name:      "code_total",
		Help:      "The total number of processed requests",
	}, []string{"kind", "operation", "code", "reason"})

	_metricSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "server",
		Subsystem: "requests",
		Name:      "duration_sec",
		Help:      "server requests duration(sec).",
		Buckets:   []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.250, 0.5, 1},
	}, []string{"kind", "operation"})

	prometheus.MustRegister(_metricRequests, _metricSeconds)
	return nil
}

func handleMetrics(srv *http.Server) {
	srv.Handle("/metrics", promhttp.Handler())
}

func withMetricRequests() metrics.Option {
	return metrics.WithRequests(prom.NewCounter(_metricRequests))
}

func withMetricHistogram() metrics.Option {
	return metrics.WithSeconds(prom.NewHistogram(_metricSeconds))
}
