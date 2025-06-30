package middlewares

import (
	"net/http"
	"social-nework/backend/services"

	"social-network/backend/models"
	"social-network/backend/services"
	"social-network/backend/utils"
)

type LoginRegisterMiddleWare struct {
	MiddlewareHanlder http.Handler
	service           *services.AppService
}

func NewLoginMiddleware(handler http.Handler, service *services.AppService) *LoginRegisterMiddleWare {
	return &LoginRegisterMiddleWare{handler, service}
}

// could be returning a boolean but to see again
func (m *Middleware) GetAuthUserEnsureAuth(r *http.Request) (*models.Session, *models.ErrorJson) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return nil, &models.ErrorJson{Status: 401, Message: "ERROR!! Unauthorized Access"}
	}
	// check if the value of the cookie is correct and if not expired!!!
	session, errJson := m.service.GetSessionByTokenEnsureAuth(cookie.Value)
	if errJson != nil || session.IsExpired() {
		return nil, errJson
	}
	return session, nil
}

func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	_, err := m.GetAuthUserEnsureAuth(r)
	if err != nil {
		utils.WriteJsonErrors(w, *err)
		return
	}
	m.MiddlewareHanlder.ServeHTTP(w, r)
}
