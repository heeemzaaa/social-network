package chat

import (
	"encoding/json"
	"fmt"
	"net/http"
	h "social-network/backend/handlers"
	"social-network/backend/models"
	"social-network/backend/services/chat"
	"strconv"
)

type MessagesHandler struct {
	service *chat.ChatService
}

func NewMessagesHandler(service *chat.ChatService) *MessagesHandler {
	return &MessagesHandler{service: service}
}

// do we need to check the id of the receiver (if it exists in the database )
func (messages *MessagesHandler) GetMessages(w http.ResponseWriter, r *http.Request) {
	offset, errConvoff := strconv.Atoi(r.URL.Query().Get("offset"))
	if errConvoff != nil {
		h.WriteJsonErrors(w, *models.NewErrorJson(400, "", "Incorrect Format offset"))
		return
	}
	target_id := r.URL.Query().Get("target_id")

	type_ := r.URL.Query().Get("type")
	if type_ != "private" && type_ != "group" {
		h.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "type is not specified"})
		return
	}

	exists, errJson := messages.service.CheckExistance(type_, target_id)
	if errJson != nil {
		h.WriteJsonErrors(w, *errJson)
		return
	}

	if !exists {
		h.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "", Message: "the target given doesn't exist"})
		return
	}

	sender_id, errJson := messages.service.GetUserIdFromSession(r)
	if errJson != nil {
		h.WriteJsonErrors(w, *errJson)
		return
	}

	mesages, errJson := messages.service.GetMessages(sender_id, target_id, offset, type_)
	if errJson != nil {
		h.WriteJsonErrors(w, *models.NewErrorJson(errJson.Status, "", errJson.Message))
		return
	}
	err := json.NewEncoder(w).Encode(mesages)
	if err != nil {
		h.WriteJsonErrors(w, *models.NewErrorJson(500, "", fmt.Sprintf("%v", err)))
		return
	}
}

func (messages *MessagesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		h.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "ERROR!! Method Not Allowed!!"})
		return
	}
	messages.GetMessages(w, r)
}
