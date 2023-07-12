package api

import (
	"fmt"
	"net/http"
)

func CreateAuction(w http.ResponseWriter, r *http.Request) {
	// CreateBook := &models.Book{}
	// utils.ParseBody(r, CreateBook)
	// b := CreateBook.CreateBook()
	// res, _ := json.Marshal(b)
	fmt.Println("Created Auction")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte{1})
}

func GetAuction(w http.ResponseWriter, r *http.Request) {
	// CreateBook := &models.Book{}
	// utils.ParseBody(r, CreateBook)
	// b := CreateBook.CreateBook()
	// res, _ := json.Marshal(b)
	fmt.Println("Get Action")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte{1})
}

func DeleteAuction(w http.ResponseWriter, r *http.Request) {
	// CreateBook := &models.Book{}
	// utils.ParseBody(r, CreateBook)
	// b := CreateBook.CreateBook()
	// res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte{1})
}

func UpdateAuction(w http.ResponseWriter, r *http.Request) {
	// CreateBook := &models.Book{}
	// utils.ParseBody(r, CreateBook)
	// b := CreateBook.CreateBook()
	// res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte{1})
}
