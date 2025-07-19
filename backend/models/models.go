package models

import (
	"encoding/json"
	"strings"
	"time"
)

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

// First create a type alias
type Date struct {
	Date time.Time
}

// e7m constructor

func NewDate(t time.Time) *Date {
	return &Date{Date: t}
}

// decodng we can check the edge cases
// so the only way to do so is to check it here or create another function
func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*d = Date{Date: t}
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d.Date))
}

// Maybe a Format function for printing your date
func (d Date) Format(s string) string {
	t := time.Time(d.Date)
	return t.Format(s)
}
