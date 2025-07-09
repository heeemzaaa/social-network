package group

import (
	"net/http"

	"social-network/backend/models"
	gservice "social-network/backend/services/group"
	"social-network/backend/utils"
)



/***   /api/groups/{group_id}/events/    ***/
// here we'll be also querying if the user logged in is interested or not in the event!!!



type GroupEventHandler struct {
	gservice *gservice.GroupService
}

func NewGroupEventHandler(service *gservice.GroupService) *GroupEventHandler {
	return &GroupEventHandler{gservice: service}
}

func (gEventHandler *GroupEventHandler) AddGroupEvent(w http.ResponseWriter, r *http.Request) {
}

func (gEventHandler *GroupEventHandler) GetGroupEvents(w http.ResponseWriter, r *http.Request) {
}

func (gEventHandler *GroupEventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		gEventHandler.GetGroupEvents(w, r)
		return
	case http.MethodPost:
		gEventHandler.AddGroupEvent(w, r)
		return
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "ERROR!! Method not allowed!"})
		return
	}
}
