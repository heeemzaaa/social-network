package repositories

import (
	"fmt"

	"social-network/backend/models"
)

func (r *PostsRepository) CreatePost(post *models.Post) error {
	_, err := r.db.Exec(
		`INSERT INTO posts (postID, userID, content, privacy, image_url)
		 VALUES (?, ?, ?, ?, ?)`,
		post.Id, post.User.Id, post.Content, post.Privacy, post.Img,
	)
	if err != nil {
		fmt.Println("Error inserting post:", err)
		return err
	}
	if post.Privacy == "private" && len(post.SelectedUsers) > 0 {
		for _, followerID := range post.SelectedUsers {
			_, err := r.db.Exec(
				`INSERT INTO post_access (postID, userID)
				 VALUES (?, ?)`,
				post.Id, followerID,
			)
			if err != nil {
				fmt.Println("Error inserting post access:", err)
				return err
			}
		}
	}
	return nil
}
