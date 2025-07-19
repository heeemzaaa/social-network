package handlers

import (
	"net/http"
	"strings"

	"social-network/backend/models"
	ps "social-network/backend/services/post"
	"social-network/backend/utils"
)

type PostHandler struct {
	service *ps.PostService
}

func NewPostHandler(service *ps.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	if len(pathParts) < 2 || pathParts[0] != "api" || pathParts[1] != "posts" {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "Not found"})
		return
	}

	switch r.Method {
	case http.MethodGet:
		if len(pathParts) == 2 {
			h.GetAllPosts(w, r)
			return
		}
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "Not found"})

	case http.MethodPost:
		if len(pathParts) == 2 {
			h.CreatePost(w, r)
			return
		}
		if len(pathParts) == 4 && pathParts[2] == "like" {
			h.LikePost(w, r)
			return
		}
		if pathParts[2] == "comment" {

			h.CommentPost(w, r)
			return
		}

		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "Method not allowed on this endpoint"})

	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "Method not allowed"})
	}
}
