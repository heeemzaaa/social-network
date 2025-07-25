package services

import (
	"social-network/backend/models"
)

func (s *PostService) GetAllPosts(userID string) ([]models.Post, *models.ErrorJson) {
	return s.repo.GetAllPosts(userID)
}

func (s *PostService) GetPostByID(postID string) (*models.Post, *models.ErrorJson) {
	return s.repo.GetPostByID(postID)
}
