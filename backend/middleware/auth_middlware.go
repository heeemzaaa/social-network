package middleware

import (
	"context"
	"net/http"

	"social-network/backend/models"
	"social-network/backend/services/auth"
)

type Middleware struct {
	MiddlewareHanlder http.Handler
	service           *auth.AuthService
}

func NewMiddleWare(handler http.Handler, service *auth.AuthService) *Middleware {
	return &Middleware{handler, service}
}

// could be returning a boolean but to see again
func (m *Middleware) GetAuthUserEnsureAuth(r *http.Request) (*models.Session, *models.ErrorJson) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return nil, &models.ErrorJson{Status: 401, Message: "ERROR!! Unauthorized Access"}
	}
	// check if the value of the cookie is correct and if not expired!!!
	session, errJson := m.service.GetSessionByTokenEnsureAuth(cookie.Value)
	if errJson != nil || auth.IsSessionExpired(session.ExpDate) {
		return nil, errJson
	}
	return session, nil
}

func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	// path := r.URL.Path

	// _, err := m.GetUserSession(r)
	// if (path == "/auth/login" || path == "/auth/register") && err == nil {
	// 	handlers.WriteJsonErrors(w, models.ErrorJson{Status: 403, Message: "User has a session!! Access Forbiden"})
	// 	return
	// } else {
	// 	if err != nil {
	// 		handlers.WriteJsonErrors(w, *err)
	// 		return
	// 	}
	// }

	ctx := context.WithValue(r.Context(), "userID", "123456")
	m.MiddlewareHanlder.ServeHTTP(w, r.WithContext(ctx))
}
