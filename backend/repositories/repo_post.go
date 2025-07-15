package repositories

import (
	"database/sql"
	"fmt"

	"social-network/backend/models"
)

type PostsRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostsRepository {
	return &PostsRepository{db: db}
}

func (r *PostsRepository) CreatePost(post *models.Post) error {
	fmt.Println("CREATING POST")
	fmt.Println("the post privacy is ", post.Privacy)
	_, err := r.db.Exec(
		`INSERT INTO posts (postID, userID, content, privacy, image_url) VALUES (?, ?, ?, ?,?)`,
		post.ID, post.UserID, post.Content, post.Privacy, post.Img,
	)
	fmt.Println("SQL : ", err)
	return err
}

func (r *PostsRepository) GetAllPosts() ([]models.Post, error) {
	rows, err := r.db.Query(`SELECT id, user_id, content FROM posts`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var p models.Post
		if err := rows.Scan(&p.ID, &p.UserID, &p.Content); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	return posts, nil
}

func (r *PostsRepository) GetPostByID(postID string) (models.Post, error) {
	var p models.Post
	err := r.db.QueryRow(`SELECT id, user_id, content FROM posts WHERE id = ?`, postID).
		Scan(&p.ID, &p.UserID, &p.Content)
	return p, err
}

func (r *PostsRepository) DeletePost(postID string) error {
	_, err := r.db.Exec(`DELETE FROM posts WHERE id = ?`, postID)
	return err
}
