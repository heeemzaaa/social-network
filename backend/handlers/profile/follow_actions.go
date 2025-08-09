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
	authUserID, err := middleware.GetUserIDFromContext(r.Context())
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
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "Invalid data !"})
		return
	}

	data, profile, errFollow := fa.service.Follow(request.ProfileID, authUserID.String())
	if errFollow != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errFollow.Status, Error: errFollow.Error})
		return
	}

	errJson := fa.NS.PostService(data)
	if errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Error: errJson.Error})
		return
	}

	utils.WriteDataBack(w, profile)
}

// POST api/profile/id/actions/unfollow
func (fa *FollowActionHandler) Unfollow(w http.ResponseWriter, r *http.Request) {
	authUserID, err := middleware.GetUserIDFromContext(r.Context())
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
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "Invalid data !"})
		return
	}

	profile, errUnfollow := fa.service.Unfollow(request.ProfileID, authUserID.String())
	if errUnfollow != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errUnfollow.Status, Error: errUnfollow.Error})
		return
	}
	utils.WriteDataBack(w, profile)
}

func (fa *FollowActionHandler) CancelFollow(w http.ResponseWriter, r *http.Request) {
	authUserID, err := middleware.GetUserIDFromContext(r.Context())
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
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "Invalid data !"})
		return
	}

	profile, errCancel := fa.service.CancelFollow(request.ProfileID, authUserID.String())
	if errCancel != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errCancel.Status, Error: errCancel.Error})
		return
	}

	if errJson := fa.NS.DeleteService(request.ProfileID, authUserID.String(), "follow-private", ""); errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Error: errJson.Error})
		return
	}
	

	utils.WriteDataBack(w, profile)
}

// handler the three cases here
func (fa *FollowActionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "Method not allowed !"})
		return
	}

	_, path, err := GetPath(r)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "Invalid path !"})
		return
	}

	switch path {
	case "follow":
		fa.Follow(w, r)
	case "unfollow":
		fa.Unfollow(w, r)
	case "cancel":
		fa.CancelFollow(w, r)
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 404, Error: "Page not found !"})
	}
}
