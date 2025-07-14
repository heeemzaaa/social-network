package group

import "social-network/backend/models"

func (gService *GroupService) IsMemberGroup(groupId, userId string) (bool, *models.ErrorJson) {
	return   gService.gRepo.IsMemberGroup(groupId, userId)
}




