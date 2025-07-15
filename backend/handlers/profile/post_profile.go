package profile

import (
	"fmt"
	"net/http"

	"social-network/backend/middleware"
	"social-network/backend/models"
	ps "social-network/backend/services/profile"
	"social-network/backend/utils"
)

type ProfilePostHandler struct {
	service *ps.ProfileService
}

func NewProfilePostsHandler(service *ps.ProfileService) *ProfilePostHandler {
	return &ProfilePostHandler{service: service}
}

// GET /api/profile/id/data/posts
func (p *ProfilePostHandler) GetPostsOfTheProfile(w http.ResponseWriter, r *http.Request, profileID string) {
	authSessionID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: err.Error()})
		return
	}

	posts, errPosts := p.service.GetPosts(profileID, authSessionID.String())
	if errPosts != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: errPosts.Error})
	}

	utils.WriteDataBack(w, posts)
}

func (p *ProfilePostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "Method not allowed !"})
		return
	}

	profileID, path, err := GetPath(r)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: fmt.Sprintf("%v", err)})
		return
	}

	switch path {
	case "/posts":
		p.GetPostsOfTheProfile(w, r, profileID)
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "Page not found !"})
	}
}
