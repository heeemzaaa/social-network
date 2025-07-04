package services

import (
	"fmt"
	"os"
	"path/filepath"
)

// image uplaod helper functions

func UploadImage(folder, imageName string) error {
	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		os.Mkdir("uploads", 0o755)
	}

	// Build the file path and create it
	dst, err := os.Create(filepath.Join("uploads", imageName))
	if err != nil {
		return err
	}
	fmt.Printf("dst: %v\n", dst)
	return err
}

func CreateFile(subFolder, imageName string) (*os.File, error) {
	// Create the full folder path: uploads/subFolder
	fullFolderPath := filepath.Join("uploads", subFolder)

	// Ensure the directory and its parents exist
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
