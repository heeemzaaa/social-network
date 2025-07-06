// authentication serevice.
package auth

import (
	"fmt"
	"mime/multipart"
	"os"
	"reflect"
	"strings"
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

func (s *AuthService) Login(login *models.Login) (*models.User, *models.ErrorJson) {
	LoginERR := models.Login{}

	if strings.TrimSpace(login.LoginField) == "" {
		LoginERR.LoginField = "empty login field!"
	}
	if strings.TrimSpace(login.Password) == "" {
		LoginERR.Password = "empty password field!"
	}
	if LoginERR != (models.Login{}) {
		return nil, &models.ErrorJson{Status: 400, Message: LoginERR}
	}

	user, err := s.repo.GetUser(login)
	if err != nil {
		switch err.Status {
		case 401:
			return nil, err
		default:
			return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("asfdasdf %v", err)}
		}
	}

	if !CheckPasswordHash(login.Password, user.Password) {
		return nil, models.NewErrorJson(401, "Invalid login credentials.", nil)
	}
	return user, nil
}

func (s *AuthService) Logout(session *models.Session) *models.ErrorJson {
	if err := s.repo.DeleteSession(*session); err != nil {
		return err
	}
	return nil
}

func (s *AuthService) Register(user *models.User, file multipart.File) *models.ErrorJson {
	fmt.Printf("user: %v\n", user)

	// data validation
	jsonError := s.validateUserData(user)
	if jsonError != nil {
		return jsonError
	}

	user.Id = uuid.New()
	// Image handling:
	if user.ProfileImage != "" {
		ImgName := user.Id.String()+ "." + strings.Split(user.ProfileImage, ".")[1]

		imgPath, err := services.UploadImage(file, ImgName)
		if err != nil {
			return models.NewErrorJson(500, "Error saving profile image", nil)
		}
		user.ProfileImage = imgPath
	}

	hash, err := HashPassword(user.Password)

	if err != nil {
		return models.NewErrorJson(500, err.Error(), nil)
	} else {
		user.Password = hash
	}

	errJson := s.repo.CreateUser(user)
	if errJson != nil {
		os.Remove(user.ProfileImage)
		return errJson
	}

	return nil
}

func (s *AuthService) validateUserData(user *models.User) *models.ErrorJson {
	fmt.Println("inside the user data validation function :| ")

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
	userErrorJson.ProfileImage = isValidImg(user.ProfileImage, user.ProfileImgSize)

	// Add more validation as needed (e.g., email format, password strength)
	if !reflect.DeepEqual(userErrorJson, models.User{}) {
		return &models.ErrorJson{Status: 400, Message: userErrorJson}
	}

	return nil
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

func (s *AuthService) CreateOrUpdateSession(user *models.User) (*models.Session, *models.ErrorJson) {
	session, errJson := s.GetSessionByUserId(user.Id.String())
	if errJson != nil {
		return nil, errJson
	}
	if session != nil {
		new_session, errJSON := s.UpdateUserSession(session)
		if errJSON != nil {
			return nil, errJSON
		}
		return new_session, nil

	} else {
		new_session, errJSON := s.SetUserSession(user)
		if errJSON != nil {
			return nil, errJSON
		}
		return new_session, nil
	}
}

func (s *AuthService) GetSessionByUserId(user_id string) (*models.Session, *models.ErrorJson) {
	session, err := s.repo.GetUserSessionByUserId(user_id)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *AuthService) UpdateUserSession(session *models.Session) (*models.Session, *models.ErrorJson) {
	new_session := models.NewSession()
	new_session.Token = uuid.NewString()
	new_session.ExpDate = time.Now().Add(24 * time.Hour)
	if err := s.repo.UpdateSession(session, new_session); err != nil {
		return nil, err
	}
	return new_session, nil
}

func (s *AuthService) GetSessionByTokenEnsureAuth(token string) (*models.Session, *models.ErrorJson) {
	session, err := s.repo.GetSessionbyTokenEnsureAuth(token)
	if err != nil {
		return nil, err
	}
	return session, nil
}
