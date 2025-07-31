package notification

import (
	"fmt"
	"social-network/backend/models"
	"social-network/backend/utils"
	"time"
)

// Insert new notification after
func (NS *NotificationService) PostService(data *models.Notif) *models.ErrorJson {
	if errJson := NS.DeleteService(data.RecieverId, data.SenderId, data.Type, data.GroupId); errJson != nil {
		return errJson
	}

	fullName, errJson := NS.authRepo.GetUserFullNameById(data.SenderId)
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
	return nil
}

// - follow private profile request
func (NS *NotificationService) FollowPrivateProfile(notification models.Notification, data *models.Notif) *models.ErrorJson {

	notification.GroupId = "none"
	notification.EventId = "none"
	notification.GroupName = "none"
	notification.Status = "later"

	if errJson := NS.notifRepo.InsertNewNotification(notification); errJson != nil {
		fmt.Println("error private insertion ---------> ", errJson)
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
		fmt.Println("error public insertion ---------> ", errJson)
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
		fmt.Println("error invitation insertion ---------> ", errJson)
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
		fmt.Println("error join insertion ---------> ", errJson)
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
		fmt.Println("error event insertion ---------> ", errJson)
		return errJson
	}
	return nil
}
