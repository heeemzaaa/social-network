package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"social-network/backend/middleware"
	"social-network/backend/models"
	"social-network/backend/utils"
)

func (h *PostHandler) CommentPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
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
		// fmt.Println("error unmarshaling data:", err)
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Invalid JSON in data field"})
		return
	}

	comment.Img = path
	comment.CreatedAt = time.Time{}

	if comment.Content == "" || comment.PostId == "" {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "Invalid data !"})
		// http.Error(w, "Missing content or post ID", http.StatusBadRequest)
		return
	}

	uuiUserID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 401, Error: "Unauthorized !"})
		// http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userID := uuiUserID.String()

	commID, fullName, errComment := h.service.CreateComment(userID, comment.PostId, comment.Content, comment.Img)
	if errComment != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errComment.Status, Error: errComment.Error})
		// http.Error(w, "Failed to add comment", http.StatusInternalServerError)
		return
	}
	comment.Id = commID
	comment.User.Nickname = fullName
	// w.Header().Set("Content-Type", "application/json")
	// fmt.Println(comment, "this is the comment comes form the back end ....")
	// json.NewEncoder(w).Encode(comment)
	utils.WriteDataBack(w, comment)
}
