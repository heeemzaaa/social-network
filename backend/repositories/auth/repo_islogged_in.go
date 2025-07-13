package auth

import (
	"database/sql"

	"social-network/backend/models"
)

func (appRep *AuthRepository) IsLoggedInUser(token string) (*models.UserData, *models.ErrorJson) {
	user_data := &models.UserData{}
	query := `
	SELECT users.userID, users.nickname
    FROM users INNER JOIN sessions ON users.userID = sessions.userID 
    WHERE sessionToken = ? `
	if err := appRep.db.QueryRow(query, token).Scan(&user_data.Id, &user_data.Nickname); err != nil {
		if err == sql.ErrNoRows {
			return nil, &models.ErrorJson{Status: 401, Message: "ERROR!! Unauthorized Access"}
		}
	}
	return user_data, nil
}
