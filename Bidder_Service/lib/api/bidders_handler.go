package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/shreyash-tripathi/bidder-service/lib/constants"
	"github.com/shreyash-tripathi/bidder-service/lib/payload"
	"github.com/shreyash-tripathi/bidder-service/lib/service"
	"github.com/shreyash-tripathi/bidder-service/lib/utils"
)

type Bidder_Handler struct {
	BidderServ service.BidderService
}

func (bh *Bidder_Handler) CreateBidder(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside Bidder Handler")

	var bidder payload.BidderReq
	err := json.NewDecoder(r.Body).Decode(&bidder)
	if err != nil {
		log.Println(constants.ErrorInvalidBidderRequestBody)

		errorResponse := payload.ErrorResponse{
			ErrorCode:        constants.ErrorInvalidBidderRequestBody,
			ErrorDescription: constants.ErrorInvalidBidderRequestBodyMsg,
		}
		utils.ErrorResponseWriter(w, errorResponse, http.StatusBadRequest)
		return
	}

	errResp := validateBidderPayload(bidder)
	if errResp != nil {
		utils.ErrorResponseWriter(w, *errResp, http.StatusBadRequest)
		return
	}

	err = bh.BidderServ.CreateBidder(bidder)
	if err != nil {
		log.Println("Error received from DB while creating bidder")

		errorResponse := payload.ErrorResponse{
			ErrorCode:        constants.ErrorFromDBInCreatingNewBidder,
			ErrorDescription: fmt.Sprintf(constants.ErrorFromDBInCreatingNewBidderMsg, err),
		}
		utils.ErrorResponseWriter(w, errorResponse, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func validateBidderPayload(bidder payload.BidderReq) *payload.ErrorResponse {
	if bidder.Name == "" {
		log.Println(constants.ErrorBidderNameEmptyMsg)

		errorResponse := payload.ErrorResponse{
			ErrorCode:        constants.ErrorBidderNameEmpty,
			ErrorDescription: constants.ErrorBidderNameEmptyMsg,
		}

		return &errorResponse
	}
	return nil
}
