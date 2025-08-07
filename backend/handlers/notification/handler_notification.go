package notification

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network/backend/middleware"
	"social-network/backend/models"
	NS "social-network/backend/services/notification"
	"social-network/backend/utils"
)

type NotificationHandler struct {
	NS *NS.NotificationService
}

func NewNotificationHandler( ns *NS.NotificationService) *NotificationHandler {
	return &NotificationHandler{NS: ns}
}

func (NH *NotificationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")


	if r.Method != "GET" {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "ERROR!! Method Not Allowed!"})
		return
	}
	NH.GetNotifications(w, r)
}

func (NH *NotificationHandler) GetNotifications(w http.ResponseWriter, r *http.Request) {
	user_Id, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: err.Error()})
		return
	}

	queryParam := r.URL.Query().Get("Count")
	if queryParam == "" {
		hasSeen, errJson := NH.NS.IsHasSeenFalse(user_Id.String())
		if errJson != nil {
			utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Message: errJson.Message})
			return

		}
		

		data := models.HasSeen{
			Status: hasSeen,
			Message: "has new notifications",
		}
		utils.WriteDataBack(w, data)
		return
	}

	if !utils.IsValidQueryParam(queryParam) {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Incorrect QueryParam by field!!"})
		return
	}

	notifications, errJson := NH.NS.GetService(user_Id.String(), queryParam)
	if errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Message: errJson.Message})
		return
	}
	errJson = NH.NS.ToggleAllSeenFalse(notifications)
	if errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)})
		return 
	}

	if err = json.NewEncoder(w).Encode(notifications); err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)})
		return
	}

}

