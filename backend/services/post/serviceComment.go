package services

import (
	"social-network/backend/models"
)

func (s *PostService) CreateComment(userID string, postID string, content string, image_url string) (string, string, *models.ErrorJson) {
	return s.repo.CreateComment(userID, postID, content, image_url)
}
