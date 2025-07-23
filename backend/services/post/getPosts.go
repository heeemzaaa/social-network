package services

import (
	"social-network/backend/models"

	"github.com/google/uuid"
)

func (s *PostService) GetAllPosts(userID uuid.UUID) ([]models.Post, *models.ErrorJson) {
	return s.repo.GetAllPosts(userID)
}

func (s *PostService) GetPostByID(postID string) (*models.Post, *models.ErrorJson) {
	return s.repo.GetPostByID(postID)
}
