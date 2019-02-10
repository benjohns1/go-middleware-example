package middleware

import (
	"log"
	"net/http"
)

type logResponseWriter struct {
	http.ResponseWriter
	status int
}

func (w *logResponseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

// Logger logs basic http request/response data
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %v", r.URL.Path)
		lw := &logResponseWriter{ResponseWriter: w}
		next(lw, r)
		log.Printf("Request complete: %v (%v)", r.URL.Path, lw.status)
	}
}
