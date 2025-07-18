package handlers

import (
	"net/http"

	group "social-network/backend/services/group"
	profile "social-network/backend/services/profile"
)

type NotificationsHandler struct {
	gService *group.GroupService
	pService *profile.ProfileService
}

func NewNotificationsHandler(gService *group.GroupService, pService *profile.ProfileService) *NotificationsHandler {
	return &NotificationsHandler{gService: gService, pService: pService}
}

func (Nhandler *NotificationsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	groupID := "0a3b66fc-367c-43e3-bb42-c28c5ec161d3"
	Nhandler.gService.GetGroupInfo(groupID)
}
