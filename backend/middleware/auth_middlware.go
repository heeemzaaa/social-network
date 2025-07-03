package middleware

import (
	"context"
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
func (m *Middleware) GetUserSession(r *http.Request) (*models.Session, *models.ErrorJson) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return nil, &models.ErrorJson{Status: 401, Message: "ERROR!! Unauthorized Access"}
	}
	session, errJson := m.service.GetSession(cookie.Value)
	if errJson != nil {
		return nil, errJson
	}
	return session, nil
}

func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    path := r.URL.Path

    if path == "/api/auth/login" || path == "/api/auth/register" {
        if session, err := m.GetUserSession(r); err == nil && session != nil {
            handlers.WriteJsonErrors(w, models.ErrorJson{Status: 403, Message: "User has a session!! Access Forbidden"})
            return
        }
        m.MiddlewareHanlder.ServeHTTP(w, r)
        return
    }

    session, err := m.GetUserSession(r)
    if err != nil {
        handlers.WriteJsonErrors(w, *err)
        return
    }

    // Add userID to context for protected routes
    ctx := context.WithValue(r.Context(), "userID", session.UserId)
    m.MiddlewareHanlder.ServeHTTP(w, r.WithContext(ctx))
}
