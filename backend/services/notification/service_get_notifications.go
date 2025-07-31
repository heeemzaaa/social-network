package notification

import (
	"social-network/backend/models"
	"social-network/backend/utils"
	"sort"
)

// get notification by 10
func (NS *NotificationService) GetService(user_id, notificationId string) ([]models.Notification, *models.ErrorJson) {

	all, err := NS.notifRepo.SelectAllNotification(user_id)
	if err != nil {
		return nil, err
	}

	maxLen := len(all)

	sort.Slice(all, func(i, j int) bool {
		// Check if one is "later" and the other isn't
		if all[i].Status == "later" && all[j].Status != "later" {
			return true
		}
		if all[i].Status != "later" && all[j].Status == "later" {
			return false
		}
		// If both are "later" or both are not, sort by CreatedAt descending
		return all[i].CreatedAt.After(all[j].CreatedAt)
	})

	// get nbr when equal to index of last notification in container
	if notificationId == "0" {
		// fmt.Println("nbr ===> ", "0")
		if maxLen <= 10 {
			return all, nil
		}
		return all[:10], nil
	}
	// fmt.Println("ID ===> ", notificationId)

	nbr := utils.GetIndexOf(all, notificationId)
	// fmt.Println("nbr ===> ", nbr)

	if nbr == -1 {
		return []models.Notification{}, &models.ErrorJson{Status: 400, Message: "Invalid Operation", Error: "400 - Bad Request"}
	}
	nbr++

	// fmt.Println("HERE !!!")

	switch {
	case maxLen <= nbr:
		return []models.Notification{}, nil
	case maxLen < nbr+10:
		return all[nbr:], nil
	}
	return all[nbr : nbr+10], nil
}
