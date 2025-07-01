package models


type ErrorJson struct {
	Status  int `json:"status"`
	Message any `json:"errors"`
}

func NewErrorJson(status int, message any) *ErrorJson {
	return &ErrorJson{
		Status:  status,
		Message: message,
	}
}