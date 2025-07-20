package chat

import (
	"database/sql"
	"fmt"

	"social-network/backend/models"
)

func (repo *ChatRepository) GetID(sessionID string) (string, *models.ErrorJson) {
	query := `SELECT userID FROM sessions WHERE userID = ?`
	var userID string
	err := repo.db.QueryRow(query, sessionID).Scan(&userID)
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
	row := repo.db.QueryRow(query, token).Scan(&session.UserId, &session.Token, &session.FirstName, &session.LastName)
	if row == sql.ErrNoRows {
		return nil, &models.ErrorJson{Status: 401, Message: " Unauthorized Access"}
	}
	return &session, nil
}

// get the userID from the session , we will need it also
func (repo *ChatRepository) GetUserIdFromSession(sessionID string) (string, *models.ErrorJson) {
	var userID string
	query := `SELECT userID FROM sessions WHERE sessionToken = ?`
	errQuery := repo.db.QueryRow(query, sessionID).Scan(&userID)
	if errQuery != nil {
		return "", &models.ErrorJson{Status: 500, Error: "", Message: fmt.Sprintf("%v", errQuery)}
	}

	return userID, nil
}
