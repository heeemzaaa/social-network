package group

import (
	"fmt"
	"net/http"

	"social-network/backend/middleware"
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
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: "Incorrect type of userID value!"})
		return
	}

	groupID, err := utils.GetUUIDFromPath(r, "group_id")
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "ERROR!! Incorrect UUID Format!"})
		return
	}

	eventID, err := utils.GetUUIDFromPath(r, "event_id")
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: "ERROR!! Incorrect UUID Format!"})
		return
	}

	fmt.Println("", eventID, userID, groupID)
}

func (gEventIDHandler *GroupEventIDHandler) GetEventDetails(w http.ResponseWriter, r *http.Request) {
	userID, errParse := middleware.GetUserIDFromContext(r.Context())
	if errParse != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: errParse.Error()})
		return
	}

	groupID, err := utils.GetUUIDFromPath(r, "group_id")
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "ERROR!! Incorrect UUID Format!"})
		return
	}

	eventID, err := utils.GetUUIDFromPath(r, "event_id")
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "ERROR!! Incorrect UUID Format!"})
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
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "ERROR!! Method not allowed!"})
		return
	}
}
