package profile

import (
	"social-network/backend/models"
	ns "social-network/backend/services/notification"
)

// here we will handle the logic of following a user
func (s *ProfileService) Follow(userID string, authUserID string, NS *ns.NotificationService) (*models.Profile, *models.ErrorJson) {
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

	if isFollower {
		return nil, &models.ErrorJson{Status: 403, Error: "You're already a follower !"}
	}

	profile.User.Visibility, err = s.repo.Visibility(userID)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	data := models.Notif{
		SenderId:       authUserID,
		RecieverId:     userID,
		// SenderFullName: profile.User.FullName,
	}

	switch profile.User.Visibility {
	case "private":
		err := s.repo.FollowPrivate(userID, authUserID)
		if err != nil {
			return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
		}

		data.Type = "follow-private"

		// insert new private notification for recieverId = userID
		errJson := NS.PostService(data)
		if errJson != nil {
			// return nil, &models.ErrorJson{Status: err.Status, Error: err.Error, Message: err.Message}
		}

	case "public":
		err := s.repo.FollowDone(userID, authUserID)
		if err != nil {
			return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
		}
		profile.IsFollower = !isFollower

		// data := models.Notif{
		// 	SenderId: authUserID,
		// 	RecieverId: userID,
		// 	Type: "follow-public",
		// }
		data.Type = "follow-public"

		// insert new public notification for recieverId = userID
		errJson := NS.PostService(data)
		if errJson != nil {
			return nil, errJson
		}

		// insert new public notification for recieverId = userID
		NS.PostService(models.Notif{SenderId: authUserID, RecieverId: userID, Type: "follow-public"})

	default:
		return nil, &models.ErrorJson{Status: 500, Error: "This is not a valid status of visibility"}
	}

	profile.Access, err = s.CheckProfileAccess(userID, authUserID)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	profile.IsRequested, err = s.repo.IsRequested(userID, authUserID)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	return &profile, nil
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

func (s *ProfileService) CancelFollow(userID string, authUserID string, NS *ns.NotificationService) (*models.Profile, *models.ErrorJson) {
	var profile models.Profile

		exists, err := s.repo.UserExists(userID)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	if !exists {
		return nil, &models.ErrorJson{Status: 400, Error: "User Id doesn't exists !"}
	}

	err = s.repo.CancelFollow(userID, authUserID)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	/// /// HERE REMOVE NOTIFICATION FOLLOW PRIVATE /// ///
	errJson := NS.DeleteService(userID, authUserID, "follow-private", "later")
	if errJson != nil {
		return nil, errJson
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

	profile.IsRequested, err = s.repo.IsRequested(userID, authUserID)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	return &profile, nil
}
