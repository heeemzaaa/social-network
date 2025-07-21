package utils

import "strconv"

func IsValidQueryParam(str string) bool {
	_, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return true
}
