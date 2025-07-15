package profile

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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

// GET api/profile/id
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

// GET api/profile/id/followers
func (PrHandler *ProfileHandler) GetFollowers(w http.ResponseWriter, r *http.Request, profileID string) {
	fmt.Println("here*********************")
	authSessionID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: err.Error()})
		return
	}

	users, errService := PrHandler.service.GetFollowers(profileID, authSessionID.String())
	if errService != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errService.Status, Message: errService.Message})
		return
	}

	utils.WriteDataBack(w, users)
}

// GET api/profile/id/following
func (PrHandler *ProfileHandler) GetFollowing(w http.ResponseWriter, r *http.Request, profileID string) {
	authSessionID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: err.Error()})
		return
	}

	users, errService := PrHandler.service.GetFollowing(profileID, authSessionID.String())
	if errService != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errService.Status, Message: errService.Message})
		return
	}

	utils.WriteDataBack(w, users)
}

// POST api/profile/id/follow
func (PrHandler *ProfileHandler) Follow(w http.ResponseWriter, r *http.Request) {
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

	errFollow := PrHandler.service.Follow(request.ProfileID, authSessionID.String())
	if errFollow != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errFollow.Status, Message: errFollow.Message})
		return
	}

	utils.WriteDataBack(w, "Done")
}

// POST api/profile/id/accepted
func (PrHandler *ProfileHandler) AcceptedRequest(w http.ResponseWriter, r *http.Request, profileID string) {
	type RequestBody struct {
		Requestor string `json:"requestor_id"`
	}

	var request RequestBody
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Invalid data !"})
		return
	}

	errRequest := PrHandler.service.AcceptedRequest(profileID, request.Requestor)
	if errRequest != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errRequest.Status, Message: errRequest.Message})
		return
	}
	utils.WriteDataBack(w, "done")
}

// POST api/profile/id/rejected
func (PrHandler *ProfileHandler) RejectedRequest(w http.ResponseWriter, r *http.Request, profileID string) {
	type RequestBody struct {
		Requestor string `json:"requestor_id"`
	}

	var request RequestBody
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Invalid data !"})
		return
	}

	errRequest := PrHandler.service.RejectedRequest(profileID, request.Requestor)
	if errRequest != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errRequest.Status, Message: errRequest.Message})
		return
	}
	utils.WriteDataBack(w, "done")
}

// POST api/profile/id/unfollow
func (PrHandler *ProfileHandler) Unfollow(w http.ResponseWriter, r *http.Request) {
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

	errUnfollow := PrHandler.service.Unfollow(request.ProfileID, authSessionID.String())
	if errUnfollow != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errUnfollow.Status, Message: errUnfollow.Message})
		return
	}
}

// PATCH api/profile/id/update-privacy
func (PrHandler *ProfileHandler) UpdatePrivacy(w http.ResponseWriter, r *http.Request) {
	authSessionID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: err.Error()})
		return
	}

	type RequestBody struct {
		ProfileID    string `json:"profile_id"`
		WantedStatus string `json:"wanted_status"`
	}

	var request RequestBody
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Invalid data !"})
		return
	}

	errUpdate := PrHandler.service.UpdatePrivacy(request.ProfileID, authSessionID.String(), request.WantedStatus)
	if errUpdate != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errUpdate.Status, Message: errUpdate.Message})
		return
	}

	utils.WriteDataBack(w, "done !")
}

// PATCH api/profile/id/edit
func (PrHandler *ProfileHandler) UpdateProfileData(w http.ResponseWriter, r *http.Request, profileID string) {
	fmt.Println("UpdateProfileData")
}

// get the posts with all cases , with one query as oumayma said
func (PrHandler *ProfileHandler) GetPostsOfTheProfile(w http.ResponseWriter, r *http.Request, profileID string) {
	authSessionID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: err.Error()})
		return
	}

	posts , errPosts := PrHandler.service.GetPosts(profileID, authSessionID.String())
	if errPosts != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: errPosts.Error})
	}

	utils.WriteDataBack(w, posts)
}

// global handler
func (PrHandler *ProfileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	splittedPath := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(splittedPath) < 3 {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "Path not found"})
		return
	}

	profileID, err := utils.GetUUIDFromPath(r, "id")
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: err.Error()})
		return
	}

	request := ""
	if len(splittedPath) > 3 {
		request = splittedPath[3]
	}

	switch r.Method {

	case http.MethodGet:
		switch request {
		case "":
			PrHandler.GetProfileData(w, r, profileID.String())
		case "followers":
			PrHandler.GetFollowers(w, r, profileID.String())
		case "following":
			PrHandler.GetFollowing(w, r, profileID.String())
		case "get-posts":
			PrHandler.GetPostsOfTheProfile(w, r, profileID.String())
		default:
			utils.WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "Page not found !"})
		}
	case http.MethodPatch:
		switch request {
		case "update-privacy":
			PrHandler.UpdatePrivacy(w, r)
		case "update-profile":
			PrHandler.UpdateProfileData(w, r, profileID.String())
		default:
			utils.WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "Page not found !"})
		}
	case http.MethodPost:
		switch request {
		case "follow":
			PrHandler.Follow(w, r)
		case "unfollow":
			PrHandler.Unfollow(w, r)
		default:
			utils.WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "Page not found !"})
		}
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "Method not allowed"})
	}
}
