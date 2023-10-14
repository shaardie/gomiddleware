package log

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-logr/logr"
	"github.com/stretchr/testify/assert"
	"github.com/tonglil/buflogr"
)

func TestWithLoggingWithoutConfig(t *testing.T) {
	var buf bytes.Buffer
	s := httptest.NewServer(
		WithLogging(func(w http.ResponseWriter, r *http.Request) {}, buflogr.NewWithBuffer(&buf), Config{}),
	)
	defer s.Close()
	assert.Equal(t, buf.String(), "")
}

func TestWithLoggingWithAccessLog(t *testing.T) {
	var buf bytes.Buffer
	s := httptest.NewServer(
		WithLogging(
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(400)
			},
			buflogr.NewWithBuffer(&buf), Config{AccessLog: true}),
	)
	defer s.Close()

	_, err := http.Get(s.URL)
	assert.Nil(t, err)

	t.Logf("Log: %v", buf.String())

	assert.Contains(t, buf.String(), "Server access")
	assert.Contains(t, buf.String(), "statusCode 400")
}

func TestWithLoggingWithContextLog(t *testing.T) {
	var buf bytes.Buffer
	s := httptest.NewServer(
		WithLogging(
			func(w http.ResponseWriter, r *http.Request) {
				logr.FromContextOrDiscard(r.Context()).Info("testString")
			},
			buflogr.NewWithBuffer(&buf), Config{ContextLog: true}),
	)
	defer s.Close()

	_, err := http.Get(s.URL)
	assert.Nil(t, err)

	t.Logf("Log: %v", buf.String())

	assert.Contains(t, buf.String(), "testString")
}
