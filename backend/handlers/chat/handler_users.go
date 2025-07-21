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

func NewChatNavigation(service *chat.ChatService) *ChatNavigation {
	return &ChatNavigation{service: service}
}

func (chatNav *ChatNavigation) GetUsers(w http.ResponseWriter, r *http.Request) {
	authUserID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
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

func (ChatNav *ChatNavigation) GetGroups(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching groups for authenticated user")
	authUserID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: fmt.Sprintf("%v", err)})
		return
	}

	groups, errGroups := ChatNav.service.GetUsers(authUserID.String())
	if errGroups != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: fmt.Sprintf("%v", err)})
		return
	}

	utils.WriteDataBack(w, groups)
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
	case "/api/get-groups/":
		chatNav.GetGroups(w, r)
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "ERROR!! Page Not Found!"})
		return
	}
}
