package handlers

import (
	"encoding/json"
	"net/http"

	"social-network/backend/middleware"
	"social-network/backend/models"
	"social-network/backend/utils"
)

func (h *PostHandler) LikePost(w http.ResponseWriter, r *http.Request) {
	postID, err := utils.GetUUIDFromPath(r, "id")
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: http.StatusBadRequest, Message: "Invalid post ID"})
		return
	}

	userID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: http.StatusUnauthorized, Message: "Unauthorized"})
		return
	}

	liked, totalLikes, err := h.service.HandleLike(postID, userID)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: http.StatusInternalServerError, Message: "Failed to like post"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"success":     true,
		"message":     "Post like updated successfully",
		"postId":      postID,
		"liked":       liked,
		"total_likes": totalLikes,
	})
}
