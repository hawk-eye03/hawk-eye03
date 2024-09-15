package utils

import (
	"encoding/json"
	"net/http"

	"github.com/shreyash-tripathi/ad-service/lib/payload"
)

func ErrorResponseWriter(w http.ResponseWriter, errorResponse payload.ErrorResponse, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errorResponse)
}
