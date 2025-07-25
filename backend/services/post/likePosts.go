package services

import (
	"social-network/backend/models"
)

func (s *PostService) HandleLike(postID string, userID string) (bool, int, *models.ErrorJson) {
	return s.repo.HandleLike(postID, userID)
}
