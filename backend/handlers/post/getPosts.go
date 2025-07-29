package handlers

import (
	"fmt"
	"net/http"

	"social-network/backend/middleware"
	"social-network/backend/models"
	"social-network/backend/utils"
)

func (h *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	usID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)})
		return
	}
	posts, errPosts := h.service.GetAllPosts(usID.String())
	if errPosts != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errPosts.Status, Error: errPosts.Error})
		return
	}
	utils.WriteDataBack(w, posts)
}
