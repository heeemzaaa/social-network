package profile

import (
	"encoding/json"
	"fmt"
	"net/http"

	"social-network/backend/middleware"
	"social-network/backend/models"
	ns "social-network/backend/services/notification"
	ps "social-network/backend/services/profile"
	"social-network/backend/utils"
)

type EditProfileHandler struct {
	service *ps.ProfileService
	NS      *ns.NotificationService
}

func NewEditProfileHandler(service *ps.ProfileService, NS *ns.NotificationService) *EditProfileHandler {
	return &EditProfileHandler{
		service: service,
		NS:      NS,
	}
}

// PATCH api/profile/id/edit/update-privacy
func (ep *EditProfileHandler) UpdatePrivacy(w http.ResponseWriter, r *http.Request) {
	authUserID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: err.Error()})
		return
	}

	type RequestBody struct {
		WantedStatus string `json:"wanted_status"`
	}

	var request RequestBody
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "Invalid data !"})
		return
	}

	profile, errUpdate := ep.service.UpdatePrivacy(authUserID.String(), request.WantedStatus)
	if errUpdate != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errUpdate.Status, Error: errUpdate.Error})
		return
	}

	if profile.User.Visibility == "public" {
		all, errJson := ep.NS.GetAllNotificationByType(authUserID.String(), "follow-private")
		if errJson != nil {
			utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Error: errJson.Error})
			return
		}

		errJson = ep.NS.ToggleAllStaus(all, "accept", "follow-private")
		if errJson != nil {
			utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Error: errJson.Error})
			return
		}
	}

	utils.WriteDataBack(w, profile)
}

func (ep *EditProfileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPatch {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "Method not allowed !"})
		return
	}

	_, path, err := GetPath(r)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: fmt.Sprintf("%v", err)})
		return
	}

	switch path {
	case "update-privacy":
		ep.UpdatePrivacy(w, r)

	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 404, Error: "Page not found !"})
	}
}
