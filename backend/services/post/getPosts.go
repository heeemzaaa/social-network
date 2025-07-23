package services

import (
	"social-network/backend/models"

	"github.com/google/uuid"
)

func (s *PostService) GetAllPosts(userID uuid.UUID) ([]models.Post, error) {
	return s.repo.GetAllPosts(userID)
}

func (s *PostService) GetPostByID(postID string) (models.Post, error) {
	return s.repo.GetPostByID(postID)
}
