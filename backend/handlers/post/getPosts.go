package handlers

import (
	"net/http"

	"social-network/backend/middleware"
	"social-network/backend/models"
	"social-network/backend/utils"
)

func (h *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	usID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: "Failed to get userID"})
		return
	}
	posts, errPosts := h.service.GetAllPosts(usID.String())
	if errPosts != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errPosts.Status, Error: errPosts.Error})
		return
	}
	utils.WriteDataBack(w, posts)
}

func (h *PostHandler) GetPostByID(w http.ResponseWriter, r *http.Request, postID string) {
	post, err := h.service.GetPostByID(postID)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 404, Error: "Post not found"})
		return
	}
	utils.WriteDataBack(w, post)
}
