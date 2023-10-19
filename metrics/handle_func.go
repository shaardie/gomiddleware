package metrics

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	rw "github.com/shaardie/gomiddleware/pkg/responsewriter"
)

func WithMetrics(h http.HandlerFunc, cfg Config) http.HandlerFunc {
	cfg = cfg.WithDefaults()
	requestsTotal := promauto.NewCounterVec(prometheus.CounterOpts{
		Name: fmt.Sprintf("%v_requests_total", cfg.Prefix),
		Help: "The total number of processed requests",
	}, []string{
		"method",
		"status_code",
	})
	requestsDuration := promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    fmt.Sprintf("%v_requests_duration_seconds", cfg.Prefix),
		Help:    "Duration of the requests",
		Buckets: cfg.RequestDurationBuckets,
	})
	requestBodySize := promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    fmt.Sprintf("%v_requests_body_bytes", cfg.Prefix),
		Help:    "Total number of bytes in requests bodies",
		Buckets: cfg.RequestBodySizeBuckets,
	})
	responseBodySize := promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    fmt.Sprintf("%v_response_body_bytes", cfg.Prefix),
		Help:    "Total number of bytes in response bodies",
		Buckets: cfg.ResponseBodySizeBuckets,
	})
	return func(w http.ResponseWriter, r *http.Request) {
		cw := &responsewriter{
			ResponseWriter: rw.ResponseWriter{
				ResponseWriter: w,
				StatusCode:     http.StatusOK,
			},
		}
		before := time.Now()
		h(cw, r)
		duration := time.Since(before)
		promLabels := prometheus.Labels{
			"method":      r.Method,
			"status_code": fmt.Sprint(cw.StatusCode),
		}
		requestsDuration.Observe(duration.Seconds())
		requestsTotal.With(promLabels).Inc()
		requestBodySize.Observe(float64(r.ContentLength))
		responseBodySize.Observe(float64(cw.size))
	}
}
