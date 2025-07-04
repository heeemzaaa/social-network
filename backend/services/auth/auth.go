// authentication serevice.
package auth

import (
	"fmt"
	"mime/multipart"
	"reflect"
	"time"

	"social-network/backend/models"
	"social-network/backend/repositories/auth"
	"social-network/backend/services"

	"github.com/google/uuid"
)

type AuthService struct {
	repo *auth.AuthRepository
}

func NewAuthServer(authRepo *auth.AuthRepository) *AuthService {
	return &AuthService{repo: authRepo}
}

func (s *AuthService) IsLoggedInUser(token string) (*models.IsLoggedIn, *models.ErrorJson) {
	isLoggedIn, err := s.repo.IsLoggedInUser(token)
	if err != nil {
		return nil, err
	}
	return isLoggedIn, nil
}

func (s *AuthService) Login(login *models.Login) *models.ErrorJson {
	// check for valid credentials
	return nil
}

func (s *AuthService) Logout(session *models.Session) *models.ErrorJson {
	if err := s.repo.DeleteSession(*session); err != nil {
		return err
	}
	return nil
}

func (s *AuthService) GetSession(token string) (*models.Session, *models.ErrorJson) {
	session, err := s.repo.GetSession(token)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *AuthService) Register(user *models.User, file multipart.File) *models.ErrorJson {
	// fmt.Println("inside the register service: ", user)
	jsonError := &models.ErrorJson{}
	fmt.Printf("initial jsonError: %v\n", jsonError)

	// data validation
	s.validateUserData(user, jsonError)
	fmt.Printf("jsonError: %v\n", jsonError)
	if !reflect.DeepEqual(*jsonError, models.ErrorJson{}) {
		fmt.Println("errors while validating user data")
		jsonError.Status = 400
		return jsonError
	}
	user.Id = uuid.New().String()

	dst, err := services.CreateFile("profile", user.ProfileImage)
	if err != nil {
		fmt.Println("error creating file : ", err.Error())
		jsonError.Status = 500
		jsonError.Error = "Error saving profile image"
	}
	defer dst.Close()

	// Copy the uploaded file to the destination file
	if _, err := dst.ReadFrom(file); err != nil {
		fmt.Println("error reading file : ", err.Error())
		jsonError.Status = 500
		jsonError.Error = "Error saving profile image: " + err.Error()
	}

	if !reflect.DeepEqual(*jsonError, models.ErrorJson{}) {
		fmt.Println("heeeere")
		jsonError.Status = 500
		return jsonError
	}
	hash, err := HashPassword(user.Password)

	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v just to verify 1", err)}
	} else {
		user.Password = hash
	}

	errJson := s.repo.CreateUser(user)
	if errJson != nil {
		return errJson
	}

	return nil
}

func (s *AuthService) validateUserData(user *models.User, jsonError *models.ErrorJson) {
	fmt.Println("inside the user data validation function :| ")

	userErrorJson := models.User{}

	if user.FirstName == "" {
		userErrorJson.FirstName = "first name is required"
	}

	// if user.FirstName != "" {
	// 	userErrorJson.FirstName = "first name is required"
	// }

	if user.LastName == "" {
		userErrorJson.LastName = "last name is required"
	}

	if user.BirthDate == "" {
		userErrorJson.BirthDate = "birth date is required"
	}

	if user.Email == "" {
		userErrorJson.Email = "email is required"
	}

	if user.Password == "" {
		userErrorJson.Password = "password is required"
	}

	// Add more validation as needed (e.g., email format, password strength)
	if !reflect.DeepEqual(userErrorJson, models.User{}) {
		jsonError.Message = userErrorJson
	}
}

func (s *AuthService) GetUser(login *models.Login) (*models.User, *models.ErrorJson) {
	user, err := s.repo.GetUser(login)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Message: err.Message}
	}
	return user, nil
}

func (s *AuthService) SetUserSession(user *models.User) (*models.Session, *models.ErrorJson) {
	session := &models.Session{}
	session.Token = uuid.NewString()
	session.ExpDate = time.Now().Add(24 * time.Hour)
	errJson := s.repo.CreateUserSession(session, user)
	if errJson != nil {
		return nil, errJson
	}
	return session, nil
}

