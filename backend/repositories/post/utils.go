package repositories

import (
	"database/sql"
)

func (r *PostsRepository)GetUserName(userID string) (string, error) {
	var firstName, lastName string
	query := `SELECT firstName , lastName FROM users WHERE userID = ?`

	err := r.db.QueryRow(query, userID).Scan(&firstName, &lastName)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
	}
	return firstName + " " + lastName, nil
}
