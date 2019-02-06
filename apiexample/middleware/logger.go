package middleware

import (
	"log"
	"net/http"
)

type logResponseWriter struct {
	rw     http.ResponseWriter
	status int
}

func (w *logResponseWriter) WriteHeader(status int) {
	w.status = status
	w.rw.WriteHeader(status)
}

func (w *logResponseWriter) Header() http.Header {
	return w.rw.Header()
}

func (w *logResponseWriter) Write(data []byte) (int, error) {
	return w.rw.Write(data)
}

// Logger logs basic http request/response data
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %v", r.URL.Path)
		lw := &logResponseWriter{rw: w}
		next(lw, r)
		log.Printf("Request complete: %v (%v)", r.URL.Path, lw.status)
	}
}
