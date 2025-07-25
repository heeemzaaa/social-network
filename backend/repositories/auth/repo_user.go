package auth

import (
	"database/sql"
	"fmt"

	models "social-network/backend/models"
)

// DB wash create user hya register hnayaa wlla hadak service aykllf ???  ;(
// hadshi taaafh
// y9dr ay wa7d ydiiruuu

func (repo *AuthRepository) CreateUser(user *models.User) *models.ErrorJson {
	fmt.Println("user in repo:" , user.AboutMe)
	query := `INSERT INTO users (userID, email, firstName, lastName, password, birthDate, nickname, avatarPath, aboutMe, visibility) VALUES (?,?,?,?,?,?,?,?,?,?)`
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()
	if _, err = stmt.Exec(user.Id, user.Email, user.FirstName, user.LastName,
		user.Password, user.BirthDate, user.Nickname, user.ImagePath, user.AboutMe, "private"); err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}

// TO GET THE USERS

// chosen_field ( it may be the nickname or the email )
// the query must not include the password entered by the user
func (appRep *AuthRepository) GetUser(login *models.Login) (*models.User, *models.ErrorJson) {
	user := models.NewUser()
	query := `SELECT userID, nickname, password 
	FROM users where nickname=? OR email =? `
	stmt, err := appRep.db.Prepare(query)
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

// get the username from the userId
func (appRep *AuthRepository) GetUserNameById(user_id string) (string, *models.ErrorJson) {
	var username string
	query := `SELECT nickname FROM users WHERE userID = ?`
	err := appRep.db.QueryRow(query, user_id).Scan(&username)
	if err != nil {
		return "", &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return username, nil
}

func (appRepo *AuthRepository) UserExists(id int) (bool, *models.ErrorJson) {
	var exists bool
	query := ` SELECT EXISTS(SELECT 1 FROM users WHERE userID = ?);`
	stmt, err := appRepo.db.Prepare(query)
	if err != nil {
		return false, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&exists)
	if err == sql.ErrNoRows {
		return false, &models.ErrorJson{Status: 400, Message: "user not found"}
	}
	return exists, nil
}
