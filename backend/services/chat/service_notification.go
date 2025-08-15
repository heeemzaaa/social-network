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
	if errJson := service.DeleteService(notification.TargetID, notification.SenderID, notification.Type, notification.GroupID); errJson != nil {
		return errJson
	}

	var errJson *models.ErrorJson

	switch notification.Type {
	case "follow-private":
		errJson = service.InsertNotification(notification, "later")
	case "follow-public":
		errJson = service.InsertNotification(notification, "accept")
	case "group-invitation":
		errJson = service.InsertNotification(notification, "later")
	case "group-join":
		errJson = service.InsertNotification(notification, "later")
	case "group-event":
		members, errJson := service.GetMembersOfGroup(notification.GroupID)
		if errJson != nil {
			return errJson
		}
		errorArray := []models.ErrorJson{}
		for _, userId := range members {
			if userId == notification.SenderID {
				continue
			}
			// SenderId:   event.EventCreator.Id,
			notification.TargetID = userId
			// Type:       "group-event",
			// GroupId:    event.Group.GroupId,
			// EventId:    event.EventId,
			// GroupName:  event.Group.Title,
			errJson = service.InsertNotification(notification, "none")
			if errJson != nil {
				errorArray = append(errorArray, errJson.PointErrorJson())
			}
		}
		fmt.Println("hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhheeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee")
		if len(errorArray) > 0 {
			fmt.Println("ERROR EVENT LOOP:", errorArray[0])
		}
		
	default:
		return &models.ErrorJson{
			Status: 400,
			Error:  "invalid notification type",
		}
	}

	if errJson != nil {
		return errJson
	}

	return nil
}

func (service *ChatService) InsertNotification(notification models.Notification, status string) *models.ErrorJson {
	notification.Status = status
	if errJson := service.repo.InsertNewNotification(notification); errJson != nil {
		return errJson
	}
	return nil
}
