package profile

import (
	"fmt"

	"social-network/backend/models"
)

// here we will handle the logic of updating the privacy of a user
func (s *ProfileService) UpdatePrivacy(userID string, requestorID string, wantedStatus string) (*models.Profile, *models.ErrorJson) {
	profile := &models.Profile{}
	if userID == "" || requestorID == "" {
		return nil, &models.ErrorJson{Status: 400, Error: "Invalid data !"}
	}

	if userID != requestorID {
		return nil, &models.ErrorJson{Status: 403, Error: "You can't update another user's profile"}
	}

	if wantedStatus == "" || (wantedStatus != "public" && wantedStatus != "private") {
		return nil, &models.ErrorJson{Status: 400, Error: "Invalid data !"}
	}

	visibility, err := s.repo.Visibility(userID)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	if visibility == wantedStatus {
		return nil, &models.ErrorJson{Status: 403, Error: fmt.Sprintf("Your account is already %s", wantedStatus)}
	}

	switch wantedStatus {
	case "public":
		err := s.repo.ToPublicAccount(userID)
		if err != nil {
			return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
		}

		err = s.repo.AcceptAllrequest(userID)
		if err != nil {
			return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
		}

		profile, err = s.GetProfileData(userID, requestorID)
		if err != nil {
			return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
		}		
		fmt.Println("profile: ", profile)
	case "private":
		err := s.repo.ToPrivateAccount(userID)
		if err != nil {
			return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
		}
	}
	return profile, nil
}
