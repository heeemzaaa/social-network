package group

import (
	"fmt"

	"social-network/backend/models"

	"github.com/mattn/go-sqlite3"
)

// in the group membership table we have a composite key
// (groupID, userID ) had l combinaison khssha tkun unique
// to detect if a user aleady a part of group and returns a 403
func (repo *GroupRepository) JoinGroup(group *models.Group, userId string) *models.ErrorJson {
	query := `
	INSERT INTO group_membership (groupID, userID)
	VALUES (?, ?)
	`
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		if sqlite3Err, ok := err.(sqlite3.Error); ok {
			if sqlite3Err.Code == sqlite3.ErrConstraint {
				return &models.ErrorJson{Status: 403, Message: "user already joined the group!"}
			}
		}

		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	_, err = stmt.Exec(group.GroupId, userId)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}



//  group title and description and up to 20 memebers joined the group and the number of the members
// the number of posts created !!
func (repo *GroupRepository) GetGroupDetails(groupId string) (*models.Group, *models.ErrorJson) {
	var groupDetails *models.Group
     



	return groupDetails, nil
}
