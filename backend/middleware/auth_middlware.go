package middleware

import (
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
	session, errJson := m.service.GetSessionByTokenEnsureAuth(cookie.Value)
	if errJson != nil {
		return nil, errJson
	}
	return session, nil
}

func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	_, err := m.GetAuthUserEnsureAuth(r)
	if err != nil {
		return
	}
	m.MiddlewareHanlder.ServeHTTP(w, r)
}
