package models

type Session struct {
	Id       int    `json:"id,omitempty"`
	Token    string `json:"token"`
	UserId   int    `json:"user_id"`
	Username string `json:"username,omitempty"`
}

func NewSession() *Session {
	return &Session{}
}
