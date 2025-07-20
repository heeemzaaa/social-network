package chat

import (
	"fmt"
	"log"

	"social-network/backend/models"
)

// check if the id givven is a member inside the group givven
func (repo *ChatRepository) IsMember(userID, groupID string) (bool, *models.ErrorJson) {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM group_membership WHERE userID = ? AND groupID = ? LIMIT 1)`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to check if the user is a member: ", err)
		return false, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	err = stmt.QueryRow(userID, groupID).Scan(&exists)
	if err != nil {
		log.Println("Error getting the member of the group: ", err)
		return false, &models.ErrorJson{Status: 500, Error: "", Message: fmt.Sprintf("%v", err)}
	}
	return exists, nil
}

// check if the user exist, I will need this a lot hhh
func (repo *ChatRepository) UserExists(userID string) (bool, *models.ErrorJson) {
	var exists bool

	query := `SELECT EXISTS (SELECT 1 FROM users WHERE userID = ? LIMIT 1)`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to check if the user exists: ", err)
		return false, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	err = stmt.QueryRow(userID).Scan(&exists)
	if err != nil {
		log.Println("Error checking if the user exists: ", err)
		return false, &models.ErrorJson{Status: 500, Error: "", Message: fmt.Sprintf("%v", err)}
	}

	return exists, nil
}

// check if the group Id givven exists
func (repo *ChatRepository) GroupExists(groupID string) (bool, *models.ErrorJson) {
	var exists bool

	query := `SELECT EXISTS (SELECT 1 FROM groups WHERE groupID = ? LIMIT 1)`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to check if the group exists: ", err)
		return false, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	err = stmt.QueryRow(groupID).Scan(&exists)
	if err != nil {
		log.Println("Error checking if the group exists: ", err)
		return false, &models.ErrorJson{Status: 500, Error: "", Message: fmt.Sprintf("%v", err)}
	}
	return exists, nil
}
