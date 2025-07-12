package middleware

import (
	"fmt"
	"net/http"

	"social-network/backend/models"
	"social-network/backend/utils"
)

// Login and Register middlwares
// e7m e7m wash hakka hadshi khassu ykun ??
// allahu a3laaam

func (LogRegM *LoginRegisterMiddleWare) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	err := LogRegM.GetAuthUser(r)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{
			Status:  err.Status,
			Message: fmt.Sprintf("%v", err.Message),
		})
		return
	}

	LogRegM.MiddlewareHanlder.ServeHTTP(w, r)
}

// we need to set the context heeeeere !!!!
// and pass the data of the user to it

func (LogRegM *LoginRegisterMiddleWare) GetAuthUser(r *http.Request) *models.ErrorJson {
	cookie, err := r.Cookie("session")
	if err != nil {
		if err == http.ErrNoCookie {
			return nil
		}
		return &models.ErrorJson{Status: 400, Message: "ERROR!! There was an error in the Request!!"}
	}
	has_session, _ := LogRegM.service.CheckUserSession(cookie.Value)
	if has_session {
		return &models.ErrorJson{Status: 403, Message: "ERROR!! The User has a session!! Access Forbiden"}
	}
	return nil
}
