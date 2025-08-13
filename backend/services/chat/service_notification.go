package chat

import (
	"strings"
	"time"

	"social-network/backend/models"
	"social-network/backend/utils"
)

// DeleteService deletes a notification based on the provided parameters.
func (service *ChatService) DeleteService(recieverId, senderId, notifType, groupId string) *models.ErrorJson {
	if strings.HasPrefix(notifType, "follow") {
		if errJson := service.repo.DeleteFollowNotification(senderId, recieverId, notifType); errJson != nil {
			return errJson
		}
	} else if notifType != "group-event" {
		if errJson := service.repo.DeleteGroupNotification(senderId, recieverId, notifType, groupId); errJson != nil {
			return errJson
		}
	}

	return nil
}

// PostService handles the creation of a new notification based on the provided data.
func (service *ChatService) PostService(data *models.Notif) *models.ErrorJson {
	if errJson := service.DeleteService(data.RecieverId, data.SenderId, data.Type, data.GroupId); errJson != nil {
		return errJson
	}

	// fullName, errJson := service.authService.GetUserFullName(data.SenderId)
	// if errJson != nil {
	// 	return errJson
	// }

	notification := models.Notification{
		Id: utils.NewUUID(),

		SenderId:   data.SenderId,
		RecieverId: data.RecieverId,

		Type: data.Type,
		Seen: false,

		// SenderFullName: fullName,

		CreatedAt: time.Now(),
	}
	var errJson *models.ErrorJson
	switch data.Type {
	case "follow-private":
		errJson = service.FollowPrivateProfile(notification, data)
	case "follow-public":
		errJson = service.FollowPublicProfile(notification, data)
	case "group-invitation":
		errJson = service.GroupInvitationRequest(notification, data)
	case "group-join":
		errJson = service.GroupJoinRequest(notification, data)
	case "group-event":
		errJson = service.GroupEventRequest(notification, data)
	default:
		return models.NewErrorJson(400, "400 - Bad Request", "invalid notification type")
	}

	if errJson != nil {
		return errJson
	}

	return nil
}

// - follow private profile request
func (service *ChatService) FollowPrivateProfile(notification models.Notification, data *models.Notif) *models.ErrorJson {
	notification.GroupId = "none"
	notification.EventId = "none"
	notification.GroupName = "none"
	notification.Status = "later"

	if errJson := service.repo.InsertNewNotification(notification); errJson != nil {
		return errJson
	}
	return nil
}

// - follow public profile request
func (service *ChatService) FollowPublicProfile(notification models.Notification, data *models.Notif) *models.ErrorJson {
	notification.GroupId = "none"
	notification.EventId = "none"
	notification.GroupName = "none"
	notification.Status = "accept"

	if errJson := service.repo.InsertNewNotification(notification); errJson != nil {
		return errJson
	}
	return nil
}

// - group invitation request
func (service *ChatService) GroupInvitationRequest(notification models.Notification, data *models.Notif) *models.ErrorJson {
	notification.GroupId = data.GroupId
	notification.EventId = "none"
	notification.GroupName = data.GroupName
	notification.Status = "later"

	if errJson := service.repo.InsertNewNotification(notification); errJson != nil {
		return errJson
	}
	return nil
}

// - group join request [admin]
func (service *ChatService) GroupJoinRequest(notification models.Notification, data *models.Notif) *models.ErrorJson {
	notification.GroupId = data.GroupId
	notification.EventId = "none"
	notification.GroupName = data.GroupName
	notification.Status = "later"

	if errJson := service.repo.InsertNewNotification(notification); errJson != nil {
		return errJson
	}
	return nil
}

// - group event created [group-members]
func (service *ChatService) GroupEventRequest(notification models.Notification, data *models.Notif) *models.ErrorJson {
	notification.GroupId = data.GroupId
	notification.EventId = data.EventId
	notification.GroupName = data.GroupName
	notification.Status = "none"

	if errJson := service.repo.InsertNewNotification(notification); errJson != nil {
		return errJson
	}
	return nil
}
