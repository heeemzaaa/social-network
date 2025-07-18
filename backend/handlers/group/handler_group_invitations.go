package group

import (
	"net/http"

	gservice "social-network/backend/services/group"
)

type GroupInvitationHandler struct {
	gService *gservice.GroupService
}

func NewGroupInvitationHandler(service *gservice.GroupService) *GroupInvitationHandler {
	return &GroupInvitationHandler{gService: service}
}

func (GrpInviatuionH *GroupInvitationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
