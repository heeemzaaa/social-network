package profile

import (
	"social-network/backend/models"
)

// here I will get the list of followers
func (s *ProfileService) GetFollowers(profileID string, authUserID string) ([]models.User, *models.ErrorJson) {
	users := []models.User{}

	exists, err := s.repo.UserExists(profileID)
	if err != nil {
		return users, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	if !exists {
		return users, &models.ErrorJson{Status: 400, Error: "User Id doesn't exists !"}
	}

	access, accessErr := s.CheckProfileAccess(profileID, authUserID)
	if !access && accessErr == nil {
		return users, &models.ErrorJson{Status: 403, Error: "the user is not a follower !"}
	} else if !access && accessErr != nil {
		return users, &models.ErrorJson{Status: accessErr.Status, Error: accessErr.Error}
	}

	users, err = s.repo.GetFollowers(profileID)
	if err != nil {
		return []models.User{}, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	return users, nil
}

// here I will get the list of followings
func (s *ProfileService) GetFollowing(profileID string, authUserID string) ([]models.User, *models.ErrorJson) {
	users := []models.User{}

	exists, err := s.repo.UserExists(profileID)
	if err != nil {
		return users, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	if !exists {
		return users, &models.ErrorJson{Status: 400, Error: "User Id doesn't exists !"}
	}

	access, accessErr := s.CheckProfileAccess(profileID, authUserID)
	if !access && accessErr == nil {
		return users, &models.ErrorJson{Status: 403, Error: "the user is not a follower !"}
	} else if !access && accessErr != nil {
		return users, &models.ErrorJson{Status: accessErr.Status, Error: accessErr.Error}
	}

	users, err = s.repo.GetFollowing(profileID)
	if err != nil {
		return []models.User{}, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	return users, nil
}
