package models

import "time"

// database notification structure
type Notification struct {
	Id             string 
	SenderId       string
	RecieverId     string
	GroupId        string
	EventId        string
	Type           string
	SenderFullName string
	GroupName      string
	Status         string
	Seen           bool
	CreatedAt      time.Time
}

// NewNotification creates a new instance of Notification.
func NewNotification() *Notification {
	return &Notification{}
}

// structue of new notification
type Notif struct {
	SenderId       string
	RecieverId     string
	GroupId        string
	EventId        string
	Type           string
	SenderFullName string
	GroupName      string
}

// NewNotif creates a new instance of Notif.
func NewNotif() *Notif {
	return &Notif{}
}

// structure of update notification
type Unotif struct {
	NotifId string
	Type    string
	Status  string
}

// NewUnotif creates a new instance of Unotif.
func UpdateNotif() *Unotif {
	return &Unotif{}
}

// structure of response message after notification action
type ResponseMsg struct {
	Status  bool
	Message string
}

// NewResponseMsg creates a new instance of ResponseMsg.
func NewResponseMsg() *ResponseMsg {
	return &ResponseMsg{}
}
