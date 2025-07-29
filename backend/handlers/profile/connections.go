package profile

import (
	"fmt"
	"net/http"

	"social-network/backend/middleware"
	"social-network/backend/models"
	ps "social-network/backend/services/profile"
	"social-network/backend/utils"
)

type UserConnectionHandler struct {
	service *ps.ProfileService
}

func NewUserConnectionHandler(service *ps.ProfileService) *UserConnectionHandler {
	return &UserConnectionHandler{service: service}
}

// GET api/profile/id/connections/followers
func (uc *UserConnectionHandler) GetFollowers(w http.ResponseWriter, r *http.Request, profileID string) {
	authUserID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)})
		return
	}

	users, errService := uc.service.GetFollowers(profileID, authUserID.String())
	if errService != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errService.Status, Message: errService.Message})
		return
	}

	utils.WriteDataBack(w, users)
}

// GET api/profile/id/connections/following
func (uc *UserConnectionHandler) GetFollowing(w http.ResponseWriter, r *http.Request, profileID string) {
	authSessionID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: err.Error()})
		return
	}

	users, errService := uc.service.GetFollowing(profileID, authSessionID.String())
	if errService != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errService.Status, Message: errService.Message})
		return
	}

	utils.WriteDataBack(w, users)
}

func (uc *UserConnectionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	if r.Method != http.MethodGet {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "Method not allowed !"})
		return
	}

	profileID, path, err := GetPath(r)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Invalid path"})
	}

	switch path {
	case "followers":
		uc.GetFollowers(w, r, profileID)
	case "following":
		uc.GetFollowing(w, r, profileID)
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "Page not found !"})
	}
}
