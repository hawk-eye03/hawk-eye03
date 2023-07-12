package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shreyash-tripathi/sellerAppAssignment/lib/api"
)

func main() {
	router := mux.NewRouter()
	api.RegisterBookStoreRoutes(router)
	log.Println("Started HTTP server on localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
