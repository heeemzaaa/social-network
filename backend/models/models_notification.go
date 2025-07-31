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

func NewNotification() *Notification {
	return &Notification{}
}

// structue of new notification 
type Notif struct {
	SenderId         string
	RecieverId       string
	GroupId          string
	EventId          string
	Type             string
	SenderFullName   string
	GroupName        string
}

func NewNotif() *Notif {
	return &Notif{}
}

// structure of update notification
type Unotif struct {
	NotifId string
	Type    string
	Status  string
}

func UpdateNotif() *Unotif {
	return &Unotif{}
}

type ResponseMsg struct {
	Status  bool
	Message string
}

func NewResponseMsg() *ResponseMsg {
	return &ResponseMsg{}
}