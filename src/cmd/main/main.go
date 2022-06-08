package main

import (
	"github.com/sean-cunniffe/go-bookstore/src/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Println("Starting server")
	log.Fatal(http.ListenAndServe(":8010", r))
}
