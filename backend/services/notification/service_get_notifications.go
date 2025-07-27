package notification

import (
	"social-network/backend/models"
	"sort"
	"strconv"
)

// get notification by 10
func (NS *NotificationService) GetService(user_id, queryParam string) ([]models.Notification, *models.ErrorJson) {

	all, err := NS.repo.SelectAllNotification(user_id)
	if err != nil {
		return nil, err
	}
	len := len(all)
	nbr, _ := strconv.Atoi(queryParam)
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

	switch {
	case len <= 10:
		return all, nil
	case nbr <= 0:
		return all[:10], nil
	case len <= nbr:
		return []models.Notification{}, nil
	case len < nbr+10:
		return all[nbr:], nil
	}
	return all[nbr : nbr+10], nil
}
