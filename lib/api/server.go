package api

import (
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/CreateAuction", CreateAuction).Methods("POST")
	router.HandleFunc("/GetAuction", GetAuction).Methods("GET")
	router.HandleFunc("/DeleteAuction/{auctionID}", DeleteAuction).Methods("DELETE")
	router.HandleFunc("/UpdateAuction/{auctionID}", UpdateAuction).Methods("PUT")
	// router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}

// func CreateAuction(w http.ResponseWriter, r *http.Request) {
// 	// CreateBook := &models.Book{}
// 	// utils.ParseBody(r, CreateBook)
// 	// b := CreateBook.CreateBook()
// 	// res, _ := json.Marshal(b)
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte{1})
// }
