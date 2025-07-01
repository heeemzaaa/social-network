package models

import (
	"time"

	"github.com/google/uuid"
)

type Group struct {
	GroupId        uuid.UUID `json:"group_id"`
	GroupCreatorId uuid.UUID `json:"group_creator_id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	ImagePath      string    `json:"image_path"`
	Members        []User
	Posts          []Post
	Events         []Event
}

type Event struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	EventDate   time.Time `json:"event_date"`
}
type Post struct {
	Id            int       `json:"id,omitempty"`
	UserId        int       `json:"user_id,omitempty"`
	Username      string    `json:"user_name,omitempty"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	TotalComments int       `json:"total_comments"`
	TotalLikes    int       `json:"total_likes"`
	Liked         int       `json:"liked"`
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

type ErrGroup struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
