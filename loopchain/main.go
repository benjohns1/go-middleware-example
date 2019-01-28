package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func chain(chain ...http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, h := range chain {
			h(w, r)
		}
	})
}

func main() {
	http.Handle("/api/v1", chain(logger, businessLogic("99999", runTimeFunc)))
	log.Print("Listening (loopchain) ...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func logger(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request %v", r.URL.Path)
}

func runTimeFunc() int {
	return rand.Intn(1000)
}

func businessLogic(p1 string, runTimeFunc func() int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v := fmt.Sprintf("Logic with server-start-time param: %v, run-time param: %v", p1, runTimeFunc())
		log.Print(v)
		fmt.Fprint(w, v)
	}
}
