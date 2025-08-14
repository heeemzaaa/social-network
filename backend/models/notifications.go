package models

import "time"

// database notification structure
type Notification struct {
	Id        string    `json:"id,omitempty"`
	SenderID  string    `json:"sender_id,omitempty"`
	TargetID  string    `json:"target_id,omitempty"`
	Seen      bool      `json:"seen,omitempty"`
	Type      string    `json:"type_notification,omitempty"`
	Status    string    `json:"status,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"craeted_at"`
	GroupID   string    `json:"group_id,omitempty"`

	SenderFullName string `json:"sender_fullname,omitempty"`
	GroupName      string `json:"group_name,omitempty"`
	EventID        string `json:"event_id,omitempty"`
	EventName      string `json:"event_name,omitempty"`
}

// NewNotification creates a new instance of Notification.
func NewNotification() *Notification {
	return &Notification{}
}

func (n *Notification) PointerToSimple() Notification {
	return Notification{
		Id:        n.Id,
		SenderID:  n.SenderID,
		TargetID:  n.TargetID,
		Seen:      n.Seen,
		Type:      n.Type,
		Status:    n.Status,
		Content:   n.Content,
		CreatedAt: n.CreatedAt,

		SenderFullName: n.SenderFullName,
		GroupName:      n.GroupName,
		EventID:        n.EventID,
		EventName:      n.EventName,
	}
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
	Content        string
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

// structure of response message after notification action
type ResponseData struct {
	Status        bool
	Message       string
	Notifications []Notification
}

// NewResponseData creates a new instance of ResponseData.
func NewResponseData() *ResponseData {
	return &ResponseData{}
}

type Data struct {
	Data         any          `json:"data"`
	Notification Notification `json:"notification"`
}
