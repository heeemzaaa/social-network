package models

import "time"

type Comment struct {
	Id         string    `json:"id,omitempty"`
	PostId     string    `json:"post_id"`
	User       User      `json:"user"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	Content    string    `json:"content"`
	Media      string    `json:"media,omitempty"`
	TotalLikes int       `json:"total_likes,omitempty"`
	Liked      int       `json:"liked,omitempty"`
}
