package models

type Message struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	SenderID   string `json:"sender_id"`
	SenderName string `json:"sender_name"`
	TargetID   string `json:"target_id"`
	Content    string `json:"content"`
	CreatedAt  string `json:"created_at,omitempty"`
}

type MessageErr struct {
	Content    string `json:"content"`
	ReceiverID string `json:"receiver_id"`
	Type       string `json:"type"`
	CreatedAt  string `json:"created_at"`
}

func NewMessageErr() *MessageErr {
	return &MessageErr{}
}