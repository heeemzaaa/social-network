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
	jsonError := s.validateUserData(user)
	if jsonError != nil {
		return jsonError
	}

	user.Id = utils.NewUUID()
	hash, err := HashPassword(user.Password)

	if err != nil {
		return models.NewErrorJson(500, err.Error(), nil)
	} else {
		user.Password = hash
	}

	errJson := s.repo.CreateUser(user)
	if errJson != nil {
		// needs another checking 
		os.Remove(user.ImagePath)
		return errJson
	}

	return nil
}




func (s *AuthService) validateUserData(user *models.User) *models.ErrorJson {
	userErrorJson := models.User{}
	if err := isValidName(user.FirstName); err != nil {
		userErrorJson.FirstName = err.Error()
	}
	if strings.TrimSpace(user.FirstName) == "" {
		userErrorJson.FirstName = "First name is required"
	}

	if err := isValidName(user.LastName); err != nil {
		userErrorJson.LastName = err.Error()
	}

	if strings.TrimSpace(user.LastName) == "" {
		userErrorJson.LastName = "Last name is required"
	}


	if err := isValidName(user.LastName); err != nil {
		userErrorJson.LastName = err.Error()
	}

	if err := isValidBirthDate(user.BirthDate); err != nil {
		userErrorJson.LastName = err.Error()
	}

	if err := isValidBirthDate(user.BirthDate); err != nil {
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
		return &models.ErrorJson{Status: 400, Message: userErrorJson}
	}

	return nil
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

func (s *AuthService) EmailVerification(email string) error {
	if len(email) > 255 {
		return fmt.Errorf("email too long")
	}
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,20}$`
	if match, _ := regexp.MatchString(emailRegex, email); !match {
		return fmt.Errorf("invalid email format")
	}
	_, has_email, _ := s.repo.GetItem("users", "email", email)
	if has_email {
		return fmt.Errorf("email already in use")
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
