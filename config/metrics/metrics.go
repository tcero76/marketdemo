package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	HTTPDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duración de las peticiones HTTP.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "route", "status"},
	)

	HTTPRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Número total de peticiones.",
		},
		[]string{"method", "route", "status"},
	)

	HTTPErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_errors_total",
			Help: "Número total de errores.",
		},
		[]string{"method", "route", "status"},
	)

	HTTPInFlight = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "http_requests_in_flight",
			Help: "Peticiones actualmente en ejecución.",
		},
	)

	ExternalDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "external_request_duration_seconds",
			Help:    "Duración de llamadas a servicios externos.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"service"},
	)
)

func Register() {
	prometheus.MustRegister(
		HTTPDuration,
		HTTPRequests,
		HTTPErrors,
		HTTPInFlight,
		ExternalDuration,
	)
}