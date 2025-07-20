package group

import (
	"social-network/backend/models"
)

func (gService *GroupService) RequestToCancel(userId, groupId string) *models.ErrorJson {
	if errJson := gService.gRepo.GetGroupById(groupId); errJson != nil {
		return &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	// always check the membership and also the the group is a valid one
	if errMembership := gService.CheckNotMember(groupId, userId); errMembership != nil {
		return &models.ErrorJson{Status: errMembership.Status, Error: errMembership.Error, Message: errMembership.Message}
	}

	if errJson := gService.gRepo.RequestToCancel(userId, groupId); errJson != nil {
		return &models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message}
	}
	return nil
}

func (gService *GroupService) RequestToJoin(userId, groupId string) *models.ErrorJson {
	if errJson := gService.gRepo.GetGroupById(groupId); errJson != nil {
		return &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	// always check the membership and also the the group is a valid one
	if errMembership := gService.CheckNotMember(groupId, userId); errMembership != nil {
		return &models.ErrorJson{Status: errMembership.Status, Error: errMembership.Error, Message: errMembership.Message}
	}

	if errJson := gService.gRepo.RequestToJoin(userId, groupId); errJson != nil {
		return &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	return nil
}

func (gService *GroupService) GetRequests(userId, groupId string) ([]models.User, *models.ErrorJson) {
	if errJson := gService.gRepo.GetGroupById(groupId); errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	// WE NEED to check that the userID is the one of the admin otherwise
	// we need to return unothorized
	isAdmin, errJson := gService.gRepo.IsAdmin(groupId, userId)
	if errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	// if is not the admin he has no right to see the resources
	
	if !isAdmin {
		return nil, &models.ErrorJson{Status: 403, Error: "ERROR!! Access Forbidden"}
	}
	users, errJson := gService.gRepo.GetRequests(groupId)
	if errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message}
	}
	return users, nil
}
