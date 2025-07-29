package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"social-network/backend/middleware"
	"social-network/backend/models"
	"social-network/backend/utils"
)

func (h *PostHandler) CommentPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "Method not allowed !"})
		return
	}

	path, errUploadImg := utils.HanldeUploadImage(r, "img", "posts/commentImg")
	if errUploadImg != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errUploadImg.Status, Message: errUploadImg.Message})
		return
	}

	dataStr := r.FormValue("data")

	var comment models.Comment
	if err := json.Unmarshal([]byte(dataStr), &comment); err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Invalid JSON in data field"})
		return
	}

	comment.Img = path
	comment.CreatedAt = time.Time{}

	if comment.PostId == "" {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "Invalid data !"})
		return
	}

	uuiUserID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)})
		return
	}
	userID := uuiUserID.String()

	comment_created, errComment := h.service.CreateComment(userID, comment.PostId, comment.Content, comment.Img)
	if errComment != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errComment.Status, Error: errComment.Error})
		return
	}

	utils.WriteDataBack(w, comment_created)
}
