package chat

import (
	"fmt"
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
	lastMessageStr := r.URL.Query().Get("last")
	target_id := r.URL.Query().Get("target_id")
	fmt.Println("Target ID:", target_id)

	type_ := r.URL.Query().Get("type")
	if type_ != "private" && type_ != "group" {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "type is not specified"})
		return
	}

	exists, errJson := messages.service.CheckExistance(type_, target_id)
	if errJson != nil {
		utils.WriteJsonErrors(w, *errJson)
		return
	}

	if !exists {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "", Message: "the target given doesn't exist"})
		return
	}

	sender_id, errJson := messages.service.GetUserIdFromSession(r)
	if errJson != nil {
		utils.WriteJsonErrors(w, *errJson)
		return
	}

	mesages, errJson := messages.service.GetMessages(sender_id, target_id, lastMessageStr, type_)
	if errJson != nil {
		utils.WriteJsonErrors(w, *models.NewErrorJson(errJson.Status, "", errJson.Message))
		return
	}

	utils.WriteDataBack(w, mesages)
}

func (messages *MessagesHandler) UpdataReadStatus(w http.ResponseWriter, r *http.Request) {
	sender_id, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Invalid userID format !"})
		return
	}

	target_id := r.URL.Query().Get("target_id")

	exists, errJson := messages.service.UserExists(target_id)
	if errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Message: errJson.Message})
		return
	}

	// the one who is logged in is the one who opens  the tab
	// so basically the messages sent by the other (target_id) must be marked read
	if !exists {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "target_id Incorrect"})
		return
	}

	errJson = messages.service.EditReadStatus(sender_id.String(), target_id)
	if errJson != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Message: errJson.Message})
		return
	}
}


func (messages *MessagesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "ERROR!! Method Not Allowed!!"})
		return
	}
	fmt.Println(r.URL.Path)
	switch r.URL.Path {
	case "/api/messages":
		messages.GetMessages(w, r)
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "ERROR!! Page Not Found!"})
		return
	}
}
