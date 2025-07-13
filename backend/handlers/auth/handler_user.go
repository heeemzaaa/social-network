package auth

import (
	"net/http"

	"social-network/backend/models"
	"social-network/backend/services/auth"
	"social-network/backend/utils"
)

type AuthHandler struct {
	service *auth.AuthService
}

func NewAuthHandler(service *auth.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// add the endpoint of getusers
func (authHandler *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch {
	case r.Method == http.MethodPost && r.URL.Path == "/api/auth/login":
		authHandler.Login(w, r)
		return
	case r.Method == http.MethodPost && r.URL.Path == "/api/auth/register":
		authHandler.Register(w, r)
		return
	case r.Method != http.MethodPost:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "ERROR!! Method Not allowed!!"})
		return
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "ERROR!! Page not Found!!"})
		return
	}
}
