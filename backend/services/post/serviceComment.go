package services

import (
	"social-network/backend/models"
)

// create the comment service .... 
func (s *PostService) CreateComment(userID string, postID string, content string, image_url string) (*models.Comment, *models.ErrorJson) {
	// only if the post exist 
	exists, err := s.repo.PostExist(postID)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	if !exists {
		return nil, &models.ErrorJson{Status: 400, Error: "This post ID doesn't exist !"}
	}
	
	return s.repo.CreateComment(userID, postID, content, image_url)
}
