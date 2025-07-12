package auth

import (
	"fmt"
	"strings"

	"social-network/backend/models"
)

// so we need to check that the password entered from the API is the same stored in the database !!

func (s *AuthService) Login(login *models.Login) (*models.User, *models.ErrorJson) {
	LoginERR := models.LoginERR{}

	if strings.TrimSpace(login.LoginField) == "" {
		LoginERR.LoginField = "empty login field!"
	}
	if strings.TrimSpace(login.Password) == "" {
		LoginERR.Password = "empty password field!"
	}
	//
	if LoginERR != (models.LoginERR{}) {
		return nil, &models.ErrorJson{Status: 400, Message: LoginERR}
	}

	// we need to check also if the user has the 401 error
	// check if the password and the login are wrong both
	user, err := s.repo.GetUser(login)
	if err != nil {
		switch err.Status {
		case 401:
			return nil, err

		default:
			return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}
	}
	// if only the password
	if CheckPasswordHash(login.Password, user.Password) {
		return nil, &models.ErrorJson{
			Status: 401,
			Message: models.LoginERR{
				LoginField: "invalid login credentials!",
				Password:   "invalid login credentials!",
			},
		}
	}
	return user, nil
}



