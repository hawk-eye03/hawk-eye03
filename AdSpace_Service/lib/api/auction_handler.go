package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/shreyash-tripathi/ad-service/lib/constants"
	"github.com/shreyash-tripathi/ad-service/lib/payload"
	"github.com/shreyash-tripathi/ad-service/lib/service"
	"github.com/shreyash-tripathi/ad-service/lib/utils"
)

type AuctionHandler struct {
	AuctionServ service.AuctionService
}

func (ah *AuctionHandler) CreateAuction(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside Create Auction Handler")

	var auction payload.AuctionReq
	err := json.NewDecoder(r.Body).Decode(&auction)
	if err != nil {
		log.Println(constants.ErrorInvalidAuctionRequestBodyMsg)

		errorResponse := payload.ErrorResponse{
			ErrorCode:        constants.ErrorInvalidAuctionRequestBody,
			ErrorDescription: constants.ErrorInvalidAuctionRequestBodyMsg,
		}
		utils.ErrorResponseWriter(w, errorResponse, http.StatusBadRequest)
		return
	}

	errResp := validateAuctionPayload(auction)
	if errResp != nil {
		utils.ErrorResponseWriter(w, *errResp, http.StatusBadRequest)
		return
	}
	// change error message
	err = ah.AuctionServ.CreateAuction(auction)
	if err != nil {
		log.Println("Error received from DB while creating auction")

		errorResponse := payload.ErrorResponse{
			ErrorCode:        constants.ErrorFromDBInCreatingNewAuction,
			ErrorDescription: fmt.Sprintf(constants.ErrorFromDBInCreatingNewAuctionMsg, err),
		}
		utils.ErrorResponseWriter(w, errorResponse, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	log.Println("Create Auction call Successful")
}

func validateAuctionPayload(auction payload.AuctionReq) *payload.ErrorResponse {
	if auction.AdSpaceID == nil {
		log.Println(constants.ErrorAdSpaceIDEmptyMsg)

		errorResponse := payload.ErrorResponse{
			ErrorCode:        constants.ErrorAdSpaceIDEmpty,
			ErrorDescription: constants.ErrorAdSpaceIDEmptyMsg,
		}

		return &errorResponse
	}
	if auction.EndTime == nil {
		log.Println(constants.ErrorAuctionEndTimeEmptyMsg)

		errorResponse := payload.ErrorResponse{
			ErrorCode:        constants.ErrorAuctionEndTimeEmpty,
			ErrorDescription: constants.ErrorAuctionEndTimeEmptyMsg,
		}

		return &errorResponse
	}
	return nil
}

func (ah *AuctionHandler) GetAllAuctions(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside Get All Auctions Handler")

	resp, err := ah.AuctionServ.GetAllAuctions()
	if err != nil {
		log.Println("Error received from DB while getting all auctions")

		errorResponse := payload.ErrorResponse{
			ErrorCode:        constants.ErrorFromDBInFetchingAllAuctions,
			ErrorDescription: fmt.Sprintf(constants.ErrorFromDBInFetchingAllAuctionsMsg, err),
		}
		utils.ErrorResponseWriter(w, errorResponse, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	log.Println("Get All Auction call Successful")
}
