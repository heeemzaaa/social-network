package profile

import (
	"social-network/backend/models"
)

// here we will check if the profile has the access to see the data of the user
func (s *ProfileService) CheckProfileAccess(profileID string, authUserID string) (bool, *models.ErrorJson) {
	if profileID == authUserID {
		return true, nil
	}

	access, err := s.repo.CheckProfileAccess(profileID, authUserID)
	if err != nil {
		return false, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	return access, nil
}

func (s *ProfileService) IsFollower(userID, authUserID string) (bool, *models.ErrorJson) {
	isFollower, err := s.repo.IsFollower(userID, authUserID)
	if err != nil {
		return false, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}
	return isFollower, nil
}
