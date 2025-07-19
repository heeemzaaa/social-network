package profile

import (
	"fmt"
	"log"

	"social-network/backend/models"
)

// here I will insert the follow request if the user has a public account
func (repo *ProfileRepository) FollowDone(userID string, authUserID string) *models.ErrorJson {
	query := `INSERT INTO followers (userID, followerID) VALUES(?,?) `

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to do the follow action: ", err)
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	_, err = stmt.Exec(userID, authUserID)
	if err != nil {
		log.Println("Error completing the follow action: ", err)
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	return nil
}

// here I will just insert the request into the table of the followrequests
func (repo *ProfileRepository) FollowPrivate(userID string, authUserID string) *models.ErrorJson {
	query := `INSERT INTO follow_requests (userID, requestorID) VALUES(?, ?) `

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to do the follow private action: ", err)
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	_, err = stmt.Exec(userID, authUserID)
	if err != nil {
		log.Println("Error adding the request follow to the table: ", err)
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	return nil
}

// here we unfollow the user chosen , we delete it from the table of followers
func (repo *ProfileRepository) Unfollow(userID string, authUserID string) *models.ErrorJson {
	query := `DELETE FROM followers WHERE userID = ? AND followerID = ?`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to do the unfollow action: ", err)
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	_, err = stmt.Exec(userID, authUserID)
	if err != nil {
		log.Println("Error executing the unfollow request: ", err)
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	return nil
}
