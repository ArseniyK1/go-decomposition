package middleware

import (
	"decomposition/pkg/logger"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	logger.SetLogger(&logger.ConsoleLogger{})
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("Incoming request", "method", r.Method, "path", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
