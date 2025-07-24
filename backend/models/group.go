package models

import (
	"time"
)

type Group struct {
	GroupId              string    `json:"group_id,omitempty"`
	GroupCreatorId       string    `json:"group_creator_id,omitempty"`
	Title                string    `json:"title"`
	GroupCreatorFullName string    `json:"group_creator,omitempty"`
	GroupCreatorNickname string    `json:"group_creator_nickname,omitempty"`
	Description          string    `json:"description"`
	ImagePath            string    `json:"image_path,omitempty"`
	Image                string    `json:"image,omitempty"`
	CreatedAt            time.Time `json:"created_at,omitempty"`
	Total_Members        int       `json:"total_members,omitempty"`
	Members              []User    `json:"members,omitempty"`
	Posts                []Post    `json:"posts,omitempty"`
	Events               []Event   `json:"events,omitempty"`
	LastInteraction      string    `json:"last_interaction,omitempty"`
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
	EventId      string    `json:"event_id,omitempty"`
	GroupId      string    `json:"group_id,omitempty"`
	EventCreator User      `json:"event_creator,omitempty"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	EventDate    string    `json:"event_date"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	Going        int       `json:"going"`
}

type ErrEventGroup struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	EventDate   string `json:"event_date"`
}

type PostGroup struct {
	Id            string    `json:"id,omitempty"`
	GroupId       string    `json:"group_id,omitempty"`
	User          User      `json:"user"`
	Content       string    `json:"content"`
	ImagePath     string    `json:"image_path,omitempty"`
	Image         string    `json:"image,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	TotalComments int       `json:"total_comments"`
	TotalLikes    int       `json:"total_likes"`
	Liked         string    `json:"liked"`
}
type PostGroupErr struct {
	Content string `json:"content"`
}

type CommentGroup struct {
	Id         string    `json:"id,omitempty"`
	GroupId    string    `json:"group_id,omitempty"`
	PostId     string    `json:"post_id"`
	User       User      `json:"user"`
	Content    string    `json:"content"`
	ImagePath  string    `json:"image_path,omitempty"`
	Image      string    `json:"image,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	TotalLikes int       `json:"total_likes"`
	Liked      int       `json:"liked"`
}

type CommentGroupErr struct {
	Content string `json:"content"`
	PostId  string `json:"post_id"`
}

type GroupReaction struct {
	Id         string `json:"id,omitempty"`
	EntityType string `json:"entity_type,omitempty"`
	EntityId   string `json:"entity_id,omitempty"`
	Reaction   int    `json:"reaction"`
	UserId     string `json:"user_id,omitempty"`
	GroupId    string `json:"group_id,omitempty"`
}

type GroupReactionErr struct {
	EntityId   string `json:"entity_id"`
	EntityType string `json:"entity_type"`
	Reaction   string `json:"reaction,omitempty"`
}

type UserEventAction struct {
	Id      string `json:"id,omitempty"`
	UserId  string `json:"user_id,omitempty"`
	GroupId string `json:"group_id,omitempty"`
	EventId string `json:"event_id,omitempty"`
	Action  int    `json:"action"` // it has to remain without omiempty
	// if it is done 0 won't be displayed because it is considered empty
}

type UserErr struct {
	UserId string `json:"user"`
}

type UserEventActionErr struct {
	Action string `json:"action,omitempty"`
}

// models for the group to send the requests

type Request struct {
	SenderId         string
	RecieverId       string
	Type             string
	SenderFullName   string
	ReceiverFullName string
	SenderNickname   string
	ReceiverNickname string
	GroupId          string
	GroupTitle       string
}
