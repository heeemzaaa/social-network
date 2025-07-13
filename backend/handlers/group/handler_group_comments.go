package group

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	"social-network/backend/models"
	gservice "social-network/backend/services/group"
	"social-network/backend/utils"

	"github.com/google/uuid"
)

/***  /api/groups/{group_id}/posts/{post_id}/comments  Route to work with ***/
// not tested yet

type GroupCommentHandler struct {
	gService *gservice.GroupService
}

func NewGroupCommentHandler(service *gservice.GroupService) *GroupCommentHandler {
	return &GroupCommentHandler{gService: service}
}

func (gCommentHandler *GroupCommentHandler) AddGroupComment(w http.ResponseWriter, r *http.Request) {
	var comment *models.CommentGroup
	userID, errParse := utils.GetUserIDFromContext(r.Context())
	if errParse != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Incorrect type of userID value!"})
		return
	}

	groupId := r.PathValue("group_id")
	groupID, err := uuid.Parse(groupId)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Incorrect type of groupID value!"})
		return
	}

	postId := r.PathValue("post_id")
	postID, err := uuid.Parse(postId)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Incorrect type of postID value!"})
		return
	}

	data := r.FormValue("data")
	if err := json.Unmarshal([]byte(data), &comment); err != nil {
		if err == io.EOF {
			utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: &models.PostGroupErr{
				Content: "empty Content field!",
			}})
			return
		}
		// which status code to return
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: fmt.Sprintf("%v", err)})
		return

	}

	// handle the image encoding in the phase that comes before the adding process
	path, errUploadImg := utils.HanldeUploadImage(r, "comment", filepath.Join("groups", "comments"), false)
	if errUploadImg != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errUploadImg.Status, Message: errUploadImg.Message})
		return
	}
	comment.UserId, comment.GroupId, comment.PostId, comment.ImagePath = userID.String(), groupID.String(), postID.String(), path
	// even if the userid is given wrong we insert the correct one
	postCreated, err_ := gCommentHandler.gService.AddComment(comment)
	if err_ != nil {
		utils.WriteJsonErrors(w, *err_)
		return
	}
	utils.WriteDataBack(w, postCreated)
}

func (gCommentHandler *GroupCommentHandler) GetGroupComments(w http.ResponseWriter, r *http.Request) {
	_, errParse := utils.GetUserIDFromContext(r.Context())
	if errParse != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Incorrect type of userID value!"})
		return
	}
}

func (gCommentHandler *GroupCommentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		gCommentHandler.GetGroupComments(w, r)
		return
	case http.MethodPost:
		gCommentHandler.AddGroupComment(w, r)
		return
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "ERROR!! Method not allowed!"})
		return
	}
}
