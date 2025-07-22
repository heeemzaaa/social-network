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
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("entered handler comment 1")

	path, errUploadImg := utils.HanldeUploadImage(r, "img", "posts/commentImg")
	if errUploadImg != nil {
		fmt.Println(errUploadImg)
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errUploadImg.Status, Message: errUploadImg.Message})
		return
	}

	dataStr := r.FormValue("data")

	var comment models.Comment
	if err := json.Unmarshal([]byte(dataStr), &comment); err != nil {
		fmt.Println("error unmarshaling data:", err)
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Invalid JSON in data field"})
		return
	}

	fmt.Println("entered handler comment 2")

	comment.Img = path
	comment.CreatedAt = time.Time{}

	if comment.Content == "" || comment.PostId == "" {
		http.Error(w, "Missing content or post ID", http.StatusBadRequest)
		return
	}

	uuiUserID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userID := uuiUserID.String()

	commID, fullName, err := h.service.CreateComment(userID, comment.PostId, comment.Content, comment.Img)
	if err != nil {
		http.Error(w, "Failed to add comment", http.StatusInternalServerError)
		return
	}
	comment.Id = commID
	comment.User.Nickname = fullName
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(comment, "this is the comment comes form the back end ....")
	json.NewEncoder(w).Encode(comment)
}
