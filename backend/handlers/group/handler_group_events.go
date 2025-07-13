package group

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"social-network/backend/models"
	gservice "social-network/backend/services/group"
	"social-network/backend/utils"

	"github.com/google/uuid"
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
	groupId := r.PathValue("group_id")
	groupID, err := uuid.Parse(groupId)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Incorrect type of groupID value!"})
		return
	}
	userIDVal := r.Context().Value("userID")
	userID, errParse := uuid.Parse(userIDVal.(string))
	if errParse != nil {
		fmt.Println("errors", errParse)
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Incorrect type of userID value!"})
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
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: fmt.Sprintf("%v", err)})
		return
	}
	event.EventCreatorId, event.GroupId = userID.String(), groupID.String()
	event, errJson := gEventHandler.gService.AddGroupEvent(event)
	if errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Message: errJson.Message})
		return
	}
	utils.WriteDataBack(w, event)
}

// we'll be working with exists to check if a user is member before proceeding in any action!!
func (gEventHandler *GroupEventHandler) GetGroupEvents(w http.ResponseWriter, r *http.Request) {
	userID , errParse := utils.GetUserIDFromContext(r.Context())
	if errParse != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Incorrect type of userID value!"})
		return
	}
	groupId := r.PathValue("group_id")
	offset, errOffset := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 64)
	if errOffset != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Incorrect Offset Format!"})
		return
	}

	events, errJson := gEventHandler.gService.GetGroupEvents(groupId, userID.String(), offset)
	if errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Message: errJson.Message})
		return
	}

	if err := json.NewEncoder(w).Encode(events); err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)})
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
