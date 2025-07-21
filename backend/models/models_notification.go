package models

import "time"

// notification structure
type Notification struct {
	Id          string
	Reciever_Id string
	Sender_Id   string
	Seen        bool
	Type        string // follow request // invite group // join group to admin // event notif for members group // public request . bonus
	Status      string // accepted or rejected or "" for public request
	Content     string
	CreatedAt   time.Time
}

func NewNotification() *Notification {
	return &Notification{}
}

// new notification request data structure
type Notif struct {
	Sender_Id        string // user-id // credencials
	SenderFullName   string
	Reciever_Id      string // user-profile-id // user-profile-id // group-id [AllGpMb] // group-id [admin] // user-target-id
	ReceiverFullName string
	Type             string // follow-private // follow-public // group-event // group-join // group-invitation
	Content          string // notification informations
}

func NewNotif() *Notif {
	return &Notif{}
}

// update notification request data structure
type Unotif struct {
	Notif_Id string // notification id
	Status   string // accept || reject

}

func UpdateNotif() *Unotif {
	return &Unotif{}
}

type HasSeen struct {
	Status  bool
	Message string
}
