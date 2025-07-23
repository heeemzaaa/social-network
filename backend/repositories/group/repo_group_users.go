package group

import (
	"fmt"

	"social-network/backend/models"
)

// SELECT EXISTS(SELECT 1 FROM users WHERE userID = ?);
func (gRepo *GroupRepository) IsMemberGroup(groupId, userId string) (bool, *models.ErrorJson) { ///// check sender or reciever if already in group
	var exists bool
	query := ` 
		SELECT EXISTS(SELECT 1 FROM  group_membership
		WHERE groupID = ? AND userID = ?);
     `
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return false, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()
	if err = stmt.QueryRow(groupId, userId).Scan(&exists); err != nil {
		return false, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	return exists, nil
}
