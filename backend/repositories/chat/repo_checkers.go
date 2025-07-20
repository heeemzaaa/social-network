package chat

import (
	"fmt"

	"social-network/backend/models"
)

// check if the id givven is a member inside the group givven
func (repo *ChatRepository) IsMember(userID, groupID string) (bool, *models.ErrorJson) {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM group_membership WHERE userID = ? AND groupID = ? LIMIT 1)`
	err := repo.db.QueryRow(query, userID, groupID).Scan(&exists)
	if err != nil {
		return false, &models.ErrorJson{Status: 500, Error: "", Message: fmt.Sprintf("%v", err)}
	}
	return exists, nil
}

// check if the user exist, I will need this a lot hhh
func (repo *ChatRepository) UserExists(userID string) (bool, *models.ErrorJson) {
	var exists bool
	fmt.Println(userID)
	query := `SELECT EXISTS (SELECT 1 FROM users WHERE userID = ? LIMIT 1)`
	err := repo.db.QueryRow(query, userID).Scan(&exists)
	if err != nil {
		return false, &models.ErrorJson{Status: 500, Error: "", Message: fmt.Sprintf("%v", err)}
	}
	return exists, nil
}

// check if the group Id givven exists
func (repo *ChatRepository) GroupExists(groupID string) (bool, *models.ErrorJson) {
	var exists bool

	query := `SELECT EXISTS (SELECT 1 FROM groups WHERE groupID = ? LIMIT 1)`
	err := repo.db.QueryRow(query, groupID).Scan(&exists)
	if err != nil {
		return false, &models.ErrorJson{Status: 500, Error: "", Message: fmt.Sprintf("%v", err)}
	}
	return exists, nil
}
