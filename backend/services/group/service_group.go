package group

import (
	"social-network/backend/models"
)

func (gService *GroupService) JoinGroup(group *models.Group, userId string) {
	gService.grepo.JoinGroup(group , userId)
}



func (gServicde *GroupService) GetGroupInfo() {
	
}