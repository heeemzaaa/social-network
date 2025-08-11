package auth

import (
	"net/http"

	"social-network/backend/models"
	"social-network/backend/services/auth"
	"social-network/backend/utils"
)

type UserData struct {
	service *auth.AuthService
}

func NewLoggedInHanlder(service *auth.AuthService) *UserData {
	return &UserData{service: service}
}

func (loggedin *UserData) GetLoggedIn(w http.ResponseWriter, r *http.Request) {
	user_data := &models.UserData{}
	cookie, err := r.Cookie("session")
	if err != nil {
		user_data.IsLoggedIn = false
		utils.WriteDataBack(w, user_data)
		return
	}

	userData, errJson := loggedin.service.IsLoggedInUser(cookie.Value)

	if errJson != nil {
		user_data.IsLoggedIn = false
		utils.WriteDataBack(w, user_data)
		return
	}
	user_data.IsLoggedIn = true
	user_data.Id = userData.Id
	user_data.FullName = userData.FullName
	user_data.Nickname = userData.Nickname
	user_data.Token = cookie.Value
	utils.WriteDataBack(w, user_data)
}

func (loggedin *UserData) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch {
	case r.Method != http.MethodGet && r.URL.Path == "/api/loggedin":
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "ERROR!! Method Not Allowed!"})
		return
	case r.Method == http.MethodGet && r.URL.Path == "/api/loggedin":
		loggedin.GetLoggedIn(w, r)
		return
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "ERROR!! Page Not Found!"})
		return

	}
}
