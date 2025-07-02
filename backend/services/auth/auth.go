package auth

import (
	"social-network/backend/models"
	"social-network/backend/repositories/auth"
)

type AuthService struct {
	repo auth.AuthRepository
}

func (s *AuthService) GetSessionByTokenEnsureAuth(token string) (*models.Session, *models.ErrorJson) {
	session, err := s.repo.GetSessionbyTokenEnsureAuth(token)
	if err != nil {
		return nil, err
	}
	return session, nil
}
