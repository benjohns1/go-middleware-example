package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

type middlewareFunc func(next http.HandlerFunc) http.HandlerFunc

func chain(chain ...middlewareFunc) http.Handler {
	return http.HandlerFunc(recurseChain(chain))
}

func recurseChain(chain []middlewareFunc) http.HandlerFunc {
	if len(chain) <= 0 {
		return func(_ http.ResponseWriter, _ *http.Request) {}
	}
	return chain[0](recurseChain(chain[1:]))
}

func main() {
	http.Handle("/api/v1", chain(logger, businessLogic(func() bool { return rand.Intn(2) == 1 })))
	log.Print("Listening (customstate) ...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

type logResponseWriter struct {
	rw     http.ResponseWriter
	status int
	extra  string
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

func logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %v", r.URL.Path)
		lw := &logResponseWriter{rw: w}
		next(lw, r)
		log.Printf("Request complete: %v (%v status): %v", r.URL.Path, lw.status, lw.extra)
	}
}

func businessLogic(decider func() bool) middlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if decider() {
				w.WriteHeader(200)
				v := "Ran business logic"
				log.Print(v)
				fmt.Fprint(w, v)

				// Add custom state to response writer
				if lw, ok := w.(*logResponseWriter); ok {
					lw.extra = "Extra response data applied by business logic"
					next(lw, r)
				} else {
					log.Print("No logResponseWriter sent to business logic, no extra data applied")
					next(w, r)
				}

			} else {
				w.WriteHeader(403)
			}
		}
	}
}
