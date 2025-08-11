package auth

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"social-network/backend/models"
	"social-network/backend/utils"
)

func (s *AuthService) Register(user *models.User) *models.ErrorJson {
	// data validation
	UserData, jsonError := s.validateUserData(user)
	if jsonError != nil {
		return jsonError
	}

	UserData.Id = utils.NewUUID()
	hash, err := HashPassword(UserData.Password)

	if err != nil {
		return models.NewErrorJson(500, err.Error(), nil)
	} else {
		UserData.Password = hash
	}

	errJson := s.repo.CreateUser(UserData)
	if errJson != nil {
		// needs another checking
		os.Remove(UserData.ImagePath)
		return errJson
	}

	return nil
}

func (s *AuthService) validateUserData(user *models.User) (*models.User, *models.ErrorJson) {
	trimmedFirstName := strings.TrimSpace(user.FirstName)
	trimmedLastName := strings.TrimSpace(user.LastName)
	trimmedUsername := strings.TrimSpace(user.Nickname)
	trimmedAboutMe := strings.TrimSpace(user.AboutMe)
	userErrorJson := models.User{}
	
	if err := isValidName(trimmedFirstName, "firstname"); err != nil {
		userErrorJson.FirstName = err.Error()
	}

	if err := isValidName(trimmedLastName, "lastname"); err != nil {
		userErrorJson.LastName = err.Error()
	}

	if err := ValidateDateRegister(user.BirthDate); err != nil {
		userErrorJson.BirthDate = err.Error()
	}

	if err := s.isValidEmail(user.Email); err != nil {
		userErrorJson.Email = err.Error()
	}

	if err := isValidPwd(user.Password); err != nil {
		userErrorJson.Password = err.Error()
	}

	// optianal user data
	if err := s.isValidNickname(user.Nickname); err != nil {
		userErrorJson.Nickname = err.Error()
	}

	if err := isValidAboutme(user.AboutMe); err != nil {
		userErrorJson.AboutMe = err.Error()
	}

	if userErrorJson != (models.User{}) {
		return nil, &models.ErrorJson{Status: 400, Message: userErrorJson}
	}
	// WE need to trim the data before insert it in the database

	UserData := user
	UserData.FirstName, UserData.LastName, UserData.Nickname, UserData.AboutMe = trimmedFirstName, trimmedLastName, trimmedUsername, trimmedAboutMe

	return UserData, nil
}

func (s *AuthService) IsValidNickname(nickname string) error {
	if len(nickname) < 3 {
		return fmt.Errorf("username is too short")
	}
	if len(nickname) > 30 {
		return fmt.Errorf("username is too long")
	}
	usernameRegex := `^[a-zA-Z0-9_.-]+$`
	if match, _ := regexp.MatchString(usernameRegex, nickname); !match {
		return fmt.Errorf("username can only contain letters, digits, underscores, dots, and hyphens")
	}
	_, has_nickname, _ := s.repo.GetItem("users", "nickname", nickname)
	if has_nickname {
		return fmt.Errorf("username already exists")
	}
	return nil
}

// lookarounds are not possible
func isValidPwd(password string) error {
	if strings.TrimSpace(password) == "" {
		return fmt.Errorf("password is required")
	}
	if len(password) < 8 {
		return fmt.Errorf("password is too short")
	}
	if len(password) > 50 {
		return fmt.Errorf("password is too long")
	}
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[\W_]`).MatchString(password)
	if !hasLower || !hasUpper || !hasDigit || !hasSpecial {
		return fmt.Errorf("password must contain at least one lowercase letter, one uppercase letter, one digit, and one special character")
	}
	return nil
}
