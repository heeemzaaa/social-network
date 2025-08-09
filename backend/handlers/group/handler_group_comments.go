package group

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	"social-network/backend/middleware"
	"social-network/backend/models"
	gservice "social-network/backend/services/group"
	"social-network/backend/utils"
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
	userID, errParse := middleware.GetUserIDFromContext(r.Context())
	if errParse != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: "Incorrect type of userID value!"})
		return
	}

	groupID, err := utils.GetUUIDFromPath(r, "group_id")
	if err != nil {
		fmt.Println("1111111111111111111111111111111111111111:", err)
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: ""})
		return
	}

	postID, err := utils.GetUUIDFromPath(r, "post_id")
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "ERROR!! Incorrect UUID Format!"})
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
	path, errUploadImg := utils.HanldeUploadImage(r, "comment", filepath.Join("groups", "comments"))
	if errUploadImg != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errUploadImg.Status, Message: errUploadImg.Message})
		return
	}
	comment.User.Id, comment.GroupId, comment.PostId, comment.ImagePath = userID.String(), groupID.String(), postID.String(), path

	// even if the userid is given wrong we insert the correct one
	commentCreated, err_ := gCommentHandler.gService.AddComment(comment)
	if err_ != nil {
		utils.WriteJsonErrors(w, *err_)
		return
	}
	fmt.Println("commentCreated:   ", commentCreated)
	utils.WriteDataBack(w, commentCreated)
}

func (gCommentHandler *GroupCommentHandler) GetGroupComments(w http.ResponseWriter, r *http.Request) {
	userID, errParse := middleware.GetUserIDFromContext(r.Context())
	if errParse != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Incorrect type of userID value!"})
		return
	}
	groupID, err := utils.GetUUIDFromPath(r, "group_id")
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: ""})
	}
	
	postID, err := utils.GetUUIDFromPath(r, "post_id")
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "ERROR!! Incorrect UUID Format!"})
		return
	}

	offset := r.URL.Query().Get("offset")
	if offset != "0" {
		if errUUID := utils.IsValidUUID(offset); errUUID != nil {
			utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: fmt.Sprintf("%v", errUUID)})
			return
		}
	}

	comments, errJson := gCommentHandler.gService.GetComments(groupID.String(), userID.String(), postID.String(), offset)
	if errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message})
		return
	}

	if err := json.NewEncoder(w).Encode(comments); err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)})
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
