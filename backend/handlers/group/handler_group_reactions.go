package group

import (
	"encoding/json"
	"io"
	"net/http"

	"social-network/backend/middleware"
	"social-network/backend/models"
	gservice "social-network/backend/services/group"
	"social-network/backend/utils"
)

type GroupReactionHanlder struct {
	gService *gservice.GroupService
}

func NewReactionHandler(service *gservice.GroupService) *GroupReactionHanlder {
	return &GroupReactionHanlder{gService: service}
}

func (Rhanlder *GroupReactionHanlder) LikeEntity(w http.ResponseWriter, r *http.Request) {
	userID, errParse := middleware.GetUserIDFromContext(r.Context())
	if errParse != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: errParse.Error()})
		return
	}

	groupID, err := utils.GetUUIDFromPath(r, "group_id")
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "ERROR!! Incorrect UUID Format!"})
		return
	}

	if errJson := Rhanlder.gService.GroupExists(groupID.String()); errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{
			Status:  errJson.Status,
			Error:   errJson.Error,
			Message: errJson.Message,
		})
		return
	}

	liked := models.GroupReaction{}
	if err := json.NewDecoder(r.Body).Decode(&liked); err != nil {
		if err == io.EOF {
			utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: &models.GroupReactionErr{
				EntityId:   "empty EntityID field!",
				EntityType: "empty EntityType field!",
			}})
			return
		}
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: " Bad Request!"})
		return
	}
	liked.UserId = userID.String()
	reaction, errJson := Rhanlder.gService.HanldeReaction(&liked, 1)
	if errJson != nil {
		utils.WriteJsonErrors(w, *errJson)
		return
	}

	utils.WriteDataBack(w, reaction)
}

func (RHanlder *GroupReactionHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodPost:
		RHanlder.LikeEntity(w, r)
		return
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "ERROR! Method Not allowed!!"})
		return
	}
}

// another one has to take up this task not me !!!
