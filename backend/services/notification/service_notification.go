package notification

import (
	"social-network/backend/models"
	ra "social-network/backend/repositories/auth"
	rg "social-network/backend/repositories/group"
	rn "social-network/backend/repositories/notification"
	rp "social-network/backend/repositories/profile"
	"strings"
)

type NotificationService struct {
	authRepo    *ra.AuthRepository
	notifRepo   *rn.NotifRepository
	profileRepo *rp.ProfileRepository
	groupRepo   *rg.GroupRepository
}

func NewNotifService(notifRepo *rn.NotifRepository, authRepo *ra.AuthRepository, profileRepo *rp.ProfileRepository, groupRepo *rg.GroupRepository) *NotificationService {
	return &NotificationService{
		notifRepo:   notifRepo,
		authRepo:    authRepo,
		profileRepo: profileRepo,
		groupRepo:   groupRepo,
	}
}

func (NS *NotificationService) ToggleAllSeenFalse(notifications []models.Notification) *models.ErrorJson {
	for _, notification := range notifications {
		if errJson := NS.notifRepo.UpdateSeen(notification.Id); errJson != nil {
			return errJson
		}
	}
	return nil
}

func (NS *NotificationService) ToggleAllStaus(notifications []models.Notification, value, notifType string) *models.ErrorJson {
	for _, notification := range notifications {
		if errJson := NS.notifRepo.UpdateStatusById(notification.Id, value); errJson != nil {
			return errJson
		}
	}
	return nil
}

func (NS *NotificationService) GetAllNotificationByType(user_id, notifType string) ([]models.Notification, *models.ErrorJson) {
	all, err := NS.notifRepo.SelectAllNotificationByType(user_id, notifType)
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (NS *NotificationService) DeleteService(recieverId, senderId, notifType, groupId string) *models.ErrorJson {
	if strings.HasPrefix(notifType, "follow") {
		if errJson := NS.notifRepo.DeleteFollowNotification(senderId, recieverId, notifType); errJson != nil {
			return errJson
		}
	} else if notifType != "group-event" {
		if errJson := NS.notifRepo.DeleteGroupNotification(senderId, recieverId, notifType, groupId); errJson != nil {
			return errJson
		}
	}
	return nil
}

func (NS *NotificationService) IsHasSeenFalse(user_id string) (bool, *models.ErrorJson) {
	isValid, errJson := NS.notifRepo.IsHasSeenFalse(user_id)
	if errJson != nil {
		return false, errJson
	}
	return isValid, nil
}
