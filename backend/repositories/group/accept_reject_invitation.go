package group

import (
	"fmt"

	"social-network/backend/models"
)

func (gRepo *GroupRepository) Accept(userId, groupId, invitedUserId string) *models.ErrorJson {
	query := `
	INSERT INTO group_membership (groupID,userID) 
	VALUES (?,?)
	`
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	_, err = stmt.Exec(groupId, invitedUserId)
	if err != nil {
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	if errJson := gRepo.CancelTheInvitation(userId, groupId, invitedUserId); errJson != nil {
		return &models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message}
	}
	return nil
}

func (gRepo *GroupRepository) Reject(userId, groupId, userToBeRejectedId string) *models.ErrorJson {
	if errJson := gRepo.CancelTheInvitation(userId, groupId, userToBeRejectedId); errJson != nil {
		return &models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message}
	}
	return nil
}
