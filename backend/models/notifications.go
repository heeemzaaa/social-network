package models

import "time"

// database notification structure
type Notification struct {
	Id        string
	Sender    User
	Target    any
	Type      string
	Data      any
	CreatedAt time.Time
}

// target here could be a user of even a group ( in the case of the creation of the event the target is the group)

// NewNotification creates a new instance of Notification.
func NewNotification() *Notification {
	return &Notification{}
}
