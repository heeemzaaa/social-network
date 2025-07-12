// authentication serevice.
package auth

import (
	"fmt"
	"mime/multipart"
	"os"
	"strings"
	"time"

	"social-network/backend/models"
	"social-network/backend/repositories/auth"
	"social-network/backend/services"
	"social-network/backend/utils"

	"github.com/google/uuid"
)








func (s *AuthService) Register(user *models.User, file multipart.File) *models.ErrorJson {
	// data validation
	jsonError := s.validateUserData(user, file)
	if jsonError != nil {
		return jsonError
	}

	user.Id = utils.NewUUID()
	// Image handling:
	if user.ImagePath != "" {
		ImgName := user.Id + "." + strings.Split(user.ImagePath, ".")[1]
		imgPath, err := services.UploadImage(file, ImgName)
		if err != nil {
			return models.NewErrorJson(500, "Error saving profile image", nil)
		}
		user.ImagePath = imgPath
	}

	hash, err := HashPassword(user.Password)

	if err != nil {
		return models.NewErrorJson(500, err.Error(), nil)
	} else {
		user.Password = hash
	}

	errJson := s.repo.CreateUser(user)
	if errJson != nil {
		os.Remove(user.ImagePath)
		return errJson
	}

	return nil
}

func (s *AuthService) validateUserData(user *models.User, file multipart.File) *models.ErrorJson {
	userErrorJson := models.User{}

	if user.FirstName == "" {
		userErrorJson.FirstName = "first name is required"
	} else {
		userErrorJson.FirstName = isValidName(user.FirstName)
	}

	if user.LastName == "" {
		userErrorJson.LastName = "last name is required"
	} else {
		userErrorJson.LastName = isValidName(user.LastName)
	}

	userErrorJson.BirthDate = isValidBirthDate(user.BirthDate)
	userErrorJson.Email = s.isValidEmail(user.Email)
	userErrorJson.Password = isValidPwd(user.Password)

	// optianal user data
	userErrorJson.Nickname = s.isValidNickname(user.Nickname)
	userErrorJson.AboutMe = isValidAboutme(user.AboutMe)
	if file != nil {
		userErrorJson.ImagePath = isValidImg(file)
	}

	if userErrorJson != (models.User{}) {
		return &models.ErrorJson{Status: 400, Message: userErrorJson}
	}

	return nil
}











