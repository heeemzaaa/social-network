package utils

import (
	"errors"
	"time"
)

func ValidateTitle(title string) error {
	if title == "" {
		return errors.New("title can not be empty")
	}
	if len(title) < 3 {
		return errors.New("title is too short! 3 characters min")
	}
	if len(title) > 100 {
		return errors.New("title is too large! 100 characters max")
	}
	return nil
}

func ValidateDesc(desc string) error {
	if desc == "" {
		return errors.New("description can not be empty")
	}
	if len(desc) < 10 {
		return errors.New("description is too short! 10 characters min")
	}
	if len(desc) > 1000 {
		return errors.New("description is too large! 1000 characters max")
	}
	return nil
}

func ValidateDate(date time.Time) error {
	if date.IsZero() {
		return errors.New("the date is not set up")
	}
	if date.Before(time.Now()) {
		return errors.New("please set a date that comes after ")
	}

	return nil
}

func IsValidFilter(filter string) bool {
	return filter == "owned" || filter == "available" || filter == "joined"
}
