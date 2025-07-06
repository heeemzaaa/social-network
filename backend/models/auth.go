package models

import "time"

type User struct {
	Id             string `json:"id,omitempty"`
	FirstName      string `json:"firstname,omitempty"`
	LastName       string `json:"lastname,omitempty"`
	BirthDate      string `json:"birthdate,omitempty"`
	Email          string `json:"email,omitempty"`
	Password       string `json:"password,omitempty"`
	Nickname       string `json:"nickname,omitempty"`
	AboutMe        string `json:"about_me,omitempty"`
	ProfileImage   string `json:"avatar,omitempty"`
	Visibility     string `json:"visibility,omitempty"`
	ProfileImgSize int64  `json:"img_size,omitempty"`
}

func NewUser() *User {
	return &User{}
}

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
	UserId   string    `json:"user_id"`
	Username string    `json:"username,omitempty"`
	ExpDate  time.Time `json:"expiration_date,omitempty"`
}

func NewSession() *Session {
	return &Session{}
}

type IsLoggedIn struct {
	IsLoggedIn bool `json:"is_logged_in"`
}

func NewIsLoggedIn() *IsLoggedIn {
	return &IsLoggedIn{}
}
