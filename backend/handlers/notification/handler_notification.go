package notification

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network/backend/middleware"
	"social-network/backend/models"
	US "social-network/backend/services/auth"
	NS "social-network/backend/services/notification"
	"social-network/backend/utils"
)

type NotificationHandler struct {
	NS *NS.NotificationService
	Us *US.AuthService
}

func NewNotificationHandler(ns *NS.NotificationService) *NotificationHandler {
	return &NotificationHandler{NS: ns}
}

func (NH *NotificationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("requested main path:", r.URL.Path)
	fmt.Println("method", r.Method)

	switch r.Method {
	case http.MethodGet:
		NH.GetNotifications(w, r)
		return
	case http.MethodPost:
		NH.CreateNotification(w, r) // web socket here //
		return
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "ERROR!! Method Not Allowed!"})
		return
	}
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
		fmt.Println("has new Notification====" , hasSeen)

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
	errJson = NH.NS.ToggleSeenFalse(notifications)
	if errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)})
		return 
	}

	if err = json.NewEncoder(w).Encode(notifications); err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)})
		return
	}

}

func (NH *NotificationHandler) CreateNotification(w http.ResponseWriter, r *http.Request) {
	user_Id, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: err.Error()})
		return
	}

	var Data models.Notif

	err = json.NewDecoder(r.Body).Decode(&Data)
	if err != nil {
		fmt.Println("invalide decode lol", Data)
		
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "bad request", Message: fmt.Sprintf("%v", err)})
		return
	}
	
	errJson := NH.NS.PostService(Data, user_Id.String())
	if errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Message: errJson.Message})
		return
	}

	utils.WriteDataBack(w, nil)
}

