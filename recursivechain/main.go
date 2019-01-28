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
	http.Handle("/api/v1", chain(logger, shortCircuitLogic(deciderFunc), businessLogic("12345")))
	log.Print("Listening (recursivechain) ...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %v", r.URL.Path)
		next(w, r)
		log.Printf("Request complete: %v", r.URL.Path)
	}
}

func deciderFunc() bool {
	return rand.Intn(2) == 1
}

func shortCircuitLogic(deciderFunc func() bool) middlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			d := deciderFunc()
			var decision string
			if d {
				decision = "continue with"
			} else {
				decision = "halt"
			}

			v := fmt.Sprintf("Randomly decided to %v the request", decision)
			log.Print(v)
			fmt.Fprint(w, v+"\n")

			if !d {
				return
			}

			next(w, r)
		}
	}
}

func businessLogic(p1 string) middlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			v := fmt.Sprintf("Ran business logic with server-start-time param: %v", p1)
			log.Print(v)
			fmt.Fprint(w, v)
			next(w, r)
		}
	}
}
