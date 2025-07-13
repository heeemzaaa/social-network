package auth

import (
	"social-network/backend/models"

	"github.com/google/uuid"
)

// used only for the registration phase
func (s *AuthService) SetUserSession(user *models.User) (*models.Session, *models.ErrorJson) {
	session := &models.Session{}
	session.Token = uuid.NewString()
	errJson := s.repo.CreateUserSession(session, user)
	if errJson != nil {
		return nil, errJson
	}
	return session, nil
}

func (s *AuthService) GetSessionByTokenEnsureAuth(token string) (*models.Session, *models.ErrorJson) {
	session, err := s.repo.GetSessionbyTokenEnsureAuth(token)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *AuthService) CheckUserSession(token string) (bool, *models.Session) {
	has_session, session := s.repo.HasValidToken(token)
	if has_session {
		return true, session
	}
	return false, nil
}

func (s *AuthService) GetSessionByUserId(user_id string) (*models.Session, *models.ErrorJson) {
	session, err := s.repo.GetUserSessionByUserId(user_id)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *AuthService) DeleteSession(session *models.Session) *models.ErrorJson {
	if err := s.repo.DeleteSession(*session); err != nil {
		return err
	}
	return nil
}





