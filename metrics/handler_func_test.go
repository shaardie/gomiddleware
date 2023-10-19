package metrics

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
)

func TestWithLoggingWithoutConfig(t *testing.T) {
	before, err := testutil.GatherAndCount(prometheus.DefaultGatherer)
	assert.Nil(t, err)

	s := httptest.NewServer(
		WithMetrics(
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("OK"))
			},
			Config{
				Prefix: "test",
			},
		),
	)
	defer s.Close()

	_, err = http.Get(s.URL)
	assert.Nil(t, err)

	after, err := testutil.GatherAndCount(prometheus.DefaultGatherer)
	assert.Nil(t, err)
	assert.True(t, after-before > 0)
}
