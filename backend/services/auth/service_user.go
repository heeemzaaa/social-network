package auth

import (
	"net/http"

	"social-network/backend/models"
)

func (s *AuthService) GetUser(login *models.Login) (*models.User, *models.ErrorJson) {
	user, err := s.repo.GetUser(login)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Message: err.Message}
	}
	return user, nil
}


func (s *AuthService) UserExists(id int) (bool, *models.ErrorJson) {
	exists, errJson := s.repo.UserExists(id)
	if errJson != nil {
		return false, errJson
	}
	return exists, nil
}

func (service *AuthService) GetUserIdFromSession(r *http.Request) string {
	cookie, _ := r.Cookie("session")
	session, _ := service.GetSessionByTokenEnsureAuth(cookie.Value)
	return session.UserId
}


