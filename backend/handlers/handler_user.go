package handlers

import (
	"net/http"

	service "social-network/backend/services"
)

type UserHandler struct {
	service *service.AppService
}

func NewUserHandler(service *service.AppService) *UserHandler {
	return &UserHandler{service: service}
}

func (Ghandler *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
}
