package repositories

import (
	"fmt"
	"log"

	"social-network/backend/models"
)

func (r *PostsRepository) GetUserName(userID string) (string, *models.ErrorJson) {
	var firstName, lastName string
	query := `SELECT firstName , lastName FROM users WHERE userID = ?`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to get the full name: ", err)
		return "", &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	err = stmt.QueryRow(userID).Scan(&firstName, &lastName)
	if err != nil {
		log.Println("Error selecting the full name: ", err)
		return "", &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	return firstName + " " + lastName, nil
}
