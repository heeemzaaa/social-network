package profile

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network/backend/middleware"
	"social-network/backend/models"
	"social-network/backend/utils"
	ps "social-network/backend/services/profile"
)

type EditProfileHandler struct {
	service *ps.ProfileService
}

func NewEditProfileHandler(service *ps.ProfileService) *EditProfileHandler {
	return &EditProfileHandler{
		service: service,
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

	utils.WriteDataBack(w, profile)
}

func (ep *EditProfileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPatch {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "Method not allowed !"})
		return
	}

	_, path, err := GetPath(r)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: fmt.Sprintf("%v", err)})
		return
	}

	switch path {
	case "update-privacy":
		ep.UpdatePrivacy(w, r)

	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "Page not found !"})
	}
}
