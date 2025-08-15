package utils

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	_ "image/gif" //  these imports are used for the init function
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"strings"
	"time"

	"social-network/backend/models"
)

// set the default image based on the specific business logic
// so the default image MUST change based on if the group or the profile for example!!!

func HanldeUploadImage(r *http.Request, fileName, subDirectoryName string) (string, *models.ErrorJson) {
	
	file, _, err := r.FormFile(fileName)
	if err != nil {
		if err == http.ErrMissingFile {
			
			return "", nil
		}
		return "", &models.ErrorJson{Status: 400, Error: fmt.Sprintf("%v", err)}
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	written, err := io.Copy(buf, file)
	// so the 500 is the more convenient way to handle it
	if err != nil {
		return "", &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	if written == 0 {
		return "", &models.ErrorJson{Status: 400, Error: "No content is being detected!!"}
	}

	mimeType := http.DetectContentType(buf.Bytes())
	if !IsValidImageType(mimeType) {
		return "", &models.ErrorJson{Status: 400, Error: "Error!! Only PNG, JPEG and GIF images are allowed"}
	}
	// the string returned here is the actual format (we can do another check to see if the mimeType is
	// on harmony with the format but no need !!
	_, _, err = image.Decode(bytes.NewReader(buf.Bytes()))
	if err != nil {
		return "", &models.ErrorJson{Status: 400, Error: "Error!! Invalid Image Content"}
	}

	if len(buf.Bytes()) > (3 * 1024 * 1024) {
		return "", &models.ErrorJson{Status: 400, Error: "File too large"}
	}

	path, errJson := CreateDirectoryForUploads(subDirectoryName, mimeType, buf.Bytes())
	if errJson != nil {
		return "", &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	return path, nil
}

func ValidateTitle(title string) error {
	if title == "" {
		return errors.New("title can not be empty")
	}
	if len(title) < 3 {
		return errors.New("title is too short! 3 characters min")
	}
	if len(title) > 100 {
		return errors.New("title is too large! 100 characters max")
	}
	return nil
}

func ValidateDesc(desc string) error {
	if desc == "" {
		return errors.New("description can not be empty")
	}
	if len(desc) < 10 {
		return errors.New("description is too short! 10 characters min")
	}
	if len(desc) > 1000 {
		return errors.New("description is too large! 1000 characters max")
	}
	return nil
}

// mn b3d had  l validation ana t2kdt belli blassti mashi f zone!!

func ValidateDateEvent(date string) error {
	s := strings.Trim(date, `"`)
	timeParsed, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return errors.New("date format is incorrect: YYYY-MM-DD HH:MM:SS")
	}

	if timeParsed.IsZero() {
		return errors.New("the date is not set up")
	}
	if timeParsed.Before(time.Now()) {
		return fmt.Errorf("please set a date that comes after %v", models.NewDate(time.Now()).Format("2006-01-02"))
	}

	return nil
}

func IsValidFilter(filter string) bool {
	return filter == "owned" || filter == "available" || filter == "joined"
}
