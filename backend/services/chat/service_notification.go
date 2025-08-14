package chat

import (
	"fmt"
	"strings"

	"social-network/backend/models"
)

// DeleteService deletes a notification based on the provided parameters.
func (service *ChatService) DeleteService(targetId, senderId, notifType, groupId string) *models.ErrorJson {
	if strings.HasPrefix(notifType, "follow") {
		if errJson := service.repo.DeleteFollowNotification(senderId, targetId, notifType); errJson != nil {
			return errJson
		}
	} else if notifType != "group-event" {
		if errJson := service.repo.DeleteGroupNotification(senderId, targetId, notifType, groupId); errJson != nil {
			return errJson
		}
	}
	return nil
}

// PostService handles the creation of a new notification based on the provided data.
func (service *ChatService) PostService(notification models.Notification) *models.ErrorJson {
	// if errJson := service.DeleteService(data.TargetID, data.SenderID, data.Type, data.GroupId); errJson != nil {
	// 	return errJson
	// }						for delete duplicate notification if exist /////

	var errJson *models.ErrorJson

	fmt.Println("PostService - notification:", notification)

	switch notification.Type {
	case "follow-private":
		errJson = service.FollowPrivateProfile(notification)
	case "follow-public":
		errJson = service.FollowPublicProfile(notification)
	case "group-invitation":
		errJson = service.GroupInvitationRequest(notification)
	case "group-join":
		errJson = service.GroupJoinRequest(notification)
	case "group-event":
		errJson = service.GroupEventRequest(notification)
	default:
		return models.NewErrorJson(400, "400 - Bad Request", "invalid notification type")
	}

	if errJson != nil {
		return errJson
	}

	return nil
}

// - follow private profile request
func (service *ChatService) FollowPrivateProfile(notification models.Notification) *models.ErrorJson {
	// notification.GroupId = "none"
	// notification.EventID = "none"
	// notification.GroupName = "none"
	// notification.Status = "later"

	if errJson := service.repo.InsertNewNotification(notification); errJson != nil {
		return errJson
	}
	return nil
}

// - follow public profile request
func (service *ChatService) FollowPublicProfile(notification models.Notification) *models.ErrorJson {
	// notification.GroupId = "none"
	// notification.EventID = "none"
	// notification.GroupName = "none"
	notification.Status = "accept"

	if errJson := service.repo.InsertNewNotification(notification); errJson != nil {
		return errJson
	}
	return nil
}

// - group invitation request
func (service *ChatService) GroupInvitationRequest(notification models.Notification) *models.ErrorJson {
	// notification.GroupId = data.GroupId
	// notification.EventID = "none"
	// notification.GroupName = data.GroupName
	notification.Status = "later"

	if errJson := service.repo.InsertNewNotification(notification); errJson != nil {
		return errJson
	}
	return nil
}

// - group join request [admin]
func (service *ChatService) GroupJoinRequest(notification models.Notification) *models.ErrorJson {
	// notification.GroupId = data.GroupId
	// notification.EventID = "none"
	// notification.GroupName = data.GroupName
	notification.Status = "later"

	if errJson := service.repo.InsertNewNotification(notification); errJson != nil {
		return errJson
	}
	return nil
}

// - group event created [group-members]
func (service *ChatService) GroupEventRequest(notification models.Notification) *models.ErrorJson {
	// notification.GroupId = data.GroupId
	// notification.EventID = data.EventId
	// notification.GroupName = data.GroupName
	notification.Status = "none"

	if errJson := service.repo.InsertNewNotification(notification); errJson != nil {
		return errJson
	}
	return nil
}
