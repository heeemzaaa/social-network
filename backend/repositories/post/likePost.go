package repositories

import (
	"fmt"

	"github.com/google/uuid"
)

func (r *PostsRepository) HandleLike(postID uuid.UUID, userID uuid.UUID) error {
	var exists bool

	entityType := "post"
	reactionValue := 1
	checkQuery := `
		SELECT EXISTS (
			SELECT 1 FROM reactions 
			WHERE userID = ? AND entityType = ? AND entityID = ?
		)
	`
	err := r.db.QueryRow(checkQuery, userID.String(), entityType, postID.String()).Scan(&exists)
	if err != nil {
		fmt.Println("Error checking existing reaction:", err)
		return err
	}

	if !exists {

		insertQuery := `
		INSERT INTO reactions (reactionID, entityType, entityID, reaction, userID)
		VALUES (?, ?, ?, ?, ?)
		`
		_, err := r.db.Exec(
			insertQuery,
			uuid.New().String(),
			entityType,
			postID.String(),
			reactionValue,
			userID.String(),
		)
		if err != nil {
			fmt.Println("Error inserting reaction:", err)
			return err
		}
	} else {
		updateQuery := `
		UPDATE reactions 
		SET reaction = CASE WHEN reaction = 1 THEN 0 ELSE 1 END
		WHERE userID = ? AND entityType = ? AND entityID = ?
		`
		_, err := r.db.Exec(updateQuery, userID.String(), entityType, postID.String())
		if err != nil {
			fmt.Println("Error deleting reaction:", err)
			return err
		}
	}

	return nil
}
