package chat

import (
	"net/http"

	"social-network/backend/middleware"
	"social-network/backend/models"
	"social-network/backend/services/chat"
	"social-network/backend/utils"
)

type MessagesHandler struct {
	service *chat.ChatService
}

func NewMessagesHandler(service *chat.ChatService) *MessagesHandler {
	return &MessagesHandler{service: service}
}

// do we need to check the id of the receiver (if it exists in the database )
func (messages *MessagesHandler) GetMessages(w http.ResponseWriter, r *http.Request) {
	senderId, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "Invalid userId !"})
		return
	}

	targetId := r.URL.Query().Get("target_id")

	lastMessageID := r.URL.Query().Get("message_id")

	type_ := r.URL.Query().Get("type")
	if type_ != "private" && type_ != "group" {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "type is not specified"})
		return
	}

	exists, errJson := messages.service.CheckExistance(type_, targetId)
	if errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Error: errJson.Error})
		return
	}

	if !exists {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "the target given doesn't exist"})
		return
	}

	mesages, errJson := messages.service.GetMessages(senderId.String(), targetId, lastMessageID, type_)
	if errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Error: errJson.Error})
		return
	}

	utils.WriteDataBack(w, mesages)
}

func (messages *MessagesHandler) UpdateReadStatus(w http.ResponseWriter, r *http.Request) {
	senderId, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "Invalid userID format !"})
		return
	}

	targetId := r.URL.Query().Get("target_id")

	exists, errJson := messages.service.UserExists(targetId)
	if errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Error: errJson.Error})
		return
	}

	// the one who is logged in is the one who opens  the tab
	// so basically the messages sent by the other (target_id) must be marked read
	if !exists {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "targetId Incorrect"})
		return
	}

	errJson = messages.service.EditReadStatus(senderId.String(), targetId)
	if errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Error: errJson.Error})
		return
	}
}

func (messages *MessagesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "ERROR!! Method Not Allowed!!"})
		return
	}

	switch r.URL.Path {
	case "/api/messages":
		messages.GetMessages(w, r)
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 404, Error: "ERROR!! Page Not Found!"})
		return
	}
}
