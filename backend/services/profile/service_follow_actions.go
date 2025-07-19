package profile

import (
	"social-network/backend/models"
)

// here we will handle the logic of following a user
func (s *ProfileService) Follow(userID string, authUserID string) *models.ErrorJson {
	if userID == "" || authUserID == "" {
		return &models.ErrorJson{Status: 400, Error: "Invalid data !"}
	}

	isFollower, errFollowers := s.repo.IsFollower(userID, authUserID)
	if errFollowers != nil {
		return &models.ErrorJson{Status: errFollowers.Status, Error: errFollowers.Error}
	}

	if isFollower {
		return &models.ErrorJson{Status: 401, Error: "You're already a follower !"}
	}

	visibility, err := s.repo.Visibility(userID)
	if err != nil {
		return &models.ErrorJson{Status: err.Status, Error: err.Error}
	}
	switch visibility {
	case "private":
		err := s.repo.FollowPrivate(userID, authUserID)
		if err != nil {
			return &models.ErrorJson{Status: err.Status, Error: err.Error}
		}
	case "public":
		err := s.repo.FollowDone(userID, authUserID)
		if err != nil {
			return &models.ErrorJson{Status: err.Status, Error: err.Error}
		}
	default:
		return &models.ErrorJson{Status: 500, Error: "This is not a valid status of visibility"}
	}

	return nil
}

// here the user can unfollow the user that he already follows
func (s *ProfileService) Unfollow(userID string, authUserID string) *models.ErrorJson {
	if userID == "" || authUserID == "" {
		return &models.ErrorJson{Status: 400, Error: "Invalid data !"}
	}

	isFollower, errFollowers := s.repo.IsFollower(userID, authUserID)
	if errFollowers != nil {
		return &models.ErrorJson{Status: errFollowers.Status, Error: errFollowers.Error}
	}

	if !isFollower {
		return &models.ErrorJson{Status: 401, Error: "You're already not following this user !"}
	}

	err := s.repo.Unfollow(userID, authUserID)
	if err != nil {
		return &models.ErrorJson{Status: err.Status, Error: err.Error}
	}
	return nil
}
