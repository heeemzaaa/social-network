package middleware

import (
	"context"
	"net/http"

	"social-network/backend/models"
	"social-network/backend/utils"
)

// could be returning a boolean but to see again
func (m *Middleware) GetAuthUserEnsureAuth(r *http.Request) (*models.Session, *models.ErrorJson) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return nil, &models.ErrorJson{Status: 401, Message: "ERROR!! Unauthorized Access"}
	}
	// check if the value of the cookie is correct and if not expired!!!
	session, errJson := m.service.GetSessionByTokenEnsureAuth(cookie.Value)
	if errJson != nil {
		return nil, errJson
	}

	return session, nil
}

func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	session, err := m.GetAuthUserEnsureAuth(r)
	if err != nil {
		utils.WriteJsonErrors(w, *err)
		return
	}
	ctx := context.WithValue(r.Context(), models.NewContextKey("userID"), session.UserId)
	m.MiddlewareHanlder.ServeHTTP(w, r.WithContext(ctx))
}
