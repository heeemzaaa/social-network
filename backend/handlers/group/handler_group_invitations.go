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

type GroupInvitationHandler struct {
	gService *gservice.GroupService
}

func NewGroupInvitationHandler(service *gservice.GroupService) *GroupInvitationHandler {
	return &GroupInvitationHandler{gService: service}
}

func (invHanlder *GroupInvitationHandler) InviteToJoin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HEREEEE inside the invite join!!!!")
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
	usersToInvite := []models.User{}
	fmt.Println("users to invite before", usersToInvite)
	if err := json.NewDecoder(r.Body).Decode(&usersToInvite); err != nil {
		if err == io.EOF {
			utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: models.UserErr{
				UserId: "empty user_id field",
			}})
			return
		}
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: fmt.Sprintf("%v", err)})
		return

	}

	fmt.Println("users to invite after", usersToInvite)

	if errJson := invHanlder.gService.InviteToJoin(userID.String(), groupID.String(), usersToInvite); errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message})
		return
	}
}

func (invHanlder *GroupInvitationHandler) CancelTheInvitation(w http.ResponseWriter, r *http.Request) {
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

	var invitedUser *models.User
	if err := json.NewDecoder(r.Body).Decode(invitedUser); err != nil {
		if err == io.EOF {
			utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: models.UserErr{
				UserId: "empty user_id field",
			}})
			return
		}
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: fmt.Sprintf("%v", err)})
		return

	}
	if errJson := invHanlder.gService.CancelTheInvitation(userID.String(), groupID.String(), invitedUser); errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message})
		return
	}

	// delete the notification from the database
}

func (invHanlder *GroupInvitationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodDelete:
		invHanlder.CancelTheInvitation(w, r)
		return
	case http.MethodPost:
		invHanlder.InviteToJoin(w, r)
		return
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "ERROR!! Method not allowed!"})
		return
	}
}
