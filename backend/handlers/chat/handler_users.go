package chat

import (
	"fmt"
	"net/http"

	"social-network/backend/middleware"
	"social-network/backend/models"
	"social-network/backend/services/chat"
	"social-network/backend/utils"
)

type ChatNavigation struct {
	service *chat.ChatService
}

func NewChatNavigation(service *chat.ChatService) *MessagesHandler {
	return &MessagesHandler{service: service}
}

func (chatNav *ChatNavigation) GetUsers(w http.ResponseWriter, r *http.Request) {
	authUserID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		fmt.Println("here an error: ", err)
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: fmt.Sprintf("%v", err)})
		return
	}

	users, errUsers := chatNav.service.GetUsers(authUserID.String())
	if errUsers != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errUsers.Status, Message: errUsers.Message})
		return
	}

	utils.WriteDataBack(w, users)
}

func (chatNav *ChatNavigation) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "Method not allowed !"})
		return
	}
	switch r.URL.Path {
	case "/api/get-users/":
		chatNav.GetUsers(w, r)
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "ERROR!! Page Not Found!"})
		return
	}
}
