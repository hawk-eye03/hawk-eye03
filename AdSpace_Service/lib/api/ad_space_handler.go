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

type Ad_Space_Handler struct {
	AdSpaceServ service.AdSpaceService
}

func (ah *Ad_Space_Handler) CreateAdSpace(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside Ad Space Handler")

	var adSpace payload.AdSpaceReq
	err := json.NewDecoder(r.Body).Decode(&adSpace)
	if err != nil {
		log.Println(constants.ErrorInvalidAdSpaceRequestBody)

		errorResponse := payload.ErrorResponse{
			ErrorCode:        constants.ErrorInvalidAdSpaceRequestBody,
			ErrorDescription: constants.ErrorInvalidAdSpaceRequestBodyMsg,
		}
		utils.ErrorResponseWriter(w, errorResponse, http.StatusBadRequest)
		return
	}

	errResp := validateAdSpacePayload(adSpace)
	if errResp != nil {
		utils.ErrorResponseWriter(w, *errResp, http.StatusBadRequest)
		return
	}

	err = ah.AdSpaceServ.CreateAdSpace(adSpace)
	if err != nil {
		log.Println("Error received from DB while creating ad space")

		errorResponse := payload.ErrorResponse{
			ErrorCode:        constants.ErrorFromDBInCreatingNewAdSpace,
			ErrorDescription: fmt.Sprintf(constants.ErrorFromDBInCreatingNewAdSpaceMsg, err),
		}
		utils.ErrorResponseWriter(w, errorResponse, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func validateAdSpacePayload(adSpace payload.AdSpaceReq) *payload.ErrorResponse {
	if adSpace.Name == "" {
		log.Println(constants.ErrorAdSpaceNameEmptyMsg)

		errorResponse := payload.ErrorResponse{
			ErrorCode:        constants.ErrorAdSpaceNameEmpty,
			ErrorDescription: constants.ErrorAdSpaceNameEmptyMsg,
		}

		return &errorResponse
	}
	if adSpace.BasePrice == "" {
		log.Println(constants.ErrorAdSpaceBasePriceEmpty)

		errorResponse := payload.ErrorResponse{
			ErrorCode:        constants.ErrorAdSpaceBasePriceEmpty,
			ErrorDescription: constants.ErrorAdSpaceBasePriceEmptyMsg,
		}

		return &errorResponse
	}
	return nil
}
