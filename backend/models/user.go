package models

import "time"

type User struct{}

type Session struct {
	Id       int       `json:"id,omitempty"`
	Token    string    `json:"token"`
	UserId   int       `json:"user_id"`
	Username string    `json:"username,omitempty"`
	ExpDate  time.Time `json:"expiration_date,omitempty"`
}
