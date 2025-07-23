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
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: fmt.Sprintf("%v", err)})
		return
	}

	users, errUsers := chatNav.service.GetUsers(authUserID.String())
	if errUsers != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errUsers.Status, Error: errUsers.Error})
		return
	}

	utils.WriteDataBack(w, users)
}

func (ChatNav *ChatNavigation) GetGroups(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching groups for authenticated user")
	authUserID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: fmt.Sprintf("%v", err)})
		return
	}

	groups, errGroups := ChatNav.service.GetGroups(authUserID.String())
	if errGroups != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: errGroups.Status, Error: errGroups.Error})
		return
	}

	utils.WriteDataBack(w, groups)
}

func (chatNav *ChatNavigation) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "Method not allowed !"})
		return
	}
	switch r.URL.Path {
	case "/api/get-users/":
		chatNav.GetUsers(w, r)
	case "/api/get-groups/":
		chatNav.GetGroups(w, r)
	default:
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 404, Error: "ERROR!! Page Not Found!"})
		return
	}
}
