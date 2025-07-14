package group

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"social-network/backend/middleware"
	"social-network/backend/models"
	gservice "social-network/backend/services/group"
	"social-network/backend/utils"
)

// Not the latest version yet just zrbt 3liha
// need to come back at night

/***   /api/groups/   ***/

type GroupHanlder struct {
	gService *gservice.GroupService
}

func NewGroupHandler(gservice *gservice.GroupService) *GroupHanlder {
	return &GroupHanlder{gService: gservice}
}

// we only need the userId and filter based on owned, availabe and created
func (Ghandler *GroupHanlder) GetGroups(w http.ResponseWriter, r *http.Request) {
	userID, errParse := middleware.GetUserIDFromContext(r.Context())
	if errParse != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: errParse.Error()})
		return
	}
	filter := r.URL.Query().Get("filter")
	if !utils.IsValidFilter(filter) {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Incorrect filter by field!!"})
		return
	}
	offset, errOffset := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 64)
	if errOffset != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Incorrect Offset Format!"})
		return
	}
	groups, errJson := Ghandler.gService.GetGroups(filter, offset, userID.String())
	if errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Message: errJson.Message})
		return
	}
	if err := json.NewEncoder(w).Encode(groups); err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)})
		return
	}
}

// tested
// but needs the context to be there to test out other things

func (Ghandler *GroupHanlder) CreateGroup(w http.ResponseWriter, r *http.Request) {
	userID, errParse := middleware.GetUserIDFromContext(r.Context())
	if errParse != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: errParse.Error()})
		return
	}
	var group_to_create *models.Group

	data := r.FormValue("data")
	if err := json.Unmarshal([]byte(data), &group_to_create); err != nil {
		if err == io.EOF  || group_to_create ==(&models.Group{}){
			utils.WriteJsonErrors(w, models.ErrorJson{
				Status: 400,
				Message: models.ErrGroup{
					Title:       "empty title field!",
					Description: "empty description field!",
				},
			})
			return
		}

		utils.WriteJsonErrors(w, models.ErrorJson{
			Status:  400,
			Message: "ERROR!! Can not Unmarshal the data!",
		})
		return
	}

	// handle the image encoding in the phase that comes before the adding process
	path, errUploadImg := utils.HanldeUploadImage(r, "group", "groups", true)
	if errUploadImg != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errUploadImg.Status, Message: errUploadImg.Message})
		return
	}

	group_to_create.GroupCreatorId, group_to_create.ImagePath = userID.String(), path
	group, errJson := Ghandler.gService.AddGroup(group_to_create)
	if errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Message: errJson.Message})
		return
	}
	utils.WriteDataBack(w, group)
}

func (Ghandler *GroupHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		Ghandler.GetGroups(w, r)
		return
	case http.MethodPost:
		Ghandler.CreateGroup(w, r)
		return
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "ERROR!! Method Not Allowed!"})
		return
	}
}
