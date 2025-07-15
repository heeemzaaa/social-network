package profile

import (
	"fmt"
	"net/http"

	"social-network/backend/middleware"
	"social-network/backend/models"
	ps "social-network/backend/services/profile"
	"social-network/backend/utils"
)

type ProfileHandler struct {
	service *ps.ProfileService
}

func NewProfileHandler(service *ps.ProfileService) *ProfileHandler {
	return &ProfileHandler{service: service}
}

// GET api/profile/id/info
func (PrHandler *ProfileHandler) GetProfileData(w http.ResponseWriter, r *http.Request, profileID string) {
	authSessionID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: err.Error()})
		return
	}

	profile, errService := PrHandler.service.GetProfileData(profileID, authSessionID.String())
	if errService != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errService.Status, Message: errService.Message})
		return
	}

	utils.WriteDataBack(w, profile)
}

func (PrHandler *ProfileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("path:" , r.URL.Path)
	if r.Method != http.MethodGet {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "Method not allowed !"})
		return
	}

	profileID, path, err := GetPath(r)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: fmt.Sprintf("%v", err)})
		return
	}

	fmt.Println("path: ", path)

	switch path {
	case "":
		PrHandler.GetProfileData(w, r, profileID)
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "Page not found !"})
	}
}
