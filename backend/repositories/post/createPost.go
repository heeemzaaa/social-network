package post

import (
	"fmt"
	"log"

	"social-network/backend/models"
)

func (r *PostsRepository) CreatePost(post *models.Post) (*models.Post, *models.ErrorJson) {
	var post_created models.Post

	query := `
		INSERT INTO posts (postID, userID, content, privacy, image_url)
			VALUES (?, ?, ?, ?, ?)
		RETURNING postID, userID, content, privacy, image_url, createdAt,
	(
        SELECT
            concat (firstName, ' ', lastName)
        FROM
            users
        WHERE
            users.userID = ?
    ) AS fullName,
    (
        SELECT
            nickname
        FROM
            users
        WHERE
            users.userID = ?
    ),(
        SELECT
            avatarPath
        FROM
            users
        WHERE
            users.userID = ?
    )
	`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing the query to create post: ", err)
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()

	err = stmt.QueryRow(post.Id, post.User.Id, post.Content, post.Privacy, post.Img, post.User.Id, post.User.Id, post.User.Id).Scan(
		&post_created.Id,
		&post_created.User.Id,
		&post_created.Content,
		&post_created.Privacy,
		&post_created.Img,
		&post_created.CreatedAt,
		&post_created.User.FullName,
		&post_created.User.Nickname,
		&post_created.User.ImagePath,
	)
	if err != nil {
		log.Println("Error inserting post: ", err)
		return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
	}

	if post.Privacy == "private" && len(post.SelectedUsers) > 0 {
		for _, followerID := range post.SelectedUsers {
			query := `INSERT INTO post_access (postID, userID)
				 VALUES (?, ?)`

			stmt, err := r.db.Prepare(query)
			if err != nil {
				log.Println("Error preparing the query to insert the allowed users: ", err)
				return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
			}
			defer stmt.Close()

			_, err = stmt.Exec(post.Id, followerID)
			if err != nil {
				log.Println("Error scanning the allowed users: ", err)
				return nil, &models.ErrorJson{Status: 500, Error: fmt.Sprintf("%v", err)}
			}
		}
	}
	return &post_created, nil
}
