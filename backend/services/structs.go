package services

import repository "social-network/backend/repositories"

type AppService struct {
	repository *repository.AppRepository
}

// NewPostRepository creates a new repository
func NewAppService(repository *repository.AppRepository) *AppService {
	return &AppService{repository: repository}
}



