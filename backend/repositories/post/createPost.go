package repositories

import (
	"fmt"
	"log"

	"social-network/backend/models"
)

func (r *PostsRepository) CreatePost(post *models.Post) *models.ErrorJson {
	query := `INSERT INTO posts (postID, userID, content, privacy, image_url)
		 VALUES (?, ?, ?, ?, ?)`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to create post: ", err)
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	_, err = stmt.Exec(post.Id, post.User.Id, post.Content, post.Privacy, post.Img)
	if err != nil {
		log.Println("Error inserting post: ", err)
		return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	fullName, errName := r.GetUserName(post.User.Id)
	if errName != nil {
		return &models.ErrorJson{Status: errName.Status, Error: errName.Error}
	}
	post.User.FirstName = fullName
	if post.Privacy == "private" && len(post.SelectedUsers) > 0 {
		for _, followerID := range post.SelectedUsers {
			query := `INSERT INTO post_access (postID, userID)
				 VALUES (?, ?)`

			stmt, err := r.db.Prepare(query)
			if err != nil {
				log.Println("Error preparing the query to insert the allowed users: ", err)
				return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
			}
			defer stmt.Close()

			_, err = stmt.Exec(post.Id, followerID)
			if err != nil {
				log.Println("Error scanning the allowed users: ", err)
				return &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
			}
		}
	}
	return nil
}
