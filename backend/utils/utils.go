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

func IsValidFilter(filter string) bool {
	return filter == "owned" || filter == "availabe" || filter == "joined"
}
