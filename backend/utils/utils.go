package utils

import (
	"encoding/json"
	"net/http"

	"social-network/backend/models"
)

func WriteJsonErrors(w http.ResponseWriter, errJson models.ErrorJson) {
	w.WriteHeader(errJson.Status)
	json.NewEncoder(w).Encode(errJson)
}

func WriteDataBack(w http.ResponseWriter, data any) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&data)
}

func IsValidImageType(mimeType string) bool {
	return mimeType == "image/jpeg" || mimeType == "image/png" || mimeType == "gif"
}
