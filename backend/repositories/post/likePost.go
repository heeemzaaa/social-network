package repositories

import (
	"fmt"
	"log"

	"social-network/backend/models"

	"github.com/google/uuid"
)

func (r *PostsRepository) HandleLike(postID string, userID string) (bool, int, *models.ErrorJson) {
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

	stmt, err := r.db.Prepare(checkQuery)
	if err != nil {
		log.Println("Error preparing the query to handle like: ", err)
		return false, 0, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	err = stmt.QueryRow(userID, entityType, postID).Scan(&exists)
	if err != nil {
		log.Println("Error checking existing reaction:", err)
		return false, 0, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	if !exists {
		// Insert a new like
		insertQuery := `
			INSERT INTO reactions (reactionID, entityType, entityID, reaction, userID)
			VALUES (?, ?, ?, ?, ?)
		`

		stmt, err := r.db.Prepare(insertQuery)
		if err != nil {
			log.Println("Error preparing the query to handle like: ", err)
			return false, 0, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
		}
		defer stmt.Close()

		_, err = stmt.Exec(
			uuid.New().String(),
			entityType,
			postID,
			reactionValue,
			userID,
		)
		if err != nil {
			log.Println("Error inserting reaction:", err)
			return false, 0, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
		}
	} else {
		// Toggle existing like
		updateQuery := `
			UPDATE reactions 
			SET reaction = CASE WHEN reaction = 1 THEN 0 ELSE 1 END
			WHERE userID = ? AND entityType = ? AND entityID = ?
		`

		stmt, err := r.db.Prepare(updateQuery)
		if err != nil {
			log.Println("Error preparing the query to update like: ", err)
			return false, 0, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
		}
		defer stmt.Close()

		_, err = stmt.Exec(userID, entityType, postID)
		if err != nil {
			log.Println("Error toggling reaction:", err)
			return false, 0, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
		}
	}

	var liked bool
	likeCheckQuery := `
		SELECT reaction = 1 FROM reactions
		WHERE userID = ? AND entityType = ? AND entityID = ?
	`

	stmt, err = r.db.Prepare(likeCheckQuery)
	if err != nil {
		log.Println("Error preparing the query to check like: ", err)
		return false, 0, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	err = stmt.QueryRow(userID, entityType, postID).Scan(&liked)
	if err != nil {
		log.Println("Error fetching liked state: ", err)
		return false, 0, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	var totalLikes int
	countQuery := `
		SELECT COUNT(*) FROM reactions
		WHERE entityType = ? AND entityID = ? AND reaction = 1
	`

	stmt, err = r.db.Prepare(countQuery)
	if err != nil {
		log.Println("Error preparing the query to count like: ", err)
		return false, 0, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	err = stmt.QueryRow(entityType, postID).Scan(&totalLikes)
	if err != nil {
		log.Println("Error fetching total likes:", err)
		return false, 0, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	return liked, totalLikes, nil
}
