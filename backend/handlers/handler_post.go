package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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
	var post models.Post

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Invalid request body"})
		return
	}

	if err := h.service.CreatePost(&post); err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Failed to create post"})
		return
	}
	
	utils.WriteDataBack(w, map[string]string{"message": "Post created successfully"})
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
