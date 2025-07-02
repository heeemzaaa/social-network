package profile

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	h "social-network/backend/handlers"
	"social-network/backend/models"
	ps "social-network/backend/services/profile"
)

type ProfileHandler struct {
	service *ps.ProfileService
}

func NewProfileHandler(service *ps.ProfileService) *ProfileHandler {
	return &ProfileHandler{service: service}
}

// GET api/profile/id
func (PrHandler *ProfileHandler) GetProfileData(w http.ResponseWriter, r *http.Request, profileID string) {
	authSessionID, err := GetSessionID(r)
	if err != nil {
		h.WriteJsonErrors(w, *err)
		return
	}

	profile, errService := PrHandler.service.GetProfileData(profileID, authSessionID)
	if errService != nil {
		h.WriteJsonErrors(w, *errService)
		return
	}

	h.WriteDataBack(w, profile)
}

// GET api/profile/id/followers
func (PrHandler *ProfileHandler) GetFollowers(w http.ResponseWriter, r *http.Request, profileID string) {
	authSessionID, err := GetSessionID(r)
	if err != nil {
		h.WriteJsonErrors(w, *err)
		return
	}

	users, errService := PrHandler.service.GetFollowers(profileID, authSessionID)
	if errService != nil {
		h.WriteJsonErrors(w, *errService)
		return
	}

	h.WriteDataBack(w, users)
}

// GET api/profile/id/following
func (PrHandler *ProfileHandler) GetFollowing(w http.ResponseWriter, r *http.Request, profileID string) {
	authSessionID, err := GetSessionID(r)
	if err != nil {
		h.WriteJsonErrors(w, *err)
		return
	}

	users, errService := PrHandler.service.GetFollowing(profileID, authSessionID)
	if errService != nil {
		h.WriteJsonErrors(w, *errService)
		return
	}

	h.WriteDataBack(w, users)
}

// POST api/profile/id/follow
func (PrHandler *ProfileHandler) Follow(w http.ResponseWriter, r *http.Request) {
	authSessionID, errSession := GetSessionID(r)
	if errSession != nil {
		h.WriteJsonErrors(w, *errSession)
		return
	}

	type RequestBody struct {
		ProfileID string `json:"profile_id"`
	}

	var request RequestBody
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Invalid data !"})
		return
	}

	errFollow := PrHandler.service.Follow(request.ProfileID, authSessionID)
	if errFollow != nil {
		h.WriteJsonErrors(w, *errFollow)
		return
	}

	h.WriteDataBack(w, "Done")
}

// POST api/profile/id/accepted
func (PrHandler *ProfileHandler) AcceptedRequest(w http.ResponseWriter, r *http.Request, profileID string) {
	type RequestBody struct {
		Requestor string `json:"requestor_id"`
	}

	var request RequestBody
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Invalid data !"})
		return
	}

	errRequest := PrHandler.service.AcceptedRequest(profileID, request.Requestor)
	if errRequest != nil {
		h.WriteJsonErrors(w, *errRequest)
		return
	}
	h.WriteDataBack(w, "done")
}

// POST api/profile/id/rejected
func (PrHandler *ProfileHandler) RejectedRequest(w http.ResponseWriter, r *http.Request, profileID string) {
	type RequestBody struct {
		Requestor string `json:"requestor_id"`
	}

	var request RequestBody
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Invalid data !"})
		return
	}

	errRequest := PrHandler.service.RejectedRequest(profileID, request.Requestor)
	if errRequest != nil {
		h.WriteJsonErrors(w, *errRequest)
		return
	}
	h.WriteDataBack(w, "done")
}

// POST api/profile/id/unfollow
func (PrHandler *ProfileHandler) Unfollow(w http.ResponseWriter, r *http.Request) {
	authSessionID, errSession := GetSessionID(r)
	if errSession != nil {
		h.WriteJsonErrors(w, *errSession)
		return
	}

	type RequestBody struct {
		ProfileID string `json:"profile_id"`
	}

	var request RequestBody
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Invalid data !"})
		return
	}

	errUnfollow := PrHandler.service.Unfollow(request.ProfileID, authSessionID)
	if errUnfollow != nil {
		h.WriteJsonErrors(w, *errUnfollow)
		return
	}
}

// PATCH api/profile/id/update-privacy
func (PrHandler *ProfileHandler) UpdatePrivacy(w http.ResponseWriter, r *http.Request, profileID string) {
	authSessionID, errSession := GetSessionID(r)
	if errSession != nil {
		h.WriteJsonErrors(w, *errSession)
		return
	}

	type RequestBody struct {
		ProfileID string `json:"profile_id"`
		ToStatus  string `json:"to_status"`
	}

	var request RequestBody
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Invalid data !"})
		return
	}

	errUpdate := PrHandler.service.UpdatePrivacy(request.ProfileID, authSessionID, request.ToStatus)
	if errUpdate != nil {
		h.WriteJsonErrors(w, *errUpdate)
		return
	}

	h.WriteDataBack(w, "done !")
}

// PATCH api/profile/id/edit
func (PrHandler *ProfileHandler) UpdateProfileData(w http.ResponseWriter, r *http.Request, profileID string) {
	fmt.Println("UpdateProfileData")
}

func (PrHandler *ProfileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	splittedPath := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	fmt.Println(splittedPath)

	if len(splittedPath) < 3 {
		h.WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "Path not found"})
		return
	}

	profileID := splittedPath[2]
	request := ""
	if len(splittedPath) > 3 {
		request = splittedPath[3]
	}
	switch r.Method {

	case http.MethodGet:
		switch request {
		case "":
			PrHandler.GetProfileData(w, r, profileID)
		case "followers":
			PrHandler.GetFollowers(w, r, profileID)
		case "following":
			PrHandler.GetFollowing(w, r, profileID)
		default:
			h.WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "Page not found !"})
		}
	case http.MethodPatch:
		PrHandler.UpdateProfileData(w, r, profileID)
	case http.MethodPost:
		switch request {
		case "follow":
			PrHandler.Follow(w, r)
		case "unfollow":
			PrHandler.Unfollow(w, r)
		case "update-privacy":
			PrHandler.UpdatePrivacy(w, r, profileID)
		default:
			h.WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "Page not found !"})
		}
	default:
		h.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "Method not allowed"})
	}
}
