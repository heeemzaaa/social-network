package models

type ErrorJson struct {
	Status  int `json:"status"`
	Message any `json:"errors"`
}
