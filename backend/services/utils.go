package services

import (
	"mime/multipart"
	"os"
	"path/filepath"
)

// image uplaod helper functions

func UploadImage(image multipart.File, imageName string) (imgPath string, err error) {
	dst, err := CreateFile("profile", imageName)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy the uploaded file to the destination file
	if _, err := dst.ReadFrom(image); err != nil {
		return "", err
	}
	return dst.Name(), nil
}

func CreateFile(subFolder, imageName string) (*os.File, error) {
	// Create the full folder path: uploads/subFolder
	fullFolderPath := filepath.Join("uploads", subFolder)

	if err := os.MkdirAll(fullFolderPath, 0o755); err != nil {
		return nil, err
	}

	// Create the file in the specified directory
	filePath := filepath.Join(fullFolderPath, imageName)
	dst, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}

	return dst, nil
}
