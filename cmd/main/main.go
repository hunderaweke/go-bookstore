package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hunderaweke/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	r.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
