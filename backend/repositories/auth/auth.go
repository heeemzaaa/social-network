package auth

import (
	"database/sql"
	"fmt"

	"social-network/backend/models"
)

type AuthRepository struct {
	db *sql.DB
}

// NewPostRepository creates a new repository
func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (appRep *AuthRepository) IsLoggedInUser(token string) (*models.IsLoggedIn, *models.ErrorJson) {
	user_data := &models.IsLoggedIn{}
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

func (repo *AuthRepository) GetSession(token string) (*models.Session, *models.ErrorJson) {
	session := models.Session{}
	query := `SELECT sessions.userID, sessions.sessionToken
	FROM sessions RIGHT JOIN users ON users.userID = sessions.userID
	WHERE sessionToken = ?`

	row := repo.db.QueryRow(query, token).Scan(&session.UserId, &session.Token)
	if row == sql.ErrNoRows {
		return nil, &models.ErrorJson{Status: 401, Message: " Unauthorized Access"}
	}
	return &session, nil
}

func (repo *AuthRepository) DeleteSession(session models.Session) *models.ErrorJson {
	query := `DELETE FROM sessions WHERE sessionToken = ?`
	_, err := repo.db.Exec(query, session.Token)
	if err != nil {
		return models.NewErrorJson(500, fmt.Sprintf("%v", err), nil)
	}
	return nil
}
