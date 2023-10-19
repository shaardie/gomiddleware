package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	DefaultRequestDurationBuckets  = prometheus.ExponentialBuckets(0.005, 2, 10)
	DefaultRequestBodySizeBuckets  = prometheus.ExponentialBuckets(8, 2, 10)
	DefaultResponseBodySizeBuckets = prometheus.ExponentialBuckets(8, 2, 10)
)

type Config struct {
	Prefix                  string
	RequestDurationBuckets  []float64
	RequestBodySizeBuckets  []float64
	ResponseBodySizeBuckets []float64
}

func (cfg Config) WithDefaults() Config {
	if len(cfg.RequestDurationBuckets) == 0 {
		cfg.RequestDurationBuckets = DefaultRequestDurationBuckets
	}
	if len(cfg.RequestBodySizeBuckets) == 0 {
		cfg.RequestBodySizeBuckets = DefaultRequestBodySizeBuckets
	}
	if len(cfg.ResponseBodySizeBuckets) == 0 {
		cfg.RequestBodySizeBuckets = DefaultResponseBodySizeBuckets
	}
	return cfg
}
