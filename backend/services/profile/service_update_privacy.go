package profile

import (
	"fmt"

	"social-network/backend/models"
	ns "social-network/backend/services/notification"
)

// here we will handle the logic of updating the privacy of a user
func (s *ProfileService) UpdatePrivacy(userID string, requestorID string, wantedStatus string, NS *ns.NotificationService) (*models.Profile, *models.ErrorJson) {
	profile := &models.Profile{}
	exists, err := s.repo.UserExists(userID)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	if !exists {
		return nil, &models.ErrorJson{Status: 400, Error: "User Id doesn't exists !"}
	}

	if userID != requestorID {
		return nil, &models.ErrorJson{Status: 403, Error: "You can't update another user's profile"}
	}

	if wantedStatus == "" || (wantedStatus != "public" && wantedStatus != "private") {
		return nil, &models.ErrorJson{Status: 400, Error: "Invalid wanted visibility !"}
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

		// get all notifications that has type follow-private and toggle status "accept"
		all, errJson := NS.GetAllNotifService(userID, "follow-private")
		if errJson != nil {
			return nil, errJson
		}
		errJson = NS.ToggleAllStaus(all, "accept", "follow-private")
		if errJson != nil {
			return nil, errJson
		}

		profile, err = s.GetProfileData(userID, requestorID)
		if err != nil {
			return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
		}
	case "private":
		err := s.repo.ToPrivateAccount(userID)
		if err != nil {
			return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
		}
	}
	return profile, nil
}
