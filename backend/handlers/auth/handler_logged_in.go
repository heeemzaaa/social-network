package auth

import (
	"net/http"

	"social-network/backend/models"
	"social-network/backend/utils"
)

func (d *UserData) GetLoggedIn(w http.ResponseWriter, r *http.Request) {
	user_data := &models.UserData{}
	if r.Method != http.MethodGet {
		utils.WriteJsonErrors(w, *&models.ErrorJson{Status: 405, Message: "ERROR!! Method Not Allowed!!"})
		return
	}

	cookie, err := r.Cookie("session")
	if err != nil {
		user_data.IsLoggedIn = false
		utils.WriteDataBack(w, user_data)
		return
	}

	userData, errJson := d.service.IsLoggedInUser(cookie.Value)

	if errJson != nil {
		user_data.IsLoggedIn = false
		utils.WriteDataBack(w, user_data)
		return
	}
	user_data.IsLoggedIn = true
	user_data.Id = userData.Id
	user_data.Nickname = userData.Nickname
	utils.WriteDataBack(w, user_data)
}
