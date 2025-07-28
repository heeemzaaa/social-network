package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"social-network/backend/middleware"
	"social-network/backend/models"
	"social-network/backend/utils"
)

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	usID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)})
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
	fmt.Println("posts: ", post)
	post.User.Id = usID.String()
	post.Img = path
	post.Id = utils.NewUUID()

	post_created, errPost := h.service.CreatePost(&post)
	if errPost != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errPost.Status, Error: errPost.Error})
		return
	}
	utils.WriteDataBack(w, post_created)
}
