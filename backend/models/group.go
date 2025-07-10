package models

import (
	"time"

	"github.com/google/uuid"
)

type Group struct {
	GroupId        *uuid.UUID `json:"group_id,omitempty"`
	GroupCreatorId *uuid.UUID `json:"group_creator_id,omitempty"`
	Title          string     `json:"title"`
	Description    string     `json:"description"`
	ImagePath      string     `json:"image_path,omitempty"`
	Image          string     `json:"image,omitempty"`
	Members        []User     `json:"members,omitempty"`
	Posts          []Post     `json:"posts,omitempty"`
	Events         []Event    `json:"events,omitempty"`
}

// when trying to  create a group
type ErrGroup struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// when trying to join a group
type ErrJoinGroup struct {
	GroupId string `json:"group_id"`
}

type Event struct {
	GroupId        *uuid.UUID `json:"group_id,omitempty"`
	EventCreator   string     `json:"event_creator"`
	EventCreatorId *uuid.UUID `json:"event_creator_id"`
	Title          string     `json:"title"`
	Description    string     `json:"description"`
	EventDate      time.Time  `json:"event_date"`
	Going          int        `json:"going"`
}

type ErrEventGroup struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	EventDate   string `json:"event_date"`
}

type PostGroup struct {
	Id            *uuid.UUID `json:"id,omitempty"`
	GroupId       *uuid.UUID `json:"group_id,omitempty"`
	UserId        *uuid.UUID `json:"user_id,omitempty"`
	Username      string     `json:"user_name,omitempty"`
	Content       string     `json:"content"`
	ImagePath     string     `json:"image_path,omitempty"`
	Image         string     `json:"image,omitempty"`
	CreatedAt     time.Time  `json:"created_at,omitempty"`
	TotalComments int        `json:"total_comments"`
	TotalLikes    int        `json:"total_likes"`
	Liked         int        `json:"liked"`
}
type PostGroupErr struct {
	Content string `json:"content"`
}

type CommentGroup struct {
	Id         *uuid.UUID `json:"id,omitempty"`
	GroupId    *uuid.UUID `json:"group_id,omitempty"`
	PostId     *uuid.UUID `json:"post_id"`
	UserId     *uuid.UUID `json:"user_id,omitempty"`
	Username   string     `json:"user_name,omitempty"`
	Content    string     `json:"content"`
	ImagePath  string     `json:"image_path,omitempty"`
	Image      string     `json:"image,omitempty"`
	CreatedAt  time.Time  `json:"created_at,omitempty"`
	TotalLikes int        `json:"total_likes"`
	Liked      int        `json:"liked"`
}

type CommentGroupErr struct {
	Content string `json:"content"`
	PostId  string `json:"post_id"`
}
