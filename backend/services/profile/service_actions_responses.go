package profile

import (
	"social-network/backend/models"
)

// here we will handle the case of an accepted request
func (s *ProfileService) AcceptedRequest(userID string, authUserID string) *models.ErrorJson {
	if userID == "" || authUserID == "" {
		return &models.ErrorJson{Status: 400, Error: "Invalid data !"}
	}

	isFollower, errFollowers := s.repo.IsFollower(userID, authUserID)
	if errFollowers != nil {
		return &models.ErrorJson{Status: errFollowers.Status, Error: errFollowers.Error}
	}

	// still need to check if its 409 or 403
	if isFollower {
		return &models.ErrorJson{Status: 403, Error: "The user is already following you !"}
	}

	err := s.repo.AcceptedRequest(userID, authUserID)
	if err != nil {
		return &models.ErrorJson{Status: err.Status, Error: err.Error}
	}
	return nil
}

// here we will handle the case of a refused request
func (s *ProfileService) RejectedRequest(userID string, authUserID string) *models.ErrorJson {
	if userID == "" || authUserID == "" {
		return &models.ErrorJson{Status: 400, Error: "Invalid data !"}
	}

	isFollower, errFollowers := s.repo.IsFollower(userID, authUserID)
	if errFollowers != nil {
		return &models.ErrorJson{Status: errFollowers.Status, Error: errFollowers.Error}
	}

	if isFollower {
		return &models.ErrorJson{Status: 403, Error: "The user is already following you !"}
	}

	err := s.repo.RejectedRequest(userID, authUserID)
	if err != nil {
		return &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	return nil
}
