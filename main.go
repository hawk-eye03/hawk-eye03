package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shreyash-tripathi/sellerAppAssignment/lib/api"
	"github.com/shreyash-tripathi/sellerAppAssignment/lib/db"
)

func main() {
	router := mux.NewRouter()
	api.RegisterBookStoreRoutes(router)
	log.Println("Started HTTP server on localhost:8080")

	db.Connect()
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
