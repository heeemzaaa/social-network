package profile

import (
	"encoding/json"
	"net/http"

	"social-network/backend/middleware"
	"social-network/backend/models"
	ps "social-network/backend/services/profile"
	ns "social-network/backend/services/notification"
	"social-network/backend/utils"
)

type FollowActionHandler struct {
	service *ps.ProfileService
	NS *ns.NotificationService
}

func NewFollowActionHandler(service *ps.ProfileService, NS *ns.NotificationService) *FollowActionHandler {
	return &FollowActionHandler{
		service: service,
		NS: NS,
	}
}

// POST api/profile/id/actions/follow
func (fa *FollowActionHandler) Follow(w http.ResponseWriter, r *http.Request) {
	authSessionID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: err.Error()})
		return
	}

	type RequestBody struct {
		ProfileID string `json:"profile_id"`
	}

	var request RequestBody
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Invalid data !"})
		return
	}

	errFollow := fa.service.Follow(request.ProfileID, authSessionID.String(), fa.NS)
	if errFollow != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errFollow.Status, Message: errFollow.Message})
		return
	}

	utils.WriteDataBack(w, "done !")
}

// POST api/profile/id/actions/unfollow
func (fa *FollowActionHandler) Unfollow(w http.ResponseWriter, r *http.Request) {
	authSessionID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: err.Error()})
		return
	}

	type RequestBody struct {
		ProfileID string `json:"profile_id"`
	}

	var request RequestBody
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Invalid data !"})
		return
	}

	errUnfollow := fa.service.Unfollow(request.ProfileID, authSessionID.String())
	if errUnfollow != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errUnfollow.Status, Message: errUnfollow.Message})
		return
	}

	utils.WriteDataBack(w, "Done!")
}

func (fa *FollowActionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "Method not allowed !"})
		return
	}

	_, path, err := GetPath(r)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Invalid path !"})
		return
	}

	switch path {
	case "follow":
		fa.Follow(w, r)
	case "unfollow":
		fa.Unfollow(w, r)
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "Page not found !"})
	}
}
