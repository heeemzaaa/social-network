package images

import (
	"social-network/backend/models"
)

func (s *ServiceImages) AccessToPostImage(userID string, creatorID string, postID string) (bool, *models.ErrorJson) {
	if userID == creatorID {
		return true , nil
	}
	isFollower, err := s.profileService.IsFollower(creatorID, userID)
	if err != nil {
		return false, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	postPrivacy, err := s.postService.PostPrivacy(postID)
	if err != nil {
		return false, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	profileVisibility, err := s.profileService.Visibility(creatorID)
	if err != nil {
		return false, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	switch isFollower {
	case true:
		// here I should check the privacy of the post to give the image access
		switch postPrivacy {
		case "private":
			access, err := s.postService.PrivatePostAccess(userID, postID)
			if err != nil {
				return false, &models.ErrorJson{Status: err.Status, Error: err.Error}
			}
			return access, nil
		case "almost-private", "public":
			return  true, nil
		}
	case false:
		// here I should check the privacy of the post and the profile
		switch profileVisibility {
		case "private":
			return  false, nil
		case "public":
			switch postPrivacy {
			case "private", "almost private":
				return false, nil
			case "public":
				return true, nil
			}
		}
	}
	return false, &models.ErrorJson{Status: 400, Error: "Invalid data !"}
}
