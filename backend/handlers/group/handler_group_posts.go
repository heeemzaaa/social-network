package group

import (
	"net/http"

	"social-network/backend/models"
	gservice "social-network/backend/services/group"
	"social-network/backend/utils"
)

/***    /api/groups/{group_id}/posts/    ***/

type GroupPostHandler struct {
	gservice *gservice.GroupService
}

func NewGroupPostHandler(service *gservice.GroupService) *GroupPostHandler {
	return &GroupPostHandler{gservice: service}
}

func (gPostHandler *GroupPostHandler) AddGroupPost(w http.ResponseWriter, r *http.Request) {
}

func (gPostHandler *GroupPostHandler) GetGroupPosts(w http.ResponseWriter, r *http.Request) {
}

func (gPostHandler *GroupPostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		gPostHandler.GetGroupPosts(w, r)
		return
	case http.MethodPost:
		gPostHandler.AddGroupPost(w, r)
		return
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "ERROR!! Method not allowed!"})
		return
	}
}
