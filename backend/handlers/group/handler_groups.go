package group

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"social-network/backend/models"
	gservice "social-network/backend/services/group"
	"social-network/backend/utils"

	"github.com/google/uuid"
)

type GroupHanlder struct {
	gservice *gservice.GroupService
}

func NewGroupHandler(service *gservice.GroupService) *GroupHanlder {
	return &GroupHanlder{gservice: service}
}

func (Ghandler *GroupHanlder) GetGroups(w http.ResponseWriter, r *http.Request) {
	userIDVal := r.Context().Value("userID")
	userID, ok := userIDVal.(uuid.UUID)
	if !ok {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Incorrect type of userID value!"})
		return
	}
	filter := r.URL.Query().Get("filter")
	offset, errOffset := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 64)
	if errOffset != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Incorrect Offset Format!"})
		return
	}
	if !utils.IsValidFilter(filter) {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Incorrect filter by field!!"})
		return
	}
	Ghandler.gservice.GetGroups(filter, offset, userID.String())
}

func (Ghandler *GroupHanlder) CreateGroup(w http.ResponseWriter, r *http.Request) {
	userIDVal := r.Context().Value("userID")
	userID, ok := userIDVal.(uuid.UUID)
	if !ok {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: "Incorrect type of userID value!"})
		return
	}
	var group_to_create *models.Group
	err := json.NewDecoder(r.Body).Decode(&group_to_create)
	if err != nil {
		if err == io.EOF {
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
			Message: "an error occured while trying to decode the json!",
		})
		return
	}

	group_to_create.GroupCreatorId = userID
	// handle the image encoding in the phase that comes before the adding process
	if group_to_create.ImageEncoded != "" {
	}

	Ghandler.gservice.AddGroup(group_to_create)
}

func (Ghandler *GroupHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appplication/json")
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
