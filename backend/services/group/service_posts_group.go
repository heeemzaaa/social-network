package group

import (
	"strings"

	"social-network/backend/models"
	"social-network/backend/utils"
)

func (s *GroupService) AddPost(post *models.PostGroup) (*models.PostGroup, *models.ErrorJson) {
	if errJson := s.gRepo.GetGroupById(post.GroupId); errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	// always check the membership and also the the group is a valid one
	if errMembership := s.CheckMembership(post.GroupId, post.User.Id); errMembership != nil {
		return nil, &models.ErrorJson{Status: errMembership.Status, Error: errMembership.Error, Message: errMembership.Message}
	}
	errPostGroup := models.PostGroupErr{}
	if strings.TrimSpace(post.Content) == "" && strings.TrimSpace(post.ImagePath) == "" {
		errPostGroup.Content = "empty Content field!"
	}

	if errPostGroup != (models.PostGroupErr{}) {
		if post.ImagePath != "" {
			if errRemoveImg := utils.RemoveImage(post.ImagePath); errRemoveImg != nil {
				return nil, &models.ErrorJson{Status: 500, Error: errRemoveImg.Error()}
			}
		}
		return nil, &models.ErrorJson{Status: 400, Message: errPostGroup}
	}

	post_created, err := s.gRepo.CreatePost(post)
	if err != nil {
		if post.ImagePath != "" {
			errPostGroup.Content = "empty Content field!"
			if errRemoveImg := utils.RemoveImage(post.ImagePath); errRemoveImg != nil {
				return nil, &models.ErrorJson{Status: 500, Error: errRemoveImg.Error()}
			}
		}
		return nil, &models.ErrorJson{Status: err.Status, Error: err.Error, Message: err.Message}
	}
	return post_created, nil
}

func (s *GroupService) GetPosts(userId, groupId, offset string) ([]models.PostGroup, *models.ErrorJson) {
	if errJson := s.gRepo.GetGroupById(groupId); errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	// always check the membership and also the the group is a valid one
	if errMembership := s.CheckMembership(groupId, userId); errMembership != nil {
		return nil, &models.ErrorJson{Status: errMembership.Status, Error: errMembership.Error, Message: errMembership.Message}
	}
	posts, err := s.gRepo.GetPosts(userId, groupId, offset)
	if err != nil {
		return nil, &models.ErrorJson{Status: err.Status, Message: err.Message, Error: err.Error}
	}
	return posts, nil
}
