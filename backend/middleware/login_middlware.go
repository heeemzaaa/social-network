package middleware

import (
	"fmt"
	"net/http"

	"social-network/backend/models"
	"social-network/backend/services/auth"
)

type LoginRegisterMiddleWare struct {
	MiddlewareHanlder http.Handler
	service           *auth.AuthService
}

func NewLoginMiddleware(handler http.Handler, service *auth.AuthService) *LoginRegisterMiddleWare {
	return &LoginRegisterMiddleWare{handler, service}
}

func (LogRegM *LoginRegisterMiddleWare) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := LogRegM.GetAuthUser(r)
	if err != nil {
		handler.WriteJsonErrors(w, models.ErrorJson{
			Status:  err.Status,
			Message: fmt.Sprintf("%v", err.Message),
		})
		return
	}

	LogRegM.MiddlewareHanlder.ServeHTTP(w, r)
}

//

func (m *LoginRegisterMiddleWare) GetAuthUser(r *http.Request) *models.ErrorJson {
	cookie, err := r.Cookie("session")
	if err != nil {
		if err == http.ErrNoCookie {
			return nil
		}
		return models.NewErrorJson(400, "ERROR!! There was an error in the Request!!")
	}
	has_session, session := m.service.CheckUserSession(cookie.Value)
	if has_session {
		if !session.IsExpired() {
			return models.NewErrorJson(403, "ERROR!! The User has a session!! Access Forbiden")
		}
	}
	return nil
}
