package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shreyash-tripathi/bidder-service/lib/api"
	"github.com/shreyash-tripathi/bidder-service/lib/constants"
	"github.com/shreyash-tripathi/bidder-service/lib/db"
)

func main() {
	router := mux.NewRouter()
	api.RegisterBidderRoutes(router)

	port := constants.Port
	log.Println("Started HTTP server on localhost:", port)

	db.Connect()
	log.Fatal(http.ListenAndServe("localhost:"+port, router))
}
