package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shreyash-tripathi/ad-service/lib/api"
	"github.com/shreyash-tripathi/ad-service/lib/constants"
	"github.com/shreyash-tripathi/ad-service/lib/db"
)

func main() {
	router := mux.NewRouter()
	api.RegisterAuctionRoutes(router)
	api.RegisterAdSpaceRoutes(router)

	port := constants.Port
	log.Println("Started HTTP server on localhost:", port)

	db.Connect()
	log.Fatal(http.ListenAndServe("localhost:"+port, router))
}
