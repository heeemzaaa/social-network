package notification

import (
	"fmt"
	"log"
	"social-network/backend/models"
	rg "social-network/backend/repositories/group"
	"social-network/backend/repositories/notification"
	"social-network/backend/repositories/profile"
	// GS "social-network/backend/service/group"
)

type NotificationServiceUpdate struct {
	repo  *notification.NotifRepository
	rProfile *profile.ProfileRepository
	rGroup   *rg.GroupRepository
}

func NewNotifServiceUpdate(repo *notification.NotifRepository, rProfile *profile.ProfileRepository, rGroup *rg.GroupRepository) *NotificationServiceUpdate {
	return &NotificationServiceUpdate{
		repo:  repo,
		rProfile: rProfile,
		rGroup:    rGroup,
	}
}

func (NUS *NotificationServiceUpdate) UpdateService(data models.Unotif, user_id string) *models.ErrorJson {
	log.Println("START UPDATE SERVICE ----- REQUEST DATA = ", data)
	fmt.Println("dataaaa", data.NotifId)

	notification, errJson := NUS.repo.SelectNotification(data.NotifId)
	if errJson != nil {
		return models.NewErrorJson(errJson.Status, errJson.Error, errJson.Message)
	}

	if user_id != notification.RecieverId {
		return &models.ErrorJson{Status: 403, Error: "ERROR 403 Acces Forbidden", Message: "Invalid----Operation"}
	}
	if notification.Status != "later" {
		return models.NewErrorJson(400, "Bad-Request 400", "Invalid----Operation")
	}

	switch data.Type {
	case "follow-private":
		errJson = NUS.UpdateFollowPrivateProfile(data, notification)
	case "follow-public":
		// errJson = NUS.UpdateFollowPublicProfile(data, notification)
	case "group-invitation":
		errJson = NUS.UpdateGroupInvitationRequest(data, notification)
	case "group-join":
		errJson = NUS.UpdateGroupJoinRequest(data, notification)
	case "group-event":
		errJson = NUS.UpdateGroupEventRequest(data, notification)
	default:
		return models.NewErrorJson(400, "Bad Request - 400", "invalid type")
	}
	if errJson != nil {
		return errJson
	}

	if errJson = NUS.repo.UpdateStatusById(notification.Id, data.Status); errJson != nil {
		return errJson
	}
	if errJson := NUS.repo.UpdateSeen(notification.Id); errJson != nil {
		return errJson
	}

	// should be return notification updated
	return nil
}

func (NUS *NotificationServiceUpdate) UpdateFollowPrivateProfile(data models.Unotif, notification models.Notification) *models.ErrorJson {

	switch data.Status {
	case "accept":
		err := NUS.rProfile.AcceptedRequest(notification.RecieverId, notification.SenderId)
		if err != nil {
			return models.NewErrorJson(500, "500 - cannot accept request", err)
		}
	case "reject":
		err := NUS.rProfile.RejectedRequest(notification.RecieverId, notification.SenderId)
		if err != nil {
			return models.NewErrorJson(500, "500 - cannot reject request", err)
		}
	default:
		return models.NewErrorJson(400, "400 - Bad Request", "Invalid Status:"+data.Status)
	}
	return nil
}

func (NUS *NotificationServiceUpdate) UpdateGroupJoinRequest(data models.Unotif, notification models.Notification) *models.ErrorJson {
	switch data.Status {
	case "accept":
		err := NUS.rGroup.Approve(notification.GroupId, notification.SenderId)
		if err != nil {
			return models.NewErrorJson(500, "500 - cannot accept request", err)
		}
	case "reject":
		err := NUS.rGroup.Decline(notification.GroupId, notification.SenderId)
		if err != nil {
			return models.NewErrorJson(500, "500 - cannot decline request", err)
		}
	default:
		return models.NewErrorJson(400, "Bad-Request 400", "Invalid----Status")
	}
	return nil
}

func (NUS *NotificationServiceUpdate) UpdateGroupInvitationRequest(data models.Unotif, notification models.Notification) *models.ErrorJson {
	switch data.Status {
	case "accept":
		err := NUS.rGroup.Accept(notification.SenderId, notification.GroupId, notification.RecieverId)
		if err != nil {
			return models.NewErrorJson(500, "500 - cannot accept request", err)
		}
	case "reject":
		err := NUS.rGroup.Reject(notification.SenderId, notification.GroupId, notification.RecieverId)
		if err != nil {
			return models.NewErrorJson(500, "500 - cannot accept request", err)
		}
	default:
		return models.NewErrorJson(400, "Bad-Request", "Invalid----Status")
	}
	return nil
}

func (NUS *NotificationServiceUpdate) UpdateGroupEventRequest(data models.Unotif, notification models.Notification) *models.ErrorJson {
	switch data.Status {
	case "accept":
		err := NUS.rGroup.Accept(notification.SenderId, notification.GroupId, notification.RecieverId)
		if err != nil {
			return models.NewErrorJson(500, "500 - cannot accept request", err)
		}
	case "reject":
		err := NUS.rGroup.Reject(notification.SenderId, notification.GroupId, notification.RecieverId)
		if err != nil {
			return models.NewErrorJson(500, "500 - cannot accept request", err)
		}
	default:
		return models.NewErrorJson(400, "Bad-Request 400", "Invalid----Status")
	}
	return nil
}
