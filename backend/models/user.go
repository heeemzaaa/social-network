package models

type User struct {
	ID         string `json:"id,omitempty"`
	Email      string `json:"email,omitempty"`
	FirstName  string `json:"first_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	BirthDate  string `json:"birth_date,omitempty"`
	Nickname   string `json:"nickname,omitempty"`
	AvatarPath string `json:"avatar,omitempty"`
	AboutMe    string `json:"about_me,omitempty"`
	Visibility string `json:"visibility,omitempty"`
	Password   string `json:"password,omitempty"`
}
