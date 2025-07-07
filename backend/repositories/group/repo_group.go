package group

import (
	"fmt"

	"social-network/backend/models"
)

func (repo *GroupRepository) JoinGroup(group *models.Group, userId string) *models.ErrorJson {
	query := `
	INSERT INTO group_membership (groupID, userID)
	VALUES (?, ?)
	`
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	_, err = stmt.Exec(group.GroupId, userId)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}


func (repo *GroupRepository) GetGroupDetails(group *models.Group, userId string) *models.ErrorJson {
   query := ``
}