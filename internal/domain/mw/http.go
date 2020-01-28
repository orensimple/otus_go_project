package mw

import (
	"context"
	"net/http"
	"time"

	"github.com/orensimple/otus_go_project/internal/logger"
)

func HTTPLogger(h http.HandlerFunc) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		uri := req.URL.RequestURI()
		logger.ContextLogger.Infof("Debug info do", "uri", uri)
		h(resp, req)
	}
}

func WithTimeout(h http.HandlerFunc, timeout time.Duration) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		ctx, cancel := context.WithTimeout(req.Context(), timeout)
		defer cancel()
		req = req.WithContext(ctx)
		h(resp, req)
	}
}
