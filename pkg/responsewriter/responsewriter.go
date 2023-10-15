package responsewriter

import (
	"errors"
	"net/http"
)

var ErrTypeError = errors.New("not the proper custom responseWriter")

type ResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func Tranform(w http.ResponseWriter) *ResponseWriter {
	if cw, ok := w.(*ResponseWriter); ok {
		return cw
	}
	return &ResponseWriter{w, http.StatusOK}
}

func (w *ResponseWriter) WriteHeader(code int) {
	w.StatusCode = code
	w.ResponseWriter.WriteHeader(code)
}
