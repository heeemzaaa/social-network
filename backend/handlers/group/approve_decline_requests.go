package group

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"social-network/backend/middleware"
	"social-network/backend/models"
	gservice "social-network/backend/services/group"
	"social-network/backend/utils"
)

type ApproveDeclineReqHandler struct {
	gService *gservice.GroupService
}

func NewApproveDeclineReqHandler(service *gservice.GroupService) *ApproveDeclineReqHandler {
	return &ApproveDeclineReqHandler{gService: service}
}

func (decApprHandler *ApproveDeclineReqHandler) Approve(w http.ResponseWriter, r *http.Request) {
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
	userToBeAdded := &models.User{}

	if err := json.NewDecoder(r.Body).Decode(userToBeAdded); err != nil {
		if err == io.EOF {
			utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: models.UserErr{
				UserId: "empty user_id field",
			}})
			return
		}
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: fmt.Sprintf("%v", err)})
		return
	}

	if errJson := decApprHandler.gService.Approve(userID.String(), groupID.String(), userToBeAdded); errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message})
		return
	}
}

func (decApprHandler *ApproveDeclineReqHandler) Decline(w http.ResponseWriter, r *http.Request) {
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
	userToBeRejected := &models.User{}

	if err := json.NewDecoder(r.Body).Decode(userToBeRejected); err != nil {
		if err == io.EOF {
			utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: models.UserErr{
				UserId: "empty user_id field",
			}})
			return
		}
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: fmt.Sprintf("%v", err)})
		return
	}

	if errJson := decApprHandler.gService.Decline(userID.String(), groupID.String(), userToBeRejected); errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message})
		return
	}
}

func (decApprHandler *ApproveDeclineReqHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodPost:
		decApprHandler.Approve(w, r)
		return
	case http.MethodDelete:
		decApprHandler.Decline(w, r)
		return
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "ERROR!! Method Not Allowed!"})
		return

	}
}
