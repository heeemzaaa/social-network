package group

import (
	"fmt"
	"net/http"

	"social-network/backend/models"
	gservice "social-network/backend/services/group"
	"social-network/backend/utils"

	"github.com/google/uuid"
)

/***  /api/groups/{group_id}/events/{event-id}/  ***/
// we'll be needing a kind of toggle like sytem of the reactions
// wttf hadshi bzzzzzf
type GroupEventIDHandler struct {
	gService *gservice.GroupService
}

func NewGroupEventIDHandler(service *gservice.GroupService) *GroupEventIDHandler {
	return &GroupEventIDHandler{gService: service}
}

func (gEventIDHandler *GroupEventIDHandler) AddInterestIntoEvent(w http.ResponseWriter, r *http.Request) {
	userIDVal := r.Context().Value("userID")
	userID, errParse := uuid.Parse(userIDVal.(string))
	if errParse != nil {
		fmt.Println("errors", errParse)
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Incorrect type of userID value!"})
		return
	}

	groupId := r.PathValue("group_id")
	groupID, err := uuid.Parse(groupId)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Incorrect type of groupID value!"})
		return
	}

	eventId := r.PathValue("event_id")
	eventID, err := uuid.Parse(eventId)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Incorrect type of postID value!"})
		return
	}

	fmt.Println("", eventID, userID, groupID)
}

func (gEventIDHandler *GroupEventIDHandler) GetEventDetails(w http.ResponseWriter, r *http.Request) {
	userIDVal := r.Context().Value("userID")
	userID, errParse := uuid.Parse(userIDVal.(string))
	if errParse != nil {
		fmt.Println("errors", errParse)
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Incorrect type of userID value!"})
		return
	}

	groupId := r.PathValue("group_id")
	groupID, err := uuid.Parse(groupId)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Incorrect type of groupID value!"})
		return
	}

	eventId := r.PathValue("event_id")
	eventID, err := uuid.Parse(eventId)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Incorrect type of postID value!"})
		return
	}
	gEventIDHandler.gService.GetEventDetails(eventID.String(), userID.String(), groupID.String())
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
