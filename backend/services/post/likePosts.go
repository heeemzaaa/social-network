package services

import (
	"social-network/backend/models"

	"github.com/google/uuid"
)

func (s *PostService) HandleLike(postID uuid.UUID, userID uuid.UUID) (bool, int, *models.ErrorJson) {
	return s.repo.HandleLike(postID, userID)
}
