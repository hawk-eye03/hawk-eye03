package api

import (
	"github.com/gorilla/mux"
	"github.com/shreyash-tripathi/ad-service/lib/service"
)

var RegisterAuctionRoutes = func(router *mux.Router) {
	ah := AuctionHandler{
		AuctionServ: service.AuctionService{},
	}
	router.HandleFunc("/CreateAuction", ah.CreateAuction).Methods("POST")
	router.HandleFunc("/GetAllAuctions", ah.GetAllAuctions).Methods("GET")
	// router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}

var RegisterAdSpaceRoutes = func(router *mux.Router) {
	as := Ad_Space_Handler{
		AdSpaceServ: service.AdSpaceService{},
	}
	router.HandleFunc("/CreateAdSpace", as.CreateAdSpace).Methods("POST")
	// router.HandleFunc("/GetAuction", GetAuction).Methods("GET")
	// router.HandleFunc("/DeleteAuction/{auctionID}", DeleteAuction).Methods("DELETE")
	// router.HandleFunc("/UpdateAuction/{auctionID}", UpdateAuction).Methods("PUT")
	// router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
