package main

import (
	"log"
	"net/http"

	"spanish-conjugation-api/api/router"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	router.Initialize(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
