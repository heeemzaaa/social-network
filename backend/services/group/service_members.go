package group

import "social-network/backend/models"

func (gService *GroupService) GetGroupMembers(groupId, userId string) ([]models.User, *models.ErrorJson) {
	if errJson := gService.gRepo.GetGroupById(groupId); errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	// always check the membership and also the the group is a valid one
	if errMembership := gService.CheckMembership(groupId, userId); errMembership != nil {
		return nil, &models.ErrorJson{Status: errMembership.Status, Error: errMembership.Error, Message: errMembership.Message}
	}

	users, errJson := gService.gRepo.GetGroupMembers(groupId)
	if errJson != nil {
		return nil, errJson
	}
	return users, nil
}
