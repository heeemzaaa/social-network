package notification

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network/backend/middleware"
	"social-network/backend/models"

	"social-network/backend/utils"
	ns "social-network/backend/services/notification"
)

type UpdateHandler struct {
	NS *ns.NotificationService
}

func NewUpdateHandler(NS *ns.NotificationService) *UpdateHandler {
	return &UpdateHandler{NS: NS}
}

func (HUN *UpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "ERROR!! Method Not Allowed!", Error: "405 - Method Not Allowed"})
		return
	}
	HUN.UpdateNotification(w, r)
}

func (HUN *UpdateHandler) UpdateNotification(w http.ResponseWriter, r *http.Request) {
	user_Id, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: err.Error()})
		return
	}

	var Data models.Unotif
	if err = json.NewDecoder(r.Body).Decode(&Data); err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "400 - Bad Request", Message: fmt.Sprintf("%v", err)})
		return
	}

	errJson := HUN.NS.UpdateService(Data, user_Id.String());
	if errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Message: errJson.Message})
		return
	}

	data := models.ResponseMsg{
		Status: true,
		Message: "Your action was successful",
	}
	utils.WriteDataBack(w, data)
}
