package profile

import (
	"fmt"
	"social-network/backend/models"
	pr "social-network/backend/repositories/profile"
)

type ProfileService struct {
	repo *pr.ProfileRepository
}

func NewProfileService(repo *pr.ProfileRepository) *ProfileService {
	return &ProfileService{repo: repo}
}

func (s *ProfileService) CheckProfileAccess(profileID string, autSessionID string) (bool, *models.ErrorJson) {

	AuthUserID, err := s.repo.GetID(autSessionID)
	if err != nil {
		return false, &models.ErrorJson{Status: 401, Message: fmt.Sprintf("%v", err)}
	}

	access, err := s.repo.CheckProfileAccess(profileID, AuthUserID)
	if err != nil {
		return false, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}

	return access, nil
}

// here I will have the data of the user to pass it to the handler
func (s *ProfileService) GetProfileData(profileID string, autSessionID string) (*models.Profile, *models.ErrorJson) {
		if profileID == "" {
		return nil, &models.ErrorJson{Status: 400, Message: "Empty data !"}
	}

	var profile *models.Profile

	access, accessErr := s.CheckProfileAccess(profileID, autSessionID)
	if !access && accessErr != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", accessErr)}
	}

	profile, err := s.repo.GetProfileData(profileID, access)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return profile, nil
}

// here I will get the list of followers
func (s *ProfileService) GetFollowers(profileID string, authSessionID string) (*[]models.User, *models.ErrorJson) {
	var users *[]models.User

	access, accessErr := s.CheckProfileAccess(profileID, authSessionID)
	if !access && accessErr == nil {
		return nil, &models.ErrorJson{Status: 401, Message: "the user is not a follower !"}
	} else if !access && accessErr != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", accessErr)}
	}

	users, err := s.repo.GetFollowers(profileID)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return users, nil
}

func (s *ProfileService) GetFollowing(profileID string, authSessionID string) (*[]models.User, *models.ErrorJson) {
	var users *[]models.User

	access, accessErr := s.CheckProfileAccess(profileID, authSessionID)
	if !access && accessErr == nil {
		return nil, &models.ErrorJson{Status: 401, Message: "the user is not a follower !"}
	} else if !access && accessErr != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", accessErr)}
	}

	users, err := s.repo.GetFollowing(profileID)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}

	return users, nil
}
