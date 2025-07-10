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
	SELECT "true"
    FROM sessions
    WHERE sessionToken = ?
	limit 1`
	if err := repo.db.QueryRow(query, token).Scan(user_data.IsLoggedIn); err != nil {
		if err == sql.ErrNoRows {
			return nil, &models.ErrorJson{Status: 401, Message: "ERROR!! Unauthorized Access"}
		}
	}
	return user_data, nil
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
	fmt.Printf("login: %v\n", login)
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
			Error:  "invalid credentials!",
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

func (appRep *AuthRepository) GetItem(typ string, field string, value string) ([]any, bool, *models.ErrorJson) {
	data := make([]any, 0)
	query := fmt.Sprintf(`SELECT %v FROM %v WHERE %v=?`, field, typ, field)
	stmt, err := appRep.db.Prepare(query)
	if err != nil {
		return nil, false, models.NewErrorJson(500, "Internal Server error", nil)
	}
	defer stmt.Close()
	rows, err := stmt.Query(value)
	if err != nil {
		return nil, false, models.NewErrorJson(500, "Internal Server error", nil)
	}
	for rows.Next() {
		var row any
		rows.Scan(&row)
		data = append(data, row)
	}

	defer rows.Close()

	if len(data) != 0 {
		return data, true, nil
	}
	return nil, false, nil
}

func (repo *AuthRepository) GetUserSessionByUserId(user_id string) (*models.Session, *models.ErrorJson) {
	session := &models.Session{}
	query := `SELECT * FROM sessions WHERE userID = ?`
	row := repo.db.QueryRow(query, user_id)
	err := row.Scan(&session.Id, &session.UserId, &session.Token, &session.ExpDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return session, nil
}

func (repo *AuthRepository) UpdateSession(session *models.Session, new_session *models.Session) *models.ErrorJson {
	query := `UPDATE sessions SET sessionToken = ? , expiresAt = ? where sessionToken= ?`
	_, err := repo.db.Exec(query, new_session.Token, new_session.ExpDate, session.Token)
	if err != nil {
		return models.NewErrorJson(500, fmt.Sprintf("%v", err), nil)
	}
	return nil
}

func (repo *AuthRepository) GetSessionbyTokenEnsureAuth(token string) (*models.Session, *models.ErrorJson) {
	session := models.Session{}
	query := `SELECT sessions.userID, sessions.sessionToken , sessions.expiresAt, users.nickname 
	FROM sessions INNER JOIN users ON users.userID = sessions.userID
	WHERE sessionToken = ?`
	row := repo.db.QueryRow(query, token).Scan(&session.UserId, &session.Token, &session.ExpDate, &session.Username)
	if row == sql.ErrNoRows {
		return nil, &models.ErrorJson{Status: 401, Message: " Unauthorized Access"}
	}
	return &session, nil
}
