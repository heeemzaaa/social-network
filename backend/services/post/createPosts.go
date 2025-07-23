package services

import "social-network/backend/models"

func (s *PostService) CreatePost(post *models.Post) error {
	return s.repo.CreatePost(post)
}
