package notification

import (
	"time"

	"social-network/backend/models"
	"social-network/backend/utils"
)

// insert new notification after event hapen
func (NS *NotificationService) PostService(data models.Notif) *models.ErrorJson {
	fullName, errJson := NS.repo2.GetUserFullNameById(data.SenderId)
	if errJson != nil {
		return errJson
	}
	data.SenderFullName = fullName

	switch data.Type {
	case "follow-private":
		errJson = NS.FollowPrivateProfile(data)
	case "follow-public":
		errJson = NS.FollowPublicProfile(data)
	case "group-invitation":
		errJson = NS.GroupInvitationRequest(data)
	case "group-join":
		errJson = NS.GroupJoinRequest(data)
	case "group-event":
		errJson = NS.GroupEventRequest(data)
	default:
		return models.NewErrorJson(400, "Bad Request - 400", "invalid type")
	}

	if errJson != nil {
		return errJson
	}

	return nil
}

// - follow private profile request
func (NS *NotificationService) FollowPrivateProfile(data models.Notif) *models.ErrorJson {
	if err := NS.repo.InsertNewNotification(models.Notification{
		Id:             utils.NewUUID(),
		SenderId:       data.SenderId,
		RecieverId:     data.RecieverId,
		GroupId:        "none",
		EventId:        "none",
		Type:           data.Type,
		SenderFullName: data.SenderFullName,
		GroupName:      "none",
		Status:         "later",
		Seen:           false,
		CreatedAt:      time.Now(),
	}); err != nil {
		return err
	}
	return nil
}

// - follow public profile request
func (NS *NotificationService) FollowPublicProfile(data models.Notif) *models.ErrorJson {
	///////////////////////////////////////////////////  golna madich nkhedmo 3la had l case //////////

	if err := NS.repo.InsertNewNotification(models.Notification{
		Id:             utils.NewUUID(),
		SenderId:       data.SenderId,
		RecieverId:     data.RecieverId,
		GroupId:        "none",
		EventId:        "none",
		Type:           data.Type,
		SenderFullName: data.SenderFullName,
		GroupName:      "none",
		Status:         "later",
		Seen:           false,
		CreatedAt:      time.Now(),
	}); err != nil {
		return err
	}
	return nil
}

// - group invitation request
func (NS *NotificationService) GroupInvitationRequest(data models.Notif) *models.ErrorJson {
	if errJson := NS.repo.InsertNewNotification(models.Notification{
		Id:             utils.NewUUID(),
		SenderId:       data.SenderId,
		RecieverId:     data.RecieverId,
		GroupId:        data.GroupId,
		EventId:        "none",
		Type:           data.Type,
		SenderFullName: data.SenderFullName,
		GroupName:      data.GroupName,
		Status:         "later",
		Seen:           false,
		CreatedAt:      time.Now(),
	}); errJson != nil {
		return errJson
	}
	return nil
}

// - group join request [admin]
func (NS *NotificationService) GroupJoinRequest(data models.Notif) *models.ErrorJson {
	if errJson := NS.repo.InsertNewNotification(models.Notification{
		Id:             utils.NewUUID(),
		SenderId:       data.SenderId,
		RecieverId:     data.RecieverId,
		GroupId:        data.GroupId,
		EventId:        "none",
		Type:           data.Type,
		SenderFullName: data.SenderFullName,
		GroupName:      data.GroupName,
		Status:         "later",
		Seen:           false,
		CreatedAt:      time.Now(),
	}); errJson != nil {
		return errJson
	}
	return nil
}

// - group event created [group-members]
func (NS *NotificationService) GroupEventRequest(data models.Notif) *models.ErrorJson {
	if err := NS.repo.InsertNewNotification(models.Notification{
		Id:             utils.NewUUID(),
		SenderId:       data.SenderId,
		RecieverId:     data.RecieverId,
		GroupId:        data.GroupId,
		EventId:        data.EventId,
		Type:           data.Type,
		SenderFullName: data.SenderFullName,
		GroupName:      data.GroupName,
		Status:         "later",
		Seen:           false,
		CreatedAt:      time.Now(),
	}); err != nil {
		return err
	}
	return nil
}
