package notification

import (
	sc "social-network/backend/handlers/chat"
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
	chatServer  *sc.ChatServer //
}

func NewNotifService(notifRepo *rn.NotifRepository, authRepo *ra.AuthRepository, profileRepo *rp.ProfileRepository, groupRepo *rg.GroupRepository, chatServer *sc.ChatServer) *NotificationService {
	return &NotificationService{
		notifRepo:   notifRepo,
		authRepo:    authRepo,
		profileRepo: profileRepo,
		groupRepo:   groupRepo,
		chatServer:  chatServer, //
	}
}

func (NS *NotificationService) ToggleAllSeenFalse(notifications []models.Notification) *models.ErrorJson {
	for _, notification := range notifications {
		if errJson := NS.notifRepo.UpdateSeen(notification.Id); errJson != nil {
			return errJson
		}
	}

	if errJson := NS.broadcast(notifications[0].RecieverId); errJson != nil {
		return errJson
	}
	return nil
}

func (NS *NotificationService) ToggleAllStaus(notifications []models.Notification, value, notifType string) *models.ErrorJson {
	for _, notification := range notifications {
		if errJson := NS.notifRepo.UpdateStatus(notification.Id, value); errJson != nil {
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

	NS.broadcast(recieverId)
	return nil
}

func (NS *NotificationService) IsHasSeenFalse(userId string) (bool, *models.ErrorJson) {
	seen, errJson := NS.notifRepo.IsHasSeenFalse(userId)
	if errJson != nil {
		return false, errJson
	}
	return seen, nil
}

func (NS *NotificationService) broadcast(recieverId string) *models.ErrorJson {
	hasSeen, errJson := NS.IsHasSeenFalse(recieverId)
	if errJson != nil {
		return errJson
	}
	if hasSeen {
		errJson = NS.chatServer.SendNotificationToUser(recieverId, "has new notification", "true")
	} else {
		errJson = NS.chatServer.SendNotificationToUser(recieverId, "dont have new notification", "false")
	}
	return errJson
}
