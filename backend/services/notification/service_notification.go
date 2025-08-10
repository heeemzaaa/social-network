package notification

import (
	"strings"

	sc "social-network/backend/handlers/chat"
	"social-network/backend/models"
	rg "social-network/backend/repositories/group"
	rn "social-network/backend/repositories/notification"
	sa "social-network/backend/services/auth"
	sp "social-network/backend/services/profile"
)

type NotificationService struct {
	authService    *sa.AuthService
	notifRepo      *rn.NotifRepository
	profileService *sp.ProfileService
	groupRepo      *rg.GroupRepository
	chatServer     *sc.ChatServer
}

func NewNotifService(notifRepo *rn.NotifRepository, authService *sa.AuthService, profileService *sp.ProfileService, groupRepo *rg.GroupRepository, chatServer *sc.ChatServer) *NotificationService {
	return &NotificationService{
		notifRepo:      notifRepo,
		authService:    authService,
		profileService: profileService,
		groupRepo:      groupRepo,
		chatServer:     chatServer,
	}
}

// ToggleAllSeenFalse sets the seen status of all notifications to true for a user.
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

// ToggleAllStaus updates the status of all notifications to the specified value.
func (NS *NotificationService) ToggleAllStaus(notifications []models.Notification, value, notifType string) *models.ErrorJson {
	for _, notification := range notifications {
		if errJson := NS.notifRepo.UpdateStatus(notification.Id, value); errJson != nil {
			return errJson
		}
	}
	return nil
}

// GetAllNotificationByType retrieves all notifications of a specific type for a user.
func (NS *NotificationService) GetAllNotificationByType(user_id, notifType string) ([]models.Notification, *models.ErrorJson) {
	all, err := NS.notifRepo.SelectAllNotificationByType(user_id, notifType)
	if err != nil {
		return nil, err
	}
	return all, nil
}

// DeleteService deletes a notification based on the provided parameters.
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

	if errJson := NS.broadcast(recieverId); errJson != nil {
		return errJson
	}
	return nil
}

// IsHasSeenFalse checks if the user has any notifications that are not seen.
func (NS *NotificationService) IsHasSeenFalse(userId string) (bool, *models.ErrorJson) {
	seen, errJson := NS.notifRepo.IsHasSeenFalse(userId)
	if errJson != nil {
		return false, errJson
	}
	return seen, nil
}

// broadcast sends a notification to the user about new notifications or no new notifications.
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
