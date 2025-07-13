package auth

import "social-network/backend/models"

func (s *AuthService) IsLoggedInUser(token string) (*models.UserData, *models.ErrorJson) {
	user, err := s.repo.IsLoggedInUser(token)
	if err != nil {
		return nil, err
	}
	return user, nil
}
