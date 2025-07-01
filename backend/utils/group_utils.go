package utils

import "errors"

func ValidateTitle(title string) error {
	if len(title) < 3 {
		return errors.New("title is too short! 3 characters min")
	}
	if len(title) > 100 {
		return errors.New("title is too large! 100 characters max")
	}
	return nil
}

func ValidateDesc(desc string) error {
	if len(desc) < 10 {
		return errors.New("description is too short! 10 characters min")
	}
	if len(desc) > 1000 {
		return errors.New("description is too large! 1000 characters max")
	}
	return nil
}
