package models

import "time"

// notification structure
type Notification struct {
	Id             string
	SenderId       string
	RecieverId     string
	GroupId        string
	EventId        string
	Type           string // follow request // invite group // join group to admin // event notif for members group
	SenderFullName string
	GroupName      string // notification informations
	Status         string // accepted or rejected or later
	Seen           bool
	CreatedAt      time.Time
}

func NewNotification() *Notification {
	return &Notification{}
}

// data structue of new notification 
type Notif struct {
	SenderId         string // user-id // credencials
	RecieverId       string // user-profile-id // user-profile-id // group-id [AllGpMb] // group-id [admin] // user-target-id
	GroupId          string
	EventId          string
	Type             string // follow-private // follow-public // group-event // group-join // group-invitation
	SenderFullName   string
	GroupName        string // notification informations
}

func NewNotif() *Notif {
	return &Notif{}
}

// update notification request data structure
type Unotif struct {
	NotifId string // notification id
	Type    string
	Status  string // accept || reject
}

func UpdateNotif() *Unotif {
	return &Unotif{}
}

type HasSeen struct {
	Status  bool
	Message string
}

func NewResponseMsg() *Unotif {
	return &Unotif{}
}

type ResponseMsg struct {
	Status  bool
	Message string
}
