package group

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"social-network/backend/middleware"
	"social-network/backend/models"
	gservice "social-network/backend/services/group"
	"social-network/backend/utils"
)

/***    /api/groups/{group_id}/posts/    ***/
// not tested yet

type GroupPostHandler struct {
	gService *gservice.GroupService
}

func NewGroupPostHandler(service *gservice.GroupService) *GroupPostHandler {
	return &GroupPostHandler{gService: service}
}

func (gPostHandler *GroupPostHandler) AddGroupPost(w http.ResponseWriter, r *http.Request) {
	var post *models.PostGroup
	userID, errParse := middleware.GetUserIDFromContext(r.Context())
	if errParse != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: "Incorrect type of userID value!"})
		return
	}
	groupID, err := utils.GetUUIDFromPath(r, "group_id")
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: "Incorrect type of groupID value!"})
		return
	}

	data := r.FormValue("data")
	if err := json.Unmarshal([]byte(data), &post); err != nil {
		if err == io.EOF {
			utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: &models.PostGroupErr{
				Content: "empty Content field!",
			}})
			return
		}
		if strings.TrimSpace(data) == "" || post == (&models.PostGroup{}) {
			utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "data is empty"})
			return
		}
		// which status code to return
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: fmt.Sprintf("%v hhh", err)})
		return

	}

	// handle the image encoding in the phase that comes before the adding process
	path, errUploadImg := utils.HanldeUploadImage(r, "post", filepath.Join("groups", "posts"))
	if errUploadImg != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errUploadImg.Status, Message: errUploadImg.Message})
		return
	}
	post.UserId, post.GroupId, post.ImagePath = userID.String(), groupID.String(), path
	// even if the userid is given wrong we insert the correct one
	postCreated, err_ := gPostHandler.gService.AddPost(post)

	fmt.Println("post created", postCreated, err_)
	if err_ != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: err_.Status, Error: err_.Error, Message: err_.Message})
		return
	}
	utils.WriteDataBack(w, postCreated)
}

func (gPostHandler *GroupPostHandler) GetGroupPosts(w http.ResponseWriter, r *http.Request) {
	userID, errParse := middleware.GetUserIDFromContext(r.Context())
	if errParse != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: "Incorrect type of userID value!"})
		return
	}
	groupID, err := utils.GetUUIDFromPath(r, "group_id")
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: "Incorrect type of groupID value!"})
		return
	}

	offset, errConvoff := strconv.Atoi(r.URL.Query().Get("offset"))
	if errConvoff != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "ERROR!! Incorrect offset"})
		return
	}

	posts, err_get := gPostHandler.gService.GetPosts(userID.String(), groupID.String(), offset)
	if err_get != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: err_get.Status, Error: err_get.Error, Message: err_get.Message})
		return
	}

	err_ := json.NewEncoder(w).Encode(posts)
	if err_ != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err_)})
		return
	}
}

func (gPostHandler *GroupPostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		gPostHandler.GetGroupPosts(w, r)
		return
	case http.MethodPost:
		gPostHandler.AddGroupPost(w, r)
		return
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "ERROR!! Method not allowed!"})
		return
	}
}
