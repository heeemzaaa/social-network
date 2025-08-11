package post

import (
	"social-network/backend/models"
)

func (s *PostService) GetComments(postID string) ([]models.Comment, *models.ErrorJson) {
	exists, err := s.repo.PostExist(postID)
	if err != nil {
		return []models.Comment{}, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	if !exists {
		return []models.Comment{}, &models.ErrorJson{Status: 400, Error: "This post ID doesn't exist !"}
	}

	comments, errComments := s.repo.GetComments(postID)
	if errComments != nil {
		return []models.Comment{}, &models.ErrorJson{Status: errComments.Status, Error: errComments.Error}
	}
	return comments, nil
}
