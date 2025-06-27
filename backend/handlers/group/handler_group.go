package group

import (
	"net/http"

	service "social-network/backend/services"
)

type GroupHanlder struct {
	service *service.AppService
}

func NewGroupHandler(service *service.AppService) *GroupHanlder {
	return &GroupHanlder{service: service}
}

func (Ghandler *GroupHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
}
