package models




// we can make the message interface and then accpet all of them but for now let's work so
type ErrorJson struct {
	Status  int `json:"status"`
	Message any `json:"errors"`
}


