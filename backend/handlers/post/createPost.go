package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"social-network/backend/middleware"
	"social-network/backend/models"
	"social-network/backend/utils"
)

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	usID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 401, Error: "Unauthorized"})
		return
	}

	path, errUploadImg := utils.HanldeUploadImage(r, "img", "posts")
	if errUploadImg != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errUploadImg.Status, Error: errUploadImg.Error})
		return
	}

	dataStr := r.FormValue("data")
	if dataStr == "" {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "Missing data field"})
		return
	}

	var post models.Post
	if err := json.Unmarshal([]byte(dataStr), &post); err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "Invalid JSON in data field"})
		return
	}

	post.User.Id = usID.String()
	post.Img = path
	post.Id = utils.NewUUID()
	post.CreatedAt = time.Now().Format(time.RFC3339)

	if err := h.service.CreatePost(&post); err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: err.Status, Error: err.Error})
		return
	}
	utils.WriteDataBack(w, post)
}
