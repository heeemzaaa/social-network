package group

import (
	"fmt"

	"social-network/backend/models"
)

func (gRepo *GroupRepository) Approve(groupId string, userIdToBeAdded string) *models.ErrorJson {
	query := `
	INSERT INTO group_membership (groupID,userID) 
	VALUES (?,?)
	`
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	_, err = stmt.Exec(groupId, userIdToBeAdded)
	if err != nil {
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	if errJson := gRepo.RequestToCancel(userIdToBeAdded, groupId); errJson != nil {
		return &models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message}
	}
	return nil
}

func (gRepo *GroupRepository) Decline(groupId string, userToBeRejected string) *models.ErrorJson {
	if errJson := gRepo.RequestToCancel(userToBeRejected, groupId); errJson != nil {
		return &models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message}
	}
	return nil
}
