package models

type RegisterError struct {
	Nickname      string `json:"nickname"`
	Age           string `json:"age"`
	Gender        string `json:"gender"`
	FirstName     string `json:"firstname"`
	LastName      string `json:"lastname"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	VerifPassword string `json:"password2"`
}

type LoginERR struct {
	LoginField string `json:"login"`
	Password   string `json:"password"`
}
