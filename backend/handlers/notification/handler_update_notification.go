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

type UpdateHandler struct {
	NSU *NS.NotificationServiceUpdate
	Us  *US.AuthService
}

func NewUpdateHandler(nsu *NS.NotificationServiceUpdate) *UpdateHandler {
	return &UpdateHandler{NSU: nsu}
}

func (NUH *UpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("requested update path:", r.URL.Path)
	fmt.Println("method", r.Method)

	if r.Method != http.MethodPost {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "ERROR!! Method Not Allowed!"})
		return
	}
	NUH.UpdateNotification(w, r)
}

func (NUH *UpdateHandler) UpdateNotification(w http.ResponseWriter, r *http.Request) {
	user_Id, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: err.Error()})
		return
	}

	var Data models.Unotif

	err = json.NewDecoder(r.Body).Decode(&Data)
	if err != nil {
		fmt.Println("invalide decode lol", Data)

		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "bad request", Message: fmt.Sprintf("%v", err)})
		return
	}

	if errJson := NUH.NSU.UpdateService(Data, user_Id.String()); errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Message: errJson.Message})
		return
	}
	utils.WriteDataBack(w, nil)
}
