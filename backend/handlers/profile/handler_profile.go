package profile

import (
	"fmt"
	"net/http"
	h "social-network/backend/handlers"
	"social-network/backend/models"
	ps "social-network/backend/services/profile"
	"strings"
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

// POST api/profile/id/edit
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
	default:
		h.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "Method not allowed"})
	}
}
