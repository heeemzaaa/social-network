package group

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strconv"

	"social-network/backend/models"
	gservice "social-network/backend/services/group"
	"social-network/backend/utils"

	"github.com/google/uuid"
)

/***    /api/groups/{group_id}/posts/    ***/

type GroupPostHandler struct {
	gService *gservice.GroupService
}

func NewGroupPostHandler(service *gservice.GroupService) *GroupPostHandler {
	return &GroupPostHandler{gService: service}
}

func (gPostHandler *GroupPostHandler) AddGroupPost(w http.ResponseWriter, r *http.Request) {
	var post *models.PostGroup
	groupId := r.PathValue("group_id")
	groupID, err := uuid.Parse(groupId)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Incorrect type of groupID value!"})
		return
	}
	userIDVal := r.Context().Value("userID")
	userID, errParse := uuid.Parse(userIDVal.(string))
	if errParse != nil {
		fmt.Println("errors", errParse)
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Incorrect type of userID value!"})
		return
	}

	data := r.FormValue("data")
	if err := json.Unmarshal([]byte(data), &post); err != nil {
		if err != nil {
			if err == io.EOF {
				utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: &models.PostGroupErr{
					Content: "empty Content field!",
				}})
				return
			}
			// which status code to return
			utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: fmt.Sprintf("%v", err)})
			return
		}
	}

	// handle the image encoding in the phase that comes before the adding process
	path, errUploadImg := utils.HanldeUploadImage(r, "post", filepath.Join("groups", "posts"), false)
	if errUploadImg != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errUploadImg.Status, Message: errUploadImg.Message})
		return
	}
	post.UserId, post.GroupId,post.ImagePath = &userID,&groupID ,path
	// even if the userid is given wrong we insert the correct one
	postCreated, err_ := gPostHandler.gService.AddPost(post)
	if err_ != nil {
		utils.WriteJsonErrors(w, *err_)
		return
	}
	utils.WriteDataBack(w, postCreated)
}

func (gPostHandler *GroupPostHandler) GetGroupPosts(w http.ResponseWriter, r *http.Request) {
	userIDVal := r.Context().Value("userID")
	userID, errParse := uuid.Parse(userIDVal.(string))
	if errParse != nil {
		fmt.Println("errors", errParse)
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Incorrect type of userID value!"})
		return
	}
	offset, errConvoff := strconv.Atoi(r.URL.Query().Get("offset"))
	if errConvoff != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "ERROR!! Incorrect offset"})
		return
	}

	posts, err_get := gPostHandler.gService.GetPosts(userID.String(), offset)
	if err_get != nil {
		utils.WriteJsonErrors(w, *err_get)
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
