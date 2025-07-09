package group

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"social-network/backend/models"
	gservice "social-network/backend/services/group"
	"social-network/backend/utils"

	"github.com/google/uuid"
)

// for the one group only

/***  /api/groups/{group_id}/   ***/

type GroupIDHanlder struct {
	gservice *gservice.GroupService
}

func NewGroupIDHandler(service *gservice.GroupService) *GroupIDHanlder {
	return &GroupIDHanlder{gservice: service}
}

// if the user has already joined the group : unauthorized important case
func (gIdHanlder *GroupIDHanlder) JoinGroup(w http.ResponseWriter, r *http.Request) {
	// userIDVal := r.Context().Value("userID")
	// userID, ok := userIDVal.(uuid.UUID)
	// if !ok {
	// 	utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Incorrect type of userID value!"})
	// 	return
	// }
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
			Message: "an error occured while trying to decode the json! hna",
		})
		return
	}

	if errJson := gIdHanlder.gservice.JoinGroup(group_to_join, "c95e60d2-3a86-4ec8-8120-8c9f85fbc172"); errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Message: errJson.Message})
		return
	}
}

// always get the groupInfo (general info like if the number of users and name and description)
func (gIdHanlder *GroupIDHanlder) GetGroupInfo(w http.ResponseWriter, r *http.Request) {
	userIDVal := r.Context().Value("userID")
	_, ok := userIDVal.(uuid.UUID)
	if !ok {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Incorrect type of userID value!"})
		return
	}
	groupId := r.PathValue("groupd_id")

	groupDetails, errJson := gIdHanlder.gservice.GetGroupInfo(groupId)
	if errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Message: errJson.Message})
		return
	}

	if err := json.NewEncoder(w).Encode(groupDetails); err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)})
		return
	}
}

func (gIdHanlder *GroupIDHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("hnaaaa inside the group id handler!!!!")
	switch r.Method {
	case http.MethodGet:
		gIdHanlder.GetGroupInfo(w, r)
		return
	case http.MethodPost:
		gIdHanlder.JoinGroup(w, r)
		return
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "ERROR!! Method Not Allowed!"})
		return
	}
}
