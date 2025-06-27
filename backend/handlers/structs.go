package handlers

import service "social-network/backend/services"

type AppService struct {
	service *service.AppService
}

// NewPostRepository creates a new repository
func NewAppRepository(service *service.AppService) *AppService {
	return &AppService{service: service}
}
