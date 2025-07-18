package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"social-network/backend/middleware"
	"social-network/backend/models"
	ps "social-network/backend/services/post"
	"social-network/backend/utils"
)

type PostHandler struct {
	service *ps.PostService
}

func NewPostHandler(service *ps.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	usID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Failed to get userID"})
		return
	}
	posts, err := h.service.GetAllPosts(usID)
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
	usID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 401, Message: "Unauthorized"})
		return
	}

	path, errUploadImg := utils.HanldeUploadImage(r, "img", "posts")
	if errUploadImg != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errUploadImg.Status, Message: errUploadImg.Message})
		return
	}

	dataStr := r.FormValue("data")
	if dataStr == "" {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Missing data field"})
		return
	}

	var post models.Post
	if err := json.Unmarshal([]byte(dataStr), &post); err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Invalid JSON in data field"})
		return
	}

	post.User = models.User{Id: usID.String()}
	post.Img = path
	post.Id = utils.NewUUID()
	post.CreatedAt = time.Now().Format(time.RFC3339)

	if err := h.service.CreatePost(&post); err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Failed to create post"})
		return
	}
	utils.WriteDataBack(w, post)
}

func (h *PostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	if len(pathParts) < 2 || pathParts[0] != "api" || pathParts[1] != "posts" {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "Not found"})
		return
	}

	switch r.Method {
	case http.MethodGet:
		if len(pathParts) == 2 {
			h.GetAllPosts(w, r)
			return
		}
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "Not found"})
	case http.MethodPost:
		if len(pathParts) == 2 {
			h.CreatePost(w, r)
			return
		}
		if len(pathParts) == 4 && pathParts[2] == "like" {
			h.LikePost(w, r)
			return
		}
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "Method not allowed on this endpoint"})
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "Method not allowed"})
	}
}

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

	err = h.service.HandleLike(postID, userID)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: http.StatusInternalServerError, Message: "Failed to like post"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "Post like updated successfully",
		"postId":  postID,
	})
}
