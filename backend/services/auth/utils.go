package auth

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func isValidName(name string) error {
	if len(name) > 50 {
		return errors.New("name must be 50 characters or less")
	}
	if !regexp.MustCompile(`^[a-zA-Z\s-]+$`).MatchString(name) {
		return errors.New("name can only contain letters, spaces, or hyphens")
	}
	return nil
}

func isValidBirthDate(date string) error {
	if strings.TrimSpace(date) == "" {
		return errors.New("BirthDate is required")
	}

	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return errors.New("birth date must be in YYYY-MM-DD format and valid")
	}

	return nil
}

func (s *AuthService) isValidEmail(email string) error {
	if strings.TrimSpace(email) == "" {
		return errors.New("email is required")
	}
	if len(email) > 100 {
		return errors.New("email must be 100 characters or less")
	}
	// Basic email regex
	if !regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(email) {
		return errors.New("invalid email format")
	}

	_, has_email, _ := s.repo.GetItem("users", "email", email)
	if has_email {
		return errors.New("email already in use")
	}

	return nil
}

func (s *AuthService) isValidNickname(nickname string) error {
	if strings.TrimSpace(nickname) == "" {
		return nil
	}

	if len(nickname) < 3 {
		return errors.New("nickname must be 3 characters or higher")
	}

	if len(nickname) > 30 {
		return errors.New("nickname must be 30 characters or less")
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9_-]+$`).MatchString(nickname) {
		return errors.New("nickname can only contain letters, numbers, underscores, or hyphens")
	}

	_, has_nickname, _ := s.repo.GetItem("users", "nickname", nickname)
	if has_nickname {
		return errors.New("username already exists")
	}

	return nil
}

func isValidAboutme(aboutme string) error {
	if strings.TrimSpace(aboutme) == "" {
		return nil
	}
	if len(aboutme) > 500 {
		return errors.New("about me must be 500 characters or less")
	}
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Println("error comparing the password and the hash", err)
	}
	return err == nil
}
