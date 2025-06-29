package auth

import (
	"net/http"

	"social-network/backend/services"
)

type AuthHandler struct {
	service *services.AppService
}

func NewAuthHandler(service *services.AppService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (Ghandler *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	

}
