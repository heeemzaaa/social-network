package post

import (
	"social-network/backend/models"
)

func (s *PostService) HandleLike(postID string, userID string) (bool, int, *models.ErrorJson) {
	exists, err := s.repo.PostExist(postID)
	if err != nil {
		return false, 0, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	if !exists {
		return false, 0, &models.ErrorJson{Status: 400, Error: "This post ID doesn't exist !"}
	}
	
	return s.repo.HandleLike(postID, userID)
}
