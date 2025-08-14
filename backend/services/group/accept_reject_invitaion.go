package group

import (
	"social-network/backend/models"
	"social-network/backend/utils"
)

func (gService *GroupService) Accept(userId, groupId string, userToBeAddedId string) *models.ErrorJson {
	// to approve wheter it we need
	// tanyan l user_id lakhur ykun valid (format , and aslo kayn f db)
	// talitan add the user_id to the the members of the group
	// rabi3an delete the request from the table of the requets-join
	// must wrap this validation inside a function and then wrap the logic of validation for each one

	if errJson := gService.gRepo.GetGroupById(groupId); errJson != nil {
		return &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}

	// validate the format of the user to be added
	if err := utils.IsValidUUID(userToBeAddedId); err != nil {
		return &models.ErrorJson{Status: 400, Message: models.UserErr{
			UserId: "Incorrect UUID Format!",
		}}
	}

	_, exists, _ := gService.gRepo.GetItem("users", "userID", userToBeAddedId)
	if !exists {
		return &models.ErrorJson{Status: 400, Message: models.UserErr{
			UserId: "user not found",
		}}
	}	
	// validate if wheter exists or not !!

	if errJson := gService.gRepo.Accept(userId, groupId, userToBeAddedId); errJson != nil {
		return &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	return nil
}

func (gService *GroupService) Reject(userId, groupId string, userToBeRejectedId string) *models.ErrorJson {
	if errJson := gService.gRepo.GetGroupById(groupId); errJson != nil {
		return &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	isAdmin, errJson := gService.gRepo.IsAdmin(groupId, userId)
	if errJson != nil {
		return &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	if !isAdmin {
		return &models.ErrorJson{Status: 403, Error: "ERROR!! Access Forbidden"}
	}
	// validate the format of the user to be added
	if err := utils.IsValidUUID(userToBeRejectedId); err != nil {
		return &models.ErrorJson{Status: 400, Message: models.UserErr{
			UserId: "ERROR!! Incorrect UUID Format!",
		}}
	}

	_, exists, _ := gService.gRepo.GetItem("users", "userID", userToBeRejectedId)
	if !exists {
		return &models.ErrorJson{Status: 400, Message: models.UserErr{
			UserId: "ERROR!! user not found",
		}}
	}
	if errJson := gService.gRepo.Reject(userId, groupId, userToBeRejectedId); errJson != nil {
		return &models.ErrorJson{Status: errJson.Status, Message: errJson.Message, Error: errJson.Error}
	}
	return nil
}
