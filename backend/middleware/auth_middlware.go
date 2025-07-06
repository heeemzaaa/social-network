package middleware

import (
	"context"
	"fmt"
	"net/http"

	"social-network/backend/models"
	"social-network/backend/services/auth"
	"social-network/backend/utils"
)

type Middleware struct {
	MiddlewareHanlder http.Handler
	service           any
}

func NewMiddleWare(handler http.Handler, service any) *Middleware {
	return &Middleware{handler, service}
}

// could be returning a boolean but to see again
func (m *Middleware) GetAuthUserEnsureAuth(r *http.Request) (*models.Session, *models.ErrorJson) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return nil, &models.ErrorJson{Status: 401, Message: "ERROR!! Unauthorized Access"}
	}
	// check if the value of the cookie is correct and if not expired!!!
	session, errJson := m.service.(*auth.AuthService).GetSessionByTokenEnsureAuth(cookie.Value)
	if errJson != nil || auth.IsSessionExpired(session.ExpDate) {
		return nil, errJson
	}
	return session, nil
}

func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	path := r.URL.Path
    fmt.Println("path", path)
	session, err := m.GetAuthUserEnsureAuth(r)
	fmt.Println("err", err)
	if (path == "/api/auth/login" || path == "/api/auth/register") && err == nil {
		
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 403, Message: "User has a session!! Access Forbiden"})
		return
	} else {
		fmt.Println("how!!!!!")
		if err != nil {
			fmt.Println("salaaaam", err.Status, err.Message)
			utils.WriteJsonErrors(w, *err)
			return
		}
	}

	fmt.Printf("session.UserId: %v\n", session.UserId)
	ctx := context.WithValue(r.Context(), "userID", session.UserId)
	m.MiddlewareHanlder.ServeHTTP(w, r.WithContext(ctx))
}
