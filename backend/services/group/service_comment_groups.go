package group

import (
	"fmt"
	"strings"

	"social-network/backend/models"
)

// add the offset and the limit thing after
func (s *GroupService) GetComments(user_id, postId, offset int) ([]models.CommentGroup, *models.ErrorJson) {
	comments, err := s.gRepo.GetComments(user_id, postId, offset)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return comments, nil
}

// check if the content is null
func (s *GroupService) AddComment(comment *models.CommentGroup) (*models.CommentGroup, *models.ErrorJson) {
	message := models.CommentGroupErr{}
	if strings.TrimSpace(comment.Content) == "" {
		message.Content = "empty body comment!"
	}
	if comment.PostId == "" {
		message.PostId = "Post ID is incorrect or did you mean post_id?"
	}
	if message.Content != "" || message.PostId != "" {
		return nil, &models.ErrorJson{Status: 400, Message: message}
	}
	comment_created, err := s.gRepo.CreateComment(comment)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return comment_created, nil
}
