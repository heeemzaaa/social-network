package chat

import (
	"fmt"
	"net/http"

	"social-network/backend/models"
)


func (s *ChatService) GetSessionByTokenEnsureAuth(token string) (*models.Session, *models.ErrorJson) {
	session, err := s.repo.GetSessionbyTokenEnsureAuth(token)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}
	return session, nil
}


// get the userID from the session after pass it to the repo
func (service *ChatService) GetUserIdFromSession(r *http.Request) (string, *models.ErrorJson) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return "", &models.ErrorJson{Status: 401, Error: "", Message: fmt.Sprintf("%v", err)}
	}

	userID, errQuery := service.repo.GetUserIdFromSession(cookie.Value)
	if errQuery != nil {
		return "", &models.ErrorJson{Status: errQuery.Status, Error: errQuery.Error}
	}

	return userID, nil
}
