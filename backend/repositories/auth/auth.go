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

func (repo *AuthRepository) CreateUser(user *models.User) *models.ErrorJson {
	query := `INSERT INTO users (userID, email, firstName, lastName, password, birthDate, nickname, avatarPath, aboutMe, visibility) VALUES (?,?,?,?,?,?,?,?,?,?)`
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()
	if _, err = stmt.Exec(user.Id, user.Email, user.FirstName, user.LastName, user.Password, user.BirthDate, user.Nickname, user.ProfileImage, user.AboutMe, "private"); err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}

func (repo *AuthRepository) IsLoggedInUser(token string) (*models.IsLoggedIn, *models.ErrorJson) {
	user_data := &models.IsLoggedIn{}
	query := `
	SELECT users.userID, users.nickname
    FROM users INNER JOIN sessions ON users.userID = sessions.userID 
    WHERE sessionToken = ? `
	if err := repo.db.QueryRow(query, token).Scan(&user_data.Id, &user_data.Nickname); err != nil {
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

func (repo *AuthRepository) GetUser(login *models.Login) (*models.User, *models.ErrorJson) {
	user := models.NewUser()
	query := `SELECT userID, nickname, password 
	FROM users where nickname=? OR email = ? `
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()
	row := stmt.QueryRow(login.LoginField, login.LoginField)
	err = row.Scan(&user.Id, &user.Nickname, &user.Password)
	if err == sql.ErrNoRows {
		return nil, &models.ErrorJson{
			Status: 401,
			Message: models.Login{
				LoginField: "invalid login credentials!",
				Password:   "invalid login credentials!",
			},
		}
	}
	return user, nil
}

func (appRep *AuthRepository) CreateUserSession(session *models.Session, user *models.User) *models.ErrorJson {
	query := `INSERT INTO sessions (userID, sessionToken, expiresAt) VALUES (?,?,?)`
	_, err := appRep.db.Exec(query, user.Id, session.Token, session.ExpDate)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}
