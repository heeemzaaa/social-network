package group

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"social-network/backend/middleware"
	"social-network/backend/models"
	gservice "social-network/backend/services/group"
	"social-network/backend/utils"
)

/***   /api/groups/{group_id}/events/    ***/
// here we'll be also querying if the user logged in is interested or not in the event!!!
// not tested yet
type GroupEventHandler struct {
	gService *gservice.GroupService
}

func NewGroupEventHandler(service *gservice.GroupService) *GroupEventHandler {
	return &GroupEventHandler{gService: service}
}

func (gEventHandler *GroupEventHandler) AddGroupEvent(w http.ResponseWriter, r *http.Request) {
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

	var event *models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		if err == io.EOF {
			utils.WriteJsonErrors(w, models.ErrorJson{
				Status: 400,
				Message: models.ErrEventGroup{
					Title:       "title field is empty!",
					Description: "description field is empty!",
					EventDate:   "event date is not set up",
				},
			})
			return
		}
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: fmt.Sprintf("%v hhh", err)})
		return
	}
	event.EventCreatorId, event.GroupId = userID.String(), groupID.String()
	event, errJson := gEventHandler.gService.AddGroupEvent(event)
	if errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error})
		return
	}
	utils.WriteDataBack(w, event)
}

// we'll be working with exists to check if a user is member before proceeding in any action!!
func (gEventHandler *GroupEventHandler) GetGroupEvents(w http.ResponseWriter, r *http.Request) {
	userID, errParse := middleware.GetUserIDFromContext(r.Context())
	if errParse != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: "Incorrect type of userID value!"})
		return
	}
	groupID, err := utils.GetUUIDFromPath(r, "group_id")
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "ERROR!! Incorrect UUID Format!"})
		return
	}
	offset, errOffset := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 64)
	if errOffset != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "Incorrect Offset Format!"})
		return
	}

	events, errJson := gEventHandler.gService.GetGroupEvents(groupID.String(), userID.String(), offset)
	if errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error})
		return
	}

	if err := json.NewEncoder(w).Encode(events); err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)})
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
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "ERROR!! Method not allowed!"})
		return
	}
}
