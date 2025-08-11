package notification

import (
	"time"

	"social-network/backend/models"
	"social-network/backend/utils"
)

// PostService handles the creation of a new notification based on the provided data.
func (NS *NotificationService) PostService(data *models.Notif) *models.ErrorJson {
	if errJson := NS.DeleteService(data.RecieverId, data.SenderId, data.Type, data.GroupId); errJson != nil {
		return errJson
	}

	fullName, errJson := NS.authService.GetUserFullName(data.SenderId)
	if errJson != nil {
		return errJson
	}

	notification := models.Notification{
		Id: utils.NewUUID(),

		SenderId:   data.SenderId,
		RecieverId: data.RecieverId,

		Type: data.Type,
		Seen: false,

		SenderFullName: fullName,

		CreatedAt: time.Now(),
	}

	switch data.Type {
	case "follow-private":
		errJson = NS.FollowPrivateProfile(notification, data)
	case "follow-public":
		errJson = NS.FollowPublicProfile(notification, data)
	case "group-invitation":
		errJson = NS.GroupInvitationRequest(notification, data)
	case "group-join":
		errJson = NS.GroupJoinRequest(notification, data)
	case "group-event":
		errJson = NS.GroupEventRequest(notification, data)
	default:
		return models.NewErrorJson(400, "400 - Bad Request", "invalid notification type")
	}

	if errJson != nil {
		return errJson
	}

	if errJson := NS.broadcast(notification.RecieverId); errJson != nil {
		return errJson
	}
	return nil
}

// - follow private profile request
func (NS *NotificationService) FollowPrivateProfile(notification models.Notification, data *models.Notif) *models.ErrorJson {

	notification.GroupId = "none"
	notification.EventId = "none"
	notification.GroupName = "none"
	notification.Status = "later"

	if errJson := NS.notifRepo.InsertNewNotification(notification); errJson != nil {
		return errJson
	}
	return nil
}

// - follow public profile request
func (NS *NotificationService) FollowPublicProfile(notification models.Notification, data *models.Notif) *models.ErrorJson {

	notification.GroupId = "none"
	notification.EventId = "none"
	notification.GroupName = "none"
	notification.Status = "accept"

	if errJson := NS.notifRepo.InsertNewNotification(notification); errJson != nil {
		return errJson
	}
	return nil
}

// - group invitation request
func (NS *NotificationService) GroupInvitationRequest(notification models.Notification, data *models.Notif) *models.ErrorJson {

	notification.GroupId = data.GroupId
	notification.EventId = "none"
	notification.GroupName = data.GroupName
	notification.Status = "later"

	if errJson := NS.notifRepo.InsertNewNotification(notification); errJson != nil {
		return errJson
	}
	return nil
}

// - group join request [admin]
func (NS *NotificationService) GroupJoinRequest(notification models.Notification, data *models.Notif) *models.ErrorJson {

	notification.GroupId = data.GroupId
	notification.EventId = "none"
	notification.GroupName = data.GroupName
	notification.Status = "later"

	if errJson := NS.notifRepo.InsertNewNotification(notification); errJson != nil {
		return errJson
	}
	return nil
}

// - group event created [group-members]
func (NS *NotificationService) GroupEventRequest(notification models.Notification, data *models.Notif) *models.ErrorJson {

	notification.GroupId = data.GroupId
	notification.EventId = data.EventId
	notification.GroupName = data.GroupName
	notification.Status = "none"

	if errJson := NS.notifRepo.InsertNewNotification(notification); errJson != nil {
		return errJson
	}
	return nil
}
