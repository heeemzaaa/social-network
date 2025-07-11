package models

import "time"

type Post struct {
	Id            int       `json:"id,omitempty"`
	UserId        int       `json:"user_id,omitempty"`
	Username      string    `json:"user_name,omitempty"`
	Content       string    `json:"content,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	TotalComments int       `json:"total_comments,omitempty"`
	TotalLikes    int       `json:"total_likes,omitempty"`
	Liked         int       `json:"liked,omitempty"`
}

func NewPost() *Post {
	return &Post{}
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