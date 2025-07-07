package profile

import (
	"fmt"
	"net/http"
	"social-network/backend/models"
)

func GetSessionID(r *http.Request) (string, *models.ErrorJson) {
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println(err)
		return "", &models.ErrorJson{Status: 401, Message: "Unauthorized !"}
	}
	return cookie.Value, nil
}
