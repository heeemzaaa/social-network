package group

import (
	"strings"

	"social-network/backend/models"
)

// add the offset and the limit thing after
func (s *GroupService) GetComments(groupId, userId, postId string, offset int) ([]models.CommentGroup, *models.ErrorJson) {
	if errJson := s.gRepo.GetGroupById(groupId); errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	// always check the membership and also the the group is a valid one
	if errMembership := s.CheckMembership(groupId, userId); errMembership != nil {
		return nil, &models.ErrorJson{Status: errMembership.Status, Error: errMembership.Error, Message: errMembership.Message}
	}
	comments, errJson := s.gRepo.GetComments(userId, postId, offset)
	if errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message}
	}
	return comments, nil
}

// check if the content is null
func (s *GroupService) AddComment(comment *models.CommentGroup) (*models.CommentGroup, *models.ErrorJson) {
	if errJson := s.gRepo.GetGroupById(comment.GroupId); errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	// always check the membership and also the the group is a valid one
	if errMembership := s.CheckMembership(comment.GroupId, comment.UserId); errMembership != nil {
		return nil, &models.ErrorJson{Status: errMembership.Status, Error: errMembership.Error, Message: errMembership.Message}
	}
	message := models.CommentGroupErr{}
	if strings.TrimSpace(comment.Content) == "" {
		message.Content = "empty body comment!"
	}

	if message.Content != "" {
		return nil, &models.ErrorJson{Status: 400, Message: message}
	}

	comment_created, errjson := s.gRepo.CreateComment(comment)
	if errjson != nil {
		return nil, &models.ErrorJson{Status: errjson.Status, Error: errjson.Error, Message: errjson.Message}
	}
	return comment_created, nil
}
