package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"social-network/backend/middleware"
	"social-network/backend/models"
	ps "social-network/backend/services"
	"social-network/backend/utils"
)

type PostHandler struct {
	service *ps.PostService
}

func NewPostHandler(service *ps.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.service.GetAllPosts()
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Failed to get posts"})
		return
	}
	utils.WriteDataBack(w, posts)
}

func (h *PostHandler) GetPostByID(w http.ResponseWriter, r *http.Request, postID string) {
	post, err := h.service.GetPostByID(postID)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "Post not found"})
		return
	}
	utils.WriteDataBack(w, post)
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*********** creating posts ***********")

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Invalid form data"})
		return
	}

	// Optional file upload
	file, handler, err := r.FormFile("img")
	imagePath := ""
	if err == nil {
		defer file.Close()
		imagePath = "uploads/posts/" + handler.Filename
		dst, err := os.Create(imagePath)
		if err != nil {
			utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Failed to save image"})
			return
		}
		defer dst.Close()
		io.Copy(dst, file)
	} else {
		imagePath = "" // no image uploaded
	}

	usID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 401, Message: "Unauthorized"})
		return
	}

	// Parse selected followers if privacy is private
	selectedFollowers := r.FormValue("selectedFollowers")
	var selectedUserIDs []string
	if selectedFollowers != "" {
		err := json.Unmarshal([]byte(selectedFollowers), &selectedUserIDs)
		if err != nil {
			utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Invalid selectedFollowers"})
			return
		}
	}

	// Create post model
	post := models.Post{
		ID:            utils.NewUUID(),
		UserID:        usID.String(),
		Content:       r.FormValue("content"),
		Privacy:       r.FormValue("privacy"),
		Img:           "/" + imagePath,
		SelectedUsers: selectedUserIDs, // ✅ <-- Don't forget this!
	}

	// Save to DB
	if err := h.service.CreatePost(&post); err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Failed to create post"})
		return
	}

	utils.WriteDataBack(w, map[string]string{
		"message": "Post created successfully",
	})
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request, postID string) {
	if err := h.service.DeletePost(postID); err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Failed to delete post"})
		return
	}
	utils.WriteDataBack(w, map[string]string{"message": "Post deleted successfully"})
}

func (h *PostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*************** serving request  ****************")
	w.Header().Set("Content-Type", "application/json")

	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	if len(pathParts) < 2 || pathParts[0] != "api" || pathParts[1] != "posts" {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "Not found"})
		return
	}

	postID := ""
	if len(pathParts) >= 3 {
		postID = pathParts[2]
	}

	switch r.Method {

	case http.MethodGet:
		if postID == "" {
			h.GetAllPosts(w, r)
		} else {
			h.GetPostByID(w, r, postID)
		}
	case http.MethodPost:
		if postID == "" {
			h.CreatePost(w, r)
		} else {
			utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "Method not allowed on this endpoint"})
		}
	case http.MethodDelete:
		if postID != "" {
			h.DeletePost(w, r, postID)
		} else {
			utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "Method not allowed"})
		}
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "Method not allowed"})
	}
}
