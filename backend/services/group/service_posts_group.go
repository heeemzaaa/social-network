package group

import (
	"strings"

	"social-network/backend/models"
)

func (s *GroupService) AddPost(post *models.PostGroup) (*models.PostGroup, *models.ErrorJson) {
	errPostGroup := models.PostGroupErr{}
	if strings.TrimSpace(post.Content) == "" {
		errPostGroup.Content = "empty post content!!"
	}

	if errPostGroup != (models.PostGroupErr{}) {
		return nil, &models.ErrorJson{Status: 400, Message: errPostGroup}
	}

	post_created, err := s.gRepo.CreatePost(post)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error, Message: err.Message}
	}
	return post_created, nil
}

func (s *GroupService) GetPosts(user_id string, offset int) ([]models.PostGroup, *models.ErrorJson) {
	posts, err := s.gRepo.GetPosts(user_id, offset)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Message: err.Message, Error: err.Error}
	}
	return posts, nil
}
