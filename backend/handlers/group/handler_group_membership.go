package group

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"social-network/backend/middleware"
	"social-network/backend/models"
	gservice "social-network/backend/services/group"
	"social-network/backend/utils"
)

// for the one group only
// POST request will help us join the group
// GET request will help us get the details of the a specific group

/***  /api/groups/{group_id}/   ***/

// DONE tooo 

type GroupIDHanlder struct {
	gservice *gservice.GroupService
}

func NewGroupIDHandler(service *gservice.GroupService) *GroupIDHanlder {
	return &GroupIDHanlder{gservice: service}
}

// if the user has already joined the group : unauthorized important case
func (gIdHanlder *GroupIDHanlder) JoinGroup(w http.ResponseWriter, r *http.Request) {
	userID, errParse := middleware.GetUserIDFromContext(r.Context())
	if errParse != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: "Incorrect format of userID UUID!"})
		return
	}
	var group_to_join *models.Group
	err := json.NewDecoder(r.Body).Decode(&group_to_join)
	if err != nil {
		if err == io.EOF {
			utils.WriteJsonErrors(w, models.ErrorJson{
				Status: 400,
				Message: models.ErrJoinGroup{
					GroupId: "empty group_id field!",
				},
			})
			return
		}

		utils.WriteJsonErrors(w, models.ErrorJson{
			Status:  400,
			Message: "an error occured while trying to decode the json!",
		})
		return
	}

	fmt.Println("userId", userID.String(), group_to_join.GroupId)

	if errJson := gIdHanlder.gservice.JoinGroup(group_to_join, userID.String()); errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error})
		return
	}
}

// always get the groupInfo (general info like if the number of users and name and description)
func (gIdHanlder *GroupIDHanlder) GetGroupInfo(w http.ResponseWriter, r *http.Request) {
	groupId, err := utils.GetUUIDFromPath(r, "group_id")
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "ERROR!! Incorrect UUID Format!"})
		return
	}

	groupDetails, errJson := gIdHanlder.gservice.GetGroupInfo(groupId.String())
	if errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error})
		return
	}

	if err := json.NewEncoder(w).Encode(groupDetails); err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)})
		return
	}
}

func (gIdHanlder *GroupIDHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		gIdHanlder.GetGroupInfo(w, r)
		return
	case http.MethodPost:
		gIdHanlder.JoinGroup(w, r)
		return
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "ERROR!! Method Not Allowed!"})
		return
	}
}
