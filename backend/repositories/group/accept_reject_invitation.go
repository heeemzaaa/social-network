package group

import (
	"fmt"

	"social-network/backend/models"
)

func (gRepo *GroupRepository) Accept(userId, groupId string, invitedUser *models.User) *models.ErrorJson {
	query := `
	INSERT INTO group_membership (groupID,userID) 
	VALUES (?,?)
	`
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	_, err = stmt.Exec(groupId, invitedUser.Id)
	if err != nil {
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	if errJson := gRepo.CancelTheInvitation(userId, groupId, invitedUser); errJson != nil {
		return &models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message}
	}
	return nil
}

func (gRepo *GroupRepository) Reject(userId, groupId string, userToBeRejected *models.User) *models.ErrorJson {
	if errJson := gRepo.CancelTheInvitation(userId, groupId, userToBeRejected); errJson != nil {
		return &models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message}
	}
	return nil
}
