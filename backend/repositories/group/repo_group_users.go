package group

import (
	"fmt"

	"social-network/backend/models"
)

func (gRepo *GroupRepository) IsMemberGroup(groupId, userId string) (bool, *models.ErrorJson) {
	query := ` 
		SELECT 1 FROM  group_membership
		WHERE groupID = ? AND userID = ?;
     `
	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return false, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v 1", err)}
	}
	defer stmt.Close()
	err = stmt.QueryRow(groupId, userId).Scan()
}
