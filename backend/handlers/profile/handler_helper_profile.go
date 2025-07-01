package profile

import (
	"net/http"
	"social-network/backend/models"
)

func GetSessionID(r *http.Request) (string, *models.ErrorJson) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return "", &models.ErrorJson{Status: 401, Message: "Unauthorized !"}
	}
	return cookie.Value, nil
}
