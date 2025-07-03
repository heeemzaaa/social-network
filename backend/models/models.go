package models

type ErrorJson struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message any    `json:"errors"`
}

func NewErrorJson(status int, err string, message any) *ErrorJson {
	return &ErrorJson{
		Status:  status,
		Error:   err,
		Message: message,
	}
}
