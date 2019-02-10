package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/benjohns1/go-middleware-example/apiexample/businessdomain"
	"github.com/benjohns1/go-middleware-example/apiexample/middleware"
	"github.com/go-chi/render"
)

func main() {
	prefix := "/api/v1/"
	http.Handle(prefix+"random/", middleware.Chain(middleware.Logger, middleware.Auth(middleware.RandAuthStrategy), requestHandler))
	http.Handle(prefix+"public/", middleware.Chain(middleware.Logger, middleware.Auth(middleware.PublicAuthStrategy), requestHandler))
	log.Print("Listening (apiexample) ...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

type response struct {
	Data string `json:"data" xml:"data"`
}

func requestHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var v string
		responseData, err := businessdomain.BusinessLogic()
		if err == nil {
			w.WriteHeader(http.StatusOK)
			v = fmt.Sprintf("Ran business logic")
		} else {
			w.WriteHeader(http.StatusBadRequest)
			v = fmt.Sprintf("Error running business logic: " + err.Error())
		}
		log.Print(v)

		if strings.HasSuffix(r.URL.Path, ".xml") {
			render.XML(w, r, response(*responseData))
		} else {
			render.JSON(w, r, response(*responseData))
		}

		next(w, r)
	}
}
