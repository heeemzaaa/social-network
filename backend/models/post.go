package models

import "time"

type Post struct {
	Id            string    `json:"id,omitempty"`
	User          User      `json:"user"`
	Content       string    `json:"content,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	Media         string    `json:"media,omitempty"`
	TotalComments int       `json:"total_comments,omitempty"`
	TotalLikes    int       `json:"total_likes,omitempty"`
	Liked         int       `json:"liked,omitempty"`
}

func NewPost() *Post {
	return &Post{}
}

type Reaction struct {
	Id         string `json:"id,omitempty"`
	EntityType string `json:"entity_type,omitempty"`
	EntityId   string `json:"entity_id,omitempty"`
	Reaction   int    `json:"reaction"`
	UserId     string `json:"user_id,omitempty"`
}

func NewReaction() *Reaction {
	return &Reaction{}
}
