package utils

import (
	"social-network/backend/models"
)

func GetIndexOf(Slice []models.Notification, id string) int {
	for i, item := range Slice {
		if item.Id == id {
			return i
		}
	}
	return -1
}
