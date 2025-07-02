package models

import "github.com/google/uuid"

type Login struct {
	LoginField string `json:"login"`
	Password   string `json:"password"`
}

func NewLogin() *Login {
	return &Login{}
}

type Session struct {
	Id       int       `json:"id,omitempty"`
	Token    string    `json:"token"`
	UserId   string `json:"user_id"`
	Username string    `json:"username,omitempty"`
}

func NewSession() *Session {
	return &Session{}
}

type IsLoggedIn struct {
	IsLoggedIn bool   `json:"is_logged_in"`
	Id         int    `json:"id,omitempty"`
	Nickname   string `json:"nickname,omitempty"`
}

func NewIsLoggedIn() *IsLoggedIn {
	return &IsLoggedIn{}
}
