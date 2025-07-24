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

//    Login

// Middlware = checks if there is a session
// but if not we need check the body of the request o
// reset the sesssion

func (authHandler *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	login := &models.Login{}

	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		if err == io.EOF {
			// case if the body sent if empty
			utils.WriteJsonErrors(w, models.ErrorJson{
				Status: 400,
				Message: models.Login{
					LoginField: "empty login field!!",
					Password:   "empty password field!!",
				},
			})
			return
		}
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: fmt.Sprintf("%v 1", err)})
		return
	}
	user, errJson := authHandler.service.Login(login)
	if errJson != nil {
		utils.WriteJsonErrors(w, *errJson)
		return
	}

	// We are kinda sure that if the user has a token he cannot be here
	// we need now
	// before setting the session we need the actual id of the user
	// if there is a session update it
	session, errJSON := authHandler.service.SetUserSession(user)
	if errJSON != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{
			Status:  errJSON.Status,
			Message: errJSON.Message,
		})
		return
	}
	fmt.Println("setting the cookie", session.Token)
	http.SetCookie(w, &http.Cookie{
		Name:    "session",
		Value:   session.Token,
		Expires: time.Now().Add(365 * 24 * time.Hour),
		Path:    "/",
	})
	utils.WriteDataBack(w, models.UserData{
		IsLoggedIn: true,
		Id:         user.Id,
		Nickname:   user.Nickname,
	})
}
