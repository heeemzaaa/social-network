package models

import "time"

type Comment struct {
	Id         int       `json:"id,omitempty"`
	PostId     int       `json:"post_id"`
	UserId     int       `json:"user_id,omitempty"`
	Username   string    `json:"user_name,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	Content    string    `json:"content"`
	TotalLikes int       `json:"total_likes"`
	Liked      int       `json:"liked"`
}
