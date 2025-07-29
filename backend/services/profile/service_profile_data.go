package profile

import (
	"social-network/backend/models"
)

// here I will have the data of the user to pass it to the handler
func (s *ProfileService) GetProfileData(profileID string, authUserID string) (*models.Profile, *models.ErrorJson) {
	var profile *models.Profile
	
	exists, err := s.repo.UserExists(profileID)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	if !exists {
		return nil, &models.ErrorJson{Status: 400, Error: "User Id doesn't exists !"}
	}

	access, accessErr := s.CheckProfileAccess(profileID, authUserID)
	if !access && accessErr != nil {
		return nil, &models.ErrorJson{Status: accessErr.Status, Error: accessErr.Error}
	}

	profile, err = s.repo.GetProfileData(profileID, access)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	if profileID == authUserID {
		profile.IsMyProfile = true
	}

	profile.IsFollower, err = s.repo.IsFollower(profileID, authUserID)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	profile.IsRequested, err = s.repo.IsRequested(profileID, authUserID)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	profile.Access = access

	profile.User.Visibility, err = s.repo.Visibility(profileID)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}
	return profile, nil
}
