package services

import (
	"social-network/backend/models"
)

func (s *PostService) CreatePost(post *models.Post) (*models.Post, *models.ErrorJson) {
	return s.repo.CreatePost(post)
}
