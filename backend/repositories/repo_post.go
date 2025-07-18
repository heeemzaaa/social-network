package repositories

import (
	"database/sql"
	"fmt"

	"social-network/backend/models"

	"github.com/google/uuid"
)

type PostsRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostsRepository {
	return &PostsRepository{db: db}
}

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

func (r *PostsRepository) GetAllPosts(userID uuid.UUID) ([]models.Post, error) {
	query := `
SELECT DISTINCT 
    p.postID, p.userID, p.content, p.createdAt, p.privacy, p.image_url,
    u.nickname
    FROM posts p
    JOIN users u ON p.userID = u.userID
    LEFT JOIN post_access pa ON p.postID = pa.postID
	LEFT JOIN followers f ON p.userID = f.userID  -- the author being followed
WHERE
    p.privacy = 'public'
    OR p.userID = ? -- author's own posts
    OR (p.privacy = 'private' AND pa.userID = ?)
    OR (p.privacy = 'almost private' AND f.followerID = ?)
ORDER BY p.createdAt DESC;`

	rows, err := r.db.Query(query, userID.String(), userID.String(), userID.String())
	if err != nil {
		fmt.Println("error in executing", err)
		return nil, err
	}
	defer rows.Close()
	var posts []models.Post
	for rows.Next() {
		var p models.Post
		if err := rows.Scan(&p.Id, &p.User.Id, &p.Content, &p.CreatedAt, &p.Privacy, &p.Img, &p.User.Nickname); err != nil {
			fmt.Println("err  in scaning", err)
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}

func (r *PostsRepository) GetPostByID(postID string) (models.Post, error) {
	var p models.Post
	err := r.db.QueryRow(`SELECT id, user_id, content FROM posts WHERE id = ?`, postID).
		Scan(&p.Id, &p.User.Id, &p.Content)
	return p, err
}
