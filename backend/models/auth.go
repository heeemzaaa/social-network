package models

type User struct {
	Id        string `json:"id,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	BirthDate string `json:"birthdate,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	UserName  string `json:"nickname,omitempty"`
	AboutMe   string `json:"about_me,omitempty"`
}

type Login struct {
	LoginField string `json:"login"`
	Password   string `json:"password"`
}

func NewLogin() *Login {
	return &Login{}
}

type Session struct {
	Id       int    `json:"id,omitempty"`
	Token    string `json:"token"`
	UserId   string `json:"user_id"`
	Username string `json:"username,omitempty"`
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
