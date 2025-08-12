package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"social-network/backend/models"

	"github.com/google/uuid"
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
	return mimeType == "image/jpeg" || mimeType == "image/png" || mimeType == "image/gif"
}

// function that creates the directory for the files for the groups and profiles and so on

func CreateDirectoryForUploads(subDirectoryName, mimeType string, data []byte) (string, *models.ErrorJson) {
	baseDir := "static/uploads/"
	err := os.MkdirAll(filepath.Join(baseDir, subDirectoryName), 0o755)
	if err != nil {
		return "", &models.ErrorJson{Status: 500, Error: "Failed to create the directory"}
	}
	ext := ".jpg"
	switch mimeType {
	case "image/png":
		ext = ".png"
	case "image/gif":
		ext = ".gif"
	}

	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	path := filepath.Join(baseDir, subDirectoryName, filename)
	err = os.WriteFile(path, data, 0o644)
	if err != nil {
		return "", &models.ErrorJson{Status: 500, Error: "Failed to write data into the file"}
	}

	// Return relative path for use in frontend / API
	relativePath := strings.TrimPrefix(path, "static/")
	relativePath = "/" + strings.ReplaceAll(relativePath, "\\", "/") // cross-platform

	return relativePath, nil
}

func GetUUIDFromPath(r *http.Request, key string) (uuid.UUID, error) {
	val := r.PathValue(key)
	return uuid.Parse(val)
}

func IsValidUUID(id string) error {
	return uuid.Validate(id)
}

func NewUUID() string {
	return uuid.New().String()
}

func RemoveImage(ImagePath string) error {
	if err := os.Remove(filepath.Join("static", ImagePath)); err != nil {
		return err
	}
	return nil
}
