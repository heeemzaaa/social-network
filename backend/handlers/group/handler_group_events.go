package group

import (
	"fmt"
	"net/http"
	"strconv"

	"social-network/backend/models"
	gservice "social-network/backend/services/group"
	"social-network/backend/utils"

	"github.com/google/uuid"
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
	userIDVal := r.Context().Value("userID")
	userID, errParse := uuid.Parse(userIDVal.(string))
	if errParse != nil {
		fmt.Println("errors", errParse)
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Incorrect type of userID value!"})
		return
	}
	offset, errOffset := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 64)
	if errOffset != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Incorrect Offset Format!"})
		return
	}

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
