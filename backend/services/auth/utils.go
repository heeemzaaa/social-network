package auth

import (
	"regexp"
	"slices"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func isValidName(name string) string {
	if len(name) > 50 {
		return "name must be 50 characters or less"
	}
	if !regexp.MustCompile(`^[a-zA-Z\s-]+$`).MatchString(name) {
		return "name can only contain letters, spaces, or hyphens"
	}
	return ""
}

func isValidBirthDate(date string) string {
	if date == "" {
		return "BirthDate is required"
	}

	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "birth date must be in YYYY-MM-DD format and valid"
	}

	return ""
}

func (s *AuthService) isValidEmail(email string) string {
	if email == "" {
		return "email is requierd"
	}
	if len(email) > 100 {
		return "email must be 100 characters or less"
	}
	// Basic email regex
	if !regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(email) {
		return "invalid email format"
	}

	_, has_email, _ := s.repo.GetItem("users", "email", email)
	if has_email {
		return "email already in use"
	}

	return ""
}

func isValidPwd(pwd string) string {
	if pwd == "" {
		return "password is required"
	}
	if len(pwd) < 8 {
		return "password must be at least 8 characters"
	}
	if len(pwd) > 128 {
		return "password must be 128 characters or less"
	}
	// Require at least one uppercase, one lowercase, one number, and one special character
	if !regexp.MustCompile(`[A-Z]`).MatchString(pwd) ||
		!regexp.MustCompile(`[a-z]`).MatchString(pwd) ||
		!regexp.MustCompile(`[0-9]`).MatchString(pwd) ||
		!regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(pwd) {
		return "password must contain at least one uppercase, one lowercase, one number, and one special character"
	}

	return ""
}

func (s *AuthService) isValidNickname(nickname string) string {
	if nickname == "" {
		return ""
	}

	if len(nickname) < 3 {
		return "nickname must be 3 characters or higher"
	}

	if len(nickname) > 30 {
		return "nickname must be 30 characters or less"
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9_-]+$`).MatchString(nickname) {
		return "nickname can only contain letters, numbers, underscores, or hyphens"
	}

	_, has_nickname, _ := s.repo.GetItem("users", "nickname", nickname)
	if has_nickname {
		return "username already exists"
	}

	return ""
}

func isValidAboutme(aboutme string) string {
	if aboutme == "" {
		return ""
	}

	if len(aboutme) > 500 {
		return "about me must be 500 characters or less"
	}

	return ""
}

func isValidImg(imgName string, size int64) string {
	if imgName == "" {
		return ""
	}

	validExtensions := []string{"png", "jpeg", "jpg", "svg"}
	imageExtensions := strings.Split(imgName, ".")[1]
	if !slices.Contains(validExtensions, imageExtensions) {
		return "profile image must be a .png, .jpeg, .jpg, or .svg file"
	}

	if size > 5*1024*1024 {
		return "profile image must be 5MB or smaller"
	}
	return ""
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
