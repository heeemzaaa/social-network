package auth

import (
	"database/sql"
	"fmt"

	"social-network/backend/models"
)

func (appRep *AuthRepository) CreateUserSession(session *models.Session, user *models.User) *models.ErrorJson {
	query := `INSERT INTO sessions (userID, sessionToken) VALUES (?,?)`
	_, err := appRep.db.Exec(query, user.Id, session.Token)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}

// get the session by the user id or the user nickname !!

func (appRepo *AuthRepository) GetSessionbyTokenEnsureAuth(token string) (*models.Session, *models.ErrorJson) {
	session := models.Session{}
	query := `SELECT sessions.userID, sessions.sessionToken , users.nickname 
	FROM sessions INNER JOIN users ON users.userID = sessions.userID
	WHERE sessionToken = ?`
	row := appRepo.db.QueryRow(query, token).Scan(&session.UserId, &session.Token, &session.Username)
	if row == sql.ErrNoRows {
		return nil, &models.ErrorJson{Status: 401, Message: " Unauthorized Access"}
	}
	return &session, nil
}

func (appRepo *AuthRepository) HasValidToken(token string) (bool, *models.Session) {
	session := models.NewSession()
	query := `SELECT userID, sessionToken , expiresAt FROM sessions WHERE sessionToken = ?`
	row := appRepo.db.QueryRow(query, token).Scan(&session.UserId, &session.Token)

	if row == sql.ErrNoRows {
		return false, nil
	}

	if (session != &models.Session{}) {
		return true, session
	}
	return false, nil
}

func (appRep *AuthRepository) GetUserSessionByUserId(user_id string) (*models.Session, *models.ErrorJson) {
	session := &models.Session{}
	query := `SELECT * FROM sessions WHERE userID = ?`
	row := appRep.db.QueryRow(query, user_id)
	err := row.Scan(&session.Id, &session.UserId, &session.Token)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return session, nil
}



func (appRep *AuthRepository) DeleteSession(session models.Session) *models.ErrorJson {
	query := `DELETE FROM sessions WHERE sessionToken = ?`
	_, err := appRep.db.Exec(query, session.Token)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}
