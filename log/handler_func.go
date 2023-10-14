package log

import (
	"net/http"

	"github.com/go-logr/logr"
)

func WithLogging(h http.HandlerFunc, logger logr.Logger, cfg Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if cfg.AccessLog {
			w = NewResponseWriter(w)
		}
		if cfg.ContextLog {
			r = r.WithContext(logr.NewContext(r.Context(), logger))
		}

		h(w, r)

		if cfg.AccessLog {
			if rw, ok := w.(*responseWriter); ok {
				logger = logger.WithValues("statusCode", rw.statusCode)
			}
			logger.Info("Server access",
				"method", r.Method,
				"proto", r.Proto,
				"header", r.Header.Clone(),
				"contentLength", r.ContentLength,
				"host", r.Host,
				"remoteAddr", r.RemoteAddr,
				"requestURI", r.RequestURI,
			)

		}
	}
}
