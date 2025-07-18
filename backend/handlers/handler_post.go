package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

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
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Invalid form data"})
		return
	}

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
		imagePath = ""
	}

	usID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 401, Message: "Unauthorized"})
		return
	}

	selectedFollowers := r.FormValue("selectedFollowers")
	var selectedUserIDs []string
	if selectedFollowers != "" {
		err := json.Unmarshal([]byte(selectedFollowers), &selectedUserIDs)
		if err != nil {
			utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Invalid selectedFollowers"})
			return
		}
	}
	var user models.User
	user.Id = usID.String()
	post := models.Post{
		Id:            utils.NewUUID(),
		User:          user,
		Content:       r.FormValue("content"),
		Privacy:       r.FormValue("privacy"),
		Img:           "/" + imagePath,
		SelectedUsers: selectedUserIDs,
		CreatedAt:     time.Now().String(),
	}

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

	thirdEndPoint := ""
	if len(pathParts) >= 3 {
		thirdEndPoint = pathParts[2]
	}
	switch r.Method {

	case http.MethodGet:
		if thirdEndPoint == "" {
			h.GetAllPosts(w, r)
		} else {
			h.GetPostByID(w, r, thirdEndPoint)
		}
	case http.MethodPost:
		if thirdEndPoint != "" {
			if len(pathParts) >= 3 && pathParts[2] == "like" {
				h.LikePost(w, r)
				return
			}
			h.CreatePost(w, r)
		} else {
			utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "Method not allowed on this endpoint"})
		}
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
		"postId" : postID,
		})
}
