package auth

import "social-network/backend/repositories/auth"

type AuthService struct {
	repo *auth.AuthRepository
}

// NewPostService creates a new service
func NewAuthService(repo *auth.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}
