package group

import (
	"net/http"

	"social-network/backend/models"
	gservice "social-network/backend/services/group"
	"social-network/backend/utils"
)

/***  /api/groups/{group_id}/events/{event-id}/  ***/

type GroupEventIDHandler struct {
	gservice *gservice.GroupService
}

func NewGroupEventIDHandler(service *gservice.GroupService) *GroupEventIDHandler {
	return &GroupEventIDHandler{gservice: service}
}

func (gEventIDHandler *GroupEventIDHandler) AddInterestIntoEvent(w http.ResponseWriter, r *http.Request) {
}

func (gEventIDHandler *GroupEventIDHandler) GetEventDetails(w http.ResponseWriter, r *http.Request) {
}

func (gEventIDHandler *GroupEventIDHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodPost:
		gEventIDHandler.AddInterestIntoEvent(w, r)
		return
	case http.MethodGet:
		gEventIDHandler.GetEventDetails(w, r)
		return
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "ERROR!! Method not allowed!"})
		return
	}
}
