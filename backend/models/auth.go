package models

type User struct {
	Id              string `json:"id,omitempty"`
	FullName        string `json:"fullname,omitempty"`
	FirstName       string `json:"firstname,omitempty"`
	LastName        string `json:"lastname,omitempty"`
	BirthDate       string `json:"birthdate,omitempty"`
	Email           string `json:"email,omitempty"`
	Password        string `json:"password,omitempty"`
	Nickname        string `json:"nickname,omitempty"`
	AboutMe         string `json:"about_me,omitempty"`
	ImagePath       string `json:"avatar,omitempty"`
	Visibility      string `json:"visibility,omitempty"`
	LastMessage     string `json:"last_message,omitempty"`
	LastInteraction string `json:"last_interaction,omitempty"`
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
	Id        int    `json:"id,omitempty"`
	Token     string `json:"token"`
	UserId    string `json:"user_id"`
	Username  string `json:"username,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
}

func NewSession() *Session {
	return &Session{}
}

type UserData struct {
	IsLoggedIn bool   `json:"is_logged_in"`
	Id         string `json:"id,omitempty"`
	Nickname   string `json:"nickname,omitempty"`
}

type ContextKey struct {
	Key string
}

// there is a  problem when doing this with contexts

func NewContextKey(key string) *ContextKey {
	return &ContextKey{Key: key}
}
