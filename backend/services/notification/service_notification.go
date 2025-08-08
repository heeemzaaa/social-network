package notification

import (
	"social-network/backend/models"
	"social-network/backend/repositories/auth"
	"social-network/backend/repositories/notification"
)

type NotificationService struct {
	repo2 *auth.AuthRepository
	repo  *notification.NotifRepository
}

func NewNotifService(repo *notification.NotifRepository, repo2 *auth.AuthRepository) *NotificationService {
	return &NotificationService{
		repo:  repo,
		repo2: repo2,
	}
}

func (NS *NotificationService) ToggleAllSeenFalse(notifications []models.Notification) *models.ErrorJson {
	for _, notification := range notifications {
		if errJson := NS.repo.UpdateSeen(notification.Id); errJson != nil {
			return errJson
		}
	}
	return nil
}

// toggle all notifications status by type
func (NS *NotificationService) ToggleAllStaus(notifications []models.Notification, value, notifType string) *models.ErrorJson {
	for _, notification := range notifications {
		if errJson := NS.repo.UpdateStatusById(notification.Id, value); errJson != nil {
			return errJson
		}
	}
	return nil
}

// toggle notifications status by type
func (NS *NotificationService) ToggleStaus(userID, reciever, value, notifType string) *models.ErrorJson {
	if errJson := NS.repo.UpdateStatusByType(userID, reciever, value, notifType); errJson != nil {
		return errJson
	}
	return nil
}

// get all notification by type
func (NS *NotificationService) GetAllNotifService(user_id, notifType string) ([]models.Notification, *models.ErrorJson) {
	all, err := NS.repo.SelectAllNotification(user_id)
	if err != nil {
		return nil, err
	}

	return all, nil
}

// should be add for delete event notification
func (NS *NotificationService) DeleteService(reciever, sender, notifType, value string) *models.ErrorJson {
	if notifType != "follow-private" {
		if errJson := NS.repo.DeleteGroupNotification(sender, reciever, notifType, value); errJson != nil {
			return errJson
		}
	} else {
		if errJson := NS.repo.DeleteFollowNotification(sender, reciever, notifType, value); errJson != nil {
			return errJson
		}
	}
	return nil
}

func (NS *NotificationService) IsHasSeenFalse(user_id string) (bool, *models.ErrorJson) {
	isValid, errJson := NS.repo.IsHasSeenFalse(user_id)
	if errJson != nil {
		return false, errJson
	}
	return isValid, nil
}
