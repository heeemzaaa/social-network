package services

import (
	"social-network/backend/models"
)

func (ps *PostService) GetComments(postID string) ([]models.Comment, *models.ErrorJson) {
	if postID == "" {
		return []models.Comment{}, &models.ErrorJson{Status: 400, Error: "Invalid data !"}
	}

	comments, errComments := ps.repo.GetComments(postID)
	if errComments != nil {
		return []models.Comment{}, &models.ErrorJson{Status: errComments.Status, Error: errComments.Error}
	}
	return comments, nil
}
