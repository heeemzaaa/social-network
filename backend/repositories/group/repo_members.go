package group

import (
	"fmt"

	"social-network/backend/models"
)

func (gRepo *GroupRepository) GetGroupMembers(groupId string) ([]models.User, *models.ErrorJson) {
	users := []models.User{}
	query := `

	SELECT  users.userID,
	concat(users.firstName, " ", users.lastName),
	users.nickname, users.avatarPath 
	FROM users INNER JOIN group_membership ON 
	users.userID = group_membership.userID 
	WHERE groupID = ? 
	`

	stmt, err := gRepo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	rows, err := stmt.Query(groupId)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer rows.Close()

	for rows.Next() {
		user := models.User{}
		if errScan := rows.Scan(&user.Id, &user.FullName, &user.Nickname, &user.ImagePath); errScan != nil {
			return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
		}
		users = append(users, user)
	}
	return users, nil
}
