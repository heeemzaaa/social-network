package group

import (
	"social-network/backend/models"
)

func (gService *GroupService) InviteToJoin(userId, groupId string, userToInvite models.User) *models.ErrorJson {
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

	isFollower, errJson := gService.sProfile.IsFollower(userId, userToInvite.Id)
	if errJson != nil {
		return &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	if !isFollower {
		return &models.ErrorJson{Status: 403, Error: "ERROR!! it is not from your followers!"}
	}
	if err := gService.gRepo.InviteToJoin(userId, groupId, userToInvite.Id); err != nil {
		return &models.ErrorJson{Status: err.Status, Error: err.Error, Message: err.Message}
	}

	// this was a slight edit for the user to see only :)
	// i hope it works

	// add the notification service method to be able to add a user
	// {sneder_id, receiver_id , "group-invitation"}
	return nil
}

func (gService *GroupService) CancelTheInvitation(userId, groupId string, invitedUser *models.User) *models.ErrorJson {
	// check the group if a valid one
	// check the user is member before he can invite
	// we need to check if the request of invitation is there before canceling it
	// delete  the invitation from the table of the requests
	if errJson := gService.gRepo.GetGroupById(groupId); errJson != nil {
		return &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	// always check the membership and also the the group is a valid one
	if errMembership := gService.CheckMembership(groupId, userId); errMembership != nil {
		return &models.ErrorJson{Status: errMembership.Status, Error: errMembership.Error, Message: errMembership.Message}
	}

	if errJson := gService.gRepo.CancelTheInvitation(userId, groupId, invitedUser); errJson != nil {
		return &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	return nil
}

/*
:[follower1 : {
fullname
nickname invited ot not (1/0)
}]
*/
func (gService *GroupService) GetUsersToInvite(userID, groupID string) ([]models.User, *models.ErrorJson) {
	// check if thye group is valid
	//  we need to check if the user is a member of the group before he proceeds to invite them
	//  we need to check if
	if errJson := gService.gRepo.GetGroupById(groupID); errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	// always check the membership and also the the group is a valid one
	if errMembership := gService.CheckMembership(groupID, userID); errMembership != nil {
		return nil, &models.ErrorJson{Status: errMembership.Status, Error: errMembership.Error, Message: errMembership.Message}
	}

	users, errJson := gService.gRepo.GetUsersToInvite(userID, groupID)
	if errJson != nil {
		return nil, &models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message}
	}

	return users, nil
}
