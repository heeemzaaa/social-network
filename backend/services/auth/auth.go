package auth

import (
	"social-network/backend/models"
	"social-network/backend/repositories/auth"
)

type AuthService struct {
	repo *auth.AuthRepository
}

func NewAuthServer(authRepo *auth.AuthRepository) *AuthService {
	return &AuthService{repo: authRepo}
}

func (s *AuthService) IsLoggedInUser(token string) (*models.IsLoggedIn, *models.ErrorJson) {
	isLoggedIn, err := s.repo.IsLoggedInUser(token)
	if err != nil {
		return nil, err
	}
	return isLoggedIn, nil
}

func (s *AuthService) Login(login *models.Login) *models.ErrorJson {
	// check for valid credentials
	return nil
}

func (s *AuthService) Register(register *models.Login) *models.ErrorJson {
	return nil
}

func (s *AuthService) Logout(session *models.Session) *models.ErrorJson {
	if err := s.repo.DeleteSession(*session); err != nil {
		return err
	}
	return nil
}

func (s *AuthService) GetSession(token string) (*models.Session, *models.ErrorJson) {
	session, err := s.repo.GetSession(token)
	if err != nil {
		return nil, err
	}
	return session, nil
}
