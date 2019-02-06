package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/benjohns1/go-middleware-example/apiexample/businessdomain"
	"github.com/benjohns1/go-middleware-example/apiexample/middleware"
	"github.com/go-chi/render"
)

func main() {
	http.Handle("/api/v1", middleware.Chain(middleware.Logger, middleware.Auth(middleware.RandAuthStrategy), requestHandler))
	log.Print("Listening (apiexample) ...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

type jsonData struct {
	Data string `json:"data"`
}

func requestHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var v string
		responseData, err := businessdomain.BusinessLogic()
		if err == nil {
			w.WriteHeader(200)
			v = fmt.Sprintf("Ran business logic")
		} else {
			w.WriteHeader(400)
			v = fmt.Sprintf("Failed to run business logic")
		}
		log.Print(v)
		render.JSON(w, r, jsonData(*responseData))
		next(w, r)
	}
}
