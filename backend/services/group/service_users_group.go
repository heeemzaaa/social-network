package group

import "social-network/backend/models"

func (gService *GroupService) IsMemberGroup(groupId, userId string) (bool, *models.ErrorJson) {
	isMember, errJson := gService.gRepo.IsMemberGroup(groupId, userId)
	if errJson != nil {
		return isMember, &models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message}
	}
	return isMember, nil
}
