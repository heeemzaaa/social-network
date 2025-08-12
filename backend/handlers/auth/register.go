package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"social-network/backend/models"
	"social-network/backend/utils"
)

func (authHandler *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	data := r.FormValue("data")
	if err := json.Unmarshal([]byte(data), &user); err != nil {
		if err == io.EOF || user == (models.User{}) {
			utils.WriteJsonErrors(w, models.ErrorJson{
				Status: 400,
				Message: models.User{
					FirstName: "First Name is required",
					LastName:  "Last Name is required",
					BirthDate: "Birthdate is required",
					Email:     "email is required",
					Password:  "password is required",
				},
			})
			return
		}

		utils.WriteJsonErrors(w, models.ErrorJson{
			Status:  400,
			Message: "ERROR!! Can not Unmarshal the data!",
		})
		return
	}
	path, errUploadImg := utils.HanldeUploadImage(r, "profile_img", "avatars/profile")
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

	fmt.Println("USER", user)
	// before setting the session we need the actual id of the user
	userData, errJson := authHandler.service.GetUser(&models.Login{LoginField: user.Email})
	fmt.Printf("userData: %v\n", userData)
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
		Name:    "session",
		Value:   session.Token,
		Expires: time.Now().Add(365 * 24 * time.Hour),
		Path:    "/",
	})
	// we don't need to write back the data for the repsonse ( sentitive data ;)
	utils.WriteDataBack(w, models.UserData{
		IsLoggedIn: true,
		Id:         userData.Id,
		Nickname:   userData.Nickname,
	})
}
