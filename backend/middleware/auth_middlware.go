package middleware

import (
	"context"
	"fmt"
	"net/http"

	"social-network/backend/handlers"
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
		return nil, &models.ErrorJson{Status: 401, Error: "Unauthorized Access"}
	}
	// check if the value of the cookie is correct and if not expired!!!

	fmt.Printf("cookie.Value: %v\n", cookie.Value)
	session, errJson := m.service.GetSessionByTokenEnsureAuth(cookie.Value)
    fmt.Println("errrrrrr", errJson)
	if errJson != nil || auth.IsSessionExpired(session.ExpDate) {
		return nil, errJson
	}
	fmt.Printf("session: %v\n", session)
	return session, nil
}

func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	session, err := m.GetAuthUserEnsureAuth(r)
	if path == "/api/auth/login" || path == "/api/auth/register" {
		if err == nil {
			// Already has a session, prevent login/register again
			handlers.WriteJsonErrors(w, models.ErrorJson{Status: 403, Message: "User already logged in. Access forbidden"})
			return
		}

		// Allow access to login/register if no session
		m.MiddlewareHanlder.ServeHTTP(w, r)
		return
	}

	if err != nil && path != "/api/auth/islogged" {
		handlers.WriteJsonErrors(w, *err)
		return
	}

	if session != nil {
		ctx := context.WithValue(r.Context(), "userID", session.UserId)
		m.MiddlewareHanlder.ServeHTTP(w, r.WithContext(ctx))
	} else {
		m.MiddlewareHanlder.ServeHTTP(w, r)
	}
}
