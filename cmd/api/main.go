package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type apiHandler struct{}

func (apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// api logic goes here
	_, err := fmt.Fprintf(w, "api endpoint reached")
	if err != nil {
		log.Println("error writing response:", err)
	}
}

func main() {
	r := chi.NewRouter()

	r.Handle("/api/*", apiHandler{})
	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		_, err := fmt.Fprintf(w, "ITS ALIVEEE LOOOLL!")
		if err != nil {
			log.Println("error writing response:", err)
		}
	})

	server := &http.Server{
		Addr:         ":5500",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
