package services

import "github.com/google/uuid"

func (s *PostService) HandleLike(postID uuid.UUID, userID uuid.UUID) error {
	return s.repo.HandleLike(postID, userID)
}
