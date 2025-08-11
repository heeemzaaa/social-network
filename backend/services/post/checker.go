package post

import (
	"social-network/backend/models"
)

func (s *PostService) PostPrivacy(postID string) (string, *models.ErrorJson) {
	exists, err := s.repo.PostExist(postID)
	if err != nil {
		return "", &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	if !exists {
		return "", &models.ErrorJson{Status: 400, Error: "Invalid post id !"}
	}

	privacy, err := s.repo.PostPrivacy(postID)
	if err != nil {
		return "", &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	return privacy, nil
}

func (s *PostService) PrivatePostAccess(userID, postID string) (bool, *models.ErrorJson) {
	return s.repo.PrivatePostAccess(userID, postID)
}