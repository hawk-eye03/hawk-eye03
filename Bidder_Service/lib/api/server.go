package api

import (
	"github.com/gorilla/mux"
	"github.com/shreyash-tripathi/bidder-service/lib/service"
)

var RegisterBidderRoutes = func(router *mux.Router) {
	bh := Bidder_Handler{
		BidderServ: service.BidderService{},
	}
	router.HandleFunc("/CreateBidder", bh.CreateBidder).Methods("POST")

	// router.HandleFunc("/DeleteAuction/{auctionID}", DeleteAuction).Methods("DELETE")
	// router.HandleFunc("/UpdateAuction/{auctionID}", UpdateAuction).Methods("PUT")

}
