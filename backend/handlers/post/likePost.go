package handlers

import (
	"net/http"

	"social-network/backend/middleware"
	"social-network/backend/models"
	"social-network/backend/utils"
)

func (h *PostHandler) LikePost(w http.ResponseWriter, r *http.Request) {
	postID, err := utils.GetUUIDFromPath(r, "id")
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: http.StatusBadRequest, Error: "Invalid post ID"})
		return
	}

	userID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: http.StatusUnauthorized, Error: "Unauthorized"})
		return
	}

	liked, totalLikes, errLike := h.service.HandleLike(postID.String(), userID.String())
	if errLike != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errLike.Status, Error: errLike.Error})
		return
	}

	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(map[string]any{
	// 	"success":     true,
	// 	"message":     "Post like updated successfully",
	// 	"postId":      postID,
	// 	"liked":       liked,
	// 	"total_likes": totalLikes,
	// })

	utils.WriteDataBack(w, map[string]any{
		"success":     true,
		"message":     "Post like updated successfully",
		"postId":      postID,
		"liked":       liked,
		"total_likes": totalLikes,
	})
}
