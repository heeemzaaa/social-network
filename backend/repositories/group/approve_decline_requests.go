package group

import (
	"fmt"

	"social-network/backend/models"
)

func (gRepo *GroupRepository) Approve(groupId string, userToBeAdded *models.User) *models.ErrorJson {
	query := `
	INSERT INTO group_membership (groupID,userID) 
	VALUES (?,?)
	`
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	_, err = stmt.Exec(groupId, userToBeAdded.Id)
	if err != nil {
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	if errJson := gRepo.RequestToCancel(userToBeAdded.Id, groupId); errJson != nil {
		return &models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message}
	}
	return nil
}

func (gRepo *GroupRepository) Decline(groupId string, userToBeRejected *models.User) *models.ErrorJson {
	if errJson := gRepo.RequestToCancel(userToBeRejected.Id, groupId); errJson != nil {
		return &models.ErrorJson{Status: errJson.Status, Error: errJson.Error, Message: errJson.Message}
	}
	return nil
}
