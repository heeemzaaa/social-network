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

type Event struct {
	GroupId        *uuid.UUID `json:"group_id,omitempty"`
	EventCreator   string     `json:"event_creator"`
	EventCreatorId *uuid.UUID `json:"event_creator_id"`
	Title          string     `json:"title"`
	Description    string     `json:"description"`
	EventDate      time.Time  `json:"event_date"`
	Going          int        `json:"going"`
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

type ErrEventGroup struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	EventDate   string `json:"event_date"`
}
