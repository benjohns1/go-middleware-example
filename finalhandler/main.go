package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

type middlewareFunc func(next http.HandlerFunc) http.HandlerFunc

func chain(handler http.HandlerFunc, chain ...middlewareFunc) http.Handler {
	return http.HandlerFunc(recurseChain(handler, chain))
}

func recurseChain(handler http.HandlerFunc, chain []middlewareFunc) http.HandlerFunc {
	if len(chain) <= 0 {
		return handler
	}
	return chain[0](recurseChain(handler, chain[1:]))
}

func main() {
	http.Handle("/api/v1", chain(businessLogic("12345"), logger, shortCircuitLogic(deciderFunc)))
	log.Print("Listening (finalhandler) ...")
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

			if deciderFunc() {
				log.Print("Continuing request")
				next(w, r)
			} else {
				log.Print("Halting request")
				fmt.Fprint(w, "Request halted")
			}
		}
	}
}

func businessLogic(p1 string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v := fmt.Sprintf("Ran business logic with server-start-time param: %v", p1)
		log.Print(v)
		fmt.Fprint(w, v)
	}
}
