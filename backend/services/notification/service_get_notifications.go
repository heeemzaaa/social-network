package notification

import (
	"social-network/backend/models"
	"social-network/backend/utils"
	"sort"
)

// get notification by 10
func (NS *NotificationService) GetService(userId, notificationId string) ([]models.Notification, *models.ErrorJson) {

	all, err := NS.notifRepo.SelectAllNotification(userId)
	if err != nil {
		return nil, err
	}

	maxLen := len(all)
	if maxLen == 0 {
		return []models.Notification{}, nil
	}

	sort.Slice(all, func(i, j int) bool {
		if all[i].Status == "later" && all[j].Status != "later" {
			return true
		}
		if all[i].Status != "later" && all[j].Status == "later" {
			return false
		}

		return all[i].CreatedAt.After(all[j].CreatedAt)
	})

	// get nbr when equal to index of last notification in container
	if notificationId == "0" {
		if maxLen <= 10 {
			return all, nil
		}
		return all[:10], nil
	}

	nbr := utils.GetIndexOf(all, notificationId)
	if nbr == -1 {
		return []models.Notification{}, &models.ErrorJson{Status: 400, Message: "Invalid Operation", Error: "400 - Bad Request"}
	}
	nbr++

	switch {
	case maxLen <= nbr:
		return []models.Notification{}, nil
	case maxLen < nbr+10:
		return all[nbr:], nil
	}
	return all[nbr : nbr+10], nil
}
