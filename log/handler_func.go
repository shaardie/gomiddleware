package log

import (
	"net/http"

	"github.com/go-logr/logr"
	"github.com/shaardie/gomiddleware/pkg/responsewriter"
)

func WithLogging(h http.HandlerFunc, logger logr.Logger, cfg Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cw = &responsewriter.ResponseWriter{}
		if cfg.AccessLog {
			cw = responsewriter.Tranform(w)
		}
		if cfg.ContextLog {
			r = r.WithContext(logr.NewContext(r.Context(), logger))
		}

		h(cw, r)

		if cfg.AccessLog {
			logger = logger.WithValues("statusCode", cw.StatusCode)
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
