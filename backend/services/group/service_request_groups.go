package group

import (
	"social-network/backend/models"
)

func (gService *GroupService) RequestToCancel(userId, groupId string) (*models.Notif, *models.ErrorJson) {
	group, errJson := gService.gRepo.GetGroupDetails(groupId)
	if errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}

	// return Invitation not found if the user has not requested to join
	if errJson := gService.gRepo.RequestToCancel(userId, groupId); errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message}
	}

	// if the user is not a member, we return an error
	// always check the membership and also the the group is a valid one
	if errMembership := gService.CheckNotMember(groupId, userId); errMembership != nil {
		return nil, &models.ErrorJson{Status: errMembership.Status, Error: errMembership.Error, Message: errMembership.Message}
	}

	return &models.Notif{
		SenderId:   userId,
		RecieverId: group.GroupCreatorId,
		GroupId:    groupId,
		Type:       "group-join",
	}, nil
}

func (gService *GroupService) RequestToJoin(userId, groupId string) (*models.Notification, *models.ErrorJson) {
	if errJson := gService.gRepo.GetGroupById(groupId); errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}

	// return ERROR!! Already a member! if the user is already a member
	// always check the membership and also the the group is a valid one
	if errMembership := gService.CheckNotMember(groupId, userId); errMembership != nil {
		return nil, &models.ErrorJson{Status: errMembership.Status, Error: errMembership.Error, Message: errMembership.Message}
	}

	notification, errJson := gService.gRepo.RequestToJoin(userId, groupId)
	if errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}

	return notification, nil
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
