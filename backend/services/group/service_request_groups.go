package group

import "social-network/backend/models"

func (gService *GroupService) RequestToCancel(userId, groupId string) *models.ErrorJson {
	if errJson := gService.gRepo.GetGroupById(groupId); errJson != nil {
		return  &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	// always check the membership and also the the group is a valid one
	if errMembership := gService.CheckNotMember(groupId, userId); errMembership != nil {
		return &models.ErrorJson{Status: errMembership.Status, Error: errMembership.Error, Message: errMembership.Message}
	}

	gService.gRepo.RequestToCancel()
	return nil
}

func (gService *GroupService) RequestToJoin(userId, groupId string) *models.ErrorJson {
	if errJson := gService.gRepo.GetGroupById(groupId); errJson != nil {
		return  &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	// always check the membership and also the the group is a valid one
	if errMembership := gService.CheckNotMember(groupId, userId); errMembership != nil {
		return &models.ErrorJson{Status: errMembership.Status, Error: errMembership.Error, Message: errMembership.Message}
	}

	gService.gRepo.RequestToJoin()
	return nil
}
