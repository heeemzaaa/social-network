package utils

import (
	"social-network/backend/models"
)

// GetIndexOf returns the index of a notification with the specified ID in the slice.
func GetIndexOf(Slice []models.Notification, id string) int {
	for i, item := range Slice {
		if item.Id == id {
			return i
		}
	}
	return -1
}
