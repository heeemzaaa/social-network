package repositories

import (
	"fmt"

	"social-network/backend/models"

	"github.com/google/uuid"
)

func (r *PostsRepository) GetAllPosts(userID uuid.UUID) ([]models.Post, error) {
	query := `
SELECT DISTINCT p.postID,p.userID, p.content, p.createdAt, p.privacy, p.image_url, u.firstName, u.lastName,
    COUNT(DISTINCT r1.reactionID) AS total_likes,
    CASE WHEN r2.reaction = 1 THEN 1 ELSE 0 END AS liked
FROM posts p
JOIN users u ON p.userID = u.userID
LEFT JOIN post_access pa ON p.postID = pa.postID
LEFT JOIN followers f ON p.userID = f.userID
LEFT JOIN reactions r1 ON r1.entityID = p.postID AND r1.entityType = 'post' AND r1.reaction = 1
LEFT JOIN reactions r2 ON r2.entityID = p.postID AND r2.entityType = 'post' AND r2.userID = ?  -- <== current user
WHERE
    p.privacy = 'public'
    OR p.userID = ?                    -- author's own posts
    OR (p.privacy = 'private' AND pa.userID = ?)
    OR (p.privacy = 'almost private' AND f.followerID = ?)
GROUP BY
    p.postID, p.userID, p.content, p.createdAt, p.privacy, p.image_url, u.firstName, u.lastName, r2.reaction
ORDER BY
    p.createdAt DESC;
`
	rows, err := r.db.Query(query, userID.String(), userID.String(), userID.String(), userID.String())
	if err != nil {
		fmt.Println("error in executing", err)
		return nil, err
	}
	defer rows.Close()
	var posts []models.Post
	for rows.Next() {
		var p models.Post
		if err := rows.Scan(&p.Id, &p.User.Id, &p.Content, &p.CreatedAt, &p.Privacy, &p.Img, &p.User.FirstName, &p.User.LastName, &p.TotalLikes, &p.Liked); err != nil {
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
