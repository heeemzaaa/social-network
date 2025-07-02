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

// here we will check if the profile has the access to see the data of the user
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

// here we will handle the logic of following a user
func (s *ProfileService) Follow(userID string, authSessionID string) *models.ErrorJson {
	if userID == "" {
		return &models.ErrorJson{Status: 400, Message: "Invalid data !"}
	}

	authUserID, err := s.repo.GetID(authSessionID)
	if err != nil {
		return &models.ErrorJson{Status: 401, Message: fmt.Sprintf("%v", err)}
	}
	
	isFollower, errFollowers := s.repo.IsFollower(userID, authUserID)
	if errFollowers != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", errFollowers)}
	}

	if isFollower {
		return &models.ErrorJson{Status: 401, Message: "You're already a follower !"}
	}

	visibility, err := s.repo.Visibility(userID)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}

	switch visibility {
	case "private":
		err := s.repo.FollowPrivate(userID, authUserID)
		if err != nil {
			return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}
	case "public":
		err := s.repo.FollowDone(userID, authUserID)
		if err != nil {
			return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}
	default:
		return &models.ErrorJson{Status: 500, Message: "This is not a valid status of visibility"}
	}

	return nil
}

// here we will handle the case of an accepted request
func (s *ProfileService) AcceptedRequest(userID string, authUserID string) *models.ErrorJson {
	
	isFollower, errFollowers := s.repo.IsFollower(userID, authUserID)
	if errFollowers != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", errFollowers)}
	}

	if isFollower {
		return &models.ErrorJson{Status: 401, Message: "The user is already following you !"}
	}

	err := s.repo.AcceptedRequest(userID, authUserID)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}

// here we will handle the case of a refused request
func (s *ProfileService) RejectedRequest(userID string, authUserID string) *models.ErrorJson {
	isFollower, errFollowers := s.repo.IsFollower(userID, authUserID)
	if errFollowers != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", errFollowers)}
	}

	if isFollower {
		return &models.ErrorJson{Status: 401, Message: "The user is already following you !"}
	}

	err := s.repo.RejectedRequest(userID, authUserID)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}

	return nil
}

// here the user can unfollow the user that he already follows
func (s *ProfileService) Unfollow(userID string, authSessionID string) *models.ErrorJson {
	if userID == "" {
		return &models.ErrorJson{Status: 400, Message: "Invalid data !"}
	}

	authUserID, err := s.repo.GetID(authSessionID)
	if err != nil {
		return &models.ErrorJson{Status: 401, Message: fmt.Sprintf("%v", err)}
	}
	
	isFollower, errFollowers := s.repo.IsFollower(userID, authUserID)
	if errFollowers != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", errFollowers)}
	}

	if !isFollower {
		return &models.ErrorJson{Status: 401, Message: "You're already not following this user !"}
	}

	err = s.repo.Unfollow(userID, authUserID)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}

// here we will handle the logic of updating the privacy of a user
func (s *ProfileService) UpdatePrivacy(userID string , authSessionID string, wantedStatus string) *models.ErrorJson {
	if userID != authSessionID {
		return &models.ErrorJson{Status: 401, Message: "You can't update another user's profile"}
	}

	if wantedStatus == "" || (wantedStatus != "public" && wantedStatus != "private") {
		return &models.ErrorJson{Status: 400, Message: "Invalid data !"}
	}

	visibility, err := s.repo.Visibility(userID)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v" , err)}
	}

	if visibility == wantedStatus {
		return &models.ErrorJson{Status: 401, Message: fmt.Sprintf("Your account is already %s", wantedStatus)}
	}

	switch wantedStatus {
	case "public":
		err := s.repo.ToPublicAccount(userID)
		if err != nil {
			return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v" , err)}
		}

		err = s.repo.AcceptAllrequest(userID)
		if err != nil {
			return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v" , err)}
		}
	case "private":
		err := s.repo.ToPrivateAccount(userID)
		if err != nil {
			return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}
	}
	return nil
}
