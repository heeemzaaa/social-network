package group

import (
	"fmt"

	"social-network/backend/models"
	"social-network/backend/utils"
)

//  This is will be for design and then implementing the requests in the correct way
//  admin of the group has the right to invite its followers
// BUT we also add the admin as a memeber of the group
// so the whole list of members will be able to invite its followers
// RECEIVER ID is always the admin of the group!!

func (gRepo *GroupRepository) RequestToJoin(userId, groupId string) *models.ErrorJson {
	requestId := utils.NewUUID()
	query := `
	INSERT INTO
    group_requests (
        requestID,
        senderID,
        receiverID,
        groupID,
        typeRequest
    )
	VALUES
    (?, ?, (SELECT groups.groupCreatorID FROM groups WHERE groups.groupID = ?), ?, ?);
	`
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v 1", err)}
	}
	defer stmt.Close()

	_, err = stmt.Exec(requestId, userId, groupId, groupId, "join-request")
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v 1", err)}
	}
	return nil
}

func (gRepo *GroupRepository) RequestToCancel(userId, groupId string) *models.ErrorJson {
	query := `
	DELETE FROM group_requests
	WHERE senderID = ? AND groupID = ? AND receiverID = (SELECT groups.groupCreatorID FROM groups WHERE groups.groupID =? )
	AND typeRequest = ?;
	`
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v 1", err)}
	}
	defer stmt.Close()

	_, err = stmt.Exec(userId, groupId, groupId, "join-request")
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v 1", err)}
	}
	return nil
}
