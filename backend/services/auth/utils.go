package auth

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"regexp"
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

func isValidImg(file multipart.File) string {
	const maxImgSize = 3 << 20

	// Read up to 512 bytes for MIME detection
	buffer := make([]byte, 512)
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return "Could not read uploaded file"
	}
	// Trim buffer to actual bytes read
	buffer = buffer[:n]

	// Reset file cursor for further reading
	file.Seek(0, io.SeekStart)
	mime := http.DetectContentType(buffer)
	fmt.Printf("mime: %v\n", mime)
	allowed := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/gif":  true,
	}

	if !allowed[mime] {
		return "Unsupported file type. Please upload an image (PNG, JPG, JPGE, or GIF)"
	}

	limitedReader := io.LimitReader(file, maxImgSize+1)
	buf := new(bytes.Buffer)
	n2, err := buf.ReadFrom(limitedReader)
	if err != nil {
		return "Could not read uploaded image"
	}
	if n2 > maxImgSize {
		return "Uploaded image is too large (max 1MB)"
	}

	return ""
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	fmt.Printf("error here: \npassword = %v \nhash = %v\n", password, hash)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Println("error comparing the password and the hash", err)
	}
	return err == nil
}

func IsSessionExpired(expDate time.Time) bool {
	return expDate.Before(time.Now())
}
