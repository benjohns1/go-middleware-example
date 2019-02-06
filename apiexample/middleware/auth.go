package middleware

import (
	"log"
	"math/rand"
	"net/http"
)

// RandAuthStrategy randomly authorizes a request (this is highly recommended for use in production to maximize user frustration and minimize revenue)
func RandAuthStrategy() bool {
	return rand.Intn(2) == 1
}

// Auth generates an http request auth method with the given strategy
func Auth(authStrategy func() bool) Factory {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			d := authStrategy()
			if !d {
				w.WriteHeader(401)
				log.Print("Unauthorized")
				return
			}

			log.Print("Authorized")
			next(w, r)
		}
	}
}
