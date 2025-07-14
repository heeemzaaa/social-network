package models

type Message struct {
	ID           string `json:"id,omitempty"`
	Type         string `json:"type"`
	SenderID     string `json:"sender_id"`
	SenderName   string `json:"sender_name,omitempty"`
	ReceiverName string `json:"receiver_name,omitempty"`
	TargetID     string `json:"target_id"`
	Content      string `json:"content"`
	CreatedAt    string `json:"created_at,omitempty"`
}

type MessageErr struct {
	Content   string `json:"content"`
	TargetID  string `json:"target_id"`
	Type      string `json:"type"`
}

func NewMessageErr() *MessageErr {
	return &MessageErr{}
}
