package notification

import (
	"fmt"
	"net/http"
	"social-network/backend/middleware"
	"social-network/backend/models"
	"social-network/backend/services/notification"
	"social-network/backend/utils"
)

type NotificationHandler struct {
	NS *notification.NotificationService
}

func NewNotificationHandler(ns *notification.NotificationService) *NotificationHandler {
	return &NotificationHandler{NS: ns}
}

func (HN *NotificationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "405 - Method Not Allowed", Message: "ERROR!! Method Not Allowed!"})
		return
	}
	HN.GetNotifications(w, r)
}

func (HN *NotificationHandler) GetNotifications(w http.ResponseWriter, r *http.Request) {
	user_Id, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: err.Error()})
		return
	}

	queryParam := r.URL.Query().Get("Id")

	if queryParam == "" {
		hasSeen, errJson := HN.NS.IsHasSeenFalse(user_Id.String())
		if errJson != nil {
			utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error})
			return
		}

		data := models.ResponseMsg{
			Status:  hasSeen,
			Message: fmt.Sprintf("has new notifications: %v", hasSeen),
		}
		utils.WriteDataBack(w, data)
		return
	}

	notifications, errJson := HN.NS.GetService(user_Id.String(), queryParam)
	if errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error})
		return
	}

	if errJson = HN.NS.ToggleAllSeenFalse(notifications); errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err), Error: errJson.Error})
		return
	}

	utils.WriteDataBack(w, notifications)
}
