package profile

import (
	"fmt"
	"log"

	"social-network/backend/models"
)

// here we accept all the requests where the userID
// change its visibility to public only
func (repo *ProfileRepository) AcceptAllrequest(userID string) *models.ErrorJson {
	var users []models.User
	query := `SELECT requestorID FROM follow_requests WHERE userID = ?`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to get the profile data: ", err)
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	rows, err := stmt.Query(userID)
	if err != nil {
		log.Println("Error accepting all requests: ", err)
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id)
		if err != nil {
			log.Println("Error scanning the accepted request: ", err)
			return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error in the whole process of scan => in accept all requests: ", err)
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	for _, user := range users {
		query := `INSERT INTO followers (userID, followerID) VALUES(?,?)`

		stmt, err := repo.db.Prepare(query)
		if err != nil {
			log.Println("Error preparing the query to add the user to the followers: ", err)
			return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
		}

		_, err = stmt.Exec(userID, user.Id)
		if err != nil {
			log.Println("Error inserting the new followers to the table: ", err)
			return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
		}
	}

	query = `DELETE FROM follow_requests WHERE userID = ?`

	stmt, err = repo.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to delete the request: ", err)
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	_, err = stmt.Exec(userID)
	if err != nil {
		log.Println("Error deleting the user from the request table: ", err)
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	return nil
}

// here I will update the visibility of a private status
func (repo *ProfileRepository) ToPublicAccount(userID string) *models.ErrorJson {
	query := `UPDATE users SET visibility = ? WHERE userID = ?`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to change visibility to public: ", err)
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	_, err = stmt.Exec("public", userID)
	if err != nil {
		log.Println("Error updating the visibility to public: ")
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	return nil
}

// change the visibility to private
func (repo *ProfileRepository) ToPrivateAccount(userID string) *models.ErrorJson {
	query := `UPDATE users SET visibility = ? WHERE userID = ?`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to change visibility to private: ", err)
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	_, err = stmt.Exec("private", userID)
	if err != nil {
		log.Println("Error updating the visibility to private: ", err)
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	return nil
}
