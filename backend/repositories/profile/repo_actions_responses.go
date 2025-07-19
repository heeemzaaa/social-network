package profile

import (
	"fmt"
	"log"
	"social-network/backend/models"
)

// delete the request works in both cases , accept and decline
func (repo *ProfileRepository) DeleteRequest(userID string, authUserID string) *models.ErrorJson {
	query := `DELETE FROM follow_requests WHERE userID = ? AND requestorID = ?`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to delete the request: ", err)
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	_, err = stmt.Exec(userID, authUserID)
	if err != nil {
		log.Println("Error executing the query to delete the request:", err)
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	return nil
}


// if the user accepted the follow request
// I will delete it from the table and add him to the following list
func (repo *ProfileRepository) AcceptedRequest(userID string, authUserID string) *models.ErrorJson {
	err := repo.DeleteRequest(userID, authUserID)
	if err != nil {
		return &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	err = repo.FollowDone(userID, authUserID)
	if err != nil {
		return &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	return nil
}

// here I will just delete it from the table of the follow requests
func (repo *ProfileRepository) RejectedRequest(userID string, authUserID string) *models.ErrorJson {
	err := repo.DeleteRequest(userID, authUserID)
	if err != nil {
		return &models.ErrorJson{Status: err.Status, Error: err.Error}
	}

	return nil
}
