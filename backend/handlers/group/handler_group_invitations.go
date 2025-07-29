package group

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"social-network/backend/middleware"
	"social-network/backend/models"
	gservice "social-network/backend/services/group"
	"social-network/backend/services/notification"
	"social-network/backend/utils"
)

type GroupInvitationHandler struct {
	gService *gservice.GroupService
	sNotif   *notification.NotificationService
}

func NewGroupInvitationHandler(service *gservice.GroupService, sNotif *notification.NotificationService) *GroupInvitationHandler {
	return &GroupInvitationHandler{gService: service, sNotif: sNotif}
}

func (invHanlder *GroupInvitationHandler) InviteToJoin(w http.ResponseWriter, r *http.Request) {
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
	userToInvite := models.User{}
	if err := json.NewDecoder(r.Body).Decode(&userToInvite); err != nil {
		if err == io.EOF {
			utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: models.UserErr{
				UserId: "empty user_id field",
			}})
			return
		}
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: fmt.Sprintf("%v", err)})
		return

	}

	newNotif, errJson := invHanlder.gService.InviteToJoin(userID.String(), groupID.String(), userToInvite);
	if errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message})
		return
	}
	fmt.Println("pass NewDataNotification to postService =======>", newNotif)
	
	if errJson := invHanlder.sNotif.PostService(newNotif); errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message})
		return
	}
	utils.WriteDataBack(w, "done")
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

	invitedUser := &models.User{}

	if err := json.NewDecoder(r.Body).Decode(&invitedUser); err != nil {
		if err == io.EOF {
			utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: models.UserErr{
				UserId: "empty user_id field",
			}})
			return
		}
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: fmt.Sprintf("%v", err)})
		return

	}
	fmt.Println("invitedUser", invitedUser)
	if errJson := invHanlder.gService.CancelTheInvitation(userID.String(), groupID.String(), invitedUser); errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message})
		return
	}

	fmt.Println("pass NewDataNotification to postService =======>")
	
	if errJson := invHanlder.sNotif.DeleteService(invitedUser.Id, userID.String(), "group-invitation", groupID.String()); errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message})
		return
	}

	utils.WriteDataBack(w, "done")
	// delete the notification from the database
}

func (invHanlder *GroupInvitationHandler) GetUsersToInvite(w http.ResponseWriter, r *http.Request) {
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

	users, errJson := invHanlder.gService.GetUsersToInvite(userID.String(), groupID.String())
	if errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message})
		return
	}
	if err := json.NewEncoder(w).Encode(users); err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: err.Error()})
		return
	}
}

func (invHanlder *GroupInvitationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("r.Method ==> "+ r.Method)


	switch r.Method {
	case http.MethodDelete:
		invHanlder.CancelTheInvitation(w, r)
		return
	case http.MethodPost:
		invHanlder.InviteToJoin(w, r)
		return
	case http.MethodGet:
		invHanlder.GetUsersToInvite(w, r)
		return
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "ERROR!! Method not allowed!"})
		return
	}
}
