package models

type Notifications struct {
	Id          string
	Reciever_Id string // or name
	Sender_Id   string
	Seen        bool
	Type        string // follow request // invite group // join group to admin // event notif for members group // public request . bonus
	Status      string // accepted or rejected or "" for public request
	Content     string
	// Created_at time.Time
}
func NewNotification() *Notifications {
	return &Notifications{}
}

// for add new notifications // request data structure
type Notif struct {
	Sender_Id   string // user-id // credencials
	Reciever_Id string // user-profile-id // user-profile-id // group-id [AllGpMb] // group-id [admin] // user-target-id
	Type        string // follow-private // follow-public // group-event // group-join // group-invitation
	Content     string // notification informations
}
func NewNotif() *Notif {
	return &Notif{}
}