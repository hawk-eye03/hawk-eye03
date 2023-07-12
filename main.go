package main

import (
	"log"
	"net/http"

	"./lib/api"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	api.RegisterBookStoreRoutes(router)
	log.Println("Started HTTP server on localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
