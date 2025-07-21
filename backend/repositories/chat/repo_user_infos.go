package chat

import (
	"database/sql"
	"fmt"
	"log"

	"social-network/backend/models"
)

func (repo *ChatRepository) GetID(sessionID string) (string, *models.ErrorJson) {
	var userID string
	query := `SELECT userID FROM sessions WHERE userID = ?`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to get the userID: ", err)
		return "", &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	err = stmt.QueryRow(sessionID).Scan(&userID)
	if err != nil {
		return "", &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return userID, nil
}

// here I will get the session Id by the token given by the browser to check the auth
func (repo *ChatRepository) GetSessionbyTokenEnsureAuth(token string) (*models.Session, *models.ErrorJson) {
	session := models.Session{}

	query := `SELECT sessions.userID, sessions.sessionToken , users.firstName, users.lastName 
	FROM sessions INNER JOIN users ON users.userID = sessions.userID
	WHERE sessionToken = ?`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to get session by token: ", err)
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	row := stmt.QueryRow(token).Scan(&session.UserId, &session.Token, &session.FirstName, &session.LastName)
	if row == sql.ErrNoRows {
		log.Println("there is no token !")
		return nil, &models.ErrorJson{Status: 401, Message: " Unauthorized Access"}
	}
	return &session, nil
}

// get the userID from the session , we will need it also
func (repo *ChatRepository) GetUserIdFromSession(sessionID string) (string, *models.ErrorJson) {
	var userID string
	query := `SELECT userID FROM sessions WHERE sessionToken = ?`
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to get the id from session: ", err)
		return "", &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	errQuery := stmt.QueryRow(sessionID).Scan(&userID)
	if errQuery != nil {
		log.Println("Error getting the user id: ", err)
		return "", &models.ErrorJson{Status: 500, Error: "", Message: fmt.Sprintf("%v", errQuery)}
	}

	return userID, nil
}
