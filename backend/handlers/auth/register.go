package auth

import (
	"encoding/json"
	"io"
	"net/http"

	"social-network/backend/models"
	"social-network/backend/utils"
)

func (authHandler *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	data := r.FormValue("data")
	if err := json.Unmarshal([]byte(data), &user); err != nil {
		if err == io.EOF {
			utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: &models.RegisterError{
				Nickname:      " empty username field",
				Age:           " empty username field",
				Gender:        " empty gender field",
				FirstName:     " empty first Name field",
				LastName:      " empty lastName field",
				Email:         " empty email field",
				Password:      " empty password  field",
				VerifPassword: " empty verification Password  field",
			}})
			return
		}

		utils.WriteJsonErrors(w, models.ErrorJson{
			Status:  400,
			Message: "ERROR!! Can not Unmarshal the data!",
		})
		return
	}
	path, errUploadImg := utils.HanldeUploadImage(r, "group", "groups", true)
	if errUploadImg != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errUploadImg.Status, Message: errUploadImg.Message})
		return
	}

	user.ImagePath = path
	errJson := authHandler.service.Register(&user)
	if errJson != nil {
		utils.WriteJsonErrors(w, *errJson)
		return
	}

	// before setting the session we need the actual id of the user
	userData, errJson := authHandler.service.GetUser(&models.Login{LoginField: user.Nickname})
	if errJson != nil {
		utils.WriteJsonErrors(w, *errJson)
		return
	}
	// var login = models.Login{LoginField: user.Nickname}
	session, err_ := authHandler.service.SetUserSession(userData)
	if err_ != nil {
		utils.WriteJsonErrors(w, *err_)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: session.Token,
		Path:  "/",
	})
	// we don't need to write back the data for the repsonse ( sentitive data ;)
	utils.WriteDataBack(w, models.UserData{
		IsLoggedIn: true,
		Id:         userData.Id,
		Nickname:   userData.Nickname,
	})
}
