package post

import (
	"social-network/backend/models"
	pr "social-network/backend/repositories/post"
)



type PostService struct {
	repo *pr.PostsRepository
}

func NewPostService(repo *pr.PostsRepository) *PostService {
	return &PostService{repo: repo}
}

func (ps *PostService) GetComments(postID string) ([]models.Comment, *models.ErrorJson) {
	comments, errComments := ps.repo.GetComments(postID)
	if errComments != nil {
		return []models.Comment{}, &models.ErrorJson{Status: errComments.Status, Error: errComments.Error}
	}
	return comments, nil
}
