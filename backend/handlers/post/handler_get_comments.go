package handlers

import (
	"net/http"

	"social-network/backend/models"
	s "social-network/backend/services/post"
	"social-network/backend/utils"
)

type CommentsHandler struct {
	service *s.PostService
}

func NewCommentsHandler(service *s.PostService) *CommentsHandler {
	return &CommentsHandler{service: service}
}

func (h *CommentsHandler) GetComments(w http.ResponseWriter, r *http.Request) {
	postID, err := utils.GetUUIDFromPath(r, "id")
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "invalid path"})
		return
	}

	comments, errComments := h.service.GetComments(postID.String())
	if errComments != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errComments.Status, Error: errComments.Error})
		return
	}

	utils.WriteDataBack(w, comments)
}

func (h *CommentsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetComments(w, r)
	case http.MethodPost:
		// hna zid dyal add comment
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "Method not allowed !"})
		return
	}
}
