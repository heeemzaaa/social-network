package group

import (
	"net/http"

	"social-network/backend/models"
	gservice "social-network/backend/services/group"
	"social-network/backend/utils"
)

/***  /api/groups/{group_id}/posts/{post_id}/comments  Route to work with ***/

type GroupCommentHandler struct {
	gservice *gservice.GroupService
}

func NewGroupCommentHandler(service *gservice.GroupService) *GroupCommentHandler {
	return &GroupCommentHandler{gservice: service}
}

func (gCommentHandler *GroupCommentHandler) AddGroupComment(w http.ResponseWriter, r *http.Request) {
	
}

func (gCommentHandler *GroupCommentHandler) GetGroupComments(w http.ResponseWriter, r *http.Request) {
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
