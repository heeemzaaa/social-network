package group

import (
	"strings"

	"social-network/backend/models"
	"social-network/backend/utils"
)

// add the offset and the limit thing after
func (gService *GroupService) GetComments(groupId, userId, postId string, offset int) ([]models.CommentGroup, *models.ErrorJson) {
	if errJson := gService.gRepo.GetGroupById(groupId); errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	// always check the membership and also the the group is a valid one
	if errMembership := gService.CheckMembership(groupId, userId); errMembership != nil {
		return nil, &models.ErrorJson{Status: errMembership.Status, Error: errMembership.Error, Message: errMembership.Message}
	}
	comments, errJson := gService.gRepo.GetComments(userId, postId, groupId, offset)
	if errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message}
	}
	return comments, nil
}

// check if the content is null
func (gService *GroupService) AddComment(comment *models.CommentGroup) (*models.CommentGroup, *models.ErrorJson) {
	if errJson := gService.gRepo.GetGroupById(comment.GroupId); errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	// always check the membership and also the the group is a valid one
	if errMembership := gService.CheckMembership(comment.GroupId, comment.User.Id); errMembership != nil {
		return nil, &models.ErrorJson{Status: errMembership.Status, Error: errMembership.Error, Message: errMembership.Message}
	}

	_, postsExists, _ := gService.gRepo.GetItem("group_posts", "postID", comment.PostId)
	if !postsExists {
		return nil, &models.ErrorJson{Status: 400, Error: "ERROR!! Post not Found!!"}
	}

	message := models.CommentGroupErr{}
	if strings.TrimSpace(comment.Content) == "" {
		message.Content = "empty body comment!"
	}

	if message.Content != "" {
		if comment.ImagePath != "" {
			if errRemoveImg := utils.RemoveImage(comment.ImagePath); errRemoveImg != nil {
				return nil, &models.ErrorJson{Status: 500, Error: errRemoveImg.Error()}
			}
		}
		return nil, &models.ErrorJson{Status: 400, Message: message}
	}

	comment_created, errjson := gService.gRepo.CreateComment(comment)
	if errjson != nil {
		if comment.ImagePath != "" {
			if errRemoveImg := utils.RemoveImage(comment.ImagePath); errRemoveImg != nil {
				return nil, &models.ErrorJson{Status: 500, Error: errRemoveImg.Error()}
			}
		}
		return nil, &models.ErrorJson{Status: errjson.Status, Error: errjson.Error, Message: errjson.Message}
	}
	return comment_created, nil
}
