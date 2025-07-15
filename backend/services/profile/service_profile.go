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
func (s *ProfileService) CheckProfileAccess(profileID string, authUserID string) (bool, *models.ErrorJson) {
	if profileID == authUserID {
		return true, nil
	}

	access, err := s.repo.CheckProfileAccess(profileID, authUserID)
	if err != nil {
		return false, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}

	return access, nil
}

// here I will have the data of the user to pass it to the handler
func (s *ProfileService) GetProfileData(profileID string, authUserID string) (*models.Profile, *models.ErrorJson) {
	var profile *models.Profile
	if profileID == "" || authUserID == "" {
		return nil, &models.ErrorJson{Status: 400, Message: "Data is invalid !"}
	}

	access, accessErr := s.CheckProfileAccess(profileID, authUserID)
	if !access && accessErr != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", accessErr)}
	}

	profile, err := s.repo.GetProfileData(profileID, access)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}

	if profileID == authUserID {
		profile.IsMyProfile = true
	}

	profile.IsFollower, err = s.repo.IsFollower(profileID, authUserID)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}

	profile.IsRequested, err = s.repo.IsRequested(profileID, authUserID)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}

	profile.Visibility, err = s.repo.Visibility(profileID)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}

	return profile, nil
}

// here I will get the list of followers
func (s *ProfileService) GetFollowers(profileID string, authUserID string) (*[]models.User, *models.ErrorJson) {
	var users *[]models.User

	if profileID == "" || authUserID == "" {
		return nil, &models.ErrorJson{Status: 400, Message: "Data is invalid !"}
	}

	access, accessErr := s.CheckProfileAccess(profileID, authUserID)
	if !access && accessErr == nil {
		return nil, &models.ErrorJson{Status: 403, Message: "you don't have the access to see the followers of this profile"}
	} else if !access && accessErr != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", accessErr)}
	}

	users, err := s.repo.GetFollowers(profileID)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}

	return users, nil
}

func (s *ProfileService) GetFollowing(profileID string, authUserID string) (*[]models.User, *models.ErrorJson) {
	var users *[]models.User

	if profileID == "" || authUserID == "" {
		return nil, &models.ErrorJson{Status: 400, Message: "Data is invalid !"}
	}

	access, accessErr := s.CheckProfileAccess(profileID, authUserID)
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
func (s *ProfileService) Follow(userID string, authUserID string) *models.ErrorJson {
	if userID == "" || authUserID == "" {
		return &models.ErrorJson{Status: 400, Message: "Invalid data !"}
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
	if userID == "" || authUserID == "" {
		return &models.ErrorJson{Status: 400, Message: "Invalid data !"}
	}

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
	if userID == "" || authUserID == "" {
		return &models.ErrorJson{Status: 400, Message: "Invalid data !"}
	}

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
func (s *ProfileService) Unfollow(userID string, authUserID string) *models.ErrorJson {
	if userID == "" || authUserID == "" {
		return &models.ErrorJson{Status: 400, Message: "Invalid data !"}
	}

	isFollower, errFollowers := s.repo.IsFollower(userID, authUserID)
	if errFollowers != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", errFollowers)}
	}

	if !isFollower {
		return &models.ErrorJson{Status: 401, Message: "You're already not following this user !"}
	}

	err := s.repo.Unfollow(userID, authUserID)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}

// here we will handle the logic of updating the privacy of a user
func (s *ProfileService) UpdatePrivacy(userID string, requestorID any, wantedStatus string) *models.ErrorJson {
	if userID == "" || requestorID == "" {
		return &models.ErrorJson{Status: 400, Message: "Invalid data !"}
	}

	if userID != requestorID {
		return &models.ErrorJson{Status: 401, Message: "You can't update another user's profile"}
	}

	if wantedStatus == "" || (wantedStatus != "public" && wantedStatus != "private") {
		return &models.ErrorJson{Status: 400, Message: "Invalid data !"}
	}

	visibility, err := s.repo.Visibility(userID)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}

	if visibility == wantedStatus {
		return &models.ErrorJson{Status: 403, Message: fmt.Sprintf("Your account is already %s", wantedStatus)}
	}

	switch wantedStatus {
	case "public":
		err := s.repo.ToPublicAccount(userID)
		if err != nil {
			return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}

		err = s.repo.AcceptAllrequest(userID)
		if err != nil {
			return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}
	case "private":
		err := s.repo.ToPrivateAccount(userID)
		if err != nil {
			return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}
	}
	return nil
}

// custom posts to each users lil2assaf
func (s *ProfileService) GetPosts(profileID string, authSessionID string) (*[]models.Post, *models.ErrorJson) {
	var posts *[]models.Post

	authUserID, err := s.repo.GetID(authSessionID)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}

	posts, err = s.repo.GetPosts(profileID, authUserID)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}

	return posts, nil
}
