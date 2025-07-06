package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

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

// validate the image Type
func IsValidImageType(mimeType string) bool {
	return mimeType == "image/jpeg" || mimeType == "image/png" || mimeType == "gif"
}

// function that creates the directory for the files for the groups and profiles and so on

func CreateDirectoryForUploads(subDirectoryName, mimeType string, data []byte) (string, *models.ErrorJson) {
	err := os.MkdirAll("static/uploads/"+subDirectoryName, 0o755)
	if err != nil {
		return "", &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	ext := ".jpg"
	switch mimeType {
	case "image/png":
		ext = ".png"
	case "image/gif":
		ext = ".gif"
	}

	filename := fmt.Sprintf("static/uploads/%s/%d%s", subDirectoryName, time.Now().UnixNano(), ext)
	err = os.WriteFile(filename, data, 0o644)
	if err != nil {
		return "", &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return filename, nil
}
