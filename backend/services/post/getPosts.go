package post

import (
	"social-network/backend/models"
)

func (s *PostService) GetAllPosts(userID string) ([]models.Post, *models.ErrorJson) {
	return s.repo.GetAllPosts(userID)
}
