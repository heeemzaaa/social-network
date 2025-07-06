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
	Image          string    `json:"image"`
	Members        []User
	Posts          []Post
	Events         []Event
}

type Event struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	EventDate   time.Time `json:"event_date"`
}

// when trying to  create a group
type ErrGroup struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

//  when trying to join a group

type ErrJoinGroup struct {
	GroupId string `json:"group_id"`
}
