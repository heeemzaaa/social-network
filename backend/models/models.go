package models

type ErrorJson struct {
	Status  int    `json:"status"`
	Error   string `json:"error,omitempty"`
	Message any    `json:"errors,omitempty"`
}

func NewErrorJson(status int, err string, message any) *ErrorJson {
	return &ErrorJson{
		Status:  status,
		Error:   err,
		Message: message,
	}
}
