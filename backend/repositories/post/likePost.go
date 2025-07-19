package repositories

import (
	"fmt"

	"github.com/google/uuid"
)

func (r *PostsRepository) HandleLike(postID uuid.UUID, userID uuid.UUID) (bool, int, error) {
	var exists bool
	entityType := "post"
	reactionValue := 1

	// Check if reaction already exists
	checkQuery := `
		SELECT EXISTS (
			SELECT 1 FROM reactions 
			WHERE userID = ? AND entityType = ? AND entityID = ?
		)
	`
	err := r.db.QueryRow(checkQuery, userID.String(), entityType, postID.String()).Scan(&exists)
	if err != nil {
		fmt.Println("Error checking existing reaction:", err)
		return false, 0, err
	}

	if !exists {
		// Insert a new like
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
			return false, 0, err
		}
	} else {
		// Toggle existing like
		updateQuery := `
			UPDATE reactions 
			SET reaction = CASE WHEN reaction = 1 THEN 0 ELSE 1 END
			WHERE userID = ? AND entityType = ? AND entityID = ?
		`
		_, err := r.db.Exec(updateQuery, userID.String(), entityType, postID.String())
		if err != nil {
			fmt.Println("Error toggling reaction:", err)
			return false, 0, err
		}
	}

	var liked bool
	likeCheckQuery := `
		SELECT reaction = 1 FROM reactions
		WHERE userID = ? AND entityType = ? AND entityID = ?
	`
	err = r.db.QueryRow(likeCheckQuery, userID.String(), entityType, postID.String()).Scan(&liked)
	if err != nil {
		fmt.Println("Error fetching liked state:", err)
		return false, 0, err
	}


	var totalLikes int
	countQuery := `
		SELECT COUNT(*) FROM reactions
		WHERE entityType = ? AND entityID = ? AND reaction = 1
	`
	err = r.db.QueryRow(countQuery, entityType, postID.String()).Scan(&totalLikes)
	if err != nil {
		fmt.Println("Error fetching total likes:", err)
		return false, 0, err
	}

	return liked, totalLikes, nil
}
