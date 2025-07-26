package models

import "time"

// notification structure
type Notification struct {
	Id         string
	RecieverId string
	SenderId   string
	Seen       bool
	Type       string // follow request // invite group // join group to admin // event notif for members group
	Status     string // accepted or rejected or later
	Content    string
	GroupId    string
	CreatedAt  time.Time
}

func NewNotification() *Notification {
	return &Notification{}
}

// new notification request data structure
type Notif struct {
	SenderId         string // user-id // credencials
	SenderFullName   string
	RecieverId       string // user-profile-id // user-profile-id // group-id [AllGpMb] // group-id [admin] // user-target-id
	ReceiverFullName string
	Type             string // follow-private // follow-public // group-event // group-join // group-invitation
	GroupName        string // notification informations
	GroupId          string
}

func NewNotif() *Notif {
	return &Notif{}
}

// update notification request data structure
type Unotif struct {
	NotifId string // notification id
	Status  string // accept || reject
	Type    string
	// GroupId string
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
