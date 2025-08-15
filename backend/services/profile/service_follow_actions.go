package profile

import (
	"social-network/backend/models"
	// ns "social-network/backend/services/notification"
)

// here we will handle the logic of following a user
func (s *ProfileService) Follow(userID string, authUserID string) (*models.Notification, *models.Profile, *models.ErrorJson) {
	var profile models.Profile

	exists, err := s.repo.UserExists(userID)
	if err != nil {
		return nil, nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	if !exists {
		return nil, nil, &models.ErrorJson{Status: 400, Error: "User Id doesn't exists !"}
	}

	isFollower, err := s.repo.IsFollower(userID, authUserID)
	if err != nil {
		return nil, nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	if isFollower {
		return nil, nil, &models.ErrorJson{Status: 403, Error: "You're already a follower !"}
	}

	profile.User.Visibility, err = s.repo.Visibility(userID)
	if err != nil {
		return nil, nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	notification := &models.Notification{}
	switch profile.User.Visibility {
	case "private":
		newNotif, err := s.repo.FollowPrivate(userID, authUserID)
		if err != nil {
			return nil, nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
		}
		notification = newNotif

	case "public":
		newNotif, err := s.repo.FollowDone(userID, authUserID)
		if err != nil {
			return nil, nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
		}
		profile.IsFollower = !isFollower
		notification = newNotif

	default:
		return nil, nil, &models.ErrorJson{Status: 500, Error: "This is not a valid status of visibility"}
	}

	profile.Access, err = s.CheckProfileAccess(userID, authUserID)
	if err != nil {
		return nil, nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	profile.IsRequested, err = s.repo.IsRequested(userID, authUserID)
	if err != nil {
		return nil, nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	return notification, &profile, nil
}

// here the user can unfollow the user that he already follows
func (s *ProfileService) Unfollow(userID string, authUserID string) (*models.Profile, *models.ErrorJson) {
	var profile models.Profile

	exists, err := s.repo.UserExists(userID)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	if !exists {
		return nil, &models.ErrorJson{Status: 400, Error: "User Id doesn't exists !"}
	}

	isFollower, err := s.repo.IsFollower(userID, authUserID)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	if !isFollower {
		return nil, &models.ErrorJson{Status: 403, Error: "You're already not following this user !"}
	}

	err = s.repo.Unfollow(userID, authUserID)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	profile.Access, err = s.CheckProfileAccess(userID, authUserID)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	profile.User.Visibility, err = s.repo.Visibility(userID)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	profile.IsRequested, err = s.repo.IsRequested(userID, authUserID)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	profile.IsFollower = !isFollower

	return &profile, nil
}

func (s *ProfileService) CancelFollow(userID string, authUserID string) (*models.Profile, *models.ErrorJson) {
	var profile models.Profile

	exists, err := s.repo.UserExists(userID)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	if !exists {
		return nil, &models.ErrorJson{Status: 400, Error: "User Id doesn't exists !"}
	}

	profile.IsRequested, err = s.repo.IsRequested(userID, authUserID)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	err = s.repo.CancelFollow(userID, authUserID)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	profile.Access, err = s.CheckProfileAccess(userID, authUserID)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	profile.User.Visibility, err = s.repo.Visibility(userID)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	profile.IsFollower, err = s.repo.IsFollower(userID, authUserID)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	return &profile, nil
}
