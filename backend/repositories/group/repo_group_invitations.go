package group

import (
	"fmt"

	"social-network/backend/models"
	"social-network/backend/utils"
)

func (gRepo *GroupRepository) InviteToJoin(userId, groupId string, userToInvite models.User) *models.ErrorJson {
	invitationID := utils.NewUUID()
	query := `
	INSERT INTO group_requests (requestID, senderID, receiverID, groupID, typeRequest)
	VALUES (?,?,?,?,?)
	`
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v 1", err)}
	}
	defer stmt.Close()

	_, err = stmt.Exec(invitationID, userId, userToInvite.Id, groupId, "invitation-request")
	if err != nil {
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v 1", err)}
	}

	return nil
}

func (gRepo *GroupRepository) CancelTheInvitation(userId, groupId string, invitedUser *models.User) *models.ErrorJson {
	query := `
	DELETE FROM group_requests WHERE 
	senderID = ? AND receiverID = ? AND groupID = ? AND typeRequest = ? 
	`
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v 1", err)}
	}
	defer stmt.Close()

	res, err := stmt.Exec(userId, invitedUser.Id, groupId, "invitation-request")
	if err != nil {
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v 1", err)}
	}
	// it must be here :)
	if count, _ := res.RowsAffected(); count == 0 {
		return &models.ErrorJson{Status: 404, Error: "Invitation not found"}
	}

	return nil
}

func (gRepo *GroupRepository) GetInvitations() ([]models.User, *models.ErrorJson) {
	query := `
	
	
	`
	fmt.Printf("query: %v\n", query)
	return nil, nil
}
