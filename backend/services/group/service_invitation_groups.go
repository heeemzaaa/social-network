package group

import "social-network/backend/models"

func (gService *GroupService) InviteToJoin(userId, groupId string) *models.ErrorJson {
	// check the group if a valid one
	// check the user is member before he can invite
	// check if the invited one is one of the followers of the user
	// add the invitation to the table of the requests
	if errJson := gService.gRepo.GetGroupById(groupId); errJson != nil {
		return &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	// always check the membership and also the the group is a valid one
	if errMembership := gService.CheckMembership(groupId, userId); errMembership != nil {
		return &models.ErrorJson{Status: errMembership.Status, Error: errMembership.Error, Message: errMembership.Message}
	}

	gService.gRepo.InviteToJoin(userId, groupId)
	return nil
}

func (gService *GroupService) CancelTheInvitation(userId, groupId string) *models.ErrorJson {
	// check the group if a valid one
	// check the user is member before he can invite
	// check if the invited one is one of the followers of the user
	// delete  the invitation from the table of the requests
	gService.gRepo.CancelTheInvitation(userId, groupId)
	return nil
}

func (gService *GroupService) GetInvitations(userId, groupId string) ([]models.User, *models.ErrorJson) {
	return nil, nil
}
