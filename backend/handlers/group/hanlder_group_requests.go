package group

import (
	"encoding/json"
	"fmt"
	"net/http"

	"social-network/backend/middleware"
	"social-network/backend/models"
	gservice "social-network/backend/services/group"
	"social-network/backend/utils"
)

type GroupRequestsHandler struct {
	gService *gservice.GroupService
}

func NewGroupRequestsHandler(service *gservice.GroupService) *GroupRequestsHandler {
	return &GroupRequestsHandler{gService: service}
}

// POST   /groups/{group_id}/join-request  (the userID here is gotten from the context the one who
//
//	is sending the request and the one who will be processing it (the receiver_id) is the admin of the group
//
// DELETE  /groups/{group_id}/join-request  (the same here)
// So in this case we will be needing to check if the user is a member or not
// here the request is not processed yet so he's allowed to cancel it
// GET /groups/{group_id}/join-request
func (GrpReqHandler *GroupRequestsHandler) RequestToJoin(w http.ResponseWriter, r *http.Request) {
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
	if errJson := GrpReqHandler.gService.RequestToJoin(userID.String(), groupID.String()); errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message})
		return
	}

	// 	we need to call the function of nortification of mellagui
   // {sender_id , receiver_id, "group-join"}
	
}

func (GrpReqHandler *GroupRequestsHandler) RequestToCancel(w http.ResponseWriter, r *http.Request) {
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

	if errJson := GrpReqHandler.gService.RequestToCancel(userID.String(), groupID.String()); errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message})
		return
	}

	// delete notification  
	// {sender_id , receiver_id, "group-join"}
	// 
}

func (GrpReqHandler *GroupRequestsHandler) GetRequests(w http.ResponseWriter, r *http.Request) {
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
	users, errJson := GrpReqHandler.gService.GetRequests(userID.String(), groupID.String())
	if errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message})
		return
	}

	if err := json.NewEncoder(w).Encode(&users); err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)})
		return
	}
}

func (GrpReqHandler *GroupRequestsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodPost:
		GrpReqHandler.RequestToJoin(w, r)
		return
	case http.MethodDelete:
		GrpReqHandler.RequestToCancel(w, r)
		return
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "ERROR!! Method Not Allowed!"})
		return

	}
}
