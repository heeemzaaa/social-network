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

func (repo *AuthRepository) GetSessionbyTokenEnsureAuth(token string) (*models.Session, *models.ErrorJson) {
	session := models.Session{}

	query := `SELECT sessions.userID, sessions.sessionToken , users.nickname 
	FROM sessions INNER JOIN users ON users.userID = sessions.userID
	WHERE sessionToken = ?`

	row := repo.db.QueryRow(query, token).Scan(&session.UserId, &session.Token, &session.Username)
	if row == sql.ErrNoRows {
		return nil, &models.ErrorJson{Status: 401, Message: " Unauthorized Access"}
	}
	return &session, nil
}

func (repo *AuthRepository) DeleteSession(session models.Session) *models.ErrorJson {
	query := `DELETE FROM sessions WHERE sessionToken = ?`
	_, err := repo.db.Exec(query, session.Token)
	if err != nil {
		return models.NewErrorJson(500, fmt.Sprintf("%v", err))
	}
	return nil
}
