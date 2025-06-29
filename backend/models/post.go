package models

import "time"

type Post struct {
	Id             int       `json:"id,omitempty"`
	UserId         int       `json:"user_id,omitempty"`
	Username       string    `json:"user_name,omitempty"`
	Content        string    `json:"content"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	TotalComments  int       `json:"total_comments"`
	TotalLikes     int       `json:"total_likes"`
	Liked          int       `json:"liked"`
}

func NewPost() *Post {
	return &Post{}
}

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

func NewComment() *Comment {
	return  &Comment{}
}

type Reaction struct {
	Id           int    `json:"id,omitempty"`
	EntityTypeId int    `json:"entity_type_id,omitempty"`
	EntityType   string `json:"entity_type,omitempty"`
	EntityId     int    `json:"entity_id,omitempty"`
	Reaction     int    `json:"reaction"`
	UserId       int    `json:"user_id,omitempty"`
}

func NewReaction() *Reaction {
	return &Reaction{}
}